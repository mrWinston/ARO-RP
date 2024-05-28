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

// SubnetsClient is a minimal interface for azure-sdk-for-go subnets client
type SubnetsClient interface {
	Get(ctx context.Context, resourceGroupName, virtualNetworkName, subnetName string, options *armnetwork.SubnetsClientGetOptions) (armnetwork.SubnetsClientGetResponse, error)
	SubnetsClientAddons
}

type subnetsClient struct {
	*armnetwork.SubnetsClient
}

func NewSubnetsClient(environment *azureclient.AROEnvironment, subscriptionID string, credential azcore.TokenCredential) (SubnetsClient, error) {
	options := common.ClientOptions
	options.Cloud = environment.Cloud
	clientFactory, err := armnetwork.NewClientFactory(subscriptionID, credential, &options)

	if err != nil {
		return nil, err
	}
	return &subnetsClient{clientFactory.NewSubnetsClient()}, err
}
