/*
Dedicated Servers

This is the description of the Dedicated Server API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
	"fmt"
)

// checks if the Partition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Partition{}

// Partition struct for Partition
type Partition struct {
	// The partition mount point (eg `/home`). Mandatory for the root partition (`/`) and not intended to be used in swap partition
	Mountpoint NullableString `json:"mountpoint,omitempty"`
	// File system in which partition would be mounted
	Filesystem string `json:"filesystem"`
	// Size of the partition (Normally in MB, but this is OS-specific)
	Size string `json:"size"`
	AdditionalProperties map[string]interface{}
}

type _Partition Partition

// NewPartition instantiates a new Partition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartition(filesystem string, size string) *Partition {
	this := Partition{}
	this.Filesystem = filesystem
	this.Size = size
	return &this
}

// NewPartitionWithDefaults instantiates a new Partition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartitionWithDefaults() *Partition {
	this := Partition{}
	return &this
}

// GetMountpoint returns the Mountpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Partition) GetMountpoint() string {
	if o == nil || IsNil(o.Mountpoint.Get()) {
		var ret string
		return ret
	}
	return *o.Mountpoint.Get()
}

// GetMountpointOk returns a tuple with the Mountpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Partition) GetMountpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mountpoint.Get(), o.Mountpoint.IsSet()
}

// HasMountpoint returns a boolean if a field has been set.
func (o *Partition) HasMountpoint() bool {
	if o != nil && o.Mountpoint.IsSet() {
		return true
	}

	return false
}

// SetMountpoint gets a reference to the given NullableString and assigns it to the Mountpoint field.
func (o *Partition) SetMountpoint(v string) {
	o.Mountpoint.Set(&v)
}
// SetMountpointNil sets the value for Mountpoint to be an explicit nil
func (o *Partition) SetMountpointNil() {
	o.Mountpoint.Set(nil)
}

// UnsetMountpoint ensures that no value is present for Mountpoint, not even an explicit nil
func (o *Partition) UnsetMountpoint() {
	o.Mountpoint.Unset()
}

// GetFilesystem returns the Filesystem field value
func (o *Partition) GetFilesystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filesystem
}

// GetFilesystemOk returns a tuple with the Filesystem field value
// and a boolean to check if the value has been set.
func (o *Partition) GetFilesystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filesystem, true
}

// SetFilesystem sets field value
func (o *Partition) SetFilesystem(v string) {
	o.Filesystem = v
}

// GetSize returns the Size field value
func (o *Partition) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *Partition) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *Partition) SetSize(v string) {
	o.Size = v
}

func (o Partition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Partition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Mountpoint.IsSet() {
		toSerialize["mountpoint"] = o.Mountpoint.Get()
	}
	toSerialize["filesystem"] = o.Filesystem
	toSerialize["size"] = o.Size

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Partition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filesystem",
		"size",
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

	varPartition := _Partition{}

	err = json.Unmarshal(data, &varPartition)

	if err != nil {
		return err
	}

	*o = Partition(varPartition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mountpoint")
		delete(additionalProperties, "filesystem")
		delete(additionalProperties, "size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartition struct {
	value *Partition
	isSet bool
}

func (v NullablePartition) Get() *Partition {
	return v.value
}

func (v *NullablePartition) Set(val *Partition) {
	v.value = val
	v.isSet = true
}

func (v NullablePartition) IsSet() bool {
	return v.isSet
}

func (v *NullablePartition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartition(val *Partition) *NullablePartition {
	return &NullablePartition{value: val, isSet: true}
}

func (v NullablePartition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


