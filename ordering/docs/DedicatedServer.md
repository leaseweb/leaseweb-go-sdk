# DedicatedServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the server | [optional] 
**Name** | Pointer to **string** | Name of the server | [optional] 
**Chassis** | Pointer to **string** | Chassis of the server | [optional] 
**Cpu** | Pointer to [**Cpu**](Cpu.md) |  | [optional] 
**DeliveryMethod** | Pointer to **string** | Time to get the server ready | [optional] 
**Hdd** | Pointer to [**[]Hdd**](Hdd.md) | HDD info of the server | [optional] 
**Ram** | Pointer to [**Ram**](Ram.md) |  | [optional] 
**Location** | Pointer to **[]string** | Location of the server | [optional] 
**Price** | Pointer to [**DedicatedServerPricePrice**](DedicatedServerPricePrice.md) |  | [optional] 

## Methods

### NewDedicatedServer

`func NewDedicatedServer() *DedicatedServer`

NewDedicatedServer instantiates a new DedicatedServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedServerWithDefaults

`func NewDedicatedServerWithDefaults() *DedicatedServer`

NewDedicatedServerWithDefaults instantiates a new DedicatedServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DedicatedServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DedicatedServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DedicatedServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DedicatedServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DedicatedServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DedicatedServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DedicatedServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DedicatedServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetChassis

`func (o *DedicatedServer) GetChassis() string`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *DedicatedServer) GetChassisOk() (*string, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *DedicatedServer) SetChassis(v string)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *DedicatedServer) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetCpu

`func (o *DedicatedServer) GetCpu() Cpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *DedicatedServer) GetCpuOk() (*Cpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *DedicatedServer) SetCpu(v Cpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *DedicatedServer) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDeliveryMethod

`func (o *DedicatedServer) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *DedicatedServer) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *DedicatedServer) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *DedicatedServer) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### GetHdd

`func (o *DedicatedServer) GetHdd() []Hdd`

GetHdd returns the Hdd field if non-nil, zero value otherwise.

### GetHddOk

`func (o *DedicatedServer) GetHddOk() (*[]Hdd, bool)`

GetHddOk returns a tuple with the Hdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdd

`func (o *DedicatedServer) SetHdd(v []Hdd)`

SetHdd sets Hdd field to given value.

### HasHdd

`func (o *DedicatedServer) HasHdd() bool`

HasHdd returns a boolean if a field has been set.

### GetRam

`func (o *DedicatedServer) GetRam() Ram`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *DedicatedServer) GetRamOk() (*Ram, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *DedicatedServer) SetRam(v Ram)`

SetRam sets Ram field to given value.

### HasRam

`func (o *DedicatedServer) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetLocation

`func (o *DedicatedServer) GetLocation() []string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DedicatedServer) GetLocationOk() (*[]string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DedicatedServer) SetLocation(v []string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DedicatedServer) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPrice

`func (o *DedicatedServer) GetPrice() DedicatedServerPricePrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DedicatedServer) GetPriceOk() (*DedicatedServerPricePrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DedicatedServer) SetPrice(v DedicatedServerPricePrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *DedicatedServer) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


