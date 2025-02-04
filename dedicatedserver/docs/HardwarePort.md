# HardwarePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoNegotiation** | Pointer to [**AutoNegotiation**](AutoNegotiation.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewHardwarePort

`func NewHardwarePort() *HardwarePort`

NewHardwarePort instantiates a new HardwarePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardwarePortWithDefaults

`func NewHardwarePortWithDefaults() *HardwarePort`

NewHardwarePortWithDefaults instantiates a new HardwarePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoNegotiation

`func (o *HardwarePort) GetAutoNegotiation() AutoNegotiation`

GetAutoNegotiation returns the AutoNegotiation field if non-nil, zero value otherwise.

### GetAutoNegotiationOk

`func (o *HardwarePort) GetAutoNegotiationOk() (*AutoNegotiation, bool)`

GetAutoNegotiationOk returns a tuple with the AutoNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNegotiation

`func (o *HardwarePort) SetAutoNegotiation(v AutoNegotiation)`

SetAutoNegotiation sets AutoNegotiation field to given value.

### HasAutoNegotiation

`func (o *HardwarePort) HasAutoNegotiation() bool`

HasAutoNegotiation returns a boolean if a field has been set.

### GetDescription

`func (o *HardwarePort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HardwarePort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HardwarePort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HardwarePort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


