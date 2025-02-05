# JobList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Jobs** | Pointer to [**[]ServerJob**](ServerJob.md) | An array of jobs | [optional] 

## Methods

### NewJobList

`func NewJobList() *JobList`

NewJobList instantiates a new JobList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobListWithDefaults

`func NewJobListWithDefaults() *JobList`

NewJobListWithDefaults instantiates a new JobList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *JobList) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *JobList) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *JobList) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *JobList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetJobs

`func (o *JobList) GetJobs() []ServerJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *JobList) GetJobsOk() (*[]ServerJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *JobList) SetJobs(v []ServerJob)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *JobList) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


