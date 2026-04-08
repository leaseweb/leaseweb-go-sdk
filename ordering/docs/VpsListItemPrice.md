# VpsListItemPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Price&#39;s currency | [optional] 
**BasePrice** | Pointer to **float32** | Base price | [optional] 
**Discount** | Pointer to **float32** | Discount amount | [optional] 
**Total** | Pointer to **float32** | Final price | [optional] 

## Methods

### NewVpsListItemPrice

`func NewVpsListItemPrice() *VpsListItemPrice`

NewVpsListItemPrice instantiates a new VpsListItemPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsListItemPriceWithDefaults

`func NewVpsListItemPriceWithDefaults() *VpsListItemPrice`

NewVpsListItemPriceWithDefaults instantiates a new VpsListItemPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *VpsListItemPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VpsListItemPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VpsListItemPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VpsListItemPrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBasePrice

`func (o *VpsListItemPrice) GetBasePrice() float32`

GetBasePrice returns the BasePrice field if non-nil, zero value otherwise.

### GetBasePriceOk

`func (o *VpsListItemPrice) GetBasePriceOk() (*float32, bool)`

GetBasePriceOk returns a tuple with the BasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePrice

`func (o *VpsListItemPrice) SetBasePrice(v float32)`

SetBasePrice sets BasePrice field to given value.

### HasBasePrice

`func (o *VpsListItemPrice) HasBasePrice() bool`

HasBasePrice returns a boolean if a field has been set.

### GetDiscount

`func (o *VpsListItemPrice) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *VpsListItemPrice) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *VpsListItemPrice) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *VpsListItemPrice) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetTotal

`func (o *VpsListItemPrice) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *VpsListItemPrice) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *VpsListItemPrice) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *VpsListItemPrice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


