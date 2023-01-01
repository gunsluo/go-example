# SubmitLoginFlowBodyWithPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | Sending the anti-csrf token is only required for browser login flows. | [optional] 
**Identifier** | **string** | Identifier is the email or username of the user trying to log in. | 
**Method** | **string** | Method should be set to \&quot;password\&quot; when logging in using the identifier and password strategy. | 
**Password** | **string** | The user&#39;s password. | 

## Methods

### NewSubmitLoginFlowBodyWithPassword

`func NewSubmitLoginFlowBodyWithPassword(identifier string, method string, password string, ) *SubmitLoginFlowBodyWithPassword`

NewSubmitLoginFlowBodyWithPassword instantiates a new SubmitLoginFlowBodyWithPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitLoginFlowBodyWithPasswordWithDefaults

`func NewSubmitLoginFlowBodyWithPasswordWithDefaults() *SubmitLoginFlowBodyWithPassword`

NewSubmitLoginFlowBodyWithPasswordWithDefaults instantiates a new SubmitLoginFlowBodyWithPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *SubmitLoginFlowBodyWithPassword) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *SubmitLoginFlowBodyWithPassword) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *SubmitLoginFlowBodyWithPassword) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *SubmitLoginFlowBodyWithPassword) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetIdentifier

`func (o *SubmitLoginFlowBodyWithPassword) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *SubmitLoginFlowBodyWithPassword) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *SubmitLoginFlowBodyWithPassword) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetMethod

`func (o *SubmitLoginFlowBodyWithPassword) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubmitLoginFlowBodyWithPassword) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubmitLoginFlowBodyWithPassword) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPassword

`func (o *SubmitLoginFlowBodyWithPassword) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SubmitLoginFlowBodyWithPassword) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SubmitLoginFlowBodyWithPassword) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


