# GetObjectStorageUserListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]User**](User.md) | A list of tenant users. | [optional] 

## Methods

### NewGetObjectStorageUserListResult

`func NewGetObjectStorageUserListResult() *GetObjectStorageUserListResult`

NewGetObjectStorageUserListResult instantiates a new GetObjectStorageUserListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectStorageUserListResultWithDefaults

`func NewGetObjectStorageUserListResultWithDefaults() *GetObjectStorageUserListResult`

NewGetObjectStorageUserListResultWithDefaults instantiates a new GetObjectStorageUserListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *GetObjectStorageUserListResult) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetObjectStorageUserListResult) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetObjectStorageUserListResult) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GetObjectStorageUserListResult) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


