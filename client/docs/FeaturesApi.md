# \FeaturesApi

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeature**](FeaturesApi.md#GetFeature) | **Get** /features/{feature_id} | 
[**GetFeatures**](FeaturesApi.md#GetFeatures) | **Get** /features | Get all features
[**GetReleaseFeatures**](FeaturesApi.md#GetReleaseFeatures) | **Get** /releases/{release_id}/features | Get all features for a release
[**UpdateFeature**](FeaturesApi.md#UpdateFeature) | **Put** /features/{feature_id} | Update a feature&#39;s custom fields with tag-like value


# **GetFeature**
> FeatureWrap GetFeature(ctx, featureId)


Get a specific feature

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **featureId** | **string**| Numeric ID, or key of the feature to be retrieved | 

### Return type

[**FeatureWrap**](FeatureWrap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatures**
> FeaturesResponse GetFeatures(ctx, optional)
Get all features

Get all features

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string**| Sub-string to match against feature name or ID | 
 **updatedSince** | **time.Time**| UTC timestamp (in ISO8601 format) that the updated_at field must be larger than. | 
 **tag** | **string**| A string tag value. | 
 **assignedToUser** | **string**| The ID or email address of user to return assigned features for. | 
 **page** | **int32**| A specific page of results. | 
 **perPage** | **int32**| Number of results per page. | 

### Return type

[**FeaturesResponse**](FeaturesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseFeatures**
> FeaturesResponse GetReleaseFeatures(ctx, releaseId)
Get all features for a release

Get all features for a release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **releaseId** | **string**| Numeric ID, or key of the release to retrieve features for | 

### Return type

[**FeaturesResponse**](FeaturesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFeature**
> FeatureWrap UpdateFeature(ctx, featureId, feature)
Update a feature's custom fields with tag-like value

Update a feature's custom fields with tag-like value

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **featureId** | **string**| Numeric ID, or key of the feature to be retrieved | 
  **feature** | [**FeatureUpdate**](FeatureUpdate.md)| Feature properties to update | 

### Return type

[**FeatureWrap**](FeatureWrap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

