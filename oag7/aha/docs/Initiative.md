# Initiative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ReferenceNum** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to [**DescriptionObject**](DescriptionObject.md) |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**Effort** | Pointer to **float64** |  | [optional] 
**Presented** | Pointer to **bool** |  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **NullableString** |  | [optional] 
**Progress** | Pointer to **NullableFloat64** |  | [optional] 
**ProgressSource** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | Web UI URL | [optional] 
**Resource** | Pointer to **string** | API URL | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**WorkflowStatus** | Pointer to [**WorkflowStatus**](WorkflowStatus.md) |  | [optional] 
**Epic** | Pointer to [**EpicMeta**](EpicMeta.md) |  | [optional] 
**Features** | Pointer to [**[]FeatureMeta**](FeatureMeta.md) |  | [optional] 
**CustomFields** | Pointer to [**[]CustomField**](CustomField.md) |  | [optional] 

## Methods

### NewInitiative

`func NewInitiative(id string, referenceNum string, name string, createdAt time.Time, ) *Initiative`

NewInitiative instantiates a new Initiative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiativeWithDefaults

`func NewInitiativeWithDefaults() *Initiative`

NewInitiativeWithDefaults instantiates a new Initiative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Initiative) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Initiative) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Initiative) SetId(v string)`

SetId sets Id field to given value.


### GetReferenceNum

`func (o *Initiative) GetReferenceNum() string`

GetReferenceNum returns the ReferenceNum field if non-nil, zero value otherwise.

### GetReferenceNumOk

`func (o *Initiative) GetReferenceNumOk() (*string, bool)`

GetReferenceNumOk returns a tuple with the ReferenceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNum

`func (o *Initiative) SetReferenceNum(v string)`

SetReferenceNum sets ReferenceNum field to given value.


### GetName

`func (o *Initiative) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Initiative) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Initiative) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Initiative) GetDescription() DescriptionObject`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Initiative) GetDescriptionOk() (*DescriptionObject, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Initiative) SetDescription(v DescriptionObject)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Initiative) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetColor

`func (o *Initiative) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Initiative) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Initiative) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Initiative) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPosition

`func (o *Initiative) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Initiative) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Initiative) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Initiative) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetValue

`func (o *Initiative) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Initiative) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Initiative) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *Initiative) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEffort

`func (o *Initiative) GetEffort() float64`

GetEffort returns the Effort field if non-nil, zero value otherwise.

### GetEffortOk

`func (o *Initiative) GetEffortOk() (*float64, bool)`

GetEffortOk returns a tuple with the Effort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffort

`func (o *Initiative) SetEffort(v float64)`

SetEffort sets Effort field to given value.

### HasEffort

`func (o *Initiative) HasEffort() bool`

HasEffort returns a boolean if a field has been set.

### GetPresented

`func (o *Initiative) GetPresented() bool`

GetPresented returns the Presented field if non-nil, zero value otherwise.

### GetPresentedOk

`func (o *Initiative) GetPresentedOk() (*bool, bool)`

GetPresentedOk returns a tuple with the Presented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresented

`func (o *Initiative) SetPresented(v bool)`

SetPresented sets Presented field to given value.

### HasPresented

`func (o *Initiative) HasPresented() bool`

HasPresented returns a boolean if a field has been set.

### GetStartDate

`func (o *Initiative) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Initiative) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Initiative) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Initiative) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *Initiative) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Initiative) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *Initiative) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Initiative) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Initiative) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Initiative) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *Initiative) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *Initiative) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetProgress

`func (o *Initiative) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Initiative) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Initiative) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Initiative) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgressNil

`func (o *Initiative) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *Initiative) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetProgressSource

`func (o *Initiative) GetProgressSource() string`

GetProgressSource returns the ProgressSource field if non-nil, zero value otherwise.

### GetProgressSourceOk

`func (o *Initiative) GetProgressSourceOk() (*string, bool)`

GetProgressSourceOk returns a tuple with the ProgressSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressSource

`func (o *Initiative) SetProgressSource(v string)`

SetProgressSource sets ProgressSource field to given value.

### HasProgressSource

`func (o *Initiative) HasProgressSource() bool`

HasProgressSource returns a boolean if a field has been set.

### GetUrl

`func (o *Initiative) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Initiative) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Initiative) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Initiative) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetResource

`func (o *Initiative) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Initiative) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Initiative) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Initiative) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Initiative) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Initiative) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Initiative) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Initiative) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Initiative) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Initiative) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Initiative) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *Initiative) GetWorkflowStatus() WorkflowStatus`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *Initiative) GetWorkflowStatusOk() (*WorkflowStatus, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *Initiative) SetWorkflowStatus(v WorkflowStatus)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *Initiative) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetEpic

`func (o *Initiative) GetEpic() EpicMeta`

GetEpic returns the Epic field if non-nil, zero value otherwise.

### GetEpicOk

`func (o *Initiative) GetEpicOk() (*EpicMeta, bool)`

GetEpicOk returns a tuple with the Epic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpic

`func (o *Initiative) SetEpic(v EpicMeta)`

SetEpic sets Epic field to given value.

### HasEpic

`func (o *Initiative) HasEpic() bool`

HasEpic returns a boolean if a field has been set.

### GetFeatures

`func (o *Initiative) GetFeatures() []FeatureMeta`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Initiative) GetFeaturesOk() (*[]FeatureMeta, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Initiative) SetFeatures(v []FeatureMeta)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Initiative) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetCustomFields

`func (o *Initiative) GetCustomFields() []CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Initiative) GetCustomFieldsOk() (*[]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Initiative) SetCustomFields(v []CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Initiative) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


