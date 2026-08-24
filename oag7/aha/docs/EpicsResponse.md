# EpicsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Epics** | Pointer to [**[]EpicMeta**](EpicMeta.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewEpicsResponse

`func NewEpicsResponse() *EpicsResponse`

NewEpicsResponse instantiates a new EpicsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpicsResponseWithDefaults

`func NewEpicsResponseWithDefaults() *EpicsResponse`

NewEpicsResponseWithDefaults instantiates a new EpicsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpics

`func (o *EpicsResponse) GetEpics() []EpicMeta`

GetEpics returns the Epics field if non-nil, zero value otherwise.

### GetEpicsOk

`func (o *EpicsResponse) GetEpicsOk() (*[]EpicMeta, bool)`

GetEpicsOk returns a tuple with the Epics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpics

`func (o *EpicsResponse) SetEpics(v []EpicMeta)`

SetEpics sets Epics field to given value.

### HasEpics

`func (o *EpicsResponse) HasEpics() bool`

HasEpics returns a boolean if a field has been set.

### GetPagination

`func (o *EpicsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *EpicsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *EpicsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *EpicsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


