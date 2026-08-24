# WorkflowsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflows** | Pointer to [**[]Workflow**](Workflow.md) |  | [optional] 

## Methods

### NewWorkflowsResponse

`func NewWorkflowsResponse() *WorkflowsResponse`

NewWorkflowsResponse instantiates a new WorkflowsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsResponseWithDefaults

`func NewWorkflowsResponseWithDefaults() *WorkflowsResponse`

NewWorkflowsResponseWithDefaults instantiates a new WorkflowsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflows

`func (o *WorkflowsResponse) GetWorkflows() []Workflow`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *WorkflowsResponse) GetWorkflowsOk() (*[]Workflow, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *WorkflowsResponse) SetWorkflows(v []Workflow)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *WorkflowsResponse) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


