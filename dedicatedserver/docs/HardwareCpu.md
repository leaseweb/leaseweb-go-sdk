# HardwareCpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**CpuCapabilities**](CpuCapabilities.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Hz** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the server | [optional] 
**Settings** | Pointer to [**CpuSettings**](CpuSettings.md) |  | [optional] 
**Slot** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** | Vendor of the server | [optional] 

## Methods

### NewHardwareCpu

`func NewHardwareCpu() *HardwareCpu`

NewHardwareCpu instantiates a new HardwareCpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardwareCpuWithDefaults

`func NewHardwareCpuWithDefaults() *HardwareCpu`

NewHardwareCpuWithDefaults instantiates a new HardwareCpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *HardwareCpu) GetCapabilities() CpuCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *HardwareCpu) GetCapabilitiesOk() (*CpuCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *HardwareCpu) SetCapabilities(v CpuCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *HardwareCpu) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetDescription

`func (o *HardwareCpu) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HardwareCpu) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HardwareCpu) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HardwareCpu) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHz

`func (o *HardwareCpu) GetHz() string`

GetHz returns the Hz field if non-nil, zero value otherwise.

### GetHzOk

`func (o *HardwareCpu) GetHzOk() (*string, bool)`

GetHzOk returns a tuple with the Hz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHz

`func (o *HardwareCpu) SetHz(v string)`

SetHz sets Hz field to given value.

### HasHz

`func (o *HardwareCpu) HasHz() bool`

HasHz returns a boolean if a field has been set.

### GetSerialNumber

`func (o *HardwareCpu) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *HardwareCpu) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *HardwareCpu) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *HardwareCpu) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSettings

`func (o *HardwareCpu) GetSettings() CpuSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *HardwareCpu) GetSettingsOk() (*CpuSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *HardwareCpu) SetSettings(v CpuSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *HardwareCpu) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSlot

`func (o *HardwareCpu) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *HardwareCpu) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *HardwareCpu) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *HardwareCpu) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetVendor

`func (o *HardwareCpu) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HardwareCpu) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HardwareCpu) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HardwareCpu) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


