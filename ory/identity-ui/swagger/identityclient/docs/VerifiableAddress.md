# VerifiableAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID | 
**Identifier** | **string** | The identifier  example foo@user.com | 
**Primary** | **bool** | Primary  Is identifier&#39;s primary identity | 
**Verified** | **bool** | Verified  Is the identifier been verified | 

## Methods

### NewVerifiableAddress

`func NewVerifiableAddress(id string, identifier string, primary bool, verified bool, ) *VerifiableAddress`

NewVerifiableAddress instantiates a new VerifiableAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifiableAddressWithDefaults

`func NewVerifiableAddressWithDefaults() *VerifiableAddress`

NewVerifiableAddressWithDefaults instantiates a new VerifiableAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VerifiableAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerifiableAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerifiableAddress) SetId(v string)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *VerifiableAddress) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VerifiableAddress) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VerifiableAddress) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetPrimary

`func (o *VerifiableAddress) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *VerifiableAddress) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *VerifiableAddress) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


### GetVerified

`func (o *VerifiableAddress) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *VerifiableAddress) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *VerifiableAddress) SetVerified(v bool)`

SetVerified sets Verified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


