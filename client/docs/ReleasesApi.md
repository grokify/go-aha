# \ReleasesApi

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRelease**](ReleasesApi.md#GetRelease) | **Get** /releases/{release_id} | 
[**UpdateProductRelease**](ReleasesApi.md#UpdateProductRelease) | **Put** /products/{product_id}/releases/{release_id} | Update a release


# **GetRelease**
> ReleaseWrap GetRelease(ctx, releaseId)


Get a specific release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **releaseId** | **string**| Numeric ID, or key of the release to be retrieved | 

### Return type

[**ReleaseWrap**](ReleaseWrap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProductRelease**
> ReleaseWrap UpdateProductRelease(ctx, productId, releaseId, release)
Update a release

Update a release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **productId** | **string**| Numeric ID, or key of the product to create the release in | 
  **releaseId** | **string**| Numeric ID, or key of the release to be updated | 
  **release** | [**ReleaseUpdateWrap**](ReleaseUpdateWrap.md)| Release properties to update | 

### Return type

[**ReleaseWrap**](ReleaseWrap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

