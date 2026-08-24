# IdeaCategoriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdeaCategories** | Pointer to [**[]IdeaCategory**](IdeaCategory.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewIdeaCategoriesResponse

`func NewIdeaCategoriesResponse() *IdeaCategoriesResponse`

NewIdeaCategoriesResponse instantiates a new IdeaCategoriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdeaCategoriesResponseWithDefaults

`func NewIdeaCategoriesResponseWithDefaults() *IdeaCategoriesResponse`

NewIdeaCategoriesResponseWithDefaults instantiates a new IdeaCategoriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdeaCategories

`func (o *IdeaCategoriesResponse) GetIdeaCategories() []IdeaCategory`

GetIdeaCategories returns the IdeaCategories field if non-nil, zero value otherwise.

### GetIdeaCategoriesOk

`func (o *IdeaCategoriesResponse) GetIdeaCategoriesOk() (*[]IdeaCategory, bool)`

GetIdeaCategoriesOk returns a tuple with the IdeaCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeaCategories

`func (o *IdeaCategoriesResponse) SetIdeaCategories(v []IdeaCategory)`

SetIdeaCategories sets IdeaCategories field to given value.

### HasIdeaCategories

`func (o *IdeaCategoriesResponse) HasIdeaCategories() bool`

HasIdeaCategories returns a boolean if a field has been set.

### GetPagination

`func (o *IdeaCategoriesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *IdeaCategoriesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *IdeaCategoriesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *IdeaCategoriesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


