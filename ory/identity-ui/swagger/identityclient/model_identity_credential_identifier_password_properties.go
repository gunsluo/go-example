/*
 * Identity
 *
 * Welcome to the Identity HTTP API documentation! You will find documentation for all HTTP APIs here.
 *
 * API version: latest
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package identityclient

import (
	"encoding/json"
)

// IdentityCredentialIdentifierPasswordProperties struct for IdentityCredentialIdentifierPasswordProperties
type IdentityCredentialIdentifierPasswordProperties struct {
	// Primary  Is identifier's primary identity
	Primary bool `json:"primary"`
	Type *string `json:"type,omitempty"`
	// Verified  Is identifier been verified
	Verified bool `json:"verified"`
}

// NewIdentityCredentialIdentifierPasswordProperties instantiates a new IdentityCredentialIdentifierPasswordProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityCredentialIdentifierPasswordProperties(primary bool, verified bool) *IdentityCredentialIdentifierPasswordProperties {
	this := IdentityCredentialIdentifierPasswordProperties{}
	this.Primary = primary
	this.Verified = verified
	return &this
}

// NewIdentityCredentialIdentifierPasswordPropertiesWithDefaults instantiates a new IdentityCredentialIdentifierPasswordProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityCredentialIdentifierPasswordPropertiesWithDefaults() *IdentityCredentialIdentifierPasswordProperties {
	this := IdentityCredentialIdentifierPasswordProperties{}
	return &this
}

// GetPrimary returns the Primary field value
func (o *IdentityCredentialIdentifierPasswordProperties) GetPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *IdentityCredentialIdentifierPasswordProperties) GetPrimaryOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value
func (o *IdentityCredentialIdentifierPasswordProperties) SetPrimary(v bool) {
	o.Primary = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityCredentialIdentifierPasswordProperties) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentialIdentifierPasswordProperties) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityCredentialIdentifierPasswordProperties) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdentityCredentialIdentifierPasswordProperties) SetType(v string) {
	o.Type = &v
}

// GetVerified returns the Verified field value
func (o *IdentityCredentialIdentifierPasswordProperties) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *IdentityCredentialIdentifierPasswordProperties) GetVerifiedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *IdentityCredentialIdentifierPasswordProperties) SetVerified(v bool) {
	o.Verified = v
}

func (o IdentityCredentialIdentifierPasswordProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["primary"] = o.Primary
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["verified"] = o.Verified
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityCredentialIdentifierPasswordProperties struct {
	value *IdentityCredentialIdentifierPasswordProperties
	isSet bool
}

func (v NullableIdentityCredentialIdentifierPasswordProperties) Get() *IdentityCredentialIdentifierPasswordProperties {
	return v.value
}

func (v *NullableIdentityCredentialIdentifierPasswordProperties) Set(val *IdentityCredentialIdentifierPasswordProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityCredentialIdentifierPasswordProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityCredentialIdentifierPasswordProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityCredentialIdentifierPasswordProperties(val *IdentityCredentialIdentifierPasswordProperties) *NullableIdentityCredentialIdentifierPasswordProperties {
	return &NullableIdentityCredentialIdentifierPasswordProperties{value: val, isSet: true}
}

func (v NullableIdentityCredentialIdentifierPasswordProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityCredentialIdentifierPasswordProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

