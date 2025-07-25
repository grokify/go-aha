/*
Aha.io API

Articles that matter on social publishing platform

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aha

import (
	"encoding/json"
)

// checks if the FeatureWrap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureWrap{}

// FeatureWrap struct for FeatureWrap
type FeatureWrap struct {
	Feature *Feature `json:"feature,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FeatureWrap FeatureWrap

// NewFeatureWrap instantiates a new FeatureWrap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureWrap() *FeatureWrap {
	this := FeatureWrap{}
	return &this
}

// NewFeatureWrapWithDefaults instantiates a new FeatureWrap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureWrapWithDefaults() *FeatureWrap {
	this := FeatureWrap{}
	return &this
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *FeatureWrap) GetFeature() Feature {
	if o == nil || IsNil(o.Feature) {
		var ret Feature
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureWrap) GetFeatureOk() (*Feature, bool) {
	if o == nil || IsNil(o.Feature) {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *FeatureWrap) HasFeature() bool {
	if o != nil && !IsNil(o.Feature) {
		return true
	}

	return false
}

// SetFeature gets a reference to the given Feature and assigns it to the Feature field.
func (o *FeatureWrap) SetFeature(v Feature) {
	o.Feature = &v
}

func (o FeatureWrap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureWrap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Feature) {
		toSerialize["feature"] = o.Feature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FeatureWrap) UnmarshalJSON(data []byte) (err error) {
	varFeatureWrap := _FeatureWrap{}

	err = json.Unmarshal(data, &varFeatureWrap)

	if err != nil {
		return err
	}

	*o = FeatureWrap(varFeatureWrap)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "feature")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFeatureWrap struct {
	value *FeatureWrap
	isSet bool
}

func (v NullableFeatureWrap) Get() *FeatureWrap {
	return v.value
}

func (v *NullableFeatureWrap) Set(val *FeatureWrap) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureWrap) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureWrap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureWrap(val *FeatureWrap) *NullableFeatureWrap {
	return &NullableFeatureWrap{value: val, isSet: true}
}

func (v NullableFeatureWrap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureWrap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


