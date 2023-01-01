# SubmitSettingsFlowBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | The Anti-CSRF Token  This token is only required when performing browser flows. | [optional] 
**Method** | **string** | Method  Should be set to profile when trying to update a profile. | 
**Password** | **string** | Password is the updated password | 
**RecoveryCodesRegenerate** | Pointer to **bool** | RecoveryCodesRegenerate if true will generate new recovery codes | [optional] 
**TotpCode** | Pointer to **string** | Code must contain a valid TOTP based on the | [optional] 
**TotpUnlink** | Pointer to **bool** | UnlinkTOTP if true will remove the TOTP pairing, effectively removing the credential. This can be used to set up a new TOTP device. | [optional] 
**Token** | **string** | Token must contain a valid token based on the | 
**Traits** | [**IdentityTraits**](IdentityTraits.md) |  | 

## Methods

### NewSubmitSettingsFlowBody

`func NewSubmitSettingsFlowBody(method string, password string, token string, traits IdentityTraits, ) *SubmitSettingsFlowBody`

NewSubmitSettingsFlowBody instantiates a new SubmitSettingsFlowBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitSettingsFlowBodyWithDefaults

`func NewSubmitSettingsFlowBodyWithDefaults() *SubmitSettingsFlowBody`

NewSubmitSettingsFlowBodyWithDefaults instantiates a new SubmitSettingsFlowBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *SubmitSettingsFlowBody) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *SubmitSettingsFlowBody) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *SubmitSettingsFlowBody) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *SubmitSettingsFlowBody) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMethod

`func (o *SubmitSettingsFlowBody) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubmitSettingsFlowBody) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubmitSettingsFlowBody) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPassword

`func (o *SubmitSettingsFlowBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SubmitSettingsFlowBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SubmitSettingsFlowBody) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRecoveryCodesRegenerate

`func (o *SubmitSettingsFlowBody) GetRecoveryCodesRegenerate() bool`

GetRecoveryCodesRegenerate returns the RecoveryCodesRegenerate field if non-nil, zero value otherwise.

### GetRecoveryCodesRegenerateOk

`func (o *SubmitSettingsFlowBody) GetRecoveryCodesRegenerateOk() (*bool, bool)`

GetRecoveryCodesRegenerateOk returns a tuple with the RecoveryCodesRegenerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCodesRegenerate

`func (o *SubmitSettingsFlowBody) SetRecoveryCodesRegenerate(v bool)`

SetRecoveryCodesRegenerate sets RecoveryCodesRegenerate field to given value.

### HasRecoveryCodesRegenerate

`func (o *SubmitSettingsFlowBody) HasRecoveryCodesRegenerate() bool`

HasRecoveryCodesRegenerate returns a boolean if a field has been set.

### GetTotpCode

`func (o *SubmitSettingsFlowBody) GetTotpCode() string`

GetTotpCode returns the TotpCode field if non-nil, zero value otherwise.

### GetTotpCodeOk

`func (o *SubmitSettingsFlowBody) GetTotpCodeOk() (*string, bool)`

GetTotpCodeOk returns a tuple with the TotpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpCode

`func (o *SubmitSettingsFlowBody) SetTotpCode(v string)`

SetTotpCode sets TotpCode field to given value.

### HasTotpCode

`func (o *SubmitSettingsFlowBody) HasTotpCode() bool`

HasTotpCode returns a boolean if a field has been set.

### GetTotpUnlink

`func (o *SubmitSettingsFlowBody) GetTotpUnlink() bool`

GetTotpUnlink returns the TotpUnlink field if non-nil, zero value otherwise.

### GetTotpUnlinkOk

`func (o *SubmitSettingsFlowBody) GetTotpUnlinkOk() (*bool, bool)`

GetTotpUnlinkOk returns a tuple with the TotpUnlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpUnlink

`func (o *SubmitSettingsFlowBody) SetTotpUnlink(v bool)`

SetTotpUnlink sets TotpUnlink field to given value.

### HasTotpUnlink

`func (o *SubmitSettingsFlowBody) HasTotpUnlink() bool`

HasTotpUnlink returns a boolean if a field has been set.

### GetToken

`func (o *SubmitSettingsFlowBody) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SubmitSettingsFlowBody) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SubmitSettingsFlowBody) SetToken(v string)`

SetToken sets Token field to given value.


### GetTraits

`func (o *SubmitSettingsFlowBody) GetTraits() IdentityTraits`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *SubmitSettingsFlowBody) GetTraitsOk() (*IdentityTraits, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *SubmitSettingsFlowBody) SetTraits(v IdentityTraits)`

SetTraits sets Traits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


