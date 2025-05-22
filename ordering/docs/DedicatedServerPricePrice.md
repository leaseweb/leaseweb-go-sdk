# DedicatedServerPricePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Price&#39;s currency | [optional] 
**BasePrice** | Pointer to **float32** | Inicial price | [optional] 
**Tax** | Pointer to **float32** | Tax | [optional] 
**SetupFee** | Pointer to **float32** | Setup Fee | [optional] 
**Fee** | Pointer to **float32** | Ram size | [optional] 
**Discounts** | Pointer to [**DedicatedServerPricePriceDiscounts**](DedicatedServerPricePriceDiscounts.md) |  | [optional] 
**Total** | Pointer to **float32** | Final price | [optional] 
**BillingCycle** | Pointer to **string** | Billing periodicity | [optional] 
**ContractTerms** | Pointer to [**[]ContractTerm**](ContractTerm.md) | Contract terms | [optional] 

## Methods

### NewDedicatedServerPricePrice

`func NewDedicatedServerPricePrice() *DedicatedServerPricePrice`

NewDedicatedServerPricePrice instantiates a new DedicatedServerPricePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedServerPricePriceWithDefaults

`func NewDedicatedServerPricePriceWithDefaults() *DedicatedServerPricePrice`

NewDedicatedServerPricePriceWithDefaults instantiates a new DedicatedServerPricePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *DedicatedServerPricePrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DedicatedServerPricePrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DedicatedServerPricePrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DedicatedServerPricePrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBasePrice

`func (o *DedicatedServerPricePrice) GetBasePrice() float32`

GetBasePrice returns the BasePrice field if non-nil, zero value otherwise.

### GetBasePriceOk

`func (o *DedicatedServerPricePrice) GetBasePriceOk() (*float32, bool)`

GetBasePriceOk returns a tuple with the BasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePrice

`func (o *DedicatedServerPricePrice) SetBasePrice(v float32)`

SetBasePrice sets BasePrice field to given value.

### HasBasePrice

`func (o *DedicatedServerPricePrice) HasBasePrice() bool`

HasBasePrice returns a boolean if a field has been set.

### GetTax

`func (o *DedicatedServerPricePrice) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *DedicatedServerPricePrice) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *DedicatedServerPricePrice) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *DedicatedServerPricePrice) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetSetupFee

`func (o *DedicatedServerPricePrice) GetSetupFee() float32`

GetSetupFee returns the SetupFee field if non-nil, zero value otherwise.

### GetSetupFeeOk

`func (o *DedicatedServerPricePrice) GetSetupFeeOk() (*float32, bool)`

GetSetupFeeOk returns a tuple with the SetupFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupFee

`func (o *DedicatedServerPricePrice) SetSetupFee(v float32)`

SetSetupFee sets SetupFee field to given value.

### HasSetupFee

`func (o *DedicatedServerPricePrice) HasSetupFee() bool`

HasSetupFee returns a boolean if a field has been set.

### GetFee

`func (o *DedicatedServerPricePrice) GetFee() float32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *DedicatedServerPricePrice) GetFeeOk() (*float32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *DedicatedServerPricePrice) SetFee(v float32)`

SetFee sets Fee field to given value.

### HasFee

`func (o *DedicatedServerPricePrice) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetDiscounts

`func (o *DedicatedServerPricePrice) GetDiscounts() DedicatedServerPricePriceDiscounts`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *DedicatedServerPricePrice) GetDiscountsOk() (*DedicatedServerPricePriceDiscounts, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *DedicatedServerPricePrice) SetDiscounts(v DedicatedServerPricePriceDiscounts)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *DedicatedServerPricePrice) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetTotal

`func (o *DedicatedServerPricePrice) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DedicatedServerPricePrice) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DedicatedServerPricePrice) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DedicatedServerPricePrice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetBillingCycle

`func (o *DedicatedServerPricePrice) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *DedicatedServerPricePrice) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *DedicatedServerPricePrice) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *DedicatedServerPricePrice) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetContractTerms

`func (o *DedicatedServerPricePrice) GetContractTerms() []ContractTerm`

GetContractTerms returns the ContractTerms field if non-nil, zero value otherwise.

### GetContractTermsOk

`func (o *DedicatedServerPricePrice) GetContractTermsOk() (*[]ContractTerm, bool)`

GetContractTermsOk returns a tuple with the ContractTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerms

`func (o *DedicatedServerPricePrice) SetContractTerms(v []ContractTerm)`

SetContractTerms sets ContractTerms field to given value.

### HasContractTerms

`func (o *DedicatedServerPricePrice) HasContractTerms() bool`

HasContractTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


