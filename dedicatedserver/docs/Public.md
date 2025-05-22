# Public

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IP address of the server | [optional] 
**Mac** | Pointer to **string** | The MAC address of the server&#39;s first network interface | [optional] 
**Cidr** | Pointer to **string** | The CIDR of the server | [optional] 
**Prefix** | Pointer to **int32** | The IP&#39;s prefix | [optional] 
**Gateway** | Pointer to **string** | The network gateway of the server | [optional] 
**Netmask** | Pointer to **string** | The netmask of the server | [optional] 
**Network** | Pointer to **string** | The network of the server | [optional] 
**IsCustom** | Pointer to **bool** | Whether the network is custom | [optional] 
**Broadcast** | Pointer to **string** | The broadcast network of the server | [optional] 
**Nameservers** | Pointer to **[]string** | The nameservers of the server | [optional] 
**NetmaskHex** | Pointer to **string** | The netmask of the server in hex | [optional] 

## Methods

### NewPublic

`func NewPublic() *Public`

NewPublic instantiates a new Public object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicWithDefaults

`func NewPublicWithDefaults() *Public`

NewPublicWithDefaults instantiates a new Public object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *Public) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Public) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Public) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Public) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *Public) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Public) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Public) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *Public) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetCidr

`func (o *Public) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *Public) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *Public) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *Public) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetPrefix

`func (o *Public) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *Public) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *Public) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *Public) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetGateway

`func (o *Public) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Public) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Public) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Public) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetNetmask

`func (o *Public) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *Public) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *Public) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *Public) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *Public) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Public) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Public) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Public) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIsCustom

`func (o *Public) GetIsCustom() bool`

GetIsCustom returns the IsCustom field if non-nil, zero value otherwise.

### GetIsCustomOk

`func (o *Public) GetIsCustomOk() (*bool, bool)`

GetIsCustomOk returns a tuple with the IsCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustom

`func (o *Public) SetIsCustom(v bool)`

SetIsCustom sets IsCustom field to given value.

### HasIsCustom

`func (o *Public) HasIsCustom() bool`

HasIsCustom returns a boolean if a field has been set.

### GetBroadcast

`func (o *Public) GetBroadcast() string`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *Public) GetBroadcastOk() (*string, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *Public) SetBroadcast(v string)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *Public) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetNameservers

`func (o *Public) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *Public) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *Public) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *Public) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetNetmaskHex

`func (o *Public) GetNetmaskHex() string`

GetNetmaskHex returns the NetmaskHex field if non-nil, zero value otherwise.

### GetNetmaskHexOk

`func (o *Public) GetNetmaskHexOk() (*string, bool)`

GetNetmaskHexOk returns a tuple with the NetmaskHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskHex

`func (o *Public) SetNetmaskHex(v string)`

SetNetmaskHex sets NetmaskHex field to given value.

### HasNetmaskHex

`func (o *Public) HasNetmaskHex() bool`

HasNetmaskHex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


