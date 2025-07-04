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

// checks if the AttachIsoOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachIsoOpts{}

// AttachIsoOpts struct for AttachIsoOpts
type AttachIsoOpts struct {
	// The ISO ID, obtained using the ISO endpoints
	IsoId string `json:"isoId"`
	AdditionalProperties map[string]interface{}
}

type _AttachIsoOpts AttachIsoOpts

// NewAttachIsoOpts instantiates a new AttachIsoOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachIsoOpts(isoId string) *AttachIsoOpts {
	this := AttachIsoOpts{}
	this.IsoId = isoId
	return &this
}

// NewAttachIsoOptsWithDefaults instantiates a new AttachIsoOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachIsoOptsWithDefaults() *AttachIsoOpts {
	this := AttachIsoOpts{}
	return &this
}

// GetIsoId returns the IsoId field value
func (o *AttachIsoOpts) GetIsoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoId
}

// GetIsoIdOk returns a tuple with the IsoId field value
// and a boolean to check if the value has been set.
func (o *AttachIsoOpts) GetIsoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsoId, true
}

// SetIsoId sets field value
func (o *AttachIsoOpts) SetIsoId(v string) {
	o.IsoId = v
}

func (o AttachIsoOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachIsoOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isoId"] = o.IsoId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AttachIsoOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isoId",
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

	varAttachIsoOpts := _AttachIsoOpts{}

	err = json.Unmarshal(data, &varAttachIsoOpts)

	if err != nil {
		return err
	}

	*o = AttachIsoOpts(varAttachIsoOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "isoId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAttachIsoOpts struct {
	value *AttachIsoOpts
	isSet bool
}

func (v NullableAttachIsoOpts) Get() *AttachIsoOpts {
	return v.value
}

func (v *NullableAttachIsoOpts) Set(val *AttachIsoOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachIsoOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachIsoOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachIsoOpts(val *AttachIsoOpts) *NullableAttachIsoOpts {
	return &NullableAttachIsoOpts{value: val, isSet: true}
}

func (v NullableAttachIsoOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachIsoOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


