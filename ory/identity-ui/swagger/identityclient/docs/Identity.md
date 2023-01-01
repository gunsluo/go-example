# Identity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is a helper struct field for. | [optional] 
**Credentials** | Pointer to [**map[string]IdentityCredential**](IdentityCredential.md) | Credentials represents all credentials that can be used for authenticating this identity. | [optional] 
**Id** | **string** | ID is the identity&#39;s unique identifier.  The Identity ID can not be changed and can not be chosen. This ensures future compatibility and optimization for distributed stores such as CockroachDB. | 
**Locked** | Pointer to **bool** | Is the user been locked | [optional] 
**Source** | Pointer to **string** | the identity&#39;s source, e.g. local or github... | [optional] 
**State** | Pointer to [**IdentityState**](IdentityState.md) |  | [optional] 
**Subject** | Pointer to **string** | generate by id | [optional] 
**Traits** | [**IdentityTraits**](IdentityTraits.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is a helper struct field for. | [optional] 
**VerifiableAddresses** | Pointer to [**[]VerifiableAddress**](VerifiableAddress.md) | VerifiableAddresses contains all the addresses that can be verified by the user. | [optional] 

## Methods

### NewIdentity

`func NewIdentity(id string, traits IdentityTraits, ) *Identity`

NewIdentity instantiates a new Identity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityWithDefaults

`func NewIdentityWithDefaults() *Identity`

NewIdentityWithDefaults instantiates a new Identity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Identity) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Identity) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Identity) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Identity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentials

`func (o *Identity) GetCredentials() map[string]IdentityCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Identity) GetCredentialsOk() (*map[string]IdentityCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Identity) SetCredentials(v map[string]IdentityCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Identity) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetId

`func (o *Identity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Identity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Identity) SetId(v string)`

SetId sets Id field to given value.


### GetLocked

`func (o *Identity) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Identity) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Identity) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Identity) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetSource

`func (o *Identity) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Identity) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Identity) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Identity) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *Identity) GetState() IdentityState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Identity) GetStateOk() (*IdentityState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Identity) SetState(v IdentityState)`

SetState sets State field to given value.

### HasState

`func (o *Identity) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubject

`func (o *Identity) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Identity) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Identity) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Identity) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTraits

`func (o *Identity) GetTraits() IdentityTraits`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *Identity) GetTraitsOk() (*IdentityTraits, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *Identity) SetTraits(v IdentityTraits)`

SetTraits sets Traits field to given value.


### GetUpdatedAt

`func (o *Identity) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Identity) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Identity) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Identity) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVerifiableAddresses

`func (o *Identity) GetVerifiableAddresses() []VerifiableAddress`

GetVerifiableAddresses returns the VerifiableAddresses field if non-nil, zero value otherwise.

### GetVerifiableAddressesOk

`func (o *Identity) GetVerifiableAddressesOk() (*[]VerifiableAddress, bool)`

GetVerifiableAddressesOk returns a tuple with the VerifiableAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiableAddresses

`func (o *Identity) SetVerifiableAddresses(v []VerifiableAddress)`

SetVerifiableAddresses sets VerifiableAddresses field to given value.

### HasVerifiableAddresses

`func (o *Identity) HasVerifiableAddresses() bool`

HasVerifiableAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


