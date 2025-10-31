# GetSecurityGroupFirewallRulesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirewallRules** | Pointer to [**[]SecurityGroupFirewallRuleResponse**](SecurityGroupFirewallRuleResponse.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetSecurityGroupFirewallRulesResult

`func NewGetSecurityGroupFirewallRulesResult() *GetSecurityGroupFirewallRulesResult`

NewGetSecurityGroupFirewallRulesResult instantiates a new GetSecurityGroupFirewallRulesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecurityGroupFirewallRulesResultWithDefaults

`func NewGetSecurityGroupFirewallRulesResultWithDefaults() *GetSecurityGroupFirewallRulesResult`

NewGetSecurityGroupFirewallRulesResultWithDefaults instantiates a new GetSecurityGroupFirewallRulesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirewallRules

`func (o *GetSecurityGroupFirewallRulesResult) GetFirewallRules() []SecurityGroupFirewallRuleResponse`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *GetSecurityGroupFirewallRulesResult) GetFirewallRulesOk() (*[]SecurityGroupFirewallRuleResponse, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *GetSecurityGroupFirewallRulesResult) SetFirewallRules(v []SecurityGroupFirewallRuleResponse)`

SetFirewallRules sets FirewallRules field to given value.

### HasFirewallRules

`func (o *GetSecurityGroupFirewallRulesResult) HasFirewallRules() bool`

HasFirewallRules returns a boolean if a field has been set.

### GetMetadata

`func (o *GetSecurityGroupFirewallRulesResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetSecurityGroupFirewallRulesResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetSecurityGroupFirewallRulesResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetSecurityGroupFirewallRulesResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


