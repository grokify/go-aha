# CommentMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewCommentMeta

`func NewCommentMeta() *CommentMeta`

NewCommentMeta instantiates a new CommentMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentMetaWithDefaults

`func NewCommentMetaWithDefaults() *CommentMeta`

NewCommentMetaWithDefaults instantiates a new CommentMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommentMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommentMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommentMeta) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommentMeta) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBody

`func (o *CommentMeta) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CommentMeta) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CommentMeta) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CommentMeta) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommentMeta) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommentMeta) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommentMeta) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommentMeta) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *CommentMeta) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CommentMeta) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CommentMeta) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CommentMeta) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetResource

`func (o *CommentMeta) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CommentMeta) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CommentMeta) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *CommentMeta) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetUser

`func (o *CommentMeta) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CommentMeta) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CommentMeta) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *CommentMeta) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


