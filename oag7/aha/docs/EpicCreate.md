# EpicCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Epic name | 
**Description** | Pointer to **string** | Epic description (HTML allowed) | [optional] 
**WorkflowStatus** | Pointer to **string** | Workflow status ID or name | [optional] 
**StartDate** | Pointer to **NullableString** | Epic start date | [optional] 
**DueDate** | Pointer to **NullableString** | Epic due date | [optional] 
**Color** | Pointer to **string** | Epic color | [optional] 
**Initiative** | Pointer to **string** | Initiative ID or reference | [optional] 

## Methods

### NewEpicCreate

`func NewEpicCreate(name string, ) *EpicCreate`

NewEpicCreate instantiates a new EpicCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpicCreateWithDefaults

`func NewEpicCreateWithDefaults() *EpicCreate`

NewEpicCreateWithDefaults instantiates a new EpicCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EpicCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EpicCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EpicCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EpicCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EpicCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EpicCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EpicCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *EpicCreate) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *EpicCreate) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *EpicCreate) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *EpicCreate) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *EpicCreate) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EpicCreate) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EpicCreate) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EpicCreate) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *EpicCreate) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *EpicCreate) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDueDate

`func (o *EpicCreate) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *EpicCreate) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *EpicCreate) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *EpicCreate) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *EpicCreate) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *EpicCreate) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetColor

`func (o *EpicCreate) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *EpicCreate) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *EpicCreate) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *EpicCreate) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetInitiative

`func (o *EpicCreate) GetInitiative() string`

GetInitiative returns the Initiative field if non-nil, zero value otherwise.

### GetInitiativeOk

`func (o *EpicCreate) GetInitiativeOk() (*string, bool)`

GetInitiativeOk returns a tuple with the Initiative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiative

`func (o *EpicCreate) SetInitiative(v string)`

SetInitiative sets Initiative field to given value.

### HasInitiative

`func (o *EpicCreate) HasInitiative() bool`

HasInitiative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


