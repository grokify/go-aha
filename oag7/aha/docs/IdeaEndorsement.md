# IdeaEndorsement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**IdeaId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**Link** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**EndorsedByPortalUser** | Pointer to [**EndorsementPortalUser**](EndorsementPortalUser.md) |  | [optional] 
**EndorsedByIdeaUser** | Pointer to [**EndorsementIdeaUser**](EndorsementIdeaUser.md) |  | [optional] 

## Methods

### NewIdeaEndorsement

`func NewIdeaEndorsement() *IdeaEndorsement`

NewIdeaEndorsement instantiates a new IdeaEndorsement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdeaEndorsementWithDefaults

`func NewIdeaEndorsementWithDefaults() *IdeaEndorsement`

NewIdeaEndorsementWithDefaults instantiates a new IdeaEndorsement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdeaEndorsement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdeaEndorsement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdeaEndorsement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdeaEndorsement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdeaId

`func (o *IdeaEndorsement) GetIdeaId() string`

GetIdeaId returns the IdeaId field if non-nil, zero value otherwise.

### GetIdeaIdOk

`func (o *IdeaEndorsement) GetIdeaIdOk() (*string, bool)`

GetIdeaIdOk returns a tuple with the IdeaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeaId

`func (o *IdeaEndorsement) SetIdeaId(v string)`

SetIdeaId sets IdeaId field to given value.

### HasIdeaId

`func (o *IdeaEndorsement) HasIdeaId() bool`

HasIdeaId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IdeaEndorsement) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdeaEndorsement) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdeaEndorsement) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdeaEndorsement) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IdeaEndorsement) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IdeaEndorsement) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IdeaEndorsement) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IdeaEndorsement) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *IdeaEndorsement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IdeaEndorsement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IdeaEndorsement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IdeaEndorsement) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *IdeaEndorsement) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *IdeaEndorsement) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetLink

`func (o *IdeaEndorsement) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *IdeaEndorsement) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *IdeaEndorsement) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *IdeaEndorsement) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *IdeaEndorsement) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *IdeaEndorsement) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetWeight

`func (o *IdeaEndorsement) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *IdeaEndorsement) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *IdeaEndorsement) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *IdeaEndorsement) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetEndorsedByPortalUser

`func (o *IdeaEndorsement) GetEndorsedByPortalUser() EndorsementPortalUser`

GetEndorsedByPortalUser returns the EndorsedByPortalUser field if non-nil, zero value otherwise.

### GetEndorsedByPortalUserOk

`func (o *IdeaEndorsement) GetEndorsedByPortalUserOk() (*EndorsementPortalUser, bool)`

GetEndorsedByPortalUserOk returns a tuple with the EndorsedByPortalUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndorsedByPortalUser

`func (o *IdeaEndorsement) SetEndorsedByPortalUser(v EndorsementPortalUser)`

SetEndorsedByPortalUser sets EndorsedByPortalUser field to given value.

### HasEndorsedByPortalUser

`func (o *IdeaEndorsement) HasEndorsedByPortalUser() bool`

HasEndorsedByPortalUser returns a boolean if a field has been set.

### GetEndorsedByIdeaUser

`func (o *IdeaEndorsement) GetEndorsedByIdeaUser() EndorsementIdeaUser`

GetEndorsedByIdeaUser returns the EndorsedByIdeaUser field if non-nil, zero value otherwise.

### GetEndorsedByIdeaUserOk

`func (o *IdeaEndorsement) GetEndorsedByIdeaUserOk() (*EndorsementIdeaUser, bool)`

GetEndorsedByIdeaUserOk returns a tuple with the EndorsedByIdeaUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndorsedByIdeaUser

`func (o *IdeaEndorsement) SetEndorsedByIdeaUser(v EndorsementIdeaUser)`

SetEndorsedByIdeaUser sets EndorsedByIdeaUser field to given value.

### HasEndorsedByIdeaUser

`func (o *IdeaEndorsement) HasEndorsedByIdeaUser() bool`

HasEndorsedByIdeaUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


