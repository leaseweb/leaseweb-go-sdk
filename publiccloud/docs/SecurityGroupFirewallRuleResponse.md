# SecurityGroupFirewallRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | The protocol of the firewall rule | 
**StartPort** | **NullableInt32** | Start port for TCP/UDP rules, null for ICMP rules | 
**EndPort** | **NullableInt32** | End port for TCP/UDP rules, null for ICMP rules | 
**IcmpType** | **NullableInt32** | ICMP type for ICMP protocol, null for TCP/UDP rules | 
**IcmpCode** | **NullableInt32** | ICMP code for ICMP protocol, null for TCP/UDP rules | 
**Source** | **string** | Source IP address or CIDR block | 
**State** | [**FirewallRuleState**](FirewallRuleState.md) |  | 

## Methods

### NewSecurityGroupFirewallRuleResponse

`func NewSecurityGroupFirewallRuleResponse(protocol string, startPort NullableInt32, endPort NullableInt32, icmpType NullableInt32, icmpCode NullableInt32, source string, state FirewallRuleState, ) *SecurityGroupFirewallRuleResponse`

NewSecurityGroupFirewallRuleResponse instantiates a new SecurityGroupFirewallRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupFirewallRuleResponseWithDefaults

`func NewSecurityGroupFirewallRuleResponseWithDefaults() *SecurityGroupFirewallRuleResponse`

NewSecurityGroupFirewallRuleResponseWithDefaults instantiates a new SecurityGroupFirewallRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *SecurityGroupFirewallRuleResponse) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityGroupFirewallRuleResponse) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityGroupFirewallRuleResponse) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetStartPort

`func (o *SecurityGroupFirewallRuleResponse) GetStartPort() int32`

GetStartPort returns the StartPort field if non-nil, zero value otherwise.

### GetStartPortOk

`func (o *SecurityGroupFirewallRuleResponse) GetStartPortOk() (*int32, bool)`

GetStartPortOk returns a tuple with the StartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPort

`func (o *SecurityGroupFirewallRuleResponse) SetStartPort(v int32)`

SetStartPort sets StartPort field to given value.


### SetStartPortNil

`func (o *SecurityGroupFirewallRuleResponse) SetStartPortNil(b bool)`

 SetStartPortNil sets the value for StartPort to be an explicit nil

### UnsetStartPort
`func (o *SecurityGroupFirewallRuleResponse) UnsetStartPort()`

UnsetStartPort ensures that no value is present for StartPort, not even an explicit nil
### GetEndPort

`func (o *SecurityGroupFirewallRuleResponse) GetEndPort() int32`

GetEndPort returns the EndPort field if non-nil, zero value otherwise.

### GetEndPortOk

`func (o *SecurityGroupFirewallRuleResponse) GetEndPortOk() (*int32, bool)`

GetEndPortOk returns a tuple with the EndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPort

`func (o *SecurityGroupFirewallRuleResponse) SetEndPort(v int32)`

SetEndPort sets EndPort field to given value.


### SetEndPortNil

`func (o *SecurityGroupFirewallRuleResponse) SetEndPortNil(b bool)`

 SetEndPortNil sets the value for EndPort to be an explicit nil

### UnsetEndPort
`func (o *SecurityGroupFirewallRuleResponse) UnsetEndPort()`

UnsetEndPort ensures that no value is present for EndPort, not even an explicit nil
### GetIcmpType

`func (o *SecurityGroupFirewallRuleResponse) GetIcmpType() int32`

GetIcmpType returns the IcmpType field if non-nil, zero value otherwise.

### GetIcmpTypeOk

`func (o *SecurityGroupFirewallRuleResponse) GetIcmpTypeOk() (*int32, bool)`

GetIcmpTypeOk returns a tuple with the IcmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpType

`func (o *SecurityGroupFirewallRuleResponse) SetIcmpType(v int32)`

SetIcmpType sets IcmpType field to given value.


### SetIcmpTypeNil

`func (o *SecurityGroupFirewallRuleResponse) SetIcmpTypeNil(b bool)`

 SetIcmpTypeNil sets the value for IcmpType to be an explicit nil

### UnsetIcmpType
`func (o *SecurityGroupFirewallRuleResponse) UnsetIcmpType()`

UnsetIcmpType ensures that no value is present for IcmpType, not even an explicit nil
### GetIcmpCode

`func (o *SecurityGroupFirewallRuleResponse) GetIcmpCode() int32`

GetIcmpCode returns the IcmpCode field if non-nil, zero value otherwise.

### GetIcmpCodeOk

`func (o *SecurityGroupFirewallRuleResponse) GetIcmpCodeOk() (*int32, bool)`

GetIcmpCodeOk returns a tuple with the IcmpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpCode

`func (o *SecurityGroupFirewallRuleResponse) SetIcmpCode(v int32)`

SetIcmpCode sets IcmpCode field to given value.


### SetIcmpCodeNil

`func (o *SecurityGroupFirewallRuleResponse) SetIcmpCodeNil(b bool)`

 SetIcmpCodeNil sets the value for IcmpCode to be an explicit nil

### UnsetIcmpCode
`func (o *SecurityGroupFirewallRuleResponse) UnsetIcmpCode()`

UnsetIcmpCode ensures that no value is present for IcmpCode, not even an explicit nil
### GetSource

`func (o *SecurityGroupFirewallRuleResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SecurityGroupFirewallRuleResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SecurityGroupFirewallRuleResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetState

`func (o *SecurityGroupFirewallRuleResponse) GetState() FirewallRuleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SecurityGroupFirewallRuleResponse) GetStateOk() (*FirewallRuleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SecurityGroupFirewallRuleResponse) SetState(v FirewallRuleState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


