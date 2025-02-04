# GetBandwidthNotificationSettingListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**BandwidthNotificationSettings** | Pointer to [**[]BandwidthNotificationSetting**](BandwidthNotificationSetting.md) | An array of Bandwidth Notification Settings | [optional] 

## Methods

### NewGetBandwidthNotificationSettingListResult

`func NewGetBandwidthNotificationSettingListResult() *GetBandwidthNotificationSettingListResult`

NewGetBandwidthNotificationSettingListResult instantiates a new GetBandwidthNotificationSettingListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBandwidthNotificationSettingListResultWithDefaults

`func NewGetBandwidthNotificationSettingListResultWithDefaults() *GetBandwidthNotificationSettingListResult`

NewGetBandwidthNotificationSettingListResultWithDefaults instantiates a new GetBandwidthNotificationSettingListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetBandwidthNotificationSettingListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetBandwidthNotificationSettingListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetBandwidthNotificationSettingListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetBandwidthNotificationSettingListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetBandwidthNotificationSettings

`func (o *GetBandwidthNotificationSettingListResult) GetBandwidthNotificationSettings() []BandwidthNotificationSetting`

GetBandwidthNotificationSettings returns the BandwidthNotificationSettings field if non-nil, zero value otherwise.

### GetBandwidthNotificationSettingsOk

`func (o *GetBandwidthNotificationSettingListResult) GetBandwidthNotificationSettingsOk() (*[]BandwidthNotificationSetting, bool)`

GetBandwidthNotificationSettingsOk returns a tuple with the BandwidthNotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthNotificationSettings

`func (o *GetBandwidthNotificationSettingListResult) SetBandwidthNotificationSettings(v []BandwidthNotificationSetting)`

SetBandwidthNotificationSettings sets BandwidthNotificationSettings field to given value.

### HasBandwidthNotificationSettings

`func (o *GetBandwidthNotificationSettingListResult) HasBandwidthNotificationSettings() bool`

HasBandwidthNotificationSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


