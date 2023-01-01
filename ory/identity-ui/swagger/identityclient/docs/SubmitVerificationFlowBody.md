# SubmitVerificationFlowBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | Sending the anti-csrf token is only required for browser login flows. | [optional] 
**Method** | **string** | Method supports &#x60;captcha&#x60; only right now. | 
**Token** | Pointer to **string** | Verification Token | [optional] 
**Traits** | [**VerificationTraits**](VerificationTraits.md) |  | 

## Methods

### NewSubmitVerificationFlowBody

`func NewSubmitVerificationFlowBody(method string, traits VerificationTraits, ) *SubmitVerificationFlowBody`

NewSubmitVerificationFlowBody instantiates a new SubmitVerificationFlowBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitVerificationFlowBodyWithDefaults

`func NewSubmitVerificationFlowBodyWithDefaults() *SubmitVerificationFlowBody`

NewSubmitVerificationFlowBodyWithDefaults instantiates a new SubmitVerificationFlowBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *SubmitVerificationFlowBody) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *SubmitVerificationFlowBody) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *SubmitVerificationFlowBody) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *SubmitVerificationFlowBody) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMethod

`func (o *SubmitVerificationFlowBody) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubmitVerificationFlowBody) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubmitVerificationFlowBody) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetToken

`func (o *SubmitVerificationFlowBody) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SubmitVerificationFlowBody) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SubmitVerificationFlowBody) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SubmitVerificationFlowBody) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTraits

`func (o *SubmitVerificationFlowBody) GetTraits() VerificationTraits`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *SubmitVerificationFlowBody) GetTraitsOk() (*VerificationTraits, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *SubmitVerificationFlowBody) SetTraits(v VerificationTraits)`

SetTraits sets Traits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


