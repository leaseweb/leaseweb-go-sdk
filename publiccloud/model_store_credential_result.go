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

// checks if the StoreCredentialResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoreCredentialResult{}

// StoreCredentialResult struct for StoreCredentialResult
type StoreCredentialResult struct {
	Type *CredentialType `json:"type,omitempty"`
	// The provided username
	Username *string `json:"username,omitempty"`
	// The provided password
	Password *string `json:"password,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreCredentialResult StoreCredentialResult

// NewStoreCredentialResult instantiates a new StoreCredentialResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreCredentialResult() *StoreCredentialResult {
	this := StoreCredentialResult{}
	return &this
}

// NewStoreCredentialResultWithDefaults instantiates a new StoreCredentialResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreCredentialResultWithDefaults() *StoreCredentialResult {
	this := StoreCredentialResult{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StoreCredentialResult) GetType() CredentialType {
	if o == nil || IsNil(o.Type) {
		var ret CredentialType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreCredentialResult) GetTypeOk() (*CredentialType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StoreCredentialResult) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CredentialType and assigns it to the Type field.
func (o *StoreCredentialResult) SetType(v CredentialType) {
	o.Type = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *StoreCredentialResult) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreCredentialResult) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *StoreCredentialResult) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *StoreCredentialResult) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *StoreCredentialResult) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreCredentialResult) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *StoreCredentialResult) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *StoreCredentialResult) SetPassword(v string) {
	o.Password = &v
}

func (o StoreCredentialResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreCredentialResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StoreCredentialResult) UnmarshalJSON(data []byte) (err error) {
	varStoreCredentialResult := _StoreCredentialResult{}

	err = json.Unmarshal(data, &varStoreCredentialResult)

	if err != nil {
		return err
	}

	*o = StoreCredentialResult(varStoreCredentialResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreCredentialResult struct {
	value *StoreCredentialResult
	isSet bool
}

func (v NullableStoreCredentialResult) Get() *StoreCredentialResult {
	return v.value
}

func (v *NullableStoreCredentialResult) Set(val *StoreCredentialResult) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreCredentialResult) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreCredentialResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreCredentialResult(val *StoreCredentialResult) *NullableStoreCredentialResult {
	return &NullableStoreCredentialResult{value: val, isSet: true}
}

func (v NullableStoreCredentialResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreCredentialResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


