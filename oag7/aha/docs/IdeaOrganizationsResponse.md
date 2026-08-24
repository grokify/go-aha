# IdeaOrganizationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdeaOrganizations** | Pointer to [**[]IdeaOrganizationRef**](IdeaOrganizationRef.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewIdeaOrganizationsResponse

`func NewIdeaOrganizationsResponse() *IdeaOrganizationsResponse`

NewIdeaOrganizationsResponse instantiates a new IdeaOrganizationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdeaOrganizationsResponseWithDefaults

`func NewIdeaOrganizationsResponseWithDefaults() *IdeaOrganizationsResponse`

NewIdeaOrganizationsResponseWithDefaults instantiates a new IdeaOrganizationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdeaOrganizations

`func (o *IdeaOrganizationsResponse) GetIdeaOrganizations() []IdeaOrganizationRef`

GetIdeaOrganizations returns the IdeaOrganizations field if non-nil, zero value otherwise.

### GetIdeaOrganizationsOk

`func (o *IdeaOrganizationsResponse) GetIdeaOrganizationsOk() (*[]IdeaOrganizationRef, bool)`

GetIdeaOrganizationsOk returns a tuple with the IdeaOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeaOrganizations

`func (o *IdeaOrganizationsResponse) SetIdeaOrganizations(v []IdeaOrganizationRef)`

SetIdeaOrganizations sets IdeaOrganizations field to given value.

### HasIdeaOrganizations

`func (o *IdeaOrganizationsResponse) HasIdeaOrganizations() bool`

HasIdeaOrganizations returns a boolean if a field has been set.

### GetPagination

`func (o *IdeaOrganizationsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *IdeaOrganizationsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *IdeaOrganizationsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *IdeaOrganizationsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


