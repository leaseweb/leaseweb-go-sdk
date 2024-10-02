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

// checks if the UpdateLoadBalancerOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLoadBalancerOpts{}

// UpdateLoadBalancerOpts struct for UpdateLoadBalancerOpts
type UpdateLoadBalancerOpts struct {
	Type *TypeName `json:"type,omitempty"`
	// An identifying name you can refer to the load balancer
	Reference *string `json:"reference,omitempty"`
	ContractType *ContractType `json:"contractType,omitempty"`
	StickySession NullableStickySession `json:"stickySession,omitempty"`
	Balance *Balance `json:"balance,omitempty"`
	// Is xForwardedFor enabled or not
	XForwardedFor *bool `json:"xForwardedFor,omitempty"`
	// Time to close the connection if load balancer is idle
	IdleTimeOut *int32 `json:"idleTimeOut,omitempty"`
	// Port on which the backend (target) servers are listening to handle incoming requests
	TargetPort *int32 `json:"targetPort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateLoadBalancerOpts UpdateLoadBalancerOpts

// NewUpdateLoadBalancerOpts instantiates a new UpdateLoadBalancerOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLoadBalancerOpts() *UpdateLoadBalancerOpts {
	this := UpdateLoadBalancerOpts{}
	return &this
}

// NewUpdateLoadBalancerOptsWithDefaults instantiates a new UpdateLoadBalancerOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLoadBalancerOptsWithDefaults() *UpdateLoadBalancerOpts {
	this := UpdateLoadBalancerOpts{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateLoadBalancerOpts) GetType() TypeName {
	if o == nil || IsNil(o.Type) {
		var ret TypeName
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerOpts) GetTypeOk() (*TypeName, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateLoadBalancerOpts) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeName and assigns it to the Type field.
func (o *UpdateLoadBalancerOpts) SetType(v TypeName) {
	o.Type = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *UpdateLoadBalancerOpts) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerOpts) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *UpdateLoadBalancerOpts) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *UpdateLoadBalancerOpts) SetReference(v string) {
	o.Reference = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *UpdateLoadBalancerOpts) GetContractType() ContractType {
	if o == nil || IsNil(o.ContractType) {
		var ret ContractType
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerOpts) GetContractTypeOk() (*ContractType, bool) {
	if o == nil || IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *UpdateLoadBalancerOpts) HasContractType() bool {
	if o != nil && !IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given ContractType and assigns it to the ContractType field.
func (o *UpdateLoadBalancerOpts) SetContractType(v ContractType) {
	o.ContractType = &v
}

// GetStickySession returns the StickySession field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateLoadBalancerOpts) GetStickySession() StickySession {
	if o == nil || IsNil(o.StickySession.Get()) {
		var ret StickySession
		return ret
	}
	return *o.StickySession.Get()
}

// GetStickySessionOk returns a tuple with the StickySession field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateLoadBalancerOpts) GetStickySessionOk() (*StickySession, bool) {
	if o == nil {
		return nil, false
	}
	return o.StickySession.Get(), o.StickySession.IsSet()
}

// HasStickySession returns a boolean if a field has been set.
func (o *UpdateLoadBalancerOpts) HasStickySession() bool {
	if o != nil && o.StickySession.IsSet() {
		return true
	}

	return false
}

// SetStickySession gets a reference to the given NullableStickySession and assigns it to the StickySession field.
func (o *UpdateLoadBalancerOpts) SetStickySession(v StickySession) {
	o.StickySession.Set(&v)
}
// SetStickySessionNil sets the value for StickySession to be an explicit nil
func (o *UpdateLoadBalancerOpts) SetStickySessionNil() {
	o.StickySession.Set(nil)
}

// UnsetStickySession ensures that no value is present for StickySession, not even an explicit nil
func (o *UpdateLoadBalancerOpts) UnsetStickySession() {
	o.StickySession.Unset()
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *UpdateLoadBalancerOpts) GetBalance() Balance {
	if o == nil || IsNil(o.Balance) {
		var ret Balance
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerOpts) GetBalanceOk() (*Balance, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *UpdateLoadBalancerOpts) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given Balance and assigns it to the Balance field.
func (o *UpdateLoadBalancerOpts) SetBalance(v Balance) {
	o.Balance = &v
}

// GetXForwardedFor returns the XForwardedFor field value if set, zero value otherwise.
func (o *UpdateLoadBalancerOpts) GetXForwardedFor() bool {
	if o == nil || IsNil(o.XForwardedFor) {
		var ret bool
		return ret
	}
	return *o.XForwardedFor
}

// GetXForwardedForOk returns a tuple with the XForwardedFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerOpts) GetXForwardedForOk() (*bool, bool) {
	if o == nil || IsNil(o.XForwardedFor) {
		return nil, false
	}
	return o.XForwardedFor, true
}

// HasXForwardedFor returns a boolean if a field has been set.
func (o *UpdateLoadBalancerOpts) HasXForwardedFor() bool {
	if o != nil && !IsNil(o.XForwardedFor) {
		return true
	}

	return false
}

// SetXForwardedFor gets a reference to the given bool and assigns it to the XForwardedFor field.
func (o *UpdateLoadBalancerOpts) SetXForwardedFor(v bool) {
	o.XForwardedFor = &v
}

// GetIdleTimeOut returns the IdleTimeOut field value if set, zero value otherwise.
func (o *UpdateLoadBalancerOpts) GetIdleTimeOut() int32 {
	if o == nil || IsNil(o.IdleTimeOut) {
		var ret int32
		return ret
	}
	return *o.IdleTimeOut
}

// GetIdleTimeOutOk returns a tuple with the IdleTimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerOpts) GetIdleTimeOutOk() (*int32, bool) {
	if o == nil || IsNil(o.IdleTimeOut) {
		return nil, false
	}
	return o.IdleTimeOut, true
}

// HasIdleTimeOut returns a boolean if a field has been set.
func (o *UpdateLoadBalancerOpts) HasIdleTimeOut() bool {
	if o != nil && !IsNil(o.IdleTimeOut) {
		return true
	}

	return false
}

// SetIdleTimeOut gets a reference to the given int32 and assigns it to the IdleTimeOut field.
func (o *UpdateLoadBalancerOpts) SetIdleTimeOut(v int32) {
	o.IdleTimeOut = &v
}

// GetTargetPort returns the TargetPort field value if set, zero value otherwise.
func (o *UpdateLoadBalancerOpts) GetTargetPort() int32 {
	if o == nil || IsNil(o.TargetPort) {
		var ret int32
		return ret
	}
	return *o.TargetPort
}

// GetTargetPortOk returns a tuple with the TargetPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerOpts) GetTargetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetPort) {
		return nil, false
	}
	return o.TargetPort, true
}

// HasTargetPort returns a boolean if a field has been set.
func (o *UpdateLoadBalancerOpts) HasTargetPort() bool {
	if o != nil && !IsNil(o.TargetPort) {
		return true
	}

	return false
}

// SetTargetPort gets a reference to the given int32 and assigns it to the TargetPort field.
func (o *UpdateLoadBalancerOpts) SetTargetPort(v int32) {
	o.TargetPort = &v
}

func (o UpdateLoadBalancerOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLoadBalancerOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.ContractType) {
		toSerialize["contractType"] = o.ContractType
	}
	if o.StickySession.IsSet() {
		toSerialize["stickySession"] = o.StickySession.Get()
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.XForwardedFor) {
		toSerialize["xForwardedFor"] = o.XForwardedFor
	}
	if !IsNil(o.IdleTimeOut) {
		toSerialize["idleTimeOut"] = o.IdleTimeOut
	}
	if !IsNil(o.TargetPort) {
		toSerialize["targetPort"] = o.TargetPort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateLoadBalancerOpts) UnmarshalJSON(data []byte) (err error) {
	varUpdateLoadBalancerOpts := _UpdateLoadBalancerOpts{}

	err = json.Unmarshal(data, &varUpdateLoadBalancerOpts)

	if err != nil {
		return err
	}

	*o = UpdateLoadBalancerOpts(varUpdateLoadBalancerOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "contractType")
		delete(additionalProperties, "stickySession")
		delete(additionalProperties, "balance")
		delete(additionalProperties, "xForwardedFor")
		delete(additionalProperties, "idleTimeOut")
		delete(additionalProperties, "targetPort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateLoadBalancerOpts struct {
	value *UpdateLoadBalancerOpts
	isSet bool
}

func (v NullableUpdateLoadBalancerOpts) Get() *UpdateLoadBalancerOpts {
	return v.value
}

func (v *NullableUpdateLoadBalancerOpts) Set(val *UpdateLoadBalancerOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLoadBalancerOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLoadBalancerOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLoadBalancerOpts(val *UpdateLoadBalancerOpts) *NullableUpdateLoadBalancerOpts {
	return &NullableUpdateLoadBalancerOpts{value: val, isSet: true}
}

func (v NullableUpdateLoadBalancerOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLoadBalancerOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


