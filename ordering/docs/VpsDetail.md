# VpsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the VPS | [optional] 
**Name** | Pointer to **string** | Name of the VPS | [optional] 
**VCpu** | Pointer to **string** | Number of virtual CPUs | [optional] 
**VRam** | Pointer to **string** | Virtual RAM in GB | [optional] 
**NvmeStorage** | Pointer to **string** | Included NVMe storage | [optional] 
**Traffic** | Pointer to **string** | Traffic allowance | [optional] 
**Location** | Pointer to **[]string** | Available locations for the VPS | [optional] 

## Methods

### NewVpsDetail

`func NewVpsDetail() *VpsDetail`

NewVpsDetail instantiates a new VpsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsDetailWithDefaults

`func NewVpsDetailWithDefaults() *VpsDetail`

NewVpsDetailWithDefaults instantiates a new VpsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpsDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpsDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpsDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpsDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VpsDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpsDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpsDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpsDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVCpu

`func (o *VpsDetail) GetVCpu() string`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *VpsDetail) GetVCpuOk() (*string, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *VpsDetail) SetVCpu(v string)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *VpsDetail) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetVRam

`func (o *VpsDetail) GetVRam() string`

GetVRam returns the VRam field if non-nil, zero value otherwise.

### GetVRamOk

`func (o *VpsDetail) GetVRamOk() (*string, bool)`

GetVRamOk returns a tuple with the VRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVRam

`func (o *VpsDetail) SetVRam(v string)`

SetVRam sets VRam field to given value.

### HasVRam

`func (o *VpsDetail) HasVRam() bool`

HasVRam returns a boolean if a field has been set.

### GetNvmeStorage

`func (o *VpsDetail) GetNvmeStorage() string`

GetNvmeStorage returns the NvmeStorage field if non-nil, zero value otherwise.

### GetNvmeStorageOk

`func (o *VpsDetail) GetNvmeStorageOk() (*string, bool)`

GetNvmeStorageOk returns a tuple with the NvmeStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvmeStorage

`func (o *VpsDetail) SetNvmeStorage(v string)`

SetNvmeStorage sets NvmeStorage field to given value.

### HasNvmeStorage

`func (o *VpsDetail) HasNvmeStorage() bool`

HasNvmeStorage returns a boolean if a field has been set.

### GetTraffic

`func (o *VpsDetail) GetTraffic() string`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *VpsDetail) GetTrafficOk() (*string, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *VpsDetail) SetTraffic(v string)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *VpsDetail) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetLocation

`func (o *VpsDetail) GetLocation() []string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VpsDetail) GetLocationOk() (*[]string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VpsDetail) SetLocation(v []string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VpsDetail) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


