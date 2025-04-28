# NotificationSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Threshold** | [**NotificationSettingThreshold**](NotificationSettingThreshold.md) |  | 
**Type** | **string** | Type of the notification | 
**TimePeriod** | [**TimePeriod**](TimePeriod.md) |  | 
**Action** | [**NullableAction**](Action.md) |  | 
**Channels** | [**[]Channel**](Channel.md) |  | 

## Methods

### NewNotificationSetting

`func NewNotificationSetting(id string, threshold NotificationSettingThreshold, type_ string, timePeriod TimePeriod, action NullableAction, channels []Channel, ) *NotificationSetting`

NewNotificationSetting instantiates a new NotificationSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingWithDefaults

`func NewNotificationSettingWithDefaults() *NotificationSetting`

NewNotificationSettingWithDefaults instantiates a new NotificationSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationSetting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationSetting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationSetting) SetId(v string)`

SetId sets Id field to given value.


### GetThreshold

`func (o *NotificationSetting) GetThreshold() NotificationSettingThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *NotificationSetting) GetThresholdOk() (*NotificationSettingThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *NotificationSetting) SetThreshold(v NotificationSettingThreshold)`

SetThreshold sets Threshold field to given value.


### GetType

`func (o *NotificationSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationSetting) SetType(v string)`

SetType sets Type field to given value.


### GetTimePeriod

`func (o *NotificationSetting) GetTimePeriod() TimePeriod`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *NotificationSetting) GetTimePeriodOk() (*TimePeriod, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *NotificationSetting) SetTimePeriod(v TimePeriod)`

SetTimePeriod sets TimePeriod field to given value.


### GetAction

`func (o *NotificationSetting) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NotificationSetting) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NotificationSetting) SetAction(v Action)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *NotificationSetting) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *NotificationSetting) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetChannels

`func (o *NotificationSetting) GetChannels() []Channel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *NotificationSetting) GetChannelsOk() (*[]Channel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *NotificationSetting) SetChannels(v []Channel)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


