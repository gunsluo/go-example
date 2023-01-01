# SubmitRecoveryFlowBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | Sending the anti-csrf token is only required for browser login flows. | [optional] 
**Method** | **string** | Method supports &#x60;captcha&#x60; only right now. | 
**Token** | Pointer to **string** | Recovery Token | [optional] 
**Traits** | [**RecoveryTraits**](RecoveryTraits.md) |  | 

## Methods

### NewSubmitRecoveryFlowBody

`func NewSubmitRecoveryFlowBody(method string, traits RecoveryTraits, ) *SubmitRecoveryFlowBody`

NewSubmitRecoveryFlowBody instantiates a new SubmitRecoveryFlowBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRecoveryFlowBodyWithDefaults

`func NewSubmitRecoveryFlowBodyWithDefaults() *SubmitRecoveryFlowBody`

NewSubmitRecoveryFlowBodyWithDefaults instantiates a new SubmitRecoveryFlowBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *SubmitRecoveryFlowBody) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *SubmitRecoveryFlowBody) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *SubmitRecoveryFlowBody) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *SubmitRecoveryFlowBody) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMethod

`func (o *SubmitRecoveryFlowBody) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubmitRecoveryFlowBody) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubmitRecoveryFlowBody) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetToken

`func (o *SubmitRecoveryFlowBody) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SubmitRecoveryFlowBody) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SubmitRecoveryFlowBody) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SubmitRecoveryFlowBody) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTraits

`func (o *SubmitRecoveryFlowBody) GetTraits() RecoveryTraits`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *SubmitRecoveryFlowBody) GetTraitsOk() (*RecoveryTraits, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *SubmitRecoveryFlowBody) SetTraits(v RecoveryTraits)`

SetTraits sets Traits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


