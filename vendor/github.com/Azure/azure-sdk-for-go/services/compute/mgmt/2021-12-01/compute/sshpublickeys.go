package compute

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SSHPublicKeysClient is the compute Client
type SSHPublicKeysClient struct {
	BaseClient
}

// NewSSHPublicKeysClient creates an instance of the SSHPublicKeysClient client.
func NewSSHPublicKeysClient(subscriptionID string) SSHPublicKeysClient {
	return NewSSHPublicKeysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSSHPublicKeysClientWithBaseURI creates an instance of the SSHPublicKeysClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSSHPublicKeysClientWithBaseURI(baseURI string, subscriptionID string) SSHPublicKeysClient {
	return SSHPublicKeysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a new SSH public key resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// SSHPublicKeyName - the name of the SSH public key.
// parameters - parameters supplied to create the SSH public key.
func (client SSHPublicKeysClient) Create(ctx context.Context, resourceGroupName string, SSHPublicKeyName string, parameters SSHPublicKeyResource) (result SSHPublicKeyResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SSHPublicKeysClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroupName, SSHPublicKeyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SSHPublicKeysClient) CreatePreparer(ctx context.Context, resourceGroupName string, SSHPublicKeyName string, parameters SSHPublicKeyResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sshPublicKeyName":  autorest.Encode("path", SSHPublicKeyName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SSHPublicKeysClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SSHPublicKeysClient) CreateResponder(resp *http.Response) (result SSHPublicKeyResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an SSH public key.
// Parameters:
// resourceGroupName - the name of the resource group.
// SSHPublicKeyName - the name of the SSH public key.
func (client SSHPublicKeysClient) Delete(ctx context.Context, resourceGroupName string, SSHPublicKeyName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SSHPublicKeysClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, SSHPublicKeyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SSHPublicKeysClient) DeletePreparer(ctx context.Context, resourceGroupName string, SSHPublicKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sshPublicKeyName":  autorest.Encode("path", SSHPublicKeyName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SSHPublicKeysClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SSHPublicKeysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GenerateKeyPair generates and returns a public/private key pair and populates the SSH public key resource with the
// public key. The length of the key will be 3072 bits. This operation can only be performed once per SSH public key
// resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// SSHPublicKeyName - the name of the SSH public key.
func (client SSHPublicKeysClient) GenerateKeyPair(ctx context.Context, resourceGroupName string, SSHPublicKeyName string) (result SSHPublicKeyGenerateKeyPairResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SSHPublicKeysClient.GenerateKeyPair")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GenerateKeyPairPreparer(ctx, resourceGroupName, SSHPublicKeyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "GenerateKeyPair", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateKeyPairSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "GenerateKeyPair", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateKeyPairResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "GenerateKeyPair", resp, "Failure responding to request")
		return
	}

	return
}

// GenerateKeyPairPreparer prepares the GenerateKeyPair request.
func (client SSHPublicKeysClient) GenerateKeyPairPreparer(ctx context.Context, resourceGroupName string, SSHPublicKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sshPublicKeyName":  autorest.Encode("path", SSHPublicKeyName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}/generateKeyPair", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GenerateKeyPairSender sends the GenerateKeyPair request. The method will close the
// http.Response Body if it receives an error.
func (client SSHPublicKeysClient) GenerateKeyPairSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GenerateKeyPairResponder handles the response to the GenerateKeyPair request. The method always
// closes the http.Response Body.
func (client SSHPublicKeysClient) GenerateKeyPairResponder(resp *http.Response) (result SSHPublicKeyGenerateKeyPairResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get retrieves information about an SSH public key.
// Parameters:
// resourceGroupName - the name of the resource group.
// SSHPublicKeyName - the name of the SSH public key.
func (client SSHPublicKeysClient) Get(ctx context.Context, resourceGroupName string, SSHPublicKeyName string) (result SSHPublicKeyResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SSHPublicKeysClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, SSHPublicKeyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SSHPublicKeysClient) GetPreparer(ctx context.Context, resourceGroupName string, SSHPublicKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sshPublicKeyName":  autorest.Encode("path", SSHPublicKeyName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SSHPublicKeysClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SSHPublicKeysClient) GetResponder(resp *http.Response) (result SSHPublicKeyResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup lists all of the SSH public keys in the specified resource group. Use the nextLink property in
// the response to get the next page of SSH public keys.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client SSHPublicKeysClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result SSHPublicKeysGroupListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SSHPublicKeysClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.spkglr.Response.Response != nil {
				sc = result.spkglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.spkglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.spkglr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.spkglr.hasNextLink() && result.spkglr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client SSHPublicKeysClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SSHPublicKeysClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client SSHPublicKeysClient) ListByResourceGroupResponder(resp *http.Response) (result SSHPublicKeysGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client SSHPublicKeysClient) listByResourceGroupNextResults(ctx context.Context, lastResults SSHPublicKeysGroupListResult) (result SSHPublicKeysGroupListResult, err error) {
	req, err := lastResults.sSHPublicKeysGroupListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SSHPublicKeysClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result SSHPublicKeysGroupListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SSHPublicKeysClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListBySubscription lists all of the SSH public keys in the subscription. Use the nextLink property in the response
// to get the next page of SSH public keys.
func (client SSHPublicKeysClient) ListBySubscription(ctx context.Context) (result SSHPublicKeysGroupListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SSHPublicKeysClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.spkglr.Response.Response != nil {
				sc = result.spkglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.spkglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.spkglr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.spkglr.hasNextLink() && result.spkglr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client SSHPublicKeysClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/sshPublicKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client SSHPublicKeysClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client SSHPublicKeysClient) ListBySubscriptionResponder(resp *http.Response) (result SSHPublicKeysGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client SSHPublicKeysClient) listBySubscriptionNextResults(ctx context.Context, lastResults SSHPublicKeysGroupListResult) (result SSHPublicKeysGroupListResult, err error) {
	req, err := lastResults.sSHPublicKeysGroupListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client SSHPublicKeysClient) ListBySubscriptionComplete(ctx context.Context) (result SSHPublicKeysGroupListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SSHPublicKeysClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx)
	return
}

// Update updates a new SSH public key resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// SSHPublicKeyName - the name of the SSH public key.
// parameters - parameters supplied to update the SSH public key.
func (client SSHPublicKeysClient) Update(ctx context.Context, resourceGroupName string, SSHPublicKeyName string, parameters SSHPublicKeyUpdateResource) (result SSHPublicKeyResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SSHPublicKeysClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, SSHPublicKeyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SSHPublicKeysClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SSHPublicKeysClient) UpdatePreparer(ctx context.Context, resourceGroupName string, SSHPublicKeyName string, parameters SSHPublicKeyUpdateResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sshPublicKeyName":  autorest.Encode("path", SSHPublicKeyName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SSHPublicKeysClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SSHPublicKeysClient) UpdateResponder(resp *http.Response) (result SSHPublicKeyResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
