# FeatureCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Feature name | 
**Description** | Pointer to **string** | Feature description (HTML allowed) | [optional] 
**WorkflowStatus** | Pointer to **string** | Workflow status ID or name | [optional] 
**AssignedToUser** | Pointer to **string** | User email to assign | [optional] 
**Tags** | Pointer to **string** | Comma-separated tags | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**DueDate** | Pointer to **NullableString** |  | [optional] 
**OriginalEstimateText** | Pointer to **string** | Effort estimate (e.g., 2d, 4h, 30min, 5p) | [optional] 
**Initiative** | Pointer to **string** | Initiative ID or name | [optional] 

## Methods

### NewFeatureCreate

`func NewFeatureCreate(name string, ) *FeatureCreate`

NewFeatureCreate instantiates a new FeatureCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureCreateWithDefaults

`func NewFeatureCreateWithDefaults() *FeatureCreate`

NewFeatureCreateWithDefaults instantiates a new FeatureCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeatureCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FeatureCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeatureCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *FeatureCreate) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *FeatureCreate) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *FeatureCreate) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *FeatureCreate) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetAssignedToUser

`func (o *FeatureCreate) GetAssignedToUser() string`

GetAssignedToUser returns the AssignedToUser field if non-nil, zero value otherwise.

### GetAssignedToUserOk

`func (o *FeatureCreate) GetAssignedToUserOk() (*string, bool)`

GetAssignedToUserOk returns a tuple with the AssignedToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToUser

`func (o *FeatureCreate) SetAssignedToUser(v string)`

SetAssignedToUser sets AssignedToUser field to given value.

### HasAssignedToUser

`func (o *FeatureCreate) HasAssignedToUser() bool`

HasAssignedToUser returns a boolean if a field has been set.

### GetTags

`func (o *FeatureCreate) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeatureCreate) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeatureCreate) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeatureCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStartDate

`func (o *FeatureCreate) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FeatureCreate) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FeatureCreate) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FeatureCreate) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *FeatureCreate) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *FeatureCreate) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDueDate

`func (o *FeatureCreate) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *FeatureCreate) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *FeatureCreate) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *FeatureCreate) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *FeatureCreate) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *FeatureCreate) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetOriginalEstimateText

`func (o *FeatureCreate) GetOriginalEstimateText() string`

GetOriginalEstimateText returns the OriginalEstimateText field if non-nil, zero value otherwise.

### GetOriginalEstimateTextOk

`func (o *FeatureCreate) GetOriginalEstimateTextOk() (*string, bool)`

GetOriginalEstimateTextOk returns a tuple with the OriginalEstimateText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEstimateText

`func (o *FeatureCreate) SetOriginalEstimateText(v string)`

SetOriginalEstimateText sets OriginalEstimateText field to given value.

### HasOriginalEstimateText

`func (o *FeatureCreate) HasOriginalEstimateText() bool`

HasOriginalEstimateText returns a boolean if a field has been set.

### GetInitiative

`func (o *FeatureCreate) GetInitiative() string`

GetInitiative returns the Initiative field if non-nil, zero value otherwise.

### GetInitiativeOk

`func (o *FeatureCreate) GetInitiativeOk() (*string, bool)`

GetInitiativeOk returns a tuple with the Initiative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiative

`func (o *FeatureCreate) SetInitiative(v string)`

SetInitiative sets Initiative field to given value.

### HasInitiative

`func (o *FeatureCreate) HasInitiative() bool`

HasInitiative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


