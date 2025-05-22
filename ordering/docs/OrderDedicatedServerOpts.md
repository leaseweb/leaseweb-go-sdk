# OrderDedicatedServerOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | The location of the server | 
**ConnectedToAggregationPool** | Pointer to **bool** | Whether the server is connected to an aggregation pool | [optional] [default to false]
**ContractTerm** | Pointer to **string** | The contract term of the server | [optional] [default to "1_MONTH"]

## Methods

### NewOrderDedicatedServerOpts

`func NewOrderDedicatedServerOpts(location string, ) *OrderDedicatedServerOpts`

NewOrderDedicatedServerOpts instantiates a new OrderDedicatedServerOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDedicatedServerOptsWithDefaults

`func NewOrderDedicatedServerOptsWithDefaults() *OrderDedicatedServerOpts`

NewOrderDedicatedServerOptsWithDefaults instantiates a new OrderDedicatedServerOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *OrderDedicatedServerOpts) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OrderDedicatedServerOpts) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OrderDedicatedServerOpts) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetConnectedToAggregationPool

`func (o *OrderDedicatedServerOpts) GetConnectedToAggregationPool() bool`

GetConnectedToAggregationPool returns the ConnectedToAggregationPool field if non-nil, zero value otherwise.

### GetConnectedToAggregationPoolOk

`func (o *OrderDedicatedServerOpts) GetConnectedToAggregationPoolOk() (*bool, bool)`

GetConnectedToAggregationPoolOk returns a tuple with the ConnectedToAggregationPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedToAggregationPool

`func (o *OrderDedicatedServerOpts) SetConnectedToAggregationPool(v bool)`

SetConnectedToAggregationPool sets ConnectedToAggregationPool field to given value.

### HasConnectedToAggregationPool

`func (o *OrderDedicatedServerOpts) HasConnectedToAggregationPool() bool`

HasConnectedToAggregationPool returns a boolean if a field has been set.

### GetContractTerm

`func (o *OrderDedicatedServerOpts) GetContractTerm() string`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *OrderDedicatedServerOpts) GetContractTermOk() (*string, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *OrderDedicatedServerOpts) SetContractTerm(v string)`

SetContractTerm sets ContractTerm field to given value.

### HasContractTerm

`func (o *OrderDedicatedServerOpts) HasContractTerm() bool`

HasContractTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


