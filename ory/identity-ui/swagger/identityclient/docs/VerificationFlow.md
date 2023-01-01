# VerificationFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is a helper struct field for. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | ExpiresAt is the time (UTC) when the request expires. If the user still wishes to verify the address, a new request has to be initiated. | [optional] 
**Id** | **string** | ID represents the request&#39;s unique ID. When performing the verification flow, this represents the id in the verify ui&#39;s query parameter: http://&lt;selfservice.flows.verification.ui_url&gt;?request&#x3D;&lt;id&gt;  type: string | 
**IssuedAt** | Pointer to **time.Time** | IssuedAt is the time (UTC) when the request occurred. | [optional] 
**Method** | Pointer to [**VerificationMethod**](VerificationMethod.md) |  | [optional] 
**RequestUrl** | Pointer to **string** | RequestURL is the initial URL that was requested from Identity. It can be used to forward information contained in the URL&#39;s path or query for example. | [optional] 
**ReturnTo** | Pointer to **string** | ReturnTo contains the requested return_to URL. | [optional] 
**State** | [**VerificationState**](VerificationState.md) |  | 
**Type** | [**FlowType**](FlowType.md) |  | 
**Ui** | [**UiContainer**](UiContainer.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is a helper struct field for. | [optional] 

## Methods

### NewVerificationFlow

`func NewVerificationFlow(id string, state VerificationState, type_ FlowType, ui UiContainer, ) *VerificationFlow`

NewVerificationFlow instantiates a new VerificationFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationFlowWithDefaults

`func NewVerificationFlowWithDefaults() *VerificationFlow`

NewVerificationFlowWithDefaults instantiates a new VerificationFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *VerificationFlow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VerificationFlow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VerificationFlow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VerificationFlow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *VerificationFlow) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *VerificationFlow) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *VerificationFlow) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *VerificationFlow) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *VerificationFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerificationFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerificationFlow) SetId(v string)`

SetId sets Id field to given value.


### GetIssuedAt

`func (o *VerificationFlow) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *VerificationFlow) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *VerificationFlow) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *VerificationFlow) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### GetMethod

`func (o *VerificationFlow) GetMethod() VerificationMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *VerificationFlow) GetMethodOk() (*VerificationMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *VerificationFlow) SetMethod(v VerificationMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *VerificationFlow) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRequestUrl

`func (o *VerificationFlow) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *VerificationFlow) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *VerificationFlow) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.

### HasRequestUrl

`func (o *VerificationFlow) HasRequestUrl() bool`

HasRequestUrl returns a boolean if a field has been set.

### GetReturnTo

`func (o *VerificationFlow) GetReturnTo() string`

GetReturnTo returns the ReturnTo field if non-nil, zero value otherwise.

### GetReturnToOk

`func (o *VerificationFlow) GetReturnToOk() (*string, bool)`

GetReturnToOk returns a tuple with the ReturnTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTo

`func (o *VerificationFlow) SetReturnTo(v string)`

SetReturnTo sets ReturnTo field to given value.

### HasReturnTo

`func (o *VerificationFlow) HasReturnTo() bool`

HasReturnTo returns a boolean if a field has been set.

### GetState

`func (o *VerificationFlow) GetState() VerificationState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VerificationFlow) GetStateOk() (*VerificationState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VerificationFlow) SetState(v VerificationState)`

SetState sets State field to given value.


### GetType

`func (o *VerificationFlow) GetType() FlowType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VerificationFlow) GetTypeOk() (*FlowType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VerificationFlow) SetType(v FlowType)`

SetType sets Type field to given value.


### GetUi

`func (o *VerificationFlow) GetUi() UiContainer`

GetUi returns the Ui field if non-nil, zero value otherwise.

### GetUiOk

`func (o *VerificationFlow) GetUiOk() (*UiContainer, bool)`

GetUiOk returns a tuple with the Ui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUi

`func (o *VerificationFlow) SetUi(v UiContainer)`

SetUi sets Ui field to given value.


### GetUpdatedAt

`func (o *VerificationFlow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VerificationFlow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VerificationFlow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VerificationFlow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


