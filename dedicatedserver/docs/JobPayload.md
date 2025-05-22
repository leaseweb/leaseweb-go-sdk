# JobPayload

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
**AiFrameworks** | Pointer to **[]string** | The AI frameworks that are being installed on the server | [optional] 
**CallbackUrl** | Pointer to **NullableString** | The URL to be called when the job is finished | [optional] 
**DoEmailNotification** | Pointer to **NullableBool** | Whether to send email notifications | [optional] 
**Configurable** | Pointer to **NullableBool** | Whether the job is configurable | [optional] 
**RescueImageId** | Pointer to **string** | The ID of the rescue image if job is for rescue mode | [optional] 

## Methods

### NewJobPayload

`func NewJobPayload() *JobPayload`

NewJobPayload instantiates a new JobPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobPayloadWithDefaults

`func NewJobPayloadWithDefaults() *JobPayload`

NewJobPayloadWithDefaults instantiates a new JobPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerCycle

`func (o *JobPayload) GetPowerCycle() bool`

GetPowerCycle returns the PowerCycle field if non-nil, zero value otherwise.

### GetPowerCycleOk

`func (o *JobPayload) GetPowerCycleOk() (*bool, bool)`

GetPowerCycleOk returns a tuple with the PowerCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycle

`func (o *JobPayload) SetPowerCycle(v bool)`

SetPowerCycle sets PowerCycle field to given value.

### HasPowerCycle

`func (o *JobPayload) HasPowerCycle() bool`

HasPowerCycle returns a boolean if a field has been set.

### GetNetwork

`func (o *JobPayload) GetNetwork() DefaultPayloadNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *JobPayload) GetNetworkOk() (*DefaultPayloadNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *JobPayload) SetNetwork(v DefaultPayloadNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *JobPayload) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSite

`func (o *JobPayload) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *JobPayload) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *JobPayload) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *JobPayload) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetInitiatedBy

`func (o *JobPayload) GetInitiatedBy() string`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *JobPayload) GetInitiatedByOk() (*string, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedBy

`func (o *JobPayload) SetInitiatedBy(v string)`

SetInitiatedBy sets InitiatedBy field to given value.

### HasInitiatedBy

`func (o *JobPayload) HasInitiatedBy() bool`

HasInitiatedBy returns a boolean if a field has been set.

### GetServerBrand

`func (o *JobPayload) GetServerBrand() string`

GetServerBrand returns the ServerBrand field if non-nil, zero value otherwise.

### GetServerBrandOk

`func (o *JobPayload) GetServerBrandOk() (*string, bool)`

GetServerBrandOk returns a tuple with the ServerBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBrand

`func (o *JobPayload) SetServerBrand(v string)`

SetServerBrand sets ServerBrand field to given value.

### HasServerBrand

`func (o *JobPayload) HasServerBrand() bool`

HasServerBrand returns a boolean if a field has been set.

### GetServerChassis

`func (o *JobPayload) GetServerChassis() string`

GetServerChassis returns the ServerChassis field if non-nil, zero value otherwise.

### GetServerChassisOk

`func (o *JobPayload) GetServerChassisOk() (*string, bool)`

GetServerChassisOk returns a tuple with the ServerChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerChassis

`func (o *JobPayload) SetServerChassis(v string)`

SetServerChassis sets ServerChassis field to given value.

### HasServerChassis

`func (o *JobPayload) HasServerChassis() bool`

HasServerChassis returns a boolean if a field has been set.

### GetFileserverBaseUrl

`func (o *JobPayload) GetFileserverBaseUrl() string`

GetFileserverBaseUrl returns the FileserverBaseUrl field if non-nil, zero value otherwise.

### GetFileserverBaseUrlOk

`func (o *JobPayload) GetFileserverBaseUrlOk() (*string, bool)`

GetFileserverBaseUrlOk returns a tuple with the FileserverBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileserverBaseUrl

`func (o *JobPayload) SetFileserverBaseUrl(v string)`

SetFileserverBaseUrl sets FileserverBaseUrl field to given value.

### HasFileserverBaseUrl

`func (o *JobPayload) HasFileserverBaseUrl() bool`

HasFileserverBaseUrl returns a boolean if a field has been set.

### GetServerHardwareRaid

`func (o *JobPayload) GetServerHardwareRaid() bool`

GetServerHardwareRaid returns the ServerHardwareRaid field if non-nil, zero value otherwise.

### GetServerHardwareRaidOk

`func (o *JobPayload) GetServerHardwareRaidOk() (*bool, bool)`

GetServerHardwareRaidOk returns a tuple with the ServerHardwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHardwareRaid

`func (o *JobPayload) SetServerHardwareRaid(v bool)`

SetServerHardwareRaid sets ServerHardwareRaid field to given value.

### HasServerHardwareRaid

`func (o *JobPayload) HasServerHardwareRaid() bool`

HasServerHardwareRaid returns a boolean if a field has been set.

### GetOs

`func (o *JobPayload) GetOs() Os`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *JobPayload) GetOsOk() (*Os, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *JobPayload) SetOs(v Os)`

SetOs sets Os field to given value.

### HasOs

`func (o *JobPayload) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetRaid

`func (o *JobPayload) GetRaid() RaidPayload`

GetRaid returns the Raid field if non-nil, zero value otherwise.

### GetRaidOk

`func (o *JobPayload) GetRaidOk() (*RaidPayload, bool)`

GetRaidOk returns a tuple with the Raid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaid

`func (o *JobPayload) SetRaid(v RaidPayload)`

SetRaid sets Raid field to given value.

### HasRaid

`func (o *JobPayload) HasRaid() bool`

HasRaid returns a boolean if a field has been set.

### SetRaidNil

`func (o *JobPayload) SetRaidNil(b bool)`

 SetRaidNil sets the value for Raid to be an explicit nil

### UnsetRaid
`func (o *JobPayload) UnsetRaid()`

UnsetRaid ensures that no value is present for Raid, not even an explicit nil
### GetDevice

`func (o *JobPayload) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *JobPayload) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *JobPayload) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *JobPayload) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetHostname

`func (o *JobPayload) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *JobPayload) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *JobPayload) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *JobPayload) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetTimezone

`func (o *JobPayload) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *JobPayload) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *JobPayload) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *JobPayload) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetPartitions

`func (o *JobPayload) GetPartitions() []Partition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *JobPayload) GetPartitionsOk() (*[]Partition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *JobPayload) SetPartitions(v []Partition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *JobPayload) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetOperatingSystemId

`func (o *JobPayload) GetOperatingSystemId() string`

GetOperatingSystemId returns the OperatingSystemId field if non-nil, zero value otherwise.

### GetOperatingSystemIdOk

`func (o *JobPayload) GetOperatingSystemIdOk() (*string, bool)`

GetOperatingSystemIdOk returns a tuple with the OperatingSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemId

`func (o *JobPayload) SetOperatingSystemId(v string)`

SetOperatingSystemId sets OperatingSystemId field to given value.

### HasOperatingSystemId

`func (o *JobPayload) HasOperatingSystemId() bool`

HasOperatingSystemId returns a boolean if a field has been set.

### GetFeatures

`func (o *JobPayload) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *JobPayload) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *JobPayload) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *JobPayload) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesUtilized

`func (o *JobPayload) GetFeaturesUtilized() []string`

GetFeaturesUtilized returns the FeaturesUtilized field if non-nil, zero value otherwise.

### GetFeaturesUtilizedOk

`func (o *JobPayload) GetFeaturesUtilizedOk() (*[]string, bool)`

GetFeaturesUtilizedOk returns a tuple with the FeaturesUtilized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesUtilized

`func (o *JobPayload) SetFeaturesUtilized(v []string)`

SetFeaturesUtilized sets FeaturesUtilized field to given value.

### HasFeaturesUtilized

`func (o *JobPayload) HasFeaturesUtilized() bool`

HasFeaturesUtilized returns a boolean if a field has been set.

### GetSshKeys

`func (o *JobPayload) GetSshKeys() string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *JobPayload) GetSshKeysOk() (*string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *JobPayload) SetSshKeys(v string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *JobPayload) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### SetSshKeysNil

`func (o *JobPayload) SetSshKeysNil(b bool)`

 SetSshKeysNil sets the value for SshKeys to be an explicit nil

### UnsetSshKeys
`func (o *JobPayload) UnsetSshKeys()`

UnsetSshKeys ensures that no value is present for SshKeys, not even an explicit nil
### GetDatabase

`func (o *JobPayload) GetDatabase() Database`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *JobPayload) GetDatabaseOk() (*Database, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *JobPayload) SetDatabase(v Database)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *JobPayload) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### SetDatabaseNil

`func (o *JobPayload) SetDatabaseNil(b bool)`

 SetDatabaseNil sets the value for Database to be an explicit nil

### UnsetDatabase
`func (o *JobPayload) UnsetDatabase()`

UnsetDatabase ensures that no value is present for Database, not even an explicit nil
### GetAiFrameworks

`func (o *JobPayload) GetAiFrameworks() []string`

GetAiFrameworks returns the AiFrameworks field if non-nil, zero value otherwise.

### GetAiFrameworksOk

`func (o *JobPayload) GetAiFrameworksOk() (*[]string, bool)`

GetAiFrameworksOk returns a tuple with the AiFrameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiFrameworks

`func (o *JobPayload) SetAiFrameworks(v []string)`

SetAiFrameworks sets AiFrameworks field to given value.

### HasAiFrameworks

`func (o *JobPayload) HasAiFrameworks() bool`

HasAiFrameworks returns a boolean if a field has been set.

### SetAiFrameworksNil

`func (o *JobPayload) SetAiFrameworksNil(b bool)`

 SetAiFrameworksNil sets the value for AiFrameworks to be an explicit nil

### UnsetAiFrameworks
`func (o *JobPayload) UnsetAiFrameworks()`

UnsetAiFrameworks ensures that no value is present for AiFrameworks, not even an explicit nil
### GetCallbackUrl

`func (o *JobPayload) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *JobPayload) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *JobPayload) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *JobPayload) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *JobPayload) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *JobPayload) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetDoEmailNotification

`func (o *JobPayload) GetDoEmailNotification() bool`

GetDoEmailNotification returns the DoEmailNotification field if non-nil, zero value otherwise.

### GetDoEmailNotificationOk

`func (o *JobPayload) GetDoEmailNotificationOk() (*bool, bool)`

GetDoEmailNotificationOk returns a tuple with the DoEmailNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoEmailNotification

`func (o *JobPayload) SetDoEmailNotification(v bool)`

SetDoEmailNotification sets DoEmailNotification field to given value.

### HasDoEmailNotification

`func (o *JobPayload) HasDoEmailNotification() bool`

HasDoEmailNotification returns a boolean if a field has been set.

### SetDoEmailNotificationNil

`func (o *JobPayload) SetDoEmailNotificationNil(b bool)`

 SetDoEmailNotificationNil sets the value for DoEmailNotification to be an explicit nil

### UnsetDoEmailNotification
`func (o *JobPayload) UnsetDoEmailNotification()`

UnsetDoEmailNotification ensures that no value is present for DoEmailNotification, not even an explicit nil
### GetConfigurable

`func (o *JobPayload) GetConfigurable() bool`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *JobPayload) GetConfigurableOk() (*bool, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *JobPayload) SetConfigurable(v bool)`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *JobPayload) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.

### SetConfigurableNil

`func (o *JobPayload) SetConfigurableNil(b bool)`

 SetConfigurableNil sets the value for Configurable to be an explicit nil

### UnsetConfigurable
`func (o *JobPayload) UnsetConfigurable()`

UnsetConfigurable ensures that no value is present for Configurable, not even an explicit nil
### GetRescueImageId

`func (o *JobPayload) GetRescueImageId() string`

GetRescueImageId returns the RescueImageId field if non-nil, zero value otherwise.

### GetRescueImageIdOk

`func (o *JobPayload) GetRescueImageIdOk() (*string, bool)`

GetRescueImageIdOk returns a tuple with the RescueImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescueImageId

`func (o *JobPayload) SetRescueImageId(v string)`

SetRescueImageId sets RescueImageId field to given value.

### HasRescueImageId

`func (o *JobPayload) HasRescueImageId() bool`

HasRescueImageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


