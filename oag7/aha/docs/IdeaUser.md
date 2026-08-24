# IdeaUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**IdeaOrganizations** | Pointer to [**[]IdeaOrganizationRef**](IdeaOrganizationRef.md) |  | [optional] 

## Methods

### NewIdeaUser

`func NewIdeaUser() *IdeaUser`

NewIdeaUser instantiates a new IdeaUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdeaUserWithDefaults

`func NewIdeaUserWithDefaults() *IdeaUser`

NewIdeaUserWithDefaults instantiates a new IdeaUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdeaUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdeaUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdeaUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdeaUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdeaUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdeaUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdeaUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdeaUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *IdeaUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdeaUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdeaUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IdeaUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IdeaUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdeaUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdeaUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdeaUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIdeaOrganizations

`func (o *IdeaUser) GetIdeaOrganizations() []IdeaOrganizationRef`

GetIdeaOrganizations returns the IdeaOrganizations field if non-nil, zero value otherwise.

### GetIdeaOrganizationsOk

`func (o *IdeaUser) GetIdeaOrganizationsOk() (*[]IdeaOrganizationRef, bool)`

GetIdeaOrganizationsOk returns a tuple with the IdeaOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeaOrganizations

`func (o *IdeaUser) SetIdeaOrganizations(v []IdeaOrganizationRef)`

SetIdeaOrganizations sets IdeaOrganizations field to given value.

### HasIdeaOrganizations

`func (o *IdeaUser) HasIdeaOrganizations() bool`

HasIdeaOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


