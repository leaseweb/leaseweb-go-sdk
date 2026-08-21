# UpdateObjectStorageUserOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** | The full name of the user. | 
**Groups** | **[]string** | Array of group UUIDs to assign the user to. | 

## Methods

### NewUpdateObjectStorageUserOpts

`func NewUpdateObjectStorageUserOpts(fullName string, groups []string, ) *UpdateObjectStorageUserOpts`

NewUpdateObjectStorageUserOpts instantiates a new UpdateObjectStorageUserOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateObjectStorageUserOptsWithDefaults

`func NewUpdateObjectStorageUserOptsWithDefaults() *UpdateObjectStorageUserOpts`

NewUpdateObjectStorageUserOptsWithDefaults instantiates a new UpdateObjectStorageUserOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *UpdateObjectStorageUserOpts) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UpdateObjectStorageUserOpts) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UpdateObjectStorageUserOpts) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetGroups

`func (o *UpdateObjectStorageUserOpts) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UpdateObjectStorageUserOpts) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UpdateObjectStorageUserOpts) SetGroups(v []string)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


