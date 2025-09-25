# LineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** | A string to indicate what kind of product this is | [optional] 
**UnitAmount** | Pointer to **float32** | The amount it cost for a single product unit. | [optional] 
**TotalAmount** | Pointer to **float32** | The total amount of the product | [optional] 
**FromDate** | Pointer to **string** | The product start date (UTC) | [optional] 
**ToDate** | Pointer to **string** | The product end date (UTC) | [optional] 
**EquipmentId** | Pointer to **string** | The id of the equipment | [optional] 
**ContractId** | Pointer to **string** | The id of the contract | [optional] 
**Quantity** | Pointer to **int32** | The quantity of a product | [optional] 

## Methods

### NewLineItem

`func NewLineItem() *LineItem`

NewLineItem instantiates a new LineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemWithDefaults

`func NewLineItemWithDefaults() *LineItem`

NewLineItemWithDefaults instantiates a new LineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *LineItem) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *LineItem) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *LineItem) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *LineItem) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetUnitAmount

`func (o *LineItem) GetUnitAmount() float32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *LineItem) GetUnitAmountOk() (*float32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *LineItem) SetUnitAmount(v float32)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmount

`func (o *LineItem) HasUnitAmount() bool`

HasUnitAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *LineItem) GetTotalAmount() float32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *LineItem) GetTotalAmountOk() (*float32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *LineItem) SetTotalAmount(v float32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *LineItem) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetFromDate

`func (o *LineItem) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *LineItem) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *LineItem) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *LineItem) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetToDate

`func (o *LineItem) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *LineItem) GetToDateOk() (*string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *LineItem) SetToDate(v string)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *LineItem) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetEquipmentId

`func (o *LineItem) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *LineItem) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *LineItem) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.

### HasEquipmentId

`func (o *LineItem) HasEquipmentId() bool`

HasEquipmentId returns a boolean if a field has been set.

### GetContractId

`func (o *LineItem) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *LineItem) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *LineItem) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *LineItem) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetQuantity

`func (o *LineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *LineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


