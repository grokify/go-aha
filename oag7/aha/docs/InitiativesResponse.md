# InitiativesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Initiatives** | Pointer to [**[]InitiativeMeta**](InitiativeMeta.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewInitiativesResponse

`func NewInitiativesResponse() *InitiativesResponse`

NewInitiativesResponse instantiates a new InitiativesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiativesResponseWithDefaults

`func NewInitiativesResponseWithDefaults() *InitiativesResponse`

NewInitiativesResponseWithDefaults instantiates a new InitiativesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiatives

`func (o *InitiativesResponse) GetInitiatives() []InitiativeMeta`

GetInitiatives returns the Initiatives field if non-nil, zero value otherwise.

### GetInitiativesOk

`func (o *InitiativesResponse) GetInitiativesOk() (*[]InitiativeMeta, bool)`

GetInitiativesOk returns a tuple with the Initiatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatives

`func (o *InitiativesResponse) SetInitiatives(v []InitiativeMeta)`

SetInitiatives sets Initiatives field to given value.

### HasInitiatives

`func (o *InitiativesResponse) HasInitiatives() bool`

HasInitiatives returns a boolean if a field has been set.

### GetPagination

`func (o *InitiativesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *InitiativesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *InitiativesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *InitiativesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


