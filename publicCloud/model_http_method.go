/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"fmt"
)

// HttpMethod the model 'HttpMethod'
type HttpMethod string

// List of httpMethod
const (
	HTTPMETHOD_GET HttpMethod = "GET"
	HTTPMETHOD_HEAD HttpMethod = "HEAD"
	HTTPMETHOD_POST HttpMethod = "POST"
	HTTPMETHOD_OPTIONS HttpMethod = "OPTIONS"
)

// All allowed values of HttpMethod enum
var AllowedHttpMethodEnumValues = []HttpMethod{
	"GET",
	"HEAD",
	"POST",
	"OPTIONS",
}

func (v *HttpMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HttpMethod(value)
	for _, existing := range AllowedHttpMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HttpMethod", value)
}

// NewHttpMethodFromValue returns a pointer to a valid HttpMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHttpMethodFromValue(v string) (*HttpMethod, error) {
	ev := HttpMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HttpMethod: valid values are %v", v, AllowedHttpMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HttpMethod) IsValid() bool {
	for _, existing := range AllowedHttpMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to httpMethod value
func (v HttpMethod) Ptr() *HttpMethod {
	return &v
}

type NullableHttpMethod struct {
	value *HttpMethod
	isSet bool
}

func (v NullableHttpMethod) Get() *HttpMethod {
	return v.value
}

func (v *NullableHttpMethod) Set(val *HttpMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpMethod(val *HttpMethod) *NullableHttpMethod {
	return &NullableHttpMethod{value: val, isSet: true}
}

func (v NullableHttpMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

