# LoadBalancers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancers** | Pointer to [**[]LoadBalancer**](LoadBalancer.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewLoadBalancers

`func NewLoadBalancers() *LoadBalancers`

NewLoadBalancers instantiates a new LoadBalancers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancersWithDefaults

`func NewLoadBalancersWithDefaults() *LoadBalancers`

NewLoadBalancersWithDefaults instantiates a new LoadBalancers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancers

`func (o *LoadBalancers) GetLoadBalancers() []LoadBalancer`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *LoadBalancers) GetLoadBalancersOk() (*[]LoadBalancer, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *LoadBalancers) SetLoadBalancers(v []LoadBalancer)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *LoadBalancers) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetMetadata

`func (o *LoadBalancers) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LoadBalancers) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LoadBalancers) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LoadBalancers) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


