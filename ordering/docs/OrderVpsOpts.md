# OrderVpsOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | The datacenter location | 
**DiskUpgrade** | Pointer to **string** | Disk upgrade option | [optional] 
**OperatingSystem** | Pointer to **string** | Operating system option | [optional] 
**ControlPanel** | Pointer to **string** | Control panel option | [optional] 
**ServiceLevelAgreement** | Pointer to [**ServiceLevelAgreement**](ServiceLevelAgreement.md) |  | [optional] 
**ContractTerm** | Pointer to [**ContractTerm**](ContractTerm.md) |  | [optional] [default to CONTRACTTERM__1_YEAR]
**BillingCycle** | Pointer to [**BillingCycle**](BillingCycle.md) |  | [optional] [default to BILLINGCYCLE__1_MONTH]

## Methods

### NewOrderVpsOpts

`func NewOrderVpsOpts(location string, ) *OrderVpsOpts`

NewOrderVpsOpts instantiates a new OrderVpsOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderVpsOptsWithDefaults

`func NewOrderVpsOptsWithDefaults() *OrderVpsOpts`

NewOrderVpsOptsWithDefaults instantiates a new OrderVpsOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *OrderVpsOpts) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OrderVpsOpts) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OrderVpsOpts) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetDiskUpgrade

`func (o *OrderVpsOpts) GetDiskUpgrade() string`

GetDiskUpgrade returns the DiskUpgrade field if non-nil, zero value otherwise.

### GetDiskUpgradeOk

`func (o *OrderVpsOpts) GetDiskUpgradeOk() (*string, bool)`

GetDiskUpgradeOk returns a tuple with the DiskUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUpgrade

`func (o *OrderVpsOpts) SetDiskUpgrade(v string)`

SetDiskUpgrade sets DiskUpgrade field to given value.

### HasDiskUpgrade

`func (o *OrderVpsOpts) HasDiskUpgrade() bool`

HasDiskUpgrade returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *OrderVpsOpts) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *OrderVpsOpts) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *OrderVpsOpts) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *OrderVpsOpts) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetControlPanel

`func (o *OrderVpsOpts) GetControlPanel() string`

GetControlPanel returns the ControlPanel field if non-nil, zero value otherwise.

### GetControlPanelOk

`func (o *OrderVpsOpts) GetControlPanelOk() (*string, bool)`

GetControlPanelOk returns a tuple with the ControlPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPanel

`func (o *OrderVpsOpts) SetControlPanel(v string)`

SetControlPanel sets ControlPanel field to given value.

### HasControlPanel

`func (o *OrderVpsOpts) HasControlPanel() bool`

HasControlPanel returns a boolean if a field has been set.

### GetServiceLevelAgreement

`func (o *OrderVpsOpts) GetServiceLevelAgreement() ServiceLevelAgreement`

GetServiceLevelAgreement returns the ServiceLevelAgreement field if non-nil, zero value otherwise.

### GetServiceLevelAgreementOk

`func (o *OrderVpsOpts) GetServiceLevelAgreementOk() (*ServiceLevelAgreement, bool)`

GetServiceLevelAgreementOk returns a tuple with the ServiceLevelAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelAgreement

`func (o *OrderVpsOpts) SetServiceLevelAgreement(v ServiceLevelAgreement)`

SetServiceLevelAgreement sets ServiceLevelAgreement field to given value.

### HasServiceLevelAgreement

`func (o *OrderVpsOpts) HasServiceLevelAgreement() bool`

HasServiceLevelAgreement returns a boolean if a field has been set.

### GetContractTerm

`func (o *OrderVpsOpts) GetContractTerm() ContractTerm`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *OrderVpsOpts) GetContractTermOk() (*ContractTerm, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *OrderVpsOpts) SetContractTerm(v ContractTerm)`

SetContractTerm sets ContractTerm field to given value.

### HasContractTerm

`func (o *OrderVpsOpts) HasContractTerm() bool`

HasContractTerm returns a boolean if a field has been set.

### GetBillingCycle

`func (o *OrderVpsOpts) GetBillingCycle() BillingCycle`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *OrderVpsOpts) GetBillingCycleOk() (*BillingCycle, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *OrderVpsOpts) SetBillingCycle(v BillingCycle)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *OrderVpsOpts) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


