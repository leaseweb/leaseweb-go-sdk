# DiscountItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Discount description | [optional] 
**Value** | Pointer to **float32** | Discount value | [optional] 

## Methods

### NewDiscountItems

`func NewDiscountItems() *DiscountItems`

NewDiscountItems instantiates a new DiscountItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountItemsWithDefaults

`func NewDiscountItemsWithDefaults() *DiscountItems`

NewDiscountItemsWithDefaults instantiates a new DiscountItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DiscountItems) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscountItems) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscountItems) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscountItems) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValue

`func (o *DiscountItems) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DiscountItems) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DiscountItems) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *DiscountItems) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


