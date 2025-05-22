# ContractTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Contract term key | [optional] 
**Description** | Pointer to **string** | Contract term description | [optional] 
**DiscountText** | Pointer to **string** | Discount in percentage | [optional] 
**DiscountValue** | Pointer to **float32** | Discount value | [optional] 
**Total** | Pointer to **float32** | Total price with discount | [optional] 

## Methods

### NewContractTerm

`func NewContractTerm() *ContractTerm`

NewContractTerm instantiates a new ContractTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractTermWithDefaults

`func NewContractTermWithDefaults() *ContractTerm`

NewContractTermWithDefaults instantiates a new ContractTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ContractTerm) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ContractTerm) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ContractTerm) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ContractTerm) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDescription

`func (o *ContractTerm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContractTerm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContractTerm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContractTerm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountText

`func (o *ContractTerm) GetDiscountText() string`

GetDiscountText returns the DiscountText field if non-nil, zero value otherwise.

### GetDiscountTextOk

`func (o *ContractTerm) GetDiscountTextOk() (*string, bool)`

GetDiscountTextOk returns a tuple with the DiscountText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountText

`func (o *ContractTerm) SetDiscountText(v string)`

SetDiscountText sets DiscountText field to given value.

### HasDiscountText

`func (o *ContractTerm) HasDiscountText() bool`

HasDiscountText returns a boolean if a field has been set.

### GetDiscountValue

`func (o *ContractTerm) GetDiscountValue() float32`

GetDiscountValue returns the DiscountValue field if non-nil, zero value otherwise.

### GetDiscountValueOk

`func (o *ContractTerm) GetDiscountValueOk() (*float32, bool)`

GetDiscountValueOk returns a tuple with the DiscountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountValue

`func (o *ContractTerm) SetDiscountValue(v float32)`

SetDiscountValue sets DiscountValue field to given value.

### HasDiscountValue

`func (o *ContractTerm) HasDiscountValue() bool`

HasDiscountValue returns a boolean if a field has been set.

### GetTotal

`func (o *ContractTerm) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ContractTerm) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ContractTerm) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ContractTerm) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


