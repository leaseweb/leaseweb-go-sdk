# DedicatedServerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatingSystem** | Pointer to **string** | Selected operating system | [optional] 
**InstallOsOnHdd** | Pointer to **string** | Disk set where the OS is installed | [optional] 
**ControlPanel** | Pointer to **string** | Selected control panel | [optional] 
**IpConnectivityType** | Pointer to **string** | IP connectivity type | [optional] 
**DataPackConfiguration** | Pointer to **string** | Data pack configuration | [optional] 
**UplinkPortSpeed** | Pointer to **string** | Uplink port speed | [optional] 
**PrivateNetwork1xvlan** | Pointer to **string** | Private network configuration | [optional] 
**IpV4Configuration** | Pointer to **string** | IPv4 configuration | [optional] 
**AdditionalServices** | Pointer to **[]string** | Additional services | [optional] 
**MonitoringType** | Pointer to **string** | Monitoring type | [optional] 
**DdosIpProtection** | Pointer to **string** | DDoS IP protection | [optional] 
**ServiceLevelAgreement** | Pointer to **string** | Service level agreement | [optional] 

## Methods

### NewDedicatedServerConfiguration

`func NewDedicatedServerConfiguration() *DedicatedServerConfiguration`

NewDedicatedServerConfiguration instantiates a new DedicatedServerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedServerConfigurationWithDefaults

`func NewDedicatedServerConfigurationWithDefaults() *DedicatedServerConfiguration`

NewDedicatedServerConfigurationWithDefaults instantiates a new DedicatedServerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatingSystem

`func (o *DedicatedServerConfiguration) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *DedicatedServerConfiguration) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *DedicatedServerConfiguration) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *DedicatedServerConfiguration) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetInstallOsOnHdd

`func (o *DedicatedServerConfiguration) GetInstallOsOnHdd() string`

GetInstallOsOnHdd returns the InstallOsOnHdd field if non-nil, zero value otherwise.

### GetInstallOsOnHddOk

`func (o *DedicatedServerConfiguration) GetInstallOsOnHddOk() (*string, bool)`

GetInstallOsOnHddOk returns a tuple with the InstallOsOnHdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallOsOnHdd

`func (o *DedicatedServerConfiguration) SetInstallOsOnHdd(v string)`

SetInstallOsOnHdd sets InstallOsOnHdd field to given value.

### HasInstallOsOnHdd

`func (o *DedicatedServerConfiguration) HasInstallOsOnHdd() bool`

HasInstallOsOnHdd returns a boolean if a field has been set.

### GetControlPanel

`func (o *DedicatedServerConfiguration) GetControlPanel() string`

GetControlPanel returns the ControlPanel field if non-nil, zero value otherwise.

### GetControlPanelOk

`func (o *DedicatedServerConfiguration) GetControlPanelOk() (*string, bool)`

GetControlPanelOk returns a tuple with the ControlPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPanel

`func (o *DedicatedServerConfiguration) SetControlPanel(v string)`

SetControlPanel sets ControlPanel field to given value.

### HasControlPanel

`func (o *DedicatedServerConfiguration) HasControlPanel() bool`

HasControlPanel returns a boolean if a field has been set.

### GetIpConnectivityType

`func (o *DedicatedServerConfiguration) GetIpConnectivityType() string`

GetIpConnectivityType returns the IpConnectivityType field if non-nil, zero value otherwise.

### GetIpConnectivityTypeOk

`func (o *DedicatedServerConfiguration) GetIpConnectivityTypeOk() (*string, bool)`

GetIpConnectivityTypeOk returns a tuple with the IpConnectivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConnectivityType

`func (o *DedicatedServerConfiguration) SetIpConnectivityType(v string)`

SetIpConnectivityType sets IpConnectivityType field to given value.

### HasIpConnectivityType

`func (o *DedicatedServerConfiguration) HasIpConnectivityType() bool`

HasIpConnectivityType returns a boolean if a field has been set.

### GetDataPackConfiguration

`func (o *DedicatedServerConfiguration) GetDataPackConfiguration() string`

GetDataPackConfiguration returns the DataPackConfiguration field if non-nil, zero value otherwise.

### GetDataPackConfigurationOk

`func (o *DedicatedServerConfiguration) GetDataPackConfigurationOk() (*string, bool)`

GetDataPackConfigurationOk returns a tuple with the DataPackConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPackConfiguration

`func (o *DedicatedServerConfiguration) SetDataPackConfiguration(v string)`

SetDataPackConfiguration sets DataPackConfiguration field to given value.

### HasDataPackConfiguration

`func (o *DedicatedServerConfiguration) HasDataPackConfiguration() bool`

HasDataPackConfiguration returns a boolean if a field has been set.

### GetUplinkPortSpeed

`func (o *DedicatedServerConfiguration) GetUplinkPortSpeed() string`

GetUplinkPortSpeed returns the UplinkPortSpeed field if non-nil, zero value otherwise.

### GetUplinkPortSpeedOk

`func (o *DedicatedServerConfiguration) GetUplinkPortSpeedOk() (*string, bool)`

GetUplinkPortSpeedOk returns a tuple with the UplinkPortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPortSpeed

`func (o *DedicatedServerConfiguration) SetUplinkPortSpeed(v string)`

SetUplinkPortSpeed sets UplinkPortSpeed field to given value.

### HasUplinkPortSpeed

`func (o *DedicatedServerConfiguration) HasUplinkPortSpeed() bool`

HasUplinkPortSpeed returns a boolean if a field has been set.

### GetPrivateNetwork1xvlan

`func (o *DedicatedServerConfiguration) GetPrivateNetwork1xvlan() string`

GetPrivateNetwork1xvlan returns the PrivateNetwork1xvlan field if non-nil, zero value otherwise.

### GetPrivateNetwork1xvlanOk

`func (o *DedicatedServerConfiguration) GetPrivateNetwork1xvlanOk() (*string, bool)`

GetPrivateNetwork1xvlanOk returns a tuple with the PrivateNetwork1xvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork1xvlan

`func (o *DedicatedServerConfiguration) SetPrivateNetwork1xvlan(v string)`

SetPrivateNetwork1xvlan sets PrivateNetwork1xvlan field to given value.

### HasPrivateNetwork1xvlan

`func (o *DedicatedServerConfiguration) HasPrivateNetwork1xvlan() bool`

HasPrivateNetwork1xvlan returns a boolean if a field has been set.

### GetIpV4Configuration

`func (o *DedicatedServerConfiguration) GetIpV4Configuration() string`

GetIpV4Configuration returns the IpV4Configuration field if non-nil, zero value otherwise.

### GetIpV4ConfigurationOk

`func (o *DedicatedServerConfiguration) GetIpV4ConfigurationOk() (*string, bool)`

GetIpV4ConfigurationOk returns a tuple with the IpV4Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Configuration

`func (o *DedicatedServerConfiguration) SetIpV4Configuration(v string)`

SetIpV4Configuration sets IpV4Configuration field to given value.

### HasIpV4Configuration

`func (o *DedicatedServerConfiguration) HasIpV4Configuration() bool`

HasIpV4Configuration returns a boolean if a field has been set.

### GetAdditionalServices

`func (o *DedicatedServerConfiguration) GetAdditionalServices() []string`

GetAdditionalServices returns the AdditionalServices field if non-nil, zero value otherwise.

### GetAdditionalServicesOk

`func (o *DedicatedServerConfiguration) GetAdditionalServicesOk() (*[]string, bool)`

GetAdditionalServicesOk returns a tuple with the AdditionalServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalServices

`func (o *DedicatedServerConfiguration) SetAdditionalServices(v []string)`

SetAdditionalServices sets AdditionalServices field to given value.

### HasAdditionalServices

`func (o *DedicatedServerConfiguration) HasAdditionalServices() bool`

HasAdditionalServices returns a boolean if a field has been set.

### GetMonitoringType

`func (o *DedicatedServerConfiguration) GetMonitoringType() string`

GetMonitoringType returns the MonitoringType field if non-nil, zero value otherwise.

### GetMonitoringTypeOk

`func (o *DedicatedServerConfiguration) GetMonitoringTypeOk() (*string, bool)`

GetMonitoringTypeOk returns a tuple with the MonitoringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringType

`func (o *DedicatedServerConfiguration) SetMonitoringType(v string)`

SetMonitoringType sets MonitoringType field to given value.

### HasMonitoringType

`func (o *DedicatedServerConfiguration) HasMonitoringType() bool`

HasMonitoringType returns a boolean if a field has been set.

### GetDdosIpProtection

`func (o *DedicatedServerConfiguration) GetDdosIpProtection() string`

GetDdosIpProtection returns the DdosIpProtection field if non-nil, zero value otherwise.

### GetDdosIpProtectionOk

`func (o *DedicatedServerConfiguration) GetDdosIpProtectionOk() (*string, bool)`

GetDdosIpProtectionOk returns a tuple with the DdosIpProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosIpProtection

`func (o *DedicatedServerConfiguration) SetDdosIpProtection(v string)`

SetDdosIpProtection sets DdosIpProtection field to given value.

### HasDdosIpProtection

`func (o *DedicatedServerConfiguration) HasDdosIpProtection() bool`

HasDdosIpProtection returns a boolean if a field has been set.

### GetServiceLevelAgreement

`func (o *DedicatedServerConfiguration) GetServiceLevelAgreement() string`

GetServiceLevelAgreement returns the ServiceLevelAgreement field if non-nil, zero value otherwise.

### GetServiceLevelAgreementOk

`func (o *DedicatedServerConfiguration) GetServiceLevelAgreementOk() (*string, bool)`

GetServiceLevelAgreementOk returns a tuple with the ServiceLevelAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelAgreement

`func (o *DedicatedServerConfiguration) SetServiceLevelAgreement(v string)`

SetServiceLevelAgreement sets ServiceLevelAgreement field to given value.

### HasServiceLevelAgreement

`func (o *DedicatedServerConfiguration) HasServiceLevelAgreement() bool`

HasServiceLevelAgreement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


