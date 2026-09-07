# BucketQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **float32** | The value of the quota. | [optional] 
**Unit** | Pointer to **string** | The unit of the quota, e.g., GB. | [optional] 

## Methods

### NewBucketQuota

`func NewBucketQuota() *BucketQuota`

NewBucketQuota instantiates a new BucketQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketQuotaWithDefaults

`func NewBucketQuotaWithDefaults() *BucketQuota`

NewBucketQuotaWithDefaults instantiates a new BucketQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *BucketQuota) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BucketQuota) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BucketQuota) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *BucketQuota) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUnit

`func (o *BucketQuota) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BucketQuota) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BucketQuota) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BucketQuota) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


