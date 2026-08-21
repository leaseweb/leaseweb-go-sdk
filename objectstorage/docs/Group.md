# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The group id. | [optional] 
**DisplayName** | Pointer to **string** | The group name. | [optional] 
**UniqueName** | Pointer to **string** | The group unique name. | [optional] 
**S3Policies** | Pointer to **NullableString** | A JSON-encoded IAM-style policy document (e.g. &#x60;{\&quot;Statement\&quot;:[{\&quot;Effect\&quot;:\&quot;Allow\&quot;,\&quot;Action\&quot;:\&quot;s3:*\&quot;,\&quot;Resource\&quot;:\&quot;arn:aws:s3:::*\&quot;}]}&#x60;), or null if no policy is set. | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *Group) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Group) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Group) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Group) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUniqueName

`func (o *Group) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *Group) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *Group) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *Group) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetS3Policies

`func (o *Group) GetS3Policies() string`

GetS3Policies returns the S3Policies field if non-nil, zero value otherwise.

### GetS3PoliciesOk

`func (o *Group) GetS3PoliciesOk() (*string, bool)`

GetS3PoliciesOk returns a tuple with the S3Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Policies

`func (o *Group) SetS3Policies(v string)`

SetS3Policies sets S3Policies field to given value.

### HasS3Policies

`func (o *Group) HasS3Policies() bool`

HasS3Policies returns a boolean if a field has been set.

### SetS3PoliciesNil

`func (o *Group) SetS3PoliciesNil(b bool)`

 SetS3PoliciesNil sets the value for S3Policies to be an explicit nil

### UnsetS3Policies
`func (o *Group) UnsetS3Policies()`

UnsetS3Policies ensures that no value is present for S3Policies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


