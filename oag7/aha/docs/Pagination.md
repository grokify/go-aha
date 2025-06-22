# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRecords** | Pointer to **int64** |  | [optional] 
**TotalPages** | Pointer to **int64** |  | [optional] 
**CurrentPage** | Pointer to **int64** |  | [optional] 

## Methods

### NewPagination

`func NewPagination() *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRecords

`func (o *Pagination) GetTotalRecords() int64`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *Pagination) GetTotalRecordsOk() (*int64, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *Pagination) SetTotalRecords(v int64)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *Pagination) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.

### GetTotalPages

`func (o *Pagination) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *Pagination) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *Pagination) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *Pagination) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetCurrentPage

`func (o *Pagination) GetCurrentPage() int64`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *Pagination) GetCurrentPageOk() (*int64, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *Pagination) SetCurrentPage(v int64)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *Pagination) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


