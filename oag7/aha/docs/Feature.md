# Feature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**CommentsCount** | Pointer to **int64** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomField**](CustomField.md) |  | [optional] 
**DueDate** | Pointer to **string** | Due date in YYYY-MM-DD format. | [optional] 
**Name** | **string** |  | 
**ProductId** | Pointer to **string** |  | [optional] 
**ProgressSource** | Pointer to **string** |  | [optional] 
**ReferenceNum** | **string** |  | 
**Release** | Pointer to [**Release**](Release.md) |  | [optional] 
**Resource** | Pointer to **string** | API URL for the resource. | [optional] 
**StartDate** | Pointer to **string** | Start date in YYYY-MM-DD format. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Url** | Pointer to **string** | Web UI URL for the resource. | [optional] 
**UseRequirementsEstimates** | Pointer to **bool** |  | [optional] 
**WorkUnits** | Pointer to **int64** |  | [optional] 
**WorkflowStatus** | Pointer to [**FeatureWorkflowStatus**](FeatureWorkflowStatus.md) |  | [optional] 
**IntegrationFields** | Pointer to [**[]IntegrationField**](IntegrationField.md) |  | [optional] 

## Methods

### NewFeature

`func NewFeature(id string, createdAt time.Time, name string, referenceNum string, ) *Feature`

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


### GetCommentsCount

`func (o *Feature) GetCommentsCount() int64`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *Feature) GetCommentsCountOk() (*int64, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *Feature) SetCommentsCount(v int64)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *Feature) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetCustomFields

`func (o *Feature) GetCustomFields() []CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Feature) GetCustomFieldsOk() (*[]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Feature) SetCustomFields(v []CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Feature) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

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


### GetProductId

`func (o *Feature) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Feature) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Feature) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Feature) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProgressSource

`func (o *Feature) GetProgressSource() string`

GetProgressSource returns the ProgressSource field if non-nil, zero value otherwise.

### GetProgressSourceOk

`func (o *Feature) GetProgressSourceOk() (*string, bool)`

GetProgressSourceOk returns a tuple with the ProgressSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressSource

`func (o *Feature) SetProgressSource(v string)`

SetProgressSource sets ProgressSource field to given value.

### HasProgressSource

`func (o *Feature) HasProgressSource() bool`

HasProgressSource returns a boolean if a field has been set.

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

### GetUseRequirementsEstimates

`func (o *Feature) GetUseRequirementsEstimates() bool`

GetUseRequirementsEstimates returns the UseRequirementsEstimates field if non-nil, zero value otherwise.

### GetUseRequirementsEstimatesOk

`func (o *Feature) GetUseRequirementsEstimatesOk() (*bool, bool)`

GetUseRequirementsEstimatesOk returns a tuple with the UseRequirementsEstimates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRequirementsEstimates

`func (o *Feature) SetUseRequirementsEstimates(v bool)`

SetUseRequirementsEstimates sets UseRequirementsEstimates field to given value.

### HasUseRequirementsEstimates

`func (o *Feature) HasUseRequirementsEstimates() bool`

HasUseRequirementsEstimates returns a boolean if a field has been set.

### GetWorkUnits

`func (o *Feature) GetWorkUnits() int64`

GetWorkUnits returns the WorkUnits field if non-nil, zero value otherwise.

### GetWorkUnitsOk

`func (o *Feature) GetWorkUnitsOk() (*int64, bool)`

GetWorkUnitsOk returns a tuple with the WorkUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkUnits

`func (o *Feature) SetWorkUnits(v int64)`

SetWorkUnits sets WorkUnits field to given value.

### HasWorkUnits

`func (o *Feature) HasWorkUnits() bool`

HasWorkUnits returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *Feature) GetWorkflowStatus() FeatureWorkflowStatus`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *Feature) GetWorkflowStatusOk() (*FeatureWorkflowStatus, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *Feature) SetWorkflowStatus(v FeatureWorkflowStatus)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *Feature) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetIntegrationFields

`func (o *Feature) GetIntegrationFields() []IntegrationField`

GetIntegrationFields returns the IntegrationFields field if non-nil, zero value otherwise.

### GetIntegrationFieldsOk

`func (o *Feature) GetIntegrationFieldsOk() (*[]IntegrationField, bool)`

GetIntegrationFieldsOk returns a tuple with the IntegrationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationFields

`func (o *Feature) SetIntegrationFields(v []IntegrationField)`

SetIntegrationFields sets IntegrationFields field to given value.

### HasIntegrationFields

`func (o *Feature) HasIntegrationFields() bool`

HasIntegrationFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


