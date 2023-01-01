# SubmitLoginFlowBodyWithTotp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeAuthCode** | Pointer to **string** | The ChangeAuthCode  recovery_code | totp_code | [optional] 
**CsrfToken** | Pointer to **string** | Sending the anti-csrf token is only required for browser login flows. | [optional] 
**Method** | **string** | Method should be set to \&quot;totp\&quot; when logging in using the TOTP strategy. | 
**RecoveryCode** | Pointer to **string** | The Recovery code. | [optional] 
**TotpCode** | Pointer to **string** | The TOTP code. | [optional] 

## Methods

### NewSubmitLoginFlowBodyWithTotp

`func NewSubmitLoginFlowBodyWithTotp(method string, ) *SubmitLoginFlowBodyWithTotp`

NewSubmitLoginFlowBodyWithTotp instantiates a new SubmitLoginFlowBodyWithTotp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitLoginFlowBodyWithTotpWithDefaults

`func NewSubmitLoginFlowBodyWithTotpWithDefaults() *SubmitLoginFlowBodyWithTotp`

NewSubmitLoginFlowBodyWithTotpWithDefaults instantiates a new SubmitLoginFlowBodyWithTotp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeAuthCode

`func (o *SubmitLoginFlowBodyWithTotp) GetChangeAuthCode() string`

GetChangeAuthCode returns the ChangeAuthCode field if non-nil, zero value otherwise.

### GetChangeAuthCodeOk

`func (o *SubmitLoginFlowBodyWithTotp) GetChangeAuthCodeOk() (*string, bool)`

GetChangeAuthCodeOk returns a tuple with the ChangeAuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAuthCode

`func (o *SubmitLoginFlowBodyWithTotp) SetChangeAuthCode(v string)`

SetChangeAuthCode sets ChangeAuthCode field to given value.

### HasChangeAuthCode

`func (o *SubmitLoginFlowBodyWithTotp) HasChangeAuthCode() bool`

HasChangeAuthCode returns a boolean if a field has been set.

### GetCsrfToken

`func (o *SubmitLoginFlowBodyWithTotp) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *SubmitLoginFlowBodyWithTotp) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *SubmitLoginFlowBodyWithTotp) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *SubmitLoginFlowBodyWithTotp) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMethod

`func (o *SubmitLoginFlowBodyWithTotp) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubmitLoginFlowBodyWithTotp) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubmitLoginFlowBodyWithTotp) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetRecoveryCode

`func (o *SubmitLoginFlowBodyWithTotp) GetRecoveryCode() string`

GetRecoveryCode returns the RecoveryCode field if non-nil, zero value otherwise.

### GetRecoveryCodeOk

`func (o *SubmitLoginFlowBodyWithTotp) GetRecoveryCodeOk() (*string, bool)`

GetRecoveryCodeOk returns a tuple with the RecoveryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCode

`func (o *SubmitLoginFlowBodyWithTotp) SetRecoveryCode(v string)`

SetRecoveryCode sets RecoveryCode field to given value.

### HasRecoveryCode

`func (o *SubmitLoginFlowBodyWithTotp) HasRecoveryCode() bool`

HasRecoveryCode returns a boolean if a field has been set.

### GetTotpCode

`func (o *SubmitLoginFlowBodyWithTotp) GetTotpCode() string`

GetTotpCode returns the TotpCode field if non-nil, zero value otherwise.

### GetTotpCodeOk

`func (o *SubmitLoginFlowBodyWithTotp) GetTotpCodeOk() (*string, bool)`

GetTotpCodeOk returns a tuple with the TotpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpCode

`func (o *SubmitLoginFlowBodyWithTotp) SetTotpCode(v string)`

SetTotpCode sets TotpCode field to given value.

### HasTotpCode

`func (o *SubmitLoginFlowBodyWithTotp) HasTotpCode() bool`

HasTotpCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


