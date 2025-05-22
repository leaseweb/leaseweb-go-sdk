# Hdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Amount of disks | [optional] 
**Size** | Pointer to **string** | Disk size | [optional] 

## Methods

### NewHdd

`func NewHdd() *Hdd`

NewHdd instantiates a new Hdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHddWithDefaults

`func NewHddWithDefaults() *Hdd`

NewHddWithDefaults instantiates a new Hdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Hdd) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Hdd) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Hdd) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Hdd) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSize

`func (o *Hdd) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Hdd) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Hdd) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Hdd) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


