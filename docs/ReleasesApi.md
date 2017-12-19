# \ReleasesApi

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductsProductIdReleasesGet**](ReleasesApi.md#ProductsProductIdReleasesGet) | **Get** /products/{product_id}/releases | Releases API
[**ProductsProductIdReleasesReleaseIdPut**](ReleasesApi.md#ProductsProductIdReleasesReleaseIdPut) | **Put** /products/{product_id}/releases/{release_id} | Update a release
[**ReleasesReleaseIdGet**](ReleasesApi.md#ReleasesReleaseIdGet) | **Get** /releases/{release_id} | 


# **ProductsProductIdReleasesGet**
> ReleasesResponse ProductsProductIdReleasesGet(ctx, productId)
Releases API

Create a release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **productId** | **string**| The id of the company being queried | 

### Return type

[**ReleasesResponse**](ReleasesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductsProductIdReleasesReleaseIdPut**
> ReleaseWrap ProductsProductIdReleasesReleaseIdPut(ctx, productId, releaseId, release)
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

# **ReleasesReleaseIdGet**
> ReleaseWrap ReleasesReleaseIdGet(ctx, releaseId)


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

