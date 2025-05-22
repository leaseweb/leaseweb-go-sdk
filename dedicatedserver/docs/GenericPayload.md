# GenericPayload

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

### NewGenericPayload

`func NewGenericPayload() *GenericPayload`

NewGenericPayload instantiates a new GenericPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericPayloadWithDefaults

`func NewGenericPayloadWithDefaults() *GenericPayload`

NewGenericPayloadWithDefaults instantiates a new GenericPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerCycle

`func (o *GenericPayload) GetPowerCycle() bool`

GetPowerCycle returns the PowerCycle field if non-nil, zero value otherwise.

### GetPowerCycleOk

`func (o *GenericPayload) GetPowerCycleOk() (*bool, bool)`

GetPowerCycleOk returns a tuple with the PowerCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycle

`func (o *GenericPayload) SetPowerCycle(v bool)`

SetPowerCycle sets PowerCycle field to given value.

### HasPowerCycle

`func (o *GenericPayload) HasPowerCycle() bool`

HasPowerCycle returns a boolean if a field has been set.

### GetNetwork

`func (o *GenericPayload) GetNetwork() DefaultPayloadNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GenericPayload) GetNetworkOk() (*DefaultPayloadNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GenericPayload) SetNetwork(v DefaultPayloadNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GenericPayload) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSite

`func (o *GenericPayload) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GenericPayload) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GenericPayload) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *GenericPayload) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetInitiatedBy

`func (o *GenericPayload) GetInitiatedBy() string`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *GenericPayload) GetInitiatedByOk() (*string, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedBy

`func (o *GenericPayload) SetInitiatedBy(v string)`

SetInitiatedBy sets InitiatedBy field to given value.

### HasInitiatedBy

`func (o *GenericPayload) HasInitiatedBy() bool`

HasInitiatedBy returns a boolean if a field has been set.

### GetServerBrand

`func (o *GenericPayload) GetServerBrand() string`

GetServerBrand returns the ServerBrand field if non-nil, zero value otherwise.

### GetServerBrandOk

`func (o *GenericPayload) GetServerBrandOk() (*string, bool)`

GetServerBrandOk returns a tuple with the ServerBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBrand

`func (o *GenericPayload) SetServerBrand(v string)`

SetServerBrand sets ServerBrand field to given value.

### HasServerBrand

`func (o *GenericPayload) HasServerBrand() bool`

HasServerBrand returns a boolean if a field has been set.

### GetServerChassis

`func (o *GenericPayload) GetServerChassis() string`

GetServerChassis returns the ServerChassis field if non-nil, zero value otherwise.

### GetServerChassisOk

`func (o *GenericPayload) GetServerChassisOk() (*string, bool)`

GetServerChassisOk returns a tuple with the ServerChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerChassis

`func (o *GenericPayload) SetServerChassis(v string)`

SetServerChassis sets ServerChassis field to given value.

### HasServerChassis

`func (o *GenericPayload) HasServerChassis() bool`

HasServerChassis returns a boolean if a field has been set.

### GetFileserverBaseUrl

`func (o *GenericPayload) GetFileserverBaseUrl() string`

GetFileserverBaseUrl returns the FileserverBaseUrl field if non-nil, zero value otherwise.

### GetFileserverBaseUrlOk

`func (o *GenericPayload) GetFileserverBaseUrlOk() (*string, bool)`

GetFileserverBaseUrlOk returns a tuple with the FileserverBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileserverBaseUrl

`func (o *GenericPayload) SetFileserverBaseUrl(v string)`

SetFileserverBaseUrl sets FileserverBaseUrl field to given value.

### HasFileserverBaseUrl

`func (o *GenericPayload) HasFileserverBaseUrl() bool`

HasFileserverBaseUrl returns a boolean if a field has been set.

### GetServerHardwareRaid

`func (o *GenericPayload) GetServerHardwareRaid() bool`

GetServerHardwareRaid returns the ServerHardwareRaid field if non-nil, zero value otherwise.

### GetServerHardwareRaidOk

`func (o *GenericPayload) GetServerHardwareRaidOk() (*bool, bool)`

GetServerHardwareRaidOk returns a tuple with the ServerHardwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHardwareRaid

`func (o *GenericPayload) SetServerHardwareRaid(v bool)`

SetServerHardwareRaid sets ServerHardwareRaid field to given value.

### HasServerHardwareRaid

`func (o *GenericPayload) HasServerHardwareRaid() bool`

HasServerHardwareRaid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


