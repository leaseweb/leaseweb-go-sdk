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

// checks if the Port type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Port{}

// Port struct for Port
type Port struct {
	Name *string `json:"name,omitempty"`
	Port *string `json:"port,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Port Port

// NewPort instantiates a new Port object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPort() *Port {
	this := Port{}
	return &this
}

// NewPortWithDefaults instantiates a new Port object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortWithDefaults() *Port {
	this := Port{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Port) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Port) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Port) SetName(v string) {
	o.Name = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Port) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Port) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *Port) SetPort(v string) {
	o.Port = &v
}

func (o Port) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Port) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Port) UnmarshalJSON(data []byte) (err error) {
	varPort := _Port{}

	err = json.Unmarshal(data, &varPort)

	if err != nil {
		return err
	}

	*o = Port(varPort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "port")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePort struct {
	value *Port
	isSet bool
}

func (v NullablePort) Get() *Port {
	return v.value
}

func (v *NullablePort) Set(val *Port) {
	v.value = val
	v.isSet = true
}

func (v NullablePort) IsSet() bool {
	return v.isSet
}

func (v *NullablePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePort(val *Port) *NullablePort {
	return &NullablePort{value: val, isSet: true}
}

func (v NullablePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


