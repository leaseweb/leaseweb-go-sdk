/*
Dedicated Network Equipments

This is the description of the Dedicated Network Equipment API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkequipment

import (
	"encoding/json"
)

// checks if the NetworkEquipmentSpecs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkEquipmentSpecs{}

// NetworkEquipmentSpecs Hardware information of the network equipment
type NetworkEquipmentSpecs struct {
	// The brand of the network equipment
	Brand *string `json:"brand,omitempty"`
	// The model of the network equipment
	Model *string `json:"model,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkEquipmentSpecs NetworkEquipmentSpecs

// NewNetworkEquipmentSpecs instantiates a new NetworkEquipmentSpecs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkEquipmentSpecs() *NetworkEquipmentSpecs {
	this := NetworkEquipmentSpecs{}
	return &this
}

// NewNetworkEquipmentSpecsWithDefaults instantiates a new NetworkEquipmentSpecs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkEquipmentSpecsWithDefaults() *NetworkEquipmentSpecs {
	this := NetworkEquipmentSpecs{}
	return &this
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *NetworkEquipmentSpecs) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEquipmentSpecs) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *NetworkEquipmentSpecs) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *NetworkEquipmentSpecs) SetBrand(v string) {
	o.Brand = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *NetworkEquipmentSpecs) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEquipmentSpecs) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *NetworkEquipmentSpecs) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *NetworkEquipmentSpecs) SetModel(v string) {
	o.Model = &v
}

func (o NetworkEquipmentSpecs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkEquipmentSpecs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkEquipmentSpecs) UnmarshalJSON(data []byte) (err error) {
	varNetworkEquipmentSpecs := _NetworkEquipmentSpecs{}

	err = json.Unmarshal(data, &varNetworkEquipmentSpecs)

	if err != nil {
		return err
	}

	*o = NetworkEquipmentSpecs(varNetworkEquipmentSpecs)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "brand")
		delete(additionalProperties, "model")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkEquipmentSpecs struct {
	value *NetworkEquipmentSpecs
	isSet bool
}

func (v NullableNetworkEquipmentSpecs) Get() *NetworkEquipmentSpecs {
	return v.value
}

func (v *NullableNetworkEquipmentSpecs) Set(val *NetworkEquipmentSpecs) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkEquipmentSpecs) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkEquipmentSpecs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkEquipmentSpecs(val *NetworkEquipmentSpecs) *NullableNetworkEquipmentSpecs {
	return &NullableNetworkEquipmentSpecs{value: val, isSet: true}
}

func (v NullableNetworkEquipmentSpecs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkEquipmentSpecs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


