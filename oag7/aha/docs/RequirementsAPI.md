# \RequirementsAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeatureRequirement**](RequirementsAPI.md#CreateFeatureRequirement) | **Post** /features/{feature_id}/requirements | Create requirement for feature
[**DeleteRequirement**](RequirementsAPI.md#DeleteRequirement) | **Delete** /requirements/{requirement_id} | Delete requirement
[**GetRequirement**](RequirementsAPI.md#GetRequirement) | **Get** /requirements/{requirement_id} | Get requirement
[**ListFeatureRequirements**](RequirementsAPI.md#ListFeatureRequirements) | **Get** /features/{feature_id}/requirements | List feature requirements
[**UpdateRequirement**](RequirementsAPI.md#UpdateRequirement) | **Put** /requirements/{requirement_id} | Update requirement



## CreateFeatureRequirement

> RequirementResponse CreateFeatureRequirement(ctx, featureId).RequirementCreateRequest(requirementCreateRequest).Execute()

Create requirement for feature



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
	featureId := "featureId_example" // string | Feature ID or reference number
	requirementCreateRequest := *openapiclient.NewRequirementCreateRequest(*openapiclient.NewRequirementCreate("Name_example")) // RequirementCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementsAPI.CreateFeatureRequirement(context.Background(), featureId).RequirementCreateRequest(requirementCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementsAPI.CreateFeatureRequirement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFeatureRequirement`: RequirementResponse
	fmt.Fprintf(os.Stdout, "Response from `RequirementsAPI.CreateFeatureRequirement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | Feature ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeatureRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requirementCreateRequest** | [**RequirementCreateRequest**](RequirementCreateRequest.md) |  | 

### Return type

[**RequirementResponse**](RequirementResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRequirement

> DeleteRequirement(ctx, requirementId).Execute()

Delete requirement



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
	requirementId := "requirementId_example" // string | Requirement ID or reference number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequirementsAPI.DeleteRequirement(context.Background(), requirementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementsAPI.DeleteRequirement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requirementId** | **string** | Requirement ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequirementRequest struct via the builder pattern


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


## GetRequirement

> RequirementResponse GetRequirement(ctx, requirementId).Execute()

Get requirement



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
	requirementId := "requirementId_example" // string | Requirement ID or reference number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementsAPI.GetRequirement(context.Background(), requirementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementsAPI.GetRequirement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequirement`: RequirementResponse
	fmt.Fprintf(os.Stdout, "Response from `RequirementsAPI.GetRequirement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requirementId** | **string** | Requirement ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequirementResponse**](RequirementResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatureRequirements

> RequirementsResponse ListFeatureRequirements(ctx, featureId).Page(page).PerPage(perPage).Execute()

List feature requirements



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
	featureId := "featureId_example" // string | Feature ID or reference number
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementsAPI.ListFeatureRequirements(context.Background(), featureId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementsAPI.ListFeatureRequirements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeatureRequirements`: RequirementsResponse
	fmt.Fprintf(os.Stdout, "Response from `RequirementsAPI.ListFeatureRequirements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | Feature ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFeatureRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**RequirementsResponse**](RequirementsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRequirement

> RequirementResponse UpdateRequirement(ctx, requirementId).RequirementUpdateRequest(requirementUpdateRequest).Execute()

Update requirement



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
	requirementId := "requirementId_example" // string | Requirement ID or reference number
	requirementUpdateRequest := *openapiclient.NewRequirementUpdateRequest(*openapiclient.NewRequirementUpdate()) // RequirementUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementsAPI.UpdateRequirement(context.Background(), requirementId).RequirementUpdateRequest(requirementUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementsAPI.UpdateRequirement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRequirement`: RequirementResponse
	fmt.Fprintf(os.Stdout, "Response from `RequirementsAPI.UpdateRequirement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requirementId** | **string** | Requirement ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requirementUpdateRequest** | [**RequirementUpdateRequest**](RequirementUpdateRequest.md) |  | 

### Return type

[**RequirementResponse**](RequirementResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

