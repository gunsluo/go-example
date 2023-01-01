# IdentityCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | CreatedAt is a helper struct field for. | [optional] 
**CredentialType** | Pointer to [**CredentialsType**](CredentialsType.md) |  | [optional] 
**Id** | Pointer to **string** | ID is the identity credential&#39;s unique identifier. | [optional] 
**Identifiers** | Pointer to [**[]IdentityCredentialIdentifier**](IdentityCredentialIdentifier.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is a helper struct field for. | [optional] 

## Methods

### NewIdentityCredential

`func NewIdentityCredential() *IdentityCredential`

NewIdentityCredential instantiates a new IdentityCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityCredentialWithDefaults

`func NewIdentityCredentialWithDefaults() *IdentityCredential`

NewIdentityCredentialWithDefaults instantiates a new IdentityCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *IdentityCredential) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IdentityCredential) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IdentityCredential) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IdentityCredential) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IdentityCredential) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdentityCredential) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdentityCredential) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdentityCredential) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentialType

`func (o *IdentityCredential) GetCredentialType() CredentialsType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *IdentityCredential) GetCredentialTypeOk() (*CredentialsType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *IdentityCredential) SetCredentialType(v CredentialsType)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *IdentityCredential) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetId

`func (o *IdentityCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityCredential) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifiers

`func (o *IdentityCredential) GetIdentifiers() []IdentityCredentialIdentifier`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *IdentityCredential) GetIdentifiersOk() (*[]IdentityCredentialIdentifier, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *IdentityCredential) SetIdentifiers(v []IdentityCredentialIdentifier)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *IdentityCredential) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IdentityCredential) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IdentityCredential) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IdentityCredential) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IdentityCredential) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


