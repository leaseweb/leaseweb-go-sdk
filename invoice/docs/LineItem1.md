# LineItem1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | Pointer to **string** | The unique id of the contract | [optional] 
**EquipmentId** | Pointer to **string** | The unique id of the equipment | [optional] 
**StartDate** | Pointer to **string** | The start date of the contract (UTC) | [optional] 
**EndDate** | Pointer to **string** | The end date of the contract (UTC) | [optional] 
**PoNumber** | Pointer to **string** | The purchase order number. | [optional] 
**Product** | Pointer to **string** | A string to indicate what kind of product this is | [optional] 
**Reference** | Pointer to **string** | The reference a customer gave to the service | [optional] 
**Price** | Pointer to **float32** | The price of the contract item. | [optional] 
**Currency** | Pointer to **string** | The currency of the service price | [optional] 

## Methods

### NewLineItem1

`func NewLineItem1() *LineItem1`

NewLineItem1 instantiates a new LineItem1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItem1WithDefaults

`func NewLineItem1WithDefaults() *LineItem1`

NewLineItem1WithDefaults instantiates a new LineItem1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *LineItem1) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *LineItem1) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *LineItem1) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *LineItem1) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetEquipmentId

`func (o *LineItem1) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *LineItem1) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *LineItem1) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.

### HasEquipmentId

`func (o *LineItem1) HasEquipmentId() bool`

HasEquipmentId returns a boolean if a field has been set.

### GetStartDate

`func (o *LineItem1) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LineItem1) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LineItem1) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LineItem1) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *LineItem1) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *LineItem1) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *LineItem1) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *LineItem1) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPoNumber

`func (o *LineItem1) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *LineItem1) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *LineItem1) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *LineItem1) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetProduct

`func (o *LineItem1) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *LineItem1) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *LineItem1) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *LineItem1) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetReference

`func (o *LineItem1) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *LineItem1) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *LineItem1) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *LineItem1) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPrice

`func (o *LineItem1) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *LineItem1) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *LineItem1) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *LineItem1) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *LineItem1) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LineItem1) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LineItem1) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *LineItem1) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


