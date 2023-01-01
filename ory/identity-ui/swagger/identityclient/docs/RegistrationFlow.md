# RegistrationFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to [**CredentialsType**](CredentialsType.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | CreatedAt is a helper struct field for. | [optional] 
**ExpiresAt** | **time.Time** | ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to log in, a new flow has to be initiated. | 
**Id** | **string** | ID represents the flow&#39;s unique ID. When performing the login flow, this represents the id in the login UI&#39;s query parameter: http://&lt;selfservice.flows.login.ui_url&gt;/?flow&#x3D;&lt;flow_id&gt; | 
**IssuedAt** | **time.Time** | IssuedAt is the time (UTC) when the flow started. | 
**Oauth2LoginChallenge** | Pointer to **string** | OAuth 2.0 Login Challenge.  This value is set using the &#x60;login_challenge&#x60; query parameter of the registration and login endpoints. If set will cooperate with OAuth2 and OpenID to act as an OAuth2 server / OpenID Provider. | [optional] 
**Oauth2LoginRequest** | Pointer to [**OAuth2LoginRequest**](OAuth2LoginRequest.md) |  | [optional] 
**RequestUrl** | **string** | RequestURL is the initial URL that was requested from Identity. It can be used to forward information contained in the URL&#39;s path or query for example. | 
**ReturnTo** | Pointer to **string** | ReturnTo contains the requested return_to URL. | [optional] 
**Type** | [**FlowType**](FlowType.md) |  | 
**Ui** | [**UiContainer**](UiContainer.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is a helper struct field for. | [optional] 

## Methods

### NewRegistrationFlow

`func NewRegistrationFlow(expiresAt time.Time, id string, issuedAt time.Time, requestUrl string, type_ FlowType, ui UiContainer, ) *RegistrationFlow`

NewRegistrationFlow instantiates a new RegistrationFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationFlowWithDefaults

`func NewRegistrationFlowWithDefaults() *RegistrationFlow`

NewRegistrationFlowWithDefaults instantiates a new RegistrationFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *RegistrationFlow) GetActive() CredentialsType`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RegistrationFlow) GetActiveOk() (*CredentialsType, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RegistrationFlow) SetActive(v CredentialsType)`

SetActive sets Active field to given value.

### HasActive

`func (o *RegistrationFlow) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RegistrationFlow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RegistrationFlow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RegistrationFlow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RegistrationFlow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *RegistrationFlow) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *RegistrationFlow) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *RegistrationFlow) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetId

`func (o *RegistrationFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistrationFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistrationFlow) SetId(v string)`

SetId sets Id field to given value.


### GetIssuedAt

`func (o *RegistrationFlow) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *RegistrationFlow) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *RegistrationFlow) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.


### GetOauth2LoginChallenge

`func (o *RegistrationFlow) GetOauth2LoginChallenge() string`

GetOauth2LoginChallenge returns the Oauth2LoginChallenge field if non-nil, zero value otherwise.

### GetOauth2LoginChallengeOk

`func (o *RegistrationFlow) GetOauth2LoginChallengeOk() (*string, bool)`

GetOauth2LoginChallengeOk returns a tuple with the Oauth2LoginChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2LoginChallenge

`func (o *RegistrationFlow) SetOauth2LoginChallenge(v string)`

SetOauth2LoginChallenge sets Oauth2LoginChallenge field to given value.

### HasOauth2LoginChallenge

`func (o *RegistrationFlow) HasOauth2LoginChallenge() bool`

HasOauth2LoginChallenge returns a boolean if a field has been set.

### GetOauth2LoginRequest

`func (o *RegistrationFlow) GetOauth2LoginRequest() OAuth2LoginRequest`

GetOauth2LoginRequest returns the Oauth2LoginRequest field if non-nil, zero value otherwise.

### GetOauth2LoginRequestOk

`func (o *RegistrationFlow) GetOauth2LoginRequestOk() (*OAuth2LoginRequest, bool)`

GetOauth2LoginRequestOk returns a tuple with the Oauth2LoginRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2LoginRequest

`func (o *RegistrationFlow) SetOauth2LoginRequest(v OAuth2LoginRequest)`

SetOauth2LoginRequest sets Oauth2LoginRequest field to given value.

### HasOauth2LoginRequest

`func (o *RegistrationFlow) HasOauth2LoginRequest() bool`

HasOauth2LoginRequest returns a boolean if a field has been set.

### GetRequestUrl

`func (o *RegistrationFlow) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *RegistrationFlow) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *RegistrationFlow) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.


### GetReturnTo

`func (o *RegistrationFlow) GetReturnTo() string`

GetReturnTo returns the ReturnTo field if non-nil, zero value otherwise.

### GetReturnToOk

`func (o *RegistrationFlow) GetReturnToOk() (*string, bool)`

GetReturnToOk returns a tuple with the ReturnTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTo

`func (o *RegistrationFlow) SetReturnTo(v string)`

SetReturnTo sets ReturnTo field to given value.

### HasReturnTo

`func (o *RegistrationFlow) HasReturnTo() bool`

HasReturnTo returns a boolean if a field has been set.

### GetType

`func (o *RegistrationFlow) GetType() FlowType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistrationFlow) GetTypeOk() (*FlowType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistrationFlow) SetType(v FlowType)`

SetType sets Type field to given value.


### GetUi

`func (o *RegistrationFlow) GetUi() UiContainer`

GetUi returns the Ui field if non-nil, zero value otherwise.

### GetUiOk

`func (o *RegistrationFlow) GetUiOk() (*UiContainer, bool)`

GetUiOk returns a tuple with the Ui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUi

`func (o *RegistrationFlow) SetUi(v UiContainer)`

SetUi sets Ui field to given value.


### GetUpdatedAt

`func (o *RegistrationFlow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RegistrationFlow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RegistrationFlow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RegistrationFlow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


