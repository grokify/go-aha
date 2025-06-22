# Release

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the release. | [optional] 
**ReferenceNum** | Pointer to **string** | Release referenence number. | [optional] 
**Name** | Pointer to **string** | Release name. | [optional] 
**StartDate** | Pointer to **string** | Start date in YYYY-MM-DD format. | [optional] 
**ReleaseDate** | Pointer to **string** | Release date in YYYY-MM-DD format. | [optional] 
**ExternalReleaseDate** | Pointer to **string** | External release date in YYYY-MM-DD format. | [optional] 
**Released** | Pointer to **bool** |  | [optional] 
**ParkingLot** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** | Web URL for release. | [optional] 
**Resource** | Pointer to **string** | API URL for release. | [optional] 

## Methods

### NewRelease

`func NewRelease() *Release`

NewRelease instantiates a new Release object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseWithDefaults

`func NewReleaseWithDefaults() *Release`

NewReleaseWithDefaults instantiates a new Release object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Release) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Release) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Release) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Release) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferenceNum

`func (o *Release) GetReferenceNum() string`

GetReferenceNum returns the ReferenceNum field if non-nil, zero value otherwise.

### GetReferenceNumOk

`func (o *Release) GetReferenceNumOk() (*string, bool)`

GetReferenceNumOk returns a tuple with the ReferenceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNum

`func (o *Release) SetReferenceNum(v string)`

SetReferenceNum sets ReferenceNum field to given value.

### HasReferenceNum

`func (o *Release) HasReferenceNum() bool`

HasReferenceNum returns a boolean if a field has been set.

### GetName

`func (o *Release) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Release) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Release) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Release) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartDate

`func (o *Release) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Release) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Release) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Release) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetReleaseDate

`func (o *Release) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *Release) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *Release) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *Release) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetExternalReleaseDate

`func (o *Release) GetExternalReleaseDate() string`

GetExternalReleaseDate returns the ExternalReleaseDate field if non-nil, zero value otherwise.

### GetExternalReleaseDateOk

`func (o *Release) GetExternalReleaseDateOk() (*string, bool)`

GetExternalReleaseDateOk returns a tuple with the ExternalReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReleaseDate

`func (o *Release) SetExternalReleaseDate(v string)`

SetExternalReleaseDate sets ExternalReleaseDate field to given value.

### HasExternalReleaseDate

`func (o *Release) HasExternalReleaseDate() bool`

HasExternalReleaseDate returns a boolean if a field has been set.

### GetReleased

`func (o *Release) GetReleased() bool`

GetReleased returns the Released field if non-nil, zero value otherwise.

### GetReleasedOk

`func (o *Release) GetReleasedOk() (*bool, bool)`

GetReleasedOk returns a tuple with the Released field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleased

`func (o *Release) SetReleased(v bool)`

SetReleased sets Released field to given value.

### HasReleased

`func (o *Release) HasReleased() bool`

HasReleased returns a boolean if a field has been set.

### GetParkingLot

`func (o *Release) GetParkingLot() bool`

GetParkingLot returns the ParkingLot field if non-nil, zero value otherwise.

### GetParkingLotOk

`func (o *Release) GetParkingLotOk() (*bool, bool)`

GetParkingLotOk returns a tuple with the ParkingLot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParkingLot

`func (o *Release) SetParkingLot(v bool)`

SetParkingLot sets ParkingLot field to given value.

### HasParkingLot

`func (o *Release) HasParkingLot() bool`

HasParkingLot returns a boolean if a field has been set.

### GetUrl

`func (o *Release) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Release) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Release) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Release) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetResource

`func (o *Release) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Release) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Release) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Release) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


