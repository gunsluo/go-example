# \OrganizationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMembersRequest**](OrganizationApi.md#AddMembersRequest) | **Post** /org/add-members | Add members to the organization



## AddMembersRequest

> AddMembersResponse AddMembersRequest(ctx).OrganizationId(organizationId).OrganizatoinMember(organizatoinMember).Execute()

Add members to the organization



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
    organizationId := int64(789) // int64 | 
    organizatoinMember := []openapiclient.OrganizatoinMember{*openapiclient.NewOrganizatoinMember("Mail_example", "Name_example")} // []OrganizatoinMember |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApi.AddMembersRequest(context.Background()).OrganizationId(organizationId).OrganizatoinMember(organizatoinMember).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.AddMembersRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMembersRequest`: AddMembersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.AddMembersRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMembersRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int64** |  | 
 **organizatoinMember** | [**[]OrganizatoinMember**](OrganizatoinMember.md) |  | 

### Return type

[**AddMembersResponse**](AddMembersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

