# Epic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ReferenceNum** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to [**DescriptionObject**](DescriptionObject.md) |  | [optional] 
**Progress** | Pointer to **NullableFloat32** | Progress percentage (0-100) | [optional] 
**ProgressSource** | Pointer to **string** | Source of progress calculation | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**DueDate** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** | Web UI URL | [optional] 
**Resource** | Pointer to **string** | API URL | [optional] 
**CommentsCount** | Pointer to **int64** |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**WorkflowStatus** | Pointer to [**WorkflowStatus**](WorkflowStatus.md) |  | [optional] 
**Release** | Pointer to [**Release**](Release.md) |  | [optional] 
**Initiative** | Pointer to [**InitiativeMeta**](InitiativeMeta.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEpic

`func NewEpic(id string, referenceNum string, name string, createdAt time.Time, ) *Epic`

NewEpic instantiates a new Epic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpicWithDefaults

`func NewEpicWithDefaults() *Epic`

NewEpicWithDefaults instantiates a new Epic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Epic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Epic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Epic) SetId(v string)`

SetId sets Id field to given value.


### GetReferenceNum

`func (o *Epic) GetReferenceNum() string`

GetReferenceNum returns the ReferenceNum field if non-nil, zero value otherwise.

### GetReferenceNumOk

`func (o *Epic) GetReferenceNumOk() (*string, bool)`

GetReferenceNumOk returns a tuple with the ReferenceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNum

`func (o *Epic) SetReferenceNum(v string)`

SetReferenceNum sets ReferenceNum field to given value.


### GetName

`func (o *Epic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Epic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Epic) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Epic) GetDescription() DescriptionObject`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Epic) GetDescriptionOk() (*DescriptionObject, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Epic) SetDescription(v DescriptionObject)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Epic) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProgress

`func (o *Epic) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Epic) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Epic) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Epic) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgressNil

`func (o *Epic) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *Epic) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetProgressSource

`func (o *Epic) GetProgressSource() string`

GetProgressSource returns the ProgressSource field if non-nil, zero value otherwise.

### GetProgressSourceOk

`func (o *Epic) GetProgressSourceOk() (*string, bool)`

GetProgressSourceOk returns a tuple with the ProgressSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressSource

`func (o *Epic) SetProgressSource(v string)`

SetProgressSource sets ProgressSource field to given value.

### HasProgressSource

`func (o *Epic) HasProgressSource() bool`

HasProgressSource returns a boolean if a field has been set.

### GetStartDate

`func (o *Epic) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Epic) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Epic) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Epic) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *Epic) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Epic) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDueDate

`func (o *Epic) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Epic) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Epic) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Epic) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *Epic) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *Epic) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetCreatedAt

`func (o *Epic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Epic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Epic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Epic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Epic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Epic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Epic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *Epic) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Epic) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Epic) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Epic) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetResource

`func (o *Epic) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Epic) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Epic) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Epic) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetCommentsCount

`func (o *Epic) GetCommentsCount() int64`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *Epic) GetCommentsCountOk() (*int64, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *Epic) SetCommentsCount(v int64)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *Epic) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetPosition

`func (o *Epic) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Epic) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Epic) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Epic) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetColor

`func (o *Epic) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Epic) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Epic) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Epic) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *Epic) GetWorkflowStatus() WorkflowStatus`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *Epic) GetWorkflowStatusOk() (*WorkflowStatus, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *Epic) SetWorkflowStatus(v WorkflowStatus)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *Epic) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetRelease

`func (o *Epic) GetRelease() Release`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *Epic) GetReleaseOk() (*Release, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *Epic) SetRelease(v Release)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *Epic) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetInitiative

`func (o *Epic) GetInitiative() InitiativeMeta`

GetInitiative returns the Initiative field if non-nil, zero value otherwise.

### GetInitiativeOk

`func (o *Epic) GetInitiativeOk() (*InitiativeMeta, bool)`

GetInitiativeOk returns a tuple with the Initiative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiative

`func (o *Epic) SetInitiative(v InitiativeMeta)`

SetInitiative sets Initiative field to given value.

### HasInitiative

`func (o *Epic) HasInitiative() bool`

HasInitiative returns a boolean if a field has been set.

### GetTags

`func (o *Epic) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Epic) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Epic) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Epic) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


