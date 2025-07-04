/*
Public Clouds

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
)

// checks if the GetIsoListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIsoListResult{}

// GetIsoListResult struct for GetIsoListResult
type GetIsoListResult struct {
	Isos []Iso `json:"isos,omitempty"`
	Metadata *Metadata `json:"_metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIsoListResult GetIsoListResult

// NewGetIsoListResult instantiates a new GetIsoListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIsoListResult() *GetIsoListResult {
	this := GetIsoListResult{}
	return &this
}

// NewGetIsoListResultWithDefaults instantiates a new GetIsoListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIsoListResultWithDefaults() *GetIsoListResult {
	this := GetIsoListResult{}
	return &this
}

// GetIsos returns the Isos field value if set, zero value otherwise.
func (o *GetIsoListResult) GetIsos() []Iso {
	if o == nil || IsNil(o.Isos) {
		var ret []Iso
		return ret
	}
	return o.Isos
}

// GetIsosOk returns a tuple with the Isos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIsoListResult) GetIsosOk() ([]Iso, bool) {
	if o == nil || IsNil(o.Isos) {
		return nil, false
	}
	return o.Isos, true
}

// HasIsos returns a boolean if a field has been set.
func (o *GetIsoListResult) HasIsos() bool {
	if o != nil && !IsNil(o.Isos) {
		return true
	}

	return false
}

// SetIsos gets a reference to the given []Iso and assigns it to the Isos field.
func (o *GetIsoListResult) SetIsos(v []Iso) {
	o.Isos = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetIsoListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIsoListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetIsoListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetIsoListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

func (o GetIsoListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIsoListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Isos) {
		toSerialize["isos"] = o.Isos
	}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetIsoListResult) UnmarshalJSON(data []byte) (err error) {
	varGetIsoListResult := _GetIsoListResult{}

	err = json.Unmarshal(data, &varGetIsoListResult)

	if err != nil {
		return err
	}

	*o = GetIsoListResult(varGetIsoListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "isos")
		delete(additionalProperties, "_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIsoListResult struct {
	value *GetIsoListResult
	isSet bool
}

func (v NullableGetIsoListResult) Get() *GetIsoListResult {
	return v.value
}

func (v *NullableGetIsoListResult) Set(val *GetIsoListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIsoListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIsoListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIsoListResult(val *GetIsoListResult) *NullableGetIsoListResult {
	return &NullableGetIsoListResult{value: val, isSet: true}
}

func (v NullableGetIsoListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIsoListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


