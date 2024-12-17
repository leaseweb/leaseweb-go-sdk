/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
)

// checks if the GetIPListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIPListResult{}

// GetIPListResult struct for GetIPListResult
type GetIPListResult struct {
	Ips []IpDetails `json:"ips,omitempty"`
	Metadata *Metadata `json:"_metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIPListResult GetIPListResult

// NewGetIPListResult instantiates a new GetIPListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIPListResult() *GetIPListResult {
	this := GetIPListResult{}
	return &this
}

// NewGetIPListResultWithDefaults instantiates a new GetIPListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIPListResultWithDefaults() *GetIPListResult {
	this := GetIPListResult{}
	return &this
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *GetIPListResult) GetIps() []IpDetails {
	if o == nil || IsNil(o.Ips) {
		var ret []IpDetails
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIPListResult) GetIpsOk() ([]IpDetails, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *GetIPListResult) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []IpDetails and assigns it to the Ips field.
func (o *GetIPListResult) SetIps(v []IpDetails) {
	o.Ips = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetIPListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIPListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetIPListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetIPListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

func (o GetIPListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIPListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetIPListResult) UnmarshalJSON(data []byte) (err error) {
	varGetIPListResult := _GetIPListResult{}

	err = json.Unmarshal(data, &varGetIPListResult)

	if err != nil {
		return err
	}

	*o = GetIPListResult(varGetIPListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ips")
		delete(additionalProperties, "_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIPListResult struct {
	value *GetIPListResult
	isSet bool
}

func (v NullableGetIPListResult) Get() *GetIPListResult {
	return v.value
}

func (v *NullableGetIPListResult) Set(val *GetIPListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIPListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIPListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIPListResult(val *GetIPListResult) *NullableGetIPListResult {
	return &NullableGetIPListResult{value: val, isSet: true}
}

func (v NullableGetIPListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIPListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


