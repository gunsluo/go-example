# \LogoutApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InitBrowserLogoutFlowRequest**](LogoutApi.md#InitBrowserLogoutFlowRequest) | **Get** /self-service/logout/browser | # Create a Logout URL for Browsers
[**SubmitLogoutFlowRequest**](LogoutApi.md#SubmitLogoutFlowRequest) | **Get** /self-service/logout | # Complete Self-Service Logout



## InitBrowserLogoutFlowRequest

> InitBrowserLogoutFlowResponse InitBrowserLogoutFlowRequest(ctx).LogoutChallenge(logoutChallenge).Cookie(cookie).Execute()

# Create a Logout URL for Browsers



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
    logoutChallenge := "logoutChallenge_example" // string | An optional Hydra logout challenge. If present, Kratos will cooperate with Ory Hydra to act as an OAuth2 identity provider.  The value for this parameter comes from `logout_challenge` URL Query parameter sent to your application (e.g. `/logout?logout_challenge=abcde`). (optional)
    cookie := "cookie_example" // string | HTTP Cookies  If you call this endpoint from a backend, please include the original Cookie header in the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogoutApi.InitBrowserLogoutFlowRequest(context.Background()).LogoutChallenge(logoutChallenge).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogoutApi.InitBrowserLogoutFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitBrowserLogoutFlowRequest`: InitBrowserLogoutFlowResponse
    fmt.Fprintf(os.Stdout, "Response from `LogoutApi.InitBrowserLogoutFlowRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitBrowserLogoutFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logoutChallenge** | **string** | An optional Hydra logout challenge. If present, Kratos will cooperate with Ory Hydra to act as an OAuth2 identity provider.  The value for this parameter comes from &#x60;logout_challenge&#x60; URL Query parameter sent to your application (e.g. &#x60;/logout?logout_challenge&#x3D;abcde&#x60;). | 
 **cookie** | **string** | HTTP Cookies  If you call this endpoint from a backend, please include the original Cookie header in the request. | 

### Return type

[**InitBrowserLogoutFlowResponse**](InitBrowserLogoutFlowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitLogoutFlowRequest

> SubmitLogoutFlowRequest(ctx).Token(token).ReturnTo(returnTo).LogoutChallenge(logoutChallenge).Execute()

# Complete Self-Service Logout



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
    token := "token_example" // string | A Valid Logout Token  If you do not have a logout token because you only have a session cookie, call `/self-service/logout/urls` to generate a URL for this endpoint. (optional)
    returnTo := "returnTo_example" // string | The URL to return to after the logout was completed. (optional)
    logoutChallenge := "logoutChallenge_example" // string | An optional Hydra logout challenge. If present, Kratos will cooperate with Ory Hydra to act as an OAuth2 identity provider.  The value for this parameter comes from `logout_challenge` URL Query parameter sent to your application (e.g. `/logout?logout_challenge=abcde`). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogoutApi.SubmitLogoutFlowRequest(context.Background()).Token(token).ReturnTo(returnTo).LogoutChallenge(logoutChallenge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogoutApi.SubmitLogoutFlowRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitLogoutFlowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | A Valid Logout Token  If you do not have a logout token because you only have a session cookie, call &#x60;/self-service/logout/urls&#x60; to generate a URL for this endpoint. | 
 **returnTo** | **string** | The URL to return to after the logout was completed. | 
 **logoutChallenge** | **string** | An optional Hydra logout challenge. If present, Kratos will cooperate with Ory Hydra to act as an OAuth2 identity provider.  The value for this parameter comes from &#x60;logout_challenge&#x60; URL Query parameter sent to your application (e.g. &#x60;/logout?logout_challenge&#x3D;abcde&#x60;). | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

