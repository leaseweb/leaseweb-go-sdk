# IpmiInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BmcUrl** | Pointer to **NullableString** | The Baseboard Management Controller URL | [optional] 
**FirmwareRevision** | Pointer to **NullableString** | The firmware revision of the IPMI controller | [optional] 
**ManufacturerId** | Pointer to **NullableString** | The manufacturer ID of the IPMI controller | [optional] 
**SystemFirmwareVersion** | Pointer to **NullableString** | The system firmware version | [optional] 

## Methods

### NewIpmiInfo

`func NewIpmiInfo() *IpmiInfo`

NewIpmiInfo instantiates a new IpmiInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiInfoWithDefaults

`func NewIpmiInfoWithDefaults() *IpmiInfo`

NewIpmiInfoWithDefaults instantiates a new IpmiInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBmcUrl

`func (o *IpmiInfo) GetBmcUrl() string`

GetBmcUrl returns the BmcUrl field if non-nil, zero value otherwise.

### GetBmcUrlOk

`func (o *IpmiInfo) GetBmcUrlOk() (*string, bool)`

GetBmcUrlOk returns a tuple with the BmcUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcUrl

`func (o *IpmiInfo) SetBmcUrl(v string)`

SetBmcUrl sets BmcUrl field to given value.

### HasBmcUrl

`func (o *IpmiInfo) HasBmcUrl() bool`

HasBmcUrl returns a boolean if a field has been set.

### SetBmcUrlNil

`func (o *IpmiInfo) SetBmcUrlNil(b bool)`

 SetBmcUrlNil sets the value for BmcUrl to be an explicit nil

### UnsetBmcUrl
`func (o *IpmiInfo) UnsetBmcUrl()`

UnsetBmcUrl ensures that no value is present for BmcUrl, not even an explicit nil
### GetFirmwareRevision

`func (o *IpmiInfo) GetFirmwareRevision() string`

GetFirmwareRevision returns the FirmwareRevision field if non-nil, zero value otherwise.

### GetFirmwareRevisionOk

`func (o *IpmiInfo) GetFirmwareRevisionOk() (*string, bool)`

GetFirmwareRevisionOk returns a tuple with the FirmwareRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareRevision

`func (o *IpmiInfo) SetFirmwareRevision(v string)`

SetFirmwareRevision sets FirmwareRevision field to given value.

### HasFirmwareRevision

`func (o *IpmiInfo) HasFirmwareRevision() bool`

HasFirmwareRevision returns a boolean if a field has been set.

### SetFirmwareRevisionNil

`func (o *IpmiInfo) SetFirmwareRevisionNil(b bool)`

 SetFirmwareRevisionNil sets the value for FirmwareRevision to be an explicit nil

### UnsetFirmwareRevision
`func (o *IpmiInfo) UnsetFirmwareRevision()`

UnsetFirmwareRevision ensures that no value is present for FirmwareRevision, not even an explicit nil
### GetManufacturerId

`func (o *IpmiInfo) GetManufacturerId() string`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *IpmiInfo) GetManufacturerIdOk() (*string, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *IpmiInfo) SetManufacturerId(v string)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *IpmiInfo) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### SetManufacturerIdNil

`func (o *IpmiInfo) SetManufacturerIdNil(b bool)`

 SetManufacturerIdNil sets the value for ManufacturerId to be an explicit nil

### UnsetManufacturerId
`func (o *IpmiInfo) UnsetManufacturerId()`

UnsetManufacturerId ensures that no value is present for ManufacturerId, not even an explicit nil
### GetSystemFirmwareVersion

`func (o *IpmiInfo) GetSystemFirmwareVersion() string`

GetSystemFirmwareVersion returns the SystemFirmwareVersion field if non-nil, zero value otherwise.

### GetSystemFirmwareVersionOk

`func (o *IpmiInfo) GetSystemFirmwareVersionOk() (*string, bool)`

GetSystemFirmwareVersionOk returns a tuple with the SystemFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFirmwareVersion

`func (o *IpmiInfo) SetSystemFirmwareVersion(v string)`

SetSystemFirmwareVersion sets SystemFirmwareVersion field to given value.

### HasSystemFirmwareVersion

`func (o *IpmiInfo) HasSystemFirmwareVersion() bool`

HasSystemFirmwareVersion returns a boolean if a field has been set.

### SetSystemFirmwareVersionNil

`func (o *IpmiInfo) SetSystemFirmwareVersionNil(b bool)`

 SetSystemFirmwareVersionNil sets the value for SystemFirmwareVersion to be an explicit nil

### UnsetSystemFirmwareVersion
`func (o *IpmiInfo) UnsetSystemFirmwareVersion()`

UnsetSystemFirmwareVersion ensures that no value is present for SystemFirmwareVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


