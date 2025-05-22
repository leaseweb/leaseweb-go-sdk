# Payload

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

## Methods

### NewPayload

`func NewPayload() *Payload`

NewPayload instantiates a new Payload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadWithDefaults

`func NewPayloadWithDefaults() *Payload`

NewPayloadWithDefaults instantiates a new Payload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerCycle

`func (o *Payload) GetPowerCycle() bool`

GetPowerCycle returns the PowerCycle field if non-nil, zero value otherwise.

### GetPowerCycleOk

`func (o *Payload) GetPowerCycleOk() (*bool, bool)`

GetPowerCycleOk returns a tuple with the PowerCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycle

`func (o *Payload) SetPowerCycle(v bool)`

SetPowerCycle sets PowerCycle field to given value.

### HasPowerCycle

`func (o *Payload) HasPowerCycle() bool`

HasPowerCycle returns a boolean if a field has been set.

### GetNetwork

`func (o *Payload) GetNetwork() DefaultPayloadNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Payload) GetNetworkOk() (*DefaultPayloadNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Payload) SetNetwork(v DefaultPayloadNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Payload) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSite

`func (o *Payload) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Payload) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Payload) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *Payload) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetInitiatedBy

`func (o *Payload) GetInitiatedBy() string`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *Payload) GetInitiatedByOk() (*string, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedBy

`func (o *Payload) SetInitiatedBy(v string)`

SetInitiatedBy sets InitiatedBy field to given value.

### HasInitiatedBy

`func (o *Payload) HasInitiatedBy() bool`

HasInitiatedBy returns a boolean if a field has been set.

### GetServerBrand

`func (o *Payload) GetServerBrand() string`

GetServerBrand returns the ServerBrand field if non-nil, zero value otherwise.

### GetServerBrandOk

`func (o *Payload) GetServerBrandOk() (*string, bool)`

GetServerBrandOk returns a tuple with the ServerBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBrand

`func (o *Payload) SetServerBrand(v string)`

SetServerBrand sets ServerBrand field to given value.

### HasServerBrand

`func (o *Payload) HasServerBrand() bool`

HasServerBrand returns a boolean if a field has been set.

### GetServerChassis

`func (o *Payload) GetServerChassis() string`

GetServerChassis returns the ServerChassis field if non-nil, zero value otherwise.

### GetServerChassisOk

`func (o *Payload) GetServerChassisOk() (*string, bool)`

GetServerChassisOk returns a tuple with the ServerChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerChassis

`func (o *Payload) SetServerChassis(v string)`

SetServerChassis sets ServerChassis field to given value.

### HasServerChassis

`func (o *Payload) HasServerChassis() bool`

HasServerChassis returns a boolean if a field has been set.

### GetFileserverBaseUrl

`func (o *Payload) GetFileserverBaseUrl() string`

GetFileserverBaseUrl returns the FileserverBaseUrl field if non-nil, zero value otherwise.

### GetFileserverBaseUrlOk

`func (o *Payload) GetFileserverBaseUrlOk() (*string, bool)`

GetFileserverBaseUrlOk returns a tuple with the FileserverBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileserverBaseUrl

`func (o *Payload) SetFileserverBaseUrl(v string)`

SetFileserverBaseUrl sets FileserverBaseUrl field to given value.

### HasFileserverBaseUrl

`func (o *Payload) HasFileserverBaseUrl() bool`

HasFileserverBaseUrl returns a boolean if a field has been set.

### GetServerHardwareRaid

`func (o *Payload) GetServerHardwareRaid() bool`

GetServerHardwareRaid returns the ServerHardwareRaid field if non-nil, zero value otherwise.

### GetServerHardwareRaidOk

`func (o *Payload) GetServerHardwareRaidOk() (*bool, bool)`

GetServerHardwareRaidOk returns a tuple with the ServerHardwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHardwareRaid

`func (o *Payload) SetServerHardwareRaid(v bool)`

SetServerHardwareRaid sets ServerHardwareRaid field to given value.

### HasServerHardwareRaid

`func (o *Payload) HasServerHardwareRaid() bool`

HasServerHardwareRaid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


