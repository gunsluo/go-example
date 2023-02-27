# \VerificationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVerificationFlowRequest**](VerificationApi.md#GetVerificationFlowRequest) | **Get** /self-service/verification/flows | # Get Verification Flow
[**InitBrowserVerificationFlowRequest**](VerificationApi.md#InitBrowserVerificationFlowRequest) | **Get** /self-service/verification/browser | # Initialize Verification Flow for Browser Clients
[**SubmitVerificationFlowRequest**](VerificationApi.md#SubmitVerificationFlowRequest) | **Post** /self-service/verification | # Complete Verification Flow



## GetVerificationFlowRequest

> VerificationFlow GetVerificationFlowRequest(ctx).Id(id).Cookie(cookie).Execute()

# Get Verification Flow



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
    id := "id_example" // string | The Flow ID  The value for this parameter comes from `request` URL Query parameter sent to your application (e.g. `/verification?flow=abcde`).
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK on the server side you must include the HTTP Cookie Header originally sent to your HTTP handler here. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VerificationApi.GetVerificationFlowRequest(context.Background()).Id(id).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerificationApi.GetVerificationFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVerificationFlowRequest`: VerificationFlow
    fmt.Fprintf(os.Stdout, "Response from `VerificationApi.GetVerificationFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVerificationFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The Flow ID  The value for this parameter comes from &#x60;request&#x60; URL Query parameter sent to your application (e.g. &#x60;/verification?flow&#x3D;abcde&#x60;). | 
 **cookie** | **string** | HTTP Cookies  When using the SDK on the server side you must include the HTTP Cookie Header originally sent to your HTTP handler here. | 

### Return type

[**VerificationFlow**](VerificationFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitBrowserVerificationFlowRequest

> VerificationFlow InitBrowserVerificationFlowRequest(ctx).ReturnTo(returnTo).LoginChallenge(loginChallenge).Execute()

# Initialize Verification Flow for Browser Clients



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
    loginChallenge := "loginChallenge_example" // string | OAuth 2.0 Login Challenge.  If set will cooperate with OAuth2 and OpenID to act as an OAuth2 server / OpenID Provider.  The value for this parameter comes from `login_challenge` URL Query parameter sent to your application (e.g. `/verification?login_challenge=abcde`).  This feature is compatible with Identity when not running on the Network. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VerificationApi.InitBrowserVerificationFlowRequest(context.Background()).ReturnTo(returnTo).LoginChallenge(loginChallenge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerificationApi.InitBrowserVerificationFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitBrowserVerificationFlowRequest`: VerificationFlow
    fmt.Fprintf(os.Stdout, "Response from `VerificationApi.InitBrowserVerificationFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitBrowserVerificationFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnTo** | **string** | The URL to return the browser to after the flow was completed. | 
 **loginChallenge** | **string** | OAuth 2.0 Login Challenge.  If set will cooperate with OAuth2 and OpenID to act as an OAuth2 server / OpenID Provider.  The value for this parameter comes from &#x60;login_challenge&#x60; URL Query parameter sent to your application (e.g. &#x60;/verification?login_challenge&#x3D;abcde&#x60;).  This feature is compatible with Identity when not running on the Network. | 

### Return type

[**VerificationFlow**](VerificationFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitVerificationFlowRequest

> VerificationFlow SubmitVerificationFlowRequest(ctx).Flow(flow).SubmitVerificationFlowBody(submitVerificationFlowBody).Token(token).Cookie(cookie).Execute()

# Complete Verification Flow



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
    flow := "flow_example" // string | The Verification Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/verification?flow=abcde`).
    submitVerificationFlowBody := openapiclient.submitVerificationFlowBody{SubmitVerificationFlowBodyWithCaptcha: openapiclient.NewSubmitVerificationFlowBodyWithCaptcha("Method_example", *openapiclient.NewVerificationTraits("Type_example"))} // SubmitVerificationFlowBody | 
    token := "token_example" // string | Verification Token  The verification token which completes the verification request. If the token is invalid (e.g. expired) an error will be shown to the end-user.  This parameter is usually set in a captcha and not used by any direct API call. (optional)
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VerificationApi.SubmitVerificationFlowRequest(context.Background()).Flow(flow).SubmitVerificationFlowBody(submitVerificationFlowBody).Token(token).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerificationApi.SubmitVerificationFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitVerificationFlowRequest`: VerificationFlow
    fmt.Fprintf(os.Stdout, "Response from `VerificationApi.SubmitVerificationFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitVerificationFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | The Verification Flow ID  The value for this parameter comes from &#x60;flow&#x60; URL Query parameter sent to your application (e.g. &#x60;/verification?flow&#x3D;abcde&#x60;). | 
 **submitVerificationFlowBody** | [**SubmitVerificationFlowBody**](SubmitVerificationFlowBody.md) |  | 
 **token** | **string** | Verification Token  The verification token which completes the verification request. If the token is invalid (e.g. expired) an error will be shown to the end-user.  This parameter is usually set in a captcha and not used by any direct API call. | 
 **cookie** | **string** | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. | 

### Return type

[**VerificationFlow**](VerificationFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

