# SubmitRegistrationFlowBodyWithOidc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | The CSRF Token | [optional] 
**Method** | **string** | Method to use  This field must be set to &#x60;oidc&#x60; when using the oidc method. | 
**Provider** | **string** | The provider to register with | 
**Traits** | [**RegistrationTraits**](RegistrationTraits.md) |  | 

## Methods

### NewSubmitRegistrationFlowBodyWithOidc

`func NewSubmitRegistrationFlowBodyWithOidc(method string, provider string, traits RegistrationTraits, ) *SubmitRegistrationFlowBodyWithOidc`

NewSubmitRegistrationFlowBodyWithOidc instantiates a new SubmitRegistrationFlowBodyWithOidc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRegistrationFlowBodyWithOidcWithDefaults

`func NewSubmitRegistrationFlowBodyWithOidcWithDefaults() *SubmitRegistrationFlowBodyWithOidc`

NewSubmitRegistrationFlowBodyWithOidcWithDefaults instantiates a new SubmitRegistrationFlowBodyWithOidc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *SubmitRegistrationFlowBodyWithOidc) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *SubmitRegistrationFlowBodyWithOidc) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *SubmitRegistrationFlowBodyWithOidc) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *SubmitRegistrationFlowBodyWithOidc) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMethod

`func (o *SubmitRegistrationFlowBodyWithOidc) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubmitRegistrationFlowBodyWithOidc) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubmitRegistrationFlowBodyWithOidc) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetProvider

`func (o *SubmitRegistrationFlowBodyWithOidc) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SubmitRegistrationFlowBodyWithOidc) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SubmitRegistrationFlowBodyWithOidc) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetTraits

`func (o *SubmitRegistrationFlowBodyWithOidc) GetTraits() RegistrationTraits`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *SubmitRegistrationFlowBodyWithOidc) GetTraitsOk() (*RegistrationTraits, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *SubmitRegistrationFlowBodyWithOidc) SetTraits(v RegistrationTraits)`

SetTraits sets Traits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


