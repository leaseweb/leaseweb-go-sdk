# ContractTermItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Contract term key | [optional] 
**Description** | Pointer to **string** | Contract term description | [optional] 
**DiscountText** | Pointer to **string** | Discount in percentage | [optional] 
**DiscountValue** | Pointer to **float32** | Discount value | [optional] 
**Total** | Pointer to **float32** | Total price with discount | [optional] 

## Methods

### NewContractTermItem

`func NewContractTermItem() *ContractTermItem`

NewContractTermItem instantiates a new ContractTermItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractTermItemWithDefaults

`func NewContractTermItemWithDefaults() *ContractTermItem`

NewContractTermItemWithDefaults instantiates a new ContractTermItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ContractTermItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ContractTermItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ContractTermItem) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ContractTermItem) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDescription

`func (o *ContractTermItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContractTermItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContractTermItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContractTermItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountText

`func (o *ContractTermItem) GetDiscountText() string`

GetDiscountText returns the DiscountText field if non-nil, zero value otherwise.

### GetDiscountTextOk

`func (o *ContractTermItem) GetDiscountTextOk() (*string, bool)`

GetDiscountTextOk returns a tuple with the DiscountText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountText

`func (o *ContractTermItem) SetDiscountText(v string)`

SetDiscountText sets DiscountText field to given value.

### HasDiscountText

`func (o *ContractTermItem) HasDiscountText() bool`

HasDiscountText returns a boolean if a field has been set.

### GetDiscountValue

`func (o *ContractTermItem) GetDiscountValue() float32`

GetDiscountValue returns the DiscountValue field if non-nil, zero value otherwise.

### GetDiscountValueOk

`func (o *ContractTermItem) GetDiscountValueOk() (*float32, bool)`

GetDiscountValueOk returns a tuple with the DiscountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountValue

`func (o *ContractTermItem) SetDiscountValue(v float32)`

SetDiscountValue sets DiscountValue field to given value.

### HasDiscountValue

`func (o *ContractTermItem) HasDiscountValue() bool`

HasDiscountValue returns a boolean if a field has been set.

### GetTotal

`func (o *ContractTermItem) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ContractTermItem) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ContractTermItem) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ContractTermItem) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


