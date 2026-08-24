# \ReleasesAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRelease**](ReleasesAPI.md#CreateRelease) | **Post** /products/{product_id}/releases | Create release
[**GetRelease**](ReleasesAPI.md#GetRelease) | **Get** /releases/{release_id} | Get release
[**ListProductReleases**](ReleasesAPI.md#ListProductReleases) | **Get** /products/{product_id}/releases | List product releases
[**UpdateRelease**](ReleasesAPI.md#UpdateRelease) | **Put** /releases/{release_id} | Update release



## CreateRelease

> ReleaseResponse CreateRelease(ctx, productId).ReleaseCreateRequest(releaseCreateRequest).Execute()

Create release



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
	releaseCreateRequest := *openapiclient.NewReleaseCreateRequest(*openapiclient.NewReleaseCreate("Name_example")) // ReleaseCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasesAPI.CreateRelease(context.Background(), productId).ReleaseCreateRequest(releaseCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.CreateRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRelease`: ReleaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.CreateRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product ID or reference prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **releaseCreateRequest** | [**ReleaseCreateRequest**](ReleaseCreateRequest.md) |  | 

### Return type

[**ReleaseResponse**](ReleaseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelease

> ReleaseResponse GetRelease(ctx, releaseId).Execute()

Get release



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
	releaseId := "releaseId_example" // string | Release ID or reference number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasesAPI.GetRelease(context.Background(), releaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.GetRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRelease`: ReleaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.GetRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** | Release ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReleaseResponse**](ReleaseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductReleases

> ReleasesResponse ListProductReleases(ctx, productId).Page(page).PerPage(perPage).Execute()

List product releases



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
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasesAPI.ListProductReleases(context.Background(), productId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.ListProductReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductReleases`: ReleasesResponse
	fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.ListProductReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product ID or reference prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProductReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**ReleasesResponse**](ReleasesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRelease

> ReleaseResponse UpdateRelease(ctx, releaseId).ReleaseUpdateRequest(releaseUpdateRequest).Execute()

Update release



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
	releaseId := "releaseId_example" // string | Release ID or reference number
	releaseUpdateRequest := *openapiclient.NewReleaseUpdateRequest(*openapiclient.NewReleaseUpdate()) // ReleaseUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasesAPI.UpdateRelease(context.Background(), releaseId).ReleaseUpdateRequest(releaseUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.UpdateRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRelease`: ReleaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.UpdateRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** | Release ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **releaseUpdateRequest** | [**ReleaseUpdateRequest**](ReleaseUpdateRequest.md) |  | 

### Return type

[**ReleaseResponse**](ReleaseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

