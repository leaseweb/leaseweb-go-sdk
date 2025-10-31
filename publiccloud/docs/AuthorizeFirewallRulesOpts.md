# AuthorizeFirewallRulesOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | [**[]AuthorizeRule**](AuthorizeRule.md) | Array of firewall rules to authorize | 

## Methods

### NewAuthorizeFirewallRulesOpts

`func NewAuthorizeFirewallRulesOpts(rules []AuthorizeRule, ) *AuthorizeFirewallRulesOpts`

NewAuthorizeFirewallRulesOpts instantiates a new AuthorizeFirewallRulesOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeFirewallRulesOptsWithDefaults

`func NewAuthorizeFirewallRulesOptsWithDefaults() *AuthorizeFirewallRulesOpts`

NewAuthorizeFirewallRulesOptsWithDefaults instantiates a new AuthorizeFirewallRulesOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *AuthorizeFirewallRulesOpts) GetRules() []AuthorizeRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AuthorizeFirewallRulesOpts) GetRulesOk() (*[]AuthorizeRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AuthorizeFirewallRulesOpts) SetRules(v []AuthorizeRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


