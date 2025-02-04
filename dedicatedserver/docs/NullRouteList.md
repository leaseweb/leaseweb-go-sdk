# NullRouteList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**NullRoutes** | Pointer to [**[]NullRoute**](NullRoute.md) | An array of server null route events | [optional] 

## Methods

### NewNullRouteList

`func NewNullRouteList() *NullRouteList`

NewNullRouteList instantiates a new NullRouteList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullRouteListWithDefaults

`func NewNullRouteListWithDefaults() *NullRouteList`

NewNullRouteListWithDefaults instantiates a new NullRouteList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *NullRouteList) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NullRouteList) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NullRouteList) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NullRouteList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNullRoutes

`func (o *NullRouteList) GetNullRoutes() []NullRoute`

GetNullRoutes returns the NullRoutes field if non-nil, zero value otherwise.

### GetNullRoutesOk

`func (o *NullRouteList) GetNullRoutesOk() (*[]NullRoute, bool)`

GetNullRoutesOk returns a tuple with the NullRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullRoutes

`func (o *NullRouteList) SetNullRoutes(v []NullRoute)`

SetNullRoutes sets NullRoutes field to given value.

### HasNullRoutes

`func (o *NullRouteList) HasNullRoutes() bool`

HasNullRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


