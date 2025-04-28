# UpdateNotificationSettingOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | Pointer to [**NotificationSettingThreshold**](NotificationSettingThreshold.md) |  | [optional] 
**TimePeriod** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**Action** | Pointer to [**NullableAction**](Action.md) |  | [optional] 
**Channels** | Pointer to [**[]UpdateNotificationSettingOptsChannelsInner**](UpdateNotificationSettingOptsChannelsInner.md) |  | [optional] 

## Methods

### NewUpdateNotificationSettingOpts

`func NewUpdateNotificationSettingOpts() *UpdateNotificationSettingOpts`

NewUpdateNotificationSettingOpts instantiates a new UpdateNotificationSettingOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationSettingOptsWithDefaults

`func NewUpdateNotificationSettingOptsWithDefaults() *UpdateNotificationSettingOpts`

NewUpdateNotificationSettingOptsWithDefaults instantiates a new UpdateNotificationSettingOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *UpdateNotificationSettingOpts) GetThreshold() NotificationSettingThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *UpdateNotificationSettingOpts) GetThresholdOk() (*NotificationSettingThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *UpdateNotificationSettingOpts) SetThreshold(v NotificationSettingThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *UpdateNotificationSettingOpts) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTimePeriod

`func (o *UpdateNotificationSettingOpts) GetTimePeriod() TimePeriod`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *UpdateNotificationSettingOpts) GetTimePeriodOk() (*TimePeriod, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *UpdateNotificationSettingOpts) SetTimePeriod(v TimePeriod)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *UpdateNotificationSettingOpts) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetAction

`func (o *UpdateNotificationSettingOpts) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateNotificationSettingOpts) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateNotificationSettingOpts) SetAction(v Action)`

SetAction sets Action field to given value.

### HasAction

`func (o *UpdateNotificationSettingOpts) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *UpdateNotificationSettingOpts) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *UpdateNotificationSettingOpts) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetChannels

`func (o *UpdateNotificationSettingOpts) GetChannels() []UpdateNotificationSettingOptsChannelsInner`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *UpdateNotificationSettingOpts) GetChannelsOk() (*[]UpdateNotificationSettingOptsChannelsInner, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *UpdateNotificationSettingOpts) SetChannels(v []UpdateNotificationSettingOptsChannelsInner)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *UpdateNotificationSettingOpts) HasChannels() bool`

HasChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


