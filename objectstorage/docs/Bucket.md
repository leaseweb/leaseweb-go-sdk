# Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The bucket name. | [optional] 
**Region** | Pointer to **string** | The bucket region. | [optional] 
**IsVersioningEnabled** | Pointer to **bool** | Indicates if the bucket versioning is enabled. | [optional] 
**Quota** | Pointer to [**NullableBucketQuota**](BucketQuota.md) |  | [optional] 
**ObjectCount** | Pointer to **NullableInt32** | The bucket object count, null if not set. | [optional] 
**Used** | Pointer to [**BucketUsedSpace**](BucketUsedSpace.md) |  | [optional] 
**CreationTime** | Pointer to **NullableTime** | The date and time the bucket was created, or null if not yet available. | [optional] 
**IsBeingDeleted** | Pointer to **bool** | Indicates if the bucket is currently being deleted. | [optional] 

## Methods

### NewBucket

`func NewBucket() *Bucket`

NewBucket instantiates a new Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketWithDefaults

`func NewBucketWithDefaults() *Bucket`

NewBucketWithDefaults instantiates a new Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Bucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Bucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *Bucket) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Bucket) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Bucket) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Bucket) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetIsVersioningEnabled

`func (o *Bucket) GetIsVersioningEnabled() bool`

GetIsVersioningEnabled returns the IsVersioningEnabled field if non-nil, zero value otherwise.

### GetIsVersioningEnabledOk

`func (o *Bucket) GetIsVersioningEnabledOk() (*bool, bool)`

GetIsVersioningEnabledOk returns a tuple with the IsVersioningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVersioningEnabled

`func (o *Bucket) SetIsVersioningEnabled(v bool)`

SetIsVersioningEnabled sets IsVersioningEnabled field to given value.

### HasIsVersioningEnabled

`func (o *Bucket) HasIsVersioningEnabled() bool`

HasIsVersioningEnabled returns a boolean if a field has been set.

### GetQuota

`func (o *Bucket) GetQuota() BucketQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Bucket) GetQuotaOk() (*BucketQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Bucket) SetQuota(v BucketQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *Bucket) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *Bucket) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *Bucket) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil
### GetObjectCount

`func (o *Bucket) GetObjectCount() int32`

GetObjectCount returns the ObjectCount field if non-nil, zero value otherwise.

### GetObjectCountOk

`func (o *Bucket) GetObjectCountOk() (*int32, bool)`

GetObjectCountOk returns a tuple with the ObjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCount

`func (o *Bucket) SetObjectCount(v int32)`

SetObjectCount sets ObjectCount field to given value.

### HasObjectCount

`func (o *Bucket) HasObjectCount() bool`

HasObjectCount returns a boolean if a field has been set.

### SetObjectCountNil

`func (o *Bucket) SetObjectCountNil(b bool)`

 SetObjectCountNil sets the value for ObjectCount to be an explicit nil

### UnsetObjectCount
`func (o *Bucket) UnsetObjectCount()`

UnsetObjectCount ensures that no value is present for ObjectCount, not even an explicit nil
### GetUsed

`func (o *Bucket) GetUsed() BucketUsedSpace`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Bucket) GetUsedOk() (*BucketUsedSpace, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Bucket) SetUsed(v BucketUsedSpace)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Bucket) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetCreationTime

`func (o *Bucket) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Bucket) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Bucket) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Bucket) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### SetCreationTimeNil

`func (o *Bucket) SetCreationTimeNil(b bool)`

 SetCreationTimeNil sets the value for CreationTime to be an explicit nil

### UnsetCreationTime
`func (o *Bucket) UnsetCreationTime()`

UnsetCreationTime ensures that no value is present for CreationTime, not even an explicit nil
### GetIsBeingDeleted

`func (o *Bucket) GetIsBeingDeleted() bool`

GetIsBeingDeleted returns the IsBeingDeleted field if non-nil, zero value otherwise.

### GetIsBeingDeletedOk

`func (o *Bucket) GetIsBeingDeletedOk() (*bool, bool)`

GetIsBeingDeletedOk returns a tuple with the IsBeingDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBeingDeleted

`func (o *Bucket) SetIsBeingDeleted(v bool)`

SetIsBeingDeleted sets IsBeingDeleted field to given value.

### HasIsBeingDeleted

`func (o *Bucket) HasIsBeingDeleted() bool`

HasIsBeingDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


