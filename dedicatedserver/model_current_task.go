/*
Dedicated Servers

This is the description of the Dedicated Server API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
)

// checks if the CurrentTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrentTask{}

// CurrentTask The current task of the job
type CurrentTask struct {
	Task *Task `json:"task,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CurrentTask CurrentTask

// NewCurrentTask instantiates a new CurrentTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentTask() *CurrentTask {
	this := CurrentTask{}
	return &this
}

// NewCurrentTaskWithDefaults instantiates a new CurrentTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentTaskWithDefaults() *CurrentTask {
	this := CurrentTask{}
	return &this
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *CurrentTask) GetTask() Task {
	if o == nil || IsNil(o.Task) {
		var ret Task
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentTask) GetTaskOk() (*Task, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *CurrentTask) HasTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given Task and assigns it to the Task field.
func (o *CurrentTask) SetTask(v Task) {
	o.Task = &v
}

func (o CurrentTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CurrentTask) UnmarshalJSON(data []byte) (err error) {
	varCurrentTask := _CurrentTask{}

	err = json.Unmarshal(data, &varCurrentTask)

	if err != nil {
		return err
	}

	*o = CurrentTask(varCurrentTask)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "task")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCurrentTask struct {
	value *CurrentTask
	isSet bool
}

func (v NullableCurrentTask) Get() *CurrentTask {
	return v.value
}

func (v *NullableCurrentTask) Set(val *CurrentTask) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentTask) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentTask(val *CurrentTask) *NullableCurrentTask {
	return &NullableCurrentTask{value: val, isSet: true}
}

func (v NullableCurrentTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


