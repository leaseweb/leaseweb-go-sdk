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

// checks if the GetCredentialListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCredentialListResult{}

// GetCredentialListResult struct for GetCredentialListResult
type GetCredentialListResult struct {
	Credentials []Credential `json:"credentials,omitempty"`
	Metadata *Metadata `json:"_metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCredentialListResult GetCredentialListResult

// NewGetCredentialListResult instantiates a new GetCredentialListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCredentialListResult() *GetCredentialListResult {
	this := GetCredentialListResult{}
	return &this
}

// NewGetCredentialListResultWithDefaults instantiates a new GetCredentialListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCredentialListResultWithDefaults() *GetCredentialListResult {
	this := GetCredentialListResult{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *GetCredentialListResult) GetCredentials() []Credential {
	if o == nil || IsNil(o.Credentials) {
		var ret []Credential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCredentialListResult) GetCredentialsOk() ([]Credential, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *GetCredentialListResult) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []Credential and assigns it to the Credentials field.
func (o *GetCredentialListResult) SetCredentials(v []Credential) {
	o.Credentials = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetCredentialListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCredentialListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetCredentialListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetCredentialListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

func (o GetCredentialListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCredentialListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCredentialListResult) UnmarshalJSON(data []byte) (err error) {
	varGetCredentialListResult := _GetCredentialListResult{}

	err = json.Unmarshal(data, &varGetCredentialListResult)

	if err != nil {
		return err
	}

	*o = GetCredentialListResult(varGetCredentialListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCredentialListResult struct {
	value *GetCredentialListResult
	isSet bool
}

func (v NullableGetCredentialListResult) Get() *GetCredentialListResult {
	return v.value
}

func (v *NullableGetCredentialListResult) Set(val *GetCredentialListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCredentialListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCredentialListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCredentialListResult(val *GetCredentialListResult) *NullableGetCredentialListResult {
	return &NullableGetCredentialListResult{value: val, isSet: true}
}

func (v NullableGetCredentialListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCredentialListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


