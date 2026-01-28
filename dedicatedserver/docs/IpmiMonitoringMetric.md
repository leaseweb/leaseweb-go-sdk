# IpmiMonitoringMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The type of metric | 
**Name** | Pointer to **NullableString** | The name of the specific metric instance | [optional] 
**Value** | **string** | The value of the metric | 

## Methods

### NewIpmiMonitoringMetric

`func NewIpmiMonitoringMetric(metric string, value string, ) *IpmiMonitoringMetric`

NewIpmiMonitoringMetric instantiates a new IpmiMonitoringMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiMonitoringMetricWithDefaults

`func NewIpmiMonitoringMetricWithDefaults() *IpmiMonitoringMetric`

NewIpmiMonitoringMetricWithDefaults instantiates a new IpmiMonitoringMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *IpmiMonitoringMetric) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *IpmiMonitoringMetric) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *IpmiMonitoringMetric) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetName

`func (o *IpmiMonitoringMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpmiMonitoringMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpmiMonitoringMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpmiMonitoringMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IpmiMonitoringMetric) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IpmiMonitoringMetric) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *IpmiMonitoringMetric) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IpmiMonitoringMetric) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IpmiMonitoringMetric) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


