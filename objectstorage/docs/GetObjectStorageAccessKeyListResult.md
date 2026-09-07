# GetObjectStorageAccessKeyListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to [**[]AccessKey**](AccessKey.md) | A list of access keys. | [optional] 

## Methods

### NewGetObjectStorageAccessKeyListResult

`func NewGetObjectStorageAccessKeyListResult() *GetObjectStorageAccessKeyListResult`

NewGetObjectStorageAccessKeyListResult instantiates a new GetObjectStorageAccessKeyListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectStorageAccessKeyListResultWithDefaults

`func NewGetObjectStorageAccessKeyListResultWithDefaults() *GetObjectStorageAccessKeyListResult`

NewGetObjectStorageAccessKeyListResultWithDefaults instantiates a new GetObjectStorageAccessKeyListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *GetObjectStorageAccessKeyListResult) GetKeys() []AccessKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *GetObjectStorageAccessKeyListResult) GetKeysOk() (*[]AccessKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *GetObjectStorageAccessKeyListResult) SetKeys(v []AccessKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *GetObjectStorageAccessKeyListResult) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


