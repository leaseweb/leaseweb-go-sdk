/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the VpsContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VpsContract{}

// VpsContract struct for VpsContract
type VpsContract struct {
	BillingFrequency BillingFrequency `json:"billingFrequency"`
	Term ContractTerm `json:"term"`
	Type ContractType `json:"type"`
	EndsAt NullableTime `json:"endsAt"`
	// Date when the contract was created
	CreatedAt time.Time `json:"createdAt"`
	State ContractState `json:"state"`
	Id string `json:"id"`
	Sla string `json:"sla"`
	ControlPanel NullableString `json:"controlPanel"`
	InModification bool `json:"inModification"`
	AdditionalProperties map[string]interface{}
}

type _VpsContract VpsContract

// NewVpsContract instantiates a new VpsContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpsContract(billingFrequency BillingFrequency, term ContractTerm, type_ ContractType, endsAt NullableTime, createdAt time.Time, state ContractState, id string, sla string, controlPanel NullableString, inModification bool) *VpsContract {
	this := VpsContract{}
	this.BillingFrequency = billingFrequency
	this.Term = term
	this.Type = type_
	this.EndsAt = endsAt
	this.CreatedAt = createdAt
	this.State = state
	this.Id = id
	this.Sla = sla
	this.ControlPanel = controlPanel
	this.InModification = inModification
	return &this
}

// NewVpsContractWithDefaults instantiates a new VpsContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpsContractWithDefaults() *VpsContract {
	this := VpsContract{}
	return &this
}

// GetBillingFrequency returns the BillingFrequency field value
func (o *VpsContract) GetBillingFrequency() BillingFrequency {
	if o == nil {
		var ret BillingFrequency
		return ret
	}

	return o.BillingFrequency
}

// GetBillingFrequencyOk returns a tuple with the BillingFrequency field value
// and a boolean to check if the value has been set.
func (o *VpsContract) GetBillingFrequencyOk() (*BillingFrequency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingFrequency, true
}

// SetBillingFrequency sets field value
func (o *VpsContract) SetBillingFrequency(v BillingFrequency) {
	o.BillingFrequency = v
}

// GetTerm returns the Term field value
func (o *VpsContract) GetTerm() ContractTerm {
	if o == nil {
		var ret ContractTerm
		return ret
	}

	return o.Term
}

// GetTermOk returns a tuple with the Term field value
// and a boolean to check if the value has been set.
func (o *VpsContract) GetTermOk() (*ContractTerm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Term, true
}

// SetTerm sets field value
func (o *VpsContract) SetTerm(v ContractTerm) {
	o.Term = v
}

// GetType returns the Type field value
func (o *VpsContract) GetType() ContractType {
	if o == nil {
		var ret ContractType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VpsContract) GetTypeOk() (*ContractType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VpsContract) SetType(v ContractType) {
	o.Type = v
}

// GetEndsAt returns the EndsAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VpsContract) GetEndsAt() time.Time {
	if o == nil || o.EndsAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.EndsAt.Get()
}

// GetEndsAtOk returns a tuple with the EndsAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpsContract) GetEndsAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndsAt.Get(), o.EndsAt.IsSet()
}

// SetEndsAt sets field value
func (o *VpsContract) SetEndsAt(v time.Time) {
	o.EndsAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *VpsContract) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *VpsContract) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *VpsContract) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetState returns the State field value
func (o *VpsContract) GetState() ContractState {
	if o == nil {
		var ret ContractState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *VpsContract) GetStateOk() (*ContractState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *VpsContract) SetState(v ContractState) {
	o.State = v
}

// GetId returns the Id field value
func (o *VpsContract) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VpsContract) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VpsContract) SetId(v string) {
	o.Id = v
}

// GetSla returns the Sla field value
func (o *VpsContract) GetSla() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sla
}

// GetSlaOk returns a tuple with the Sla field value
// and a boolean to check if the value has been set.
func (o *VpsContract) GetSlaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sla, true
}

// SetSla sets field value
func (o *VpsContract) SetSla(v string) {
	o.Sla = v
}

// GetControlPanel returns the ControlPanel field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VpsContract) GetControlPanel() string {
	if o == nil || o.ControlPanel.Get() == nil {
		var ret string
		return ret
	}

	return *o.ControlPanel.Get()
}

// GetControlPanelOk returns a tuple with the ControlPanel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpsContract) GetControlPanelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ControlPanel.Get(), o.ControlPanel.IsSet()
}

// SetControlPanel sets field value
func (o *VpsContract) SetControlPanel(v string) {
	o.ControlPanel.Set(&v)
}

// GetInModification returns the InModification field value
func (o *VpsContract) GetInModification() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InModification
}

// GetInModificationOk returns a tuple with the InModification field value
// and a boolean to check if the value has been set.
func (o *VpsContract) GetInModificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InModification, true
}

// SetInModification sets field value
func (o *VpsContract) SetInModification(v bool) {
	o.InModification = v
}

func (o VpsContract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VpsContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["billingFrequency"] = o.BillingFrequency
	toSerialize["term"] = o.Term
	toSerialize["type"] = o.Type
	toSerialize["endsAt"] = o.EndsAt.Get()
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["state"] = o.State
	toSerialize["id"] = o.Id
	toSerialize["sla"] = o.Sla
	toSerialize["controlPanel"] = o.ControlPanel.Get()
	toSerialize["inModification"] = o.InModification

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VpsContract) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"billingFrequency",
		"term",
		"type",
		"endsAt",
		"createdAt",
		"state",
		"id",
		"sla",
		"controlPanel",
		"inModification",
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

	varVpsContract := _VpsContract{}

	err = json.Unmarshal(data, &varVpsContract)

	if err != nil {
		return err
	}

	*o = VpsContract(varVpsContract)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "billingFrequency")
		delete(additionalProperties, "term")
		delete(additionalProperties, "type")
		delete(additionalProperties, "endsAt")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "state")
		delete(additionalProperties, "id")
		delete(additionalProperties, "sla")
		delete(additionalProperties, "controlPanel")
		delete(additionalProperties, "inModification")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVpsContract struct {
	value *VpsContract
	isSet bool
}

func (v NullableVpsContract) Get() *VpsContract {
	return v.value
}

func (v *NullableVpsContract) Set(val *VpsContract) {
	v.value = val
	v.isSet = true
}

func (v NullableVpsContract) IsSet() bool {
	return v.isSet
}

func (v *NullableVpsContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpsContract(val *VpsContract) *NullableVpsContract {
	return &NullableVpsContract{value: val, isSet: true}
}

func (v NullableVpsContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpsContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


