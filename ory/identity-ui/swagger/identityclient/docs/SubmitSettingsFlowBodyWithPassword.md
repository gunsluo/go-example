# SubmitSettingsFlowBodyWithPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | CSRFToken is the anti-CSRF token | [optional] 
**Method** | **string** | Method  Should be set to password when trying to update a password. | 
**Password** | **string** | Password is the updated password | 

## Methods

### NewSubmitSettingsFlowBodyWithPassword

`func NewSubmitSettingsFlowBodyWithPassword(method string, password string, ) *SubmitSettingsFlowBodyWithPassword`

NewSubmitSettingsFlowBodyWithPassword instantiates a new SubmitSettingsFlowBodyWithPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitSettingsFlowBodyWithPasswordWithDefaults

`func NewSubmitSettingsFlowBodyWithPasswordWithDefaults() *SubmitSettingsFlowBodyWithPassword`

NewSubmitSettingsFlowBodyWithPasswordWithDefaults instantiates a new SubmitSettingsFlowBodyWithPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *SubmitSettingsFlowBodyWithPassword) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *SubmitSettingsFlowBodyWithPassword) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *SubmitSettingsFlowBodyWithPassword) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *SubmitSettingsFlowBodyWithPassword) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMethod

`func (o *SubmitSettingsFlowBodyWithPassword) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubmitSettingsFlowBodyWithPassword) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubmitSettingsFlowBodyWithPassword) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPassword

`func (o *SubmitSettingsFlowBodyWithPassword) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SubmitSettingsFlowBodyWithPassword) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SubmitSettingsFlowBodyWithPassword) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


