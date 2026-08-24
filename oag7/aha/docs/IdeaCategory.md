# IdeaCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewIdeaCategory

`func NewIdeaCategory() *IdeaCategory`

NewIdeaCategory instantiates a new IdeaCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdeaCategoryWithDefaults

`func NewIdeaCategoryWithDefaults() *IdeaCategory`

NewIdeaCategoryWithDefaults instantiates a new IdeaCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdeaCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdeaCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdeaCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdeaCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdeaCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdeaCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdeaCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdeaCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentId

`func (o *IdeaCategory) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *IdeaCategory) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *IdeaCategory) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *IdeaCategory) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *IdeaCategory) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *IdeaCategory) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetProjectId

`func (o *IdeaCategory) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *IdeaCategory) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *IdeaCategory) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *IdeaCategory) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IdeaCategory) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdeaCategory) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdeaCategory) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdeaCategory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


