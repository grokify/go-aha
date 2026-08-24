# StrategicModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ReferenceNum** | **string** |  | 
**Name** | **string** |  | 
**Kind** | **string** | Type of canvas (e.g., Opportunity, Lean Canvas, Business Model) | 
**Description** | Pointer to [**DescriptionObject**](DescriptionObject.md) |  | [optional] 
**Url** | Pointer to **string** | Web UI URL | [optional] 
**Resource** | Pointer to **string** | API URL | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Project** | Pointer to [**ProductMeta**](ProductMeta.md) |  | [optional] 
**Components** | Pointer to [**[]StrategicModelComponent**](StrategicModelComponent.md) |  | [optional] 

## Methods

### NewStrategicModel

`func NewStrategicModel(id string, referenceNum string, name string, kind string, createdAt time.Time, ) *StrategicModel`

NewStrategicModel instantiates a new StrategicModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategicModelWithDefaults

`func NewStrategicModelWithDefaults() *StrategicModel`

NewStrategicModelWithDefaults instantiates a new StrategicModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StrategicModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StrategicModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StrategicModel) SetId(v string)`

SetId sets Id field to given value.


### GetReferenceNum

`func (o *StrategicModel) GetReferenceNum() string`

GetReferenceNum returns the ReferenceNum field if non-nil, zero value otherwise.

### GetReferenceNumOk

`func (o *StrategicModel) GetReferenceNumOk() (*string, bool)`

GetReferenceNumOk returns a tuple with the ReferenceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNum

`func (o *StrategicModel) SetReferenceNum(v string)`

SetReferenceNum sets ReferenceNum field to given value.


### GetName

`func (o *StrategicModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StrategicModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StrategicModel) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *StrategicModel) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StrategicModel) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StrategicModel) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDescription

`func (o *StrategicModel) GetDescription() DescriptionObject`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StrategicModel) GetDescriptionOk() (*DescriptionObject, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StrategicModel) SetDescription(v DescriptionObject)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StrategicModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *StrategicModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StrategicModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StrategicModel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StrategicModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetResource

`func (o *StrategicModel) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *StrategicModel) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *StrategicModel) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *StrategicModel) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StrategicModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StrategicModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StrategicModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *StrategicModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StrategicModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StrategicModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StrategicModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProject

`func (o *StrategicModel) GetProject() ProductMeta`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *StrategicModel) GetProjectOk() (*ProductMeta, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *StrategicModel) SetProject(v ProductMeta)`

SetProject sets Project field to given value.

### HasProject

`func (o *StrategicModel) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetComponents

`func (o *StrategicModel) GetComponents() []StrategicModelComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *StrategicModel) GetComponentsOk() (*[]StrategicModelComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *StrategicModel) SetComponents(v []StrategicModelComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *StrategicModel) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


