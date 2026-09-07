# CreateObjectStorageAccessKeyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The access key id. | [optional] 
**DisplayName** | Pointer to **string** | The access key name. | [optional] 
**AccountId** | Pointer to **string** | The account identifier this key belongs to. | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | The datetime string of the key expire date if it has been set, null otherwise. | [optional] 
**AccessKey** | Pointer to **string** | The access key. | [optional] 
**SecretAccessKey** | Pointer to **string** | The access key secret. | [optional] 

## Methods

### NewCreateObjectStorageAccessKeyResult

`func NewCreateObjectStorageAccessKeyResult() *CreateObjectStorageAccessKeyResult`

NewCreateObjectStorageAccessKeyResult instantiates a new CreateObjectStorageAccessKeyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateObjectStorageAccessKeyResultWithDefaults

`func NewCreateObjectStorageAccessKeyResultWithDefaults() *CreateObjectStorageAccessKeyResult`

NewCreateObjectStorageAccessKeyResultWithDefaults instantiates a new CreateObjectStorageAccessKeyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateObjectStorageAccessKeyResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateObjectStorageAccessKeyResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateObjectStorageAccessKeyResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateObjectStorageAccessKeyResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateObjectStorageAccessKeyResult) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateObjectStorageAccessKeyResult) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateObjectStorageAccessKeyResult) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateObjectStorageAccessKeyResult) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateObjectStorageAccessKeyResult) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateObjectStorageAccessKeyResult) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateObjectStorageAccessKeyResult) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateObjectStorageAccessKeyResult) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateObjectStorageAccessKeyResult) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateObjectStorageAccessKeyResult) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateObjectStorageAccessKeyResult) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateObjectStorageAccessKeyResult) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CreateObjectStorageAccessKeyResult) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CreateObjectStorageAccessKeyResult) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetAccessKey

`func (o *CreateObjectStorageAccessKeyResult) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *CreateObjectStorageAccessKeyResult) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *CreateObjectStorageAccessKeyResult) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *CreateObjectStorageAccessKeyResult) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *CreateObjectStorageAccessKeyResult) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *CreateObjectStorageAccessKeyResult) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *CreateObjectStorageAccessKeyResult) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *CreateObjectStorageAccessKeyResult) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


