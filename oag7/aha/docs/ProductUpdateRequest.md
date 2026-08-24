# ProductUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | [**ProductUpdate**](ProductUpdate.md) |  | 

## Methods

### NewProductUpdateRequest

`func NewProductUpdateRequest(product ProductUpdate, ) *ProductUpdateRequest`

NewProductUpdateRequest instantiates a new ProductUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductUpdateRequestWithDefaults

`func NewProductUpdateRequestWithDefaults() *ProductUpdateRequest`

NewProductUpdateRequestWithDefaults instantiates a new ProductUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *ProductUpdateRequest) GetProduct() ProductUpdate`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ProductUpdateRequest) GetProductOk() (*ProductUpdate, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ProductUpdateRequest) SetProduct(v ProductUpdate)`

SetProduct sets Product field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


