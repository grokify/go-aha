# \IdeasAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIdea**](IdeasAPI.md#DeleteIdea) | **Delete** /ideas/{idea_id} | Delete idea
[**GetIdea**](IdeasAPI.md#GetIdea) | **Get** /ideas/{idea_id} | Get idea
[**GetIdeaOrganization**](IdeasAPI.md#GetIdeaOrganization) | **Get** /idea_organizations/{idea_organization_id} | Get an idea organization by ID
[**GetIdeaUser**](IdeasAPI.md#GetIdeaUser) | **Get** /idea_users/{idea_user_id} | Get an idea user by ID
[**ListIdeaEndorsements**](IdeasAPI.md#ListIdeaEndorsements) | **Get** /ideas/{idea_id}/endorsements | List endorsements (votes) on an idea
[**ListIdeaOrganizations**](IdeasAPI.md#ListIdeaOrganizations) | **Get** /idea_organizations | List idea organizations (customer/account records)
[**ListIdeaUsers**](IdeasAPI.md#ListIdeaUsers) | **Get** /idea_users | List idea users (voter identities)
[**ListIdeas**](IdeasAPI.md#ListIdeas) | **Get** /ideas | List ideas
[**ListProductIdeaCategories**](IdeasAPI.md#ListProductIdeaCategories) | **Get** /products/{product_id}/idea_categories | List idea categories
[**UpdateIdea**](IdeasAPI.md#UpdateIdea) | **Put** /ideas/{idea_id} | Update idea



## DeleteIdea

> DeleteIdea(ctx, ideaId).Execute()

Delete idea



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	ideaId := "ideaId_example" // string | Idea ID or reference number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdeasAPI.DeleteIdea(context.Background(), ideaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdeasAPI.DeleteIdea``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ideaId** | **string** | Idea ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdeaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdea

> IdeaResponse GetIdea(ctx, ideaId).Execute()

Get idea



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	ideaId := "ideaId_example" // string | Idea ID or reference number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdeasAPI.GetIdea(context.Background(), ideaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdeasAPI.GetIdea``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdea`: IdeaResponse
	fmt.Fprintf(os.Stdout, "Response from `IdeasAPI.GetIdea`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ideaId** | **string** | Idea ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdeaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdeaResponse**](IdeaResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdeaOrganization

> IdeaOrganizationResponse GetIdeaOrganization(ctx, ideaOrganizationId).Execute()

Get an idea organization by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	ideaOrganizationId := "ideaOrganizationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdeasAPI.GetIdeaOrganization(context.Background(), ideaOrganizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdeasAPI.GetIdeaOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdeaOrganization`: IdeaOrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `IdeasAPI.GetIdeaOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ideaOrganizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdeaOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdeaOrganizationResponse**](IdeaOrganizationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdeaUser

> IdeaUserResponse GetIdeaUser(ctx, ideaUserId).Execute()

Get an idea user by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	ideaUserId := "ideaUserId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdeasAPI.GetIdeaUser(context.Background(), ideaUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdeasAPI.GetIdeaUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdeaUser`: IdeaUserResponse
	fmt.Fprintf(os.Stdout, "Response from `IdeasAPI.GetIdeaUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ideaUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdeaUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdeaUserResponse**](IdeaUserResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdeaEndorsements

> IdeaEndorsementsResponse ListIdeaEndorsements(ctx, ideaId).Page(page).PerPage(perPage).Execute()

List endorsements (votes) on an idea



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	ideaId := "ideaId_example" // string | Idea ID or reference number
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdeasAPI.ListIdeaEndorsements(context.Background(), ideaId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdeasAPI.ListIdeaEndorsements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdeaEndorsements`: IdeaEndorsementsResponse
	fmt.Fprintf(os.Stdout, "Response from `IdeasAPI.ListIdeaEndorsements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ideaId** | **string** | Idea ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIdeaEndorsementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**IdeaEndorsementsResponse**](IdeaEndorsementsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdeaOrganizations

> IdeaOrganizationsResponse ListIdeaOrganizations(ctx).Page(page).PerPage(perPage).Execute()

List idea organizations (customer/account records)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdeasAPI.ListIdeaOrganizations(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdeasAPI.ListIdeaOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdeaOrganizations`: IdeaOrganizationsResponse
	fmt.Fprintf(os.Stdout, "Response from `IdeasAPI.ListIdeaOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdeaOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**IdeaOrganizationsResponse**](IdeaOrganizationsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdeaUsers

> IdeaUsersResponse ListIdeaUsers(ctx).Page(page).PerPage(perPage).Execute()

List idea users (voter identities)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdeasAPI.ListIdeaUsers(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdeasAPI.ListIdeaUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdeaUsers`: IdeaUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `IdeasAPI.ListIdeaUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdeaUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**IdeaUsersResponse**](IdeaUsersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdeas

> IdeasResponse ListIdeas(ctx).Q(q).Spam(spam).WorkflowStatus(workflowStatus).Sort(sort).CreatedBefore(createdBefore).CreatedSince(createdSince).UpdatedSince(updatedSince).Tag(tag).UserId(userId).IdeaUserId(ideaUserId).Page(page).PerPage(perPage).Fields(fields).Execute()

List ideas



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	q := "q_example" // string | Search term to match against idea name (optional)
	spam := true // bool | When true, shows ideas marked as spam (optional)
	workflowStatus := "workflowStatus_example" // string | Filter by workflow status ID or name (optional)
	sort := "sort_example" // string | Sort order (optional)
	createdBefore := time.Now() // time.Time | UTC timestamp (ISO8601). Only ideas created before this time. (optional)
	createdSince := time.Now() // time.Time | UTC timestamp (ISO8601). Only ideas created after this time. (optional)
	updatedSince := time.Now() // time.Time | UTC timestamp (ISO8601). Only ideas updated after this time. (optional)
	tag := "tag_example" // string | Filter by tag value (optional)
	userId := "userId_example" // string | Filter by creator user ID (optional)
	ideaUserId := "ideaUserId_example" // string | Filter by idea user ID (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)
	fields := "fields_example" // string | Comma-separated list of additional idea fields to include in each list item (e.g. \"votes,categories,score\"). Aha's list endpoint omits these by default; passing this parameter overrides Aha's default field set rather than adding to it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdeasAPI.ListIdeas(context.Background()).Q(q).Spam(spam).WorkflowStatus(workflowStatus).Sort(sort).CreatedBefore(createdBefore).CreatedSince(createdSince).UpdatedSince(updatedSince).Tag(tag).UserId(userId).IdeaUserId(ideaUserId).Page(page).PerPage(perPage).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdeasAPI.ListIdeas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdeas`: IdeasResponse
	fmt.Fprintf(os.Stdout, "Response from `IdeasAPI.ListIdeas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdeasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Search term to match against idea name | 
 **spam** | **bool** | When true, shows ideas marked as spam | 
 **workflowStatus** | **string** | Filter by workflow status ID or name | 
 **sort** | **string** | Sort order | 
 **createdBefore** | **time.Time** | UTC timestamp (ISO8601). Only ideas created before this time. | 
 **createdSince** | **time.Time** | UTC timestamp (ISO8601). Only ideas created after this time. | 
 **updatedSince** | **time.Time** | UTC timestamp (ISO8601). Only ideas updated after this time. | 
 **tag** | **string** | Filter by tag value | 
 **userId** | **string** | Filter by creator user ID | 
 **ideaUserId** | **string** | Filter by idea user ID | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **fields** | **string** | Comma-separated list of additional idea fields to include in each list item (e.g. \&quot;votes,categories,score\&quot;). Aha&#39;s list endpoint omits these by default; passing this parameter overrides Aha&#39;s default field set rather than adding to it. | 

### Return type

[**IdeasResponse**](IdeasResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductIdeaCategories

> IdeaCategoriesResponse ListProductIdeaCategories(ctx, productId).Execute()

List idea categories



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	productId := "productId_example" // string | Product ID or reference prefix

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdeasAPI.ListProductIdeaCategories(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdeasAPI.ListProductIdeaCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductIdeaCategories`: IdeaCategoriesResponse
	fmt.Fprintf(os.Stdout, "Response from `IdeasAPI.ListProductIdeaCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product ID or reference prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProductIdeaCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdeaCategoriesResponse**](IdeaCategoriesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdea

> IdeaResponse UpdateIdea(ctx, ideaId).IdeaUpdateRequest(ideaUpdateRequest).Execute()

Update idea



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	ideaId := "ideaId_example" // string | Idea ID or reference number
	ideaUpdateRequest := *openapiclient.NewIdeaUpdateRequest(*openapiclient.NewIdeaUpdate()) // IdeaUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdeasAPI.UpdateIdea(context.Background(), ideaId).IdeaUpdateRequest(ideaUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdeasAPI.UpdateIdea``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdea`: IdeaResponse
	fmt.Fprintf(os.Stdout, "Response from `IdeasAPI.UpdateIdea`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ideaId** | **string** | Idea ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdeaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ideaUpdateRequest** | [**IdeaUpdateRequest**](IdeaUpdateRequest.md) |  | 

### Return type

[**IdeaResponse**](IdeaResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

