# ReleaseCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**StartDate** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**ExternalReleaseDate** | Pointer to **string** |  | [optional] 
**DevelopmentStartedOn** | Pointer to **string** |  | [optional] 
**ParkingLot** | Pointer to **bool** |  | [optional] 
**Theme** | Pointer to **string** | Theme of the release (may include HTML formatting) | [optional] 

## Methods

### NewReleaseCreate

`func NewReleaseCreate(name string, ) *ReleaseCreate`

NewReleaseCreate instantiates a new ReleaseCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseCreateWithDefaults

`func NewReleaseCreateWithDefaults() *ReleaseCreate`

NewReleaseCreateWithDefaults instantiates a new ReleaseCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReleaseCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReleaseCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReleaseCreate) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *ReleaseCreate) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ReleaseCreate) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ReleaseCreate) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ReleaseCreate) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetReleaseDate

`func (o *ReleaseCreate) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ReleaseCreate) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ReleaseCreate) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ReleaseCreate) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetExternalReleaseDate

`func (o *ReleaseCreate) GetExternalReleaseDate() string`

GetExternalReleaseDate returns the ExternalReleaseDate field if non-nil, zero value otherwise.

### GetExternalReleaseDateOk

`func (o *ReleaseCreate) GetExternalReleaseDateOk() (*string, bool)`

GetExternalReleaseDateOk returns a tuple with the ExternalReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReleaseDate

`func (o *ReleaseCreate) SetExternalReleaseDate(v string)`

SetExternalReleaseDate sets ExternalReleaseDate field to given value.

### HasExternalReleaseDate

`func (o *ReleaseCreate) HasExternalReleaseDate() bool`

HasExternalReleaseDate returns a boolean if a field has been set.

### GetDevelopmentStartedOn

`func (o *ReleaseCreate) GetDevelopmentStartedOn() string`

GetDevelopmentStartedOn returns the DevelopmentStartedOn field if non-nil, zero value otherwise.

### GetDevelopmentStartedOnOk

`func (o *ReleaseCreate) GetDevelopmentStartedOnOk() (*string, bool)`

GetDevelopmentStartedOnOk returns a tuple with the DevelopmentStartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopmentStartedOn

`func (o *ReleaseCreate) SetDevelopmentStartedOn(v string)`

SetDevelopmentStartedOn sets DevelopmentStartedOn field to given value.

### HasDevelopmentStartedOn

`func (o *ReleaseCreate) HasDevelopmentStartedOn() bool`

HasDevelopmentStartedOn returns a boolean if a field has been set.

### GetParkingLot

`func (o *ReleaseCreate) GetParkingLot() bool`

GetParkingLot returns the ParkingLot field if non-nil, zero value otherwise.

### GetParkingLotOk

`func (o *ReleaseCreate) GetParkingLotOk() (*bool, bool)`

GetParkingLotOk returns a tuple with the ParkingLot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParkingLot

`func (o *ReleaseCreate) SetParkingLot(v bool)`

SetParkingLot sets ParkingLot field to given value.

### HasParkingLot

`func (o *ReleaseCreate) HasParkingLot() bool`

HasParkingLot returns a boolean if a field has been set.

### GetTheme

`func (o *ReleaseCreate) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *ReleaseCreate) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *ReleaseCreate) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *ReleaseCreate) HasTheme() bool`

HasTheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


