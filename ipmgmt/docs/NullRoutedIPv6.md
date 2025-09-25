# NullRoutedIPv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP address | 
**NullLevel** | **int32** | Null route level | 
**UnnullingAllowed** | **bool** | Describes whether the current user is allowed to unnull the IP | 

## Methods

### NewNullRoutedIPv6

`func NewNullRoutedIPv6(ip string, nullLevel int32, unnullingAllowed bool, ) *NullRoutedIPv6`

NewNullRoutedIPv6 instantiates a new NullRoutedIPv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullRoutedIPv6WithDefaults

`func NewNullRoutedIPv6WithDefaults() *NullRoutedIPv6`

NewNullRoutedIPv6WithDefaults instantiates a new NullRoutedIPv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *NullRoutedIPv6) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NullRoutedIPv6) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NullRoutedIPv6) SetIp(v string)`

SetIp sets Ip field to given value.


### GetNullLevel

`func (o *NullRoutedIPv6) GetNullLevel() int32`

GetNullLevel returns the NullLevel field if non-nil, zero value otherwise.

### GetNullLevelOk

`func (o *NullRoutedIPv6) GetNullLevelOk() (*int32, bool)`

GetNullLevelOk returns a tuple with the NullLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullLevel

`func (o *NullRoutedIPv6) SetNullLevel(v int32)`

SetNullLevel sets NullLevel field to given value.


### GetUnnullingAllowed

`func (o *NullRoutedIPv6) GetUnnullingAllowed() bool`

GetUnnullingAllowed returns the UnnullingAllowed field if non-nil, zero value otherwise.

### GetUnnullingAllowedOk

`func (o *NullRoutedIPv6) GetUnnullingAllowedOk() (*bool, bool)`

GetUnnullingAllowedOk returns a tuple with the UnnullingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnnullingAllowed

`func (o *NullRoutedIPv6) SetUnnullingAllowed(v bool)`

SetUnnullingAllowed sets UnnullingAllowed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


