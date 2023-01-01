# \SettingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingsFlowRequest**](SettingsApi.md#GetSettingsFlowRequest) | **Get** /self-service/settings/flows | # Get Settings Flow
[**InitBrowserSettingsFlowRequest**](SettingsApi.md#InitBrowserSettingsFlowRequest) | **Get** /self-service/settings/browser | # Initialize Settings Flow for Browsers
[**SubmitSettingsFlowRequest**](SettingsApi.md#SubmitSettingsFlowRequest) | **Post** /self-service/settings | # Complete Settings Flow



## GetSettingsFlowRequest

> SettingsFlow GetSettingsFlowRequest(ctx).Id(id).XSessionToken(xSessionToken).Cookie(cookie).Execute()

# Get Settings Flow



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
    id := "id_example" // string | ID is the Settings Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/settings?flow=abcde`).
    xSessionToken := "xSessionToken_example" // string | The Session Token  When using the SDK in an app without a browser, please include the session token here. (optional)
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.GetSettingsFlowRequest(context.Background()).Id(id).XSessionToken(xSessionToken).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetSettingsFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSettingsFlowRequest`: SettingsFlow
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetSettingsFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | ID is the Settings Flow ID  The value for this parameter comes from &#x60;flow&#x60; URL Query parameter sent to your application (e.g. &#x60;/settings?flow&#x3D;abcde&#x60;). | 
 **xSessionToken** | **string** | The Session Token  When using the SDK in an app without a browser, please include the session token here. | 
 **cookie** | **string** | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. | 

### Return type

[**SettingsFlow**](SettingsFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitBrowserSettingsFlowRequest

> SettingsFlow InitBrowserSettingsFlowRequest(ctx).ReturnTo(returnTo).Execute()

# Initialize Settings Flow for Browsers



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
    resp, r, err := apiClient.SettingsApi.InitBrowserSettingsFlowRequest(context.Background()).ReturnTo(returnTo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.InitBrowserSettingsFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitBrowserSettingsFlowRequest`: SettingsFlow
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.InitBrowserSettingsFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitBrowserSettingsFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnTo** | **string** | The URL to return the browser to after the flow was completed. | 

### Return type

[**SettingsFlow**](SettingsFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitSettingsFlowRequest

> SettingsFlow SubmitSettingsFlowRequest(ctx).Flow(flow).SubmitSettingsFlowBody(submitSettingsFlowBody).XSessionToken(xSessionToken).Cookie(cookie).Execute()

# Complete Settings Flow



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
    flow := "flow_example" // string | The Settings Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/settings?flow=abcde`).
    submitSettingsFlowBody := openapiclient.submitSettingsFlowBody{SubmitSettingsFlowBodyWithPassword: openapiclient.NewSubmitSettingsFlowBodyWithPassword("Method_example", "Password_example")} // SubmitSettingsFlowBody | 
    xSessionToken := "xSessionToken_example" // string | The Session Token of the Identity performing the settings flow. (optional)
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.SubmitSettingsFlowRequest(context.Background()).Flow(flow).SubmitSettingsFlowBody(submitSettingsFlowBody).XSessionToken(xSessionToken).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.SubmitSettingsFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitSettingsFlowRequest`: SettingsFlow
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.SubmitSettingsFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitSettingsFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | The Settings Flow ID  The value for this parameter comes from &#x60;flow&#x60; URL Query parameter sent to your application (e.g. &#x60;/settings?flow&#x3D;abcde&#x60;). | 
 **submitSettingsFlowBody** | [**SubmitSettingsFlowBody**](SubmitSettingsFlowBody.md) |  | 
 **xSessionToken** | **string** | The Session Token of the Identity performing the settings flow. | 
 **cookie** | **string** | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. | 

### Return type

[**SettingsFlow**](SettingsFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

