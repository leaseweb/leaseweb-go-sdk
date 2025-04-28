# Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ContactGroup** | **string** |  | 
**Contacts** | **[]string** |  | 

## Methods

### NewChannel

`func NewChannel(type_ string, contactGroup string, contacts []string, ) *Channel`

NewChannel instantiates a new Channel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelWithDefaults

`func NewChannelWithDefaults() *Channel`

NewChannelWithDefaults instantiates a new Channel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Channel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Channel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Channel) SetType(v string)`

SetType sets Type field to given value.


### GetContactGroup

`func (o *Channel) GetContactGroup() string`

GetContactGroup returns the ContactGroup field if non-nil, zero value otherwise.

### GetContactGroupOk

`func (o *Channel) GetContactGroupOk() (*string, bool)`

GetContactGroupOk returns a tuple with the ContactGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactGroup

`func (o *Channel) SetContactGroup(v string)`

SetContactGroup sets ContactGroup field to given value.


### GetContacts

`func (o *Channel) GetContacts() []string`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *Channel) GetContactsOk() (*[]string, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *Channel) SetContacts(v []string)`

SetContacts sets Contacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


