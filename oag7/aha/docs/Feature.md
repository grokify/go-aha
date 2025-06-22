# Feature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ReferenceNum** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **string** | Start date in YYYY-MM-DD format. | [optional] 
**DueDate** | Pointer to **string** | Due date in YYYY-MM-DD format. | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**Release** | Pointer to [**Release**](Release.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFeature

`func NewFeature() *Feature`

NewFeature instantiates a new Feature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureWithDefaults

`func NewFeatureWithDefaults() *Feature`

NewFeatureWithDefaults instantiates a new Feature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Feature) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Feature) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Feature) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Feature) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferenceNum

`func (o *Feature) GetReferenceNum() string`

GetReferenceNum returns the ReferenceNum field if non-nil, zero value otherwise.

### GetReferenceNumOk

`func (o *Feature) GetReferenceNumOk() (*string, bool)`

GetReferenceNumOk returns a tuple with the ReferenceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNum

`func (o *Feature) SetReferenceNum(v string)`

SetReferenceNum sets ReferenceNum field to given value.

### HasReferenceNum

`func (o *Feature) HasReferenceNum() bool`

HasReferenceNum returns a boolean if a field has been set.

### GetName

`func (o *Feature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Feature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Feature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Feature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Feature) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Feature) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Feature) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Feature) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStartDate

`func (o *Feature) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Feature) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Feature) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Feature) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetDueDate

`func (o *Feature) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Feature) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Feature) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Feature) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetUrl

`func (o *Feature) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Feature) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Feature) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Feature) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetResource

`func (o *Feature) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Feature) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Feature) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Feature) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRelease

`func (o *Feature) GetRelease() Release`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *Feature) GetReleaseOk() (*Release, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *Feature) SetRelease(v Release)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *Feature) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetTags

`func (o *Feature) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Feature) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Feature) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Feature) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


