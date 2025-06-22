# \ReleasesApi

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProductReleases**](ReleasesApi.md#GetProductReleases) | **Get** /products/{product_id}/releases | 
[**GetRelease**](ReleasesApi.md#GetRelease) | **Get** /releases/{release_id} | 
[**UpdateProductRelease**](ReleasesApi.md#UpdateProductRelease) | **Put** /products/{product_id}/releases/{release_id} | Update a release



## GetProductReleases

> ReleasesResponse GetProductReleases(ctx, productId, optional)



Get releases for a product release

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string**| Numeric ID, or key of the product to retrieve releases for. | 
 **optional** | ***GetProductReleasesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProductReleasesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A specific page of results. | 
 **perPage** | **optional.Int32**| Number of results per page. | 

### Return type

[**ReleasesResponse**](ReleasesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelease

> ReleaseWrap GetRelease(ctx, releaseId)



Get a specific release

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string**| Numeric ID, or key of the release to be retrieved | 

### Return type

[**ReleaseWrap**](ReleaseWrap.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProductRelease

> ReleaseWrap UpdateProductRelease(ctx, productId, releaseId, release)

Update a release

Update a release

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

