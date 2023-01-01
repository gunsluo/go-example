# SettingsFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to [**CredentialsType**](CredentialsType.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | CreatedAt is a helper struct field for gobuffalo.pop. | [optional] 
**ExpiresAt** | **time.Time** | ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to update the setting, a new flow has to be initiated. | 
**Id** | **string** | ID represents the flow&#39;s unique ID. When performing the settings flow, this represents the id in the settings ui&#39;s query parameter: http://&lt;selfservice.flows.settings.ui_url&gt;?flow&#x3D;&lt;id&gt; | 
**Identity** | [**Identity**](Identity.md) |  | 
**IssuedAt** | **time.Time** | IssuedAt is the time (UTC) when the flow occurred. | 
**RequestUrl** | **string** | RequestURL is the initial URL that was requested from Identity. It can be used to forward information contained in the URL&#39;s path or query for example. | 
**ReturnTo** | Pointer to **string** | ReturnTo contains the requested return_to URL. | [optional] 
**State** | [**SettingsState**](SettingsState.md) |  | 
**Type** | [**FlowType**](FlowType.md) |  | 
**Ui** | [**UiContainer**](UiContainer.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is a helper struct field for gobuffalo.pop. | [optional] 

## Methods

### NewSettingsFlow

`func NewSettingsFlow(expiresAt time.Time, id string, identity Identity, issuedAt time.Time, requestUrl string, state SettingsState, type_ FlowType, ui UiContainer, ) *SettingsFlow`

NewSettingsFlow instantiates a new SettingsFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsFlowWithDefaults

`func NewSettingsFlowWithDefaults() *SettingsFlow`

NewSettingsFlowWithDefaults instantiates a new SettingsFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *SettingsFlow) GetActive() CredentialsType`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SettingsFlow) GetActiveOk() (*CredentialsType, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SettingsFlow) SetActive(v CredentialsType)`

SetActive sets Active field to given value.

### HasActive

`func (o *SettingsFlow) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SettingsFlow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SettingsFlow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SettingsFlow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SettingsFlow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *SettingsFlow) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SettingsFlow) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SettingsFlow) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetId

`func (o *SettingsFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SettingsFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SettingsFlow) SetId(v string)`

SetId sets Id field to given value.


### GetIdentity

`func (o *SettingsFlow) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *SettingsFlow) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *SettingsFlow) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.


### GetIssuedAt

`func (o *SettingsFlow) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *SettingsFlow) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *SettingsFlow) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.


### GetRequestUrl

`func (o *SettingsFlow) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *SettingsFlow) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *SettingsFlow) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.


### GetReturnTo

`func (o *SettingsFlow) GetReturnTo() string`

GetReturnTo returns the ReturnTo field if non-nil, zero value otherwise.

### GetReturnToOk

`func (o *SettingsFlow) GetReturnToOk() (*string, bool)`

GetReturnToOk returns a tuple with the ReturnTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTo

`func (o *SettingsFlow) SetReturnTo(v string)`

SetReturnTo sets ReturnTo field to given value.

### HasReturnTo

`func (o *SettingsFlow) HasReturnTo() bool`

HasReturnTo returns a boolean if a field has been set.

### GetState

`func (o *SettingsFlow) GetState() SettingsState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SettingsFlow) GetStateOk() (*SettingsState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SettingsFlow) SetState(v SettingsState)`

SetState sets State field to given value.


### GetType

`func (o *SettingsFlow) GetType() FlowType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SettingsFlow) GetTypeOk() (*FlowType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SettingsFlow) SetType(v FlowType)`

SetType sets Type field to given value.


### GetUi

`func (o *SettingsFlow) GetUi() UiContainer`

GetUi returns the Ui field if non-nil, zero value otherwise.

### GetUiOk

`func (o *SettingsFlow) GetUiOk() (*UiContainer, bool)`

GetUiOk returns a tuple with the Ui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUi

`func (o *SettingsFlow) SetUi(v UiContainer)`

SetUi sets Ui field to given value.


### GetUpdatedAt

`func (o *SettingsFlow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SettingsFlow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SettingsFlow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SettingsFlow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


