# BrowserLocationChangeRequiredResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**Code**](Code.md) |  | [optional] 
**Detail** | Pointer to **string** | Detail contains further information on the nature of the error. | [optional] 
**Msg** | Pointer to **string** | Message is the error message. | [optional] 
**RedirectBrowserTo** | Pointer to **string** | Since when the flow has expired | [optional] 
**TraceId** | Pointer to **string** | TraceId is the identifier for a trace. It is globally unique. | [optional] 
**Type** | Pointer to **string** | Type A URI reference that identifies the error type. | [optional] 

## Methods

### NewBrowserLocationChangeRequiredResponse

`func NewBrowserLocationChangeRequiredResponse() *BrowserLocationChangeRequiredResponse`

NewBrowserLocationChangeRequiredResponse instantiates a new BrowserLocationChangeRequiredResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserLocationChangeRequiredResponseWithDefaults

`func NewBrowserLocationChangeRequiredResponseWithDefaults() *BrowserLocationChangeRequiredResponse`

NewBrowserLocationChangeRequiredResponseWithDefaults instantiates a new BrowserLocationChangeRequiredResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *BrowserLocationChangeRequiredResponse) GetCode() Code`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BrowserLocationChangeRequiredResponse) GetCodeOk() (*Code, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BrowserLocationChangeRequiredResponse) SetCode(v Code)`

SetCode sets Code field to given value.

### HasCode

`func (o *BrowserLocationChangeRequiredResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *BrowserLocationChangeRequiredResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *BrowserLocationChangeRequiredResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *BrowserLocationChangeRequiredResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *BrowserLocationChangeRequiredResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetMsg

`func (o *BrowserLocationChangeRequiredResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *BrowserLocationChangeRequiredResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *BrowserLocationChangeRequiredResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *BrowserLocationChangeRequiredResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetRedirectBrowserTo

`func (o *BrowserLocationChangeRequiredResponse) GetRedirectBrowserTo() string`

GetRedirectBrowserTo returns the RedirectBrowserTo field if non-nil, zero value otherwise.

### GetRedirectBrowserToOk

`func (o *BrowserLocationChangeRequiredResponse) GetRedirectBrowserToOk() (*string, bool)`

GetRedirectBrowserToOk returns a tuple with the RedirectBrowserTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectBrowserTo

`func (o *BrowserLocationChangeRequiredResponse) SetRedirectBrowserTo(v string)`

SetRedirectBrowserTo sets RedirectBrowserTo field to given value.

### HasRedirectBrowserTo

`func (o *BrowserLocationChangeRequiredResponse) HasRedirectBrowserTo() bool`

HasRedirectBrowserTo returns a boolean if a field has been set.

### GetTraceId

`func (o *BrowserLocationChangeRequiredResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *BrowserLocationChangeRequiredResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *BrowserLocationChangeRequiredResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *BrowserLocationChangeRequiredResponse) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *BrowserLocationChangeRequiredResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BrowserLocationChangeRequiredResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BrowserLocationChangeRequiredResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BrowserLocationChangeRequiredResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


