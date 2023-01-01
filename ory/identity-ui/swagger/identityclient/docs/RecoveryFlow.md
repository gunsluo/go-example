# RecoveryFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is a helper struct field for gobuffalo.pop. | [optional] 
**ExpiresAt** | **time.Time** | ExpiresAt is the time (UTC) when the request expires. If the user still wishes to update the setting, a new request has to be initiated. | 
**Id** | **string** | ID represents the request&#39;s unique ID. When performing the recovery flow, this represents the id in the recovery ui&#39;s query parameter: http://&lt;selfservice.flows.recovery.ui_url&gt;?request&#x3D;&lt;id&gt; | 
**IssuedAt** | **time.Time** | IssuedAt is the time (UTC) when the request occurred. | 
**Method** | Pointer to [**RecoveryMethod**](RecoveryMethod.md) |  | [optional] 
**RequestUrl** | **string** | RequestURL is the initial URL that was requested from Identity. It can be used to forward information contained in the URL&#39;s path or query for example. | 
**ReturnTo** | Pointer to **string** | ReturnTo contains the requested return_to URL. | [optional] 
**State** | [**RecoveryState**](RecoveryState.md) |  | 
**Type** | [**FlowType**](FlowType.md) |  | 
**Ui** | [**UiContainer**](UiContainer.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is a helper struct field for gobuffalo.pop. | [optional] 

## Methods

### NewRecoveryFlow

`func NewRecoveryFlow(expiresAt time.Time, id string, issuedAt time.Time, requestUrl string, state RecoveryState, type_ FlowType, ui UiContainer, ) *RecoveryFlow`

NewRecoveryFlow instantiates a new RecoveryFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryFlowWithDefaults

`func NewRecoveryFlowWithDefaults() *RecoveryFlow`

NewRecoveryFlowWithDefaults instantiates a new RecoveryFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RecoveryFlow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RecoveryFlow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RecoveryFlow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RecoveryFlow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *RecoveryFlow) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *RecoveryFlow) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *RecoveryFlow) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetId

`func (o *RecoveryFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecoveryFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecoveryFlow) SetId(v string)`

SetId sets Id field to given value.


### GetIssuedAt

`func (o *RecoveryFlow) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *RecoveryFlow) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *RecoveryFlow) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.


### GetMethod

`func (o *RecoveryFlow) GetMethod() RecoveryMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RecoveryFlow) GetMethodOk() (*RecoveryMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RecoveryFlow) SetMethod(v RecoveryMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RecoveryFlow) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRequestUrl

`func (o *RecoveryFlow) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *RecoveryFlow) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *RecoveryFlow) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.


### GetReturnTo

`func (o *RecoveryFlow) GetReturnTo() string`

GetReturnTo returns the ReturnTo field if non-nil, zero value otherwise.

### GetReturnToOk

`func (o *RecoveryFlow) GetReturnToOk() (*string, bool)`

GetReturnToOk returns a tuple with the ReturnTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTo

`func (o *RecoveryFlow) SetReturnTo(v string)`

SetReturnTo sets ReturnTo field to given value.

### HasReturnTo

`func (o *RecoveryFlow) HasReturnTo() bool`

HasReturnTo returns a boolean if a field has been set.

### GetState

`func (o *RecoveryFlow) GetState() RecoveryState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RecoveryFlow) GetStateOk() (*RecoveryState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RecoveryFlow) SetState(v RecoveryState)`

SetState sets State field to given value.


### GetType

`func (o *RecoveryFlow) GetType() FlowType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecoveryFlow) GetTypeOk() (*FlowType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecoveryFlow) SetType(v FlowType)`

SetType sets Type field to given value.


### GetUi

`func (o *RecoveryFlow) GetUi() UiContainer`

GetUi returns the Ui field if non-nil, zero value otherwise.

### GetUiOk

`func (o *RecoveryFlow) GetUiOk() (*UiContainer, bool)`

GetUiOk returns a tuple with the Ui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUi

`func (o *RecoveryFlow) SetUi(v UiContainer)`

SetUi sets Ui field to given value.


### GetUpdatedAt

`func (o *RecoveryFlow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RecoveryFlow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RecoveryFlow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RecoveryFlow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


