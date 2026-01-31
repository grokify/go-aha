# IdeasResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ideas** | Pointer to [**[]Idea**](Idea.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewIdeasResponse

`func NewIdeasResponse() *IdeasResponse`

NewIdeasResponse instantiates a new IdeasResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdeasResponseWithDefaults

`func NewIdeasResponseWithDefaults() *IdeasResponse`

NewIdeasResponseWithDefaults instantiates a new IdeasResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdeas

`func (o *IdeasResponse) GetIdeas() []Idea`

GetIdeas returns the Ideas field if non-nil, zero value otherwise.

### GetIdeasOk

`func (o *IdeasResponse) GetIdeasOk() (*[]Idea, bool)`

GetIdeasOk returns a tuple with the Ideas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeas

`func (o *IdeasResponse) SetIdeas(v []Idea)`

SetIdeas sets Ideas field to given value.

### HasIdeas

`func (o *IdeasResponse) HasIdeas() bool`

HasIdeas returns a boolean if a field has been set.

### GetPagination

`func (o *IdeasResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *IdeasResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *IdeasResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *IdeasResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


