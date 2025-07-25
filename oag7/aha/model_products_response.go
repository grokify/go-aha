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

// checks if the ProductsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductsResponse{}

// ProductsResponse struct for ProductsResponse
type ProductsResponse struct {
	Products []ProductMeta `json:"products,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductsResponse ProductsResponse

// NewProductsResponse instantiates a new ProductsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductsResponse() *ProductsResponse {
	this := ProductsResponse{}
	return &this
}

// NewProductsResponseWithDefaults instantiates a new ProductsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductsResponseWithDefaults() *ProductsResponse {
	this := ProductsResponse{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *ProductsResponse) GetProducts() []ProductMeta {
	if o == nil || IsNil(o.Products) {
		var ret []ProductMeta
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductsResponse) GetProductsOk() ([]ProductMeta, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *ProductsResponse) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []ProductMeta and assigns it to the Products field.
func (o *ProductsResponse) SetProducts(v []ProductMeta) {
	o.Products = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ProductsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ProductsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ProductsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ProductsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductsResponse) UnmarshalJSON(data []byte) (err error) {
	varProductsResponse := _ProductsResponse{}

	err = json.Unmarshal(data, &varProductsResponse)

	if err != nil {
		return err
	}

	*o = ProductsResponse(varProductsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "products")
		delete(additionalProperties, "pagination")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductsResponse struct {
	value *ProductsResponse
	isSet bool
}

func (v NullableProductsResponse) Get() *ProductsResponse {
	return v.value
}

func (v *NullableProductsResponse) Set(val *ProductsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductsResponse(val *ProductsResponse) *NullableProductsResponse {
	return &NullableProductsResponse{value: val, isSet: true}
}

func (v NullableProductsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


