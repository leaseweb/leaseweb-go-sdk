/*
Public Clouds

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"time"
)

// checks if the Peak type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Peak{}

// Peak struct for Peak
type Peak struct {
	// The highest amount of incoming traffic given the provided aggregation and granularity
	Value *float32 `json:"value,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Peak Peak

// NewPeak instantiates a new Peak object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeak() *Peak {
	this := Peak{}
	return &this
}

// NewPeakWithDefaults instantiates a new Peak object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeakWithDefaults() *Peak {
	this := Peak{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Peak) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Peak) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Peak) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *Peak) SetValue(v float32) {
	o.Value = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Peak) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Peak) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Peak) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *Peak) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o Peak) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Peak) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Peak) UnmarshalJSON(data []byte) (err error) {
	varPeak := _Peak{}

	err = json.Unmarshal(data, &varPeak)

	if err != nil {
		return err
	}

	*o = Peak(varPeak)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePeak struct {
	value *Peak
	isSet bool
}

func (v NullablePeak) Get() *Peak {
	return v.value
}

func (v *NullablePeak) Set(val *Peak) {
	v.value = val
	v.isSet = true
}

func (v NullablePeak) IsSet() bool {
	return v.isSet
}

func (v *NullablePeak) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeak(val *Peak) *NullablePeak {
	return &NullablePeak{value: val, isSet: true}
}

func (v NullablePeak) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeak) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


