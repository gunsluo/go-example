/*
 * Configurator
 *
 * Welcome to the Configurator HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * API version: latest
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AddMembersResponse struct for AddMembersResponse
type AddMembersResponse struct {
	Data AddMembersResponseData `json:"data"`
}

// NewAddMembersResponse instantiates a new AddMembersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddMembersResponse(data AddMembersResponseData) *AddMembersResponse {
	this := AddMembersResponse{}
	this.Data = data
	return &this
}

// NewAddMembersResponseWithDefaults instantiates a new AddMembersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddMembersResponseWithDefaults() *AddMembersResponse {
	this := AddMembersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AddMembersResponse) GetData() AddMembersResponseData {
	if o == nil {
		var ret AddMembersResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddMembersResponse) GetDataOk() (*AddMembersResponseData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddMembersResponse) SetData(v AddMembersResponseData) {
	o.Data = v
}

func (o AddMembersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAddMembersResponse struct {
	value *AddMembersResponse
	isSet bool
}

func (v NullableAddMembersResponse) Get() *AddMembersResponse {
	return v.value
}

func (v *NullableAddMembersResponse) Set(val *AddMembersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddMembersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddMembersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddMembersResponse(val *AddMembersResponse) *NullableAddMembersResponse {
	return &NullableAddMembersResponse{value: val, isSet: true}
}

func (v NullableAddMembersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddMembersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

