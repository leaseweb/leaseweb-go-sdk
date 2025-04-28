# CreateNotificationSettingOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | [**NotificationSettingThreshold**](NotificationSettingThreshold.md) |  | 
**TimePeriod** | [**TimePeriod**](TimePeriod.md) |  | 
**Action** | [**NullableAction**](Action.md) |  | 
**Channels** | [**[]UpdateNotificationSettingOptsChannelsInner**](UpdateNotificationSettingOptsChannelsInner.md) |  | 

## Methods

### NewCreateNotificationSettingOpts

`func NewCreateNotificationSettingOpts(threshold NotificationSettingThreshold, timePeriod TimePeriod, action NullableAction, channels []UpdateNotificationSettingOptsChannelsInner, ) *CreateNotificationSettingOpts`

NewCreateNotificationSettingOpts instantiates a new CreateNotificationSettingOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationSettingOptsWithDefaults

`func NewCreateNotificationSettingOptsWithDefaults() *CreateNotificationSettingOpts`

NewCreateNotificationSettingOptsWithDefaults instantiates a new CreateNotificationSettingOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *CreateNotificationSettingOpts) GetThreshold() NotificationSettingThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *CreateNotificationSettingOpts) GetThresholdOk() (*NotificationSettingThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *CreateNotificationSettingOpts) SetThreshold(v NotificationSettingThreshold)`

SetThreshold sets Threshold field to given value.


### GetTimePeriod

`func (o *CreateNotificationSettingOpts) GetTimePeriod() TimePeriod`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *CreateNotificationSettingOpts) GetTimePeriodOk() (*TimePeriod, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *CreateNotificationSettingOpts) SetTimePeriod(v TimePeriod)`

SetTimePeriod sets TimePeriod field to given value.


### GetAction

`func (o *CreateNotificationSettingOpts) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateNotificationSettingOpts) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateNotificationSettingOpts) SetAction(v Action)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *CreateNotificationSettingOpts) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *CreateNotificationSettingOpts) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetChannels

`func (o *CreateNotificationSettingOpts) GetChannels() []UpdateNotificationSettingOptsChannelsInner`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CreateNotificationSettingOpts) GetChannelsOk() (*[]UpdateNotificationSettingOptsChannelsInner, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CreateNotificationSettingOpts) SetChannels(v []UpdateNotificationSettingOptsChannelsInner)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


