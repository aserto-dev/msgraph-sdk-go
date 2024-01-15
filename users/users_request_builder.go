package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie5faf868346cbd3c75e1e8329117d648c5f48d3737fa46dd1a10043c93c23754 "github.com/aserto-dev/msgraph-sdk-go/models"
    i0d805acb36b6faf1927302c3a0d4473c75d4a7188a2047e901c95d05fd534fae "github.com/aserto-dev/msgraph-sdk-go/models/odataerrors"
)

// UsersRequestBuilder builds and executes requests for operations under \users
type UsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UsersRequestBuilderGetQueryParameters retrieve a list of user objects.
type UsersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// UsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersRequestBuilderGetQueryParameters
}
// ByUserId gets an item from the github.com/aserto-dev/msgraph-sdk-go.users.item collection
func (m *UsersRequestBuilder) ByUserId(userId string)(*UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userId != "" {
        urlTplParams["user%2Did"] = userId
    }
    return NewUserItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUsersRequestBuilderInternal instantiates a new UsersRequestBuilder and sets the default values.
func NewUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersRequestBuilder) {
    m := &UsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users{?%24top,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewUsersRequestBuilder instantiates a new UsersRequestBuilder and sets the default values.
func NewUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve a list of user objects.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-list?view=graph-rest-1.0
func (m *UsersRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersRequestBuilderGetRequestConfiguration)(ie5faf868346cbd3c75e1e8329117d648c5f48d3737fa46dd1a10043c93c23754.UserCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i0d805acb36b6faf1927302c3a0d4473c75d4a7188a2047e901c95d05fd534fae.CreateODataErrorFromDiscriminatorValue,
        "5XX": i0d805acb36b6faf1927302c3a0d4473c75d4a7188a2047e901c95d05fd534fae.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie5faf868346cbd3c75e1e8329117d648c5f48d3737fa46dd1a10043c93c23754.CreateUserCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie5faf868346cbd3c75e1e8329117d648c5f48d3737fa46dd1a10043c93c23754.UserCollectionResponseable), nil
}
// ToGetRequestInformation retrieve a list of user objects.
func (m *UsersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UsersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersRequestBuilder) WithUrl(rawUrl string)(*UsersRequestBuilder) {
    return NewUsersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
