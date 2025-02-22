/*
Public Clouds

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"fmt"
)

// checks if the Prices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Prices{}

// Prices struct for Prices
type Prices struct {
	Currency string `json:"currency"`
	CurrencySymbol string `json:"currencySymbol"`
	Compute Price `json:"compute"`
	Storage Storage `json:"storage"`
	AdditionalProperties map[string]interface{}
}

type _Prices Prices

// NewPrices instantiates a new Prices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrices(currency string, currencySymbol string, compute Price, storage Storage) *Prices {
	this := Prices{}
	this.Currency = currency
	this.CurrencySymbol = currencySymbol
	this.Compute = compute
	this.Storage = storage
	return &this
}

// NewPricesWithDefaults instantiates a new Prices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricesWithDefaults() *Prices {
	this := Prices{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *Prices) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Prices) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Prices) SetCurrency(v string) {
	o.Currency = v
}

// GetCurrencySymbol returns the CurrencySymbol field value
func (o *Prices) GetCurrencySymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value
// and a boolean to check if the value has been set.
func (o *Prices) GetCurrencySymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencySymbol, true
}

// SetCurrencySymbol sets field value
func (o *Prices) SetCurrencySymbol(v string) {
	o.CurrencySymbol = v
}

// GetCompute returns the Compute field value
func (o *Prices) GetCompute() Price {
	if o == nil {
		var ret Price
		return ret
	}

	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *Prices) GetComputeOk() (*Price, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value
func (o *Prices) SetCompute(v Price) {
	o.Compute = v
}

// GetStorage returns the Storage field value
func (o *Prices) GetStorage() Storage {
	if o == nil {
		var ret Storage
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *Prices) GetStorageOk() (*Storage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *Prices) SetStorage(v Storage) {
	o.Storage = v
}

func (o Prices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Prices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["currencySymbol"] = o.CurrencySymbol
	toSerialize["compute"] = o.Compute
	toSerialize["storage"] = o.Storage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Prices) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
		"currencySymbol",
		"compute",
		"storage",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPrices := _Prices{}

	err = json.Unmarshal(data, &varPrices)

	if err != nil {
		return err
	}

	*o = Prices(varPrices)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currency")
		delete(additionalProperties, "currencySymbol")
		delete(additionalProperties, "compute")
		delete(additionalProperties, "storage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrices struct {
	value *Prices
	isSet bool
}

func (v NullablePrices) Get() *Prices {
	return v.value
}

func (v *NullablePrices) Set(val *Prices) {
	v.value = val
	v.isSet = true
}

func (v NullablePrices) IsSet() bool {
	return v.isSet
}

func (v *NullablePrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrices(val *Prices) *NullablePrices {
	return &NullablePrices{value: val, isSet: true}
}

func (v NullablePrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


