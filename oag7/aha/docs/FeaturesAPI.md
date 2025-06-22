# \FeaturesAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeature**](FeaturesAPI.md#GetFeature) | **Get** /features/{feature_id} | 
[**GetFeatures**](FeaturesAPI.md#GetFeatures) | **Get** /features | Get all features
[**GetReleaseFeatures**](FeaturesAPI.md#GetReleaseFeatures) | **Get** /releases/{release_id}/features | Get all features for a release
[**UpdateFeature**](FeaturesAPI.md#UpdateFeature) | **Put** /features/{feature_id} | Update a feature&#39;s custom fields with tag-like value



## GetFeature

> FeatureWrap GetFeature(ctx, featureId).Execute()





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
	featureId := "featureId_example" // string | Numeric ID, or key of the feature to be retrieved

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.GetFeature(context.Background(), featureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.GetFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeature`: FeatureWrap
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.GetFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | Numeric ID, or key of the feature to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeatureWrap**](FeatureWrap.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatures

> FeaturesResponse GetFeatures(ctx).Q(q).UpdatedSince(updatedSince).Tag(tag).AssignedToUser(assignedToUser).Page(page).PerPage(perPage).Execute()

Get all features



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	q := "q_example" // string | Sub-string to match against feature name or ID (optional)
	updatedSince := time.Now() // time.Time | UTC timestamp (in ISO8601 format) that the updated_at field must be larger than. (optional)
	tag := "tag_example" // string | A string tag value. (optional)
	assignedToUser := "assignedToUser_example" // string | The ID or email address of user to return assigned features for. (optional)
	page := int32(56) // int32 | A specific page of results. (optional)
	perPage := int32(56) // int32 | Number of results per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.GetFeatures(context.Background()).Q(q).UpdatedSince(updatedSince).Tag(tag).AssignedToUser(assignedToUser).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.GetFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeatures`: FeaturesResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.GetFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Sub-string to match against feature name or ID | 
 **updatedSince** | **time.Time** | UTC timestamp (in ISO8601 format) that the updated_at field must be larger than. | 
 **tag** | **string** | A string tag value. | 
 **assignedToUser** | **string** | The ID or email address of user to return assigned features for. | 
 **page** | **int32** | A specific page of results. | 
 **perPage** | **int32** | Number of results per page. | 

### Return type

[**FeaturesResponse**](FeaturesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseFeatures

> FeaturesResponse GetReleaseFeatures(ctx, releaseId).Page(page).PerPage(perPage).Execute()

Get all features for a release



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
	releaseId := "releaseId_example" // string | Numeric ID, or key of the release to retrieve features for
	page := int32(56) // int32 | A specific page of results. (optional)
	perPage := int32(56) // int32 | Number of results per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.GetReleaseFeatures(context.Background(), releaseId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.GetReleaseFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseFeatures`: FeaturesResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.GetReleaseFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** | Numeric ID, or key of the release to retrieve features for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A specific page of results. | 
 **perPage** | **int32** | Number of results per page. | 

### Return type

[**FeaturesResponse**](FeaturesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeature

> FeatureWrap UpdateFeature(ctx, featureId).FeatureUpdate(featureUpdate).Execute()

Update a feature's custom fields with tag-like value



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
	featureId := "featureId_example" // string | Numeric ID, or key of the feature to be retrieved
	featureUpdate := *openapiclient.NewFeatureUpdate() // FeatureUpdate | Feature properties to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.UpdateFeature(context.Background(), featureId).FeatureUpdate(featureUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.UpdateFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFeature`: FeatureWrap
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.UpdateFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | Numeric ID, or key of the feature to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **featureUpdate** | [**FeatureUpdate**](FeatureUpdate.md) | Feature properties to update | 

### Return type

[**FeatureWrap**](FeatureWrap.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

