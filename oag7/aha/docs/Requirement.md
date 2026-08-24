# Requirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ReferenceNum** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to [**DescriptionObject**](DescriptionObject.md) |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**OriginalEstimate** | Pointer to **NullableFloat32** |  | [optional] 
**RemainingEstimate** | Pointer to **NullableFloat32** |  | [optional] 
**WorkDone** | Pointer to **NullableFloat32** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** | Web UI URL | [optional] 
**Resource** | Pointer to **string** | API URL | [optional] 
**WorkflowStatus** | Pointer to [**WorkflowStatus**](WorkflowStatus.md) |  | [optional] 
**AssignedToUser** | Pointer to [**NullableUser**](User.md) |  | [optional] 
**Feature** | Pointer to [**FeatureMeta**](FeatureMeta.md) |  | [optional] 

## Methods

### NewRequirement

`func NewRequirement(id string, referenceNum string, name string, createdAt time.Time, ) *Requirement`

NewRequirement instantiates a new Requirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementWithDefaults

`func NewRequirementWithDefaults() *Requirement`

NewRequirementWithDefaults instantiates a new Requirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Requirement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Requirement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Requirement) SetId(v string)`

SetId sets Id field to given value.


### GetReferenceNum

`func (o *Requirement) GetReferenceNum() string`

GetReferenceNum returns the ReferenceNum field if non-nil, zero value otherwise.

### GetReferenceNumOk

`func (o *Requirement) GetReferenceNumOk() (*string, bool)`

GetReferenceNumOk returns a tuple with the ReferenceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNum

`func (o *Requirement) SetReferenceNum(v string)`

SetReferenceNum sets ReferenceNum field to given value.


### GetName

`func (o *Requirement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Requirement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Requirement) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Requirement) GetDescription() DescriptionObject`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Requirement) GetDescriptionOk() (*DescriptionObject, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Requirement) SetDescription(v DescriptionObject)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Requirement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPosition

`func (o *Requirement) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Requirement) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Requirement) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Requirement) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetOriginalEstimate

`func (o *Requirement) GetOriginalEstimate() float32`

GetOriginalEstimate returns the OriginalEstimate field if non-nil, zero value otherwise.

### GetOriginalEstimateOk

`func (o *Requirement) GetOriginalEstimateOk() (*float32, bool)`

GetOriginalEstimateOk returns a tuple with the OriginalEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEstimate

`func (o *Requirement) SetOriginalEstimate(v float32)`

SetOriginalEstimate sets OriginalEstimate field to given value.

### HasOriginalEstimate

`func (o *Requirement) HasOriginalEstimate() bool`

HasOriginalEstimate returns a boolean if a field has been set.

### SetOriginalEstimateNil

`func (o *Requirement) SetOriginalEstimateNil(b bool)`

 SetOriginalEstimateNil sets the value for OriginalEstimate to be an explicit nil

### UnsetOriginalEstimate
`func (o *Requirement) UnsetOriginalEstimate()`

UnsetOriginalEstimate ensures that no value is present for OriginalEstimate, not even an explicit nil
### GetRemainingEstimate

`func (o *Requirement) GetRemainingEstimate() float32`

GetRemainingEstimate returns the RemainingEstimate field if non-nil, zero value otherwise.

### GetRemainingEstimateOk

`func (o *Requirement) GetRemainingEstimateOk() (*float32, bool)`

GetRemainingEstimateOk returns a tuple with the RemainingEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingEstimate

`func (o *Requirement) SetRemainingEstimate(v float32)`

SetRemainingEstimate sets RemainingEstimate field to given value.

### HasRemainingEstimate

`func (o *Requirement) HasRemainingEstimate() bool`

HasRemainingEstimate returns a boolean if a field has been set.

### SetRemainingEstimateNil

`func (o *Requirement) SetRemainingEstimateNil(b bool)`

 SetRemainingEstimateNil sets the value for RemainingEstimate to be an explicit nil

### UnsetRemainingEstimate
`func (o *Requirement) UnsetRemainingEstimate()`

UnsetRemainingEstimate ensures that no value is present for RemainingEstimate, not even an explicit nil
### GetWorkDone

`func (o *Requirement) GetWorkDone() float32`

GetWorkDone returns the WorkDone field if non-nil, zero value otherwise.

### GetWorkDoneOk

`func (o *Requirement) GetWorkDoneOk() (*float32, bool)`

GetWorkDoneOk returns a tuple with the WorkDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkDone

`func (o *Requirement) SetWorkDone(v float32)`

SetWorkDone sets WorkDone field to given value.

### HasWorkDone

`func (o *Requirement) HasWorkDone() bool`

HasWorkDone returns a boolean if a field has been set.

### SetWorkDoneNil

`func (o *Requirement) SetWorkDoneNil(b bool)`

 SetWorkDoneNil sets the value for WorkDone to be an explicit nil

### UnsetWorkDone
`func (o *Requirement) UnsetWorkDone()`

UnsetWorkDone ensures that no value is present for WorkDone, not even an explicit nil
### GetCreatedAt

`func (o *Requirement) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Requirement) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Requirement) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Requirement) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Requirement) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Requirement) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Requirement) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *Requirement) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Requirement) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Requirement) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Requirement) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetResource

`func (o *Requirement) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Requirement) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Requirement) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Requirement) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *Requirement) GetWorkflowStatus() WorkflowStatus`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *Requirement) GetWorkflowStatusOk() (*WorkflowStatus, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *Requirement) SetWorkflowStatus(v WorkflowStatus)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *Requirement) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetAssignedToUser

`func (o *Requirement) GetAssignedToUser() User`

GetAssignedToUser returns the AssignedToUser field if non-nil, zero value otherwise.

### GetAssignedToUserOk

`func (o *Requirement) GetAssignedToUserOk() (*User, bool)`

GetAssignedToUserOk returns a tuple with the AssignedToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToUser

`func (o *Requirement) SetAssignedToUser(v User)`

SetAssignedToUser sets AssignedToUser field to given value.

### HasAssignedToUser

`func (o *Requirement) HasAssignedToUser() bool`

HasAssignedToUser returns a boolean if a field has been set.

### SetAssignedToUserNil

`func (o *Requirement) SetAssignedToUserNil(b bool)`

 SetAssignedToUserNil sets the value for AssignedToUser to be an explicit nil

### UnsetAssignedToUser
`func (o *Requirement) UnsetAssignedToUser()`

UnsetAssignedToUser ensures that no value is present for AssignedToUser, not even an explicit nil
### GetFeature

`func (o *Requirement) GetFeature() FeatureMeta`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *Requirement) GetFeatureOk() (*FeatureMeta, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *Requirement) SetFeature(v FeatureMeta)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *Requirement) HasFeature() bool`

HasFeature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


