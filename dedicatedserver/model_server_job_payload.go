/*
Dedicated Servers

This is the description of the Dedicated Server API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
)

// checks if the ServerJobPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerJobPayload{}

// ServerJobPayload struct for ServerJobPayload
type ServerJobPayload struct {
	FileserverBaseUrl *string `json:"fileserverBaseUrl,omitempty"`
	Pop *string `json:"pop,omitempty"`
	PowerCycle *bool `json:"powerCycle,omitempty"`
	IsUnassignedServer *bool `json:"isUnassignedServer,omitempty"`
	// Id of the server
	ServerId *string `json:"serverId,omitempty"`
	JobType *string `json:"jobType,omitempty"`
	Configurable *bool `json:"configurable,omitempty"`
	Device *string `json:"device,omitempty"`
	NumberOfDisks NullableInt32 `json:"numberOfDisks,omitempty"`
	OperatingSystemId *string `json:"operatingSystemId,omitempty"`
	Os *Os `json:"os,omitempty"`
	Partitions []Partition `json:"partitions,omitempty"`
	RaidLevel NullableInt32 `json:"raidLevel,omitempty"`
	// Timezone represented as Geographical_Area/City
	Timezone *string `json:"timezone,omitempty"`
	Raid NullableRaidPayload `json:"raid,omitempty"`
	// The hostname of the server
	Hostname *string `json:"hostname,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerJobPayload ServerJobPayload

// NewServerJobPayload instantiates a new ServerJobPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerJobPayload() *ServerJobPayload {
	this := ServerJobPayload{}
	return &this
}

// NewServerJobPayloadWithDefaults instantiates a new ServerJobPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerJobPayloadWithDefaults() *ServerJobPayload {
	this := ServerJobPayload{}
	return &this
}

// GetFileserverBaseUrl returns the FileserverBaseUrl field value if set, zero value otherwise.
func (o *ServerJobPayload) GetFileserverBaseUrl() string {
	if o == nil || IsNil(o.FileserverBaseUrl) {
		var ret string
		return ret
	}
	return *o.FileserverBaseUrl
}

// GetFileserverBaseUrlOk returns a tuple with the FileserverBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerJobPayload) GetFileserverBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FileserverBaseUrl) {
		return nil, false
	}
	return o.FileserverBaseUrl, true
}

// HasFileserverBaseUrl returns a boolean if a field has been set.
func (o *ServerJobPayload) HasFileserverBaseUrl() bool {
	if o != nil && !IsNil(o.FileserverBaseUrl) {
		return true
	}

	return false
}

// SetFileserverBaseUrl gets a reference to the given string and assigns it to the FileserverBaseUrl field.
func (o *ServerJobPayload) SetFileserverBaseUrl(v string) {
	o.FileserverBaseUrl = &v
}

// GetPop returns the Pop field value if set, zero value otherwise.
func (o *ServerJobPayload) GetPop() string {
	if o == nil || IsNil(o.Pop) {
		var ret string
		return ret
	}
	return *o.Pop
}

// GetPopOk returns a tuple with the Pop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerJobPayload) GetPopOk() (*string, bool) {
	if o == nil || IsNil(o.Pop) {
		return nil, false
	}
	return o.Pop, true
}

// HasPop returns a boolean if a field has been set.
func (o *ServerJobPayload) HasPop() bool {
	if o != nil && !IsNil(o.Pop) {
		return true
	}

	return false
}

// SetPop gets a reference to the given string and assigns it to the Pop field.
func (o *ServerJobPayload) SetPop(v string) {
	o.Pop = &v
}

// GetPowerCycle returns the PowerCycle field value if set, zero value otherwise.
func (o *ServerJobPayload) GetPowerCycle() bool {
	if o == nil || IsNil(o.PowerCycle) {
		var ret bool
		return ret
	}
	return *o.PowerCycle
}

// GetPowerCycleOk returns a tuple with the PowerCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerJobPayload) GetPowerCycleOk() (*bool, bool) {
	if o == nil || IsNil(o.PowerCycle) {
		return nil, false
	}
	return o.PowerCycle, true
}

// HasPowerCycle returns a boolean if a field has been set.
func (o *ServerJobPayload) HasPowerCycle() bool {
	if o != nil && !IsNil(o.PowerCycle) {
		return true
	}

	return false
}

// SetPowerCycle gets a reference to the given bool and assigns it to the PowerCycle field.
func (o *ServerJobPayload) SetPowerCycle(v bool) {
	o.PowerCycle = &v
}

// GetIsUnassignedServer returns the IsUnassignedServer field value if set, zero value otherwise.
func (o *ServerJobPayload) GetIsUnassignedServer() bool {
	if o == nil || IsNil(o.IsUnassignedServer) {
		var ret bool
		return ret
	}
	return *o.IsUnassignedServer
}

// GetIsUnassignedServerOk returns a tuple with the IsUnassignedServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerJobPayload) GetIsUnassignedServerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUnassignedServer) {
		return nil, false
	}
	return o.IsUnassignedServer, true
}

// HasIsUnassignedServer returns a boolean if a field has been set.
func (o *ServerJobPayload) HasIsUnassignedServer() bool {
	if o != nil && !IsNil(o.IsUnassignedServer) {
		return true
	}

	return false
}

// SetIsUnassignedServer gets a reference to the given bool and assigns it to the IsUnassignedServer field.
func (o *ServerJobPayload) SetIsUnassignedServer(v bool) {
	o.IsUnassignedServer = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *ServerJobPayload) GetServerId() string {
	if o == nil || IsNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerJobPayload) GetServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *ServerJobPayload) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *ServerJobPayload) SetServerId(v string) {
	o.ServerId = &v
}

// GetJobType returns the JobType field value if set, zero value otherwise.
func (o *ServerJobPayload) GetJobType() string {
	if o == nil || IsNil(o.JobType) {
		var ret string
		return ret
	}
	return *o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerJobPayload) GetJobTypeOk() (*string, bool) {
	if o == nil || IsNil(o.JobType) {
		return nil, false
	}
	return o.JobType, true
}

// HasJobType returns a boolean if a field has been set.
func (o *ServerJobPayload) HasJobType() bool {
	if o != nil && !IsNil(o.JobType) {
		return true
	}

	return false
}

// SetJobType gets a reference to the given string and assigns it to the JobType field.
func (o *ServerJobPayload) SetJobType(v string) {
	o.JobType = &v
}

// GetConfigurable returns the Configurable field value if set, zero value otherwise.
func (o *ServerJobPayload) GetConfigurable() bool {
	if o == nil || IsNil(o.Configurable) {
		var ret bool
		return ret
	}
	return *o.Configurable
}

// GetConfigurableOk returns a tuple with the Configurable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerJobPayload) GetConfigurableOk() (*bool, bool) {
	if o == nil || IsNil(o.Configurable) {
		return nil, false
	}
	return o.Configurable, true
}

// HasConfigurable returns a boolean if a field has been set.
func (o *ServerJobPayload) HasConfigurable() bool {
	if o != nil && !IsNil(o.Configurable) {
		return true
	}

	return false
}

// SetConfigurable gets a reference to the given bool and assigns it to the Configurable field.
func (o *ServerJobPayload) SetConfigurable(v bool) {
	o.Configurable = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *ServerJobPayload) GetDevice() string {
	if o == nil || IsNil(o.Device) {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerJobPayload) GetDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *ServerJobPayload) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *ServerJobPayload) SetDevice(v string) {
	o.Device = &v
}

// GetNumberOfDisks returns the NumberOfDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerJobPayload) GetNumberOfDisks() int32 {
	if o == nil || IsNil(o.NumberOfDisks.Get()) {
		var ret int32
		return ret
	}
	return *o.NumberOfDisks.Get()
}

// GetNumberOfDisksOk returns a tuple with the NumberOfDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerJobPayload) GetNumberOfDisksOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumberOfDisks.Get(), o.NumberOfDisks.IsSet()
}

// HasNumberOfDisks returns a boolean if a field has been set.
func (o *ServerJobPayload) HasNumberOfDisks() bool {
	if o != nil && o.NumberOfDisks.IsSet() {
		return true
	}

	return false
}

// SetNumberOfDisks gets a reference to the given NullableInt32 and assigns it to the NumberOfDisks field.
func (o *ServerJobPayload) SetNumberOfDisks(v int32) {
	o.NumberOfDisks.Set(&v)
}
// SetNumberOfDisksNil sets the value for NumberOfDisks to be an explicit nil
func (o *ServerJobPayload) SetNumberOfDisksNil() {
	o.NumberOfDisks.Set(nil)
}

// UnsetNumberOfDisks ensures that no value is present for NumberOfDisks, not even an explicit nil
func (o *ServerJobPayload) UnsetNumberOfDisks() {
	o.NumberOfDisks.Unset()
}

// GetOperatingSystemId returns the OperatingSystemId field value if set, zero value otherwise.
func (o *ServerJobPayload) GetOperatingSystemId() string {
	if o == nil || IsNil(o.OperatingSystemId) {
		var ret string
		return ret
	}
	return *o.OperatingSystemId
}

// GetOperatingSystemIdOk returns a tuple with the OperatingSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerJobPayload) GetOperatingSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.OperatingSystemId) {
		return nil, false
	}
	return o.OperatingSystemId, true
}

// HasOperatingSystemId returns a boolean if a field has been set.
func (o *ServerJobPayload) HasOperatingSystemId() bool {
	if o != nil && !IsNil(o.OperatingSystemId) {
		return true
	}

	return false
}

// SetOperatingSystemId gets a reference to the given string and assigns it to the OperatingSystemId field.
func (o *ServerJobPayload) SetOperatingSystemId(v string) {
	o.OperatingSystemId = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *ServerJobPayload) GetOs() Os {
	if o == nil || IsNil(o.Os) {
		var ret Os
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerJobPayload) GetOsOk() (*Os, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *ServerJobPayload) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given Os and assigns it to the Os field.
func (o *ServerJobPayload) SetOs(v Os) {
	o.Os = &v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *ServerJobPayload) GetPartitions() []Partition {
	if o == nil || IsNil(o.Partitions) {
		var ret []Partition
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerJobPayload) GetPartitionsOk() ([]Partition, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *ServerJobPayload) HasPartitions() bool {
	if o != nil && !IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []Partition and assigns it to the Partitions field.
func (o *ServerJobPayload) SetPartitions(v []Partition) {
	o.Partitions = v
}

// GetRaidLevel returns the RaidLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerJobPayload) GetRaidLevel() int32 {
	if o == nil || IsNil(o.RaidLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.RaidLevel.Get()
}

// GetRaidLevelOk returns a tuple with the RaidLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerJobPayload) GetRaidLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaidLevel.Get(), o.RaidLevel.IsSet()
}

// HasRaidLevel returns a boolean if a field has been set.
func (o *ServerJobPayload) HasRaidLevel() bool {
	if o != nil && o.RaidLevel.IsSet() {
		return true
	}

	return false
}

// SetRaidLevel gets a reference to the given NullableInt32 and assigns it to the RaidLevel field.
func (o *ServerJobPayload) SetRaidLevel(v int32) {
	o.RaidLevel.Set(&v)
}
// SetRaidLevelNil sets the value for RaidLevel to be an explicit nil
func (o *ServerJobPayload) SetRaidLevelNil() {
	o.RaidLevel.Set(nil)
}

// UnsetRaidLevel ensures that no value is present for RaidLevel, not even an explicit nil
func (o *ServerJobPayload) UnsetRaidLevel() {
	o.RaidLevel.Unset()
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ServerJobPayload) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerJobPayload) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ServerJobPayload) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ServerJobPayload) SetTimezone(v string) {
	o.Timezone = &v
}

// GetRaid returns the Raid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerJobPayload) GetRaid() RaidPayload {
	if o == nil || IsNil(o.Raid.Get()) {
		var ret RaidPayload
		return ret
	}
	return *o.Raid.Get()
}

// GetRaidOk returns a tuple with the Raid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerJobPayload) GetRaidOk() (*RaidPayload, bool) {
	if o == nil {
		return nil, false
	}
	return o.Raid.Get(), o.Raid.IsSet()
}

// HasRaid returns a boolean if a field has been set.
func (o *ServerJobPayload) HasRaid() bool {
	if o != nil && o.Raid.IsSet() {
		return true
	}

	return false
}

// SetRaid gets a reference to the given NullableRaidPayload and assigns it to the Raid field.
func (o *ServerJobPayload) SetRaid(v RaidPayload) {
	o.Raid.Set(&v)
}
// SetRaidNil sets the value for Raid to be an explicit nil
func (o *ServerJobPayload) SetRaidNil() {
	o.Raid.Set(nil)
}

// UnsetRaid ensures that no value is present for Raid, not even an explicit nil
func (o *ServerJobPayload) UnsetRaid() {
	o.Raid.Unset()
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ServerJobPayload) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerJobPayload) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ServerJobPayload) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ServerJobPayload) SetHostname(v string) {
	o.Hostname = &v
}

func (o ServerJobPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerJobPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileserverBaseUrl) {
		toSerialize["fileserverBaseUrl"] = o.FileserverBaseUrl
	}
	if !IsNil(o.Pop) {
		toSerialize["pop"] = o.Pop
	}
	if !IsNil(o.PowerCycle) {
		toSerialize["powerCycle"] = o.PowerCycle
	}
	if !IsNil(o.IsUnassignedServer) {
		toSerialize["isUnassignedServer"] = o.IsUnassignedServer
	}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if !IsNil(o.JobType) {
		toSerialize["jobType"] = o.JobType
	}
	if !IsNil(o.Configurable) {
		toSerialize["configurable"] = o.Configurable
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if o.NumberOfDisks.IsSet() {
		toSerialize["numberOfDisks"] = o.NumberOfDisks.Get()
	}
	if !IsNil(o.OperatingSystemId) {
		toSerialize["operatingSystemId"] = o.OperatingSystemId
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !IsNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}
	if o.RaidLevel.IsSet() {
		toSerialize["raidLevel"] = o.RaidLevel.Get()
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if o.Raid.IsSet() {
		toSerialize["raid"] = o.Raid.Get()
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerJobPayload) UnmarshalJSON(data []byte) (err error) {
	varServerJobPayload := _ServerJobPayload{}

	err = json.Unmarshal(data, &varServerJobPayload)

	if err != nil {
		return err
	}

	*o = ServerJobPayload(varServerJobPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fileserverBaseUrl")
		delete(additionalProperties, "pop")
		delete(additionalProperties, "powerCycle")
		delete(additionalProperties, "isUnassignedServer")
		delete(additionalProperties, "serverId")
		delete(additionalProperties, "jobType")
		delete(additionalProperties, "configurable")
		delete(additionalProperties, "device")
		delete(additionalProperties, "numberOfDisks")
		delete(additionalProperties, "operatingSystemId")
		delete(additionalProperties, "os")
		delete(additionalProperties, "partitions")
		delete(additionalProperties, "raidLevel")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "raid")
		delete(additionalProperties, "hostname")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerJobPayload struct {
	value *ServerJobPayload
	isSet bool
}

func (v NullableServerJobPayload) Get() *ServerJobPayload {
	return v.value
}

func (v *NullableServerJobPayload) Set(val *ServerJobPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableServerJobPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableServerJobPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerJobPayload(val *ServerJobPayload) *NullableServerJobPayload {
	return &NullableServerJobPayload{value: val, isSet: true}
}

func (v NullableServerJobPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerJobPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


