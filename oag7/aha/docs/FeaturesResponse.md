# FeaturesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**[]FeatureMeta**](FeatureMeta.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewFeaturesResponse

`func NewFeaturesResponse() *FeaturesResponse`

NewFeaturesResponse instantiates a new FeaturesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturesResponseWithDefaults

`func NewFeaturesResponseWithDefaults() *FeaturesResponse`

NewFeaturesResponseWithDefaults instantiates a new FeaturesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *FeaturesResponse) GetFeatures() []FeatureMeta`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FeaturesResponse) GetFeaturesOk() (*[]FeatureMeta, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FeaturesResponse) SetFeatures(v []FeatureMeta)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FeaturesResponse) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetPagination

`func (o *FeaturesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *FeaturesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *FeaturesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *FeaturesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


