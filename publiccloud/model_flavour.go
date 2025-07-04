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

// Flavour Standard images flavours
type Flavour string

// List of flavour
const (
	FLAVOUR_UBUNTU Flavour = "ubuntu"
	FLAVOUR_DEBIAN Flavour = "debian"
	FLAVOUR_FREEBSD Flavour = "freebsd"
	FLAVOUR_CENTOS Flavour = "centos"
	FLAVOUR_ALMALINUX Flavour = "almalinux"
	FLAVOUR_ROCKYLINUX Flavour = "rockylinux"
	FLAVOUR_ARCHLINUX Flavour = "archlinux"
	FLAVOUR_WINDOWS Flavour = "windows"
)

// All allowed values of Flavour enum
var AllowedFlavourEnumValues = []Flavour{
	"ubuntu",
	"debian",
	"freebsd",
	"centos",
	"almalinux",
	"rockylinux",
	"archlinux",
	"windows",
}

func (v *Flavour) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Flavour(value)
	for _, existing := range AllowedFlavourEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Flavour", value)
}

// NewFlavourFromValue returns a pointer to a valid Flavour
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlavourFromValue(v string) (*Flavour, error) {
	ev := Flavour(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Flavour: valid values are %v", v, AllowedFlavourEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Flavour) IsValid() bool {
	for _, existing := range AllowedFlavourEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to flavour value
func (v Flavour) Ptr() *Flavour {
	return &v
}

type NullableFlavour struct {
	value *Flavour
	isSet bool
}

func (v NullableFlavour) Get() *Flavour {
	return v.value
}

func (v *NullableFlavour) Set(val *Flavour) {
	v.value = val
	v.isSet = true
}

func (v NullableFlavour) IsSet() bool {
	return v.isSet
}

func (v *NullableFlavour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlavour(val *Flavour) *NullableFlavour {
	return &NullableFlavour{value: val, isSet: true}
}

func (v NullableFlavour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlavour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

