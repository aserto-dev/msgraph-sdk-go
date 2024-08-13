package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie5faf868346cbd3c75e1e8329117d648c5f48d3737fa46dd1a10043c93c23754 "github.com/aserto-dev/msgraph-sdk-go/models"
    i0d805acb36b6faf1927302c3a0d4473c75d4a7188a2047e901c95d05fd534fae "github.com/aserto-dev/msgraph-sdk-go/models/odataerrors"
)

// ItemMembersRequestBuilder builds and executes requests for operations under \groups\{group-id}\members
type ItemMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMembersRequestBuilderGetQueryParameters get a list of the group's direct members. A group can have users, organizational contacts, devices, service principals and other groups as members. This operation is not transitive.
type ItemMembersRequestBuilderGetQueryParameters struct {
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
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMembersRequestBuilderGetQueryParameters
}
// NewItemMembersRequestBuilderInternal instantiates a new MembersRequestBuilder and sets the default values.
func NewItemMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersRequestBuilder) {
    m := &ItemMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/members{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemMembersRequestBuilder instantiates a new MembersRequestBuilder and sets the default values.
func NewItemMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a list of the group's direct members. A group can have users, organizational contacts, devices, service principals and other groups as members. This operation is not transitive.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-list-members?view=graph-rest-1.0
func (m *ItemMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMembersRequestBuilderGetRequestConfiguration)(ie5faf868346cbd3c75e1e8329117d648c5f48d3737fa46dd1a10043c93c23754.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i0d805acb36b6faf1927302c3a0d4473c75d4a7188a2047e901c95d05fd534fae.CreateODataErrorFromDiscriminatorValue,
        "5XX": i0d805acb36b6faf1927302c3a0d4473c75d4a7188a2047e901c95d05fd534fae.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie5faf868346cbd3c75e1e8329117d648c5f48d3737fa46dd1a10043c93c23754.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie5faf868346cbd3c75e1e8329117d648c5f48d3737fa46dd1a10043c93c23754.DirectoryObjectCollectionResponseable), nil
}
// ToGetRequestInformation get a list of the group's direct members. A group can have users, organizational contacts, devices, service principals and other groups as members. This operation is not transitive.
func (m *ItemMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemMembersRequestBuilder) WithUrl(rawUrl string)(*ItemMembersRequestBuilder) {
    return NewItemMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
