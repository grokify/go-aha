# RequirementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requirements** | Pointer to [**[]RequirementMeta**](RequirementMeta.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewRequirementsResponse

`func NewRequirementsResponse() *RequirementsResponse`

NewRequirementsResponse instantiates a new RequirementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementsResponseWithDefaults

`func NewRequirementsResponseWithDefaults() *RequirementsResponse`

NewRequirementsResponseWithDefaults instantiates a new RequirementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirements

`func (o *RequirementsResponse) GetRequirements() []RequirementMeta`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *RequirementsResponse) GetRequirementsOk() (*[]RequirementMeta, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *RequirementsResponse) SetRequirements(v []RequirementMeta)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *RequirementsResponse) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetPagination

`func (o *RequirementsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RequirementsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RequirementsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RequirementsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


