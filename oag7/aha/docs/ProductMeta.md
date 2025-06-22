# ProductMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the product. | [optional] 
**ReferencePrefix** | Pointer to **string** | The reference prefix slug for the product. | [optional] 
**Name** | Pointer to **string** | The name for the product. | [optional] 
**ProductLine** | Pointer to **bool** | Whether the product is a product line or not. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date-time when the product was created. | [optional] 

## Methods

### NewProductMeta

`func NewProductMeta() *ProductMeta`

NewProductMeta instantiates a new ProductMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductMetaWithDefaults

`func NewProductMetaWithDefaults() *ProductMeta`

NewProductMetaWithDefaults instantiates a new ProductMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductMeta) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductMeta) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferencePrefix

`func (o *ProductMeta) GetReferencePrefix() string`

GetReferencePrefix returns the ReferencePrefix field if non-nil, zero value otherwise.

### GetReferencePrefixOk

`func (o *ProductMeta) GetReferencePrefixOk() (*string, bool)`

GetReferencePrefixOk returns a tuple with the ReferencePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencePrefix

`func (o *ProductMeta) SetReferencePrefix(v string)`

SetReferencePrefix sets ReferencePrefix field to given value.

### HasReferencePrefix

`func (o *ProductMeta) HasReferencePrefix() bool`

HasReferencePrefix returns a boolean if a field has been set.

### GetName

`func (o *ProductMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductLine

`func (o *ProductMeta) GetProductLine() bool`

GetProductLine returns the ProductLine field if non-nil, zero value otherwise.

### GetProductLineOk

`func (o *ProductMeta) GetProductLineOk() (*bool, bool)`

GetProductLineOk returns a tuple with the ProductLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductLine

`func (o *ProductMeta) SetProductLine(v bool)`

SetProductLine sets ProductLine field to given value.

### HasProductLine

`func (o *ProductMeta) HasProductLine() bool`

HasProductLine returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProductMeta) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductMeta) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductMeta) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProductMeta) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


