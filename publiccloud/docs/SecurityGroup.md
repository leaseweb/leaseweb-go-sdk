# SecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the security group | 
**Name** | **string** | The name of the security group | 
**Default** | **bool** | Whether this is the default security group | 
**State** | [**SecurityGroupState**](SecurityGroupState.md) |  | 

## Methods

### NewSecurityGroup

`func NewSecurityGroup(id string, name string, default_ bool, state SecurityGroupState, ) *SecurityGroup`

NewSecurityGroup instantiates a new SecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupWithDefaults

`func NewSecurityGroupWithDefaults() *SecurityGroup`

NewSecurityGroupWithDefaults instantiates a new SecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SecurityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDefault

`func (o *SecurityGroup) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SecurityGroup) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SecurityGroup) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetState

`func (o *SecurityGroup) GetState() SecurityGroupState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SecurityGroup) GetStateOk() (*SecurityGroupState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SecurityGroup) SetState(v SecurityGroupState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


