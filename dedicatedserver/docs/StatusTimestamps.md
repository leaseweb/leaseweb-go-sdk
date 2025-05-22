# StatusTimestamps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WAITING** | Pointer to **time.Time** |  | [optional] 
**PENDING** | Pointer to **time.Time** |  | [optional] 
**INPROGRESS** | Pointer to **time.Time** |  | [optional] 
**REBOOTING** | Pointer to **time.Time** |  | [optional] 
**FINISHED** | Pointer to **time.Time** |  | [optional] 
**FAILED** | Pointer to **time.Time** |  | [optional] 
**SKIPPED** | Pointer to **time.Time** |  | [optional] 
**CANCELED** | Pointer to **time.Time** |  | [optional] 
**WARNING** | Pointer to **time.Time** |  | [optional] 
**EXPIRED** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStatusTimestamps

`func NewStatusTimestamps() *StatusTimestamps`

NewStatusTimestamps instantiates a new StatusTimestamps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusTimestampsWithDefaults

`func NewStatusTimestampsWithDefaults() *StatusTimestamps`

NewStatusTimestampsWithDefaults instantiates a new StatusTimestamps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWAITING

`func (o *StatusTimestamps) GetWAITING() time.Time`

GetWAITING returns the WAITING field if non-nil, zero value otherwise.

### GetWAITINGOk

`func (o *StatusTimestamps) GetWAITINGOk() (*time.Time, bool)`

GetWAITINGOk returns a tuple with the WAITING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWAITING

`func (o *StatusTimestamps) SetWAITING(v time.Time)`

SetWAITING sets WAITING field to given value.

### HasWAITING

`func (o *StatusTimestamps) HasWAITING() bool`

HasWAITING returns a boolean if a field has been set.

### GetPENDING

`func (o *StatusTimestamps) GetPENDING() time.Time`

GetPENDING returns the PENDING field if non-nil, zero value otherwise.

### GetPENDINGOk

`func (o *StatusTimestamps) GetPENDINGOk() (*time.Time, bool)`

GetPENDINGOk returns a tuple with the PENDING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPENDING

`func (o *StatusTimestamps) SetPENDING(v time.Time)`

SetPENDING sets PENDING field to given value.

### HasPENDING

`func (o *StatusTimestamps) HasPENDING() bool`

HasPENDING returns a boolean if a field has been set.

### GetINPROGRESS

`func (o *StatusTimestamps) GetINPROGRESS() time.Time`

GetINPROGRESS returns the INPROGRESS field if non-nil, zero value otherwise.

### GetINPROGRESSOk

`func (o *StatusTimestamps) GetINPROGRESSOk() (*time.Time, bool)`

GetINPROGRESSOk returns a tuple with the INPROGRESS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINPROGRESS

`func (o *StatusTimestamps) SetINPROGRESS(v time.Time)`

SetINPROGRESS sets INPROGRESS field to given value.

### HasINPROGRESS

`func (o *StatusTimestamps) HasINPROGRESS() bool`

HasINPROGRESS returns a boolean if a field has been set.

### GetREBOOTING

`func (o *StatusTimestamps) GetREBOOTING() time.Time`

GetREBOOTING returns the REBOOTING field if non-nil, zero value otherwise.

### GetREBOOTINGOk

`func (o *StatusTimestamps) GetREBOOTINGOk() (*time.Time, bool)`

GetREBOOTINGOk returns a tuple with the REBOOTING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREBOOTING

`func (o *StatusTimestamps) SetREBOOTING(v time.Time)`

SetREBOOTING sets REBOOTING field to given value.

### HasREBOOTING

`func (o *StatusTimestamps) HasREBOOTING() bool`

HasREBOOTING returns a boolean if a field has been set.

### GetFINISHED

`func (o *StatusTimestamps) GetFINISHED() time.Time`

GetFINISHED returns the FINISHED field if non-nil, zero value otherwise.

### GetFINISHEDOk

`func (o *StatusTimestamps) GetFINISHEDOk() (*time.Time, bool)`

GetFINISHEDOk returns a tuple with the FINISHED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFINISHED

`func (o *StatusTimestamps) SetFINISHED(v time.Time)`

SetFINISHED sets FINISHED field to given value.

### HasFINISHED

`func (o *StatusTimestamps) HasFINISHED() bool`

HasFINISHED returns a boolean if a field has been set.

### GetFAILED

`func (o *StatusTimestamps) GetFAILED() time.Time`

GetFAILED returns the FAILED field if non-nil, zero value otherwise.

### GetFAILEDOk

`func (o *StatusTimestamps) GetFAILEDOk() (*time.Time, bool)`

GetFAILEDOk returns a tuple with the FAILED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFAILED

`func (o *StatusTimestamps) SetFAILED(v time.Time)`

SetFAILED sets FAILED field to given value.

### HasFAILED

`func (o *StatusTimestamps) HasFAILED() bool`

HasFAILED returns a boolean if a field has been set.

### GetSKIPPED

`func (o *StatusTimestamps) GetSKIPPED() time.Time`

GetSKIPPED returns the SKIPPED field if non-nil, zero value otherwise.

### GetSKIPPEDOk

`func (o *StatusTimestamps) GetSKIPPEDOk() (*time.Time, bool)`

GetSKIPPEDOk returns a tuple with the SKIPPED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSKIPPED

`func (o *StatusTimestamps) SetSKIPPED(v time.Time)`

SetSKIPPED sets SKIPPED field to given value.

### HasSKIPPED

`func (o *StatusTimestamps) HasSKIPPED() bool`

HasSKIPPED returns a boolean if a field has been set.

### GetCANCELED

`func (o *StatusTimestamps) GetCANCELED() time.Time`

GetCANCELED returns the CANCELED field if non-nil, zero value otherwise.

### GetCANCELEDOk

`func (o *StatusTimestamps) GetCANCELEDOk() (*time.Time, bool)`

GetCANCELEDOk returns a tuple with the CANCELED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCANCELED

`func (o *StatusTimestamps) SetCANCELED(v time.Time)`

SetCANCELED sets CANCELED field to given value.

### HasCANCELED

`func (o *StatusTimestamps) HasCANCELED() bool`

HasCANCELED returns a boolean if a field has been set.

### GetWARNING

`func (o *StatusTimestamps) GetWARNING() time.Time`

GetWARNING returns the WARNING field if non-nil, zero value otherwise.

### GetWARNINGOk

`func (o *StatusTimestamps) GetWARNINGOk() (*time.Time, bool)`

GetWARNINGOk returns a tuple with the WARNING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWARNING

`func (o *StatusTimestamps) SetWARNING(v time.Time)`

SetWARNING sets WARNING field to given value.

### HasWARNING

`func (o *StatusTimestamps) HasWARNING() bool`

HasWARNING returns a boolean if a field has been set.

### GetEXPIRED

`func (o *StatusTimestamps) GetEXPIRED() time.Time`

GetEXPIRED returns the EXPIRED field if non-nil, zero value otherwise.

### GetEXPIREDOk

`func (o *StatusTimestamps) GetEXPIREDOk() (*time.Time, bool)`

GetEXPIREDOk returns a tuple with the EXPIRED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEXPIRED

`func (o *StatusTimestamps) SetEXPIRED(v time.Time)`

SetEXPIRED sets EXPIRED field to given value.

### HasEXPIRED

`func (o *StatusTimestamps) HasEXPIRED() bool`

HasEXPIRED returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


