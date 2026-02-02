# CustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **interface{}** | Value can be string, array, or other types depending on field type | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomField

`func NewCustomField() *CustomField`

NewCustomField instantiates a new CustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldWithDefaults

`func NewCustomFieldWithDefaults() *CustomField`

NewCustomFieldWithDefaults instantiates a new CustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CustomField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomField) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CustomField) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *CustomField) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomField) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomField) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomField) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CustomField) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CustomField) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetType

`func (o *CustomField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomField) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


