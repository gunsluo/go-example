# JsonErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code responses is error code when an error occurred.  Success                       Code &#x3D; \&quot;Success\&quot; BadRequest                    Code &#x3D; \&quot;BadRequest\&quot; InvalidOrganizationID         Code &#x3D; \&quot;InvalidOrganizationID\&quot; InvalidOrganization           Code &#x3D; \&quot;InvalidOrganization\&quot; InvalidOrgUser                Code &#x3D; \&quot;InvalidOrgUser\&quot; EmptyOrgUserList              Code &#x3D; \&quot;EmptyOrgUserList\&quot; FailedToDeleteOrg             Code &#x3D; \&quot;FailedToDeleteOrg\&quot; FailedToAddSubsidiary         Code &#x3D; \&quot;FailedToAddSubsidiary\&quot; InvalidUserInfo               Code &#x3D; \&quot;InvalidUserInfo\&quot; InvalidUserProfile            Code &#x3D; \&quot;InvalidUserProfile\&quot; GrpcConnectionFailed          Code &#x3D; \&quot;GrpcConnectionFailed\&quot; InvalidCrtOrKey               Code &#x3D; \&quot;InvalidCrtOrKey\&quot; Forbidden                     Code &#x3D; \&quot;Forbidden\&quot; FailedToPostFile              Code &#x3D; \&quot;FailedToPostFile\&quot; FileCantParsed                Code &#x3D; \&quot;FileCantParsed\&quot; SheetNotFound                 Code &#x3D; \&quot;SheetNotFound\&quot; DataError                     Code &#x3D; \&quot;DataError\&quot; OrgUserDataEmpty              Code &#x3D; \&quot;OrgUserDataEmpty\&quot; TemplateError                 Code &#x3D; \&quot;TemplateError\&quot; MemberLimit100                Code &#x3D; \&quot;MemberLimit100\&quot; EmptyEmailAddress             Code &#x3D; \&quot;EmptyEmailAddress\&quot; InvalidEmailFormat            Code &#x3D; \&quot;InvalidEmailFormat\&quot; InvalidPassword               Code &#x3D; \&quot;InvalidPassword\&quot; InvalidCompany                Code &#x3D; \&quot;InvalidCompany\&quot; InvalidFeature                Code &#x3D; \&quot;InvalidFeature\&quot; InvalidFeatureDefaultRole     Code &#x3D; \&quot;InvalidFeatureDefaultRole\&quot; InvalidApp                    Code &#x3D; \&quot;InvalidApp\&quot; InvalidSubOrder               Code &#x3D; \&quot;InvalidSubOrder\&quot; InvalidAppRole                Code &#x3D; \&quot;InvalidAppRole\&quot; InvalidPlan                   Code &#x3D; \&quot;InvalidPlan\&quot; InvalidSubscription           Code &#x3D; \&quot;InvalidSubscription\&quot; InvalidOrganisationConfigs    Code &#x3D; \&quot;InvalidOrganisationConfigs\&quot; InvalidPpPRelation            Code &#x3D; \&quot;InvalidPpPRelation\&quot; InvalidPrePolicy              Code &#x3D; \&quot;InvalidPrePolicy\&quot; InvalidPhoneNumber            Code &#x3D; \&quot;InvalidPhoneNumber\&quot; InvalidContactPerson          Code &#x3D; \&quot;InvalidContactPerson\&quot; InvalidCountry                Code &#x3D; \&quot;InvalidCountry\&quot; DuplicatedEmailAddress        Code &#x3D; \&quot;DuplicatedEmailAddress\&quot; DuplicatedMember              Code &#x3D; \&quot;DuplicatedMember\&quot; EmptyMemberName               Code &#x3D; \&quot;EmptyMemberName\&quot; TheEmailAddressAlreadyExists  Code &#x3D; \&quot;TheEmailAddressAlreadyExists\&quot; NotMatchedWithEmail           Code &#x3D; \&quot;NotMatchedWithEmail\&quot; MaxMemberLimitError           Code &#x3D; \&quot;MaxMemberLimitError\&quot; DataValidationError           Code &#x3D; \&quot;DataValidationError\&quot; FailedToCreateUser            Code &#x3D; \&quot;FailedToCreateUser\&quot; FailedToCreateSubOrder        Code &#x3D; \&quot;FailedToCreateSubscriptionOrder\&quot; FailedToCreateOrganizationApp Code &#x3D; \&quot;FailedToCreateOrganizationApp\&quot; FailedToAddUserToFreePlan     Code &#x3D; \&quot;FailedToAddUserToFreePlan\&quot; FailedToAddUserToAdmin        Code &#x3D; \&quot;FailedToAddUserToAdmin\&quot; FailedToAddSubscription       Code &#x3D; \&quot;FailedToAddSubscription\&quot; FailedToCreateUserTags        Code &#x3D; \&quot;FailedToCreateUserTags\&quot; FailedToExecuteCallback       Code &#x3D; \&quot;FailedToExecuteCallback\&quot; FailedToSyncOrg               Code &#x3D; \&quot;FailedToSyncOrg\&quot; FailedToGetCompany            Code &#x3D; \&quot;FailedToGetCompany\&quot; FailedToSyncOrgApp            Code &#x3D; \&quot;FailedToSyncOrgApp\&quot; NotTheCreator                 Code &#x3D; \&quot;NotTheCreator\&quot; FailedToConvertPolicy         Code &#x3D; \&quot;FailedToConvertPolicy\&quot; FailedToSyncPolicy            Code &#x3D; \&quot;FailedToSyncPolicy\&quot; FailedToUpsertOrg             Code &#x3D; \&quot;FailedToUpsertOrg\&quot; FailedToSyncAc                Code &#x3D; \&quot;FailedToSyncAc\&quot; FailedToAddedDefaultPolicy    Code &#x3D; \&quot;FailedToAddedDefaultPolicy\&quot; FailedToUpsertAcGroups        Code &#x3D; \&quot;FailedToUpsertAcGroups\&quot; FailedToSyncSubscriptions     Code &#x3D; \&quot;FailedToSyncSubscriptions\&quot; EmailAlreadyExists            Code &#x3D; \&quot;EmailAlreadyExists\&quot; MobileAlreadyExists           Code &#x3D; \&quot;MobileAlreadyExists\&quot; InternalError                 Code &#x3D; \&quot;InternalError\&quot; OrgAlreadyExists              Code &#x3D; \&quot;OrgAlreadyExists\&quot; NotBelongThisOrg              Code &#x3D; \&quot;NotBelongThisOrg\&quot; FailedToTransferCreator       Code &#x3D; \&quot;FailedToTransferCreator\&quot; FailedToRemoveOrgMember       Code &#x3D; \&quot;FailedToRemoveOrgMember\&quot; FailedToGetAcGroup            Code &#x3D; \&quot;FailedToGetAcGroup\&quot; | [optional] 
**Detail** | Pointer to **string** | Detail contains further information on the nature of the error. | [optional] 
**Msg** | Pointer to **string** | Message is the error message. | [optional] 
**TraceId** | Pointer to **string** | TraceId is the identifier for a trace. It is globally unique. | [optional] 
**Type** | Pointer to **string** | Type A URI reference that identifies the error type. | [optional] 

## Methods

### NewJsonErrorResponse

`func NewJsonErrorResponse() *JsonErrorResponse`

NewJsonErrorResponse instantiates a new JsonErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonErrorResponseWithDefaults

`func NewJsonErrorResponseWithDefaults() *JsonErrorResponse`

NewJsonErrorResponseWithDefaults instantiates a new JsonErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *JsonErrorResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *JsonErrorResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *JsonErrorResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *JsonErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *JsonErrorResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *JsonErrorResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *JsonErrorResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *JsonErrorResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetMsg

`func (o *JsonErrorResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *JsonErrorResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *JsonErrorResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *JsonErrorResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetTraceId

`func (o *JsonErrorResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *JsonErrorResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *JsonErrorResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *JsonErrorResponse) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *JsonErrorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JsonErrorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JsonErrorResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JsonErrorResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


