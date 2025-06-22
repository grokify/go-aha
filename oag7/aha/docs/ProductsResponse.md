# ProductsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]ProductMeta**](ProductMeta.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewProductsResponse

`func NewProductsResponse() *ProductsResponse`

NewProductsResponse instantiates a new ProductsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsResponseWithDefaults

`func NewProductsResponseWithDefaults() *ProductsResponse`

NewProductsResponseWithDefaults instantiates a new ProductsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *ProductsResponse) GetProducts() []ProductMeta`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ProductsResponse) GetProductsOk() (*[]ProductMeta, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ProductsResponse) SetProducts(v []ProductMeta)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *ProductsResponse) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetPagination

`func (o *ProductsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ProductsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ProductsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ProductsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


