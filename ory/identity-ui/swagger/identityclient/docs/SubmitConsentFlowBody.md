# SubmitConsentFlowBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action is allow or deny. | 
**CsrfToken** | Pointer to **string** | Sending the anti-csrf token is only required for browser login flows. | [optional] 
**Method** | **string** | Method should be set to \&quot;consent\&quot;. | 
**NotAsk** | **bool** | Do Not Ask Me  If set to true, don&#39;t ask me again. | 
**Scope** | **[]string** | scope is selected by the user. | 

## Methods

### NewSubmitConsentFlowBody

`func NewSubmitConsentFlowBody(action string, method string, notAsk bool, scope []string, ) *SubmitConsentFlowBody`

NewSubmitConsentFlowBody instantiates a new SubmitConsentFlowBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitConsentFlowBodyWithDefaults

`func NewSubmitConsentFlowBodyWithDefaults() *SubmitConsentFlowBody`

NewSubmitConsentFlowBodyWithDefaults instantiates a new SubmitConsentFlowBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SubmitConsentFlowBody) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SubmitConsentFlowBody) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SubmitConsentFlowBody) SetAction(v string)`

SetAction sets Action field to given value.


### GetCsrfToken

`func (o *SubmitConsentFlowBody) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *SubmitConsentFlowBody) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *SubmitConsentFlowBody) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *SubmitConsentFlowBody) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMethod

`func (o *SubmitConsentFlowBody) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubmitConsentFlowBody) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubmitConsentFlowBody) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetNotAsk

`func (o *SubmitConsentFlowBody) GetNotAsk() bool`

GetNotAsk returns the NotAsk field if non-nil, zero value otherwise.

### GetNotAskOk

`func (o *SubmitConsentFlowBody) GetNotAskOk() (*bool, bool)`

GetNotAskOk returns a tuple with the NotAsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAsk

`func (o *SubmitConsentFlowBody) SetNotAsk(v bool)`

SetNotAsk sets NotAsk field to given value.


### GetScope

`func (o *SubmitConsentFlowBody) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SubmitConsentFlowBody) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SubmitConsentFlowBody) SetScope(v []string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


