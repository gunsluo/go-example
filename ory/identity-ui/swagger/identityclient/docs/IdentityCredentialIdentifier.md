# IdentityCredentialIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is a helper struct field for. | [optional] 
**Id** | Pointer to **string** | ID is the identity credential&#39;s unique identifier. | [optional] 
**Identifier** | Pointer to **string** | Identifier is the identifier, e.g. email, mobile or others. | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is a helper struct field for. | [optional] 

## Methods

### NewIdentityCredentialIdentifier

`func NewIdentityCredentialIdentifier() *IdentityCredentialIdentifier`

NewIdentityCredentialIdentifier instantiates a new IdentityCredentialIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityCredentialIdentifierWithDefaults

`func NewIdentityCredentialIdentifierWithDefaults() *IdentityCredentialIdentifier`

NewIdentityCredentialIdentifierWithDefaults instantiates a new IdentityCredentialIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IdentityCredentialIdentifier) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdentityCredentialIdentifier) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdentityCredentialIdentifier) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdentityCredentialIdentifier) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *IdentityCredentialIdentifier) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityCredentialIdentifier) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityCredentialIdentifier) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityCredentialIdentifier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *IdentityCredentialIdentifier) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *IdentityCredentialIdentifier) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *IdentityCredentialIdentifier) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *IdentityCredentialIdentifier) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetProperties

`func (o *IdentityCredentialIdentifier) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IdentityCredentialIdentifier) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IdentityCredentialIdentifier) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IdentityCredentialIdentifier) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IdentityCredentialIdentifier) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IdentityCredentialIdentifier) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IdentityCredentialIdentifier) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IdentityCredentialIdentifier) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


