/*
Abuse Reports

This API provides ways to manage the abuse reports you might receive from Leaseweb. To use this API, please request access via your account manager and/or compliance officer. **LIMITED ACCESS** 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package abuse

import (
	"encoding/json"
)

// checks if the GetReportMessageListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReportMessageListResult{}

// GetReportMessageListResult struct for GetReportMessageListResult
type GetReportMessageListResult struct {
	// An array of the posted messages.
	Messages []Message `json:"messages,omitempty"`
	Metadata *Metadata `json:"_metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetReportMessageListResult GetReportMessageListResult

// NewGetReportMessageListResult instantiates a new GetReportMessageListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReportMessageListResult() *GetReportMessageListResult {
	this := GetReportMessageListResult{}
	return &this
}

// NewGetReportMessageListResultWithDefaults instantiates a new GetReportMessageListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReportMessageListResultWithDefaults() *GetReportMessageListResult {
	this := GetReportMessageListResult{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *GetReportMessageListResult) GetMessages() []Message {
	if o == nil || IsNil(o.Messages) {
		var ret []Message
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportMessageListResult) GetMessagesOk() ([]Message, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *GetReportMessageListResult) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []Message and assigns it to the Messages field.
func (o *GetReportMessageListResult) SetMessages(v []Message) {
	o.Messages = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetReportMessageListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportMessageListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetReportMessageListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetReportMessageListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

func (o GetReportMessageListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReportMessageListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetReportMessageListResult) UnmarshalJSON(data []byte) (err error) {
	varGetReportMessageListResult := _GetReportMessageListResult{}

	err = json.Unmarshal(data, &varGetReportMessageListResult)

	if err != nil {
		return err
	}

	*o = GetReportMessageListResult(varGetReportMessageListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "messages")
		delete(additionalProperties, "_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetReportMessageListResult struct {
	value *GetReportMessageListResult
	isSet bool
}

func (v NullableGetReportMessageListResult) Get() *GetReportMessageListResult {
	return v.value
}

func (v *NullableGetReportMessageListResult) Set(val *GetReportMessageListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReportMessageListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReportMessageListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReportMessageListResult(val *GetReportMessageListResult) *NullableGetReportMessageListResult {
	return &NullableGetReportMessageListResult{value: val, isSet: true}
}

func (v NullableGetReportMessageListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReportMessageListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


