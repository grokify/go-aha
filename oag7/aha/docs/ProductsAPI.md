# \ProductsAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProduct**](ProductsAPI.md#GetProduct) | **Get** /products/{product_id} | Products API
[**GetProducts**](ProductsAPI.md#GetProducts) | **Get** /products | Products API



## GetProduct

> ProductResponse GetProduct(ctx, productId).Execute()

Products API



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
	productId := "productId_example" // string | Numeric ID, or key of the feature to be retrieved

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.GetProduct(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.GetProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProduct`: ProductResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.GetProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Numeric ID, or key of the feature to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductResponse**](ProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProducts

> ProductsResponse GetProducts(ctx).Page(page).PerPage(perPage).Execute()

Products API



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
	page := int32(56) // int32 | A specific page of results. (optional)
	perPage := int32(56) // int32 | Number of results per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.GetProducts(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.GetProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProducts`: ProductsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.GetProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A specific page of results. | 
 **perPage** | **int32** | Number of results per page. | 

### Return type

[**ProductsResponse**](ProductsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

