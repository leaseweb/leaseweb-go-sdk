# InstanceContract

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

### NewInstanceContract

`func NewInstanceContract(billingFrequency BillingFrequency, term ContractTerm, type_ ContractType, endsAt NullableTime, startsAt time.Time, state ContractState, id string, sla string, controlPanel NullableString, inModification bool, ) *InstanceContract`

NewInstanceContract instantiates a new InstanceContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContractWithDefaults

`func NewInstanceContractWithDefaults() *InstanceContract`

NewInstanceContractWithDefaults instantiates a new InstanceContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingFrequency

`func (o *InstanceContract) GetBillingFrequency() BillingFrequency`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *InstanceContract) GetBillingFrequencyOk() (*BillingFrequency, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *InstanceContract) SetBillingFrequency(v BillingFrequency)`

SetBillingFrequency sets BillingFrequency field to given value.


### GetTerm

`func (o *InstanceContract) GetTerm() ContractTerm`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *InstanceContract) GetTermOk() (*ContractTerm, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *InstanceContract) SetTerm(v ContractTerm)`

SetTerm sets Term field to given value.


### GetType

`func (o *InstanceContract) GetType() ContractType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceContract) GetTypeOk() (*ContractType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceContract) SetType(v ContractType)`

SetType sets Type field to given value.


### GetEndsAt

`func (o *InstanceContract) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *InstanceContract) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *InstanceContract) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### SetEndsAtNil

`func (o *InstanceContract) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *InstanceContract) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
### GetStartsAt

`func (o *InstanceContract) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *InstanceContract) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *InstanceContract) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetState

`func (o *InstanceContract) GetState() ContractState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InstanceContract) GetStateOk() (*ContractState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InstanceContract) SetState(v ContractState)`

SetState sets State field to given value.


### GetId

`func (o *InstanceContract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContract) SetId(v string)`

SetId sets Id field to given value.


### GetSla

`func (o *InstanceContract) GetSla() string`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *InstanceContract) GetSlaOk() (*string, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *InstanceContract) SetSla(v string)`

SetSla sets Sla field to given value.


### GetControlPanel

`func (o *InstanceContract) GetControlPanel() string`

GetControlPanel returns the ControlPanel field if non-nil, zero value otherwise.

### GetControlPanelOk

`func (o *InstanceContract) GetControlPanelOk() (*string, bool)`

GetControlPanelOk returns a tuple with the ControlPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPanel

`func (o *InstanceContract) SetControlPanel(v string)`

SetControlPanel sets ControlPanel field to given value.


### SetControlPanelNil

`func (o *InstanceContract) SetControlPanelNil(b bool)`

 SetControlPanelNil sets the value for ControlPanel to be an explicit nil

### UnsetControlPanel
`func (o *InstanceContract) UnsetControlPanel()`

UnsetControlPanel ensures that no value is present for ControlPanel, not even an explicit nil
### GetInModification

`func (o *InstanceContract) GetInModification() bool`

GetInModification returns the InModification field if non-nil, zero value otherwise.

### GetInModificationOk

`func (o *InstanceContract) GetInModificationOk() (*bool, bool)`

GetInModificationOk returns a tuple with the InModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInModification

`func (o *InstanceContract) SetInModification(v bool)`

SetInModification sets InModification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


