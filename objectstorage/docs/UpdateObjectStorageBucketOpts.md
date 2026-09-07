# UpdateObjectStorageBucketOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsVersioningEnabled** | **bool** | Indicates if versioning is enabled for the bucket. | 
**Quota** | Pointer to **NullableInt32** | The quota for the size of the bucket in GB (1–1,000,000). | [optional] 

## Methods

### NewUpdateObjectStorageBucketOpts

`func NewUpdateObjectStorageBucketOpts(isVersioningEnabled bool, ) *UpdateObjectStorageBucketOpts`

NewUpdateObjectStorageBucketOpts instantiates a new UpdateObjectStorageBucketOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateObjectStorageBucketOptsWithDefaults

`func NewUpdateObjectStorageBucketOptsWithDefaults() *UpdateObjectStorageBucketOpts`

NewUpdateObjectStorageBucketOptsWithDefaults instantiates a new UpdateObjectStorageBucketOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsVersioningEnabled

`func (o *UpdateObjectStorageBucketOpts) GetIsVersioningEnabled() bool`

GetIsVersioningEnabled returns the IsVersioningEnabled field if non-nil, zero value otherwise.

### GetIsVersioningEnabledOk

`func (o *UpdateObjectStorageBucketOpts) GetIsVersioningEnabledOk() (*bool, bool)`

GetIsVersioningEnabledOk returns a tuple with the IsVersioningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVersioningEnabled

`func (o *UpdateObjectStorageBucketOpts) SetIsVersioningEnabled(v bool)`

SetIsVersioningEnabled sets IsVersioningEnabled field to given value.


### GetQuota

`func (o *UpdateObjectStorageBucketOpts) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *UpdateObjectStorageBucketOpts) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *UpdateObjectStorageBucketOpts) SetQuota(v int32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *UpdateObjectStorageBucketOpts) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *UpdateObjectStorageBucketOpts) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *UpdateObjectStorageBucketOpts) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


