# \ReleasesAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProductReleases**](ReleasesAPI.md#GetProductReleases) | **Get** /products/{product_id}/releases | Get product releases
[**GetRelease**](ReleasesAPI.md#GetRelease) | **Get** /releases/{release_id} | Get release
[**UpdateProductRelease**](ReleasesAPI.md#UpdateProductRelease) | **Put** /products/{product_id}/releases/{release_id} | Update product release



## GetProductReleases

> ReleasesResponse GetProductReleases(ctx, productId).Page(page).PerPage(perPage).Execute()

Get product releases



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
	productId := "productId_example" // string | Numeric ID, or key of the product to retrieve releases for.
	page := int32(56) // int32 | A specific page of results. (optional)
	perPage := int32(56) // int32 | Number of results per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasesAPI.GetProductReleases(context.Background(), productId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.GetProductReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductReleases`: ReleasesResponse
	fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.GetProductReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Numeric ID, or key of the product to retrieve releases for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A specific page of results. | 
 **perPage** | **int32** | Number of results per page. | 

### Return type

[**ReleasesResponse**](ReleasesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelease

> ReleaseWrap GetRelease(ctx, releaseId).Execute()

Get release



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
	releaseId := "releaseId_example" // string | Numeric ID, or key of the release to be retrieved

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasesAPI.GetRelease(context.Background(), releaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.GetRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRelease`: ReleaseWrap
	fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.GetRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** | Numeric ID, or key of the release to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReleaseWrap**](ReleaseWrap.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProductRelease

> ReleaseWrap UpdateProductRelease(ctx, productId, releaseId).ReleaseUpdateWrap(releaseUpdateWrap).Execute()

Update product release



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
	productId := "productId_example" // string | Numeric ID, or key of the product to create the release in
	releaseId := "releaseId_example" // string | Numeric ID, or key of the release to be updated
	releaseUpdateWrap := *openapiclient.NewReleaseUpdateWrap() // ReleaseUpdateWrap | Release properties to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasesAPI.UpdateProductRelease(context.Background(), productId, releaseId).ReleaseUpdateWrap(releaseUpdateWrap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.UpdateProductRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProductRelease`: ReleaseWrap
	fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.UpdateProductRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Numeric ID, or key of the product to create the release in | 
**releaseId** | **string** | Numeric ID, or key of the release to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **releaseUpdateWrap** | [**ReleaseUpdateWrap**](ReleaseUpdateWrap.md) | Release properties to update | 

### Return type

[**ReleaseWrap**](ReleaseWrap.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

