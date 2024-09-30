/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
)

// checks if the Certificate1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Certificate1{}

// Certificate1 struct for Certificate1
type Certificate1 struct {
	// Client Private Key. Required only if protocol is HTTPS
	PrivateKey *string `json:"privateKey,omitempty"`
	// Client Certificate. Required only if protocol is HTTPS
	Certificate *string `json:"certificate,omitempty"`
	// CA certificate. Not required, but can be added if protocol is HTTPS
	Chain NullableString `json:"chain,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Certificate1 Certificate1

// NewCertificate1 instantiates a new Certificate1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificate1() *Certificate1 {
	this := Certificate1{}
	return &this
}

// NewCertificate1WithDefaults instantiates a new Certificate1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificate1WithDefaults() *Certificate1 {
	this := Certificate1{}
	return &this
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *Certificate1) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate1) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *Certificate1) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *Certificate1) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *Certificate1) GetCertificate() string {
	if o == nil || IsNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate1) GetCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *Certificate1) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *Certificate1) SetCertificate(v string) {
	o.Certificate = &v
}

// GetChain returns the Chain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Certificate1) GetChain() string {
	if o == nil || IsNil(o.Chain.Get()) {
		var ret string
		return ret
	}
	return *o.Chain.Get()
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Certificate1) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Chain.Get(), o.Chain.IsSet()
}

// HasChain returns a boolean if a field has been set.
func (o *Certificate1) HasChain() bool {
	if o != nil && o.Chain.IsSet() {
		return true
	}

	return false
}

// SetChain gets a reference to the given NullableString and assigns it to the Chain field.
func (o *Certificate1) SetChain(v string) {
	o.Chain.Set(&v)
}
// SetChainNil sets the value for Chain to be an explicit nil
func (o *Certificate1) SetChainNil() {
	o.Chain.Set(nil)
}

// UnsetChain ensures that no value is present for Chain, not even an explicit nil
func (o *Certificate1) UnsetChain() {
	o.Chain.Unset()
}

func (o Certificate1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Certificate1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrivateKey) {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if o.Chain.IsSet() {
		toSerialize["chain"] = o.Chain.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Certificate1) UnmarshalJSON(data []byte) (err error) {
	varCertificate1 := _Certificate1{}

	err = json.Unmarshal(data, &varCertificate1)

	if err != nil {
		return err
	}

	*o = Certificate1(varCertificate1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "privateKey")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "chain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificate1 struct {
	value *Certificate1
	isSet bool
}

func (v NullableCertificate1) Get() *Certificate1 {
	return v.value
}

func (v *NullableCertificate1) Set(val *Certificate1) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificate1) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificate1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificate1(val *Certificate1) *NullableCertificate1 {
	return &NullableCertificate1{value: val, isSet: true}
}

func (v NullableCertificate1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificate1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


