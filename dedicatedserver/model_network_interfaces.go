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

// checks if the NetworkInterfaces type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkInterfaces{}

// NetworkInterfaces Network interface information grouped by type
type NetworkInterfaces struct {
	Public *NetworkInterface `json:"public,omitempty"`
	Internal *NetworkInterface `json:"internal,omitempty"`
	RemoteManagement *NetworkInterface `json:"remoteManagement,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkInterfaces NetworkInterfaces

// NewNetworkInterfaces instantiates a new NetworkInterfaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkInterfaces() *NetworkInterfaces {
	this := NetworkInterfaces{}
	return &this
}

// NewNetworkInterfacesWithDefaults instantiates a new NetworkInterfaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkInterfacesWithDefaults() *NetworkInterfaces {
	this := NetworkInterfaces{}
	return &this
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *NetworkInterfaces) GetPublic() NetworkInterface {
	if o == nil || IsNil(o.Public) {
		var ret NetworkInterface
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaces) GetPublicOk() (*NetworkInterface, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *NetworkInterfaces) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given NetworkInterface and assigns it to the Public field.
func (o *NetworkInterfaces) SetPublic(v NetworkInterface) {
	o.Public = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *NetworkInterfaces) GetInternal() NetworkInterface {
	if o == nil || IsNil(o.Internal) {
		var ret NetworkInterface
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaces) GetInternalOk() (*NetworkInterface, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *NetworkInterfaces) HasInternal() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given NetworkInterface and assigns it to the Internal field.
func (o *NetworkInterfaces) SetInternal(v NetworkInterface) {
	o.Internal = &v
}

// GetRemoteManagement returns the RemoteManagement field value if set, zero value otherwise.
func (o *NetworkInterfaces) GetRemoteManagement() NetworkInterface {
	if o == nil || IsNil(o.RemoteManagement) {
		var ret NetworkInterface
		return ret
	}
	return *o.RemoteManagement
}

// GetRemoteManagementOk returns a tuple with the RemoteManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaces) GetRemoteManagementOk() (*NetworkInterface, bool) {
	if o == nil || IsNil(o.RemoteManagement) {
		return nil, false
	}
	return o.RemoteManagement, true
}

// HasRemoteManagement returns a boolean if a field has been set.
func (o *NetworkInterfaces) HasRemoteManagement() bool {
	if o != nil && !IsNil(o.RemoteManagement) {
		return true
	}

	return false
}

// SetRemoteManagement gets a reference to the given NetworkInterface and assigns it to the RemoteManagement field.
func (o *NetworkInterfaces) SetRemoteManagement(v NetworkInterface) {
	o.RemoteManagement = &v
}

func (o NetworkInterfaces) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkInterfaces) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	if !IsNil(o.RemoteManagement) {
		toSerialize["remoteManagement"] = o.RemoteManagement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkInterfaces) UnmarshalJSON(data []byte) (err error) {
	varNetworkInterfaces := _NetworkInterfaces{}

	err = json.Unmarshal(data, &varNetworkInterfaces)

	if err != nil {
		return err
	}

	*o = NetworkInterfaces(varNetworkInterfaces)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "public")
		delete(additionalProperties, "internal")
		delete(additionalProperties, "remoteManagement")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkInterfaces struct {
	value *NetworkInterfaces
	isSet bool
}

func (v NullableNetworkInterfaces) Get() *NetworkInterfaces {
	return v.value
}

func (v *NullableNetworkInterfaces) Set(val *NetworkInterfaces) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkInterfaces) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkInterfaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkInterfaces(val *NetworkInterfaces) *NullableNetworkInterfaces {
	return &NullableNetworkInterfaces{value: val, isSet: true}
}

func (v NullableNetworkInterfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkInterfaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


