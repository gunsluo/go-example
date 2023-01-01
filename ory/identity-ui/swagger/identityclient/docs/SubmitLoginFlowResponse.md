# SubmitLoginFlowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | [**Session**](Session.md) |  | 
**SessionToken** | Pointer to **string** | The Session Token  A session token is equivalent to a session cookie, but it can be sent in the HTTP Authorization Header:  Authorization: bearer ${session-token}  The session token is only issued for API flows, not for Browser flows! | [optional] 

## Methods

### NewSubmitLoginFlowResponse

`func NewSubmitLoginFlowResponse(session Session, ) *SubmitLoginFlowResponse`

NewSubmitLoginFlowResponse instantiates a new SubmitLoginFlowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitLoginFlowResponseWithDefaults

`func NewSubmitLoginFlowResponseWithDefaults() *SubmitLoginFlowResponse`

NewSubmitLoginFlowResponseWithDefaults instantiates a new SubmitLoginFlowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *SubmitLoginFlowResponse) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *SubmitLoginFlowResponse) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *SubmitLoginFlowResponse) SetSession(v Session)`

SetSession sets Session field to given value.


### GetSessionToken

`func (o *SubmitLoginFlowResponse) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *SubmitLoginFlowResponse) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *SubmitLoginFlowResponse) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *SubmitLoginFlowResponse) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


