# Idea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ReferenceNum** | **string** |  | 
**Feature** | Pointer to [**IdeaFeature**](IdeaFeature.md) |  | [optional] 
**WorkflowStatus** | Pointer to [**FeatureWorkflowStatus**](FeatureWorkflowStatus.md) |  | [optional] 
**Categories** | [**[]Category**](Category.md) |  | 
**Votes** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**StatusChangedAt** | **time.Time** |  | 

## Methods

### NewIdea

`func NewIdea(id string, name string, referenceNum string, categories []Category, votes int32, createdAt time.Time, updatedAt time.Time, statusChangedAt time.Time, ) *Idea`

NewIdea instantiates a new Idea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdeaWithDefaults

`func NewIdeaWithDefaults() *Idea`

NewIdeaWithDefaults instantiates a new Idea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Idea) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Idea) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Idea) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Idea) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Idea) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Idea) SetName(v string)`

SetName sets Name field to given value.


### GetReferenceNum

`func (o *Idea) GetReferenceNum() string`

GetReferenceNum returns the ReferenceNum field if non-nil, zero value otherwise.

### GetReferenceNumOk

`func (o *Idea) GetReferenceNumOk() (*string, bool)`

GetReferenceNumOk returns a tuple with the ReferenceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNum

`func (o *Idea) SetReferenceNum(v string)`

SetReferenceNum sets ReferenceNum field to given value.


### GetFeature

`func (o *Idea) GetFeature() IdeaFeature`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *Idea) GetFeatureOk() (*IdeaFeature, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *Idea) SetFeature(v IdeaFeature)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *Idea) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *Idea) GetWorkflowStatus() FeatureWorkflowStatus`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *Idea) GetWorkflowStatusOk() (*FeatureWorkflowStatus, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *Idea) SetWorkflowStatus(v FeatureWorkflowStatus)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *Idea) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetCategories

`func (o *Idea) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Idea) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Idea) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetVotes

`func (o *Idea) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *Idea) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *Idea) SetVotes(v int32)`

SetVotes sets Votes field to given value.


### GetCreatedAt

`func (o *Idea) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Idea) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Idea) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Idea) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Idea) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Idea) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatusChangedAt

`func (o *Idea) GetStatusChangedAt() time.Time`

GetStatusChangedAt returns the StatusChangedAt field if non-nil, zero value otherwise.

### GetStatusChangedAtOk

`func (o *Idea) GetStatusChangedAtOk() (*time.Time, bool)`

GetStatusChangedAtOk returns a tuple with the StatusChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChangedAt

`func (o *Idea) SetStatusChangedAt(v time.Time)`

SetStatusChangedAt sets StatusChangedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


