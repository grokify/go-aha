# CustomFieldOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Color** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCustomFieldOption

`func NewCustomFieldOption() *CustomFieldOption`

NewCustomFieldOption instantiates a new CustomFieldOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldOptionWithDefaults

`func NewCustomFieldOptionWithDefaults() *CustomFieldOption`

NewCustomFieldOptionWithDefaults instantiates a new CustomFieldOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFieldOption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldOption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldOption) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomFieldOption) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *CustomFieldOption) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldOption) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldOption) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldOption) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPosition

`func (o *CustomFieldOption) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CustomFieldOption) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CustomFieldOption) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CustomFieldOption) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetColor

`func (o *CustomFieldOption) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CustomFieldOption) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CustomFieldOption) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CustomFieldOption) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CustomFieldOption) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CustomFieldOption) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


