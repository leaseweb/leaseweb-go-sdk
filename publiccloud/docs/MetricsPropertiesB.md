# MetricsPropertiesB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to [**[]MetricsValues**](MetricsValues.md) |  | [optional] 
**Unit** | Pointer to **string** | Unit representing bytes | [optional] 

## Methods

### NewMetricsPropertiesB

`func NewMetricsPropertiesB() *MetricsPropertiesB`

NewMetricsPropertiesB instantiates a new MetricsPropertiesB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsPropertiesBWithDefaults

`func NewMetricsPropertiesBWithDefaults() *MetricsPropertiesB`

NewMetricsPropertiesBWithDefaults instantiates a new MetricsPropertiesB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *MetricsPropertiesB) GetValues() []MetricsValues`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MetricsPropertiesB) GetValuesOk() (*[]MetricsValues, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MetricsPropertiesB) SetValues(v []MetricsValues)`

SetValues sets Values field to given value.

### HasValues

`func (o *MetricsPropertiesB) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetUnit

`func (o *MetricsPropertiesB) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricsPropertiesB) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricsPropertiesB) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricsPropertiesB) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


