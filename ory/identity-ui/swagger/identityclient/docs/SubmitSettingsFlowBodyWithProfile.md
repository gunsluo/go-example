# SubmitSettingsFlowBodyWithProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | The Anti-CSRF Token  This token is only required when performing browser flows. | [optional] 
**Method** | **string** | Method  Should be set to profile when trying to update a profile. | 
**Token** | **string** | Token must contain a valid token based on the | 
**Traits** | [**IdentityTraits**](IdentityTraits.md) |  | 

## Methods

### NewSubmitSettingsFlowBodyWithProfile

`func NewSubmitSettingsFlowBodyWithProfile(method string, token string, traits IdentityTraits, ) *SubmitSettingsFlowBodyWithProfile`

NewSubmitSettingsFlowBodyWithProfile instantiates a new SubmitSettingsFlowBodyWithProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitSettingsFlowBodyWithProfileWithDefaults

`func NewSubmitSettingsFlowBodyWithProfileWithDefaults() *SubmitSettingsFlowBodyWithProfile`

NewSubmitSettingsFlowBodyWithProfileWithDefaults instantiates a new SubmitSettingsFlowBodyWithProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *SubmitSettingsFlowBodyWithProfile) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *SubmitSettingsFlowBodyWithProfile) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *SubmitSettingsFlowBodyWithProfile) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *SubmitSettingsFlowBodyWithProfile) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMethod

`func (o *SubmitSettingsFlowBodyWithProfile) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubmitSettingsFlowBodyWithProfile) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubmitSettingsFlowBodyWithProfile) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetToken

`func (o *SubmitSettingsFlowBodyWithProfile) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SubmitSettingsFlowBodyWithProfile) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SubmitSettingsFlowBodyWithProfile) SetToken(v string)`

SetToken sets Token field to given value.


### GetTraits

`func (o *SubmitSettingsFlowBodyWithProfile) GetTraits() IdentityTraits`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *SubmitSettingsFlowBodyWithProfile) GetTraitsOk() (*IdentityTraits, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *SubmitSettingsFlowBodyWithProfile) SetTraits(v IdentityTraits)`

SetTraits sets Traits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


