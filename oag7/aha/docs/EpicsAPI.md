# \EpicsAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReleaseEpic**](EpicsAPI.md#CreateReleaseEpic) | **Post** /releases/{release_id}/epics | Create epic in release
[**GetEpic**](EpicsAPI.md#GetEpic) | **Get** /epics/{epic_id} | Get epic
[**ListEpics**](EpicsAPI.md#ListEpics) | **Get** /epics | List epics
[**ListProductEpics**](EpicsAPI.md#ListProductEpics) | **Get** /products/{product_id}/epics | List product epics
[**UpdateEpic**](EpicsAPI.md#UpdateEpic) | **Put** /epics/{epic_id} | Update epic



## CreateReleaseEpic

> EpicResponse CreateReleaseEpic(ctx, releaseId).EpicCreateRequest(epicCreateRequest).Execute()

Create epic in release



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
	epicCreateRequest := *openapiclient.NewEpicCreateRequest(*openapiclient.NewEpicCreate("Name_example")) // EpicCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpicsAPI.CreateReleaseEpic(context.Background(), releaseId).EpicCreateRequest(epicCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpicsAPI.CreateReleaseEpic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReleaseEpic`: EpicResponse
	fmt.Fprintf(os.Stdout, "Response from `EpicsAPI.CreateReleaseEpic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** | Release ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleaseEpicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **epicCreateRequest** | [**EpicCreateRequest**](EpicCreateRequest.md) |  | 

### Return type

[**EpicResponse**](EpicResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEpic

> EpicResponse GetEpic(ctx, epicId).Execute()

Get epic



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
	epicId := "epicId_example" // string | Epic ID or reference number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpicsAPI.GetEpic(context.Background(), epicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpicsAPI.GetEpic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEpic`: EpicResponse
	fmt.Fprintf(os.Stdout, "Response from `EpicsAPI.GetEpic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicId** | **string** | Epic ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEpicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EpicResponse**](EpicResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEpics

> EpicsResponse ListEpics(ctx).Q(q).UpdatedSince(updatedSince).Page(page).PerPage(perPage).Execute()

List epics



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
	q := "q_example" // string | Sub-string to match against epic name (optional)
	updatedSince := time.Now() // time.Time | UTC timestamp (ISO8601). Only epics updated after this time. (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpicsAPI.ListEpics(context.Background()).Q(q).UpdatedSince(updatedSince).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpicsAPI.ListEpics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEpics`: EpicsResponse
	fmt.Fprintf(os.Stdout, "Response from `EpicsAPI.ListEpics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEpicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Sub-string to match against epic name | 
 **updatedSince** | **time.Time** | UTC timestamp (ISO8601). Only epics updated after this time. | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**EpicsResponse**](EpicsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductEpics

> EpicsResponse ListProductEpics(ctx, productId).Page(page).PerPage(perPage).Execute()

List product epics



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
	resp, r, err := apiClient.EpicsAPI.ListProductEpics(context.Background(), productId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpicsAPI.ListProductEpics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductEpics`: EpicsResponse
	fmt.Fprintf(os.Stdout, "Response from `EpicsAPI.ListProductEpics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product ID or reference prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProductEpicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**EpicsResponse**](EpicsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEpic

> EpicResponse UpdateEpic(ctx, epicId).EpicUpdateRequest(epicUpdateRequest).Execute()

Update epic



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
	epicId := "epicId_example" // string | Epic ID or reference number
	epicUpdateRequest := *openapiclient.NewEpicUpdateRequest(*openapiclient.NewEpicUpdate()) // EpicUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpicsAPI.UpdateEpic(context.Background(), epicId).EpicUpdateRequest(epicUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpicsAPI.UpdateEpic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEpic`: EpicResponse
	fmt.Fprintf(os.Stdout, "Response from `EpicsAPI.UpdateEpic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicId** | **string** | Epic ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEpicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **epicUpdateRequest** | [**EpicUpdateRequest**](EpicUpdateRequest.md) |  | 

### Return type

[**EpicResponse**](EpicResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

