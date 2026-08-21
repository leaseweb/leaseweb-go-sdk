# UpdateObjectStorageGroupOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The display name of the group. | 
**S3Policies** | Pointer to **NullableString** | Optional JSON string defining the S3 policies and permissions for users in this group. | [optional] 

## Methods

### NewUpdateObjectStorageGroupOpts

`func NewUpdateObjectStorageGroupOpts(displayName string, ) *UpdateObjectStorageGroupOpts`

NewUpdateObjectStorageGroupOpts instantiates a new UpdateObjectStorageGroupOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateObjectStorageGroupOptsWithDefaults

`func NewUpdateObjectStorageGroupOptsWithDefaults() *UpdateObjectStorageGroupOpts`

NewUpdateObjectStorageGroupOptsWithDefaults instantiates a new UpdateObjectStorageGroupOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UpdateObjectStorageGroupOpts) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateObjectStorageGroupOpts) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateObjectStorageGroupOpts) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetS3Policies

`func (o *UpdateObjectStorageGroupOpts) GetS3Policies() string`

GetS3Policies returns the S3Policies field if non-nil, zero value otherwise.

### GetS3PoliciesOk

`func (o *UpdateObjectStorageGroupOpts) GetS3PoliciesOk() (*string, bool)`

GetS3PoliciesOk returns a tuple with the S3Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Policies

`func (o *UpdateObjectStorageGroupOpts) SetS3Policies(v string)`

SetS3Policies sets S3Policies field to given value.

### HasS3Policies

`func (o *UpdateObjectStorageGroupOpts) HasS3Policies() bool`

HasS3Policies returns a boolean if a field has been set.

### SetS3PoliciesNil

`func (o *UpdateObjectStorageGroupOpts) SetS3PoliciesNil(b bool)`

 SetS3PoliciesNil sets the value for S3Policies to be an explicit nil

### UnsetS3Policies
`func (o *UpdateObjectStorageGroupOpts) UnsetS3Policies()`

UnsetS3Policies ensures that no value is present for S3Policies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


