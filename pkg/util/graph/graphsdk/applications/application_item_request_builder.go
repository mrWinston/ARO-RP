package applications

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"

	i3989e3e55cd52fc1cab838f02782595a577b769f28259f22e5f9969d0d8d0c34 "github.com/Azure/ARO-RP/pkg/util/graph/graphsdk/applications/item"
	i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd "github.com/Azure/ARO-RP/pkg/util/graph/graphsdk/models"
	i590dfc7f28a1fc5720c211d996119093307169ae10220ddded8912d222cbd376 "github.com/Azure/ARO-RP/pkg/util/graph/graphsdk/models/odataerrors"
)

// ApplicationItemRequestBuilder provides operations to manage the collection of application entities.
type ApplicationItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ApplicationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ApplicationItemRequestBuilderGetQueryParameters get the properties and relationships of an application object.
type ApplicationItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	// Deprecated: This property is deprecated, use ExpandAsGetExpandQueryParameterType instead
	Expand []string `uriparametername:"%24expand"`
	// Expand related entities
	ExpandAsGetExpandQueryParameterType []i3989e3e55cd52fc1cab838f02782595a577b769f28259f22e5f9969d0d8d0c34.GetExpandQueryParameterType `uriparametername:"%24expand"`
	// Select properties to be returned
	// Deprecated: This property is deprecated, use SelectAsGetSelectQueryParameterType instead
	Select []string `uriparametername:"%24select"`
	// Select properties to be returned
	SelectAsGetSelectQueryParameterType []i3989e3e55cd52fc1cab838f02782595a577b769f28259f22e5f9969d0d8d0c34.GetSelectQueryParameterType `uriparametername:"%24select"`
}

// ApplicationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ApplicationItemRequestBuilderGetQueryParameters
}

// ApplicationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// AddPassword provides operations to call the addPassword method.
// returns a *ItemAddPasswordRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) AddPassword() *ItemAddPasswordRequestBuilder {
	return NewItemAddPasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewApplicationItemRequestBuilderInternal instantiates a new ApplicationItemRequestBuilder and sets the default values.
func NewApplicationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ApplicationItemRequestBuilder {
	m := &ApplicationItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}{?%24expand,%24select}", pathParameters),
	}
	return m
}

// NewApplicationItemRequestBuilder instantiates a new ApplicationItemRequestBuilder and sets the default values.
func NewApplicationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ApplicationItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewApplicationItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete an application object. When deleted, apps are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
//
// [Find more info here]: https://learn.microsoft.com/graph/api/application-delete?view=graph-rest-1.0
func (m *ApplicationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ApplicationItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"XXX": i590dfc7f28a1fc5720c211d996119093307169ae10220ddded8912d222cbd376.CreateODataErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get get the properties and relationships of an application object.
// returns a Applicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
//
// [Find more info here]: https://learn.microsoft.com/graph/api/application-get?view=graph-rest-1.0
func (m *ApplicationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ApplicationItemRequestBuilderGetRequestConfiguration) (i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.Applicationable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"XXX": i590dfc7f28a1fc5720c211d996119093307169ae10220ddded8912d222cbd376.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.CreateApplicationFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.Applicationable), nil
}

// Patch update the properties of an application object.
// returns a Applicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
//
// [Find more info here]: https://learn.microsoft.com/graph/api/application-update?view=graph-rest-1.0
func (m *ApplicationItemRequestBuilder) Patch(ctx context.Context, body i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.Applicationable, requestConfiguration *ApplicationItemRequestBuilderPatchRequestConfiguration) (i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.Applicationable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"XXX": i590dfc7f28a1fc5720c211d996119093307169ae10220ddded8912d222cbd376.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.CreateApplicationFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.Applicationable), nil
}

// RemovePassword provides operations to call the removePassword method.
// returns a *ItemRemovePasswordRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) RemovePassword() *ItemRemovePasswordRequestBuilder {
	return NewItemRemovePasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToDeleteRequestInformation delete an application object. When deleted, apps are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted.
// returns a *RequestInformation when successful
func (m *ApplicationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ApplicationItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/applications/{application%2Did}", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation get the properties and relationships of an application object.
// returns a *RequestInformation when successful
func (m *ApplicationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ApplicationItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPatchRequestInformation update the properties of an application object.
// returns a *RequestInformation when successful
func (m *ApplicationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.Applicationable, requestConfiguration *ApplicationItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/applications/{application%2Did}", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ApplicationItemRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) WithUrl(rawUrl string) *ApplicationItemRequestBuilder {
	return NewApplicationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
