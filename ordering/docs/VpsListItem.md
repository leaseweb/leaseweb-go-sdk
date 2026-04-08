# VpsListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the VPS | [optional] 
**Name** | Pointer to **string** | Name of the VPS | [optional] 
**VCpu** | Pointer to **string** | Number of virtual CPUs | [optional] 
**VRam** | Pointer to **string** | Virtual RAM in GB | [optional] 
**NvmeStorage** | Pointer to **string** | NVMe storage capacity | [optional] 
**Traffic** | Pointer to **string** | Traffic allowance | [optional] 
**Price** | Pointer to [**VpsListItemPrice**](VpsListItemPrice.md) |  | [optional] 

## Methods

### NewVpsListItem

`func NewVpsListItem() *VpsListItem`

NewVpsListItem instantiates a new VpsListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsListItemWithDefaults

`func NewVpsListItemWithDefaults() *VpsListItem`

NewVpsListItemWithDefaults instantiates a new VpsListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpsListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpsListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpsListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpsListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VpsListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpsListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpsListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpsListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVCpu

`func (o *VpsListItem) GetVCpu() string`

GetVCpu returns the VCpu field if non-nil, zero value otherwise.

### GetVCpuOk

`func (o *VpsListItem) GetVCpuOk() (*string, bool)`

GetVCpuOk returns a tuple with the VCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpu

`func (o *VpsListItem) SetVCpu(v string)`

SetVCpu sets VCpu field to given value.

### HasVCpu

`func (o *VpsListItem) HasVCpu() bool`

HasVCpu returns a boolean if a field has been set.

### GetVRam

`func (o *VpsListItem) GetVRam() string`

GetVRam returns the VRam field if non-nil, zero value otherwise.

### GetVRamOk

`func (o *VpsListItem) GetVRamOk() (*string, bool)`

GetVRamOk returns a tuple with the VRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVRam

`func (o *VpsListItem) SetVRam(v string)`

SetVRam sets VRam field to given value.

### HasVRam

`func (o *VpsListItem) HasVRam() bool`

HasVRam returns a boolean if a field has been set.

### GetNvmeStorage

`func (o *VpsListItem) GetNvmeStorage() string`

GetNvmeStorage returns the NvmeStorage field if non-nil, zero value otherwise.

### GetNvmeStorageOk

`func (o *VpsListItem) GetNvmeStorageOk() (*string, bool)`

GetNvmeStorageOk returns a tuple with the NvmeStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvmeStorage

`func (o *VpsListItem) SetNvmeStorage(v string)`

SetNvmeStorage sets NvmeStorage field to given value.

### HasNvmeStorage

`func (o *VpsListItem) HasNvmeStorage() bool`

HasNvmeStorage returns a boolean if a field has been set.

### GetTraffic

`func (o *VpsListItem) GetTraffic() string`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *VpsListItem) GetTrafficOk() (*string, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *VpsListItem) SetTraffic(v string)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *VpsListItem) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetPrice

`func (o *VpsListItem) GetPrice() VpsListItemPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *VpsListItem) GetPriceOk() (*VpsListItemPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *VpsListItem) SetPrice(v VpsListItemPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *VpsListItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


