# GetExpensesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billing** | [**Billing**](Billing.md) |  | 
**Metadata** | [**ExpensesMetadata**](ExpensesMetadata.md) |  | 

## Methods

### NewGetExpensesResult

`func NewGetExpensesResult(billing Billing, metadata ExpensesMetadata, ) *GetExpensesResult`

NewGetExpensesResult instantiates a new GetExpensesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExpensesResultWithDefaults

`func NewGetExpensesResultWithDefaults() *GetExpensesResult`

NewGetExpensesResultWithDefaults instantiates a new GetExpensesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilling

`func (o *GetExpensesResult) GetBilling() Billing`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *GetExpensesResult) GetBillingOk() (*Billing, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *GetExpensesResult) SetBilling(v Billing)`

SetBilling sets Billing field to given value.


### GetMetadata

`func (o *GetExpensesResult) GetMetadata() ExpensesMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetExpensesResult) GetMetadataOk() (*ExpensesMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetExpensesResult) SetMetadata(v ExpensesMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


