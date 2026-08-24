# RequirementCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Requirement name | 
**Description** | Pointer to **string** | Requirement description (HTML allowed) | [optional] 
**WorkflowStatus** | Pointer to **string** | Workflow status ID or name | [optional] 
**AssignedToUser** | Pointer to **string** | User ID or email | [optional] 
**OriginalEstimate** | Pointer to **float32** | Original estimate in work units | [optional] 

## Methods

### NewRequirementCreate

`func NewRequirementCreate(name string, ) *RequirementCreate`

NewRequirementCreate instantiates a new RequirementCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementCreateWithDefaults

`func NewRequirementCreateWithDefaults() *RequirementCreate`

NewRequirementCreateWithDefaults instantiates a new RequirementCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RequirementCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequirementCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequirementCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RequirementCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequirementCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequirementCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequirementCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *RequirementCreate) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *RequirementCreate) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *RequirementCreate) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *RequirementCreate) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetAssignedToUser

`func (o *RequirementCreate) GetAssignedToUser() string`

GetAssignedToUser returns the AssignedToUser field if non-nil, zero value otherwise.

### GetAssignedToUserOk

`func (o *RequirementCreate) GetAssignedToUserOk() (*string, bool)`

GetAssignedToUserOk returns a tuple with the AssignedToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToUser

`func (o *RequirementCreate) SetAssignedToUser(v string)`

SetAssignedToUser sets AssignedToUser field to given value.

### HasAssignedToUser

`func (o *RequirementCreate) HasAssignedToUser() bool`

HasAssignedToUser returns a boolean if a field has been set.

### GetOriginalEstimate

`func (o *RequirementCreate) GetOriginalEstimate() float32`

GetOriginalEstimate returns the OriginalEstimate field if non-nil, zero value otherwise.

### GetOriginalEstimateOk

`func (o *RequirementCreate) GetOriginalEstimateOk() (*float32, bool)`

GetOriginalEstimateOk returns a tuple with the OriginalEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEstimate

`func (o *RequirementCreate) SetOriginalEstimate(v float32)`

SetOriginalEstimate sets OriginalEstimate field to given value.

### HasOriginalEstimate

`func (o *RequirementCreate) HasOriginalEstimate() bool`

HasOriginalEstimate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


