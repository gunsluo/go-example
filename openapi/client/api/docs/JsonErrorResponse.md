# JsonErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code represents the error code. | [optional] 
**Detail** | Pointer to **string** | Detail contains further information on the nature of the error. | [optional] 
**Msg** | Pointer to **string** | Message is the error message. | [optional] 
**TraceId** | Pointer to **string** | TraceId is the identifier for a trace. It is globally unique. | [optional] 
**Type** | Pointer to **string** | Type A URI reference that identifies the error type. | [optional] 

## Methods

### NewJsonErrorResponse

`func NewJsonErrorResponse() *JsonErrorResponse`

NewJsonErrorResponse instantiates a new JsonErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonErrorResponseWithDefaults

`func NewJsonErrorResponseWithDefaults() *JsonErrorResponse`

NewJsonErrorResponseWithDefaults instantiates a new JsonErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *JsonErrorResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *JsonErrorResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *JsonErrorResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *JsonErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *JsonErrorResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *JsonErrorResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *JsonErrorResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *JsonErrorResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetMsg

`func (o *JsonErrorResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *JsonErrorResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *JsonErrorResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *JsonErrorResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetTraceId

`func (o *JsonErrorResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *JsonErrorResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *JsonErrorResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *JsonErrorResponse) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *JsonErrorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JsonErrorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JsonErrorResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JsonErrorResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


