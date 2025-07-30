# MetricsMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | [**Aggregation**](Aggregation.md) |  | 
**From** | **time.Time** | Start of date interval in ISO-8601 format | 
**Granularity** | [**Granularity**](Granularity.md) |  | 
**To** | **time.Time** | End of date interval in ISO-8601 format | 

## Methods

### NewMetricsMetaData

`func NewMetricsMetaData(aggregation Aggregation, from time.Time, granularity Granularity, to time.Time, ) *MetricsMetaData`

NewMetricsMetaData instantiates a new MetricsMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsMetaDataWithDefaults

`func NewMetricsMetaDataWithDefaults() *MetricsMetaData`

NewMetricsMetaDataWithDefaults instantiates a new MetricsMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *MetricsMetaData) GetAggregation() Aggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MetricsMetaData) GetAggregationOk() (*Aggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MetricsMetaData) SetAggregation(v Aggregation)`

SetAggregation sets Aggregation field to given value.


### GetFrom

`func (o *MetricsMetaData) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MetricsMetaData) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MetricsMetaData) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetGranularity

`func (o *MetricsMetaData) GetGranularity() Granularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *MetricsMetaData) GetGranularityOk() (*Granularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *MetricsMetaData) SetGranularity(v Granularity)`

SetGranularity sets Granularity field to given value.


### GetTo

`func (o *MetricsMetaData) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MetricsMetaData) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MetricsMetaData) SetTo(v time.Time)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


