# GetQueryMetricsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsQuery** | [**DnsQuery**](DnsQuery.md) |  | 
**Metadata** | [**MetricsMetaData**](MetricsMetaData.md) |  | 

## Methods

### NewGetQueryMetricsResult

`func NewGetQueryMetricsResult(dnsQuery DnsQuery, metadata MetricsMetaData, ) *GetQueryMetricsResult`

NewGetQueryMetricsResult instantiates a new GetQueryMetricsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQueryMetricsResultWithDefaults

`func NewGetQueryMetricsResultWithDefaults() *GetQueryMetricsResult`

NewGetQueryMetricsResultWithDefaults instantiates a new GetQueryMetricsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsQuery

`func (o *GetQueryMetricsResult) GetDnsQuery() DnsQuery`

GetDnsQuery returns the DnsQuery field if non-nil, zero value otherwise.

### GetDnsQueryOk

`func (o *GetQueryMetricsResult) GetDnsQueryOk() (*DnsQuery, bool)`

GetDnsQueryOk returns a tuple with the DnsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsQuery

`func (o *GetQueryMetricsResult) SetDnsQuery(v DnsQuery)`

SetDnsQuery sets DnsQuery field to given value.


### GetMetadata

`func (o *GetQueryMetricsResult) GetMetadata() MetricsMetaData`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetQueryMetricsResult) GetMetadataOk() (*MetricsMetaData, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetQueryMetricsResult) SetMetadata(v MetricsMetaData)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


