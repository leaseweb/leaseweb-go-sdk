# GetNullRoutedListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NullRoutes** | [**[]NullRoutedIPv6**](NullRoutedIPv6.md) |  | 
**Metadata** | [**Metadata**](Metadata.md) |  | 

## Methods

### NewGetNullRoutedListResult

`func NewGetNullRoutedListResult(nullRoutes []NullRoutedIPv6, metadata Metadata, ) *GetNullRoutedListResult`

NewGetNullRoutedListResult instantiates a new GetNullRoutedListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNullRoutedListResultWithDefaults

`func NewGetNullRoutedListResultWithDefaults() *GetNullRoutedListResult`

NewGetNullRoutedListResultWithDefaults instantiates a new GetNullRoutedListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullRoutes

`func (o *GetNullRoutedListResult) GetNullRoutes() []NullRoutedIPv6`

GetNullRoutes returns the NullRoutes field if non-nil, zero value otherwise.

### GetNullRoutesOk

`func (o *GetNullRoutedListResult) GetNullRoutesOk() (*[]NullRoutedIPv6, bool)`

GetNullRoutesOk returns a tuple with the NullRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullRoutes

`func (o *GetNullRoutedListResult) SetNullRoutes(v []NullRoutedIPv6)`

SetNullRoutes sets NullRoutes field to given value.


### GetMetadata

`func (o *GetNullRoutedListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetNullRoutedListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetNullRoutedListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


