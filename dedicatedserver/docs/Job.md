# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BareMetalId** | Pointer to **string** | Id of the server (deprecated, use serverId instead) | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**Flow** | Pointer to **string** | Job flow | [optional] 
**IsRunning** | Pointer to **bool** | Describes whether the job is running | [optional] 
**Node** | Pointer to **string** | Node ID for this server | [optional] 
**Payload** | Pointer to [**JobPayload**](JobPayload.md) |  | [optional] 
**Progress** | Pointer to [**Progress**](Progress.md) |  | [optional] 
**ServerId** | Pointer to **string** | Id of the server | [optional] 
**Status** | Pointer to **string** | Status of the job | [optional] 
**Tasks** | Pointer to [**[]Task**](Task.md) | List of tasks in the job | [optional] 
**Type** | Pointer to [**JobType**](JobType.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Update timestamp | [optional] 
**Uuid** | Pointer to **string** | Unique identifier of the job | [optional] 

## Methods

### NewJob

`func NewJob() *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBareMetalId

`func (o *Job) GetBareMetalId() string`

GetBareMetalId returns the BareMetalId field if non-nil, zero value otherwise.

### GetBareMetalIdOk

`func (o *Job) GetBareMetalIdOk() (*string, bool)`

GetBareMetalIdOk returns a tuple with the BareMetalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBareMetalId

`func (o *Job) SetBareMetalId(v string)`

SetBareMetalId sets BareMetalId field to given value.

### HasBareMetalId

`func (o *Job) HasBareMetalId() bool`

HasBareMetalId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Job) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Job) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Job) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Job) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFlow

`func (o *Job) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *Job) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *Job) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *Job) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetIsRunning

`func (o *Job) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *Job) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *Job) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.

### HasIsRunning

`func (o *Job) HasIsRunning() bool`

HasIsRunning returns a boolean if a field has been set.

### GetNode

`func (o *Job) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *Job) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *Job) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *Job) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPayload

`func (o *Job) GetPayload() JobPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Job) GetPayloadOk() (*JobPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Job) SetPayload(v JobPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Job) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetProgress

`func (o *Job) GetProgress() Progress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Job) GetProgressOk() (*Progress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Job) SetProgress(v Progress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Job) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetServerId

`func (o *Job) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *Job) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *Job) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *Job) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetStatus

`func (o *Job) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Job) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTasks

`func (o *Job) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Job) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Job) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *Job) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetType

`func (o *Job) GetType() JobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Job) GetTypeOk() (*JobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Job) SetType(v JobType)`

SetType sets Type field to given value.

### HasType

`func (o *Job) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Job) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Job) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Job) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Job) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUuid

`func (o *Job) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Job) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Job) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Job) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


