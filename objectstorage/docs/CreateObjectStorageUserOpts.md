# CreateObjectStorageUserOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** | The full name of the user. | 
**UniqueName** | **string** | The unique name identifier for the user. | 
**Groups** | **[]string** | Array of group UUIDs to assign the user to. | 

## Methods

### NewCreateObjectStorageUserOpts

`func NewCreateObjectStorageUserOpts(fullName string, uniqueName string, groups []string, ) *CreateObjectStorageUserOpts`

NewCreateObjectStorageUserOpts instantiates a new CreateObjectStorageUserOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateObjectStorageUserOptsWithDefaults

`func NewCreateObjectStorageUserOptsWithDefaults() *CreateObjectStorageUserOpts`

NewCreateObjectStorageUserOptsWithDefaults instantiates a new CreateObjectStorageUserOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *CreateObjectStorageUserOpts) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *CreateObjectStorageUserOpts) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *CreateObjectStorageUserOpts) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetUniqueName

`func (o *CreateObjectStorageUserOpts) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *CreateObjectStorageUserOpts) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *CreateObjectStorageUserOpts) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.


### GetGroups

`func (o *CreateObjectStorageUserOpts) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CreateObjectStorageUserOpts) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CreateObjectStorageUserOpts) SetGroups(v []string)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


