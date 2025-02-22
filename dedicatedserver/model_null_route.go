/*
Dedicated Servers

This is the description of the Dedicated Server API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
	"time"
)

// checks if the NullRoute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NullRoute{}

// NullRoute struct for NullRoute
type NullRoute struct {
	// The time the null route was removed or will be removed.
	AutomatedUnnullingAt *time.Time `json:"automatedUnnullingAt,omitempty"`
	// An optional comment for the reason of the null route
	Comment *string `json:"comment,omitempty"`
	// The ip address that was null routed
	Ip *string `json:"ip,omitempty"`
	// The level of the null route
	NullLevel *int32 `json:"nullLevel,omitempty"`
	// The time the null route was created
	NulledAt *time.Time `json:"nulledAt,omitempty"`
	// A ticket number if available
	TicketId *string `json:"ticketId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NullRoute NullRoute

// NewNullRoute instantiates a new NullRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNullRoute() *NullRoute {
	this := NullRoute{}
	return &this
}

// NewNullRouteWithDefaults instantiates a new NullRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNullRouteWithDefaults() *NullRoute {
	this := NullRoute{}
	return &this
}

// GetAutomatedUnnullingAt returns the AutomatedUnnullingAt field value if set, zero value otherwise.
func (o *NullRoute) GetAutomatedUnnullingAt() time.Time {
	if o == nil || IsNil(o.AutomatedUnnullingAt) {
		var ret time.Time
		return ret
	}
	return *o.AutomatedUnnullingAt
}

// GetAutomatedUnnullingAtOk returns a tuple with the AutomatedUnnullingAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullRoute) GetAutomatedUnnullingAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AutomatedUnnullingAt) {
		return nil, false
	}
	return o.AutomatedUnnullingAt, true
}

// HasAutomatedUnnullingAt returns a boolean if a field has been set.
func (o *NullRoute) HasAutomatedUnnullingAt() bool {
	if o != nil && !IsNil(o.AutomatedUnnullingAt) {
		return true
	}

	return false
}

// SetAutomatedUnnullingAt gets a reference to the given time.Time and assigns it to the AutomatedUnnullingAt field.
func (o *NullRoute) SetAutomatedUnnullingAt(v time.Time) {
	o.AutomatedUnnullingAt = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NullRoute) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullRoute) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NullRoute) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NullRoute) SetComment(v string) {
	o.Comment = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *NullRoute) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullRoute) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *NullRoute) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *NullRoute) SetIp(v string) {
	o.Ip = &v
}

// GetNullLevel returns the NullLevel field value if set, zero value otherwise.
func (o *NullRoute) GetNullLevel() int32 {
	if o == nil || IsNil(o.NullLevel) {
		var ret int32
		return ret
	}
	return *o.NullLevel
}

// GetNullLevelOk returns a tuple with the NullLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullRoute) GetNullLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.NullLevel) {
		return nil, false
	}
	return o.NullLevel, true
}

// HasNullLevel returns a boolean if a field has been set.
func (o *NullRoute) HasNullLevel() bool {
	if o != nil && !IsNil(o.NullLevel) {
		return true
	}

	return false
}

// SetNullLevel gets a reference to the given int32 and assigns it to the NullLevel field.
func (o *NullRoute) SetNullLevel(v int32) {
	o.NullLevel = &v
}

// GetNulledAt returns the NulledAt field value if set, zero value otherwise.
func (o *NullRoute) GetNulledAt() time.Time {
	if o == nil || IsNil(o.NulledAt) {
		var ret time.Time
		return ret
	}
	return *o.NulledAt
}

// GetNulledAtOk returns a tuple with the NulledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullRoute) GetNulledAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NulledAt) {
		return nil, false
	}
	return o.NulledAt, true
}

// HasNulledAt returns a boolean if a field has been set.
func (o *NullRoute) HasNulledAt() bool {
	if o != nil && !IsNil(o.NulledAt) {
		return true
	}

	return false
}

// SetNulledAt gets a reference to the given time.Time and assigns it to the NulledAt field.
func (o *NullRoute) SetNulledAt(v time.Time) {
	o.NulledAt = &v
}

// GetTicketId returns the TicketId field value if set, zero value otherwise.
func (o *NullRoute) GetTicketId() string {
	if o == nil || IsNil(o.TicketId) {
		var ret string
		return ret
	}
	return *o.TicketId
}

// GetTicketIdOk returns a tuple with the TicketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullRoute) GetTicketIdOk() (*string, bool) {
	if o == nil || IsNil(o.TicketId) {
		return nil, false
	}
	return o.TicketId, true
}

// HasTicketId returns a boolean if a field has been set.
func (o *NullRoute) HasTicketId() bool {
	if o != nil && !IsNil(o.TicketId) {
		return true
	}

	return false
}

// SetTicketId gets a reference to the given string and assigns it to the TicketId field.
func (o *NullRoute) SetTicketId(v string) {
	o.TicketId = &v
}

func (o NullRoute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NullRoute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutomatedUnnullingAt) {
		toSerialize["automatedUnnullingAt"] = o.AutomatedUnnullingAt
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.NullLevel) {
		toSerialize["nullLevel"] = o.NullLevel
	}
	if !IsNil(o.NulledAt) {
		toSerialize["nulledAt"] = o.NulledAt
	}
	if !IsNil(o.TicketId) {
		toSerialize["ticketId"] = o.TicketId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NullRoute) UnmarshalJSON(data []byte) (err error) {
	varNullRoute := _NullRoute{}

	err = json.Unmarshal(data, &varNullRoute)

	if err != nil {
		return err
	}

	*o = NullRoute(varNullRoute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "automatedUnnullingAt")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "nullLevel")
		delete(additionalProperties, "nulledAt")
		delete(additionalProperties, "ticketId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNullRoute struct {
	value *NullRoute
	isSet bool
}

func (v NullableNullRoute) Get() *NullRoute {
	return v.value
}

func (v *NullableNullRoute) Set(val *NullRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableNullRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableNullRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullRoute(val *NullRoute) *NullableNullRoute {
	return &NullableNullRoute{value: val, isSet: true}
}

func (v NullableNullRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


