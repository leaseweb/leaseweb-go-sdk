# RescueModePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PowerCycle** | Pointer to **bool** | Whether the server should be powered off and on automatically before the job is started | [optional] 
**Network** | Pointer to [**DefaultPayloadNetwork**](DefaultPayloadNetwork.md) |  | [optional] 
**Site** | Pointer to **string** | Location of the server | [optional] 
**InitiatedBy** | Pointer to **string** | Who initiated the job | [optional] 
**ServerBrand** | Pointer to **string** | The brand of the server | [optional] 
**ServerChassis** | Pointer to **string** | The chassis of the server | [optional] 
**FileserverBaseUrl** | Pointer to **string** | The base URL of the fileserver | [optional] 
**ServerHardwareRaid** | Pointer to **bool** | Whether the server has hardware RAID | [optional] 
**RescueImageId** | Pointer to **string** | The ID of the rescue image if job is for rescue mode | [optional] 
**SshKeys** | Pointer to **NullableString** | SSH keys to be added to the server | [optional] 
**CallbackUrl** | Pointer to **NullableString** | The URL to be called when the job is finished | [optional] 
**DoEmailNotification** | Pointer to **NullableBool** | Whether to send email notifications | [optional] 
**Configurable** | Pointer to **NullableBool** | Whether the job is configurable | [optional] 

## Methods

### NewRescueModePayload

`func NewRescueModePayload() *RescueModePayload`

NewRescueModePayload instantiates a new RescueModePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRescueModePayloadWithDefaults

`func NewRescueModePayloadWithDefaults() *RescueModePayload`

NewRescueModePayloadWithDefaults instantiates a new RescueModePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerCycle

`func (o *RescueModePayload) GetPowerCycle() bool`

GetPowerCycle returns the PowerCycle field if non-nil, zero value otherwise.

### GetPowerCycleOk

`func (o *RescueModePayload) GetPowerCycleOk() (*bool, bool)`

GetPowerCycleOk returns a tuple with the PowerCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycle

`func (o *RescueModePayload) SetPowerCycle(v bool)`

SetPowerCycle sets PowerCycle field to given value.

### HasPowerCycle

`func (o *RescueModePayload) HasPowerCycle() bool`

HasPowerCycle returns a boolean if a field has been set.

### GetNetwork

`func (o *RescueModePayload) GetNetwork() DefaultPayloadNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RescueModePayload) GetNetworkOk() (*DefaultPayloadNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RescueModePayload) SetNetwork(v DefaultPayloadNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *RescueModePayload) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSite

`func (o *RescueModePayload) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *RescueModePayload) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *RescueModePayload) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *RescueModePayload) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetInitiatedBy

`func (o *RescueModePayload) GetInitiatedBy() string`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *RescueModePayload) GetInitiatedByOk() (*string, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedBy

`func (o *RescueModePayload) SetInitiatedBy(v string)`

SetInitiatedBy sets InitiatedBy field to given value.

### HasInitiatedBy

`func (o *RescueModePayload) HasInitiatedBy() bool`

HasInitiatedBy returns a boolean if a field has been set.

### GetServerBrand

`func (o *RescueModePayload) GetServerBrand() string`

GetServerBrand returns the ServerBrand field if non-nil, zero value otherwise.

### GetServerBrandOk

`func (o *RescueModePayload) GetServerBrandOk() (*string, bool)`

GetServerBrandOk returns a tuple with the ServerBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBrand

`func (o *RescueModePayload) SetServerBrand(v string)`

SetServerBrand sets ServerBrand field to given value.

### HasServerBrand

`func (o *RescueModePayload) HasServerBrand() bool`

HasServerBrand returns a boolean if a field has been set.

### GetServerChassis

`func (o *RescueModePayload) GetServerChassis() string`

GetServerChassis returns the ServerChassis field if non-nil, zero value otherwise.

### GetServerChassisOk

`func (o *RescueModePayload) GetServerChassisOk() (*string, bool)`

GetServerChassisOk returns a tuple with the ServerChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerChassis

`func (o *RescueModePayload) SetServerChassis(v string)`

SetServerChassis sets ServerChassis field to given value.

### HasServerChassis

`func (o *RescueModePayload) HasServerChassis() bool`

HasServerChassis returns a boolean if a field has been set.

### GetFileserverBaseUrl

`func (o *RescueModePayload) GetFileserverBaseUrl() string`

GetFileserverBaseUrl returns the FileserverBaseUrl field if non-nil, zero value otherwise.

### GetFileserverBaseUrlOk

`func (o *RescueModePayload) GetFileserverBaseUrlOk() (*string, bool)`

GetFileserverBaseUrlOk returns a tuple with the FileserverBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileserverBaseUrl

`func (o *RescueModePayload) SetFileserverBaseUrl(v string)`

SetFileserverBaseUrl sets FileserverBaseUrl field to given value.

### HasFileserverBaseUrl

`func (o *RescueModePayload) HasFileserverBaseUrl() bool`

HasFileserverBaseUrl returns a boolean if a field has been set.

### GetServerHardwareRaid

`func (o *RescueModePayload) GetServerHardwareRaid() bool`

GetServerHardwareRaid returns the ServerHardwareRaid field if non-nil, zero value otherwise.

### GetServerHardwareRaidOk

`func (o *RescueModePayload) GetServerHardwareRaidOk() (*bool, bool)`

GetServerHardwareRaidOk returns a tuple with the ServerHardwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHardwareRaid

`func (o *RescueModePayload) SetServerHardwareRaid(v bool)`

SetServerHardwareRaid sets ServerHardwareRaid field to given value.

### HasServerHardwareRaid

`func (o *RescueModePayload) HasServerHardwareRaid() bool`

HasServerHardwareRaid returns a boolean if a field has been set.

### GetRescueImageId

`func (o *RescueModePayload) GetRescueImageId() string`

GetRescueImageId returns the RescueImageId field if non-nil, zero value otherwise.

### GetRescueImageIdOk

`func (o *RescueModePayload) GetRescueImageIdOk() (*string, bool)`

GetRescueImageIdOk returns a tuple with the RescueImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescueImageId

`func (o *RescueModePayload) SetRescueImageId(v string)`

SetRescueImageId sets RescueImageId field to given value.

### HasRescueImageId

`func (o *RescueModePayload) HasRescueImageId() bool`

HasRescueImageId returns a boolean if a field has been set.

### GetSshKeys

`func (o *RescueModePayload) GetSshKeys() string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *RescueModePayload) GetSshKeysOk() (*string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *RescueModePayload) SetSshKeys(v string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *RescueModePayload) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### SetSshKeysNil

`func (o *RescueModePayload) SetSshKeysNil(b bool)`

 SetSshKeysNil sets the value for SshKeys to be an explicit nil

### UnsetSshKeys
`func (o *RescueModePayload) UnsetSshKeys()`

UnsetSshKeys ensures that no value is present for SshKeys, not even an explicit nil
### GetCallbackUrl

`func (o *RescueModePayload) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *RescueModePayload) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *RescueModePayload) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *RescueModePayload) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *RescueModePayload) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *RescueModePayload) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetDoEmailNotification

`func (o *RescueModePayload) GetDoEmailNotification() bool`

GetDoEmailNotification returns the DoEmailNotification field if non-nil, zero value otherwise.

### GetDoEmailNotificationOk

`func (o *RescueModePayload) GetDoEmailNotificationOk() (*bool, bool)`

GetDoEmailNotificationOk returns a tuple with the DoEmailNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoEmailNotification

`func (o *RescueModePayload) SetDoEmailNotification(v bool)`

SetDoEmailNotification sets DoEmailNotification field to given value.

### HasDoEmailNotification

`func (o *RescueModePayload) HasDoEmailNotification() bool`

HasDoEmailNotification returns a boolean if a field has been set.

### SetDoEmailNotificationNil

`func (o *RescueModePayload) SetDoEmailNotificationNil(b bool)`

 SetDoEmailNotificationNil sets the value for DoEmailNotification to be an explicit nil

### UnsetDoEmailNotification
`func (o *RescueModePayload) UnsetDoEmailNotification()`

UnsetDoEmailNotification ensures that no value is present for DoEmailNotification, not even an explicit nil
### GetConfigurable

`func (o *RescueModePayload) GetConfigurable() bool`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *RescueModePayload) GetConfigurableOk() (*bool, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *RescueModePayload) SetConfigurable(v bool)`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *RescueModePayload) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.

### SetConfigurableNil

`func (o *RescueModePayload) SetConfigurableNil(b bool)`

 SetConfigurableNil sets the value for Configurable to be an explicit nil

### UnsetConfigurable
`func (o *RescueModePayload) UnsetConfigurable()`

UnsetConfigurable ensures that no value is present for Configurable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


