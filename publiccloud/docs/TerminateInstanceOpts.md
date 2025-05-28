# TerminateInstanceOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonCode** | Pointer to **string** | The reason code for terminating the instance. This is a required field. Please refer at this page for the valid options: [Cancel Reasons](https://developer.leaseweb.com/docs/#tag/Services/operation/services-cancel-reasons-get) | [optional] 
**Reason** | Pointer to **string** | Required only when reasonCode is CANCEL_OTHER. | [optional] 

## Methods

### NewTerminateInstanceOpts

`func NewTerminateInstanceOpts() *TerminateInstanceOpts`

NewTerminateInstanceOpts instantiates a new TerminateInstanceOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminateInstanceOptsWithDefaults

`func NewTerminateInstanceOptsWithDefaults() *TerminateInstanceOpts`

NewTerminateInstanceOptsWithDefaults instantiates a new TerminateInstanceOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonCode

`func (o *TerminateInstanceOpts) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *TerminateInstanceOpts) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *TerminateInstanceOpts) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *TerminateInstanceOpts) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetReason

`func (o *TerminateInstanceOpts) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TerminateInstanceOpts) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TerminateInstanceOpts) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TerminateInstanceOpts) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


