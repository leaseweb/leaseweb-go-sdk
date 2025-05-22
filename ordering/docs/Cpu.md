# Cpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoresPerCPU** | Pointer to **string** | Quantity of cores per CPU | [optional] 
**TotalCores** | Pointer to **NullableString** | Sum of CPU cores | [optional] 
**Quantity** | Pointer to **string** | Quantity of CPUs | [optional] 
**Speed** | Pointer to **string** | CPU speed | [optional] 
**Type** | Pointer to **string** | CPU type | [optional] 

## Methods

### NewCpu

`func NewCpu() *Cpu`

NewCpu instantiates a new Cpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuWithDefaults

`func NewCpuWithDefaults() *Cpu`

NewCpuWithDefaults instantiates a new Cpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoresPerCPU

`func (o *Cpu) GetCoresPerCPU() string`

GetCoresPerCPU returns the CoresPerCPU field if non-nil, zero value otherwise.

### GetCoresPerCPUOk

`func (o *Cpu) GetCoresPerCPUOk() (*string, bool)`

GetCoresPerCPUOk returns a tuple with the CoresPerCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerCPU

`func (o *Cpu) SetCoresPerCPU(v string)`

SetCoresPerCPU sets CoresPerCPU field to given value.

### HasCoresPerCPU

`func (o *Cpu) HasCoresPerCPU() bool`

HasCoresPerCPU returns a boolean if a field has been set.

### GetTotalCores

`func (o *Cpu) GetTotalCores() string`

GetTotalCores returns the TotalCores field if non-nil, zero value otherwise.

### GetTotalCoresOk

`func (o *Cpu) GetTotalCoresOk() (*string, bool)`

GetTotalCoresOk returns a tuple with the TotalCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCores

`func (o *Cpu) SetTotalCores(v string)`

SetTotalCores sets TotalCores field to given value.

### HasTotalCores

`func (o *Cpu) HasTotalCores() bool`

HasTotalCores returns a boolean if a field has been set.

### SetTotalCoresNil

`func (o *Cpu) SetTotalCoresNil(b bool)`

 SetTotalCoresNil sets the value for TotalCores to be an explicit nil

### UnsetTotalCores
`func (o *Cpu) UnsetTotalCores()`

UnsetTotalCores ensures that no value is present for TotalCores, not even an explicit nil
### GetQuantity

`func (o *Cpu) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Cpu) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Cpu) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Cpu) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSpeed

`func (o *Cpu) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *Cpu) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *Cpu) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *Cpu) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetType

`func (o *Cpu) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cpu) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cpu) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Cpu) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


