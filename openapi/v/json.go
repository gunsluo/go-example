package v

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Success                       string = "Success"
	Ok                            string = "Ok"
	BadRequest                    string = "BadRequest"
	InvalidOrganizationID         string = "InvalidOrganizationID"
	InvalidOrganization           string = "InvalidOrganization"
	InvalidOrgUser                string = "InvalidOrgUser"
	EmptyOrgUserList              string = "EmptyOrgUserList"
	FailedToDeleteOrg             string = "FailedToDeleteOrg"
	FailedToAddSubsidiary         string = "FailedToAddSubsidiary"
	InvalidUserInfo               string = "InvalidUserInfo"
	InvalidUserProfile            string = "InvalidUserProfile"
	GrpcConnectionFailed          string = "GrpcConnectionFailed"
	InvalidCrtOrKey               string = "InvalidCrtOrKey"
	Forbidden                     string = "Forbidden"
	FailedToPostFile              string = "FailedToPostFile"
	FileCantParsed                string = "FileCantParsed"
	SheetNotFound                 string = "SheetNotFound"
	DataError                     string = "DataError"
	OrgUserDataEmpty              string = "OrgUserDataEmpty"
	TemplateError                 string = "TemplateError"
	MemberLimit100                string = "MemberLimit100"
	EmptyEmailAddress             string = "EmptyEmailAddress"
	InvalidEmailFormat            string = "InvalidEmailFormat"
	InvalidPassword               string = "InvalidPassword"
	InvalidCompany                string = "InvalidCompany"
	InvalidFeature                string = "InvalidFeature"
	InvalidFeatureDefaultRole     string = "InvalidFeatureDefaultRole"
	InvalidApp                    string = "InvalidApp"
	InvalidSubOrder               string = "InvalidSubOrder"
	InvalidAppRole                string = "InvalidAppRole"
	InvalidPlan                   string = "InvalidPlan"
	InvalidSubscription           string = "InvalidSubscription"
	InvalidOrganisationConfigs    string = "InvalidOrganisationConfigs"
	InvalidPpPRelation            string = "InvalidPpPRelation"
	InvalidPrePolicy              string = "InvalidPrePolicy"
	InvalidPhoneNumber            string = "InvalidPhoneNumber"
	InvalidContactPerson          string = "InvalidContactPerson"
	InvalidCountry                string = "InvalidCountry"
	DuplicatedEmailAddress        string = "DuplicatedEmailAddress"
	DuplicatedMember              string = "DuplicatedMember"
	EmptyMemberName               string = "EmptyMemberName"
	TheEmailAddressAlreadyExists  string = "TheEmailAddressAlreadyExists"
	NotMatchedWithEmail           string = "NotMatchedWithEmail"
	MaxMemberLimitError           string = "MaxMemberLimitError"
	DataValidationError           string = "DataValidationError"
	FailedToCreateUser            string = "FailedToCreateUser"
	FailedToCreateSubOrder        string = "FailedToCreateSubscriptionOrder"
	FailedToCreateOrganizationApp string = "FailedToCreateOrganizationApp"
	FailedToAddUserToFreePlan     string = "FailedToAddUserToFreePlan"
	FailedToAddUserToAdmin        string = "FailedToAddUserToAdmin"
	FailedToAddSubscription       string = "FailedToAddSubscription"
	FailedToCreateUserTags        string = "FailedToCreateUserTags"
	FailedToExecuteCallback       string = "FailedToExecuteCallback"
	FailedToSyncOrg               string = "FailedToSyncOrg"
	FailedToGetCompany            string = "FailedToGetCompany"
	FailedToSyncOrgApp            string = "FailedToSyncOrgApp"
	NotTheCreator                 string = "NotTheCreator"
	FailedToConvertPolicy         string = "FailedToConvertPolicy"
	FailedToSyncPolicy            string = "FailedToSyncPolicy"
	FailedToUpsertOrg             string = "FailedToUpsertOrg"
	FailedToSyncAc                string = "FailedToSyncAc"
	FailedToAddedDefaultPolicy    string = "FailedToAddedDefaultPolicy"
	FailedToUpsertAcGroups        string = "FailedToUpsertAcGroups"
	FailedToSyncSubscriptions     string = "FailedToSyncSubscriptions"
	EmailAlreadyExists            string = "EmailAlreadyExists"
	MobileAlreadyExists           string = "MobileAlreadyExists"
	InternalError                 string = "InternalError"
	OrgAlreadyExists              string = "OrgAlreadyExists"
	NotBelongThisOrg              string = "NotBelongThisOrg"
	FailedToTransferCreator       string = "FailedToTransferCreator"
	FailedToRemoveOrgMember       string = "FailedToRemoveOrgMember"
	FailedToGetAcGroup            string = "FailedToGetAcGroup"
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
	Code string `json:"code"`

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

func JSONError(c *gin.Context, code, format string, args ...interface{}) {
	c.JSON(http.StatusOK, ErrorResponse{
		Type:    "/v2/swagger.json",
		Code:    code,
		Message: code,
		Detail:  fmt.Sprintf(format, args...),
	})
}

func JSON(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, v)
}
