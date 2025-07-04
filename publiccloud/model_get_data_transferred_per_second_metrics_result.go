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

// checks if the GetDataTransferredPerSecondMetricsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDataTransferredPerSecondMetricsResult{}

// GetDataTransferredPerSecondMetricsResult struct for GetDataTransferredPerSecondMetricsResult
type GetDataTransferredPerSecondMetricsResult struct {
	Metrics *DataTransferredMetrics1 `json:"metrics,omitempty"`
	Metadata *MetricsMetadataProperties `json:"_metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDataTransferredPerSecondMetricsResult GetDataTransferredPerSecondMetricsResult

// NewGetDataTransferredPerSecondMetricsResult instantiates a new GetDataTransferredPerSecondMetricsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDataTransferredPerSecondMetricsResult() *GetDataTransferredPerSecondMetricsResult {
	this := GetDataTransferredPerSecondMetricsResult{}
	return &this
}

// NewGetDataTransferredPerSecondMetricsResultWithDefaults instantiates a new GetDataTransferredPerSecondMetricsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDataTransferredPerSecondMetricsResultWithDefaults() *GetDataTransferredPerSecondMetricsResult {
	this := GetDataTransferredPerSecondMetricsResult{}
	return &this
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *GetDataTransferredPerSecondMetricsResult) GetMetrics() DataTransferredMetrics1 {
	if o == nil || IsNil(o.Metrics) {
		var ret DataTransferredMetrics1
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDataTransferredPerSecondMetricsResult) GetMetricsOk() (*DataTransferredMetrics1, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *GetDataTransferredPerSecondMetricsResult) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given DataTransferredMetrics1 and assigns it to the Metrics field.
func (o *GetDataTransferredPerSecondMetricsResult) SetMetrics(v DataTransferredMetrics1) {
	o.Metrics = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetDataTransferredPerSecondMetricsResult) GetMetadata() MetricsMetadataProperties {
	if o == nil || IsNil(o.Metadata) {
		var ret MetricsMetadataProperties
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDataTransferredPerSecondMetricsResult) GetMetadataOk() (*MetricsMetadataProperties, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetDataTransferredPerSecondMetricsResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given MetricsMetadataProperties and assigns it to the Metadata field.
func (o *GetDataTransferredPerSecondMetricsResult) SetMetadata(v MetricsMetadataProperties) {
	o.Metadata = &v
}

func (o GetDataTransferredPerSecondMetricsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDataTransferredPerSecondMetricsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDataTransferredPerSecondMetricsResult) UnmarshalJSON(data []byte) (err error) {
	varGetDataTransferredPerSecondMetricsResult := _GetDataTransferredPerSecondMetricsResult{}

	err = json.Unmarshal(data, &varGetDataTransferredPerSecondMetricsResult)

	if err != nil {
		return err
	}

	*o = GetDataTransferredPerSecondMetricsResult(varGetDataTransferredPerSecondMetricsResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "metrics")
		delete(additionalProperties, "_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDataTransferredPerSecondMetricsResult struct {
	value *GetDataTransferredPerSecondMetricsResult
	isSet bool
}

func (v NullableGetDataTransferredPerSecondMetricsResult) Get() *GetDataTransferredPerSecondMetricsResult {
	return v.value
}

func (v *NullableGetDataTransferredPerSecondMetricsResult) Set(val *GetDataTransferredPerSecondMetricsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDataTransferredPerSecondMetricsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDataTransferredPerSecondMetricsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDataTransferredPerSecondMetricsResult(val *GetDataTransferredPerSecondMetricsResult) *NullableGetDataTransferredPerSecondMetricsResult {
	return &NullableGetDataTransferredPerSecondMetricsResult{value: val, isSet: true}
}

func (v NullableGetDataTransferredPerSecondMetricsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDataTransferredPerSecondMetricsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


