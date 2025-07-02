# InstanceTypesMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | Total amount of elements in this collection | 
**Offset** | **int32** | The offset used to generate this response | [default to 0]
**Limit** | **int32** | The limit used to generate this response | [default to 5]
**Currency** | **string** |  | 
**CurrencySymbol** | **string** |  | 

## Methods

### NewInstanceTypesMetadata

`func NewInstanceTypesMetadata(totalCount int32, offset int32, limit int32, currency string, currencySymbol string, ) *InstanceTypesMetadata`

NewInstanceTypesMetadata instantiates a new InstanceTypesMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypesMetadataWithDefaults

`func NewInstanceTypesMetadataWithDefaults() *InstanceTypesMetadata`

NewInstanceTypesMetadataWithDefaults instantiates a new InstanceTypesMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *InstanceTypesMetadata) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *InstanceTypesMetadata) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *InstanceTypesMetadata) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetOffset

`func (o *InstanceTypesMetadata) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *InstanceTypesMetadata) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *InstanceTypesMetadata) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *InstanceTypesMetadata) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InstanceTypesMetadata) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InstanceTypesMetadata) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetCurrency

`func (o *InstanceTypesMetadata) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InstanceTypesMetadata) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InstanceTypesMetadata) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCurrencySymbol

`func (o *InstanceTypesMetadata) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *InstanceTypesMetadata) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *InstanceTypesMetadata) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


