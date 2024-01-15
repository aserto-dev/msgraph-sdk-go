package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie5faf868346cbd3c75e1e8329117d648c5f48d3737fa46dd1a10043c93c23754 "github.com/aserto-dev/msgraph-sdk-go/models"
    i0d805acb36b6faf1927302c3a0d4473c75d4a7188a2047e901c95d05fd534fae "github.com/aserto-dev/msgraph-sdk-go/models/odataerrors"
)

// UserItemRequestBuilder builds and executes requests for operations under \users\{user-id}
type UserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserItemRequestBuilderGetQueryParameters read properties and relationships of the user object.
type UserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserItemRequestBuilderGetQueryParameters
}
// NewUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserItemRequestBuilder) {
    m := &UserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read properties and relationships of the user object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-user-get?view=graph-rest-1.0
func (m *UserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserItemRequestBuilderGetRequestConfiguration)(ie5faf868346cbd3c75e1e8329117d648c5f48d3737fa46dd1a10043c93c23754.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i0d805acb36b6faf1927302c3a0d4473c75d4a7188a2047e901c95d05fd534fae.CreateODataErrorFromDiscriminatorValue,
        "5XX": i0d805acb36b6faf1927302c3a0d4473c75d4a7188a2047e901c95d05fd534fae.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie5faf868346cbd3c75e1e8329117d648c5f48d3737fa46dd1a10043c93c23754.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie5faf868346cbd3c75e1e8329117d648c5f48d3737fa46dd1a10043c93c23754.Userable), nil
}
// MemberOf the memberOf property
func (m *UserItemRequestBuilder) MemberOf()(*ItemMemberOfRequestBuilder) {
    return NewItemMemberOfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation read properties and relationships of the user object.
func (m *UserItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *UserItemRequestBuilder) WithUrl(rawUrl string)(*UserItemRequestBuilder) {
    return NewUserItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
