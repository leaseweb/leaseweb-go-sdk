# RaidPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of RAID | [optional] 
**Level** | Pointer to **int32** | The RAID level | [optional] 
**NumberOfDisks** | Pointer to **NullableInt32** | The number of disks in the RAID | [optional] 

## Methods

### NewRaidPayload

`func NewRaidPayload() *RaidPayload`

NewRaidPayload instantiates a new RaidPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaidPayloadWithDefaults

`func NewRaidPayloadWithDefaults() *RaidPayload`

NewRaidPayloadWithDefaults instantiates a new RaidPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RaidPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RaidPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RaidPayload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RaidPayload) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLevel

`func (o *RaidPayload) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *RaidPayload) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *RaidPayload) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *RaidPayload) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetNumberOfDisks

`func (o *RaidPayload) GetNumberOfDisks() int32`

GetNumberOfDisks returns the NumberOfDisks field if non-nil, zero value otherwise.

### GetNumberOfDisksOk

`func (o *RaidPayload) GetNumberOfDisksOk() (*int32, bool)`

GetNumberOfDisksOk returns a tuple with the NumberOfDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDisks

`func (o *RaidPayload) SetNumberOfDisks(v int32)`

SetNumberOfDisks sets NumberOfDisks field to given value.

### HasNumberOfDisks

`func (o *RaidPayload) HasNumberOfDisks() bool`

HasNumberOfDisks returns a boolean if a field has been set.

### SetNumberOfDisksNil

`func (o *RaidPayload) SetNumberOfDisksNil(b bool)`

 SetNumberOfDisksNil sets the value for NumberOfDisks to be an explicit nil

### UnsetNumberOfDisks
`func (o *RaidPayload) UnsetNumberOfDisks()`

UnsetNumberOfDisks ensures that no value is present for NumberOfDisks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


