# \RecoveryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRecoveryFlowRequest**](RecoveryApi.md#GetRecoveryFlowRequest) | **Get** /self-service/recovery/flows | # Get Recovery Flow
[**InitBrowserRecoveryFlowRequest**](RecoveryApi.md#InitBrowserRecoveryFlowRequest) | **Get** /self-service/recovery/browser | # Initialize Recovery Flow for Browsers
[**SubmitRecoveryFlowRequest**](RecoveryApi.md#SubmitRecoveryFlowRequest) | **Post** /self-service/recovery | # Complete Recovery Flow



## GetRecoveryFlowRequest

> RecoveryFlow GetRecoveryFlowRequest(ctx).Id(id).Cookie(cookie).Execute()

# Get Recovery Flow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The Flow ID  The value for this parameter comes from `request` URL Query parameter sent to your application (e.g. `/recovery?flow=abcde`).
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryApi.GetRecoveryFlowRequest(context.Background()).Id(id).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryApi.GetRecoveryFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecoveryFlowRequest`: RecoveryFlow
    fmt.Fprintf(os.Stdout, "Response from `RecoveryApi.GetRecoveryFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecoveryFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The Flow ID  The value for this parameter comes from &#x60;request&#x60; URL Query parameter sent to your application (e.g. &#x60;/recovery?flow&#x3D;abcde&#x60;). | 
 **cookie** | **string** | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. | 

### Return type

[**RecoveryFlow**](RecoveryFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitBrowserRecoveryFlowRequest

> RecoveryFlow InitBrowserRecoveryFlowRequest(ctx).ReturnTo(returnTo).Execute()

# Initialize Recovery Flow for Browsers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    returnTo := "returnTo_example" // string | The URL to return the browser to after the flow was completed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryApi.InitBrowserRecoveryFlowRequest(context.Background()).ReturnTo(returnTo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryApi.InitBrowserRecoveryFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitBrowserRecoveryFlowRequest`: RecoveryFlow
    fmt.Fprintf(os.Stdout, "Response from `RecoveryApi.InitBrowserRecoveryFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitBrowserRecoveryFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnTo** | **string** | The URL to return the browser to after the flow was completed. | 

### Return type

[**RecoveryFlow**](RecoveryFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitRecoveryFlowRequest

> RecoveryFlow SubmitRecoveryFlowRequest(ctx).Flow(flow).SubmitRecoveryFlowBody(submitRecoveryFlowBody).Token(token).Cookie(cookie).Execute()

# Complete Recovery Flow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    flow := "flow_example" // string | The Recovery Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/recovery?flow=abcde`).
    submitRecoveryFlowBody := openapiclient.submitRecoveryFlowBody{SubmitRecoveryFlowBodyWithCaptcha: openapiclient.NewSubmitRecoveryFlowBodyWithCaptcha("Method_example", *openapiclient.NewRecoveryTraits("Type_example"))} // SubmitRecoveryFlowBody | 
    token := "token_example" // string | Recovery Token  The recovery token which completes the recovery request. If the token is invalid (e.g. expired) an error will be shown to the end-user.  This parameter is usually set in a link and not used by any direct API call. (optional)
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryApi.SubmitRecoveryFlowRequest(context.Background()).Flow(flow).SubmitRecoveryFlowBody(submitRecoveryFlowBody).Token(token).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryApi.SubmitRecoveryFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitRecoveryFlowRequest`: RecoveryFlow
    fmt.Fprintf(os.Stdout, "Response from `RecoveryApi.SubmitRecoveryFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRecoveryFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | The Recovery Flow ID  The value for this parameter comes from &#x60;flow&#x60; URL Query parameter sent to your application (e.g. &#x60;/recovery?flow&#x3D;abcde&#x60;). | 
 **submitRecoveryFlowBody** | [**SubmitRecoveryFlowBody**](SubmitRecoveryFlowBody.md) |  | 
 **token** | **string** | Recovery Token  The recovery token which completes the recovery request. If the token is invalid (e.g. expired) an error will be shown to the end-user.  This parameter is usually set in a link and not used by any direct API call. | 
 **cookie** | **string** | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. | 

### Return type

[**RecoveryFlow**](RecoveryFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

