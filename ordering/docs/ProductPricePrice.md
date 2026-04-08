# ProductPricePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Price&#39;s currency | [optional] 
**BasePrice** | Pointer to **float32** | Initial price | [optional] 
**Tax** | Pointer to **float32** | Tax | [optional] 
**SetupFee** | Pointer to **float32** | Setup Fee | [optional] 
**Fee** | Pointer to **float32** | Fee amount | [optional] 
**Discounts** | Pointer to [**ProductDiscounts**](ProductDiscounts.md) |  | [optional] 
**Total** | Pointer to **float32** | Final price | [optional] 
**ContractTerm** | Pointer to **string** | Selected contract term | [optional] 
**ContractTerms** | Pointer to [**[]ContractTermItem**](ContractTermItem.md) | Available contract terms | [optional] 
**BillingCycle** | Pointer to **string** | Selected billing cycle | [optional] 
**BillingCycles** | Pointer to [**[]BillingCycleItem**](BillingCycleItem.md) | Available billing cycles | [optional] 

## Methods

### NewProductPricePrice

`func NewProductPricePrice() *ProductPricePrice`

NewProductPricePrice instantiates a new ProductPricePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductPricePriceWithDefaults

`func NewProductPricePriceWithDefaults() *ProductPricePrice`

NewProductPricePriceWithDefaults instantiates a new ProductPricePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *ProductPricePrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductPricePrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductPricePrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ProductPricePrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBasePrice

`func (o *ProductPricePrice) GetBasePrice() float32`

GetBasePrice returns the BasePrice field if non-nil, zero value otherwise.

### GetBasePriceOk

`func (o *ProductPricePrice) GetBasePriceOk() (*float32, bool)`

GetBasePriceOk returns a tuple with the BasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePrice

`func (o *ProductPricePrice) SetBasePrice(v float32)`

SetBasePrice sets BasePrice field to given value.

### HasBasePrice

`func (o *ProductPricePrice) HasBasePrice() bool`

HasBasePrice returns a boolean if a field has been set.

### GetTax

`func (o *ProductPricePrice) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *ProductPricePrice) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *ProductPricePrice) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *ProductPricePrice) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetSetupFee

`func (o *ProductPricePrice) GetSetupFee() float32`

GetSetupFee returns the SetupFee field if non-nil, zero value otherwise.

### GetSetupFeeOk

`func (o *ProductPricePrice) GetSetupFeeOk() (*float32, bool)`

GetSetupFeeOk returns a tuple with the SetupFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupFee

`func (o *ProductPricePrice) SetSetupFee(v float32)`

SetSetupFee sets SetupFee field to given value.

### HasSetupFee

`func (o *ProductPricePrice) HasSetupFee() bool`

HasSetupFee returns a boolean if a field has been set.

### GetFee

`func (o *ProductPricePrice) GetFee() float32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ProductPricePrice) GetFeeOk() (*float32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ProductPricePrice) SetFee(v float32)`

SetFee sets Fee field to given value.

### HasFee

`func (o *ProductPricePrice) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetDiscounts

`func (o *ProductPricePrice) GetDiscounts() ProductDiscounts`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *ProductPricePrice) GetDiscountsOk() (*ProductDiscounts, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *ProductPricePrice) SetDiscounts(v ProductDiscounts)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *ProductPricePrice) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetTotal

`func (o *ProductPricePrice) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProductPricePrice) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProductPricePrice) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProductPricePrice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetContractTerm

`func (o *ProductPricePrice) GetContractTerm() string`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *ProductPricePrice) GetContractTermOk() (*string, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *ProductPricePrice) SetContractTerm(v string)`

SetContractTerm sets ContractTerm field to given value.

### HasContractTerm

`func (o *ProductPricePrice) HasContractTerm() bool`

HasContractTerm returns a boolean if a field has been set.

### GetContractTerms

`func (o *ProductPricePrice) GetContractTerms() []ContractTermItem`

GetContractTerms returns the ContractTerms field if non-nil, zero value otherwise.

### GetContractTermsOk

`func (o *ProductPricePrice) GetContractTermsOk() (*[]ContractTermItem, bool)`

GetContractTermsOk returns a tuple with the ContractTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerms

`func (o *ProductPricePrice) SetContractTerms(v []ContractTermItem)`

SetContractTerms sets ContractTerms field to given value.

### HasContractTerms

`func (o *ProductPricePrice) HasContractTerms() bool`

HasContractTerms returns a boolean if a field has been set.

### GetBillingCycle

`func (o *ProductPricePrice) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *ProductPricePrice) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *ProductPricePrice) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *ProductPricePrice) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetBillingCycles

`func (o *ProductPricePrice) GetBillingCycles() []BillingCycleItem`

GetBillingCycles returns the BillingCycles field if non-nil, zero value otherwise.

### GetBillingCyclesOk

`func (o *ProductPricePrice) GetBillingCyclesOk() (*[]BillingCycleItem, bool)`

GetBillingCyclesOk returns a tuple with the BillingCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycles

`func (o *ProductPricePrice) SetBillingCycles(v []BillingCycleItem)`

SetBillingCycles sets BillingCycles field to given value.

### HasBillingCycles

`func (o *ProductPricePrice) HasBillingCycles() bool`

HasBillingCycles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


