package applications

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"

	i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd "github.com/Azure/ARO-RP/pkg/util/graph/graphsdk/models"
	i590dfc7f28a1fc5720c211d996119093307169ae10220ddded8912d222cbd376 "github.com/Azure/ARO-RP/pkg/util/graph/graphsdk/models/odataerrors"
)

// ApplicationsRequestBuilder provides operations to manage the collection of application entities.
type ApplicationsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ApplicationsRequestBuilderGetQueryParameters get the list of applications in this organization.
type ApplicationsRequestBuilderGetQueryParameters struct {
	// Include count of items
	Count *bool `uriparametername:"%24count"`
	// Expand related entities
	// Deprecated: This property is deprecated, use ExpandAsGetExpandQueryParameterType instead
	Expand []string `uriparametername:"%24expand"`
	// Expand related entities
	ExpandAsGetExpandQueryParameterType []GetExpandQueryParameterType `uriparametername:"%24expand"`
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Order items by property values
	// Deprecated: This property is deprecated, use OrderbyAsGetOrderbyQueryParameterType instead
	Orderby []string `uriparametername:"%24orderby"`
	// Order items by property values
	OrderbyAsGetOrderbyQueryParameterType []GetOrderbyQueryParameterType `uriparametername:"%24orderby"`
	// Search items by search phrases
	Search *string `uriparametername:"%24search"`
	// Select properties to be returned
	// Deprecated: This property is deprecated, use SelectAsGetSelectQueryParameterType instead
	Select []string `uriparametername:"%24select"`
	// Select properties to be returned
	SelectAsGetSelectQueryParameterType []GetSelectQueryParameterType `uriparametername:"%24select"`
	// Skip the first n items
	Skip *int32 `uriparametername:"%24skip"`
	// Show only the first n items
	Top *int32 `uriparametername:"%24top"`
}

// ApplicationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ApplicationsRequestBuilderGetQueryParameters
}

// ApplicationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByApplicationId provides operations to manage the collection of application entities.
// returns a *ApplicationItemRequestBuilder when successful
func (m *ApplicationsRequestBuilder) ByApplicationId(applicationId string) *ApplicationItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if applicationId != "" {
		urlTplParams["application%2Did"] = applicationId
	}
	return NewApplicationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewApplicationsRequestBuilderInternal instantiates a new ApplicationsRequestBuilder and sets the default values.
func NewApplicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ApplicationsRequestBuilder {
	m := &ApplicationsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
	}
	return m
}

// NewApplicationsRequestBuilder instantiates a new ApplicationsRequestBuilder and sets the default values.
func NewApplicationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ApplicationsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewApplicationsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get get the list of applications in this organization.
// returns a ApplicationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
//
// [Find more info here]: https://learn.microsoft.com/graph/api/application-list?view=graph-rest-1.0
func (m *ApplicationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ApplicationsRequestBuilderGetRequestConfiguration) (i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.ApplicationCollectionResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"XXX": i590dfc7f28a1fc5720c211d996119093307169ae10220ddded8912d222cbd376.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.CreateApplicationCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.ApplicationCollectionResponseable), nil
}

// Post create a new application object.
// returns a Applicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
//
// [Find more info here]: https://learn.microsoft.com/graph/api/application-post-applications?view=graph-rest-1.0
func (m *ApplicationsRequestBuilder) Post(ctx context.Context, body i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.Applicationable, requestConfiguration *ApplicationsRequestBuilderPostRequestConfiguration) (i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.Applicationable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
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

// ToGetRequestInformation get the list of applications in this organization.
// returns a *RequestInformation when successful
func (m *ApplicationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ApplicationsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation create a new application object.
// returns a *RequestInformation when successful
func (m *ApplicationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i6a022527509c6c974d313985d6b1e1814af5796dab5da8f53d13c951e06bb0cd.Applicationable, requestConfiguration *ApplicationsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/applications", m.BaseRequestBuilder.PathParameters)
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
// returns a *ApplicationsRequestBuilder when successful
func (m *ApplicationsRequestBuilder) WithUrl(rawUrl string) *ApplicationsRequestBuilder {
	return NewApplicationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
