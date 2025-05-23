/*
Ordering

This document outlines the ordering API. The API is only available for customers with invoice payment method.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ordering

import (
	"encoding/json"
	"fmt"
)

// checks if the OrderDedicatedServerOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDedicatedServerOpts{}

// OrderDedicatedServerOpts struct for OrderDedicatedServerOpts
type OrderDedicatedServerOpts struct {
	// The location of the server
	Location string `json:"location"`
	// Whether the server is connected to an aggregation pool
	ConnectedToAggregationPool *bool `json:"connectedToAggregationPool,omitempty"`
	// The contract term of the server
	ContractTerm *string `json:"contractTerm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderDedicatedServerOpts OrderDedicatedServerOpts

// NewOrderDedicatedServerOpts instantiates a new OrderDedicatedServerOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDedicatedServerOpts(location string) *OrderDedicatedServerOpts {
	this := OrderDedicatedServerOpts{}
	this.Location = location
	var connectedToAggregationPool bool = false
	this.ConnectedToAggregationPool = &connectedToAggregationPool
	var contractTerm string = "1_MONTH"
	this.ContractTerm = &contractTerm
	return &this
}

// NewOrderDedicatedServerOptsWithDefaults instantiates a new OrderDedicatedServerOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDedicatedServerOptsWithDefaults() *OrderDedicatedServerOpts {
	this := OrderDedicatedServerOpts{}
	var connectedToAggregationPool bool = false
	this.ConnectedToAggregationPool = &connectedToAggregationPool
	var contractTerm string = "1_MONTH"
	this.ContractTerm = &contractTerm
	return &this
}

// GetLocation returns the Location field value
func (o *OrderDedicatedServerOpts) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *OrderDedicatedServerOpts) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *OrderDedicatedServerOpts) SetLocation(v string) {
	o.Location = v
}

// GetConnectedToAggregationPool returns the ConnectedToAggregationPool field value if set, zero value otherwise.
func (o *OrderDedicatedServerOpts) GetConnectedToAggregationPool() bool {
	if o == nil || IsNil(o.ConnectedToAggregationPool) {
		var ret bool
		return ret
	}
	return *o.ConnectedToAggregationPool
}

// GetConnectedToAggregationPoolOk returns a tuple with the ConnectedToAggregationPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDedicatedServerOpts) GetConnectedToAggregationPoolOk() (*bool, bool) {
	if o == nil || IsNil(o.ConnectedToAggregationPool) {
		return nil, false
	}
	return o.ConnectedToAggregationPool, true
}

// HasConnectedToAggregationPool returns a boolean if a field has been set.
func (o *OrderDedicatedServerOpts) HasConnectedToAggregationPool() bool {
	if o != nil && !IsNil(o.ConnectedToAggregationPool) {
		return true
	}

	return false
}

// SetConnectedToAggregationPool gets a reference to the given bool and assigns it to the ConnectedToAggregationPool field.
func (o *OrderDedicatedServerOpts) SetConnectedToAggregationPool(v bool) {
	o.ConnectedToAggregationPool = &v
}

// GetContractTerm returns the ContractTerm field value if set, zero value otherwise.
func (o *OrderDedicatedServerOpts) GetContractTerm() string {
	if o == nil || IsNil(o.ContractTerm) {
		var ret string
		return ret
	}
	return *o.ContractTerm
}

// GetContractTermOk returns a tuple with the ContractTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDedicatedServerOpts) GetContractTermOk() (*string, bool) {
	if o == nil || IsNil(o.ContractTerm) {
		return nil, false
	}
	return o.ContractTerm, true
}

// HasContractTerm returns a boolean if a field has been set.
func (o *OrderDedicatedServerOpts) HasContractTerm() bool {
	if o != nil && !IsNil(o.ContractTerm) {
		return true
	}

	return false
}

// SetContractTerm gets a reference to the given string and assigns it to the ContractTerm field.
func (o *OrderDedicatedServerOpts) SetContractTerm(v string) {
	o.ContractTerm = &v
}

func (o OrderDedicatedServerOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDedicatedServerOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
	if !IsNil(o.ConnectedToAggregationPool) {
		toSerialize["connectedToAggregationPool"] = o.ConnectedToAggregationPool
	}
	if !IsNil(o.ContractTerm) {
		toSerialize["contractTerm"] = o.ContractTerm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderDedicatedServerOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"location",
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

	varOrderDedicatedServerOpts := _OrderDedicatedServerOpts{}

	err = json.Unmarshal(data, &varOrderDedicatedServerOpts)

	if err != nil {
		return err
	}

	*o = OrderDedicatedServerOpts(varOrderDedicatedServerOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "location")
		delete(additionalProperties, "connectedToAggregationPool")
		delete(additionalProperties, "contractTerm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderDedicatedServerOpts struct {
	value *OrderDedicatedServerOpts
	isSet bool
}

func (v NullableOrderDedicatedServerOpts) Get() *OrderDedicatedServerOpts {
	return v.value
}

func (v *NullableOrderDedicatedServerOpts) Set(val *OrderDedicatedServerOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDedicatedServerOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDedicatedServerOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDedicatedServerOpts(val *OrderDedicatedServerOpts) *NullableOrderDedicatedServerOpts {
	return &NullableOrderDedicatedServerOpts{value: val, isSet: true}
}

func (v NullableOrderDedicatedServerOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDedicatedServerOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


