# GetInvoicesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoices** | Pointer to [**[]Invoice**](Invoice.md) | Array of Invoices | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetInvoicesResult

`func NewGetInvoicesResult() *GetInvoicesResult`

NewGetInvoicesResult instantiates a new GetInvoicesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoicesResultWithDefaults

`func NewGetInvoicesResultWithDefaults() *GetInvoicesResult`

NewGetInvoicesResultWithDefaults instantiates a new GetInvoicesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoices

`func (o *GetInvoicesResult) GetInvoices() []Invoice`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *GetInvoicesResult) GetInvoicesOk() (*[]Invoice, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *GetInvoicesResult) SetInvoices(v []Invoice)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *GetInvoicesResult) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.

### GetMetadata

`func (o *GetInvoicesResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetInvoicesResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetInvoicesResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetInvoicesResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


