# DescriptionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Plain text description | [optional] 
**HtmlBody** | Pointer to **string** | HTML formatted description | [optional] 

## Methods

### NewDescriptionObject

`func NewDescriptionObject() *DescriptionObject`

NewDescriptionObject instantiates a new DescriptionObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescriptionObjectWithDefaults

`func NewDescriptionObjectWithDefaults() *DescriptionObject`

NewDescriptionObjectWithDefaults instantiates a new DescriptionObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *DescriptionObject) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *DescriptionObject) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *DescriptionObject) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *DescriptionObject) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHtmlBody

`func (o *DescriptionObject) GetHtmlBody() string`

GetHtmlBody returns the HtmlBody field if non-nil, zero value otherwise.

### GetHtmlBodyOk

`func (o *DescriptionObject) GetHtmlBodyOk() (*string, bool)`

GetHtmlBodyOk returns a tuple with the HtmlBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBody

`func (o *DescriptionObject) SetHtmlBody(v string)`

SetHtmlBody sets HtmlBody field to given value.

### HasHtmlBody

`func (o *DescriptionObject) HasHtmlBody() bool`

HasHtmlBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


