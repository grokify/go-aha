# \StrategicModelsAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProductStrategicModel**](StrategicModelsAPI.md#CreateProductStrategicModel) | **Post** /products/{product_id}/strategy_models | Create strategic model in product
[**GetStrategicModel**](StrategicModelsAPI.md#GetStrategicModel) | **Get** /strategy_models/{strategy_model_id} | Get strategic model
[**ListProductStrategicModels**](StrategicModelsAPI.md#ListProductStrategicModels) | **Get** /products/{product_id}/strategy_models | List product strategic models
[**ListStrategicModels**](StrategicModelsAPI.md#ListStrategicModels) | **Get** /strategy_models | List strategic models
[**UpdateStrategicModel**](StrategicModelsAPI.md#UpdateStrategicModel) | **Put** /strategy_models/{strategy_model_id} | Update strategic model
[**UpdateStrategicModelComponent**](StrategicModelsAPI.md#UpdateStrategicModelComponent) | **Put** /strategy_models/{strategy_model_id}/components/{component_id} | Update strategic model component



## CreateProductStrategicModel

> StrategicModelResponse CreateProductStrategicModel(ctx, productId).StrategicModelCreateRequest(strategicModelCreateRequest).Execute()

Create strategic model in product



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
	strategicModelCreateRequest := *openapiclient.NewStrategicModelCreateRequest(*openapiclient.NewStrategicModelCreate("Name_example", "Kind_example")) // StrategicModelCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StrategicModelsAPI.CreateProductStrategicModel(context.Background(), productId).StrategicModelCreateRequest(strategicModelCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StrategicModelsAPI.CreateProductStrategicModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProductStrategicModel`: StrategicModelResponse
	fmt.Fprintf(os.Stdout, "Response from `StrategicModelsAPI.CreateProductStrategicModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product ID or reference prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductStrategicModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **strategicModelCreateRequest** | [**StrategicModelCreateRequest**](StrategicModelCreateRequest.md) |  | 

### Return type

[**StrategicModelResponse**](StrategicModelResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStrategicModel

> StrategicModelResponse GetStrategicModel(ctx, strategyModelId).Execute()

Get strategic model



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
	strategyModelId := "strategyModelId_example" // string | Strategic model ID or reference number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StrategicModelsAPI.GetStrategicModel(context.Background(), strategyModelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StrategicModelsAPI.GetStrategicModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStrategicModel`: StrategicModelResponse
	fmt.Fprintf(os.Stdout, "Response from `StrategicModelsAPI.GetStrategicModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**strategyModelId** | **string** | Strategic model ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStrategicModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StrategicModelResponse**](StrategicModelResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductStrategicModels

> StrategicModelsResponse ListProductStrategicModels(ctx, productId).Kind(kind).Page(page).PerPage(perPage).Execute()

List product strategic models



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
	kind := "kind_example" // string | Filter by model kind (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StrategicModelsAPI.ListProductStrategicModels(context.Background(), productId).Kind(kind).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StrategicModelsAPI.ListProductStrategicModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductStrategicModels`: StrategicModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `StrategicModelsAPI.ListProductStrategicModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product ID or reference prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProductStrategicModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kind** | **string** | Filter by model kind | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**StrategicModelsResponse**](StrategicModelsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStrategicModels

> StrategicModelsResponse ListStrategicModels(ctx).Kind(kind).Page(page).PerPage(perPage).Execute()

List strategic models



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
	kind := "kind_example" // string | Filter by model kind (e.g., \"Opportunity\", \"Lean Canvas\") (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StrategicModelsAPI.ListStrategicModels(context.Background()).Kind(kind).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StrategicModelsAPI.ListStrategicModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStrategicModels`: StrategicModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `StrategicModelsAPI.ListStrategicModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStrategicModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Filter by model kind (e.g., \&quot;Opportunity\&quot;, \&quot;Lean Canvas\&quot;) | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**StrategicModelsResponse**](StrategicModelsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStrategicModel

> StrategicModelResponse UpdateStrategicModel(ctx, strategyModelId).StrategicModelUpdateRequest(strategicModelUpdateRequest).Execute()

Update strategic model



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
	strategyModelId := "strategyModelId_example" // string | Strategic model ID or reference number
	strategicModelUpdateRequest := *openapiclient.NewStrategicModelUpdateRequest(*openapiclient.NewStrategicModelUpdate()) // StrategicModelUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StrategicModelsAPI.UpdateStrategicModel(context.Background(), strategyModelId).StrategicModelUpdateRequest(strategicModelUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StrategicModelsAPI.UpdateStrategicModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStrategicModel`: StrategicModelResponse
	fmt.Fprintf(os.Stdout, "Response from `StrategicModelsAPI.UpdateStrategicModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**strategyModelId** | **string** | Strategic model ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStrategicModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **strategicModelUpdateRequest** | [**StrategicModelUpdateRequest**](StrategicModelUpdateRequest.md) |  | 

### Return type

[**StrategicModelResponse**](StrategicModelResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStrategicModelComponent

> StrategicModelComponentResponse UpdateStrategicModelComponent(ctx, strategyModelId, componentId).StrategicModelComponentUpdateRequest(strategicModelComponentUpdateRequest).Execute()

Update strategic model component



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
	strategyModelId := "strategyModelId_example" // string | Strategic model ID or reference number
	componentId := "componentId_example" // string | Component ID within the strategic model
	strategicModelComponentUpdateRequest := *openapiclient.NewStrategicModelComponentUpdateRequest(*openapiclient.NewStrategicModelComponentUpdate()) // StrategicModelComponentUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StrategicModelsAPI.UpdateStrategicModelComponent(context.Background(), strategyModelId, componentId).StrategicModelComponentUpdateRequest(strategicModelComponentUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StrategicModelsAPI.UpdateStrategicModelComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStrategicModelComponent`: StrategicModelComponentResponse
	fmt.Fprintf(os.Stdout, "Response from `StrategicModelsAPI.UpdateStrategicModelComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**strategyModelId** | **string** | Strategic model ID or reference number | 
**componentId** | **string** | Component ID within the strategic model | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStrategicModelComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **strategicModelComponentUpdateRequest** | [**StrategicModelComponentUpdateRequest**](StrategicModelComponentUpdateRequest.md) |  | 

### Return type

[**StrategicModelComponentResponse**](StrategicModelComponentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

