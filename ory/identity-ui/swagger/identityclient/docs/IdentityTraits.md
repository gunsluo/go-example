# IdentityTraits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The identity&#39;s email  It&#39;s optional, have at least one email and mobile | [optional] 
**Mobile** | Pointer to **string** | The identity&#39;s mobile  It&#39;s optional, have at least one email and mobile | [optional] 

## Methods

### NewIdentityTraits

`func NewIdentityTraits() *IdentityTraits`

NewIdentityTraits instantiates a new IdentityTraits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityTraitsWithDefaults

`func NewIdentityTraitsWithDefaults() *IdentityTraits`

NewIdentityTraitsWithDefaults instantiates a new IdentityTraits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *IdentityTraits) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentityTraits) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentityTraits) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IdentityTraits) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMobile

`func (o *IdentityTraits) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *IdentityTraits) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *IdentityTraits) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *IdentityTraits) HasMobile() bool`

HasMobile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

