package service

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/database"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/metrics"
	"github.com/Azure/ARO-RP/pkg/util/azureclient/azuresdk/azcore"
	"github.com/Azure/ARO-RP/pkg/util/encryption"
	"github.com/Azure/ARO-RP/pkg/util/keyvault"
)

func NewDatabaseClientUsingMasterKey(ctx context.Context, _env env.Core, log *logrus.Entry, m metrics.Emitter, msiToken azcore.TokenCredential, aead encryption.AEAD) (cosmosdb.DatabaseClient, error) {
	dbAccountName := os.Getenv(DatabaseAccountName)

	clientOptions := &policy.ClientOptions{
		ClientOptions: _env.Environment().ManagedIdentityCredentialOptions().ClientOptions,
	}

	dbAuthorizer, err := database.NewMasterKeyAuthorizer(ctx, _env.Logger(), msiToken, clientOptions, _env.SubscriptionID(), _env.ResourceGroup(), dbAccountName)
	if err != nil {
		return nil, err
	}

	dbc, err := database.NewDatabaseClient(
		log.WithField("component", "database"),
		_env,
		dbAuthorizer,
		m,
		aead,
		dbAccountName,
	)
	if err != nil {
		return nil, err
	}
	return dbc, nil
}

func NewDatabase(ctx context.Context, _env env.Core, log *logrus.Entry, m metrics.Emitter, withAEAD bool) (cosmosdb.DatabaseClient, error) {
	var aead encryption.AEAD

	if withAEAD {
		msiKVAuthorizer, err := _env.NewMSIAuthorizer(_env.Environment().KeyVaultScope)
		if err != nil {
			return nil, err
		}

		keyVaultPrefix := os.Getenv(KeyVaultPrefix)
		// TODO: should not be using the service keyvault here
		serviceKeyvaultURI := keyvault.URI(_env, env.ServiceKeyvaultSuffix, keyVaultPrefix)
		serviceKeyvault := keyvault.NewManager(msiKVAuthorizer, serviceKeyvaultURI)

		aead, err = encryption.NewMulti(ctx, serviceKeyvault, env.EncryptionSecretV2Name, env.EncryptionSecretName)
		if err != nil {
			return nil, err
		}
	}
	msiToken, err := _env.NewMSITokenCredential()
	if err != nil {
		return nil, err
	}

	return NewDatabaseClientUsingMasterKey(ctx, _env, _env.Logger(), m, msiToken, aead)
}
