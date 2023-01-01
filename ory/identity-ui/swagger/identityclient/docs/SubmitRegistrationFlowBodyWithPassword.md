# SubmitRegistrationFlowBodyWithPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmPassword** | **string** | Password to sign the user up with | 
**CsrfToken** | Pointer to **string** | The CSRF Token | [optional] 
**Method** | **string** | Method to use  This field must be set to &#x60;password&#x60; when using the password method. | 
**Password** | **string** | Password to sign the user up with | 
**Traits** | [**RegistrationTraits**](RegistrationTraits.md) |  | 

## Methods

### NewSubmitRegistrationFlowBodyWithPassword

`func NewSubmitRegistrationFlowBodyWithPassword(confirmPassword string, method string, password string, traits RegistrationTraits, ) *SubmitRegistrationFlowBodyWithPassword`

NewSubmitRegistrationFlowBodyWithPassword instantiates a new SubmitRegistrationFlowBodyWithPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRegistrationFlowBodyWithPasswordWithDefaults

`func NewSubmitRegistrationFlowBodyWithPasswordWithDefaults() *SubmitRegistrationFlowBodyWithPassword`

NewSubmitRegistrationFlowBodyWithPasswordWithDefaults instantiates a new SubmitRegistrationFlowBodyWithPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmPassword

`func (o *SubmitRegistrationFlowBodyWithPassword) GetConfirmPassword() string`

GetConfirmPassword returns the ConfirmPassword field if non-nil, zero value otherwise.

### GetConfirmPasswordOk

`func (o *SubmitRegistrationFlowBodyWithPassword) GetConfirmPasswordOk() (*string, bool)`

GetConfirmPasswordOk returns a tuple with the ConfirmPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmPassword

`func (o *SubmitRegistrationFlowBodyWithPassword) SetConfirmPassword(v string)`

SetConfirmPassword sets ConfirmPassword field to given value.


### GetCsrfToken

`func (o *SubmitRegistrationFlowBodyWithPassword) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *SubmitRegistrationFlowBodyWithPassword) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *SubmitRegistrationFlowBodyWithPassword) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *SubmitRegistrationFlowBodyWithPassword) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMethod

`func (o *SubmitRegistrationFlowBodyWithPassword) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubmitRegistrationFlowBodyWithPassword) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubmitRegistrationFlowBodyWithPassword) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPassword

`func (o *SubmitRegistrationFlowBodyWithPassword) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SubmitRegistrationFlowBodyWithPassword) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SubmitRegistrationFlowBodyWithPassword) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetTraits

`func (o *SubmitRegistrationFlowBodyWithPassword) GetTraits() RegistrationTraits`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *SubmitRegistrationFlowBodyWithPassword) GetTraitsOk() (*RegistrationTraits, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *SubmitRegistrationFlowBodyWithPassword) SetTraits(v RegistrationTraits)`

SetTraits sets Traits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


