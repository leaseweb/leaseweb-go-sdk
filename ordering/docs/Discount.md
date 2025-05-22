# Discount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Discount description | [optional] 
**Value** | Pointer to **float32** | Discount value | [optional] 

## Methods

### NewDiscount

`func NewDiscount() *Discount`

NewDiscount instantiates a new Discount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountWithDefaults

`func NewDiscountWithDefaults() *Discount`

NewDiscountWithDefaults instantiates a new Discount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Discount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Discount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Discount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Discount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValue

`func (o *Discount) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Discount) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Discount) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Discount) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


