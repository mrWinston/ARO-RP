package item

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"errors"
)

// Provides operations to manage the collection of servicePrincipal entities.
type GetExpandQueryParameterType int

const (
	ASTERISK_GETEXPANDQUERYPARAMETERTYPE GetExpandQueryParameterType = iota
	APPMANAGEMENTPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	APPROLEASSIGNEDTO_GETEXPANDQUERYPARAMETERTYPE
	APPROLEASSIGNMENTS_GETEXPANDQUERYPARAMETERTYPE
	CLAIMSMAPPINGPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	CREATEDOBJECTS_GETEXPANDQUERYPARAMETERTYPE
	DELEGATEDPERMISSIONCLASSIFICATIONS_GETEXPANDQUERYPARAMETERTYPE
	ENDPOINTS_GETEXPANDQUERYPARAMETERTYPE
	FEDERATEDIDENTITYCREDENTIALS_GETEXPANDQUERYPARAMETERTYPE
	HOMEREALMDISCOVERYPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	MEMBEROF_GETEXPANDQUERYPARAMETERTYPE
	OAUTH2PERMISSIONGRANTS_GETEXPANDQUERYPARAMETERTYPE
	OWNEDOBJECTS_GETEXPANDQUERYPARAMETERTYPE
	OWNERS_GETEXPANDQUERYPARAMETERTYPE
	TOKENISSUANCEPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	TOKENLIFETIMEPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	TRANSITIVEMEMBEROF_GETEXPANDQUERYPARAMETERTYPE
	SYNCHRONIZATION_GETEXPANDQUERYPARAMETERTYPE
)

func (i GetExpandQueryParameterType) String() string {
	return []string{"*", "appManagementPolicies", "appRoleAssignedTo", "appRoleAssignments", "claimsMappingPolicies", "createdObjects", "delegatedPermissionClassifications", "endpoints", "federatedIdentityCredentials", "homeRealmDiscoveryPolicies", "memberOf", "oauth2PermissionGrants", "ownedObjects", "owners", "tokenIssuancePolicies", "tokenLifetimePolicies", "transitiveMemberOf", "synchronization"}[i]
}
func ParseGetExpandQueryParameterType(v string) (any, error) {
	result := ASTERISK_GETEXPANDQUERYPARAMETERTYPE
	switch v {
	case "*":
		result = ASTERISK_GETEXPANDQUERYPARAMETERTYPE
	case "appManagementPolicies":
		result = APPMANAGEMENTPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	case "appRoleAssignedTo":
		result = APPROLEASSIGNEDTO_GETEXPANDQUERYPARAMETERTYPE
	case "appRoleAssignments":
		result = APPROLEASSIGNMENTS_GETEXPANDQUERYPARAMETERTYPE
	case "claimsMappingPolicies":
		result = CLAIMSMAPPINGPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	case "createdObjects":
		result = CREATEDOBJECTS_GETEXPANDQUERYPARAMETERTYPE
	case "delegatedPermissionClassifications":
		result = DELEGATEDPERMISSIONCLASSIFICATIONS_GETEXPANDQUERYPARAMETERTYPE
	case "endpoints":
		result = ENDPOINTS_GETEXPANDQUERYPARAMETERTYPE
	case "federatedIdentityCredentials":
		result = FEDERATEDIDENTITYCREDENTIALS_GETEXPANDQUERYPARAMETERTYPE
	case "homeRealmDiscoveryPolicies":
		result = HOMEREALMDISCOVERYPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	case "memberOf":
		result = MEMBEROF_GETEXPANDQUERYPARAMETERTYPE
	case "oauth2PermissionGrants":
		result = OAUTH2PERMISSIONGRANTS_GETEXPANDQUERYPARAMETERTYPE
	case "ownedObjects":
		result = OWNEDOBJECTS_GETEXPANDQUERYPARAMETERTYPE
	case "owners":
		result = OWNERS_GETEXPANDQUERYPARAMETERTYPE
	case "tokenIssuancePolicies":
		result = TOKENISSUANCEPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	case "tokenLifetimePolicies":
		result = TOKENLIFETIMEPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	case "transitiveMemberOf":
		result = TRANSITIVEMEMBEROF_GETEXPANDQUERYPARAMETERTYPE
	case "synchronization":
		result = SYNCHRONIZATION_GETEXPANDQUERYPARAMETERTYPE
	default:
		return 0, errors.New("Unknown GetExpandQueryParameterType value: " + v)
	}
	return &result, nil
}
func SerializeGetExpandQueryParameterType(values []GetExpandQueryParameterType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i GetExpandQueryParameterType) isMultiValue() bool {
	return false
}
