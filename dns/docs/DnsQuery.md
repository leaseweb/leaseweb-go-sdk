# DnsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | The metric unit that&#39;s used | 
**Values** | [**[]MetricValue**](MetricValue.md) |  | 

## Methods

### NewDnsQuery

`func NewDnsQuery(unit string, values []MetricValue, ) *DnsQuery`

NewDnsQuery instantiates a new DnsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsQueryWithDefaults

`func NewDnsQueryWithDefaults() *DnsQuery`

NewDnsQueryWithDefaults instantiates a new DnsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *DnsQuery) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *DnsQuery) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *DnsQuery) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetValues

`func (o *DnsQuery) GetValues() []MetricValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *DnsQuery) GetValuesOk() (*[]MetricValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *DnsQuery) SetValues(v []MetricValue)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


