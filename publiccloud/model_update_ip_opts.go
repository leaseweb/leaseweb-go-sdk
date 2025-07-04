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

// checks if the UpdateIPOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIPOpts{}

// UpdateIPOpts struct for UpdateIPOpts
type UpdateIPOpts struct {
	ReverseLookup string `json:"reverseLookup"`
	AdditionalProperties map[string]interface{}
}

type _UpdateIPOpts UpdateIPOpts

// NewUpdateIPOpts instantiates a new UpdateIPOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIPOpts(reverseLookup string) *UpdateIPOpts {
	this := UpdateIPOpts{}
	this.ReverseLookup = reverseLookup
	return &this
}

// NewUpdateIPOptsWithDefaults instantiates a new UpdateIPOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIPOptsWithDefaults() *UpdateIPOpts {
	this := UpdateIPOpts{}
	return &this
}

// GetReverseLookup returns the ReverseLookup field value
func (o *UpdateIPOpts) GetReverseLookup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReverseLookup
}

// GetReverseLookupOk returns a tuple with the ReverseLookup field value
// and a boolean to check if the value has been set.
func (o *UpdateIPOpts) GetReverseLookupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReverseLookup, true
}

// SetReverseLookup sets field value
func (o *UpdateIPOpts) SetReverseLookup(v string) {
	o.ReverseLookup = v
}

func (o UpdateIPOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIPOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reverseLookup"] = o.ReverseLookup

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateIPOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reverseLookup",
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

	varUpdateIPOpts := _UpdateIPOpts{}

	err = json.Unmarshal(data, &varUpdateIPOpts)

	if err != nil {
		return err
	}

	*o = UpdateIPOpts(varUpdateIPOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reverseLookup")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateIPOpts struct {
	value *UpdateIPOpts
	isSet bool
}

func (v NullableUpdateIPOpts) Get() *UpdateIPOpts {
	return v.value
}

func (v *NullableUpdateIPOpts) Set(val *UpdateIPOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIPOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIPOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIPOpts(val *UpdateIPOpts) *NullableUpdateIPOpts {
	return &NullableUpdateIPOpts{value: val, isSet: true}
}

func (v NullableUpdateIPOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIPOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


