# DedicatedServerPricePriceDiscounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** | Total discounts | [optional] 
**Details** | Pointer to [**[]Discount**](Discount.md) | Discounts details | [optional] 

## Methods

### NewDedicatedServerPricePriceDiscounts

`func NewDedicatedServerPricePriceDiscounts() *DedicatedServerPricePriceDiscounts`

NewDedicatedServerPricePriceDiscounts instantiates a new DedicatedServerPricePriceDiscounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedServerPricePriceDiscountsWithDefaults

`func NewDedicatedServerPricePriceDiscountsWithDefaults() *DedicatedServerPricePriceDiscounts`

NewDedicatedServerPricePriceDiscountsWithDefaults instantiates a new DedicatedServerPricePriceDiscounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *DedicatedServerPricePriceDiscounts) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DedicatedServerPricePriceDiscounts) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DedicatedServerPricePriceDiscounts) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DedicatedServerPricePriceDiscounts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetDetails

`func (o *DedicatedServerPricePriceDiscounts) GetDetails() []Discount`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DedicatedServerPricePriceDiscounts) GetDetailsOk() (*[]Discount, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DedicatedServerPricePriceDiscounts) SetDetails(v []Discount)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DedicatedServerPricePriceDiscounts) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


