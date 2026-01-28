# GetHardwareMonitoringListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Servers** | Pointer to [**[]GetHardwareMonitoringResult**](GetHardwareMonitoringResult.md) | An array of servers with hardware monitoring metrics | [optional] 

## Methods

### NewGetHardwareMonitoringListResult

`func NewGetHardwareMonitoringListResult() *GetHardwareMonitoringListResult`

NewGetHardwareMonitoringListResult instantiates a new GetHardwareMonitoringListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHardwareMonitoringListResultWithDefaults

`func NewGetHardwareMonitoringListResultWithDefaults() *GetHardwareMonitoringListResult`

NewGetHardwareMonitoringListResultWithDefaults instantiates a new GetHardwareMonitoringListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetHardwareMonitoringListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetHardwareMonitoringListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetHardwareMonitoringListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetHardwareMonitoringListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetServers

`func (o *GetHardwareMonitoringListResult) GetServers() []GetHardwareMonitoringResult`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetHardwareMonitoringListResult) GetServersOk() (*[]GetHardwareMonitoringResult, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetHardwareMonitoringListResult) SetServers(v []GetHardwareMonitoringResult)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetHardwareMonitoringListResult) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


