# \FeaturesApi

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FeaturesFeatureIdGet**](FeaturesApi.md#FeaturesFeatureIdGet) | **Get** /features/{feature_id} | 
[**FeaturesFeatureIdPut**](FeaturesApi.md#FeaturesFeatureIdPut) | **Put** /features/{feature_id} | Update a feature&#39;s custom fields with tag-like value
[**FeaturesGet**](FeaturesApi.md#FeaturesGet) | **Get** /features | Get all features
[**ReleasesReleaseIdFeaturesGet**](FeaturesApi.md#ReleasesReleaseIdFeaturesGet) | **Get** /releases/{release_id}/features | Get all features for a release


# **FeaturesFeatureIdGet**
> FeatureWrap FeaturesFeatureIdGet($featureId)



Get a specific feature


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featureId** | **string**| Numeric ID, or key of the feature to be retrieved | 

### Return type

[**FeatureWrap**](FeatureWrap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeaturesFeatureIdPut**
> FeatureWrap FeaturesFeatureIdPut($featureId, $feature)

Update a feature's custom fields with tag-like value

Update a feature's custom fields with tag-like value


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

# **FeaturesGet**
> FeaturesResponse FeaturesGet($q, $updatedSince, $tag, $assignedToUser, $page, $perPage)

Get all features

Get all features


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string**| Sub-string to match against feature name or ID | [optional] 
 **updatedSince** | **time.Time**| UTC timestamp (in ISO8601 format) that the updated_at field must be larger than. | [optional] 
 **tag** | **string**| A string tag value. | [optional] 
 **assignedToUser** | **string**| The ID or email address of user to return assigned features for. | [optional] 
 **page** | **int32**| A specific page of results. | [optional] 
 **perPage** | **int32**| Number of results per page. | [optional] 

### Return type

[**FeaturesResponse**](FeaturesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReleasesReleaseIdFeaturesGet**
> FeaturesResponse ReleasesReleaseIdFeaturesGet($releaseId)

Get all features for a release

Get all features for a release


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseId** | **string**| Numeric ID, or key of the release to retrieve features for | 

### Return type

[**FeaturesResponse**](FeaturesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

