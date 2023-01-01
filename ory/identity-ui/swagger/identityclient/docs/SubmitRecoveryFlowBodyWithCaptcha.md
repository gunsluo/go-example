# SubmitRecoveryFlowBodyWithCaptcha

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | Sending the anti-csrf token is only required for browser login flows. | [optional] 
**Method** | **string** | Method supports &#x60;captcha&#x60; only right now. | 
**Token** | Pointer to **string** | Recovery Token | [optional] 
**Traits** | [**RecoveryTraits**](RecoveryTraits.md) |  | 

## Methods

### NewSubmitRecoveryFlowBodyWithCaptcha

`func NewSubmitRecoveryFlowBodyWithCaptcha(method string, traits RecoveryTraits, ) *SubmitRecoveryFlowBodyWithCaptcha`

NewSubmitRecoveryFlowBodyWithCaptcha instantiates a new SubmitRecoveryFlowBodyWithCaptcha object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRecoveryFlowBodyWithCaptchaWithDefaults

`func NewSubmitRecoveryFlowBodyWithCaptchaWithDefaults() *SubmitRecoveryFlowBodyWithCaptcha`

NewSubmitRecoveryFlowBodyWithCaptchaWithDefaults instantiates a new SubmitRecoveryFlowBodyWithCaptcha object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *SubmitRecoveryFlowBodyWithCaptcha) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *SubmitRecoveryFlowBodyWithCaptcha) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *SubmitRecoveryFlowBodyWithCaptcha) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *SubmitRecoveryFlowBodyWithCaptcha) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMethod

`func (o *SubmitRecoveryFlowBodyWithCaptcha) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubmitRecoveryFlowBodyWithCaptcha) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubmitRecoveryFlowBodyWithCaptcha) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetToken

`func (o *SubmitRecoveryFlowBodyWithCaptcha) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SubmitRecoveryFlowBodyWithCaptcha) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SubmitRecoveryFlowBodyWithCaptcha) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SubmitRecoveryFlowBodyWithCaptcha) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTraits

`func (o *SubmitRecoveryFlowBodyWithCaptcha) GetTraits() RecoveryTraits`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *SubmitRecoveryFlowBodyWithCaptcha) GetTraitsOk() (*RecoveryTraits, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *SubmitRecoveryFlowBodyWithCaptcha) SetTraits(v RecoveryTraits)`

SetTraits sets Traits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


