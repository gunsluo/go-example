# \LoginApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLoginFlowRequest**](LoginApi.md#GetLoginFlowRequest) | **Get** /self-service/login/flows | # Get Login Flow
[**InitBrowserLoginFlowRequest**](LoginApi.md#InitBrowserLoginFlowRequest) | **Get** /self-service/login/browser | # Initialize Login Flow for Browsers
[**SubmitLoginFlowRequest**](LoginApi.md#SubmitLoginFlowRequest) | **Post** /self-service/login | # Submit a Login Flow



## GetLoginFlowRequest

> LoginFlow GetLoginFlowRequest(ctx).Id(id).Cookie(cookie).Execute()

# Get Login Flow



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
    id := "id_example" // string | The Login Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/login?flow=abcde`).
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.GetLoginFlowRequest(context.Background()).Id(id).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.GetLoginFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoginFlowRequest`: LoginFlow
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.GetLoginFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The Login Flow ID  The value for this parameter comes from &#x60;flow&#x60; URL Query parameter sent to your application (e.g. &#x60;/login?flow&#x3D;abcde&#x60;). | 
 **cookie** | **string** | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. | 

### Return type

[**LoginFlow**](LoginFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitBrowserLoginFlowRequest

> LoginFlow InitBrowserLoginFlowRequest(ctx).Refresh(refresh).Aal(aal).ReturnTo(returnTo).Cookie(cookie).LoginChallenge(loginChallenge).Execute()

# Initialize Login Flow for Browsers



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
    refresh := true // bool | Refresh a login session  If set to true, this will refresh an existing login session by asking the user to sign in again. This will reset the authenticated_at time of the session. (optional)
    aal := "aal_example" // string | Request a Specific AuthenticationMethod Assurance Level  Use this parameter to upgrade an existing session's authenticator assurance level (AAL). This allows you to ask for multi-factor authentication. When an identity sign in using e.g. username+password, the AAL is 1. If you wish to \"upgrade\" the session's security by asking the user to perform TOTP / WebAuth/ ... you would set this to \"aal2\". (optional)
    returnTo := "returnTo_example" // string | The URL to return the browser to after the flow was completed. (optional)
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)
    loginChallenge := "loginChallenge_example" // string | An optional Hydra login challenge. If present, Kratos will cooperate with Ory Hydra to act as an OAuth2 identity provider.  The value for this parameter comes from `login_challenge` URL Query parameter sent to your application (e.g. `/login?login_challenge=abcde`). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.InitBrowserLoginFlowRequest(context.Background()).Refresh(refresh).Aal(aal).ReturnTo(returnTo).Cookie(cookie).LoginChallenge(loginChallenge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.InitBrowserLoginFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitBrowserLoginFlowRequest`: LoginFlow
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.InitBrowserLoginFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitBrowserLoginFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refresh** | **bool** | Refresh a login session  If set to true, this will refresh an existing login session by asking the user to sign in again. This will reset the authenticated_at time of the session. | 
 **aal** | **string** | Request a Specific AuthenticationMethod Assurance Level  Use this parameter to upgrade an existing session&#39;s authenticator assurance level (AAL). This allows you to ask for multi-factor authentication. When an identity sign in using e.g. username+password, the AAL is 1. If you wish to \&quot;upgrade\&quot; the session&#39;s security by asking the user to perform TOTP / WebAuth/ ... you would set this to \&quot;aal2\&quot;. | 
 **returnTo** | **string** | The URL to return the browser to after the flow was completed. | 
 **cookie** | **string** | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. | 
 **loginChallenge** | **string** | An optional Hydra login challenge. If present, Kratos will cooperate with Ory Hydra to act as an OAuth2 identity provider.  The value for this parameter comes from &#x60;login_challenge&#x60; URL Query parameter sent to your application (e.g. &#x60;/login?login_challenge&#x3D;abcde&#x60;). | 

### Return type

[**LoginFlow**](LoginFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitLoginFlowRequest

> SubmitLoginFlowResponse SubmitLoginFlowRequest(ctx).Flow(flow).SubmitLoginFlowBody(submitLoginFlowBody).XSessionToken(xSessionToken).Cookie(cookie).Execute()

# Submit a Login Flow



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
    flow := "flow_example" // string | The Login Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/login?flow=abcde`).
    submitLoginFlowBody := openapiclient.submitLoginFlowBody{SubmitLoginFlowBodyWithOidc: openapiclient.NewSubmitLoginFlowBodyWithOidc("Method_example", "Provider_example")} // SubmitLoginFlowBody | 
    xSessionToken := "xSessionToken_example" // string | The Session Token of the Identity performing the settings flow. (optional)
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.SubmitLoginFlowRequest(context.Background()).Flow(flow).SubmitLoginFlowBody(submitLoginFlowBody).XSessionToken(xSessionToken).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.SubmitLoginFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitLoginFlowRequest`: SubmitLoginFlowResponse
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.SubmitLoginFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitLoginFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | The Login Flow ID  The value for this parameter comes from &#x60;flow&#x60; URL Query parameter sent to your application (e.g. &#x60;/login?flow&#x3D;abcde&#x60;). | 
 **submitLoginFlowBody** | [**SubmitLoginFlowBody**](SubmitLoginFlowBody.md) |  | 
 **xSessionToken** | **string** | The Session Token of the Identity performing the settings flow. | 
 **cookie** | **string** | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. | 

### Return type

[**SubmitLoginFlowResponse**](SubmitLoginFlowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

