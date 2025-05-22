# Database

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkType** | Pointer to **string** | Network type for the database | [optional] 
**DbName** | Pointer to **string** | Name of the database | [optional] 
**Type** | Pointer to **string** | Type of database | [optional] 

## Methods

### NewDatabase

`func NewDatabase() *Database`

NewDatabase instantiates a new Database object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseWithDefaults

`func NewDatabaseWithDefaults() *Database`

NewDatabaseWithDefaults instantiates a new Database object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkType

`func (o *Database) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *Database) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *Database) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *Database) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetDbName

`func (o *Database) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *Database) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *Database) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *Database) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetType

`func (o *Database) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Database) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Database) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Database) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


