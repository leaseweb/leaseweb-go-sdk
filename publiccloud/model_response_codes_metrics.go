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

// checks if the ResponseCodesMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCodesMetrics{}

// ResponseCodesMetrics struct for ResponseCodesMetrics
type ResponseCodesMetrics struct {
	Var2xx *MetricsPropertiesResponses `json:"2xx,omitempty"`
	Var3xx *MetricsPropertiesResponses `json:"3xx,omitempty"`
	Var4xx *MetricsPropertiesResponses `json:"4xx,omitempty"`
	Var5xx *MetricsPropertiesResponses `json:"5xx,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseCodesMetrics ResponseCodesMetrics

// NewResponseCodesMetrics instantiates a new ResponseCodesMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCodesMetrics() *ResponseCodesMetrics {
	this := ResponseCodesMetrics{}
	return &this
}

// NewResponseCodesMetricsWithDefaults instantiates a new ResponseCodesMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCodesMetricsWithDefaults() *ResponseCodesMetrics {
	this := ResponseCodesMetrics{}
	return &this
}

// GetVar2xx returns the Var2xx field value if set, zero value otherwise.
func (o *ResponseCodesMetrics) GetVar2xx() MetricsPropertiesResponses {
	if o == nil || IsNil(o.Var2xx) {
		var ret MetricsPropertiesResponses
		return ret
	}
	return *o.Var2xx
}

// GetVar2xxOk returns a tuple with the Var2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCodesMetrics) GetVar2xxOk() (*MetricsPropertiesResponses, bool) {
	if o == nil || IsNil(o.Var2xx) {
		return nil, false
	}
	return o.Var2xx, true
}

// HasVar2xx returns a boolean if a field has been set.
func (o *ResponseCodesMetrics) HasVar2xx() bool {
	if o != nil && !IsNil(o.Var2xx) {
		return true
	}

	return false
}

// SetVar2xx gets a reference to the given MetricsPropertiesResponses and assigns it to the Var2xx field.
func (o *ResponseCodesMetrics) SetVar2xx(v MetricsPropertiesResponses) {
	o.Var2xx = &v
}

// GetVar3xx returns the Var3xx field value if set, zero value otherwise.
func (o *ResponseCodesMetrics) GetVar3xx() MetricsPropertiesResponses {
	if o == nil || IsNil(o.Var3xx) {
		var ret MetricsPropertiesResponses
		return ret
	}
	return *o.Var3xx
}

// GetVar3xxOk returns a tuple with the Var3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCodesMetrics) GetVar3xxOk() (*MetricsPropertiesResponses, bool) {
	if o == nil || IsNil(o.Var3xx) {
		return nil, false
	}
	return o.Var3xx, true
}

// HasVar3xx returns a boolean if a field has been set.
func (o *ResponseCodesMetrics) HasVar3xx() bool {
	if o != nil && !IsNil(o.Var3xx) {
		return true
	}

	return false
}

// SetVar3xx gets a reference to the given MetricsPropertiesResponses and assigns it to the Var3xx field.
func (o *ResponseCodesMetrics) SetVar3xx(v MetricsPropertiesResponses) {
	o.Var3xx = &v
}

// GetVar4xx returns the Var4xx field value if set, zero value otherwise.
func (o *ResponseCodesMetrics) GetVar4xx() MetricsPropertiesResponses {
	if o == nil || IsNil(o.Var4xx) {
		var ret MetricsPropertiesResponses
		return ret
	}
	return *o.Var4xx
}

// GetVar4xxOk returns a tuple with the Var4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCodesMetrics) GetVar4xxOk() (*MetricsPropertiesResponses, bool) {
	if o == nil || IsNil(o.Var4xx) {
		return nil, false
	}
	return o.Var4xx, true
}

// HasVar4xx returns a boolean if a field has been set.
func (o *ResponseCodesMetrics) HasVar4xx() bool {
	if o != nil && !IsNil(o.Var4xx) {
		return true
	}

	return false
}

// SetVar4xx gets a reference to the given MetricsPropertiesResponses and assigns it to the Var4xx field.
func (o *ResponseCodesMetrics) SetVar4xx(v MetricsPropertiesResponses) {
	o.Var4xx = &v
}

// GetVar5xx returns the Var5xx field value if set, zero value otherwise.
func (o *ResponseCodesMetrics) GetVar5xx() MetricsPropertiesResponses {
	if o == nil || IsNil(o.Var5xx) {
		var ret MetricsPropertiesResponses
		return ret
	}
	return *o.Var5xx
}

// GetVar5xxOk returns a tuple with the Var5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCodesMetrics) GetVar5xxOk() (*MetricsPropertiesResponses, bool) {
	if o == nil || IsNil(o.Var5xx) {
		return nil, false
	}
	return o.Var5xx, true
}

// HasVar5xx returns a boolean if a field has been set.
func (o *ResponseCodesMetrics) HasVar5xx() bool {
	if o != nil && !IsNil(o.Var5xx) {
		return true
	}

	return false
}

// SetVar5xx gets a reference to the given MetricsPropertiesResponses and assigns it to the Var5xx field.
func (o *ResponseCodesMetrics) SetVar5xx(v MetricsPropertiesResponses) {
	o.Var5xx = &v
}

func (o ResponseCodesMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCodesMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var2xx) {
		toSerialize["2xx"] = o.Var2xx
	}
	if !IsNil(o.Var3xx) {
		toSerialize["3xx"] = o.Var3xx
	}
	if !IsNil(o.Var4xx) {
		toSerialize["4xx"] = o.Var4xx
	}
	if !IsNil(o.Var5xx) {
		toSerialize["5xx"] = o.Var5xx
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseCodesMetrics) UnmarshalJSON(data []byte) (err error) {
	varResponseCodesMetrics := _ResponseCodesMetrics{}

	err = json.Unmarshal(data, &varResponseCodesMetrics)

	if err != nil {
		return err
	}

	*o = ResponseCodesMetrics(varResponseCodesMetrics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "2xx")
		delete(additionalProperties, "3xx")
		delete(additionalProperties, "4xx")
		delete(additionalProperties, "5xx")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseCodesMetrics struct {
	value *ResponseCodesMetrics
	isSet bool
}

func (v NullableResponseCodesMetrics) Get() *ResponseCodesMetrics {
	return v.value
}

func (v *NullableResponseCodesMetrics) Set(val *ResponseCodesMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCodesMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCodesMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCodesMetrics(val *ResponseCodesMetrics) *NullableResponseCodesMetrics {
	return &NullableResponseCodesMetrics{value: val, isSet: true}
}

func (v NullableResponseCodesMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCodesMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


