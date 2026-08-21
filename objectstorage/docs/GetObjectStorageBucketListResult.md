# GetObjectStorageBucketListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]Bucket**](Bucket.md) | A list of tenant buckets. | [optional] 

## Methods

### NewGetObjectStorageBucketListResult

`func NewGetObjectStorageBucketListResult() *GetObjectStorageBucketListResult`

NewGetObjectStorageBucketListResult instantiates a new GetObjectStorageBucketListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectStorageBucketListResultWithDefaults

`func NewGetObjectStorageBucketListResultWithDefaults() *GetObjectStorageBucketListResult`

NewGetObjectStorageBucketListResultWithDefaults instantiates a new GetObjectStorageBucketListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *GetObjectStorageBucketListResult) GetBuckets() []Bucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *GetObjectStorageBucketListResult) GetBucketsOk() (*[]Bucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *GetObjectStorageBucketListResult) SetBuckets(v []Bucket)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *GetObjectStorageBucketListResult) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


