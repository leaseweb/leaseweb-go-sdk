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

// HealthCheckStatus Health check status
type HealthCheckStatus string

// List of healthCheckStatus
const (
	HEALTHCHECKSTATUS_HEALTHY HealthCheckStatus = "HEALTHY"
	HEALTHCHECKSTATUS_UNHEALTHY HealthCheckStatus = "UNHEALTHY"
	HEALTHCHECKSTATUS_MAINTENANCE HealthCheckStatus = "MAINTENANCE"
	HEALTHCHECKSTATUS_UNKNOWN HealthCheckStatus = "UNKNOWN"
)

// All allowed values of HealthCheckStatus enum
var AllowedHealthCheckStatusEnumValues = []HealthCheckStatus{
	"HEALTHY",
	"UNHEALTHY",
	"MAINTENANCE",
	"UNKNOWN",
}

func (v *HealthCheckStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HealthCheckStatus(value)
	for _, existing := range AllowedHealthCheckStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HealthCheckStatus", value)
}

// NewHealthCheckStatusFromValue returns a pointer to a valid HealthCheckStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHealthCheckStatusFromValue(v string) (*HealthCheckStatus, error) {
	ev := HealthCheckStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HealthCheckStatus: valid values are %v", v, AllowedHealthCheckStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HealthCheckStatus) IsValid() bool {
	for _, existing := range AllowedHealthCheckStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to healthCheckStatus value
func (v HealthCheckStatus) Ptr() *HealthCheckStatus {
	return &v
}

type NullableHealthCheckStatus struct {
	value *HealthCheckStatus
	isSet bool
}

func (v NullableHealthCheckStatus) Get() *HealthCheckStatus {
	return v.value
}

func (v *NullableHealthCheckStatus) Set(val *HealthCheckStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckStatus(val *HealthCheckStatus) *NullableHealthCheckStatus {
	return &NullableHealthCheckStatus{value: val, isSet: true}
}

func (v NullableHealthCheckStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

