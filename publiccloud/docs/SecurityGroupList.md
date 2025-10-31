# SecurityGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroups** | [**[]SecurityGroup**](SecurityGroup.md) |  | 
**Metadata** | [**Metadata**](Metadata.md) |  | 

## Methods

### NewSecurityGroupList

`func NewSecurityGroupList(securityGroups []SecurityGroup, metadata Metadata, ) *SecurityGroupList`

NewSecurityGroupList instantiates a new SecurityGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupListWithDefaults

`func NewSecurityGroupListWithDefaults() *SecurityGroupList`

NewSecurityGroupListWithDefaults instantiates a new SecurityGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroups

`func (o *SecurityGroupList) GetSecurityGroups() []SecurityGroup`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *SecurityGroupList) GetSecurityGroupsOk() (*[]SecurityGroup, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *SecurityGroupList) SetSecurityGroups(v []SecurityGroup)`

SetSecurityGroups sets SecurityGroups field to given value.


### GetMetadata

`func (o *SecurityGroupList) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecurityGroupList) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecurityGroupList) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


