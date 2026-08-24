# IdeaUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdeaUsers** | Pointer to [**[]IdeaUser**](IdeaUser.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewIdeaUsersResponse

`func NewIdeaUsersResponse() *IdeaUsersResponse`

NewIdeaUsersResponse instantiates a new IdeaUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdeaUsersResponseWithDefaults

`func NewIdeaUsersResponseWithDefaults() *IdeaUsersResponse`

NewIdeaUsersResponseWithDefaults instantiates a new IdeaUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdeaUsers

`func (o *IdeaUsersResponse) GetIdeaUsers() []IdeaUser`

GetIdeaUsers returns the IdeaUsers field if non-nil, zero value otherwise.

### GetIdeaUsersOk

`func (o *IdeaUsersResponse) GetIdeaUsersOk() (*[]IdeaUser, bool)`

GetIdeaUsersOk returns a tuple with the IdeaUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeaUsers

`func (o *IdeaUsersResponse) SetIdeaUsers(v []IdeaUser)`

SetIdeaUsers sets IdeaUsers field to given value.

### HasIdeaUsers

`func (o *IdeaUsersResponse) HasIdeaUsers() bool`

HasIdeaUsers returns a boolean if a field has been set.

### GetPagination

`func (o *IdeaUsersResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *IdeaUsersResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *IdeaUsersResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *IdeaUsersResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


