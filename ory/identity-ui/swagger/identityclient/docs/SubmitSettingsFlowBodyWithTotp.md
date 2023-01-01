# SubmitSettingsFlowBodyWithTotp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | CSRFToken is the anti-CSRF token | [optional] 
**Method** | **string** | Method  Should be set to \&quot;totp\&quot; when trying to add, update, or remove a totp pairing. | 
**RecoveryCodesRegenerate** | Pointer to **bool** | RecoveryCodesRegenerate if true will generate new recovery codes | [optional] 
**TotpCode** | Pointer to **string** | Code must contain a valid TOTP based on the | [optional] 
**TotpUnlink** | Pointer to **bool** | UnlinkTOTP if true will remove the TOTP pairing, effectively removing the credential. This can be used to set up a new TOTP device. | [optional] 

## Methods

### NewSubmitSettingsFlowBodyWithTotp

`func NewSubmitSettingsFlowBodyWithTotp(method string, ) *SubmitSettingsFlowBodyWithTotp`

NewSubmitSettingsFlowBodyWithTotp instantiates a new SubmitSettingsFlowBodyWithTotp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitSettingsFlowBodyWithTotpWithDefaults

`func NewSubmitSettingsFlowBodyWithTotpWithDefaults() *SubmitSettingsFlowBodyWithTotp`

NewSubmitSettingsFlowBodyWithTotpWithDefaults instantiates a new SubmitSettingsFlowBodyWithTotp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *SubmitSettingsFlowBodyWithTotp) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *SubmitSettingsFlowBodyWithTotp) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *SubmitSettingsFlowBodyWithTotp) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *SubmitSettingsFlowBodyWithTotp) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMethod

`func (o *SubmitSettingsFlowBodyWithTotp) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubmitSettingsFlowBodyWithTotp) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubmitSettingsFlowBodyWithTotp) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetRecoveryCodesRegenerate

`func (o *SubmitSettingsFlowBodyWithTotp) GetRecoveryCodesRegenerate() bool`

GetRecoveryCodesRegenerate returns the RecoveryCodesRegenerate field if non-nil, zero value otherwise.

### GetRecoveryCodesRegenerateOk

`func (o *SubmitSettingsFlowBodyWithTotp) GetRecoveryCodesRegenerateOk() (*bool, bool)`

GetRecoveryCodesRegenerateOk returns a tuple with the RecoveryCodesRegenerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCodesRegenerate

`func (o *SubmitSettingsFlowBodyWithTotp) SetRecoveryCodesRegenerate(v bool)`

SetRecoveryCodesRegenerate sets RecoveryCodesRegenerate field to given value.

### HasRecoveryCodesRegenerate

`func (o *SubmitSettingsFlowBodyWithTotp) HasRecoveryCodesRegenerate() bool`

HasRecoveryCodesRegenerate returns a boolean if a field has been set.

### GetTotpCode

`func (o *SubmitSettingsFlowBodyWithTotp) GetTotpCode() string`

GetTotpCode returns the TotpCode field if non-nil, zero value otherwise.

### GetTotpCodeOk

`func (o *SubmitSettingsFlowBodyWithTotp) GetTotpCodeOk() (*string, bool)`

GetTotpCodeOk returns a tuple with the TotpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpCode

`func (o *SubmitSettingsFlowBodyWithTotp) SetTotpCode(v string)`

SetTotpCode sets TotpCode field to given value.

### HasTotpCode

`func (o *SubmitSettingsFlowBodyWithTotp) HasTotpCode() bool`

HasTotpCode returns a boolean if a field has been set.

### GetTotpUnlink

`func (o *SubmitSettingsFlowBodyWithTotp) GetTotpUnlink() bool`

GetTotpUnlink returns the TotpUnlink field if non-nil, zero value otherwise.

### GetTotpUnlinkOk

`func (o *SubmitSettingsFlowBodyWithTotp) GetTotpUnlinkOk() (*bool, bool)`

GetTotpUnlinkOk returns a tuple with the TotpUnlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpUnlink

`func (o *SubmitSettingsFlowBodyWithTotp) SetTotpUnlink(v bool)`

SetTotpUnlink sets TotpUnlink field to given value.

### HasTotpUnlink

`func (o *SubmitSettingsFlowBodyWithTotp) HasTotpUnlink() bool`

HasTotpUnlink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


