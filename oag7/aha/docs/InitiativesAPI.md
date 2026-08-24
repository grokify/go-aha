# \InitiativesAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProductInitiative**](InitiativesAPI.md#CreateProductInitiative) | **Post** /products/{product_id}/initiatives | Create initiative in product
[**GetInitiative**](InitiativesAPI.md#GetInitiative) | **Get** /initiatives/{initiative_id} | Get initiative
[**ListInitiatives**](InitiativesAPI.md#ListInitiatives) | **Get** /initiatives | List initiatives
[**ListProductInitiatives**](InitiativesAPI.md#ListProductInitiatives) | **Get** /products/{product_id}/initiatives | List product initiatives
[**UpdateInitiative**](InitiativesAPI.md#UpdateInitiative) | **Put** /initiatives/{initiative_id} | Update initiative



## CreateProductInitiative

> InitiativeResponse CreateProductInitiative(ctx, productId).InitiativeCreateRequest(initiativeCreateRequest).Execute()

Create initiative in product



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
	initiativeCreateRequest := *openapiclient.NewInitiativeCreateRequest(*openapiclient.NewInitiativeCreate("Name_example")) // InitiativeCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InitiativesAPI.CreateProductInitiative(context.Background(), productId).InitiativeCreateRequest(initiativeCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InitiativesAPI.CreateProductInitiative``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProductInitiative`: InitiativeResponse
	fmt.Fprintf(os.Stdout, "Response from `InitiativesAPI.CreateProductInitiative`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product ID or reference prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductInitiativeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **initiativeCreateRequest** | [**InitiativeCreateRequest**](InitiativeCreateRequest.md) |  | 

### Return type

[**InitiativeResponse**](InitiativeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInitiative

> InitiativeResponse GetInitiative(ctx, initiativeId).Execute()

Get initiative



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
	initiativeId := "initiativeId_example" // string | Initiative ID or reference number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InitiativesAPI.GetInitiative(context.Background(), initiativeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InitiativesAPI.GetInitiative``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInitiative`: InitiativeResponse
	fmt.Fprintf(os.Stdout, "Response from `InitiativesAPI.GetInitiative`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**initiativeId** | **string** | Initiative ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInitiativeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InitiativeResponse**](InitiativeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInitiatives

> InitiativesResponse ListInitiatives(ctx).Q(q).UpdatedSince(updatedSince).Page(page).PerPage(perPage).Execute()

List initiatives



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
	q := "q_example" // string | Sub-string to match against initiative name (optional)
	updatedSince := time.Now() // time.Time | UTC timestamp (ISO8601). Only initiatives updated after this time. (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InitiativesAPI.ListInitiatives(context.Background()).Q(q).UpdatedSince(updatedSince).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InitiativesAPI.ListInitiatives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInitiatives`: InitiativesResponse
	fmt.Fprintf(os.Stdout, "Response from `InitiativesAPI.ListInitiatives`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInitiativesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Sub-string to match against initiative name | 
 **updatedSince** | **time.Time** | UTC timestamp (ISO8601). Only initiatives updated after this time. | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**InitiativesResponse**](InitiativesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductInitiatives

> InitiativesResponse ListProductInitiatives(ctx, productId).Page(page).PerPage(perPage).Execute()

List product initiatives



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
	resp, r, err := apiClient.InitiativesAPI.ListProductInitiatives(context.Background(), productId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InitiativesAPI.ListProductInitiatives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductInitiatives`: InitiativesResponse
	fmt.Fprintf(os.Stdout, "Response from `InitiativesAPI.ListProductInitiatives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product ID or reference prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProductInitiativesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**InitiativesResponse**](InitiativesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInitiative

> InitiativeResponse UpdateInitiative(ctx, initiativeId).InitiativeUpdateRequest(initiativeUpdateRequest).Execute()

Update initiative



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
	initiativeId := "initiativeId_example" // string | Initiative ID or reference number
	initiativeUpdateRequest := *openapiclient.NewInitiativeUpdateRequest(*openapiclient.NewInitiativeUpdate()) // InitiativeUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InitiativesAPI.UpdateInitiative(context.Background(), initiativeId).InitiativeUpdateRequest(initiativeUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InitiativesAPI.UpdateInitiative``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInitiative`: InitiativeResponse
	fmt.Fprintf(os.Stdout, "Response from `InitiativesAPI.UpdateInitiative`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**initiativeId** | **string** | Initiative ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInitiativeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **initiativeUpdateRequest** | [**InitiativeUpdateRequest**](InitiativeUpdateRequest.md) |  | 

### Return type

[**InitiativeResponse**](InitiativeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

