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

// SubmitLoginFlowBodyWithPassword struct for SubmitLoginFlowBodyWithPassword
type SubmitLoginFlowBodyWithPassword struct {
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Identifier is the email or username of the user trying to log in.
	Identifier string `json:"identifier"`
	// Method should be set to \"password\" when logging in using the identifier and password strategy.
	Method string `json:"method"`
	// The user's password.
	Password string `json:"password"`
}

// NewSubmitLoginFlowBodyWithPassword instantiates a new SubmitLoginFlowBodyWithPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitLoginFlowBodyWithPassword(identifier string, method string, password string) *SubmitLoginFlowBodyWithPassword {
	this := SubmitLoginFlowBodyWithPassword{}
	this.Identifier = identifier
	this.Method = method
	this.Password = password
	return &this
}

// NewSubmitLoginFlowBodyWithPasswordWithDefaults instantiates a new SubmitLoginFlowBodyWithPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitLoginFlowBodyWithPasswordWithDefaults() *SubmitLoginFlowBodyWithPassword {
	this := SubmitLoginFlowBodyWithPassword{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitLoginFlowBodyWithPassword) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitLoginFlowBodyWithPassword) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitLoginFlowBodyWithPassword) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitLoginFlowBodyWithPassword) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetIdentifier returns the Identifier field value
func (o *SubmitLoginFlowBodyWithPassword) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *SubmitLoginFlowBodyWithPassword) GetIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *SubmitLoginFlowBodyWithPassword) SetIdentifier(v string) {
	o.Identifier = v
}

// GetMethod returns the Method field value
func (o *SubmitLoginFlowBodyWithPassword) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitLoginFlowBodyWithPassword) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitLoginFlowBodyWithPassword) SetMethod(v string) {
	o.Method = v
}

// GetPassword returns the Password field value
func (o *SubmitLoginFlowBodyWithPassword) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SubmitLoginFlowBodyWithPassword) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SubmitLoginFlowBodyWithPassword) SetPassword(v string) {
	o.Password = v
}

func (o SubmitLoginFlowBodyWithPassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitLoginFlowBodyWithPassword struct {
	value *SubmitLoginFlowBodyWithPassword
	isSet bool
}

func (v NullableSubmitLoginFlowBodyWithPassword) Get() *SubmitLoginFlowBodyWithPassword {
	return v.value
}

func (v *NullableSubmitLoginFlowBodyWithPassword) Set(val *SubmitLoginFlowBodyWithPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitLoginFlowBodyWithPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitLoginFlowBodyWithPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitLoginFlowBodyWithPassword(val *SubmitLoginFlowBodyWithPassword) *NullableSubmitLoginFlowBodyWithPassword {
	return &NullableSubmitLoginFlowBodyWithPassword{value: val, isSet: true}
}

func (v NullableSubmitLoginFlowBodyWithPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitLoginFlowBodyWithPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

