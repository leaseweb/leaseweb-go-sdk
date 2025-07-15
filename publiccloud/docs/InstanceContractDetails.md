# InstanceContractDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingFrequency** | [**BillingFrequency**](BillingFrequency.md) |  | 
**Term** | [**ContractTerm**](ContractTerm.md) |  | 
**Type** | [**ContractType**](ContractType.md) |  | 
**EndsAt** | **NullableTime** | End date is null for hourly instances | 
**StartsAt** | **time.Time** | Date when the contract is starting | 
**State** | [**ContractState**](ContractState.md) |  | 
**Id** | **string** | Contract id is empty for hourly instances | 
**Sla** | **string** |  | 
**ControlPanel** | **NullableString** |  | 
**InModification** | **bool** |  | 

## Methods

### NewInstanceContractDetails

`func NewInstanceContractDetails(billingFrequency BillingFrequency, term ContractTerm, type_ ContractType, endsAt NullableTime, startsAt time.Time, state ContractState, id string, sla string, controlPanel NullableString, inModification bool, ) *InstanceContractDetails`

NewInstanceContractDetails instantiates a new InstanceContractDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContractDetailsWithDefaults

`func NewInstanceContractDetailsWithDefaults() *InstanceContractDetails`

NewInstanceContractDetailsWithDefaults instantiates a new InstanceContractDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingFrequency

`func (o *InstanceContractDetails) GetBillingFrequency() BillingFrequency`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *InstanceContractDetails) GetBillingFrequencyOk() (*BillingFrequency, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *InstanceContractDetails) SetBillingFrequency(v BillingFrequency)`

SetBillingFrequency sets BillingFrequency field to given value.


### GetTerm

`func (o *InstanceContractDetails) GetTerm() ContractTerm`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *InstanceContractDetails) GetTermOk() (*ContractTerm, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *InstanceContractDetails) SetTerm(v ContractTerm)`

SetTerm sets Term field to given value.


### GetType

`func (o *InstanceContractDetails) GetType() ContractType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceContractDetails) GetTypeOk() (*ContractType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceContractDetails) SetType(v ContractType)`

SetType sets Type field to given value.


### GetEndsAt

`func (o *InstanceContractDetails) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *InstanceContractDetails) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *InstanceContractDetails) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### SetEndsAtNil

`func (o *InstanceContractDetails) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *InstanceContractDetails) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
### GetStartsAt

`func (o *InstanceContractDetails) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *InstanceContractDetails) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *InstanceContractDetails) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetState

`func (o *InstanceContractDetails) GetState() ContractState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InstanceContractDetails) GetStateOk() (*ContractState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InstanceContractDetails) SetState(v ContractState)`

SetState sets State field to given value.


### GetId

`func (o *InstanceContractDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContractDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContractDetails) SetId(v string)`

SetId sets Id field to given value.


### GetSla

`func (o *InstanceContractDetails) GetSla() string`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *InstanceContractDetails) GetSlaOk() (*string, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *InstanceContractDetails) SetSla(v string)`

SetSla sets Sla field to given value.


### GetControlPanel

`func (o *InstanceContractDetails) GetControlPanel() string`

GetControlPanel returns the ControlPanel field if non-nil, zero value otherwise.

### GetControlPanelOk

`func (o *InstanceContractDetails) GetControlPanelOk() (*string, bool)`

GetControlPanelOk returns a tuple with the ControlPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPanel

`func (o *InstanceContractDetails) SetControlPanel(v string)`

SetControlPanel sets ControlPanel field to given value.


### SetControlPanelNil

`func (o *InstanceContractDetails) SetControlPanelNil(b bool)`

 SetControlPanelNil sets the value for ControlPanel to be an explicit nil

### UnsetControlPanel
`func (o *InstanceContractDetails) UnsetControlPanel()`

UnsetControlPanel ensures that no value is present for ControlPanel, not even an explicit nil
### GetInModification

`func (o *InstanceContractDetails) GetInModification() bool`

GetInModification returns the InModification field if non-nil, zero value otherwise.

### GetInModificationOk

`func (o *InstanceContractDetails) GetInModificationOk() (*bool, bool)`

GetInModificationOk returns a tuple with the InModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInModification

`func (o *InstanceContractDetails) SetInModification(v bool)`

SetInModification sets InModification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


