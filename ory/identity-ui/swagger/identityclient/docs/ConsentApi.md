# \ConsentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConsentFlowRequest**](ConsentApi.md#GetConsentFlowRequest) | **Get** /self-service/consent/flows | # Get Consent Flow
[**InitBrowserConsentFlowRequest**](ConsentApi.md#InitBrowserConsentFlowRequest) | **Get** /self-service/consent/browser | # Initialize Consent Flow for Browsers
[**SubmitConsentFlowRequest**](ConsentApi.md#SubmitConsentFlowRequest) | **Post** /self-service/consent | # Submit a Consent Flow



## GetConsentFlowRequest

> ConsentFlow GetConsentFlowRequest(ctx).Id(id).Cookie(cookie).Execute()

# Get Consent Flow



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
    id := "id_example" // string | The Consent Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/consent?flow=abcde`).
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConsentApi.GetConsentFlowRequest(context.Background()).Id(id).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsentApi.GetConsentFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentFlowRequest`: ConsentFlow
    fmt.Fprintf(os.Stdout, "Response from `ConsentApi.GetConsentFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The Consent Flow ID  The value for this parameter comes from &#x60;flow&#x60; URL Query parameter sent to your application (e.g. &#x60;/consent?flow&#x3D;abcde&#x60;). | 
 **cookie** | **string** | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. | 

### Return type

[**ConsentFlow**](ConsentFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitBrowserConsentFlowRequest

> ConsentFlow InitBrowserConsentFlowRequest(ctx).ConsentChallenge(consentChallenge).ReturnTo(returnTo).Cookie(cookie).Execute()

# Initialize Consent Flow for Browsers



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
    consentChallenge := "consentChallenge_example" // string | An optional Hydra consent challenge. If present, Kratos will cooperate with Ory Hydra to act as an OAuth2 identity provider.  The value for this parameter comes from `consent_challenge` URL Query parameter sent to your application (e.g. `/consent?consent_challenge=abcde`).
    returnTo := "returnTo_example" // string | The URL to return the browser to after the flow was completed. (optional)
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConsentApi.InitBrowserConsentFlowRequest(context.Background()).ConsentChallenge(consentChallenge).ReturnTo(returnTo).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsentApi.InitBrowserConsentFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitBrowserConsentFlowRequest`: ConsentFlow
    fmt.Fprintf(os.Stdout, "Response from `ConsentApi.InitBrowserConsentFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitBrowserConsentFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consentChallenge** | **string** | An optional Hydra consent challenge. If present, Kratos will cooperate with Ory Hydra to act as an OAuth2 identity provider.  The value for this parameter comes from &#x60;consent_challenge&#x60; URL Query parameter sent to your application (e.g. &#x60;/consent?consent_challenge&#x3D;abcde&#x60;). | 
 **returnTo** | **string** | The URL to return the browser to after the flow was completed. | 
 **cookie** | **string** | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. | 

### Return type

[**ConsentFlow**](ConsentFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitConsentFlowRequest

> ConsentFlow SubmitConsentFlowRequest(ctx).Flow(flow).SubmitConsentFlowBody(submitConsentFlowBody).XSessionToken(xSessionToken).Cookie(cookie).Execute()

# Submit a Consent Flow



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
    flow := "flow_example" // string | The Consent Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/login?flow=abcde`).
    submitConsentFlowBody := openapiclient.submitConsentFlowBody{SubmitConsentFlowBodyWithDefault: openapiclient.NewSubmitConsentFlowBodyWithDefault("Action_example", "Method_example", false, []string{"Scope_example"})} // SubmitConsentFlowBody | 
    xSessionToken := "xSessionToken_example" // string | The Session Token of the Identity performing the consent flow. (optional)
    cookie := "cookie_example" // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConsentApi.SubmitConsentFlowRequest(context.Background()).Flow(flow).SubmitConsentFlowBody(submitConsentFlowBody).XSessionToken(xSessionToken).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsentApi.SubmitConsentFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitConsentFlowRequest`: ConsentFlow
    fmt.Fprintf(os.Stdout, "Response from `ConsentApi.SubmitConsentFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitConsentFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | The Consent Flow ID  The value for this parameter comes from &#x60;flow&#x60; URL Query parameter sent to your application (e.g. &#x60;/login?flow&#x3D;abcde&#x60;). | 
 **submitConsentFlowBody** | [**SubmitConsentFlowBody**](SubmitConsentFlowBody.md) |  | 
 **xSessionToken** | **string** | The Session Token of the Identity performing the consent flow. | 
 **cookie** | **string** | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. | 

### Return type

[**ConsentFlow**](ConsentFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

