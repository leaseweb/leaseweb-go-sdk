# VpsConfigurationOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Option name | [optional] 
**Selected** | Pointer to **bool** | Whether this option is selected | [optional] 
**Price** | Pointer to **float32** | Price for this option | [optional] 
**Currency** | Pointer to **string** | Currency for the price | [optional] 

## Methods

### NewVpsConfigurationOption

`func NewVpsConfigurationOption() *VpsConfigurationOption`

NewVpsConfigurationOption instantiates a new VpsConfigurationOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsConfigurationOptionWithDefaults

`func NewVpsConfigurationOptionWithDefaults() *VpsConfigurationOption`

NewVpsConfigurationOptionWithDefaults instantiates a new VpsConfigurationOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VpsConfigurationOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpsConfigurationOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpsConfigurationOption) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpsConfigurationOption) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelected

`func (o *VpsConfigurationOption) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *VpsConfigurationOption) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *VpsConfigurationOption) SetSelected(v bool)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *VpsConfigurationOption) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### GetPrice

`func (o *VpsConfigurationOption) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *VpsConfigurationOption) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *VpsConfigurationOption) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *VpsConfigurationOption) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *VpsConfigurationOption) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VpsConfigurationOption) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VpsConfigurationOption) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VpsConfigurationOption) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


