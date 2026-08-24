# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ReferencePrefix** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to [**DescriptionObject**](DescriptionObject.md) |  | [optional] 
**ProductLine** | Pointer to **bool** |  | [optional] 
**ParentId** | Pointer to **string** | ID of the parent product line | [optional] 
**WorkspaceType** | Pointer to **string** | Type of workspace (product_workspace, it_workspace, marketing_workspace, etc.) | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**HasIdeas** | Pointer to **bool** |  | [optional] 
**HasMasterFeatures** | Pointer to **bool** |  | [optional] 

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

### GetDescription

`func (o *Product) GetDescription() DescriptionObject`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*DescriptionObject, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v DescriptionObject)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Product) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

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

### GetParentId

`func (o *Product) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Product) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Product) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Product) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetWorkspaceType

`func (o *Product) GetWorkspaceType() string`

GetWorkspaceType returns the WorkspaceType field if non-nil, zero value otherwise.

### GetWorkspaceTypeOk

`func (o *Product) GetWorkspaceTypeOk() (*string, bool)`

GetWorkspaceTypeOk returns a tuple with the WorkspaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceType

`func (o *Product) SetWorkspaceType(v string)`

SetWorkspaceType sets WorkspaceType field to given value.

### HasWorkspaceType

`func (o *Product) HasWorkspaceType() bool`

HasWorkspaceType returns a boolean if a field has been set.

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


