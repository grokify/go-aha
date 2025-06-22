# FeatureUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the feature | [optional] 
**Description** | Pointer to **string** | Description of the feature and it can include HTML formatting. | [optional] 
**CreatedBy** | Pointer to **string** | Email address of user that created the feature. | [optional] 
**AssignedToUser** | Pointer to **string** | Email address of user that is assigned the feature. | [optional] 
**Tags** | Pointer to **string** | Tags can be automatically assigned to the new feature. If more than one tag is used then tags should be separated by commas | [optional] 
**OriginalEstimateText** | Pointer to **string** | Set the original estimated effort in a text format, you can use d, h, min (or &#39;p&#39; for points) to indicate the units to use. | [optional] 
**RemainingEstimateText** | Pointer to **string** |  Set the remaining estimated effort in a text format, you can use d, h, min (or &#39;p&#39; for points) to indicate the units to use. | [optional] 
**StartDate** | Pointer to **string** | Date that work will start on the feature in format YYYY-MM-DD. | [optional] 
**DueDate** | Pointer to **string** | Date that work is due to be completed on the feature in format YYYY-MM-DD. | [optional] 
**ReleasePhase** | Pointer to **string** | Name or id of release phase which the feature belongs to. | [optional] 
**Initiative** | Pointer to **string** | Name or id of initiative which the feature belongs to. | [optional] 
**MasterFeature** | Pointer to **string** | Name or id of master feature which the feature belongs to. | [optional] 

## Methods

### NewFeatureUpdate

`func NewFeatureUpdate() *FeatureUpdate`

NewFeatureUpdate instantiates a new FeatureUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureUpdateWithDefaults

`func NewFeatureUpdateWithDefaults() *FeatureUpdate`

NewFeatureUpdateWithDefaults instantiates a new FeatureUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeatureUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FeatureUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FeatureUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeatureUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FeatureUpdate) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FeatureUpdate) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FeatureUpdate) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FeatureUpdate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetAssignedToUser

`func (o *FeatureUpdate) GetAssignedToUser() string`

GetAssignedToUser returns the AssignedToUser field if non-nil, zero value otherwise.

### GetAssignedToUserOk

`func (o *FeatureUpdate) GetAssignedToUserOk() (*string, bool)`

GetAssignedToUserOk returns a tuple with the AssignedToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToUser

`func (o *FeatureUpdate) SetAssignedToUser(v string)`

SetAssignedToUser sets AssignedToUser field to given value.

### HasAssignedToUser

`func (o *FeatureUpdate) HasAssignedToUser() bool`

HasAssignedToUser returns a boolean if a field has been set.

### GetTags

`func (o *FeatureUpdate) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeatureUpdate) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeatureUpdate) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeatureUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOriginalEstimateText

`func (o *FeatureUpdate) GetOriginalEstimateText() string`

GetOriginalEstimateText returns the OriginalEstimateText field if non-nil, zero value otherwise.

### GetOriginalEstimateTextOk

`func (o *FeatureUpdate) GetOriginalEstimateTextOk() (*string, bool)`

GetOriginalEstimateTextOk returns a tuple with the OriginalEstimateText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEstimateText

`func (o *FeatureUpdate) SetOriginalEstimateText(v string)`

SetOriginalEstimateText sets OriginalEstimateText field to given value.

### HasOriginalEstimateText

`func (o *FeatureUpdate) HasOriginalEstimateText() bool`

HasOriginalEstimateText returns a boolean if a field has been set.

### GetRemainingEstimateText

`func (o *FeatureUpdate) GetRemainingEstimateText() string`

GetRemainingEstimateText returns the RemainingEstimateText field if non-nil, zero value otherwise.

### GetRemainingEstimateTextOk

`func (o *FeatureUpdate) GetRemainingEstimateTextOk() (*string, bool)`

GetRemainingEstimateTextOk returns a tuple with the RemainingEstimateText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingEstimateText

`func (o *FeatureUpdate) SetRemainingEstimateText(v string)`

SetRemainingEstimateText sets RemainingEstimateText field to given value.

### HasRemainingEstimateText

`func (o *FeatureUpdate) HasRemainingEstimateText() bool`

HasRemainingEstimateText returns a boolean if a field has been set.

### GetStartDate

`func (o *FeatureUpdate) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FeatureUpdate) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FeatureUpdate) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FeatureUpdate) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetDueDate

`func (o *FeatureUpdate) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *FeatureUpdate) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *FeatureUpdate) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *FeatureUpdate) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetReleasePhase

`func (o *FeatureUpdate) GetReleasePhase() string`

GetReleasePhase returns the ReleasePhase field if non-nil, zero value otherwise.

### GetReleasePhaseOk

`func (o *FeatureUpdate) GetReleasePhaseOk() (*string, bool)`

GetReleasePhaseOk returns a tuple with the ReleasePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasePhase

`func (o *FeatureUpdate) SetReleasePhase(v string)`

SetReleasePhase sets ReleasePhase field to given value.

### HasReleasePhase

`func (o *FeatureUpdate) HasReleasePhase() bool`

HasReleasePhase returns a boolean if a field has been set.

### GetInitiative

`func (o *FeatureUpdate) GetInitiative() string`

GetInitiative returns the Initiative field if non-nil, zero value otherwise.

### GetInitiativeOk

`func (o *FeatureUpdate) GetInitiativeOk() (*string, bool)`

GetInitiativeOk returns a tuple with the Initiative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiative

`func (o *FeatureUpdate) SetInitiative(v string)`

SetInitiative sets Initiative field to given value.

### HasInitiative

`func (o *FeatureUpdate) HasInitiative() bool`

HasInitiative returns a boolean if a field has been set.

### GetMasterFeature

`func (o *FeatureUpdate) GetMasterFeature() string`

GetMasterFeature returns the MasterFeature field if non-nil, zero value otherwise.

### GetMasterFeatureOk

`func (o *FeatureUpdate) GetMasterFeatureOk() (*string, bool)`

GetMasterFeatureOk returns a tuple with the MasterFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterFeature

`func (o *FeatureUpdate) SetMasterFeature(v string)`

SetMasterFeature sets MasterFeature field to given value.

### HasMasterFeature

`func (o *FeatureUpdate) HasMasterFeature() bool`

HasMasterFeature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


