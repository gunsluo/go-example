# SubmitLoginFlowBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | The CSRF Token | [optional] 
**Identifier** | **string** | Identifier is the email or username of the user trying to log in. | 
**Method** | **string** | Method to use  This field must be set to &#x60;oidc&#x60; when using the oidc method. | 
**Password** | **string** | The user&#39;s password. | 
**ChangeAuthCode** | Pointer to **string** | The ChangeAuthCode  recovery_code | totp_code | [optional] 
**RecoveryCode** | Pointer to **string** | The Recovery code. | [optional] 
**TotpCode** | Pointer to **string** | The TOTP code. | [optional] 
**Provider** | **string** | The provider to register with | 

## Methods

### NewSubmitLoginFlowBody

`func NewSubmitLoginFlowBody(identifier string, method string, password string, provider string, ) *SubmitLoginFlowBody`

NewSubmitLoginFlowBody instantiates a new SubmitLoginFlowBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitLoginFlowBodyWithDefaults

`func NewSubmitLoginFlowBodyWithDefaults() *SubmitLoginFlowBody`

NewSubmitLoginFlowBodyWithDefaults instantiates a new SubmitLoginFlowBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *SubmitLoginFlowBody) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *SubmitLoginFlowBody) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *SubmitLoginFlowBody) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *SubmitLoginFlowBody) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetIdentifier

`func (o *SubmitLoginFlowBody) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *SubmitLoginFlowBody) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *SubmitLoginFlowBody) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetMethod

`func (o *SubmitLoginFlowBody) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubmitLoginFlowBody) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubmitLoginFlowBody) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPassword

`func (o *SubmitLoginFlowBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SubmitLoginFlowBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SubmitLoginFlowBody) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetChangeAuthCode

`func (o *SubmitLoginFlowBody) GetChangeAuthCode() string`

GetChangeAuthCode returns the ChangeAuthCode field if non-nil, zero value otherwise.

### GetChangeAuthCodeOk

`func (o *SubmitLoginFlowBody) GetChangeAuthCodeOk() (*string, bool)`

GetChangeAuthCodeOk returns a tuple with the ChangeAuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAuthCode

`func (o *SubmitLoginFlowBody) SetChangeAuthCode(v string)`

SetChangeAuthCode sets ChangeAuthCode field to given value.

### HasChangeAuthCode

`func (o *SubmitLoginFlowBody) HasChangeAuthCode() bool`

HasChangeAuthCode returns a boolean if a field has been set.

### GetRecoveryCode

`func (o *SubmitLoginFlowBody) GetRecoveryCode() string`

GetRecoveryCode returns the RecoveryCode field if non-nil, zero value otherwise.

### GetRecoveryCodeOk

`func (o *SubmitLoginFlowBody) GetRecoveryCodeOk() (*string, bool)`

GetRecoveryCodeOk returns a tuple with the RecoveryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCode

`func (o *SubmitLoginFlowBody) SetRecoveryCode(v string)`

SetRecoveryCode sets RecoveryCode field to given value.

### HasRecoveryCode

`func (o *SubmitLoginFlowBody) HasRecoveryCode() bool`

HasRecoveryCode returns a boolean if a field has been set.

### GetTotpCode

`func (o *SubmitLoginFlowBody) GetTotpCode() string`

GetTotpCode returns the TotpCode field if non-nil, zero value otherwise.

### GetTotpCodeOk

`func (o *SubmitLoginFlowBody) GetTotpCodeOk() (*string, bool)`

GetTotpCodeOk returns a tuple with the TotpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpCode

`func (o *SubmitLoginFlowBody) SetTotpCode(v string)`

SetTotpCode sets TotpCode field to given value.

### HasTotpCode

`func (o *SubmitLoginFlowBody) HasTotpCode() bool`

HasTotpCode returns a boolean if a field has been set.

### GetProvider

`func (o *SubmitLoginFlowBody) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SubmitLoginFlowBody) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SubmitLoginFlowBody) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


