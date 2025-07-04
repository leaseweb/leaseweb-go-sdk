/*
Public Clouds

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"fmt"
)

// checks if the NetworkSpeed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkSpeed{}

// NetworkSpeed Network speed in Gbps
type NetworkSpeed struct {
	Value int32 `json:"value"`
	Unit string `json:"unit"`
	AdditionalProperties map[string]interface{}
}

type _NetworkSpeed NetworkSpeed

// NewNetworkSpeed instantiates a new NetworkSpeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSpeed(value int32, unit string) *NetworkSpeed {
	this := NetworkSpeed{}
	this.Value = value
	this.Unit = unit
	return &this
}

// NewNetworkSpeedWithDefaults instantiates a new NetworkSpeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSpeedWithDefaults() *NetworkSpeed {
	this := NetworkSpeed{}
	return &this
}

// GetValue returns the Value field value
func (o *NetworkSpeed) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *NetworkSpeed) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *NetworkSpeed) SetValue(v int32) {
	o.Value = v
}

// GetUnit returns the Unit field value
func (o *NetworkSpeed) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *NetworkSpeed) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *NetworkSpeed) SetUnit(v string) {
	o.Unit = v
}

func (o NetworkSpeed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkSpeed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["unit"] = o.Unit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkSpeed) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"unit",
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

	varNetworkSpeed := _NetworkSpeed{}

	err = json.Unmarshal(data, &varNetworkSpeed)

	if err != nil {
		return err
	}

	*o = NetworkSpeed(varNetworkSpeed)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "unit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkSpeed struct {
	value *NetworkSpeed
	isSet bool
}

func (v NullableNetworkSpeed) Get() *NetworkSpeed {
	return v.value
}

func (v *NullableNetworkSpeed) Set(val *NetworkSpeed) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSpeed) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSpeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSpeed(val *NetworkSpeed) *NullableNetworkSpeed {
	return &NullableNetworkSpeed{value: val, isSet: true}
}

func (v NullableNetworkSpeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSpeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


