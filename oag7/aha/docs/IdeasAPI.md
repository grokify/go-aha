# \IdeasAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdea**](IdeasAPI.md#GetIdea) | **Get** /ideas/{idea_id} | Get Idea
[**ListIdeas**](IdeasAPI.md#ListIdeas) | **Get** /ideas | List ideas



## GetIdea

> IdeaResponse GetIdea(ctx, ideaId).Execute()

Get Idea

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	ideaId := "ideaId_example" // string | 

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
**ideaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdeaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdeaResponse**](IdeaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdeas

> IdeasResponse ListIdeas(ctx).Q(q).Spam(spam).WorkflowStatus(workflowStatus).Sort(sort).CreatedBefore(createdBefore).CreatedSince(createdSince).UpdatedSince(updatedSince).Tag(tag).UserId(userId).IdeaUserId(ideaUserId).Page(page).PerPage(perPage).Execute()

List ideas



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	q := "q_example" // string | Search term to match against the idea name (optional)
	spam := true // bool | When true, shows ideas that have been marked as spam. By default, no spam ideas will be shown. (optional)
	workflowStatus := "workflowStatus_example" // string | When present, filters to ideas with the provided workflow status ID or name. (optional)
	sort := "sort_example" // string | Sorting of the list of ideas. Accepted values are recent, trending, or popular. (optional)
	createdBefore := time.Now() // time.Time | UTC timestamp (in ISO8601 format). If provided, only ideas created before the timestamp will be returned. (optional)
	createdSince := time.Now() // time.Time | UTC timestamp (in ISO8601 format). If provided, only ideas created after the timestamp will be returned. (optional)
	updatedSince := time.Now() // time.Time | UTC timestamp (in ISO8601 format). If provided, only ideas updated or created after the timestamp will be returned. (optional)
	tag := "tag_example" // string | String tag value. If provided, only ideas with the associated tag will be returned. (optional)
	userId := "userId_example" // string | ID of a user. If provided, only ideas created by that user will be returned. (optional)
	ideaUserId := "ideaUserId_example" // string | ID of an idea user. If provided, only ideas created by that idea user will be returned. (optional)
	page := int32(56) // int32 | A specific page of results. (optional)
	perPage := int32(56) // int32 | Number of results per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdeasAPI.ListIdeas(context.Background()).Q(q).Spam(spam).WorkflowStatus(workflowStatus).Sort(sort).CreatedBefore(createdBefore).CreatedSince(createdSince).UpdatedSince(updatedSince).Tag(tag).UserId(userId).IdeaUserId(ideaUserId).Page(page).PerPage(perPage).Execute()
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
 **q** | **string** | Search term to match against the idea name | 
 **spam** | **bool** | When true, shows ideas that have been marked as spam. By default, no spam ideas will be shown. | 
 **workflowStatus** | **string** | When present, filters to ideas with the provided workflow status ID or name. | 
 **sort** | **string** | Sorting of the list of ideas. Accepted values are recent, trending, or popular. | 
 **createdBefore** | **time.Time** | UTC timestamp (in ISO8601 format). If provided, only ideas created before the timestamp will be returned. | 
 **createdSince** | **time.Time** | UTC timestamp (in ISO8601 format). If provided, only ideas created after the timestamp will be returned. | 
 **updatedSince** | **time.Time** | UTC timestamp (in ISO8601 format). If provided, only ideas updated or created after the timestamp will be returned. | 
 **tag** | **string** | String tag value. If provided, only ideas with the associated tag will be returned. | 
 **userId** | **string** | ID of a user. If provided, only ideas created by that user will be returned. | 
 **ideaUserId** | **string** | ID of an idea user. If provided, only ideas created by that idea user will be returned. | 
 **page** | **int32** | A specific page of results. | 
 **perPage** | **int32** | Number of results per page. | 

### Return type

[**IdeasResponse**](IdeasResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

