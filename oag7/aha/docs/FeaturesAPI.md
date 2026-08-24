# \FeaturesAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReleaseFeature**](FeaturesAPI.md#CreateReleaseFeature) | **Post** /releases/{release_id}/features | Create feature in release
[**GetFeature**](FeaturesAPI.md#GetFeature) | **Get** /features/{feature_id} | Get feature
[**ListFeatureIdeas**](FeaturesAPI.md#ListFeatureIdeas) | **Get** /features/{feature_id}/ideas | List ideas linked to a feature
[**ListFeatures**](FeaturesAPI.md#ListFeatures) | **Get** /features | List features
[**ListReleaseFeatures**](FeaturesAPI.md#ListReleaseFeatures) | **Get** /releases/{release_id}/features | List features in release
[**UpdateFeature**](FeaturesAPI.md#UpdateFeature) | **Put** /features/{feature_id} | Update feature



## CreateReleaseFeature

> FeatureResponse CreateReleaseFeature(ctx, releaseId).FeatureCreateRequest(featureCreateRequest).Execute()

Create feature in release



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
	featureCreateRequest := *openapiclient.NewFeatureCreateRequest(*openapiclient.NewFeatureCreate("Name_example")) // FeatureCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.CreateReleaseFeature(context.Background(), releaseId).FeatureCreateRequest(featureCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.CreateReleaseFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReleaseFeature`: FeatureResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.CreateReleaseFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** | Release ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleaseFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **featureCreateRequest** | [**FeatureCreateRequest**](FeatureCreateRequest.md) |  | 

### Return type

[**FeatureResponse**](FeatureResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeature

> FeatureResponse GetFeature(ctx, featureId).Execute()

Get feature



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
	featureId := "featureId_example" // string | Numeric ID or reference number (e.g., PROD-123)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.GetFeature(context.Background(), featureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.GetFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeature`: FeatureResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.GetFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | Numeric ID or reference number (e.g., PROD-123) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeatureResponse**](FeatureResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatureIdeas

> IdeasResponse ListFeatureIdeas(ctx, featureId).Page(page).PerPage(perPage).Execute()

List ideas linked to a feature



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
	resp, r, err := apiClient.FeaturesAPI.ListFeatureIdeas(context.Background(), featureId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.ListFeatureIdeas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeatureIdeas`: IdeasResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.ListFeatureIdeas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | Feature ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFeatureIdeasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**IdeasResponse**](IdeasResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatures

> FeaturesResponse ListFeatures(ctx).Q(q).UpdatedSince(updatedSince).Tag(tag).AssignedToUser(assignedToUser).Page(page).PerPage(perPage).Execute()

List features



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
	q := "q_example" // string | Sub-string to match against feature name or ID (optional)
	updatedSince := time.Now() // time.Time | UTC timestamp (ISO8601). Only features updated after this time. (optional)
	tag := "tag_example" // string | Filter by tag value (optional)
	assignedToUser := "assignedToUser_example" // string | Filter by assigned user ID or email (optional)
	page := int32(56) // int32 | Page number (optional)
	perPage := int32(56) // int32 | Results per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.ListFeatures(context.Background()).Q(q).UpdatedSince(updatedSince).Tag(tag).AssignedToUser(assignedToUser).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.ListFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeatures`: FeaturesResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.ListFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Sub-string to match against feature name or ID | 
 **updatedSince** | **time.Time** | UTC timestamp (ISO8601). Only features updated after this time. | 
 **tag** | **string** | Filter by tag value | 
 **assignedToUser** | **string** | Filter by assigned user ID or email | 
 **page** | **int32** | Page number | 
 **perPage** | **int32** | Results per page | 

### Return type

[**FeaturesResponse**](FeaturesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReleaseFeatures

> FeaturesResponse ListReleaseFeatures(ctx, releaseId).Page(page).PerPage(perPage).Execute()

List features in release



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
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.ListReleaseFeatures(context.Background(), releaseId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.ListReleaseFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReleaseFeatures`: FeaturesResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.ListReleaseFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** | Release ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiListReleaseFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**FeaturesResponse**](FeaturesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeature

> FeatureResponse UpdateFeature(ctx, featureId).FeatureUpdateRequest(featureUpdateRequest).Execute()

Update feature



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
	featureId := "featureId_example" // string | Numeric ID or reference number
	featureUpdateRequest := *openapiclient.NewFeatureUpdateRequest(*openapiclient.NewFeatureUpdate()) // FeatureUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.UpdateFeature(context.Background(), featureId).FeatureUpdateRequest(featureUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.UpdateFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFeature`: FeatureResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.UpdateFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | Numeric ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **featureUpdateRequest** | [**FeatureUpdateRequest**](FeatureUpdateRequest.md) |  | 

### Return type

[**FeatureResponse**](FeatureResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

