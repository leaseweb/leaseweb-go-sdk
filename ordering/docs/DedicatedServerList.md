# DedicatedServerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**DedicatedServers** | Pointer to [**[]DedicatedServerDetail**](DedicatedServerDetail.md) |  | [optional] 

## Methods

### NewDedicatedServerList

`func NewDedicatedServerList() *DedicatedServerList`

NewDedicatedServerList instantiates a new DedicatedServerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedServerListWithDefaults

`func NewDedicatedServerListWithDefaults() *DedicatedServerList`

NewDedicatedServerListWithDefaults instantiates a new DedicatedServerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *DedicatedServerList) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DedicatedServerList) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DedicatedServerList) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DedicatedServerList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDedicatedServers

`func (o *DedicatedServerList) GetDedicatedServers() []DedicatedServerDetail`

GetDedicatedServers returns the DedicatedServers field if non-nil, zero value otherwise.

### GetDedicatedServersOk

`func (o *DedicatedServerList) GetDedicatedServersOk() (*[]DedicatedServerDetail, bool)`

GetDedicatedServersOk returns a tuple with the DedicatedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedServers

`func (o *DedicatedServerList) SetDedicatedServers(v []DedicatedServerDetail)`

SetDedicatedServers sets DedicatedServers field to given value.

### HasDedicatedServers

`func (o *DedicatedServerList) HasDedicatedServers() bool`

HasDedicatedServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


