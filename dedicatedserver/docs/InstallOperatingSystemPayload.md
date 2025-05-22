# InstallOperatingSystemPayload

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
**Os** | Pointer to [**Os**](Os.md) |  | [optional] 
**Raid** | Pointer to [**NullableRaidPayload**](RaidPayload.md) |  | [optional] 
**Device** | Pointer to **string** | The installation device | [optional] 
**Hostname** | Pointer to **string** | The hostname of the server | [optional] 
**Timezone** | Pointer to **string** | Timezone to be configured on the server | [optional] 
**Partitions** | Pointer to [**[]Partition**](Partition.md) |  | [optional] 
**OperatingSystemId** | Pointer to **string** | The ID of the operating system | [optional] 
**Features** | Pointer to **[]string** | The server&#39;s features | [optional] 
**FeaturesUtilized** | Pointer to **[]string** | The features that are being utilized by the server | [optional] 
**SshKeys** | Pointer to **NullableString** | SSH keys to be added to the server | [optional] 
**Database** | Pointer to [**NullableDatabase**](Database.md) |  | [optional] 
**CallbackUrl** | Pointer to **NullableString** | The URL to be called when the job is finished | [optional] 
**DoEmailNotification** | Pointer to **NullableBool** | Whether to send email notifications | [optional] 
**Configurable** | Pointer to **NullableBool** | Whether the job is configurable | [optional] 

## Methods

### NewInstallOperatingSystemPayload

`func NewInstallOperatingSystemPayload() *InstallOperatingSystemPayload`

NewInstallOperatingSystemPayload instantiates a new InstallOperatingSystemPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallOperatingSystemPayloadWithDefaults

`func NewInstallOperatingSystemPayloadWithDefaults() *InstallOperatingSystemPayload`

NewInstallOperatingSystemPayloadWithDefaults instantiates a new InstallOperatingSystemPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerCycle

`func (o *InstallOperatingSystemPayload) GetPowerCycle() bool`

GetPowerCycle returns the PowerCycle field if non-nil, zero value otherwise.

### GetPowerCycleOk

`func (o *InstallOperatingSystemPayload) GetPowerCycleOk() (*bool, bool)`

GetPowerCycleOk returns a tuple with the PowerCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycle

`func (o *InstallOperatingSystemPayload) SetPowerCycle(v bool)`

SetPowerCycle sets PowerCycle field to given value.

### HasPowerCycle

`func (o *InstallOperatingSystemPayload) HasPowerCycle() bool`

HasPowerCycle returns a boolean if a field has been set.

### GetNetwork

`func (o *InstallOperatingSystemPayload) GetNetwork() DefaultPayloadNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstallOperatingSystemPayload) GetNetworkOk() (*DefaultPayloadNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstallOperatingSystemPayload) SetNetwork(v DefaultPayloadNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InstallOperatingSystemPayload) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSite

`func (o *InstallOperatingSystemPayload) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *InstallOperatingSystemPayload) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *InstallOperatingSystemPayload) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *InstallOperatingSystemPayload) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetInitiatedBy

`func (o *InstallOperatingSystemPayload) GetInitiatedBy() string`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *InstallOperatingSystemPayload) GetInitiatedByOk() (*string, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedBy

`func (o *InstallOperatingSystemPayload) SetInitiatedBy(v string)`

SetInitiatedBy sets InitiatedBy field to given value.

### HasInitiatedBy

`func (o *InstallOperatingSystemPayload) HasInitiatedBy() bool`

HasInitiatedBy returns a boolean if a field has been set.

### GetServerBrand

`func (o *InstallOperatingSystemPayload) GetServerBrand() string`

GetServerBrand returns the ServerBrand field if non-nil, zero value otherwise.

### GetServerBrandOk

`func (o *InstallOperatingSystemPayload) GetServerBrandOk() (*string, bool)`

GetServerBrandOk returns a tuple with the ServerBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBrand

`func (o *InstallOperatingSystemPayload) SetServerBrand(v string)`

SetServerBrand sets ServerBrand field to given value.

### HasServerBrand

`func (o *InstallOperatingSystemPayload) HasServerBrand() bool`

HasServerBrand returns a boolean if a field has been set.

### GetServerChassis

`func (o *InstallOperatingSystemPayload) GetServerChassis() string`

GetServerChassis returns the ServerChassis field if non-nil, zero value otherwise.

### GetServerChassisOk

`func (o *InstallOperatingSystemPayload) GetServerChassisOk() (*string, bool)`

GetServerChassisOk returns a tuple with the ServerChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerChassis

`func (o *InstallOperatingSystemPayload) SetServerChassis(v string)`

SetServerChassis sets ServerChassis field to given value.

### HasServerChassis

`func (o *InstallOperatingSystemPayload) HasServerChassis() bool`

HasServerChassis returns a boolean if a field has been set.

### GetFileserverBaseUrl

`func (o *InstallOperatingSystemPayload) GetFileserverBaseUrl() string`

GetFileserverBaseUrl returns the FileserverBaseUrl field if non-nil, zero value otherwise.

### GetFileserverBaseUrlOk

`func (o *InstallOperatingSystemPayload) GetFileserverBaseUrlOk() (*string, bool)`

GetFileserverBaseUrlOk returns a tuple with the FileserverBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileserverBaseUrl

`func (o *InstallOperatingSystemPayload) SetFileserverBaseUrl(v string)`

SetFileserverBaseUrl sets FileserverBaseUrl field to given value.

### HasFileserverBaseUrl

`func (o *InstallOperatingSystemPayload) HasFileserverBaseUrl() bool`

HasFileserverBaseUrl returns a boolean if a field has been set.

### GetServerHardwareRaid

`func (o *InstallOperatingSystemPayload) GetServerHardwareRaid() bool`

GetServerHardwareRaid returns the ServerHardwareRaid field if non-nil, zero value otherwise.

### GetServerHardwareRaidOk

`func (o *InstallOperatingSystemPayload) GetServerHardwareRaidOk() (*bool, bool)`

GetServerHardwareRaidOk returns a tuple with the ServerHardwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHardwareRaid

`func (o *InstallOperatingSystemPayload) SetServerHardwareRaid(v bool)`

SetServerHardwareRaid sets ServerHardwareRaid field to given value.

### HasServerHardwareRaid

`func (o *InstallOperatingSystemPayload) HasServerHardwareRaid() bool`

HasServerHardwareRaid returns a boolean if a field has been set.

### GetOs

`func (o *InstallOperatingSystemPayload) GetOs() Os`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *InstallOperatingSystemPayload) GetOsOk() (*Os, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *InstallOperatingSystemPayload) SetOs(v Os)`

SetOs sets Os field to given value.

### HasOs

`func (o *InstallOperatingSystemPayload) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetRaid

`func (o *InstallOperatingSystemPayload) GetRaid() RaidPayload`

GetRaid returns the Raid field if non-nil, zero value otherwise.

### GetRaidOk

`func (o *InstallOperatingSystemPayload) GetRaidOk() (*RaidPayload, bool)`

GetRaidOk returns a tuple with the Raid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaid

`func (o *InstallOperatingSystemPayload) SetRaid(v RaidPayload)`

SetRaid sets Raid field to given value.

### HasRaid

`func (o *InstallOperatingSystemPayload) HasRaid() bool`

HasRaid returns a boolean if a field has been set.

### SetRaidNil

`func (o *InstallOperatingSystemPayload) SetRaidNil(b bool)`

 SetRaidNil sets the value for Raid to be an explicit nil

### UnsetRaid
`func (o *InstallOperatingSystemPayload) UnsetRaid()`

UnsetRaid ensures that no value is present for Raid, not even an explicit nil
### GetDevice

`func (o *InstallOperatingSystemPayload) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InstallOperatingSystemPayload) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InstallOperatingSystemPayload) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InstallOperatingSystemPayload) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetHostname

`func (o *InstallOperatingSystemPayload) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InstallOperatingSystemPayload) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InstallOperatingSystemPayload) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *InstallOperatingSystemPayload) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetTimezone

`func (o *InstallOperatingSystemPayload) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *InstallOperatingSystemPayload) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *InstallOperatingSystemPayload) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *InstallOperatingSystemPayload) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetPartitions

`func (o *InstallOperatingSystemPayload) GetPartitions() []Partition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *InstallOperatingSystemPayload) GetPartitionsOk() (*[]Partition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *InstallOperatingSystemPayload) SetPartitions(v []Partition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *InstallOperatingSystemPayload) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetOperatingSystemId

`func (o *InstallOperatingSystemPayload) GetOperatingSystemId() string`

GetOperatingSystemId returns the OperatingSystemId field if non-nil, zero value otherwise.

### GetOperatingSystemIdOk

`func (o *InstallOperatingSystemPayload) GetOperatingSystemIdOk() (*string, bool)`

GetOperatingSystemIdOk returns a tuple with the OperatingSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemId

`func (o *InstallOperatingSystemPayload) SetOperatingSystemId(v string)`

SetOperatingSystemId sets OperatingSystemId field to given value.

### HasOperatingSystemId

`func (o *InstallOperatingSystemPayload) HasOperatingSystemId() bool`

HasOperatingSystemId returns a boolean if a field has been set.

### GetFeatures

`func (o *InstallOperatingSystemPayload) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *InstallOperatingSystemPayload) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *InstallOperatingSystemPayload) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *InstallOperatingSystemPayload) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesUtilized

`func (o *InstallOperatingSystemPayload) GetFeaturesUtilized() []string`

GetFeaturesUtilized returns the FeaturesUtilized field if non-nil, zero value otherwise.

### GetFeaturesUtilizedOk

`func (o *InstallOperatingSystemPayload) GetFeaturesUtilizedOk() (*[]string, bool)`

GetFeaturesUtilizedOk returns a tuple with the FeaturesUtilized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesUtilized

`func (o *InstallOperatingSystemPayload) SetFeaturesUtilized(v []string)`

SetFeaturesUtilized sets FeaturesUtilized field to given value.

### HasFeaturesUtilized

`func (o *InstallOperatingSystemPayload) HasFeaturesUtilized() bool`

HasFeaturesUtilized returns a boolean if a field has been set.

### GetSshKeys

`func (o *InstallOperatingSystemPayload) GetSshKeys() string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *InstallOperatingSystemPayload) GetSshKeysOk() (*string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *InstallOperatingSystemPayload) SetSshKeys(v string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *InstallOperatingSystemPayload) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### SetSshKeysNil

`func (o *InstallOperatingSystemPayload) SetSshKeysNil(b bool)`

 SetSshKeysNil sets the value for SshKeys to be an explicit nil

### UnsetSshKeys
`func (o *InstallOperatingSystemPayload) UnsetSshKeys()`

UnsetSshKeys ensures that no value is present for SshKeys, not even an explicit nil
### GetDatabase

`func (o *InstallOperatingSystemPayload) GetDatabase() Database`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *InstallOperatingSystemPayload) GetDatabaseOk() (*Database, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *InstallOperatingSystemPayload) SetDatabase(v Database)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *InstallOperatingSystemPayload) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### SetDatabaseNil

`func (o *InstallOperatingSystemPayload) SetDatabaseNil(b bool)`

 SetDatabaseNil sets the value for Database to be an explicit nil

### UnsetDatabase
`func (o *InstallOperatingSystemPayload) UnsetDatabase()`

UnsetDatabase ensures that no value is present for Database, not even an explicit nil
### GetCallbackUrl

`func (o *InstallOperatingSystemPayload) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *InstallOperatingSystemPayload) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *InstallOperatingSystemPayload) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *InstallOperatingSystemPayload) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *InstallOperatingSystemPayload) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *InstallOperatingSystemPayload) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetDoEmailNotification

`func (o *InstallOperatingSystemPayload) GetDoEmailNotification() bool`

GetDoEmailNotification returns the DoEmailNotification field if non-nil, zero value otherwise.

### GetDoEmailNotificationOk

`func (o *InstallOperatingSystemPayload) GetDoEmailNotificationOk() (*bool, bool)`

GetDoEmailNotificationOk returns a tuple with the DoEmailNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoEmailNotification

`func (o *InstallOperatingSystemPayload) SetDoEmailNotification(v bool)`

SetDoEmailNotification sets DoEmailNotification field to given value.

### HasDoEmailNotification

`func (o *InstallOperatingSystemPayload) HasDoEmailNotification() bool`

HasDoEmailNotification returns a boolean if a field has been set.

### SetDoEmailNotificationNil

`func (o *InstallOperatingSystemPayload) SetDoEmailNotificationNil(b bool)`

 SetDoEmailNotificationNil sets the value for DoEmailNotification to be an explicit nil

### UnsetDoEmailNotification
`func (o *InstallOperatingSystemPayload) UnsetDoEmailNotification()`

UnsetDoEmailNotification ensures that no value is present for DoEmailNotification, not even an explicit nil
### GetConfigurable

`func (o *InstallOperatingSystemPayload) GetConfigurable() bool`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *InstallOperatingSystemPayload) GetConfigurableOk() (*bool, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *InstallOperatingSystemPayload) SetConfigurable(v bool)`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *InstallOperatingSystemPayload) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.

### SetConfigurableNil

`func (o *InstallOperatingSystemPayload) SetConfigurableNil(b bool)`

 SetConfigurableNil sets the value for Configurable to be an explicit nil

### UnsetConfigurable
`func (o *InstallOperatingSystemPayload) UnsetConfigurable()`

UnsetConfigurable ensures that no value is present for Configurable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


