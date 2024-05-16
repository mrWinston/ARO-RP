package armnetwork

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"

	"github.com/Azure/ARO-RP/pkg/util/azureclient"
	"github.com/Azure/ARO-RP/pkg/util/azureclient/azuresdk/azcore"
	"github.com/Azure/ARO-RP/pkg/util/azureclient/azuresdk/common"
)

// SecurityGroupsClient is a minimal interface for azure SecurityGroupsClient
type SecurityGroupsClient interface {
	Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *armnetwork.SecurityGroupsClientGetOptions) (armnetwork.SecurityGroupsClientGetResponse, error)
	SecurityGroupsClientAddons
}

type securityGroupsClient struct {
	*armnetwork.SecurityGroupsClient
}

// NewSecurityGroupsClient creates a new SecurityGroupsClient
func NewSecurityGroupsClient(environment *azureclient.AROEnvironment, subscriptionID string, credential azcore.TokenCredential) (SecurityGroupsClient, error) {
	options := common.ClientOptions
	options.Cloud = environment.Cloud
	clientFactory, err := armnetwork.NewClientFactory(subscriptionID, credential, &options)

	if err != nil {
		return nil, err
	}
	return &securityGroupsClient{clientFactory.NewSecurityGroupsClient()}, nil
}
