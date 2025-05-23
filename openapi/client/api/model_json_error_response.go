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

// JsonErrorResponse Error responses are sent when an error (e.g. unauthorized, bad request, ...) occurred.
type JsonErrorResponse struct {
	// Code responses is error code when an error occurred.  Success                       Code = \"Success\" BadRequest                    Code = \"BadRequest\" InvalidOrganizationID         Code = \"InvalidOrganizationID\" InvalidOrganization           Code = \"InvalidOrganization\" InvalidOrgUser                Code = \"InvalidOrgUser\" EmptyOrgUserList              Code = \"EmptyOrgUserList\" FailedToDeleteOrg             Code = \"FailedToDeleteOrg\" FailedToAddSubsidiary         Code = \"FailedToAddSubsidiary\" InvalidUserInfo               Code = \"InvalidUserInfo\" InvalidUserProfile            Code = \"InvalidUserProfile\" GrpcConnectionFailed          Code = \"GrpcConnectionFailed\" InvalidCrtOrKey               Code = \"InvalidCrtOrKey\" Forbidden                     Code = \"Forbidden\" FailedToPostFile              Code = \"FailedToPostFile\" FileCantParsed                Code = \"FileCantParsed\" SheetNotFound                 Code = \"SheetNotFound\" DataError                     Code = \"DataError\" OrgUserDataEmpty              Code = \"OrgUserDataEmpty\" TemplateError                 Code = \"TemplateError\" MemberLimit100                Code = \"MemberLimit100\" EmptyEmailAddress             Code = \"EmptyEmailAddress\" InvalidEmailFormat            Code = \"InvalidEmailFormat\" InvalidPassword               Code = \"InvalidPassword\" InvalidCompany                Code = \"InvalidCompany\" InvalidFeature                Code = \"InvalidFeature\" InvalidFeatureDefaultRole     Code = \"InvalidFeatureDefaultRole\" InvalidApp                    Code = \"InvalidApp\" InvalidSubOrder               Code = \"InvalidSubOrder\" InvalidAppRole                Code = \"InvalidAppRole\" InvalidPlan                   Code = \"InvalidPlan\" InvalidSubscription           Code = \"InvalidSubscription\" InvalidOrganisationConfigs    Code = \"InvalidOrganisationConfigs\" InvalidPpPRelation            Code = \"InvalidPpPRelation\" InvalidPrePolicy              Code = \"InvalidPrePolicy\" InvalidPhoneNumber            Code = \"InvalidPhoneNumber\" InvalidContactPerson          Code = \"InvalidContactPerson\" InvalidCountry                Code = \"InvalidCountry\" DuplicatedEmailAddress        Code = \"DuplicatedEmailAddress\" DuplicatedMember              Code = \"DuplicatedMember\" EmptyMemberName               Code = \"EmptyMemberName\" TheEmailAddressAlreadyExists  Code = \"TheEmailAddressAlreadyExists\" NotMatchedWithEmail           Code = \"NotMatchedWithEmail\" MaxMemberLimitError           Code = \"MaxMemberLimitError\" DataValidationError           Code = \"DataValidationError\" FailedToCreateUser            Code = \"FailedToCreateUser\" FailedToCreateSubOrder        Code = \"FailedToCreateSubscriptionOrder\" FailedToCreateOrganizationApp Code = \"FailedToCreateOrganizationApp\" FailedToAddUserToFreePlan     Code = \"FailedToAddUserToFreePlan\" FailedToAddUserToAdmin        Code = \"FailedToAddUserToAdmin\" FailedToAddSubscription       Code = \"FailedToAddSubscription\" FailedToCreateUserTags        Code = \"FailedToCreateUserTags\" FailedToExecuteCallback       Code = \"FailedToExecuteCallback\" FailedToSyncOrg               Code = \"FailedToSyncOrg\" FailedToGetCompany            Code = \"FailedToGetCompany\" FailedToSyncOrgApp            Code = \"FailedToSyncOrgApp\" NotTheCreator                 Code = \"NotTheCreator\" FailedToConvertPolicy         Code = \"FailedToConvertPolicy\" FailedToSyncPolicy            Code = \"FailedToSyncPolicy\" FailedToUpsertOrg             Code = \"FailedToUpsertOrg\" FailedToSyncAc                Code = \"FailedToSyncAc\" FailedToAddedDefaultPolicy    Code = \"FailedToAddedDefaultPolicy\" FailedToUpsertAcGroups        Code = \"FailedToUpsertAcGroups\" FailedToSyncSubscriptions     Code = \"FailedToSyncSubscriptions\" EmailAlreadyExists            Code = \"EmailAlreadyExists\" MobileAlreadyExists           Code = \"MobileAlreadyExists\" InternalError                 Code = \"InternalError\" OrgAlreadyExists              Code = \"OrgAlreadyExists\" NotBelongThisOrg              Code = \"NotBelongThisOrg\" FailedToTransferCreator       Code = \"FailedToTransferCreator\" FailedToRemoveOrgMember       Code = \"FailedToRemoveOrgMember\" FailedToGetAcGroup            Code = \"FailedToGetAcGroup\"
	Code *string `json:"code,omitempty"`
	// Detail contains further information on the nature of the error.
	Detail *string `json:"detail,omitempty"`
	// Message is the error message.
	Msg *string `json:"msg,omitempty"`
	// TraceId is the identifier for a trace. It is globally unique.
	TraceId *string `json:"traceId,omitempty"`
	// Type A URI reference that identifies the error type.
	Type *string `json:"type,omitempty"`
}

// NewJsonErrorResponse instantiates a new JsonErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonErrorResponse() *JsonErrorResponse {
	this := JsonErrorResponse{}
	return &this
}

// NewJsonErrorResponseWithDefaults instantiates a new JsonErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonErrorResponseWithDefaults() *JsonErrorResponse {
	this := JsonErrorResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *JsonErrorResponse) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorResponse) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *JsonErrorResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *JsonErrorResponse) SetCode(v string) {
	o.Code = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *JsonErrorResponse) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorResponse) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *JsonErrorResponse) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *JsonErrorResponse) SetDetail(v string) {
	o.Detail = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *JsonErrorResponse) GetMsg() string {
	if o == nil || o.Msg == nil {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorResponse) GetMsgOk() (*string, bool) {
	if o == nil || o.Msg == nil {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *JsonErrorResponse) HasMsg() bool {
	if o != nil && o.Msg != nil {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *JsonErrorResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *JsonErrorResponse) GetTraceId() string {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorResponse) GetTraceIdOk() (*string, bool) {
	if o == nil || o.TraceId == nil {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *JsonErrorResponse) HasTraceId() bool {
	if o != nil && o.TraceId != nil {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *JsonErrorResponse) SetTraceId(v string) {
	o.TraceId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *JsonErrorResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *JsonErrorResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *JsonErrorResponse) SetType(v string) {
	o.Type = &v
}

func (o JsonErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Msg != nil {
		toSerialize["msg"] = o.Msg
	}
	if o.TraceId != nil {
		toSerialize["traceId"] = o.TraceId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableJsonErrorResponse struct {
	value *JsonErrorResponse
	isSet bool
}

func (v NullableJsonErrorResponse) Get() *JsonErrorResponse {
	return v.value
}

func (v *NullableJsonErrorResponse) Set(val *JsonErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonErrorResponse(val *JsonErrorResponse) *NullableJsonErrorResponse {
	return &NullableJsonErrorResponse{value: val, isSet: true}
}

func (v NullableJsonErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


