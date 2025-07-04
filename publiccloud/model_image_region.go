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

// ImageRegion region name
type ImageRegion string

// List of imageRegion
const (
	IMAGEREGION_EU_WEST_3 ImageRegion = "eu-west-3"
	IMAGEREGION_US_EAST_1 ImageRegion = "us-east-1"
	IMAGEREGION_EU_CENTRAL_1 ImageRegion = "eu-central-1"
	IMAGEREGION_AP_SOUTHEAST_1 ImageRegion = "ap-southeast-1"
	IMAGEREGION_US_WEST_1 ImageRegion = "us-west-1"
	IMAGEREGION_EU_WEST_2 ImageRegion = "eu-west-2"
	IMAGEREGION_CA_CENTRAL_1 ImageRegion = "ca-central-1"
	IMAGEREGION_AP_NORTHEAST_1 ImageRegion = "ap-northeast-1"
)

// All allowed values of ImageRegion enum
var AllowedImageRegionEnumValues = []ImageRegion{
	"eu-west-3",
	"us-east-1",
	"eu-central-1",
	"ap-southeast-1",
	"us-west-1",
	"eu-west-2",
	"ca-central-1",
	"ap-northeast-1",
}

func (v *ImageRegion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImageRegion(value)
	for _, existing := range AllowedImageRegionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImageRegion", value)
}

// NewImageRegionFromValue returns a pointer to a valid ImageRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageRegionFromValue(v string) (*ImageRegion, error) {
	ev := ImageRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImageRegion: valid values are %v", v, AllowedImageRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageRegion) IsValid() bool {
	for _, existing := range AllowedImageRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to imageRegion value
func (v ImageRegion) Ptr() *ImageRegion {
	return &v
}

type NullableImageRegion struct {
	value *ImageRegion
	isSet bool
}

func (v NullableImageRegion) Get() *ImageRegion {
	return v.value
}

func (v *NullableImageRegion) Set(val *ImageRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableImageRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableImageRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageRegion(val *ImageRegion) *NullableImageRegion {
	return &NullableImageRegion{value: val, isSet: true}
}

func (v NullableImageRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

