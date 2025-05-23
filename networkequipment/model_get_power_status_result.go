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

// checks if the GetPowerStatusResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPowerStatusResult{}

// GetPowerStatusResult struct for GetPowerStatusResult
type GetPowerStatusResult struct {
	Pdu *Pdu `json:"pdu,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPowerStatusResult GetPowerStatusResult

// NewGetPowerStatusResult instantiates a new GetPowerStatusResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPowerStatusResult() *GetPowerStatusResult {
	this := GetPowerStatusResult{}
	return &this
}

// NewGetPowerStatusResultWithDefaults instantiates a new GetPowerStatusResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPowerStatusResultWithDefaults() *GetPowerStatusResult {
	this := GetPowerStatusResult{}
	return &this
}

// GetPdu returns the Pdu field value if set, zero value otherwise.
func (o *GetPowerStatusResult) GetPdu() Pdu {
	if o == nil || IsNil(o.Pdu) {
		var ret Pdu
		return ret
	}
	return *o.Pdu
}

// GetPduOk returns a tuple with the Pdu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPowerStatusResult) GetPduOk() (*Pdu, bool) {
	if o == nil || IsNil(o.Pdu) {
		return nil, false
	}
	return o.Pdu, true
}

// HasPdu returns a boolean if a field has been set.
func (o *GetPowerStatusResult) HasPdu() bool {
	if o != nil && !IsNil(o.Pdu) {
		return true
	}

	return false
}

// SetPdu gets a reference to the given Pdu and assigns it to the Pdu field.
func (o *GetPowerStatusResult) SetPdu(v Pdu) {
	o.Pdu = &v
}

func (o GetPowerStatusResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPowerStatusResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pdu) {
		toSerialize["pdu"] = o.Pdu
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPowerStatusResult) UnmarshalJSON(data []byte) (err error) {
	varGetPowerStatusResult := _GetPowerStatusResult{}

	err = json.Unmarshal(data, &varGetPowerStatusResult)

	if err != nil {
		return err
	}

	*o = GetPowerStatusResult(varGetPowerStatusResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pdu")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPowerStatusResult struct {
	value *GetPowerStatusResult
	isSet bool
}

func (v NullableGetPowerStatusResult) Get() *GetPowerStatusResult {
	return v.value
}

func (v *NullableGetPowerStatusResult) Set(val *GetPowerStatusResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPowerStatusResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPowerStatusResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPowerStatusResult(val *GetPowerStatusResult) *NullableGetPowerStatusResult {
	return &NullableGetPowerStatusResult{value: val, isSet: true}
}

func (v NullableGetPowerStatusResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPowerStatusResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


