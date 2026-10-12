# CreateReportMessageOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | Plain text body of the message, for newlines use &#39; &#39;. | 
**UploadId** | Pointer to **string** | Found in the response after ulpoading a file | [optional] 

## Methods

### NewCreateReportMessageOpts

`func NewCreateReportMessageOpts(body string, ) *CreateReportMessageOpts`

NewCreateReportMessageOpts instantiates a new CreateReportMessageOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReportMessageOptsWithDefaults

`func NewCreateReportMessageOptsWithDefaults() *CreateReportMessageOpts`

NewCreateReportMessageOptsWithDefaults instantiates a new CreateReportMessageOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *CreateReportMessageOpts) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateReportMessageOpts) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateReportMessageOpts) SetBody(v string)`

SetBody sets Body field to given value.


### GetUploadId

`func (o *CreateReportMessageOpts) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *CreateReportMessageOpts) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *CreateReportMessageOpts) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *CreateReportMessageOpts) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


