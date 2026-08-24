# CustomFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | Display name of the field | 
**Key** | **string** | Internal reference key | 
**Type** | **string** | Field type (e.g., CustomFieldDefinitions::SelectConstant) | 
**CustomFieldableType** | **string** | Record type the field applies to (Feature, Initiative, etc.) | 
**InternalName** | Pointer to **NullableString** | Internal system name | [optional] 
**Position** | Pointer to **int32** | Display order position | [optional] 
**ApiType** | Pointer to **string** | API type for the field | [optional] 
**AllowsOtherOption** | Pointer to **bool** | Whether \&quot;Other\&quot; option is allowed for select fields | [optional] 

## Methods

### NewCustomFieldDefinition

`func NewCustomFieldDefinition(id string, name string, key string, type_ string, customFieldableType string, ) *CustomFieldDefinition`

NewCustomFieldDefinition instantiates a new CustomFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldDefinitionWithDefaults

`func NewCustomFieldDefinitionWithDefaults() *CustomFieldDefinition`

NewCustomFieldDefinitionWithDefaults instantiates a new CustomFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFieldDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomFieldDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFieldDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFieldDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *CustomFieldDefinition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomFieldDefinition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomFieldDefinition) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *CustomFieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetCustomFieldableType

`func (o *CustomFieldDefinition) GetCustomFieldableType() string`

GetCustomFieldableType returns the CustomFieldableType field if non-nil, zero value otherwise.

### GetCustomFieldableTypeOk

`func (o *CustomFieldDefinition) GetCustomFieldableTypeOk() (*string, bool)`

GetCustomFieldableTypeOk returns a tuple with the CustomFieldableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldableType

`func (o *CustomFieldDefinition) SetCustomFieldableType(v string)`

SetCustomFieldableType sets CustomFieldableType field to given value.


### GetInternalName

`func (o *CustomFieldDefinition) GetInternalName() string`

GetInternalName returns the InternalName field if non-nil, zero value otherwise.

### GetInternalNameOk

`func (o *CustomFieldDefinition) GetInternalNameOk() (*string, bool)`

GetInternalNameOk returns a tuple with the InternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalName

`func (o *CustomFieldDefinition) SetInternalName(v string)`

SetInternalName sets InternalName field to given value.

### HasInternalName

`func (o *CustomFieldDefinition) HasInternalName() bool`

HasInternalName returns a boolean if a field has been set.

### SetInternalNameNil

`func (o *CustomFieldDefinition) SetInternalNameNil(b bool)`

 SetInternalNameNil sets the value for InternalName to be an explicit nil

### UnsetInternalName
`func (o *CustomFieldDefinition) UnsetInternalName()`

UnsetInternalName ensures that no value is present for InternalName, not even an explicit nil
### GetPosition

`func (o *CustomFieldDefinition) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CustomFieldDefinition) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CustomFieldDefinition) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CustomFieldDefinition) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetApiType

`func (o *CustomFieldDefinition) GetApiType() string`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *CustomFieldDefinition) GetApiTypeOk() (*string, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *CustomFieldDefinition) SetApiType(v string)`

SetApiType sets ApiType field to given value.

### HasApiType

`func (o *CustomFieldDefinition) HasApiType() bool`

HasApiType returns a boolean if a field has been set.

### GetAllowsOtherOption

`func (o *CustomFieldDefinition) GetAllowsOtherOption() bool`

GetAllowsOtherOption returns the AllowsOtherOption field if non-nil, zero value otherwise.

### GetAllowsOtherOptionOk

`func (o *CustomFieldDefinition) GetAllowsOtherOptionOk() (*bool, bool)`

GetAllowsOtherOptionOk returns a tuple with the AllowsOtherOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsOtherOption

`func (o *CustomFieldDefinition) SetAllowsOtherOption(v bool)`

SetAllowsOtherOption sets AllowsOtherOption field to given value.

### HasAllowsOtherOption

`func (o *CustomFieldDefinition) HasAllowsOtherOption() bool`

HasAllowsOtherOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


