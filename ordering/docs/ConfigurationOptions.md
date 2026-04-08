# ConfigurationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskUpgrade** | Pointer to [**[]VpsConfigurationOption**](VpsConfigurationOption.md) | Available disk upgrade options. Note: This will add additional storage on top of the base NVMe Storage included with your package | [optional] 
**OperatingSystem** | Pointer to [**[]VpsConfigurationOption**](VpsConfigurationOption.md) | Available operating system options | [optional] 
**ControlPanel** | Pointer to [**[]VpsConfigurationOption**](VpsConfigurationOption.md) | Available control panel options | [optional] 
**ServiceLevelAgreement** | Pointer to [**[]VpsConfigurationOption**](VpsConfigurationOption.md) | Available SLA options | [optional] 

## Methods

### NewConfigurationOptions

`func NewConfigurationOptions() *ConfigurationOptions`

NewConfigurationOptions instantiates a new ConfigurationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationOptionsWithDefaults

`func NewConfigurationOptionsWithDefaults() *ConfigurationOptions`

NewConfigurationOptionsWithDefaults instantiates a new ConfigurationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskUpgrade

`func (o *ConfigurationOptions) GetDiskUpgrade() []VpsConfigurationOption`

GetDiskUpgrade returns the DiskUpgrade field if non-nil, zero value otherwise.

### GetDiskUpgradeOk

`func (o *ConfigurationOptions) GetDiskUpgradeOk() (*[]VpsConfigurationOption, bool)`

GetDiskUpgradeOk returns a tuple with the DiskUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUpgrade

`func (o *ConfigurationOptions) SetDiskUpgrade(v []VpsConfigurationOption)`

SetDiskUpgrade sets DiskUpgrade field to given value.

### HasDiskUpgrade

`func (o *ConfigurationOptions) HasDiskUpgrade() bool`

HasDiskUpgrade returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *ConfigurationOptions) GetOperatingSystem() []VpsConfigurationOption`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ConfigurationOptions) GetOperatingSystemOk() (*[]VpsConfigurationOption, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ConfigurationOptions) SetOperatingSystem(v []VpsConfigurationOption)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ConfigurationOptions) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetControlPanel

`func (o *ConfigurationOptions) GetControlPanel() []VpsConfigurationOption`

GetControlPanel returns the ControlPanel field if non-nil, zero value otherwise.

### GetControlPanelOk

`func (o *ConfigurationOptions) GetControlPanelOk() (*[]VpsConfigurationOption, bool)`

GetControlPanelOk returns a tuple with the ControlPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPanel

`func (o *ConfigurationOptions) SetControlPanel(v []VpsConfigurationOption)`

SetControlPanel sets ControlPanel field to given value.

### HasControlPanel

`func (o *ConfigurationOptions) HasControlPanel() bool`

HasControlPanel returns a boolean if a field has been set.

### GetServiceLevelAgreement

`func (o *ConfigurationOptions) GetServiceLevelAgreement() []VpsConfigurationOption`

GetServiceLevelAgreement returns the ServiceLevelAgreement field if non-nil, zero value otherwise.

### GetServiceLevelAgreementOk

`func (o *ConfigurationOptions) GetServiceLevelAgreementOk() (*[]VpsConfigurationOption, bool)`

GetServiceLevelAgreementOk returns a tuple with the ServiceLevelAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelAgreement

`func (o *ConfigurationOptions) SetServiceLevelAgreement(v []VpsConfigurationOption)`

SetServiceLevelAgreement sets ServiceLevelAgreement field to given value.

### HasServiceLevelAgreement

`func (o *ConfigurationOptions) HasServiceLevelAgreement() bool`

HasServiceLevelAgreement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


