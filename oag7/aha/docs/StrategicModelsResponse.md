# StrategicModelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StrategyModels** | Pointer to [**[]StrategicModelMeta**](StrategicModelMeta.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewStrategicModelsResponse

`func NewStrategicModelsResponse() *StrategicModelsResponse`

NewStrategicModelsResponse instantiates a new StrategicModelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategicModelsResponseWithDefaults

`func NewStrategicModelsResponseWithDefaults() *StrategicModelsResponse`

NewStrategicModelsResponseWithDefaults instantiates a new StrategicModelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategyModels

`func (o *StrategicModelsResponse) GetStrategyModels() []StrategicModelMeta`

GetStrategyModels returns the StrategyModels field if non-nil, zero value otherwise.

### GetStrategyModelsOk

`func (o *StrategicModelsResponse) GetStrategyModelsOk() (*[]StrategicModelMeta, bool)`

GetStrategyModelsOk returns a tuple with the StrategyModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyModels

`func (o *StrategicModelsResponse) SetStrategyModels(v []StrategicModelMeta)`

SetStrategyModels sets StrategyModels field to given value.

### HasStrategyModels

`func (o *StrategicModelsResponse) HasStrategyModels() bool`

HasStrategyModels returns a boolean if a field has been set.

### GetPagination

`func (o *StrategicModelsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *StrategicModelsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *StrategicModelsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *StrategicModelsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


