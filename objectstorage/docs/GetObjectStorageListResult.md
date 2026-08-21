# GetObjectStorageListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectStorages** | Pointer to [**[]ObjectStorageListItem**](ObjectStorageListItem.md) | A list of object storage products. | [optional] 
**Metadata** | Pointer to [**ObjectStorageListMetadata**](ObjectStorageListMetadata.md) |  | [optional] 

## Methods

### NewGetObjectStorageListResult

`func NewGetObjectStorageListResult() *GetObjectStorageListResult`

NewGetObjectStorageListResult instantiates a new GetObjectStorageListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectStorageListResultWithDefaults

`func NewGetObjectStorageListResultWithDefaults() *GetObjectStorageListResult`

NewGetObjectStorageListResultWithDefaults instantiates a new GetObjectStorageListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectStorages

`func (o *GetObjectStorageListResult) GetObjectStorages() []ObjectStorageListItem`

GetObjectStorages returns the ObjectStorages field if non-nil, zero value otherwise.

### GetObjectStoragesOk

`func (o *GetObjectStorageListResult) GetObjectStoragesOk() (*[]ObjectStorageListItem, bool)`

GetObjectStoragesOk returns a tuple with the ObjectStorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorages

`func (o *GetObjectStorageListResult) SetObjectStorages(v []ObjectStorageListItem)`

SetObjectStorages sets ObjectStorages field to given value.

### HasObjectStorages

`func (o *GetObjectStorageListResult) HasObjectStorages() bool`

HasObjectStorages returns a boolean if a field has been set.

### GetMetadata

`func (o *GetObjectStorageListResult) GetMetadata() ObjectStorageListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetObjectStorageListResult) GetMetadataOk() (*ObjectStorageListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetObjectStorageListResult) SetMetadata(v ObjectStorageListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetObjectStorageListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


