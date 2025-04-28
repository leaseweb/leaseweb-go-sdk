# GetNotificationSettingListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationSettings** | Pointer to [**[]NotificationSetting**](NotificationSetting.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetNotificationSettingListResult

`func NewGetNotificationSettingListResult() *GetNotificationSettingListResult`

NewGetNotificationSettingListResult instantiates a new GetNotificationSettingListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationSettingListResultWithDefaults

`func NewGetNotificationSettingListResultWithDefaults() *GetNotificationSettingListResult`

NewGetNotificationSettingListResultWithDefaults instantiates a new GetNotificationSettingListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationSettings

`func (o *GetNotificationSettingListResult) GetNotificationSettings() []NotificationSetting`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *GetNotificationSettingListResult) GetNotificationSettingsOk() (*[]NotificationSetting, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *GetNotificationSettingListResult) SetNotificationSettings(v []NotificationSetting)`

SetNotificationSettings sets NotificationSettings field to given value.

### HasNotificationSettings

`func (o *GetNotificationSettingListResult) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.

### GetMetadata

`func (o *GetNotificationSettingListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetNotificationSettingListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetNotificationSettingListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetNotificationSettingListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


