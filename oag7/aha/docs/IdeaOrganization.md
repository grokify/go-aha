# IdeaOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReferenceNum** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**EndorsementsCount** | Pointer to **int32** |  | [optional] 
**EmailDomains** | Pointer to **NullableString** |  | [optional] 
**Revenue** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewIdeaOrganization

`func NewIdeaOrganization() *IdeaOrganization`

NewIdeaOrganization instantiates a new IdeaOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdeaOrganizationWithDefaults

`func NewIdeaOrganizationWithDefaults() *IdeaOrganization`

NewIdeaOrganizationWithDefaults instantiates a new IdeaOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdeaOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdeaOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdeaOrganization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdeaOrganization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdeaOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdeaOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdeaOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdeaOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferenceNum

`func (o *IdeaOrganization) GetReferenceNum() string`

GetReferenceNum returns the ReferenceNum field if non-nil, zero value otherwise.

### GetReferenceNumOk

`func (o *IdeaOrganization) GetReferenceNumOk() (*string, bool)`

GetReferenceNumOk returns a tuple with the ReferenceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNum

`func (o *IdeaOrganization) SetReferenceNum(v string)`

SetReferenceNum sets ReferenceNum field to given value.

### HasReferenceNum

`func (o *IdeaOrganization) HasReferenceNum() bool`

HasReferenceNum returns a boolean if a field has been set.

### GetUrl

`func (o *IdeaOrganization) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IdeaOrganization) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IdeaOrganization) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IdeaOrganization) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IdeaOrganization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdeaOrganization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdeaOrganization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdeaOrganization) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IdeaOrganization) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IdeaOrganization) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IdeaOrganization) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IdeaOrganization) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEndorsementsCount

`func (o *IdeaOrganization) GetEndorsementsCount() int32`

GetEndorsementsCount returns the EndorsementsCount field if non-nil, zero value otherwise.

### GetEndorsementsCountOk

`func (o *IdeaOrganization) GetEndorsementsCountOk() (*int32, bool)`

GetEndorsementsCountOk returns a tuple with the EndorsementsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndorsementsCount

`func (o *IdeaOrganization) SetEndorsementsCount(v int32)`

SetEndorsementsCount sets EndorsementsCount field to given value.

### HasEndorsementsCount

`func (o *IdeaOrganization) HasEndorsementsCount() bool`

HasEndorsementsCount returns a boolean if a field has been set.

### GetEmailDomains

`func (o *IdeaOrganization) GetEmailDomains() string`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *IdeaOrganization) GetEmailDomainsOk() (*string, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *IdeaOrganization) SetEmailDomains(v string)`

SetEmailDomains sets EmailDomains field to given value.

### HasEmailDomains

`func (o *IdeaOrganization) HasEmailDomains() bool`

HasEmailDomains returns a boolean if a field has been set.

### SetEmailDomainsNil

`func (o *IdeaOrganization) SetEmailDomainsNil(b bool)`

 SetEmailDomainsNil sets the value for EmailDomains to be an explicit nil

### UnsetEmailDomains
`func (o *IdeaOrganization) UnsetEmailDomains()`

UnsetEmailDomains ensures that no value is present for EmailDomains, not even an explicit nil
### GetRevenue

`func (o *IdeaOrganization) GetRevenue() float64`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *IdeaOrganization) GetRevenueOk() (*float64, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *IdeaOrganization) SetRevenue(v float64)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *IdeaOrganization) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### SetRevenueNil

`func (o *IdeaOrganization) SetRevenueNil(b bool)`

 SetRevenueNil sets the value for Revenue to be an explicit nil

### UnsetRevenue
`func (o *IdeaOrganization) UnsetRevenue()`

UnsetRevenue ensures that no value is present for Revenue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


