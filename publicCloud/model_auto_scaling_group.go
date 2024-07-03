/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the AutoScalingGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoScalingGroup{}

// AutoScalingGroup struct for AutoScalingGroup
type AutoScalingGroup struct {
	// The Auto Scaling Group unique identifier
	Id string `json:"id"`
	// Auto Scaling Group type
	Type string `json:"type"`
	// The Auto Scaling Group's current state.
	State string `json:"state"`
	// Number of instances that should be running
	DesiredAmount NullableInt32 `json:"desiredAmount"`
	// The region in which the Auto Scaling Group was launched
	Region string `json:"region"`
	// The identifying name set to the auto scaling group
	Reference string `json:"reference"`
	// Date and time when the Auto Scaling Group was created
	CreatedAt time.Time `json:"createdAt"`
	// Date and time when the Auto Scaling Group was last updated
	UpdatedAt time.Time `json:"updatedAt"`
	// Only for \"SCHEDULED\" auto scaling group. Date and time (UTC) that the instances need to be launched
	StartsAt NullableTime `json:"startsAt"`
	// Only for \"SCHEDULED\" auto scaling group. Date and time (UTC) that the instances need to be terminated
	EndsAt NullableTime `json:"endsAt"`
	// The minimum number of instances that should be running
	MinimumAmount NullableInt32 `json:"minimumAmount"`
	// Only for \"CPU_BASED\" auto scaling group. The maximum number of instances that can be running
	MaximumAmount NullableInt32 `json:"maximumAmount"`
	// Only for \"CPU_BASED\" auto scaling group. The target average CPU utilization for scaling
	CpuThreshold NullableInt32 `json:"cpuThreshold"`
	// Only for \"CPU_BASED\" auto scaling group. Warm-up time in seconds for new instances
	WarmupTime NullableInt32 `json:"warmupTime"`
	// Only for \"CPU_BASED\" auto scaling group. Cool-down time in seconds for new instances
	CooldownTime NullableInt32 `json:"cooldownTime"`
}

type _AutoScalingGroup AutoScalingGroup

// NewAutoScalingGroup instantiates a new AutoScalingGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoScalingGroup(id string, type_ string, state string, desiredAmount NullableInt32, region string, reference string, createdAt time.Time, updatedAt time.Time, startsAt NullableTime, endsAt NullableTime, minimumAmount NullableInt32, maximumAmount NullableInt32, cpuThreshold NullableInt32, warmupTime NullableInt32, cooldownTime NullableInt32) *AutoScalingGroup {
	this := AutoScalingGroup{}
	this.Id = id
	this.Type = type_
	this.State = state
	this.DesiredAmount = desiredAmount
	this.Region = region
	this.Reference = reference
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.StartsAt = startsAt
	this.EndsAt = endsAt
	this.MinimumAmount = minimumAmount
	this.MaximumAmount = maximumAmount
	this.CpuThreshold = cpuThreshold
	this.WarmupTime = warmupTime
	this.CooldownTime = cooldownTime
	return &this
}

// NewAutoScalingGroupWithDefaults instantiates a new AutoScalingGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoScalingGroupWithDefaults() *AutoScalingGroup {
	this := AutoScalingGroup{}
	return &this
}

// GetId returns the Id field value
func (o *AutoScalingGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutoScalingGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutoScalingGroup) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *AutoScalingGroup) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AutoScalingGroup) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AutoScalingGroup) SetType(v string) {
	o.Type = v
}

// GetState returns the State field value
func (o *AutoScalingGroup) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *AutoScalingGroup) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *AutoScalingGroup) SetState(v string) {
	o.State = v
}

// GetDesiredAmount returns the DesiredAmount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AutoScalingGroup) GetDesiredAmount() int32 {
	if o == nil || o.DesiredAmount.Get() == nil {
		var ret int32
		return ret
	}

	return *o.DesiredAmount.Get()
}

// GetDesiredAmountOk returns a tuple with the DesiredAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoScalingGroup) GetDesiredAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DesiredAmount.Get(), o.DesiredAmount.IsSet()
}

// SetDesiredAmount sets field value
func (o *AutoScalingGroup) SetDesiredAmount(v int32) {
	o.DesiredAmount.Set(&v)
}

// GetRegion returns the Region field value
func (o *AutoScalingGroup) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AutoScalingGroup) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AutoScalingGroup) SetRegion(v string) {
	o.Region = v
}

// GetReference returns the Reference field value
func (o *AutoScalingGroup) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *AutoScalingGroup) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *AutoScalingGroup) SetReference(v string) {
	o.Reference = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AutoScalingGroup) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AutoScalingGroup) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AutoScalingGroup) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AutoScalingGroup) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AutoScalingGroup) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AutoScalingGroup) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetStartsAt returns the StartsAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AutoScalingGroup) GetStartsAt() time.Time {
	if o == nil || o.StartsAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StartsAt.Get()
}

// GetStartsAtOk returns a tuple with the StartsAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoScalingGroup) GetStartsAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartsAt.Get(), o.StartsAt.IsSet()
}

// SetStartsAt sets field value
func (o *AutoScalingGroup) SetStartsAt(v time.Time) {
	o.StartsAt.Set(&v)
}

// GetEndsAt returns the EndsAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AutoScalingGroup) GetEndsAt() time.Time {
	if o == nil || o.EndsAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.EndsAt.Get()
}

// GetEndsAtOk returns a tuple with the EndsAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoScalingGroup) GetEndsAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndsAt.Get(), o.EndsAt.IsSet()
}

// SetEndsAt sets field value
func (o *AutoScalingGroup) SetEndsAt(v time.Time) {
	o.EndsAt.Set(&v)
}

// GetMinimumAmount returns the MinimumAmount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AutoScalingGroup) GetMinimumAmount() int32 {
	if o == nil || o.MinimumAmount.Get() == nil {
		var ret int32
		return ret
	}

	return *o.MinimumAmount.Get()
}

// GetMinimumAmountOk returns a tuple with the MinimumAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoScalingGroup) GetMinimumAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinimumAmount.Get(), o.MinimumAmount.IsSet()
}

// SetMinimumAmount sets field value
func (o *AutoScalingGroup) SetMinimumAmount(v int32) {
	o.MinimumAmount.Set(&v)
}

// GetMaximumAmount returns the MaximumAmount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AutoScalingGroup) GetMaximumAmount() int32 {
	if o == nil || o.MaximumAmount.Get() == nil {
		var ret int32
		return ret
	}

	return *o.MaximumAmount.Get()
}

// GetMaximumAmountOk returns a tuple with the MaximumAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoScalingGroup) GetMaximumAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaximumAmount.Get(), o.MaximumAmount.IsSet()
}

// SetMaximumAmount sets field value
func (o *AutoScalingGroup) SetMaximumAmount(v int32) {
	o.MaximumAmount.Set(&v)
}

// GetCpuThreshold returns the CpuThreshold field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AutoScalingGroup) GetCpuThreshold() int32 {
	if o == nil || o.CpuThreshold.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CpuThreshold.Get()
}

// GetCpuThresholdOk returns a tuple with the CpuThreshold field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoScalingGroup) GetCpuThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuThreshold.Get(), o.CpuThreshold.IsSet()
}

// SetCpuThreshold sets field value
func (o *AutoScalingGroup) SetCpuThreshold(v int32) {
	o.CpuThreshold.Set(&v)
}

// GetWarmupTime returns the WarmupTime field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AutoScalingGroup) GetWarmupTime() int32 {
	if o == nil || o.WarmupTime.Get() == nil {
		var ret int32
		return ret
	}

	return *o.WarmupTime.Get()
}

// GetWarmupTimeOk returns a tuple with the WarmupTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoScalingGroup) GetWarmupTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarmupTime.Get(), o.WarmupTime.IsSet()
}

// SetWarmupTime sets field value
func (o *AutoScalingGroup) SetWarmupTime(v int32) {
	o.WarmupTime.Set(&v)
}

// GetCooldownTime returns the CooldownTime field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AutoScalingGroup) GetCooldownTime() int32 {
	if o == nil || o.CooldownTime.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CooldownTime.Get()
}

// GetCooldownTimeOk returns a tuple with the CooldownTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoScalingGroup) GetCooldownTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CooldownTime.Get(), o.CooldownTime.IsSet()
}

// SetCooldownTime sets field value
func (o *AutoScalingGroup) SetCooldownTime(v int32) {
	o.CooldownTime.Set(&v)
}

func (o AutoScalingGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoScalingGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["state"] = o.State
	toSerialize["desiredAmount"] = o.DesiredAmount.Get()
	toSerialize["region"] = o.Region
	toSerialize["reference"] = o.Reference
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["startsAt"] = o.StartsAt.Get()
	toSerialize["endsAt"] = o.EndsAt.Get()
	toSerialize["minimumAmount"] = o.MinimumAmount.Get()
	toSerialize["maximumAmount"] = o.MaximumAmount.Get()
	toSerialize["cpuThreshold"] = o.CpuThreshold.Get()
	toSerialize["warmupTime"] = o.WarmupTime.Get()
	toSerialize["cooldownTime"] = o.CooldownTime.Get()
	return toSerialize, nil
}

func (o *AutoScalingGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"state",
		"desiredAmount",
		"region",
		"reference",
		"createdAt",
		"updatedAt",
		"startsAt",
		"endsAt",
		"minimumAmount",
		"maximumAmount",
		"cpuThreshold",
		"warmupTime",
		"cooldownTime",
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

	varAutoScalingGroup := _AutoScalingGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAutoScalingGroup)

	if err != nil {
		return err
	}

	*o = AutoScalingGroup(varAutoScalingGroup)

	return err
}

type NullableAutoScalingGroup struct {
	value *AutoScalingGroup
	isSet bool
}

func (v NullableAutoScalingGroup) Get() *AutoScalingGroup {
	return v.value
}

func (v *NullableAutoScalingGroup) Set(val *AutoScalingGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoScalingGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoScalingGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoScalingGroup(val *AutoScalingGroup) *NullableAutoScalingGroup {
	return &NullableAutoScalingGroup{value: val, isSet: true}
}

func (v NullableAutoScalingGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoScalingGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


