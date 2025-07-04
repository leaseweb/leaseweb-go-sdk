/*
Public Clouds

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"fmt"
)

// checks if the ReinstallResourceOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReinstallResourceOpts{}

// ReinstallResourceOpts struct for ReinstallResourceOpts
type ReinstallResourceOpts struct {
	// The Image ID
	ImageId string `json:"imageId"`
	// The Market App to be installed
	MarketAppId *string `json:"marketAppId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReinstallResourceOpts ReinstallResourceOpts

// NewReinstallResourceOpts instantiates a new ReinstallResourceOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReinstallResourceOpts(imageId string) *ReinstallResourceOpts {
	this := ReinstallResourceOpts{}
	this.ImageId = imageId
	return &this
}

// NewReinstallResourceOptsWithDefaults instantiates a new ReinstallResourceOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReinstallResourceOptsWithDefaults() *ReinstallResourceOpts {
	this := ReinstallResourceOpts{}
	return &this
}

// GetImageId returns the ImageId field value
func (o *ReinstallResourceOpts) GetImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value
// and a boolean to check if the value has been set.
func (o *ReinstallResourceOpts) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageId, true
}

// SetImageId sets field value
func (o *ReinstallResourceOpts) SetImageId(v string) {
	o.ImageId = v
}

// GetMarketAppId returns the MarketAppId field value if set, zero value otherwise.
func (o *ReinstallResourceOpts) GetMarketAppId() string {
	if o == nil || IsNil(o.MarketAppId) {
		var ret string
		return ret
	}
	return *o.MarketAppId
}

// GetMarketAppIdOk returns a tuple with the MarketAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReinstallResourceOpts) GetMarketAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketAppId) {
		return nil, false
	}
	return o.MarketAppId, true
}

// HasMarketAppId returns a boolean if a field has been set.
func (o *ReinstallResourceOpts) HasMarketAppId() bool {
	if o != nil && !IsNil(o.MarketAppId) {
		return true
	}

	return false
}

// SetMarketAppId gets a reference to the given string and assigns it to the MarketAppId field.
func (o *ReinstallResourceOpts) SetMarketAppId(v string) {
	o.MarketAppId = &v
}

func (o ReinstallResourceOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReinstallResourceOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["imageId"] = o.ImageId
	if !IsNil(o.MarketAppId) {
		toSerialize["marketAppId"] = o.MarketAppId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReinstallResourceOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"imageId",
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

	varReinstallResourceOpts := _ReinstallResourceOpts{}

	err = json.Unmarshal(data, &varReinstallResourceOpts)

	if err != nil {
		return err
	}

	*o = ReinstallResourceOpts(varReinstallResourceOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "imageId")
		delete(additionalProperties, "marketAppId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReinstallResourceOpts struct {
	value *ReinstallResourceOpts
	isSet bool
}

func (v NullableReinstallResourceOpts) Get() *ReinstallResourceOpts {
	return v.value
}

func (v *NullableReinstallResourceOpts) Set(val *ReinstallResourceOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableReinstallResourceOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableReinstallResourceOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReinstallResourceOpts(val *ReinstallResourceOpts) *NullableReinstallResourceOpts {
	return &NullableReinstallResourceOpts{value: val, isSet: true}
}

func (v NullableReinstallResourceOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReinstallResourceOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


