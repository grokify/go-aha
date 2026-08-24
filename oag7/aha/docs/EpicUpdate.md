# EpicUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Epic name | [optional] 
**Description** | Pointer to **string** | Epic description (HTML allowed) | [optional] 
**WorkflowStatus** | Pointer to **string** | Workflow status ID or name | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**DueDate** | Pointer to **NullableString** |  | [optional] 
**Progress** | Pointer to **float32** | Progress percentage (0-100) when progress_source is manual | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Initiative** | Pointer to **string** | Initiative ID or reference | [optional] 

## Methods

### NewEpicUpdate

`func NewEpicUpdate() *EpicUpdate`

NewEpicUpdate instantiates a new EpicUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpicUpdateWithDefaults

`func NewEpicUpdateWithDefaults() *EpicUpdate`

NewEpicUpdateWithDefaults instantiates a new EpicUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EpicUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EpicUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EpicUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EpicUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EpicUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EpicUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EpicUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EpicUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *EpicUpdate) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *EpicUpdate) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *EpicUpdate) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *EpicUpdate) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *EpicUpdate) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EpicUpdate) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EpicUpdate) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EpicUpdate) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *EpicUpdate) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *EpicUpdate) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDueDate

`func (o *EpicUpdate) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *EpicUpdate) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *EpicUpdate) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *EpicUpdate) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *EpicUpdate) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *EpicUpdate) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetProgress

`func (o *EpicUpdate) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *EpicUpdate) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *EpicUpdate) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *EpicUpdate) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetColor

`func (o *EpicUpdate) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *EpicUpdate) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *EpicUpdate) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *EpicUpdate) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetInitiative

`func (o *EpicUpdate) GetInitiative() string`

GetInitiative returns the Initiative field if non-nil, zero value otherwise.

### GetInitiativeOk

`func (o *EpicUpdate) GetInitiativeOk() (*string, bool)`

GetInitiativeOk returns a tuple with the Initiative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiative

`func (o *EpicUpdate) SetInitiative(v string)`

SetInitiative sets Initiative field to given value.

### HasInitiative

`func (o *EpicUpdate) HasInitiative() bool`

HasInitiative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


