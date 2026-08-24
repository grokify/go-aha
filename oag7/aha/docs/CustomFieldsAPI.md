# \CustomFieldsAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCustomFieldDefinitions**](CustomFieldsAPI.md#ListCustomFieldDefinitions) | **Get** /custom_field_definitions | List all custom field definitions
[**ListCustomFieldOptions**](CustomFieldsAPI.md#ListCustomFieldOptions) | **Get** /custom_field_definitions/{id}/options | List options for a custom field
[**ListProductCustomFieldDefinitions**](CustomFieldsAPI.md#ListProductCustomFieldDefinitions) | **Get** /products/{product_id}/custom_field_definitions | List custom field definitions for a product



## ListCustomFieldDefinitions

> CustomFieldDefinitionListResponse ListCustomFieldDefinitions(ctx).Execute()

List all custom field definitions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.ListCustomFieldDefinitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.ListCustomFieldDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomFieldDefinitions`: CustomFieldDefinitionListResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.ListCustomFieldDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomFieldDefinitionsRequest struct via the builder pattern


### Return type

[**CustomFieldDefinitionListResponse**](CustomFieldDefinitionListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomFieldOptions

> CustomFieldOptionListResponse ListCustomFieldOptions(ctx, id).Execute()

List options for a custom field



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
	id := "id_example" // string | Custom field definition ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.ListCustomFieldOptions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.ListCustomFieldOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomFieldOptions`: CustomFieldOptionListResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.ListCustomFieldOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Custom field definition ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomFieldOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomFieldOptionListResponse**](CustomFieldOptionListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductCustomFieldDefinitions

> CustomFieldDefinitionListResponse ListProductCustomFieldDefinitions(ctx, productId).Execute()

List custom field definitions for a product



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
	resp, r, err := apiClient.CustomFieldsAPI.ListProductCustomFieldDefinitions(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.ListProductCustomFieldDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductCustomFieldDefinitions`: CustomFieldDefinitionListResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.ListProductCustomFieldDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product ID or reference prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProductCustomFieldDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomFieldDefinitionListResponse**](CustomFieldDefinitionListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

