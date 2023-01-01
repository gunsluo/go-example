# \RegistrationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRegistrationFlowRequest**](RegistrationApi.md#GetRegistrationFlowRequest) | **Get** /self-service/registration/flows | # Get Registration Flow
[**InitBrowserRegistrationFlowRequest**](RegistrationApi.md#InitBrowserRegistrationFlowRequest) | **Get** /self-service/registration/browser | # Initialize Registration Flow for Browsers
[**SubmitRegistrationFlowRequest**](RegistrationApi.md#SubmitRegistrationFlowRequest) | **Post** /self-service/registration | # Submit a Registration Flow



## GetRegistrationFlowRequest

> RegistrationFlow GetRegistrationFlowRequest(ctx).Id(id).Cookie(cookie).Execute()

# Get Registration Flow



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
    id := "id_example" // string | The Registration Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/registration?flow=abcde`).
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistrationApi.GetRegistrationFlowRequest(context.Background()).Id(id).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistrationApi.GetRegistrationFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistrationFlowRequest`: RegistrationFlow
    fmt.Fprintf(os.Stdout, "Response from `RegistrationApi.GetRegistrationFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistrationFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The Registration Flow ID  The value for this parameter comes from &#x60;flow&#x60; URL Query parameter sent to your application (e.g. &#x60;/registration?flow&#x3D;abcde&#x60;). | 
 **cookie** | **string** | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. | 

### Return type

[**RegistrationFlow**](RegistrationFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitBrowserRegistrationFlowRequest

> RegistrationFlow InitBrowserRegistrationFlowRequest(ctx).ReturnTo(returnTo).LoginChallenge(loginChallenge).Execute()

# Initialize Registration Flow for Browsers



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
    loginChallenge := "loginChallenge_example" // string | OAuth 2.0 Login Challenge.  If set will cooperate with OAuth2 and OpenID to act as an OAuth2 server / OpenID Provider.  The value for this parameter comes from `login_challenge` URL Query parameter sent to your application (e.g. `/registration?login_challenge=abcde`).  This feature is compatible with Identity when not running on the Network. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistrationApi.InitBrowserRegistrationFlowRequest(context.Background()).ReturnTo(returnTo).LoginChallenge(loginChallenge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistrationApi.InitBrowserRegistrationFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitBrowserRegistrationFlowRequest`: RegistrationFlow
    fmt.Fprintf(os.Stdout, "Response from `RegistrationApi.InitBrowserRegistrationFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitBrowserRegistrationFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnTo** | **string** | The URL to return the browser to after the flow was completed. | 
 **loginChallenge** | **string** | OAuth 2.0 Login Challenge.  If set will cooperate with OAuth2 and OpenID to act as an OAuth2 server / OpenID Provider.  The value for this parameter comes from &#x60;login_challenge&#x60; URL Query parameter sent to your application (e.g. &#x60;/registration?login_challenge&#x3D;abcde&#x60;).  This feature is compatible with Identity when not running on the Network. | 

### Return type

[**RegistrationFlow**](RegistrationFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitRegistrationFlowRequest

> SubmitRegistrationFlowResponse SubmitRegistrationFlowRequest(ctx).Flow(flow).SubmitRegistrationFlowBody(submitRegistrationFlowBody).Cookie(cookie).Execute()

# Submit a Registration Flow



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
    flow := "flow_example" // string | The Registration Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/registration?flow=abcde`).
    submitRegistrationFlowBody := openapiclient.submitRegistrationFlowBody{SubmitRegistrationFlowBodyWithPassword: openapiclient.NewSubmitRegistrationFlowBodyWithPassword("ConfirmPassword_example", "Method_example", "Password_example", *openapiclient.NewRegistrationTraits("Type_example"))} // SubmitRegistrationFlowBody | 
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistrationApi.SubmitRegistrationFlowRequest(context.Background()).Flow(flow).SubmitRegistrationFlowBody(submitRegistrationFlowBody).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistrationApi.SubmitRegistrationFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitRegistrationFlowRequest`: SubmitRegistrationFlowResponse
    fmt.Fprintf(os.Stdout, "Response from `RegistrationApi.SubmitRegistrationFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRegistrationFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | The Registration Flow ID  The value for this parameter comes from &#x60;flow&#x60; URL Query parameter sent to your application (e.g. &#x60;/registration?flow&#x3D;abcde&#x60;). | 
 **submitRegistrationFlowBody** | [**SubmitRegistrationFlowBody**](SubmitRegistrationFlowBody.md) |  | 
 **cookie** | **string** | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. | 

### Return type

[**SubmitRegistrationFlowResponse**](SubmitRegistrationFlowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

