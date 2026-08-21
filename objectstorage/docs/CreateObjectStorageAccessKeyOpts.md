# CreateObjectStorageAccessKeyOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **NullableTime** | The expire date for the access key in ISO 8601 datetime format. Can be null. | [optional] 

## Methods

### NewCreateObjectStorageAccessKeyOpts

`func NewCreateObjectStorageAccessKeyOpts() *CreateObjectStorageAccessKeyOpts`

NewCreateObjectStorageAccessKeyOpts instantiates a new CreateObjectStorageAccessKeyOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateObjectStorageAccessKeyOptsWithDefaults

`func NewCreateObjectStorageAccessKeyOptsWithDefaults() *CreateObjectStorageAccessKeyOpts`

NewCreateObjectStorageAccessKeyOptsWithDefaults instantiates a new CreateObjectStorageAccessKeyOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *CreateObjectStorageAccessKeyOpts) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateObjectStorageAccessKeyOpts) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateObjectStorageAccessKeyOpts) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateObjectStorageAccessKeyOpts) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CreateObjectStorageAccessKeyOpts) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CreateObjectStorageAccessKeyOpts) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


