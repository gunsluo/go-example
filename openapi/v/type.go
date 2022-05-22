package v

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Generic Error Code
//
// Code responses is error code when an error occurred.
//
// Success                       Code = "Success"
// BadRequest                    Code = "BadRequest"
// InvalidOrganizationID         Code = "InvalidOrganizationID"
// InvalidOrganization           Code = "InvalidOrganization"
// InvalidOrgUser                Code = "InvalidOrgUser"
// EmptyOrgUserList              Code = "EmptyOrgUserList"
// FailedToDeleteOrg             Code = "FailedToDeleteOrg"
// FailedToAddSubsidiary         Code = "FailedToAddSubsidiary"
// InvalidUserInfo               Code = "InvalidUserInfo"
// InvalidUserProfile            Code = "InvalidUserProfile"
// GrpcConnectionFailed          Code = "GrpcConnectionFailed"
// InvalidCrtOrKey               Code = "InvalidCrtOrKey"
// Forbidden                     Code = "Forbidden"
// FailedToPostFile              Code = "FailedToPostFile"
// FileCantParsed                Code = "FileCantParsed"
// SheetNotFound                 Code = "SheetNotFound"
// DataError                     Code = "DataError"
// OrgUserDataEmpty              Code = "OrgUserDataEmpty"
// TemplateError                 Code = "TemplateError"
// MemberLimit100                Code = "MemberLimit100"
// EmptyEmailAddress             Code = "EmptyEmailAddress"
// InvalidEmailFormat            Code = "InvalidEmailFormat"
// InvalidPassword               Code = "InvalidPassword"
// InvalidCompany                Code = "InvalidCompany"
// InvalidFeature                Code = "InvalidFeature"
// InvalidFeatureDefaultRole     Code = "InvalidFeatureDefaultRole"
// InvalidApp                    Code = "InvalidApp"
// InvalidSubOrder               Code = "InvalidSubOrder"
// InvalidAppRole                Code = "InvalidAppRole"
// InvalidPlan                   Code = "InvalidPlan"
// InvalidSubscription           Code = "InvalidSubscription"
// InvalidOrganisationConfigs    Code = "InvalidOrganisationConfigs"
// InvalidPpPRelation            Code = "InvalidPpPRelation"
// InvalidPrePolicy              Code = "InvalidPrePolicy"
// InvalidPhoneNumber            Code = "InvalidPhoneNumber"
// InvalidContactPerson          Code = "InvalidContactPerson"
// InvalidCountry                Code = "InvalidCountry"
// DuplicatedEmailAddress        Code = "DuplicatedEmailAddress"
// DuplicatedMember              Code = "DuplicatedMember"
// EmptyMemberName               Code = "EmptyMemberName"
// TheEmailAddressAlreadyExists  Code = "TheEmailAddressAlreadyExists"
// NotMatchedWithEmail           Code = "NotMatchedWithEmail"
// MaxMemberLimitError           Code = "MaxMemberLimitError"
// DataValidationError           Code = "DataValidationError"
// FailedToCreateUser            Code = "FailedToCreateUser"
// FailedToCreateSubOrder        Code = "FailedToCreateSubscriptionOrder"
// FailedToCreateOrganizationApp Code = "FailedToCreateOrganizationApp"
// FailedToAddUserToFreePlan     Code = "FailedToAddUserToFreePlan"
// FailedToAddUserToAdmin        Code = "FailedToAddUserToAdmin"
// FailedToAddSubscription       Code = "FailedToAddSubscription"
// FailedToCreateUserTags        Code = "FailedToCreateUserTags"
// FailedToExecuteCallback       Code = "FailedToExecuteCallback"
// FailedToSyncOrg               Code = "FailedToSyncOrg"
// FailedToGetCompany            Code = "FailedToGetCompany"
// FailedToSyncOrgApp            Code = "FailedToSyncOrgApp"
// NotTheCreator                 Code = "NotTheCreator"
// FailedToConvertPolicy         Code = "FailedToConvertPolicy"
// FailedToSyncPolicy            Code = "FailedToSyncPolicy"
// FailedToUpsertOrg             Code = "FailedToUpsertOrg"
// FailedToSyncAc                Code = "FailedToSyncAc"
// FailedToAddedDefaultPolicy    Code = "FailedToAddedDefaultPolicy"
// FailedToUpsertAcGroups        Code = "FailedToUpsertAcGroups"
// FailedToSyncSubscriptions     Code = "FailedToSyncSubscriptions"
// EmailAlreadyExists            Code = "EmailAlreadyExists"
// MobileAlreadyExists           Code = "MobileAlreadyExists"
// InternalError                 Code = "InternalError"
// OrgAlreadyExists              Code = "OrgAlreadyExists"
// NotBelongThisOrg              Code = "NotBelongThisOrg"
// FailedToTransferCreator       Code = "FailedToTransferCreator"
// FailedToRemoveOrgMember       Code = "FailedToRemoveOrgMember"
// FailedToGetAcGroup            Code = "FailedToGetAcGroup"
//
// swagger:model code
type Code string

const (
	Success                       Code = "Success"
	BadRequest                    Code = "BadRequest"
	InvalidOrganizationID         Code = "InvalidOrganizationID"
	InvalidOrganization           Code = "InvalidOrganization"
	InvalidOrgUser                Code = "InvalidOrgUser"
	EmptyOrgUserList              Code = "EmptyOrgUserList"
	FailedToDeleteOrg             Code = "FailedToDeleteOrg"
	FailedToAddSubsidiary         Code = "FailedToAddSubsidiary"
	InvalidUserInfo               Code = "InvalidUserInfo"
	InvalidUserProfile            Code = "InvalidUserProfile"
	GrpcConnectionFailed          Code = "GrpcConnectionFailed"
	InvalidCrtOrKey               Code = "InvalidCrtOrKey"
	Forbidden                     Code = "Forbidden"
	FailedToPostFile              Code = "FailedToPostFile"
	FileCantParsed                Code = "FileCantParsed"
	SheetNotFound                 Code = "SheetNotFound"
	DataError                     Code = "DataError"
	OrgUserDataEmpty              Code = "OrgUserDataEmpty"
	TemplateError                 Code = "TemplateError"
	MemberLimit100                Code = "MemberLimit100"
	EmptyEmailAddress             Code = "EmptyEmailAddress"
	InvalidEmailFormat            Code = "InvalidEmailFormat"
	InvalidPassword               Code = "InvalidPassword"
	InvalidCompany                Code = "InvalidCompany"
	InvalidFeature                Code = "InvalidFeature"
	InvalidFeatureDefaultRole     Code = "InvalidFeatureDefaultRole"
	InvalidApp                    Code = "InvalidApp"
	InvalidSubOrder               Code = "InvalidSubOrder"
	InvalidAppRole                Code = "InvalidAppRole"
	InvalidPlan                   Code = "InvalidPlan"
	InvalidSubscription           Code = "InvalidSubscription"
	InvalidOrganisationConfigs    Code = "InvalidOrganisationConfigs"
	InvalidPpPRelation            Code = "InvalidPpPRelation"
	InvalidPrePolicy              Code = "InvalidPrePolicy"
	InvalidPhoneNumber            Code = "InvalidPhoneNumber"
	InvalidContactPerson          Code = "InvalidContactPerson"
	InvalidCountry                Code = "InvalidCountry"
	DuplicatedEmailAddress        Code = "DuplicatedEmailAddress"
	DuplicatedMember              Code = "DuplicatedMember"
	EmptyMemberName               Code = "EmptyMemberName"
	TheEmailAddressAlreadyExists  Code = "TheEmailAddressAlreadyExists"
	NotMatchedWithEmail           Code = "NotMatchedWithEmail"
	MaxMemberLimitError           Code = "MaxMemberLimitError"
	DataValidationError           Code = "DataValidationError"
	FailedToCreateUser            Code = "FailedToCreateUser"
	FailedToCreateSubOrder        Code = "FailedToCreateSubscriptionOrder"
	FailedToCreateOrganizationApp Code = "FailedToCreateOrganizationApp"
	FailedToAddUserToFreePlan     Code = "FailedToAddUserToFreePlan"
	FailedToAddUserToAdmin        Code = "FailedToAddUserToAdmin"
	FailedToAddSubscription       Code = "FailedToAddSubscription"
	FailedToCreateUserTags        Code = "FailedToCreateUserTags"
	FailedToExecuteCallback       Code = "FailedToExecuteCallback"
	FailedToSyncOrg               Code = "FailedToSyncOrg"
	FailedToGetCompany            Code = "FailedToGetCompany"
	FailedToSyncOrgApp            Code = "FailedToSyncOrgApp"
	NotTheCreator                 Code = "NotTheCreator"
	FailedToConvertPolicy         Code = "FailedToConvertPolicy"
	FailedToSyncPolicy            Code = "FailedToSyncPolicy"
	FailedToUpsertOrg             Code = "FailedToUpsertOrg"
	FailedToSyncAc                Code = "FailedToSyncAc"
	FailedToAddedDefaultPolicy    Code = "FailedToAddedDefaultPolicy"
	FailedToUpsertAcGroups        Code = "FailedToUpsertAcGroups"
	FailedToSyncSubscriptions     Code = "FailedToSyncSubscriptions"
	EmailAlreadyExists            Code = "EmailAlreadyExists"
	MobileAlreadyExists           Code = "MobileAlreadyExists"
	InternalError                 Code = "InternalError"
	OrgAlreadyExists              Code = "OrgAlreadyExists"
	NotBelongThisOrg              Code = "NotBelongThisOrg"
	FailedToTransferCreator       Code = "FailedToTransferCreator"
	FailedToRemoveOrgMember       Code = "FailedToRemoveOrgMember"
	FailedToGetAcGroup            Code = "FailedToGetAcGroup"
)

// Generic Error Response
//
// Error responses are sent when an error (e.g. unauthorized, bad request, ...) occurred.
//
// swagger:model jsonErrorResponse
type ErrorResponse struct {
	// Type A URI reference that identifies the error type.
	//
	// example: https://example.net/validation-error
	Type string `json:"type,omitempty"`

	// Code represents the error code.
	//
	// example: Success or BadRequest
	Code Code `json:"code"`

	// Message is the error message.
	//
	// example: The requested resource could not be found
	Message string `json:"msg,omitempty"`

	// Detail contains further information on the nature of the error.
	//
	// example: Member with ID 12345 does not exist
	Detail string `json:"detail,omitempty"`

	// TraceId is the identifier for a trace. It is globally unique.
	//
	// example: 463ac35c9f6413ad48485a3953bb6124
	TraceId string `json:"traceId,omitempty"`
}

func JSONError(c *gin.Context, code Code, format string, args ...interface{}) {
	c.JSON(http.StatusOK, ErrorResponse{
		Type:    "/v2/swagger.json",
		Code:    code,
		Message: string(code),
		Detail:  fmt.Sprintf(format, args...),
	})
}

func JSON(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, v)
}
