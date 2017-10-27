# \ReleasesApi

All URIs are relative to *https://ringcentral.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductsProductIdReleasesGet**](ReleasesApi.md#ProductsProductIdReleasesGet) | **Get** /products/{productId}/releases | Releases API
[**ReleasesIdGet**](ReleasesApi.md#ReleasesIdGet) | **Get** /releases/{id} | 


# **ProductsProductIdReleasesGet**
> Releases ProductsProductIdReleasesGet($productId)

Releases API

Create a release


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string**| The id of the company being queried | 

### Return type

[**Releases**](Releases.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReleasesIdGet**
> Release ReleasesIdGet($id)



Get a specific release


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Numeric ID, or key of the release to be retrieved | 

### Return type

[**Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

