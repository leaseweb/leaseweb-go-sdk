# CreateDhcpReservationOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bootfile** | **string** | The URL of PXE boot you want your server to boot from | 
**Hostname** | Pointer to **string** | The hostname for the server | [optional] 

## Methods

### NewCreateDhcpReservationOpts

`func NewCreateDhcpReservationOpts(bootfile string, ) *CreateDhcpReservationOpts`

NewCreateDhcpReservationOpts instantiates a new CreateDhcpReservationOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDhcpReservationOptsWithDefaults

`func NewCreateDhcpReservationOptsWithDefaults() *CreateDhcpReservationOpts`

NewCreateDhcpReservationOptsWithDefaults instantiates a new CreateDhcpReservationOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootfile

`func (o *CreateDhcpReservationOpts) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *CreateDhcpReservationOpts) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *CreateDhcpReservationOpts) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.


### GetHostname

`func (o *CreateDhcpReservationOpts) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CreateDhcpReservationOpts) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CreateDhcpReservationOpts) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CreateDhcpReservationOpts) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


