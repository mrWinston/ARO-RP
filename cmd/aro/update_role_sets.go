package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/metrics/statsd"
	"github.com/Azure/ARO-RP/pkg/util/encryption"
	"github.com/Azure/ARO-RP/pkg/util/keyvault"
)

// 1 - Get env data from agent VMs (with getEnvironemntData) and write to types created in step 1
func getPlatformWorkloadIdentityRoleSet() (*api.PlatformWorkloadIdentityRoleSet, error) {
	const envKey = envPlatformWorkloadIdentityRoleSets
	var PlatformWorkloadIdentityRoleSet api.PlatformWorkloadIdentityRoleSet

	// marshall env data into type api.PlatformWorkloadIdentityRoleSet
	if err := getEnvironmentData(envKey, &PlatformWorkloadIdentityRoleSet); err != nil {
		return nil, err
	}

	return &PlatformWorkloadIdentityRoleSet, nil
}

func getRoleSetFromEnv() ([]api.PlatformWorkloadIdentityRoleSet, error) {
	roleSet, err := getPlatformWorkloadIdentityRoleSet()
	if err != nil {
		return nil, err
	}

	finalRoleSet := []api.PlatformWorkloadIdentityRoleSet{}
	finalRoleSet = append(finalRoleSet, *roleSet)

	return finalRoleSet, nil
}

// 2 - Get the existing role set documents, if existing
// Mostly copied from update_ocp_versions.go
func getPlatformWorkloadIdentityRoleSetDatabase(ctx context.Context, log *logrus.Entry) (database.PlatformWorkloadIdentityRoleSets, error) {
	_env, err := env.NewCore(ctx, log, env.COMPONENT_UPDATE_OCP_VERSIONS)
	if err != nil {
		return nil, err
	}

	msiToken, err := _env.NewMSITokenCredential()
	if err != nil {
		return nil, fmt.Errorf("MSI Authorizer failed with: %s", err.Error())
	}

	msiKVAuthorizer, err := _env.NewMSIAuthorizer(_env.Environment().KeyVaultScope)
	if err != nil {
		return nil, fmt.Errorf("MSI KeyVault Authorizer failed with: %s", err.Error())
	}

	m := statsd.New(ctx, log.WithField("component", "update-role-sets"), _env, os.Getenv("MDM_ACCOUNT"), os.Getenv("MDM_NAMESPACE"), os.Getenv("MDM_STATSD_SOCKET"))

	keyVaultPrefix := os.Getenv(envKeyVaultPrefix)
	serviceKeyvaultURI := keyvault.URI(_env, env.ServiceKeyvaultSuffix, keyVaultPrefix)
	serviceKeyvault := keyvault.NewManager(msiKVAuthorizer, serviceKeyvaultURI)

	aead, err := encryption.NewMulti(ctx, serviceKeyvault, env.EncryptionSecretV2Name, env.EncryptionSecretName)
	if err != nil {
		return nil, err
	}

	if err := env.ValidateVars(envDatabaseAccountName); err != nil {
		return nil, err
	}
	dbAccountName := os.Getenv(envDatabaseAccountName)
	clientOptions := &policy.ClientOptions{
		ClientOptions: _env.Environment().ManagedIdentityCredentialOptions().ClientOptions,
	}

	logrusEntry := log.WithField("component", "database")
	dbAuthorizer, err := database.NewMasterKeyAuthorizer(ctx, logrusEntry, msiToken, clientOptions, _env.SubscriptionID(), _env.ResourceGroup(), dbAccountName)
	if err != nil {
		return nil, err
	}

	dbc, err := database.NewDatabaseClient(log.WithField("component", "database"), _env, dbAuthorizer, m, aead, dbAccountName)
	if err != nil {
		return nil, err
	}

	dbName, err := DBName(_env.IsLocalDevelopmentMode())
	if err != nil {
		return nil, err
	}
	dbPlatformWorkloadIdentityRoleSetsDocument, err := database.NewPlatformWorkloadIdentityRoleSets(ctx, dbc, dbName)
	if err != nil {
		return nil, err
	}

	return dbPlatformWorkloadIdentityRoleSetsDocument, nil
}

// 3 - Put/patch the new role sets to the doc, overwriting whatever is there for that version, or adding if new
// Mostly copied from update_ocp_versions.go
func updatePlatformWorkloadIdentityRoleSetsInCosmosDB(ctx context.Context, dbPlatformWorkloadIdentityRoleSets database.PlatformWorkloadIdentityRoleSets, log *logrus.Entry) error {
	dbPlatformWorkloadIdentityRoleSet, err := dbPlatformWorkloadIdentityRoleSets.ListAll(ctx)
	if err != nil {
		return nil
	}

	incomingRoleSet, err := getRoleSetFromEnv()
	if err != nil {
		return err
	}

	newRoleSets := make(map[string]api.PlatformWorkloadIdentityRoleSet)
	for _, doc := range incomingRoleSet {
		newRoleSets[doc.Properties.OpenShiftVersion] = doc
	}

	for _, doc := range dbPlatformWorkloadIdentityRoleSet.PlatformWorkloadIdentityRoleSetDocuments {
		existing, found := newRoleSets[doc.PlatformWorkloadIdentityRoleSet.Properties.OpenShiftVersion]
		if found {
			log.Printf("Found Version %q, patching", existing.Properties.OpenShiftVersion)
			_, err := dbPlatformWorkloadIdentityRoleSets.Patch(ctx, doc.ID, func(inFlightDoc *api.PlatformWorkloadIdentityRoleSetDocument) error {
				inFlightDoc.PlatformWorkloadIdentityRoleSet = &existing
				return nil
			})
			if err != nil {
				return err
			}
			log.Printf("Version %q found", existing.Properties.OpenShiftVersion)
			delete(newRoleSets, existing.Properties.OpenShiftVersion)
			continue
		}

		log.Printf("Version %q not found, deleting", doc.PlatformWorkloadIdentityRoleSet.Properties.OpenShiftVersion)
		// Delete via changefeed
		_, err := dbPlatformWorkloadIdentityRoleSets.Patch(ctx, doc.ID,
			func(d *api.PlatformWorkloadIdentityRoleSetDocument) error {
				d.PlatformWorkloadIdentityRoleSet.Deleting = true
				d.TTL = 60
				return nil
			})
		if err != nil {
			return err
		}
	}

	for _, doc := range newRoleSets {
		log.Printf("Version %q not found in database, creating", doc.Properties.OpenShiftVersion)
		newDoc := api.PlatformWorkloadIdentityRoleSetDocument{
			ID:                              dbPlatformWorkloadIdentityRoleSets.NewUUID(),
			PlatformWorkloadIdentityRoleSet: &doc,
		}
		_, err := dbPlatformWorkloadIdentityRoleSets.Create(ctx, &newDoc)
		if err != nil {
			return err
		}
	}

	return nil
}

func updatePlatformWorkloadIdentityRoleSets(ctx context.Context, log *logrus.Entry) error {
	if err := env.ValidateVars("PLATFORM_WORKLOAD_IDENTITY_ROLE_SETS"); err != nil {
		return err
	}

	if !env.IsLocalDevelopmentMode() {
		if err := env.ValidateVars("MDM_ACCOUNT", "MDM_NAMESPACE"); err != nil {
			return err
		}
	}

	dbRoleSets, err := getPlatformWorkloadIdentityRoleSetDatabase(ctx, log)
	if err != nil {
		return err
	}

	err = updatePlatformWorkloadIdentityRoleSetsInCosmosDB(ctx, dbRoleSets, log)
	if err != nil {
		return err
	}

	return nil
}