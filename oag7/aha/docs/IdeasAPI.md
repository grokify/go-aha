# \IdeasAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdea**](IdeasAPI.md#GetIdea) | **Get** /ideas/{idea_id} | Get Idea



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

