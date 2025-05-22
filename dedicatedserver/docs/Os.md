# Os

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human readable name for the operating system | [optional] 
**Type** | Pointer to **string** | The type of operating system | [optional] 
**Engine** | Pointer to **string** | Our engine used to install the operating system | [optional] 
**Family** | Pointer to **string** | The operating system family | [optional] 
**Version** | Pointer to **string** | The version of the operating system | [optional] 
**Architecture** | Pointer to **string** | The architecture of the operating system | [optional] 

## Methods

### NewOs

`func NewOs() *Os`

NewOs instantiates a new Os object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsWithDefaults

`func NewOsWithDefaults() *Os`

NewOsWithDefaults instantiates a new Os object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Os) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Os) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Os) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Os) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Os) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Os) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Os) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Os) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEngine

`func (o *Os) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *Os) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *Os) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *Os) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetFamily

`func (o *Os) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *Os) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *Os) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *Os) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetVersion

`func (o *Os) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Os) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Os) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Os) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetArchitecture

`func (o *Os) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Os) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *Os) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *Os) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


