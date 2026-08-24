# GoalsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Goals** | Pointer to [**[]GoalMeta**](GoalMeta.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGoalsResponse

`func NewGoalsResponse() *GoalsResponse`

NewGoalsResponse instantiates a new GoalsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoalsResponseWithDefaults

`func NewGoalsResponseWithDefaults() *GoalsResponse`

NewGoalsResponseWithDefaults instantiates a new GoalsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGoals

`func (o *GoalsResponse) GetGoals() []GoalMeta`

GetGoals returns the Goals field if non-nil, zero value otherwise.

### GetGoalsOk

`func (o *GoalsResponse) GetGoalsOk() (*[]GoalMeta, bool)`

GetGoalsOk returns a tuple with the Goals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoals

`func (o *GoalsResponse) SetGoals(v []GoalMeta)`

SetGoals sets Goals field to given value.

### HasGoals

`func (o *GoalsResponse) HasGoals() bool`

HasGoals returns a boolean if a field has been set.

### GetPagination

`func (o *GoalsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GoalsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GoalsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GoalsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


