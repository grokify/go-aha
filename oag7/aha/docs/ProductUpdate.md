# ProductUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the product | [optional] 
**ReferencePrefix** | Pointer to **string** | Abbreviation used as prefix on all features | [optional] 
**Description** | Pointer to **string** | Description of the product (HTML allowed) | [optional] 
**ParentId** | Pointer to **string** | Numeric ID or prefix of parent product line | [optional] 
**WorkspaceType** | Pointer to **string** | Type of workspace | [optional] 

## Methods

### NewProductUpdate

`func NewProductUpdate() *ProductUpdate`

NewProductUpdate instantiates a new ProductUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductUpdateWithDefaults

`func NewProductUpdateWithDefaults() *ProductUpdate`

NewProductUpdateWithDefaults instantiates a new ProductUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferencePrefix

`func (o *ProductUpdate) GetReferencePrefix() string`

GetReferencePrefix returns the ReferencePrefix field if non-nil, zero value otherwise.

### GetReferencePrefixOk

`func (o *ProductUpdate) GetReferencePrefixOk() (*string, bool)`

GetReferencePrefixOk returns a tuple with the ReferencePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencePrefix

`func (o *ProductUpdate) SetReferencePrefix(v string)`

SetReferencePrefix sets ReferencePrefix field to given value.

### HasReferencePrefix

`func (o *ProductUpdate) HasReferencePrefix() bool`

HasReferencePrefix returns a boolean if a field has been set.

### GetDescription

`func (o *ProductUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentId

`func (o *ProductUpdate) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ProductUpdate) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ProductUpdate) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ProductUpdate) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetWorkspaceType

`func (o *ProductUpdate) GetWorkspaceType() string`

GetWorkspaceType returns the WorkspaceType field if non-nil, zero value otherwise.

### GetWorkspaceTypeOk

`func (o *ProductUpdate) GetWorkspaceTypeOk() (*string, bool)`

GetWorkspaceTypeOk returns a tuple with the WorkspaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceType

`func (o *ProductUpdate) SetWorkspaceType(v string)`

SetWorkspaceType sets WorkspaceType field to given value.

### HasWorkspaceType

`func (o *ProductUpdate) HasWorkspaceType() bool`

HasWorkspaceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


