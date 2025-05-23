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

// SubmitRegistrationFlowBodyWithOidc SubmitRegistrationFlowBodyWithOidc is used to decode the registration form payload when using the oidc method.
type SubmitRegistrationFlowBodyWithOidc struct {
	// The CSRF Token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method to use  This field must be set to `oidc` when using the oidc method.
	Method string `json:"method"`
	// The provider to register with
	Provider string `json:"provider"`
	Traits RegistrationTraits `json:"traits"`
}

// NewSubmitRegistrationFlowBodyWithOidc instantiates a new SubmitRegistrationFlowBodyWithOidc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitRegistrationFlowBodyWithOidc(method string, provider string, traits RegistrationTraits) *SubmitRegistrationFlowBodyWithOidc {
	this := SubmitRegistrationFlowBodyWithOidc{}
	this.Method = method
	this.Provider = provider
	this.Traits = traits
	return &this
}

// NewSubmitRegistrationFlowBodyWithOidcWithDefaults instantiates a new SubmitRegistrationFlowBodyWithOidc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitRegistrationFlowBodyWithOidcWithDefaults() *SubmitRegistrationFlowBodyWithOidc {
	this := SubmitRegistrationFlowBodyWithOidc{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitRegistrationFlowBodyWithOidc) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitRegistrationFlowBodyWithOidc) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitRegistrationFlowBodyWithOidc) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitRegistrationFlowBodyWithOidc) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *SubmitRegistrationFlowBodyWithOidc) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitRegistrationFlowBodyWithOidc) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitRegistrationFlowBodyWithOidc) SetMethod(v string) {
	o.Method = v
}

// GetProvider returns the Provider field value
func (o *SubmitRegistrationFlowBodyWithOidc) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *SubmitRegistrationFlowBodyWithOidc) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *SubmitRegistrationFlowBodyWithOidc) SetProvider(v string) {
	o.Provider = v
}

// GetTraits returns the Traits field value
func (o *SubmitRegistrationFlowBodyWithOidc) GetTraits() RegistrationTraits {
	if o == nil {
		var ret RegistrationTraits
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *SubmitRegistrationFlowBodyWithOidc) GetTraitsOk() (*RegistrationTraits, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Traits, true
}

// SetTraits sets field value
func (o *SubmitRegistrationFlowBodyWithOidc) SetTraits(v RegistrationTraits) {
	o.Traits = v
}

func (o SubmitRegistrationFlowBodyWithOidc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["traits"] = o.Traits
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitRegistrationFlowBodyWithOidc struct {
	value *SubmitRegistrationFlowBodyWithOidc
	isSet bool
}

func (v NullableSubmitRegistrationFlowBodyWithOidc) Get() *SubmitRegistrationFlowBodyWithOidc {
	return v.value
}

func (v *NullableSubmitRegistrationFlowBodyWithOidc) Set(val *SubmitRegistrationFlowBodyWithOidc) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitRegistrationFlowBodyWithOidc) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitRegistrationFlowBodyWithOidc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitRegistrationFlowBodyWithOidc(val *SubmitRegistrationFlowBodyWithOidc) *NullableSubmitRegistrationFlowBodyWithOidc {
	return &NullableSubmitRegistrationFlowBodyWithOidc{value: val, isSet: true}
}

func (v NullableSubmitRegistrationFlowBodyWithOidc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitRegistrationFlowBodyWithOidc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


