# ProductCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the product | 
**ReferencePrefix** | **string** | Abbreviation used as prefix on all features | 
**Description** | Pointer to **string** | Description of the product (HTML allowed) | [optional] 
**ParentId** | Pointer to **string** | Numeric ID or prefix of parent product line | [optional] 
**WorkspaceType** | Pointer to **string** | Type of workspace (product_workspace, it_workspace, marketing_workspace, etc.) | [optional] 
**ProductLine** | Pointer to **bool** | Set to true to create a product line | [optional] 
**ProductLineType** | Pointer to **string** | Required if creating a product line | [optional] 

## Methods

### NewProductCreate

`func NewProductCreate(name string, referencePrefix string, ) *ProductCreate`

NewProductCreate instantiates a new ProductCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCreateWithDefaults

`func NewProductCreateWithDefaults() *ProductCreate`

NewProductCreateWithDefaults instantiates a new ProductCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductCreate) SetName(v string)`

SetName sets Name field to given value.


### GetReferencePrefix

`func (o *ProductCreate) GetReferencePrefix() string`

GetReferencePrefix returns the ReferencePrefix field if non-nil, zero value otherwise.

### GetReferencePrefixOk

`func (o *ProductCreate) GetReferencePrefixOk() (*string, bool)`

GetReferencePrefixOk returns a tuple with the ReferencePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencePrefix

`func (o *ProductCreate) SetReferencePrefix(v string)`

SetReferencePrefix sets ReferencePrefix field to given value.


### GetDescription

`func (o *ProductCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentId

`func (o *ProductCreate) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ProductCreate) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ProductCreate) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ProductCreate) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetWorkspaceType

`func (o *ProductCreate) GetWorkspaceType() string`

GetWorkspaceType returns the WorkspaceType field if non-nil, zero value otherwise.

### GetWorkspaceTypeOk

`func (o *ProductCreate) GetWorkspaceTypeOk() (*string, bool)`

GetWorkspaceTypeOk returns a tuple with the WorkspaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceType

`func (o *ProductCreate) SetWorkspaceType(v string)`

SetWorkspaceType sets WorkspaceType field to given value.

### HasWorkspaceType

`func (o *ProductCreate) HasWorkspaceType() bool`

HasWorkspaceType returns a boolean if a field has been set.

### GetProductLine

`func (o *ProductCreate) GetProductLine() bool`

GetProductLine returns the ProductLine field if non-nil, zero value otherwise.

### GetProductLineOk

`func (o *ProductCreate) GetProductLineOk() (*bool, bool)`

GetProductLineOk returns a tuple with the ProductLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductLine

`func (o *ProductCreate) SetProductLine(v bool)`

SetProductLine sets ProductLine field to given value.

### HasProductLine

`func (o *ProductCreate) HasProductLine() bool`

HasProductLine returns a boolean if a field has been set.

### GetProductLineType

`func (o *ProductCreate) GetProductLineType() string`

GetProductLineType returns the ProductLineType field if non-nil, zero value otherwise.

### GetProductLineTypeOk

`func (o *ProductCreate) GetProductLineTypeOk() (*string, bool)`

GetProductLineTypeOk returns a tuple with the ProductLineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductLineType

`func (o *ProductCreate) SetProductLineType(v string)`

SetProductLineType sets ProductLineType field to given value.

### HasProductLineType

`func (o *ProductCreate) HasProductLineType() bool`

HasProductLineType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


