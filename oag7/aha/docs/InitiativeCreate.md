# InitiativeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Initiative name | 
**Description** | Pointer to **string** | Initiative description (HTML allowed) | [optional] 
**WorkflowStatus** | Pointer to **string** | Workflow status ID or name | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**Effort** | Pointer to **float64** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Presented** | Pointer to **bool** |  | [optional] 

## Methods

### NewInitiativeCreate

`func NewInitiativeCreate(name string, ) *InitiativeCreate`

NewInitiativeCreate instantiates a new InitiativeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiativeCreateWithDefaults

`func NewInitiativeCreateWithDefaults() *InitiativeCreate`

NewInitiativeCreateWithDefaults instantiates a new InitiativeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InitiativeCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InitiativeCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InitiativeCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InitiativeCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InitiativeCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InitiativeCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InitiativeCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *InitiativeCreate) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *InitiativeCreate) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *InitiativeCreate) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *InitiativeCreate) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *InitiativeCreate) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InitiativeCreate) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InitiativeCreate) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InitiativeCreate) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *InitiativeCreate) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *InitiativeCreate) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *InitiativeCreate) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InitiativeCreate) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InitiativeCreate) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InitiativeCreate) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *InitiativeCreate) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *InitiativeCreate) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetValue

`func (o *InitiativeCreate) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InitiativeCreate) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InitiativeCreate) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *InitiativeCreate) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEffort

`func (o *InitiativeCreate) GetEffort() float64`

GetEffort returns the Effort field if non-nil, zero value otherwise.

### GetEffortOk

`func (o *InitiativeCreate) GetEffortOk() (*float64, bool)`

GetEffortOk returns a tuple with the Effort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffort

`func (o *InitiativeCreate) SetEffort(v float64)`

SetEffort sets Effort field to given value.

### HasEffort

`func (o *InitiativeCreate) HasEffort() bool`

HasEffort returns a boolean if a field has been set.

### GetColor

`func (o *InitiativeCreate) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *InitiativeCreate) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *InitiativeCreate) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *InitiativeCreate) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPresented

`func (o *InitiativeCreate) GetPresented() bool`

GetPresented returns the Presented field if non-nil, zero value otherwise.

### GetPresentedOk

`func (o *InitiativeCreate) GetPresentedOk() (*bool, bool)`

GetPresentedOk returns a tuple with the Presented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresented

`func (o *InitiativeCreate) SetPresented(v bool)`

SetPresented sets Presented field to given value.

### HasPresented

`func (o *InitiativeCreate) HasPresented() bool`

HasPresented returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


