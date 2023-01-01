# IdentityCredentialIdentifierPasswordProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | **bool** | Primary  Is identifier&#39;s primary identity | 
**Type** | Pointer to **string** |  | [optional] 
**Verified** | **bool** | Verified  Is identifier been verified | 

## Methods

### NewIdentityCredentialIdentifierPasswordProperties

`func NewIdentityCredentialIdentifierPasswordProperties(primary bool, verified bool, ) *IdentityCredentialIdentifierPasswordProperties`

NewIdentityCredentialIdentifierPasswordProperties instantiates a new IdentityCredentialIdentifierPasswordProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityCredentialIdentifierPasswordPropertiesWithDefaults

`func NewIdentityCredentialIdentifierPasswordPropertiesWithDefaults() *IdentityCredentialIdentifierPasswordProperties`

NewIdentityCredentialIdentifierPasswordPropertiesWithDefaults instantiates a new IdentityCredentialIdentifierPasswordProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *IdentityCredentialIdentifierPasswordProperties) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *IdentityCredentialIdentifierPasswordProperties) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *IdentityCredentialIdentifierPasswordProperties) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


### GetType

`func (o *IdentityCredentialIdentifierPasswordProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityCredentialIdentifierPasswordProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityCredentialIdentifierPasswordProperties) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityCredentialIdentifierPasswordProperties) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVerified

`func (o *IdentityCredentialIdentifierPasswordProperties) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *IdentityCredentialIdentifierPasswordProperties) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *IdentityCredentialIdentifierPasswordProperties) SetVerified(v bool)`

SetVerified sets Verified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


