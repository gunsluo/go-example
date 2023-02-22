# ConsentFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is a helper struct field for. | [optional] 
**ExpiresAt** | **time.Time** | ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to log in, a new flow has to be initiated. | 
**Id** | **string** | ID represents the flow&#39;s unique ID. When performing the consent flow, this represents the id in the consent UI&#39;s query parameter: http://&lt;selfservice.flows.consent.ui_url&gt;/?flow&#x3D;&lt;flow_id&gt; | 
**Identity** | [**Identity**](Identity.md) |  | 
**IssuedAt** | **time.Time** | IssuedAt is the time (UTC) when the flow started. | 
**Method** | Pointer to **string** |  | [optional] 
**NotAsk** | Pointer to **bool** | NotAsk stores whether this consent flow should keep user session. | [optional] 
**Oauth2ConsentChallenge** | Pointer to **string** | OAuth 2.0 Consent Challenge.  This value is set using the &#x60;consent_challenge&#x60; query parameter of the registration and consent endpoints. If set will cooperate with OAuth2 and OpenID to act as an OAuth2 server / OpenID Provider. | [optional] 
**Oauth2ConsentRequest** | Pointer to [**OAuth2ConsentRequest**](OAuth2ConsentRequest.md) |  | [optional] 
**RequestUrl** | **string** | RequestURL is the initial URL that was requested from Identity. It can be used to forward information contained in the URL&#39;s path or query for example. | 
**ReturnTo** | Pointer to **string** | ReturnTo contains the requested return_to URL. | [optional] 
**State** | **string** | The state represents the state of the consent flow.  &#x60;&#x60;&#x60; accepted: ask the user to accepted rejected: reject consent by the user &#x60;&#x60;&#x60; | 
**Type** | [**FlowType**](FlowType.md) |  | 
**Ui** | [**UiContainer**](UiContainer.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is a helper struct field for. | [optional] 

## Methods

### NewConsentFlow

`func NewConsentFlow(expiresAt time.Time, id string, identity Identity, issuedAt time.Time, requestUrl string, state string, type_ FlowType, ui UiContainer, ) *ConsentFlow`

NewConsentFlow instantiates a new ConsentFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentFlowWithDefaults

`func NewConsentFlowWithDefaults() *ConsentFlow`

NewConsentFlowWithDefaults instantiates a new ConsentFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ConsentFlow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConsentFlow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConsentFlow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConsentFlow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ConsentFlow) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ConsentFlow) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ConsentFlow) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetId

`func (o *ConsentFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsentFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsentFlow) SetId(v string)`

SetId sets Id field to given value.


### GetIdentity

`func (o *ConsentFlow) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ConsentFlow) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ConsentFlow) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.


### GetIssuedAt

`func (o *ConsentFlow) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *ConsentFlow) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *ConsentFlow) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.


### GetMethod

`func (o *ConsentFlow) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ConsentFlow) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ConsentFlow) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ConsentFlow) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNotAsk

`func (o *ConsentFlow) GetNotAsk() bool`

GetNotAsk returns the NotAsk field if non-nil, zero value otherwise.

### GetNotAskOk

`func (o *ConsentFlow) GetNotAskOk() (*bool, bool)`

GetNotAskOk returns a tuple with the NotAsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAsk

`func (o *ConsentFlow) SetNotAsk(v bool)`

SetNotAsk sets NotAsk field to given value.

### HasNotAsk

`func (o *ConsentFlow) HasNotAsk() bool`

HasNotAsk returns a boolean if a field has been set.

### GetOauth2ConsentChallenge

`func (o *ConsentFlow) GetOauth2ConsentChallenge() string`

GetOauth2ConsentChallenge returns the Oauth2ConsentChallenge field if non-nil, zero value otherwise.

### GetOauth2ConsentChallengeOk

`func (o *ConsentFlow) GetOauth2ConsentChallengeOk() (*string, bool)`

GetOauth2ConsentChallengeOk returns a tuple with the Oauth2ConsentChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ConsentChallenge

`func (o *ConsentFlow) SetOauth2ConsentChallenge(v string)`

SetOauth2ConsentChallenge sets Oauth2ConsentChallenge field to given value.

### HasOauth2ConsentChallenge

`func (o *ConsentFlow) HasOauth2ConsentChallenge() bool`

HasOauth2ConsentChallenge returns a boolean if a field has been set.

### GetOauth2ConsentRequest

`func (o *ConsentFlow) GetOauth2ConsentRequest() OAuth2ConsentRequest`

GetOauth2ConsentRequest returns the Oauth2ConsentRequest field if non-nil, zero value otherwise.

### GetOauth2ConsentRequestOk

`func (o *ConsentFlow) GetOauth2ConsentRequestOk() (*OAuth2ConsentRequest, bool)`

GetOauth2ConsentRequestOk returns a tuple with the Oauth2ConsentRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ConsentRequest

`func (o *ConsentFlow) SetOauth2ConsentRequest(v OAuth2ConsentRequest)`

SetOauth2ConsentRequest sets Oauth2ConsentRequest field to given value.

### HasOauth2ConsentRequest

`func (o *ConsentFlow) HasOauth2ConsentRequest() bool`

HasOauth2ConsentRequest returns a boolean if a field has been set.

### GetRequestUrl

`func (o *ConsentFlow) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *ConsentFlow) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *ConsentFlow) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.


### GetReturnTo

`func (o *ConsentFlow) GetReturnTo() string`

GetReturnTo returns the ReturnTo field if non-nil, zero value otherwise.

### GetReturnToOk

`func (o *ConsentFlow) GetReturnToOk() (*string, bool)`

GetReturnToOk returns a tuple with the ReturnTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTo

`func (o *ConsentFlow) SetReturnTo(v string)`

SetReturnTo sets ReturnTo field to given value.

### HasReturnTo

`func (o *ConsentFlow) HasReturnTo() bool`

HasReturnTo returns a boolean if a field has been set.

### GetState

`func (o *ConsentFlow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConsentFlow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConsentFlow) SetState(v string)`

SetState sets State field to given value.


### GetType

`func (o *ConsentFlow) GetType() FlowType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsentFlow) GetTypeOk() (*FlowType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsentFlow) SetType(v FlowType)`

SetType sets Type field to given value.


### GetUi

`func (o *ConsentFlow) GetUi() UiContainer`

GetUi returns the Ui field if non-nil, zero value otherwise.

### GetUiOk

`func (o *ConsentFlow) GetUiOk() (*UiContainer, bool)`

GetUiOk returns a tuple with the Ui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUi

`func (o *ConsentFlow) SetUi(v UiContainer)`

SetUi sets Ui field to given value.


### GetUpdatedAt

`func (o *ConsentFlow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConsentFlow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConsentFlow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ConsentFlow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


