# GetRequestsMetricsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**MetricsPropertiesRequests**](MetricsPropertiesRequests.md) |  | [optional] 
**Metadata** | Pointer to [**MetricsMetadataProperties**](MetricsMetadataProperties.md) |  | [optional] 

## Methods

### NewGetRequestsMetricsResult

`func NewGetRequestsMetricsResult() *GetRequestsMetricsResult`

NewGetRequestsMetricsResult instantiates a new GetRequestsMetricsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRequestsMetricsResultWithDefaults

`func NewGetRequestsMetricsResultWithDefaults() *GetRequestsMetricsResult`

NewGetRequestsMetricsResultWithDefaults instantiates a new GetRequestsMetricsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *GetRequestsMetricsResult) GetMetrics() MetricsPropertiesRequests`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetRequestsMetricsResult) GetMetricsOk() (*MetricsPropertiesRequests, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetRequestsMetricsResult) SetMetrics(v MetricsPropertiesRequests)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetRequestsMetricsResult) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetadata

`func (o *GetRequestsMetricsResult) GetMetadata() MetricsMetadataProperties`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetRequestsMetricsResult) GetMetadataOk() (*MetricsMetadataProperties, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetRequestsMetricsResult) SetMetadata(v MetricsMetadataProperties)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetRequestsMetricsResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


