# InstanceSecurityGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroups** | [**[]SecurityGroup**](SecurityGroup.md) | List of security groups associated with the instance | 
**Metadata** | [**Metadata**](Metadata.md) |  | 

## Methods

### NewInstanceSecurityGroups

`func NewInstanceSecurityGroups(securityGroups []SecurityGroup, metadata Metadata, ) *InstanceSecurityGroups`

NewInstanceSecurityGroups instantiates a new InstanceSecurityGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceSecurityGroupsWithDefaults

`func NewInstanceSecurityGroupsWithDefaults() *InstanceSecurityGroups`

NewInstanceSecurityGroupsWithDefaults instantiates a new InstanceSecurityGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroups

`func (o *InstanceSecurityGroups) GetSecurityGroups() []SecurityGroup`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *InstanceSecurityGroups) GetSecurityGroupsOk() (*[]SecurityGroup, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *InstanceSecurityGroups) SetSecurityGroups(v []SecurityGroup)`

SetSecurityGroups sets SecurityGroups field to given value.


### GetMetadata

`func (o *InstanceSecurityGroups) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InstanceSecurityGroups) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InstanceSecurityGroups) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


