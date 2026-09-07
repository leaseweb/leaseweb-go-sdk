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
**ConfigurationOptions** | Pointer to [**DedicatedServerConfigurationOptions1**](DedicatedServerConfigurationOptions1.md) |  | [optional] 
**Price** | Pointer to [**ProductPricePrice**](ProductPricePrice.md) |  | [optional] 

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

### GetOperatingSystem

`func (o *DedicatedServer) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *DedicatedServer) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *DedicatedServer) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *DedicatedServer) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetInstallOsOnHdd

`func (o *DedicatedServer) GetInstallOsOnHdd() string`

GetInstallOsOnHdd returns the InstallOsOnHdd field if non-nil, zero value otherwise.

### GetInstallOsOnHddOk

`func (o *DedicatedServer) GetInstallOsOnHddOk() (*string, bool)`

GetInstallOsOnHddOk returns a tuple with the InstallOsOnHdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallOsOnHdd

`func (o *DedicatedServer) SetInstallOsOnHdd(v string)`

SetInstallOsOnHdd sets InstallOsOnHdd field to given value.

### HasInstallOsOnHdd

`func (o *DedicatedServer) HasInstallOsOnHdd() bool`

HasInstallOsOnHdd returns a boolean if a field has been set.

### GetControlPanel

`func (o *DedicatedServer) GetControlPanel() string`

GetControlPanel returns the ControlPanel field if non-nil, zero value otherwise.

### GetControlPanelOk

`func (o *DedicatedServer) GetControlPanelOk() (*string, bool)`

GetControlPanelOk returns a tuple with the ControlPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPanel

`func (o *DedicatedServer) SetControlPanel(v string)`

SetControlPanel sets ControlPanel field to given value.

### HasControlPanel

`func (o *DedicatedServer) HasControlPanel() bool`

HasControlPanel returns a boolean if a field has been set.

### GetIpConnectivityType

`func (o *DedicatedServer) GetIpConnectivityType() string`

GetIpConnectivityType returns the IpConnectivityType field if non-nil, zero value otherwise.

### GetIpConnectivityTypeOk

`func (o *DedicatedServer) GetIpConnectivityTypeOk() (*string, bool)`

GetIpConnectivityTypeOk returns a tuple with the IpConnectivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConnectivityType

`func (o *DedicatedServer) SetIpConnectivityType(v string)`

SetIpConnectivityType sets IpConnectivityType field to given value.

### HasIpConnectivityType

`func (o *DedicatedServer) HasIpConnectivityType() bool`

HasIpConnectivityType returns a boolean if a field has been set.

### GetDataPackConfiguration

`func (o *DedicatedServer) GetDataPackConfiguration() string`

GetDataPackConfiguration returns the DataPackConfiguration field if non-nil, zero value otherwise.

### GetDataPackConfigurationOk

`func (o *DedicatedServer) GetDataPackConfigurationOk() (*string, bool)`

GetDataPackConfigurationOk returns a tuple with the DataPackConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPackConfiguration

`func (o *DedicatedServer) SetDataPackConfiguration(v string)`

SetDataPackConfiguration sets DataPackConfiguration field to given value.

### HasDataPackConfiguration

`func (o *DedicatedServer) HasDataPackConfiguration() bool`

HasDataPackConfiguration returns a boolean if a field has been set.

### GetUplinkPortSpeed

`func (o *DedicatedServer) GetUplinkPortSpeed() string`

GetUplinkPortSpeed returns the UplinkPortSpeed field if non-nil, zero value otherwise.

### GetUplinkPortSpeedOk

`func (o *DedicatedServer) GetUplinkPortSpeedOk() (*string, bool)`

GetUplinkPortSpeedOk returns a tuple with the UplinkPortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPortSpeed

`func (o *DedicatedServer) SetUplinkPortSpeed(v string)`

SetUplinkPortSpeed sets UplinkPortSpeed field to given value.

### HasUplinkPortSpeed

`func (o *DedicatedServer) HasUplinkPortSpeed() bool`

HasUplinkPortSpeed returns a boolean if a field has been set.

### GetPrivateNetwork1xvlan

`func (o *DedicatedServer) GetPrivateNetwork1xvlan() string`

GetPrivateNetwork1xvlan returns the PrivateNetwork1xvlan field if non-nil, zero value otherwise.

### GetPrivateNetwork1xvlanOk

`func (o *DedicatedServer) GetPrivateNetwork1xvlanOk() (*string, bool)`

GetPrivateNetwork1xvlanOk returns a tuple with the PrivateNetwork1xvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork1xvlan

`func (o *DedicatedServer) SetPrivateNetwork1xvlan(v string)`

SetPrivateNetwork1xvlan sets PrivateNetwork1xvlan field to given value.

### HasPrivateNetwork1xvlan

`func (o *DedicatedServer) HasPrivateNetwork1xvlan() bool`

HasPrivateNetwork1xvlan returns a boolean if a field has been set.

### GetIpV4Configuration

`func (o *DedicatedServer) GetIpV4Configuration() string`

GetIpV4Configuration returns the IpV4Configuration field if non-nil, zero value otherwise.

### GetIpV4ConfigurationOk

`func (o *DedicatedServer) GetIpV4ConfigurationOk() (*string, bool)`

GetIpV4ConfigurationOk returns a tuple with the IpV4Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Configuration

`func (o *DedicatedServer) SetIpV4Configuration(v string)`

SetIpV4Configuration sets IpV4Configuration field to given value.

### HasIpV4Configuration

`func (o *DedicatedServer) HasIpV4Configuration() bool`

HasIpV4Configuration returns a boolean if a field has been set.

### GetAdditionalServices

`func (o *DedicatedServer) GetAdditionalServices() []string`

GetAdditionalServices returns the AdditionalServices field if non-nil, zero value otherwise.

### GetAdditionalServicesOk

`func (o *DedicatedServer) GetAdditionalServicesOk() (*[]string, bool)`

GetAdditionalServicesOk returns a tuple with the AdditionalServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalServices

`func (o *DedicatedServer) SetAdditionalServices(v []string)`

SetAdditionalServices sets AdditionalServices field to given value.

### HasAdditionalServices

`func (o *DedicatedServer) HasAdditionalServices() bool`

HasAdditionalServices returns a boolean if a field has been set.

### GetMonitoringType

`func (o *DedicatedServer) GetMonitoringType() string`

GetMonitoringType returns the MonitoringType field if non-nil, zero value otherwise.

### GetMonitoringTypeOk

`func (o *DedicatedServer) GetMonitoringTypeOk() (*string, bool)`

GetMonitoringTypeOk returns a tuple with the MonitoringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringType

`func (o *DedicatedServer) SetMonitoringType(v string)`

SetMonitoringType sets MonitoringType field to given value.

### HasMonitoringType

`func (o *DedicatedServer) HasMonitoringType() bool`

HasMonitoringType returns a boolean if a field has been set.

### GetDdosIpProtection

`func (o *DedicatedServer) GetDdosIpProtection() string`

GetDdosIpProtection returns the DdosIpProtection field if non-nil, zero value otherwise.

### GetDdosIpProtectionOk

`func (o *DedicatedServer) GetDdosIpProtectionOk() (*string, bool)`

GetDdosIpProtectionOk returns a tuple with the DdosIpProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosIpProtection

`func (o *DedicatedServer) SetDdosIpProtection(v string)`

SetDdosIpProtection sets DdosIpProtection field to given value.

### HasDdosIpProtection

`func (o *DedicatedServer) HasDdosIpProtection() bool`

HasDdosIpProtection returns a boolean if a field has been set.

### GetServiceLevelAgreement

`func (o *DedicatedServer) GetServiceLevelAgreement() string`

GetServiceLevelAgreement returns the ServiceLevelAgreement field if non-nil, zero value otherwise.

### GetServiceLevelAgreementOk

`func (o *DedicatedServer) GetServiceLevelAgreementOk() (*string, bool)`

GetServiceLevelAgreementOk returns a tuple with the ServiceLevelAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelAgreement

`func (o *DedicatedServer) SetServiceLevelAgreement(v string)`

SetServiceLevelAgreement sets ServiceLevelAgreement field to given value.

### HasServiceLevelAgreement

`func (o *DedicatedServer) HasServiceLevelAgreement() bool`

HasServiceLevelAgreement returns a boolean if a field has been set.

### GetConfigurationOptions

`func (o *DedicatedServer) GetConfigurationOptions() DedicatedServerConfigurationOptions1`

GetConfigurationOptions returns the ConfigurationOptions field if non-nil, zero value otherwise.

### GetConfigurationOptionsOk

`func (o *DedicatedServer) GetConfigurationOptionsOk() (*DedicatedServerConfigurationOptions1, bool)`

GetConfigurationOptionsOk returns a tuple with the ConfigurationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationOptions

`func (o *DedicatedServer) SetConfigurationOptions(v DedicatedServerConfigurationOptions1)`

SetConfigurationOptions sets ConfigurationOptions field to given value.

### HasConfigurationOptions

`func (o *DedicatedServer) HasConfigurationOptions() bool`

HasConfigurationOptions returns a boolean if a field has been set.

### GetPrice

`func (o *DedicatedServer) GetPrice() ProductPricePrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DedicatedServer) GetPriceOk() (*ProductPricePrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DedicatedServer) SetPrice(v ProductPricePrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *DedicatedServer) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


