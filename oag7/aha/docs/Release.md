# Release

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ReferenceNum** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**ReleaseDate** | Pointer to **NullableString** |  | [optional] 
**ExternalReleaseDate** | Pointer to **NullableString** |  | [optional] 
**Released** | Pointer to **bool** |  | [optional] 
**ParkingLot** | Pointer to **bool** |  | [optional] 
**Theme** | Pointer to [**DescriptionObject**](DescriptionObject.md) |  | [optional] 
**ProgressSource** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **NullableFloat64** |  | [optional] 
**WorkflowStatus** | Pointer to [**WorkflowStatus**](WorkflowStatus.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 

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

### SetStartDateNil

`func (o *Release) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Release) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
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

### SetReleaseDateNil

`func (o *Release) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *Release) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
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

### SetExternalReleaseDateNil

`func (o *Release) SetExternalReleaseDateNil(b bool)`

 SetExternalReleaseDateNil sets the value for ExternalReleaseDate to be an explicit nil

### UnsetExternalReleaseDate
`func (o *Release) UnsetExternalReleaseDate()`

UnsetExternalReleaseDate ensures that no value is present for ExternalReleaseDate, not even an explicit nil
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

### GetTheme

`func (o *Release) GetTheme() DescriptionObject`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *Release) GetThemeOk() (*DescriptionObject, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *Release) SetTheme(v DescriptionObject)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *Release) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetProgressSource

`func (o *Release) GetProgressSource() string`

GetProgressSource returns the ProgressSource field if non-nil, zero value otherwise.

### GetProgressSourceOk

`func (o *Release) GetProgressSourceOk() (*string, bool)`

GetProgressSourceOk returns a tuple with the ProgressSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressSource

`func (o *Release) SetProgressSource(v string)`

SetProgressSource sets ProgressSource field to given value.

### HasProgressSource

`func (o *Release) HasProgressSource() bool`

HasProgressSource returns a boolean if a field has been set.

### GetProgress

`func (o *Release) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Release) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Release) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Release) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgressNil

`func (o *Release) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *Release) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetWorkflowStatus

`func (o *Release) GetWorkflowStatus() WorkflowStatus`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *Release) GetWorkflowStatusOk() (*WorkflowStatus, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *Release) SetWorkflowStatus(v WorkflowStatus)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *Release) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

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


