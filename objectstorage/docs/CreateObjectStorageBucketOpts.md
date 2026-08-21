# CreateObjectStorageBucketOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The bucket name. Must be 3–63 characters, alphanumeric and hyphens only. | 
**IsVersioningEnabled** | **bool** | Indicates if versioning is enabled for the bucket. | 
**Quota** | Pointer to **NullableInt32** | The quota for the size of the bucket in GB (1–1,000,000). Omit or set to null for no quota. In responses, quota is returned as an object with &#x60;value&#x60; and &#x60;unit&#x60; fields. | [optional] 

## Methods

### NewCreateObjectStorageBucketOpts

`func NewCreateObjectStorageBucketOpts(name string, isVersioningEnabled bool, ) *CreateObjectStorageBucketOpts`

NewCreateObjectStorageBucketOpts instantiates a new CreateObjectStorageBucketOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateObjectStorageBucketOptsWithDefaults

`func NewCreateObjectStorageBucketOptsWithDefaults() *CreateObjectStorageBucketOpts`

NewCreateObjectStorageBucketOptsWithDefaults instantiates a new CreateObjectStorageBucketOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateObjectStorageBucketOpts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateObjectStorageBucketOpts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateObjectStorageBucketOpts) SetName(v string)`

SetName sets Name field to given value.


### GetIsVersioningEnabled

`func (o *CreateObjectStorageBucketOpts) GetIsVersioningEnabled() bool`

GetIsVersioningEnabled returns the IsVersioningEnabled field if non-nil, zero value otherwise.

### GetIsVersioningEnabledOk

`func (o *CreateObjectStorageBucketOpts) GetIsVersioningEnabledOk() (*bool, bool)`

GetIsVersioningEnabledOk returns a tuple with the IsVersioningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVersioningEnabled

`func (o *CreateObjectStorageBucketOpts) SetIsVersioningEnabled(v bool)`

SetIsVersioningEnabled sets IsVersioningEnabled field to given value.


### GetQuota

`func (o *CreateObjectStorageBucketOpts) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *CreateObjectStorageBucketOpts) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *CreateObjectStorageBucketOpts) SetQuota(v int32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *CreateObjectStorageBucketOpts) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *CreateObjectStorageBucketOpts) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *CreateObjectStorageBucketOpts) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


