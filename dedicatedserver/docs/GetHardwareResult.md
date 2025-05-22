# GetHardwareResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of this hardware scan result | [optional] 
**ParserVersion** | Pointer to **string** | Version of the parser used for this hardware info | [optional] 
**Result** | Pointer to [**Result**](Result.md) |  | [optional] 
**ScannedAt** | Pointer to **time.Time** | Timestamp of hardware scan, the ISO-8601 format | [optional] 
**ServerId** | Pointer to **string** | Id of the server | [optional] 

## Methods

### NewGetHardwareResult

`func NewGetHardwareResult() *GetHardwareResult`

NewGetHardwareResult instantiates a new GetHardwareResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHardwareResultWithDefaults

`func NewGetHardwareResultWithDefaults() *GetHardwareResult`

NewGetHardwareResultWithDefaults instantiates a new GetHardwareResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetHardwareResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetHardwareResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetHardwareResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetHardwareResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParserVersion

`func (o *GetHardwareResult) GetParserVersion() string`

GetParserVersion returns the ParserVersion field if non-nil, zero value otherwise.

### GetParserVersionOk

`func (o *GetHardwareResult) GetParserVersionOk() (*string, bool)`

GetParserVersionOk returns a tuple with the ParserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserVersion

`func (o *GetHardwareResult) SetParserVersion(v string)`

SetParserVersion sets ParserVersion field to given value.

### HasParserVersion

`func (o *GetHardwareResult) HasParserVersion() bool`

HasParserVersion returns a boolean if a field has been set.

### GetResult

`func (o *GetHardwareResult) GetResult() Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetHardwareResult) GetResultOk() (*Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetHardwareResult) SetResult(v Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetHardwareResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetScannedAt

`func (o *GetHardwareResult) GetScannedAt() time.Time`

GetScannedAt returns the ScannedAt field if non-nil, zero value otherwise.

### GetScannedAtOk

`func (o *GetHardwareResult) GetScannedAtOk() (*time.Time, bool)`

GetScannedAtOk returns a tuple with the ScannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannedAt

`func (o *GetHardwareResult) SetScannedAt(v time.Time)`

SetScannedAt sets ScannedAt field to given value.

### HasScannedAt

`func (o *GetHardwareResult) HasScannedAt() bool`

HasScannedAt returns a boolean if a field has been set.

### GetServerId

`func (o *GetHardwareResult) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetHardwareResult) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetHardwareResult) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetHardwareResult) HasServerId() bool`

HasServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


