# ObjectStorageListMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The maximum number of results returned. | [optional] [default to 10]
**Offset** | Pointer to **int32** | Results are returned starting at the given offset. | [optional] [default to 0]
**TotalCount** | Pointer to **int32** | The total amount of results. | [optional] 

## Methods

### NewObjectStorageListMetadata

`func NewObjectStorageListMetadata() *ObjectStorageListMetadata`

NewObjectStorageListMetadata instantiates a new ObjectStorageListMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageListMetadataWithDefaults

`func NewObjectStorageListMetadataWithDefaults() *ObjectStorageListMetadata`

NewObjectStorageListMetadataWithDefaults instantiates a new ObjectStorageListMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ObjectStorageListMetadata) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ObjectStorageListMetadata) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ObjectStorageListMetadata) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ObjectStorageListMetadata) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ObjectStorageListMetadata) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ObjectStorageListMetadata) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ObjectStorageListMetadata) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ObjectStorageListMetadata) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotalCount

`func (o *ObjectStorageListMetadata) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ObjectStorageListMetadata) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ObjectStorageListMetadata) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ObjectStorageListMetadata) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


