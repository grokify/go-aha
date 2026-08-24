# IdeaUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**WorkflowStatus** | Pointer to **string** | Workflow status ID or name | [optional] 
**Categories** | Pointer to **[]string** | Category IDs or names | [optional] 
**Visibility** | Pointer to **string** | Visibility (e.g., public, private) | [optional] 

## Methods

### NewIdeaUpdate

`func NewIdeaUpdate() *IdeaUpdate`

NewIdeaUpdate instantiates a new IdeaUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdeaUpdateWithDefaults

`func NewIdeaUpdateWithDefaults() *IdeaUpdate`

NewIdeaUpdateWithDefaults instantiates a new IdeaUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdeaUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdeaUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdeaUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdeaUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *IdeaUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdeaUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdeaUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdeaUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *IdeaUpdate) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *IdeaUpdate) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *IdeaUpdate) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *IdeaUpdate) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetCategories

`func (o *IdeaUpdate) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *IdeaUpdate) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *IdeaUpdate) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *IdeaUpdate) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetVisibility

`func (o *IdeaUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *IdeaUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *IdeaUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *IdeaUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


