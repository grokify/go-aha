# WorkflowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**Complete** | Pointer to **bool** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowStatus

`func NewWorkflowStatus() *WorkflowStatus`

NewWorkflowStatus instantiates a new WorkflowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStatusWithDefaults

`func NewWorkflowStatusWithDefaults() *WorkflowStatus`

NewWorkflowStatusWithDefaults instantiates a new WorkflowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPosition

`func (o *WorkflowStatus) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WorkflowStatus) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WorkflowStatus) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *WorkflowStatus) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetComplete

`func (o *WorkflowStatus) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *WorkflowStatus) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *WorkflowStatus) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *WorkflowStatus) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetColor

`func (o *WorkflowStatus) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WorkflowStatus) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WorkflowStatus) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WorkflowStatus) HasColor() bool`

HasColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


