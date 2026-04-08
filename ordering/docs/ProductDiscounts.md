# ProductDiscounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** | Total discounts | [optional] 
**Details** | Pointer to [**[]DiscountItems**](DiscountItems.md) | Discount details | [optional] 

## Methods

### NewProductDiscounts

`func NewProductDiscounts() *ProductDiscounts`

NewProductDiscounts instantiates a new ProductDiscounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDiscountsWithDefaults

`func NewProductDiscountsWithDefaults() *ProductDiscounts`

NewProductDiscountsWithDefaults instantiates a new ProductDiscounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ProductDiscounts) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProductDiscounts) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProductDiscounts) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProductDiscounts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetDetails

`func (o *ProductDiscounts) GetDetails() []DiscountItems`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ProductDiscounts) GetDetailsOk() (*[]DiscountItems, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ProductDiscounts) SetDetails(v []DiscountItems)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ProductDiscounts) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


