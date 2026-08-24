# FeatureUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**WorkflowStatus** | Pointer to **string** |  | [optional] 
**AssignedToUser** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**DueDate** | Pointer to **NullableString** |  | [optional] 
**Release** | Pointer to **string** | Release ID to move feature to | [optional] 
**OriginalEstimateText** | Pointer to **string** |  | [optional] 
**RemainingEstimateText** | Pointer to **string** |  | [optional] 
**Initiative** | Pointer to **string** |  | [optional] 
**ReleasePhase** | Pointer to **string** |  | [optional] 
**ProgressSource** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**Epic** | Pointer to **string** | Epic ID or name | [optional] 

## Methods

### NewFeatureUpdate

`func NewFeatureUpdate() *FeatureUpdate`

NewFeatureUpdate instantiates a new FeatureUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureUpdateWithDefaults

`func NewFeatureUpdateWithDefaults() *FeatureUpdate`

NewFeatureUpdateWithDefaults instantiates a new FeatureUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeatureUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FeatureUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FeatureUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeatureUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *FeatureUpdate) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *FeatureUpdate) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *FeatureUpdate) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *FeatureUpdate) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetAssignedToUser

`func (o *FeatureUpdate) GetAssignedToUser() string`

GetAssignedToUser returns the AssignedToUser field if non-nil, zero value otherwise.

### GetAssignedToUserOk

`func (o *FeatureUpdate) GetAssignedToUserOk() (*string, bool)`

GetAssignedToUserOk returns a tuple with the AssignedToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToUser

`func (o *FeatureUpdate) SetAssignedToUser(v string)`

SetAssignedToUser sets AssignedToUser field to given value.

### HasAssignedToUser

`func (o *FeatureUpdate) HasAssignedToUser() bool`

HasAssignedToUser returns a boolean if a field has been set.

### GetTags

`func (o *FeatureUpdate) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeatureUpdate) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeatureUpdate) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeatureUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStartDate

`func (o *FeatureUpdate) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FeatureUpdate) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FeatureUpdate) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FeatureUpdate) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *FeatureUpdate) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *FeatureUpdate) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDueDate

`func (o *FeatureUpdate) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *FeatureUpdate) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *FeatureUpdate) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *FeatureUpdate) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *FeatureUpdate) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *FeatureUpdate) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetRelease

`func (o *FeatureUpdate) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *FeatureUpdate) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *FeatureUpdate) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *FeatureUpdate) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetOriginalEstimateText

`func (o *FeatureUpdate) GetOriginalEstimateText() string`

GetOriginalEstimateText returns the OriginalEstimateText field if non-nil, zero value otherwise.

### GetOriginalEstimateTextOk

`func (o *FeatureUpdate) GetOriginalEstimateTextOk() (*string, bool)`

GetOriginalEstimateTextOk returns a tuple with the OriginalEstimateText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEstimateText

`func (o *FeatureUpdate) SetOriginalEstimateText(v string)`

SetOriginalEstimateText sets OriginalEstimateText field to given value.

### HasOriginalEstimateText

`func (o *FeatureUpdate) HasOriginalEstimateText() bool`

HasOriginalEstimateText returns a boolean if a field has been set.

### GetRemainingEstimateText

`func (o *FeatureUpdate) GetRemainingEstimateText() string`

GetRemainingEstimateText returns the RemainingEstimateText field if non-nil, zero value otherwise.

### GetRemainingEstimateTextOk

`func (o *FeatureUpdate) GetRemainingEstimateTextOk() (*string, bool)`

GetRemainingEstimateTextOk returns a tuple with the RemainingEstimateText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingEstimateText

`func (o *FeatureUpdate) SetRemainingEstimateText(v string)`

SetRemainingEstimateText sets RemainingEstimateText field to given value.

### HasRemainingEstimateText

`func (o *FeatureUpdate) HasRemainingEstimateText() bool`

HasRemainingEstimateText returns a boolean if a field has been set.

### GetInitiative

`func (o *FeatureUpdate) GetInitiative() string`

GetInitiative returns the Initiative field if non-nil, zero value otherwise.

### GetInitiativeOk

`func (o *FeatureUpdate) GetInitiativeOk() (*string, bool)`

GetInitiativeOk returns a tuple with the Initiative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiative

`func (o *FeatureUpdate) SetInitiative(v string)`

SetInitiative sets Initiative field to given value.

### HasInitiative

`func (o *FeatureUpdate) HasInitiative() bool`

HasInitiative returns a boolean if a field has been set.

### GetReleasePhase

`func (o *FeatureUpdate) GetReleasePhase() string`

GetReleasePhase returns the ReleasePhase field if non-nil, zero value otherwise.

### GetReleasePhaseOk

`func (o *FeatureUpdate) GetReleasePhaseOk() (*string, bool)`

GetReleasePhaseOk returns a tuple with the ReleasePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasePhase

`func (o *FeatureUpdate) SetReleasePhase(v string)`

SetReleasePhase sets ReleasePhase field to given value.

### HasReleasePhase

`func (o *FeatureUpdate) HasReleasePhase() bool`

HasReleasePhase returns a boolean if a field has been set.

### GetProgressSource

`func (o *FeatureUpdate) GetProgressSource() string`

GetProgressSource returns the ProgressSource field if non-nil, zero value otherwise.

### GetProgressSourceOk

`func (o *FeatureUpdate) GetProgressSourceOk() (*string, bool)`

GetProgressSourceOk returns a tuple with the ProgressSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressSource

`func (o *FeatureUpdate) SetProgressSource(v string)`

SetProgressSource sets ProgressSource field to given value.

### HasProgressSource

`func (o *FeatureUpdate) HasProgressSource() bool`

HasProgressSource returns a boolean if a field has been set.

### GetProgress

`func (o *FeatureUpdate) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *FeatureUpdate) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *FeatureUpdate) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *FeatureUpdate) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetEpic

`func (o *FeatureUpdate) GetEpic() string`

GetEpic returns the Epic field if non-nil, zero value otherwise.

### GetEpicOk

`func (o *FeatureUpdate) GetEpicOk() (*string, bool)`

GetEpicOk returns a tuple with the Epic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpic

`func (o *FeatureUpdate) SetEpic(v string)`

SetEpic sets Epic field to given value.

### HasEpic

`func (o *FeatureUpdate) HasEpic() bool`

HasEpic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


