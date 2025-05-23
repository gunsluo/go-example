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
	"time"
)

// IdentityCredentialIdentifier IdentityCredentialIdentifier represents some specific identifiers
type IdentityCredentialIdentifier struct {
	// CreatedAt is a helper struct field for.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ID is the identity credential's unique identifier.
	Id *string `json:"id,omitempty"`
	// Identifier is the identifier, e.g. email, mobile or others.
	Identifier *string `json:"identifier,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	// UpdatedAt is a helper struct field for.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewIdentityCredentialIdentifier instantiates a new IdentityCredentialIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityCredentialIdentifier() *IdentityCredentialIdentifier {
	this := IdentityCredentialIdentifier{}
	return &this
}

// NewIdentityCredentialIdentifierWithDefaults instantiates a new IdentityCredentialIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityCredentialIdentifierWithDefaults() *IdentityCredentialIdentifier {
	this := IdentityCredentialIdentifier{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IdentityCredentialIdentifier) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentialIdentifier) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IdentityCredentialIdentifier) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IdentityCredentialIdentifier) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityCredentialIdentifier) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentialIdentifier) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityCredentialIdentifier) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityCredentialIdentifier) SetId(v string) {
	o.Id = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *IdentityCredentialIdentifier) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentialIdentifier) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *IdentityCredentialIdentifier) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *IdentityCredentialIdentifier) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *IdentityCredentialIdentifier) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentialIdentifier) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *IdentityCredentialIdentifier) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *IdentityCredentialIdentifier) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IdentityCredentialIdentifier) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentialIdentifier) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IdentityCredentialIdentifier) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *IdentityCredentialIdentifier) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o IdentityCredentialIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityCredentialIdentifier struct {
	value *IdentityCredentialIdentifier
	isSet bool
}

func (v NullableIdentityCredentialIdentifier) Get() *IdentityCredentialIdentifier {
	return v.value
}

func (v *NullableIdentityCredentialIdentifier) Set(val *IdentityCredentialIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityCredentialIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityCredentialIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityCredentialIdentifier(val *IdentityCredentialIdentifier) *NullableIdentityCredentialIdentifier {
	return &NullableIdentityCredentialIdentifier{value: val, isSet: true}
}

func (v NullableIdentityCredentialIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityCredentialIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


