# AuthorizeRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | The protocol of the firewall rule | 
**StartPort** | Pointer to **NullableInt32** | Start port for TCP/UDP rules. Required for TCP/UDP protocols, not applicable for ICMP | [optional] 
**EndPort** | Pointer to **NullableInt32** | End port for TCP/UDP rules. Required for TCP/UDP protocols, not applicable for ICMP | [optional] 
**IcmpType** | Pointer to **NullableInt32** | ICMP type. Required for ICMP protocol, not applicable for TCP/UDP | [optional] 
**IcmpCode** | Pointer to **NullableInt32** | ICMP code. Required for ICMP protocol, not applicable for TCP/UDP | [optional] 
**Source** | Pointer to **string** | CIDR block | [optional] 

## Methods

### NewAuthorizeRule

`func NewAuthorizeRule(protocol string, ) *AuthorizeRule`

NewAuthorizeRule instantiates a new AuthorizeRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeRuleWithDefaults

`func NewAuthorizeRuleWithDefaults() *AuthorizeRule`

NewAuthorizeRuleWithDefaults instantiates a new AuthorizeRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *AuthorizeRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AuthorizeRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AuthorizeRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetStartPort

`func (o *AuthorizeRule) GetStartPort() int32`

GetStartPort returns the StartPort field if non-nil, zero value otherwise.

### GetStartPortOk

`func (o *AuthorizeRule) GetStartPortOk() (*int32, bool)`

GetStartPortOk returns a tuple with the StartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPort

`func (o *AuthorizeRule) SetStartPort(v int32)`

SetStartPort sets StartPort field to given value.

### HasStartPort

`func (o *AuthorizeRule) HasStartPort() bool`

HasStartPort returns a boolean if a field has been set.

### SetStartPortNil

`func (o *AuthorizeRule) SetStartPortNil(b bool)`

 SetStartPortNil sets the value for StartPort to be an explicit nil

### UnsetStartPort
`func (o *AuthorizeRule) UnsetStartPort()`

UnsetStartPort ensures that no value is present for StartPort, not even an explicit nil
### GetEndPort

`func (o *AuthorizeRule) GetEndPort() int32`

GetEndPort returns the EndPort field if non-nil, zero value otherwise.

### GetEndPortOk

`func (o *AuthorizeRule) GetEndPortOk() (*int32, bool)`

GetEndPortOk returns a tuple with the EndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPort

`func (o *AuthorizeRule) SetEndPort(v int32)`

SetEndPort sets EndPort field to given value.

### HasEndPort

`func (o *AuthorizeRule) HasEndPort() bool`

HasEndPort returns a boolean if a field has been set.

### SetEndPortNil

`func (o *AuthorizeRule) SetEndPortNil(b bool)`

 SetEndPortNil sets the value for EndPort to be an explicit nil

### UnsetEndPort
`func (o *AuthorizeRule) UnsetEndPort()`

UnsetEndPort ensures that no value is present for EndPort, not even an explicit nil
### GetIcmpType

`func (o *AuthorizeRule) GetIcmpType() int32`

GetIcmpType returns the IcmpType field if non-nil, zero value otherwise.

### GetIcmpTypeOk

`func (o *AuthorizeRule) GetIcmpTypeOk() (*int32, bool)`

GetIcmpTypeOk returns a tuple with the IcmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpType

`func (o *AuthorizeRule) SetIcmpType(v int32)`

SetIcmpType sets IcmpType field to given value.

### HasIcmpType

`func (o *AuthorizeRule) HasIcmpType() bool`

HasIcmpType returns a boolean if a field has been set.

### SetIcmpTypeNil

`func (o *AuthorizeRule) SetIcmpTypeNil(b bool)`

 SetIcmpTypeNil sets the value for IcmpType to be an explicit nil

### UnsetIcmpType
`func (o *AuthorizeRule) UnsetIcmpType()`

UnsetIcmpType ensures that no value is present for IcmpType, not even an explicit nil
### GetIcmpCode

`func (o *AuthorizeRule) GetIcmpCode() int32`

GetIcmpCode returns the IcmpCode field if non-nil, zero value otherwise.

### GetIcmpCodeOk

`func (o *AuthorizeRule) GetIcmpCodeOk() (*int32, bool)`

GetIcmpCodeOk returns a tuple with the IcmpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpCode

`func (o *AuthorizeRule) SetIcmpCode(v int32)`

SetIcmpCode sets IcmpCode field to given value.

### HasIcmpCode

`func (o *AuthorizeRule) HasIcmpCode() bool`

HasIcmpCode returns a boolean if a field has been set.

### SetIcmpCodeNil

`func (o *AuthorizeRule) SetIcmpCodeNil(b bool)`

 SetIcmpCodeNil sets the value for IcmpCode to be an explicit nil

### UnsetIcmpCode
`func (o *AuthorizeRule) UnsetIcmpCode()`

UnsetIcmpCode ensures that no value is present for IcmpCode, not even an explicit nil
### GetSource

`func (o *AuthorizeRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AuthorizeRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AuthorizeRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AuthorizeRule) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


