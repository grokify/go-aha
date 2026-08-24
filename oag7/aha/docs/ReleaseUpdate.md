# ReleaseUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**ReleaseDate** | Pointer to **NullableString** |  | [optional] 
**ExternalReleaseDate** | Pointer to **NullableString** |  | [optional] 
**DevelopmentStartedOn** | Pointer to **string** |  | [optional] 
**ParkingLot** | Pointer to **bool** |  | [optional] 
**Theme** | Pointer to **string** | Theme of the release (may include HTML formatting). Also shown as the release description in the Aha! UI. | [optional] 
**ProgressSource** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**WorkflowStatus** | Pointer to **string** | Workflow status ID or name. Transitioning to the status the product&#39;s workflow defines as \&quot;released\&quot; is how a release becomes Released. | [optional] 

## Methods

### NewReleaseUpdate

`func NewReleaseUpdate() *ReleaseUpdate`

NewReleaseUpdate instantiates a new ReleaseUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseUpdateWithDefaults

`func NewReleaseUpdateWithDefaults() *ReleaseUpdate`

NewReleaseUpdateWithDefaults instantiates a new ReleaseUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReleaseUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReleaseUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReleaseUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReleaseUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartDate

`func (o *ReleaseUpdate) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ReleaseUpdate) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ReleaseUpdate) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ReleaseUpdate) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *ReleaseUpdate) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *ReleaseUpdate) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetReleaseDate

`func (o *ReleaseUpdate) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ReleaseUpdate) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ReleaseUpdate) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ReleaseUpdate) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *ReleaseUpdate) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *ReleaseUpdate) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetExternalReleaseDate

`func (o *ReleaseUpdate) GetExternalReleaseDate() string`

GetExternalReleaseDate returns the ExternalReleaseDate field if non-nil, zero value otherwise.

### GetExternalReleaseDateOk

`func (o *ReleaseUpdate) GetExternalReleaseDateOk() (*string, bool)`

GetExternalReleaseDateOk returns a tuple with the ExternalReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReleaseDate

`func (o *ReleaseUpdate) SetExternalReleaseDate(v string)`

SetExternalReleaseDate sets ExternalReleaseDate field to given value.

### HasExternalReleaseDate

`func (o *ReleaseUpdate) HasExternalReleaseDate() bool`

HasExternalReleaseDate returns a boolean if a field has been set.

### SetExternalReleaseDateNil

`func (o *ReleaseUpdate) SetExternalReleaseDateNil(b bool)`

 SetExternalReleaseDateNil sets the value for ExternalReleaseDate to be an explicit nil

### UnsetExternalReleaseDate
`func (o *ReleaseUpdate) UnsetExternalReleaseDate()`

UnsetExternalReleaseDate ensures that no value is present for ExternalReleaseDate, not even an explicit nil
### GetDevelopmentStartedOn

`func (o *ReleaseUpdate) GetDevelopmentStartedOn() string`

GetDevelopmentStartedOn returns the DevelopmentStartedOn field if non-nil, zero value otherwise.

### GetDevelopmentStartedOnOk

`func (o *ReleaseUpdate) GetDevelopmentStartedOnOk() (*string, bool)`

GetDevelopmentStartedOnOk returns a tuple with the DevelopmentStartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopmentStartedOn

`func (o *ReleaseUpdate) SetDevelopmentStartedOn(v string)`

SetDevelopmentStartedOn sets DevelopmentStartedOn field to given value.

### HasDevelopmentStartedOn

`func (o *ReleaseUpdate) HasDevelopmentStartedOn() bool`

HasDevelopmentStartedOn returns a boolean if a field has been set.

### GetParkingLot

`func (o *ReleaseUpdate) GetParkingLot() bool`

GetParkingLot returns the ParkingLot field if non-nil, zero value otherwise.

### GetParkingLotOk

`func (o *ReleaseUpdate) GetParkingLotOk() (*bool, bool)`

GetParkingLotOk returns a tuple with the ParkingLot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParkingLot

`func (o *ReleaseUpdate) SetParkingLot(v bool)`

SetParkingLot sets ParkingLot field to given value.

### HasParkingLot

`func (o *ReleaseUpdate) HasParkingLot() bool`

HasParkingLot returns a boolean if a field has been set.

### GetTheme

`func (o *ReleaseUpdate) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *ReleaseUpdate) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *ReleaseUpdate) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *ReleaseUpdate) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetProgressSource

`func (o *ReleaseUpdate) GetProgressSource() string`

GetProgressSource returns the ProgressSource field if non-nil, zero value otherwise.

### GetProgressSourceOk

`func (o *ReleaseUpdate) GetProgressSourceOk() (*string, bool)`

GetProgressSourceOk returns a tuple with the ProgressSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressSource

`func (o *ReleaseUpdate) SetProgressSource(v string)`

SetProgressSource sets ProgressSource field to given value.

### HasProgressSource

`func (o *ReleaseUpdate) HasProgressSource() bool`

HasProgressSource returns a boolean if a field has been set.

### GetProgress

`func (o *ReleaseUpdate) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ReleaseUpdate) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ReleaseUpdate) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ReleaseUpdate) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *ReleaseUpdate) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *ReleaseUpdate) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *ReleaseUpdate) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *ReleaseUpdate) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


