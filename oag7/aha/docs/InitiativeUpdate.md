# InitiativeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**WorkflowStatus** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**Effort** | Pointer to **float64** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Presented** | Pointer to **bool** |  | [optional] 
**ProgressSource** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 

## Methods

### NewInitiativeUpdate

`func NewInitiativeUpdate() *InitiativeUpdate`

NewInitiativeUpdate instantiates a new InitiativeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiativeUpdateWithDefaults

`func NewInitiativeUpdateWithDefaults() *InitiativeUpdate`

NewInitiativeUpdateWithDefaults instantiates a new InitiativeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InitiativeUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InitiativeUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InitiativeUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InitiativeUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InitiativeUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InitiativeUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InitiativeUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InitiativeUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *InitiativeUpdate) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *InitiativeUpdate) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *InitiativeUpdate) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *InitiativeUpdate) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *InitiativeUpdate) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InitiativeUpdate) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InitiativeUpdate) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InitiativeUpdate) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *InitiativeUpdate) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *InitiativeUpdate) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *InitiativeUpdate) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InitiativeUpdate) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InitiativeUpdate) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InitiativeUpdate) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *InitiativeUpdate) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *InitiativeUpdate) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetValue

`func (o *InitiativeUpdate) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InitiativeUpdate) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InitiativeUpdate) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *InitiativeUpdate) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEffort

`func (o *InitiativeUpdate) GetEffort() float64`

GetEffort returns the Effort field if non-nil, zero value otherwise.

### GetEffortOk

`func (o *InitiativeUpdate) GetEffortOk() (*float64, bool)`

GetEffortOk returns a tuple with the Effort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffort

`func (o *InitiativeUpdate) SetEffort(v float64)`

SetEffort sets Effort field to given value.

### HasEffort

`func (o *InitiativeUpdate) HasEffort() bool`

HasEffort returns a boolean if a field has been set.

### GetColor

`func (o *InitiativeUpdate) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *InitiativeUpdate) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *InitiativeUpdate) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *InitiativeUpdate) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPresented

`func (o *InitiativeUpdate) GetPresented() bool`

GetPresented returns the Presented field if non-nil, zero value otherwise.

### GetPresentedOk

`func (o *InitiativeUpdate) GetPresentedOk() (*bool, bool)`

GetPresentedOk returns a tuple with the Presented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresented

`func (o *InitiativeUpdate) SetPresented(v bool)`

SetPresented sets Presented field to given value.

### HasPresented

`func (o *InitiativeUpdate) HasPresented() bool`

HasPresented returns a boolean if a field has been set.

### GetProgressSource

`func (o *InitiativeUpdate) GetProgressSource() string`

GetProgressSource returns the ProgressSource field if non-nil, zero value otherwise.

### GetProgressSourceOk

`func (o *InitiativeUpdate) GetProgressSourceOk() (*string, bool)`

GetProgressSourceOk returns a tuple with the ProgressSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressSource

`func (o *InitiativeUpdate) SetProgressSource(v string)`

SetProgressSource sets ProgressSource field to given value.

### HasProgressSource

`func (o *InitiativeUpdate) HasProgressSource() bool`

HasProgressSource returns a boolean if a field has been set.

### GetProgress

`func (o *InitiativeUpdate) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *InitiativeUpdate) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *InitiativeUpdate) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *InitiativeUpdate) HasProgress() bool`

HasProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


