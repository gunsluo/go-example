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

// InitBrowserLogoutFlowResponse struct for InitBrowserLogoutFlowResponse
type InitBrowserLogoutFlowResponse struct {
	// LogoutToken can be used to perform logout using AJAX.
	LogoutToken string `json:"logout_token"`
	// LogoutURL can be opened in a browser to sign the user out.  format: uri
	LogoutUrl string `json:"logout_url"`
}

// NewInitBrowserLogoutFlowResponse instantiates a new InitBrowserLogoutFlowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitBrowserLogoutFlowResponse(logoutToken string, logoutUrl string) *InitBrowserLogoutFlowResponse {
	this := InitBrowserLogoutFlowResponse{}
	this.LogoutToken = logoutToken
	this.LogoutUrl = logoutUrl
	return &this
}

// NewInitBrowserLogoutFlowResponseWithDefaults instantiates a new InitBrowserLogoutFlowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitBrowserLogoutFlowResponseWithDefaults() *InitBrowserLogoutFlowResponse {
	this := InitBrowserLogoutFlowResponse{}
	return &this
}

// GetLogoutToken returns the LogoutToken field value
func (o *InitBrowserLogoutFlowResponse) GetLogoutToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoutToken
}

// GetLogoutTokenOk returns a tuple with the LogoutToken field value
// and a boolean to check if the value has been set.
func (o *InitBrowserLogoutFlowResponse) GetLogoutTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogoutToken, true
}

// SetLogoutToken sets field value
func (o *InitBrowserLogoutFlowResponse) SetLogoutToken(v string) {
	o.LogoutToken = v
}

// GetLogoutUrl returns the LogoutUrl field value
func (o *InitBrowserLogoutFlowResponse) GetLogoutUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value
// and a boolean to check if the value has been set.
func (o *InitBrowserLogoutFlowResponse) GetLogoutUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogoutUrl, true
}

// SetLogoutUrl sets field value
func (o *InitBrowserLogoutFlowResponse) SetLogoutUrl(v string) {
	o.LogoutUrl = v
}

func (o InitBrowserLogoutFlowResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["logout_token"] = o.LogoutToken
	}
	if true {
		toSerialize["logout_url"] = o.LogoutUrl
	}
	return json.Marshal(toSerialize)
}

type NullableInitBrowserLogoutFlowResponse struct {
	value *InitBrowserLogoutFlowResponse
	isSet bool
}

func (v NullableInitBrowserLogoutFlowResponse) Get() *InitBrowserLogoutFlowResponse {
	return v.value
}

func (v *NullableInitBrowserLogoutFlowResponse) Set(val *InitBrowserLogoutFlowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInitBrowserLogoutFlowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInitBrowserLogoutFlowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitBrowserLogoutFlowResponse(val *InitBrowserLogoutFlowResponse) *NullableInitBrowserLogoutFlowResponse {
	return &NullableInitBrowserLogoutFlowResponse{value: val, isSet: true}
}

func (v NullableInitBrowserLogoutFlowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitBrowserLogoutFlowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


