# StrategicModelCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Strategic model name | 
**Kind** | **string** | Type of canvas (e.g., Opportunity, Lean Canvas) | 
**Description** | Pointer to **string** | Strategic model description | [optional] 

## Methods

### NewStrategicModelCreate

`func NewStrategicModelCreate(name string, kind string, ) *StrategicModelCreate`

NewStrategicModelCreate instantiates a new StrategicModelCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategicModelCreateWithDefaults

`func NewStrategicModelCreateWithDefaults() *StrategicModelCreate`

NewStrategicModelCreateWithDefaults instantiates a new StrategicModelCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StrategicModelCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StrategicModelCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StrategicModelCreate) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *StrategicModelCreate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StrategicModelCreate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StrategicModelCreate) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDescription

`func (o *StrategicModelCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StrategicModelCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StrategicModelCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StrategicModelCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


