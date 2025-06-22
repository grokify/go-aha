# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the product. | [optional] 
**ReferencePrefix** | Pointer to **string** | The reference prefix slug for the product. | [optional] 
**Name** | Pointer to **string** | The name for the product. | [optional] 
**ProductLine** | Pointer to **bool** | Whether the product is a product line or not. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date-time when the product was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date-time when the product was updated. | [optional] 
**Url** | Pointer to **string** | User URL for the project. | [optional] 
**Resource** | Pointer to **string** | User URL for the project. | [optional] 
**HasIdeas** | Pointer to **bool** | Whether the product has ideas or not. | [optional] 
**HasMasterFeatures** | Pointer to **bool** | Whether the product has master features or not. | [optional] 

## Methods

### NewProduct

`func NewProduct() *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Product) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Product) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferencePrefix

`func (o *Product) GetReferencePrefix() string`

GetReferencePrefix returns the ReferencePrefix field if non-nil, zero value otherwise.

### GetReferencePrefixOk

`func (o *Product) GetReferencePrefixOk() (*string, bool)`

GetReferencePrefixOk returns a tuple with the ReferencePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencePrefix

`func (o *Product) SetReferencePrefix(v string)`

SetReferencePrefix sets ReferencePrefix field to given value.

### HasReferencePrefix

`func (o *Product) HasReferencePrefix() bool`

HasReferencePrefix returns a boolean if a field has been set.

### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Product) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductLine

`func (o *Product) GetProductLine() bool`

GetProductLine returns the ProductLine field if non-nil, zero value otherwise.

### GetProductLineOk

`func (o *Product) GetProductLineOk() (*bool, bool)`

GetProductLineOk returns a tuple with the ProductLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductLine

`func (o *Product) SetProductLine(v bool)`

SetProductLine sets ProductLine field to given value.

### HasProductLine

`func (o *Product) HasProductLine() bool`

HasProductLine returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Product) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Product) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Product) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Product) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Product) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Product) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Product) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Product) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *Product) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Product) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Product) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Product) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetResource

`func (o *Product) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Product) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Product) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Product) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetHasIdeas

`func (o *Product) GetHasIdeas() bool`

GetHasIdeas returns the HasIdeas field if non-nil, zero value otherwise.

### GetHasIdeasOk

`func (o *Product) GetHasIdeasOk() (*bool, bool)`

GetHasIdeasOk returns a tuple with the HasIdeas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIdeas

`func (o *Product) SetHasIdeas(v bool)`

SetHasIdeas sets HasIdeas field to given value.

### HasHasIdeas

`func (o *Product) HasHasIdeas() bool`

HasHasIdeas returns a boolean if a field has been set.

### GetHasMasterFeatures

`func (o *Product) GetHasMasterFeatures() bool`

GetHasMasterFeatures returns the HasMasterFeatures field if non-nil, zero value otherwise.

### GetHasMasterFeaturesOk

`func (o *Product) GetHasMasterFeaturesOk() (*bool, bool)`

GetHasMasterFeaturesOk returns a tuple with the HasMasterFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMasterFeatures

`func (o *Product) SetHasMasterFeatures(v bool)`

SetHasMasterFeatures sets HasMasterFeatures field to given value.

### HasHasMasterFeatures

`func (o *Product) HasHasMasterFeatures() bool`

HasHasMasterFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


