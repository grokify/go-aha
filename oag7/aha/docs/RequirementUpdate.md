# RequirementUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Requirement name | [optional] 
**Description** | Pointer to **string** | Requirement description (HTML allowed) | [optional] 
**WorkflowStatus** | Pointer to **string** | Workflow status ID or name | [optional] 
**AssignedToUser** | Pointer to **string** | User ID or email | [optional] 
**OriginalEstimate** | Pointer to **float32** |  | [optional] 
**RemainingEstimate** | Pointer to **float32** |  | [optional] 
**WorkDone** | Pointer to **float32** |  | [optional] 

## Methods

### NewRequirementUpdate

`func NewRequirementUpdate() *RequirementUpdate`

NewRequirementUpdate instantiates a new RequirementUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementUpdateWithDefaults

`func NewRequirementUpdateWithDefaults() *RequirementUpdate`

NewRequirementUpdateWithDefaults instantiates a new RequirementUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RequirementUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequirementUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequirementUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequirementUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RequirementUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequirementUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequirementUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequirementUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *RequirementUpdate) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *RequirementUpdate) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *RequirementUpdate) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *RequirementUpdate) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetAssignedToUser

`func (o *RequirementUpdate) GetAssignedToUser() string`

GetAssignedToUser returns the AssignedToUser field if non-nil, zero value otherwise.

### GetAssignedToUserOk

`func (o *RequirementUpdate) GetAssignedToUserOk() (*string, bool)`

GetAssignedToUserOk returns a tuple with the AssignedToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToUser

`func (o *RequirementUpdate) SetAssignedToUser(v string)`

SetAssignedToUser sets AssignedToUser field to given value.

### HasAssignedToUser

`func (o *RequirementUpdate) HasAssignedToUser() bool`

HasAssignedToUser returns a boolean if a field has been set.

### GetOriginalEstimate

`func (o *RequirementUpdate) GetOriginalEstimate() float32`

GetOriginalEstimate returns the OriginalEstimate field if non-nil, zero value otherwise.

### GetOriginalEstimateOk

`func (o *RequirementUpdate) GetOriginalEstimateOk() (*float32, bool)`

GetOriginalEstimateOk returns a tuple with the OriginalEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEstimate

`func (o *RequirementUpdate) SetOriginalEstimate(v float32)`

SetOriginalEstimate sets OriginalEstimate field to given value.

### HasOriginalEstimate

`func (o *RequirementUpdate) HasOriginalEstimate() bool`

HasOriginalEstimate returns a boolean if a field has been set.

### GetRemainingEstimate

`func (o *RequirementUpdate) GetRemainingEstimate() float32`

GetRemainingEstimate returns the RemainingEstimate field if non-nil, zero value otherwise.

### GetRemainingEstimateOk

`func (o *RequirementUpdate) GetRemainingEstimateOk() (*float32, bool)`

GetRemainingEstimateOk returns a tuple with the RemainingEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingEstimate

`func (o *RequirementUpdate) SetRemainingEstimate(v float32)`

SetRemainingEstimate sets RemainingEstimate field to given value.

### HasRemainingEstimate

`func (o *RequirementUpdate) HasRemainingEstimate() bool`

HasRemainingEstimate returns a boolean if a field has been set.

### GetWorkDone

`func (o *RequirementUpdate) GetWorkDone() float32`

GetWorkDone returns the WorkDone field if non-nil, zero value otherwise.

### GetWorkDoneOk

`func (o *RequirementUpdate) GetWorkDoneOk() (*float32, bool)`

GetWorkDoneOk returns a tuple with the WorkDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkDone

`func (o *RequirementUpdate) SetWorkDone(v float32)`

SetWorkDone sets WorkDone field to given value.

### HasWorkDone

`func (o *RequirementUpdate) HasWorkDone() bool`

HasWorkDone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


