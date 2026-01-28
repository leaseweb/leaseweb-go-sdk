# GetHardwareMonitoringResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** | The ID of the server | [optional] 
**ScrapedAt** | Pointer to **time.Time** | The date and time the hardware monitoring data was scraped | [optional] 
**IpmiInfo** | Pointer to [**IpmiInfo**](IpmiInfo.md) |  | [optional] 
**Metrics** | Pointer to [**[]IpmiMonitoringMetric**](IpmiMonitoringMetric.md) | IPMI monitoring metrics | [optional] 

## Methods

### NewGetHardwareMonitoringResult

`func NewGetHardwareMonitoringResult() *GetHardwareMonitoringResult`

NewGetHardwareMonitoringResult instantiates a new GetHardwareMonitoringResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHardwareMonitoringResultWithDefaults

`func NewGetHardwareMonitoringResultWithDefaults() *GetHardwareMonitoringResult`

NewGetHardwareMonitoringResultWithDefaults instantiates a new GetHardwareMonitoringResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *GetHardwareMonitoringResult) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetHardwareMonitoringResult) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetHardwareMonitoringResult) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetHardwareMonitoringResult) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetScrapedAt

`func (o *GetHardwareMonitoringResult) GetScrapedAt() time.Time`

GetScrapedAt returns the ScrapedAt field if non-nil, zero value otherwise.

### GetScrapedAtOk

`func (o *GetHardwareMonitoringResult) GetScrapedAtOk() (*time.Time, bool)`

GetScrapedAtOk returns a tuple with the ScrapedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapedAt

`func (o *GetHardwareMonitoringResult) SetScrapedAt(v time.Time)`

SetScrapedAt sets ScrapedAt field to given value.

### HasScrapedAt

`func (o *GetHardwareMonitoringResult) HasScrapedAt() bool`

HasScrapedAt returns a boolean if a field has been set.

### GetIpmiInfo

`func (o *GetHardwareMonitoringResult) GetIpmiInfo() IpmiInfo`

GetIpmiInfo returns the IpmiInfo field if non-nil, zero value otherwise.

### GetIpmiInfoOk

`func (o *GetHardwareMonitoringResult) GetIpmiInfoOk() (*IpmiInfo, bool)`

GetIpmiInfoOk returns a tuple with the IpmiInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiInfo

`func (o *GetHardwareMonitoringResult) SetIpmiInfo(v IpmiInfo)`

SetIpmiInfo sets IpmiInfo field to given value.

### HasIpmiInfo

`func (o *GetHardwareMonitoringResult) HasIpmiInfo() bool`

HasIpmiInfo returns a boolean if a field has been set.

### GetMetrics

`func (o *GetHardwareMonitoringResult) GetMetrics() []IpmiMonitoringMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetHardwareMonitoringResult) GetMetricsOk() (*[]IpmiMonitoringMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetHardwareMonitoringResult) SetMetrics(v []IpmiMonitoringMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetHardwareMonitoringResult) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


