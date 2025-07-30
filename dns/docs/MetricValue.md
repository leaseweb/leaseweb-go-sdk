# MetricValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | The time of the current metric | 
**Value** | **int32** | The value of the current metric | 

## Methods

### NewMetricValue

`func NewMetricValue(timestamp time.Time, value int32, ) *MetricValue`

NewMetricValue instantiates a new MetricValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricValueWithDefaults

`func NewMetricValueWithDefaults() *MetricValue`

NewMetricValueWithDefaults instantiates a new MetricValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *MetricValue) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MetricValue) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MetricValue) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetValue

`func (o *MetricValue) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricValue) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricValue) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


