# \PubliccloudAPI

All URIs are relative to *https://api.leaseweb.com/publicCloud/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToPrivateNetwork**](PubliccloudAPI.md#AddToPrivateNetwork) | **Put** /instances/{instanceId}/addToPrivateNetwork | Add instance to Private Network
[**AttachIso**](PubliccloudAPI.md#AttachIso) | **Post** /instances/{instanceId}/attachIso | Attach ISO to a specific Instance
[**AttachSecurityGroups**](PubliccloudAPI.md#AttachSecurityGroups) | **Post** /instances/{instanceId}/attachSecurityGroups | Attach security groups to instance
[**AuthorizeSecurityGroupFirewallRules**](PubliccloudAPI.md#AuthorizeSecurityGroupFirewallRules) | **Post** /securityGroups/{securityGroupId}/authorizeFirewallRules | Add firewall rules to security group
[**CancelInstanceTermination**](PubliccloudAPI.md#CancelInstanceTermination) | **Post** /instances/{instanceId}/cancelTermination | Cancel instance termination
[**CreateAutoScalingGroup**](PubliccloudAPI.md#CreateAutoScalingGroup) | **Post** /autoScalingGroups | Create Auto Scaling Group
[**CreateImage**](PubliccloudAPI.md#CreateImage) | **Post** /images | Create Custom Image
[**CreateLoadBalancerListener**](PubliccloudAPI.md#CreateLoadBalancerListener) | **Post** /loadBalancers/{loadBalancerId}/listeners | Create listener
[**CreateNotificationSetting**](PubliccloudAPI.md#CreateNotificationSetting) | **Post** /instances/{instanceId}/notificationSettings/dataTraffic/{notificationSettingId} | Create a notification setting
[**CreateSecurityGroup**](PubliccloudAPI.md#CreateSecurityGroup) | **Post** /securityGroups | Create Security Group
[**CreateSnapshot**](PubliccloudAPI.md#CreateSnapshot) | **Post** /instances/{instanceId}/snapshots | Create instance snapshot
[**CreateTargetGroup**](PubliccloudAPI.md#CreateTargetGroup) | **Post** /targetGroups | Create Target Group
[**DeleteAutoScalingGroup**](PubliccloudAPI.md#DeleteAutoScalingGroup) | **Delete** /autoScalingGroups/{autoScalingGroupId} | Delete Auto Scaling Group
[**DeleteCredential**](PubliccloudAPI.md#DeleteCredential) | **Delete** /instances/{instanceId}/credentials/{type}/{username} | Delete Instance credential for a given type and username
[**DeleteCredentials**](PubliccloudAPI.md#DeleteCredentials) | **Delete** /instances/{instanceId}/credentials | Delete all credentials associated with a specific Instance
[**DeleteLoadBalancerListener**](PubliccloudAPI.md#DeleteLoadBalancerListener) | **Delete** /loadBalancers/{loadBalancerId}/listeners/{listenerId} | Delete load balancer listener
[**DeleteNotificationSetting**](PubliccloudAPI.md#DeleteNotificationSetting) | **Delete** /instances/{instanceId}/notificationSettings/dataTraffic/{notificationSettingId} | Delete a notification setting setting
[**DeleteSecurityGroup**](PubliccloudAPI.md#DeleteSecurityGroup) | **Delete** /securityGroups/{securityGroupId} | Delete Security Group
[**DeleteSnapshot**](PubliccloudAPI.md#DeleteSnapshot) | **Delete** /instances/{instanceId}/snapshots/{snapshotId} | Delete instance snapshot
[**DeleteTargetGroup**](PubliccloudAPI.md#DeleteTargetGroup) | **Delete** /targetGroups/{targetGroupId} | Delete Target Group
[**DeregisterAutoScalingGroupTargetGroup**](PubliccloudAPI.md#DeregisterAutoScalingGroupTargetGroup) | **Post** /autoScalingGroups/{autoScalingGroupId}/deregisterTargetGroup | Deregister Target Group
[**DeregisterTargets**](PubliccloudAPI.md#DeregisterTargets) | **Post** /targetGroups/{targetGroupId}/deregisterTargets | Deregister Targets
[**DetachIso**](PubliccloudAPI.md#DetachIso) | **Post** /instances/{instanceId}/detachIso | Detach ISO from a specific Instance
[**DetachSecurityGroups**](PubliccloudAPI.md#DetachSecurityGroups) | **Post** /instances/{instanceId}/detachSecurityGroups | Detach security groups from instance
[**GetAutoScalingGroup**](PubliccloudAPI.md#GetAutoScalingGroup) | **Get** /autoScalingGroups/{autoScalingGroupId} | Get Auto Scaling Group details
[**GetAutoScalingGroupInstanceList**](PubliccloudAPI.md#GetAutoScalingGroupInstanceList) | **Get** /autoScalingGroups/{autoScalingGroupId}/instances | Get list of instances belonging to an Auto Scaling Group
[**GetAutoScalingGroupList**](PubliccloudAPI.md#GetAutoScalingGroupList) | **Get** /autoScalingGroups | Get Auto Scaling Group list
[**GetConnectionsMetrics**](PubliccloudAPI.md#GetConnectionsMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/connections | Get connections metrics
[**GetConnectionsPerSecondMetrics**](PubliccloudAPI.md#GetConnectionsPerSecondMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/connectionsPerSecond | Get connections per second metrics
[**GetConsoleAccess**](PubliccloudAPI.md#GetConsoleAccess) | **Get** /instances/{instanceId}/console | Get console access
[**GetCredential**](PubliccloudAPI.md#GetCredential) | **Get** /instances/{instanceId}/credentials/{type}/{username} | Get Instance credentials by type and username.
[**GetCredentialList**](PubliccloudAPI.md#GetCredentialList) | **Get** /instances/{instanceId}/credentials | List credentials stored for a specific Instance
[**GetCredentialListByType**](PubliccloudAPI.md#GetCredentialListByType) | **Get** /instances/{instanceId}/credentials/{type} | Get credentials by type for a specific Instance
[**GetDataTransferredMetrics**](PubliccloudAPI.md#GetDataTransferredMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/dataTransferred | Get load balancer data transferred metrics
[**GetDataTransferredPerSecondMetrics**](PubliccloudAPI.md#GetDataTransferredPerSecondMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/dataTransferredPerSecond | Get load balancer data transferred per second metrics
[**GetExpenses**](PubliccloudAPI.md#GetExpenses) | **Get** /equipments/{equipmentId}/expenses | Get costs for a given month.
[**GetImageList**](PubliccloudAPI.md#GetImageList) | **Get** /images | List all available Images
[**GetInstance**](PubliccloudAPI.md#GetInstance) | **Get** /instances/{instanceId} | Get instance details
[**GetInstanceCpuMetrics**](PubliccloudAPI.md#GetInstanceCpuMetrics) | **Get** /instances/{instanceId}/metrics/cpu | Get instance CPU metrics
[**GetInstanceDataTrafficMetrics**](PubliccloudAPI.md#GetInstanceDataTrafficMetrics) | **Get** /instances/{instanceId}/metrics/datatraffic | Get data traffic metrics for a specific Instance
[**GetInstanceIP**](PubliccloudAPI.md#GetInstanceIP) | **Get** /instances/{instanceId}/ips/{ip} | Get IP details for a specific Instance
[**GetInstanceIPList**](PubliccloudAPI.md#GetInstanceIPList) | **Get** /instances/{instanceId}/ips | List IP addresses associated with a specific Instance
[**GetInstanceList**](PubliccloudAPI.md#GetInstanceList) | **Get** /instances | Get instance list
[**GetInstanceSecurityGroups**](PubliccloudAPI.md#GetInstanceSecurityGroups) | **Get** /instances/{instanceId}/securityGroups | Get Instance Security Groups
[**GetInstanceTypeList**](PubliccloudAPI.md#GetInstanceTypeList) | **Get** /instanceTypes | List instance types
[**GetIsoList**](PubliccloudAPI.md#GetIsoList) | **Get** /isos | List available ISOs
[**GetLoadBalancer**](PubliccloudAPI.md#GetLoadBalancer) | **Get** /loadBalancers/{loadBalancerId} | Get load balancer details
[**GetLoadBalancerCpuMetrics**](PubliccloudAPI.md#GetLoadBalancerCpuMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/cpu | Get load balancer CPU metrics
[**GetLoadBalancerDataTrafficMetrics**](PubliccloudAPI.md#GetLoadBalancerDataTrafficMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/datatraffic | Get data traffic metrics for a specific Load Balancer
[**GetLoadBalancerIP**](PubliccloudAPI.md#GetLoadBalancerIP) | **Get** /loadBalancers/{loadBalancerId}/ips/{ip} | Get IP details for a specific Load Balancer
[**GetLoadBalancerIPList**](PubliccloudAPI.md#GetLoadBalancerIPList) | **Get** /loadBalancers/{loadBalancerId}/ips | List IP addresses associated with a specific Load Balancer
[**GetLoadBalancerList**](PubliccloudAPI.md#GetLoadBalancerList) | **Get** /loadBalancers | Get load balancer list
[**GetLoadBalancerListener**](PubliccloudAPI.md#GetLoadBalancerListener) | **Get** /loadBalancers/{loadBalancerId}/listeners/{listenerId} | Get listener details
[**GetLoadBalancerListenerList**](PubliccloudAPI.md#GetLoadBalancerListenerList) | **Get** /loadBalancers/{loadBalancerId}/listeners | Get listener list
[**GetMarketAppList**](PubliccloudAPI.md#GetMarketAppList) | **Get** /marketApps | Get marketplace apps
[**GetNotificationSetting**](PubliccloudAPI.md#GetNotificationSetting) | **Get** /instances/{instanceId}/notificationSettings/dataTraffic/{notificationSettingId} | Get details of a notification Setting
[**GetNotificationSettingList**](PubliccloudAPI.md#GetNotificationSettingList) | **Get** /instances/{instanceId}/notificationSettings/dataTraffic | List the notification settings of a customer
[**GetRegionList**](PubliccloudAPI.md#GetRegionList) | **Get** /regions | List regions
[**GetReinstallImageList**](PubliccloudAPI.md#GetReinstallImageList) | **Get** /instances/{instanceId}/reinstall/images | List images available for reinstall
[**GetRequestsMetrics**](PubliccloudAPI.md#GetRequestsMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/requests | Get load balancer requests metrics. Not available for listeners with TCP protocol
[**GetRequestsPerSecondMetrics**](PubliccloudAPI.md#GetRequestsPerSecondMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/requestsPerSecond | Get load balancer requests per second metrics. Not available for listeners with TCP protocol
[**GetResponseCodesMetrics**](PubliccloudAPI.md#GetResponseCodesMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/responseCodes | Get response codes metrics
[**GetResponseCodesPerSecondMetrics**](PubliccloudAPI.md#GetResponseCodesPerSecondMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/responseCodesPerSecond | Get response codes per second metrics
[**GetSecurityGroup**](PubliccloudAPI.md#GetSecurityGroup) | **Get** /securityGroups/{securityGroupId} | Get Security Group details
[**GetSecurityGroupFirewallRules**](PubliccloudAPI.md#GetSecurityGroupFirewallRules) | **Get** /securityGroups/{securityGroupId}/firewallRules | Get Security Group Firewall Rules
[**GetSecurityGroupList**](PubliccloudAPI.md#GetSecurityGroupList) | **Get** /securityGroups | Get Security Group list
[**GetSnapshot**](PubliccloudAPI.md#GetSnapshot) | **Get** /instances/{instanceId}/snapshots/{snapshotId} | Get snapshot detail
[**GetSnapshotList**](PubliccloudAPI.md#GetSnapshotList) | **Get** /instances/{instanceId}/snapshots | List snapshots
[**GetTargetGroup**](PubliccloudAPI.md#GetTargetGroup) | **Get** /targetGroups/{targetGroupId} | Get Target Group details
[**GetTargetGroupList**](PubliccloudAPI.md#GetTargetGroupList) | **Get** /targetGroups | Get Target Group list
[**GetTargetList**](PubliccloudAPI.md#GetTargetList) | **Get** /targetGroups/{targetGroupId}/targets | Get Targets
[**GetUpdateInstanceTypeList**](PubliccloudAPI.md#GetUpdateInstanceTypeList) | **Get** /instances/{instanceId}/instanceTypesUpdate | List available instance types for update
[**LaunchInstance**](PubliccloudAPI.md#LaunchInstance) | **Post** /instances | Launch instance
[**LaunchLoadBalancer**](PubliccloudAPI.md#LaunchLoadBalancer) | **Post** /loadBalancers | Launch Load balancer
[**NullRouteInstanceIP**](PubliccloudAPI.md#NullRouteInstanceIP) | **Post** /instances/{instanceId}/ips/{ip}/null | Null route IP address for a specific resource Instance
[**NullRouteLoadBalancerIP**](PubliccloudAPI.md#NullRouteLoadBalancerIP) | **Post** /loadBalancers/{loadBalancerId}/ips/{ip}/null | Null route IP address for a specific resource Load Balancer
[**RebootInstance**](PubliccloudAPI.md#RebootInstance) | **Post** /instances/{instanceId}/reboot | Reboot a specific Instance
[**RebootLoadBalancer**](PubliccloudAPI.md#RebootLoadBalancer) | **Post** /loadBalancers/{loadBalancerId}/reboot | Reboot a specific Load Balancer
[**RegisterAutoScalingGroupTargetGroup**](PubliccloudAPI.md#RegisterAutoScalingGroupTargetGroup) | **Post** /autoScalingGroups/{autoScalingGroupId}/registerTargetGroup | Register Target Group
[**RegisterTargets**](PubliccloudAPI.md#RegisterTargets) | **Post** /targetGroups/{targetGroupId}/registerTargets | Register Targets
[**ReinstallInstance**](PubliccloudAPI.md#ReinstallInstance) | **Put** /instances/{instanceId}/reinstall | Reinstall an Instance
[**RemoveFromPrivateNetwork**](PubliccloudAPI.md#RemoveFromPrivateNetwork) | **Delete** /instances/{instanceId}/removeFromPrivateNetwork | Remove instance from Private Network
[**RemoveInstanceIPNullRoute**](PubliccloudAPI.md#RemoveInstanceIPNullRoute) | **Post** /instances/{instanceId}/ips/{ip}/unnull | Remove an IP null route for a specific Instance
[**RemoveLoadBalancerIPNullRoute**](PubliccloudAPI.md#RemoveLoadBalancerIPNullRoute) | **Post** /loadBalancers/{loadBalancerId}/ips/{ip}/unnull | Remove an IP null route for a specific Load Balancer
[**ResetPassword**](PubliccloudAPI.md#ResetPassword) | **Post** /instances/{instanceId}/resetPassword | Reset the password for a specific Instance
[**RestoreSnapshot**](PubliccloudAPI.md#RestoreSnapshot) | **Put** /instances/{instanceId}/snapshots/{snapshotId} | Restore instance snapshot
[**RevokeSecurityGroupFirewallRules**](PubliccloudAPI.md#RevokeSecurityGroupFirewallRules) | **Post** /securityGroups/{securityGroupId}/revokeFirewallRules | Remove firewall rules from security group
[**StartInstance**](PubliccloudAPI.md#StartInstance) | **Post** /instances/{instanceId}/start | Start a specific resource Instance
[**StartLoadBalancer**](PubliccloudAPI.md#StartLoadBalancer) | **Post** /loadBalancers/{loadBalancerId}/start | Start a specific resource Load Balancer
[**StopInstance**](PubliccloudAPI.md#StopInstance) | **Post** /instances/{instanceId}/stop | Stop a specific Instance
[**StopLoadBalancer**](PubliccloudAPI.md#StopLoadBalancer) | **Post** /loadBalancers/{loadBalancerId}/stop | Stop a specific Load Balancer
[**StoreCredential**](PubliccloudAPI.md#StoreCredential) | **Post** /instances/{instanceId}/credentials | Store credentials for a specific Instance
[**TerminateInstance**](PubliccloudAPI.md#TerminateInstance) | **Delete** /instances/{instanceId} | Terminate instance
[**TerminateLoadBalancer**](PubliccloudAPI.md#TerminateLoadBalancer) | **Delete** /loadBalancers/{loadBalancerId} | Delete load balancer
[**UpdateAutoScalingGroup**](PubliccloudAPI.md#UpdateAutoScalingGroup) | **Put** /autoScalingGroups/{autoScalingGroupId} | Update Auto Scaling Group
[**UpdateCredential**](PubliccloudAPI.md#UpdateCredential) | **Put** /instances/{instanceId}/credentials/{type}/{username} | Update credentials for a given type and username
[**UpdateImage**](PubliccloudAPI.md#UpdateImage) | **Put** /images/{imageId} | Update Custom Image
[**UpdateInstance**](PubliccloudAPI.md#UpdateInstance) | **Put** /instances/{instanceId} | Update instance
[**UpdateInstanceIP**](PubliccloudAPI.md#UpdateInstanceIP) | **Put** /instances/{instanceId}/ips/{ip} | Update the IP address for a specific Instance
[**UpdateLoadBalancer**](PubliccloudAPI.md#UpdateLoadBalancer) | **Put** /loadBalancers/{loadBalancerId} | Update load balancer
[**UpdateLoadBalancerIP**](PubliccloudAPI.md#UpdateLoadBalancerIP) | **Put** /loadBalancers/{loadBalancerId}/ips/{ip} | Update the IP address for a specific Load Balancer
[**UpdateLoadBalancerListener**](PubliccloudAPI.md#UpdateLoadBalancerListener) | **Put** /loadBalancers/{loadBalancerId}/listeners/{listenerId} | Update a listener
[**UpdateNotificationSetting**](PubliccloudAPI.md#UpdateNotificationSetting) | **Put** /instances/{instanceId}/notificationSettings/dataTraffic/{notificationSettingId} | Update Notification Setting details
[**UpdateSecurityGroup**](PubliccloudAPI.md#UpdateSecurityGroup) | **Put** /securityGroups/{securityGroupId} | Update Security Group
[**UpdateTargetGroup**](PubliccloudAPI.md#UpdateTargetGroup) | **Put** /targetGroups/{targetGroupId} | Update Target Group



## AddToPrivateNetwork

> AddToPrivateNetwork(ctx, instanceId).Execute()

Add instance to Private Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.AddToPrivateNetwork(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.AddToPrivateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddToPrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachIso

> AttachIso(ctx, instanceId).AttachIsoOpts(attachIsoOpts).Execute()

Attach ISO to a specific Instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	attachIsoOpts := *openapiclient.NewAttachIsoOpts("IsoId_example") // AttachIsoOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.AttachIso(context.Background(), instanceId).AttachIsoOpts(attachIsoOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.AttachIso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachIsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachIsoOpts** | [**AttachIsoOpts**](AttachIsoOpts.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachSecurityGroups

> AttachSecurityGroups(ctx, instanceId).AttachSecurityGroupsOpts(attachSecurityGroupsOpts).Execute()

Attach security groups to instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	attachSecurityGroupsOpts := *openapiclient.NewAttachSecurityGroupsOpts([]string{"SecurityGroupIds_example"}) // AttachSecurityGroupsOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.AttachSecurityGroups(context.Background(), instanceId).AttachSecurityGroupsOpts(attachSecurityGroupsOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.AttachSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachSecurityGroupsOpts** | [**AttachSecurityGroupsOpts**](AttachSecurityGroupsOpts.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizeSecurityGroupFirewallRules

> AuthorizeSecurityGroupFirewallRules(ctx, securityGroupId).AuthorizeFirewallRulesOpts(authorizeFirewallRulesOpts).Execute()

Add firewall rules to security group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	securityGroupId := "166cc685-8273-11f0-b4f8-2687367a8a5e" // string | Security Group ID
	authorizeFirewallRulesOpts := *openapiclient.NewAuthorizeFirewallRulesOpts([]openapiclient.AuthorizeRule{*openapiclient.NewAuthorizeRule("Protocol_example")}) // AuthorizeFirewallRulesOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.AuthorizeSecurityGroupFirewallRules(context.Background(), securityGroupId).AuthorizeFirewallRulesOpts(authorizeFirewallRulesOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.AuthorizeSecurityGroupFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | Security Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeSecurityGroupFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizeFirewallRulesOpts** | [**AuthorizeFirewallRulesOpts**](AuthorizeFirewallRulesOpts.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelInstanceTermination

> CancelInstanceTermination(ctx, instanceId).Execute()

Cancel instance termination



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.CancelInstanceTermination(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CancelInstanceTermination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelInstanceTerminationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAutoScalingGroup

> AutoScalingGroupDetails CreateAutoScalingGroup(ctx).CreateAutoScalingGroupOpts(createAutoScalingGroupOpts).Execute()

Create Auto Scaling Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	createAutoScalingGroupOpts := *openapiclient.NewCreateAutoScalingGroupOpts("InstanceId_example", "Reference_example", "Type_example") // CreateAutoScalingGroupOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.CreateAutoScalingGroup(context.Background()).CreateAutoScalingGroupOpts(createAutoScalingGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CreateAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutoScalingGroup`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.CreateAutoScalingGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoScalingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAutoScalingGroupOpts** | [**CreateAutoScalingGroupOpts**](CreateAutoScalingGroupOpts.md) |  | 

### Return type

[**AutoScalingGroupDetails**](AutoScalingGroupDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateImage

> ImageDetails CreateImage(ctx).CreateImageOpts(createImageOpts).Execute()

Create Custom Image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	createImageOpts := *openapiclient.NewCreateImageOpts("Name_example", "InstanceId_example") // CreateImageOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.CreateImage(context.Background()).CreateImageOpts(createImageOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CreateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateImage`: ImageDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.CreateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createImageOpts** | [**CreateImageOpts**](CreateImageOpts.md) |  | 

### Return type

[**ImageDetails**](ImageDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoadBalancerListener

> LoadBalancerListener CreateLoadBalancerListener(ctx, loadBalancerId).LoadBalancerListenerCreateOpts(loadBalancerListenerCreateOpts).Execute()

Create listener



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	loadBalancerListenerCreateOpts := *openapiclient.NewLoadBalancerListenerCreateOpts(openapiclient.protocol("HTTP"), int32(123), *openapiclient.NewLoadBalancerListenerDefaultRule("TargetGroupId_example")) // LoadBalancerListenerCreateOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.CreateLoadBalancerListener(context.Background(), loadBalancerId).LoadBalancerListenerCreateOpts(loadBalancerListenerCreateOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CreateLoadBalancerListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoadBalancerListener`: LoadBalancerListener
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.CreateLoadBalancerListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loadBalancerListenerCreateOpts** | [**LoadBalancerListenerCreateOpts**](LoadBalancerListenerCreateOpts.md) |  | 

### Return type

[**LoadBalancerListener**](LoadBalancerListener.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationSetting

> NotificationSetting CreateNotificationSetting(ctx, instanceId, notificationSettingId).CreateNotificationSettingOpts(createNotificationSettingOpts).Execute()

Create a notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	notificationSettingId := "3a042956-0689-45dc-8322-8b8325464182" // string | Notification Setting ID
	createNotificationSettingOpts := *openapiclient.NewCreateNotificationSettingOpts(*openapiclient.NewNotificationSettingThreshold(int32(123), openapiclient.unit("MB")), openapiclient.timePeriod("DAY"), "TODO", []openapiclient.UpdateNotificationSettingOptsChannelsInner{*openapiclient.NewUpdateNotificationSettingOptsChannelsInner()}) // CreateNotificationSettingOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.CreateNotificationSetting(context.Background(), instanceId, notificationSettingId).CreateNotificationSettingOpts(createNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CreateNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotificationSetting`: NotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.CreateNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**notificationSettingId** | **string** | Notification Setting ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNotificationSettingOpts** | [**CreateNotificationSettingOpts**](CreateNotificationSettingOpts.md) |  | 

### Return type

[**NotificationSetting**](NotificationSetting.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecurityGroup

> SecurityGroup CreateSecurityGroup(ctx).CreateSecurityGroupOpts(createSecurityGroupOpts).Execute()

Create Security Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	createSecurityGroupOpts := *openapiclient.NewCreateSecurityGroupOpts("Name_example") // CreateSecurityGroupOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.CreateSecurityGroup(context.Background()).CreateSecurityGroupOpts(createSecurityGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CreateSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecurityGroup`: SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.CreateSecurityGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSecurityGroupOpts** | [**CreateSecurityGroupOpts**](CreateSecurityGroupOpts.md) |  | 

### Return type

[**SecurityGroup**](SecurityGroup.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSnapshot

> CreateSnapshot(ctx, instanceId).CreateSnapshotOpts(createSnapshotOpts).Execute()

Create instance snapshot



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	createSnapshotOpts := *openapiclient.NewCreateSnapshotOpts("Name_example") // CreateSnapshotOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.CreateSnapshot(context.Background(), instanceId).CreateSnapshotOpts(createSnapshotOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CreateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSnapshotOpts** | [**CreateSnapshotOpts**](CreateSnapshotOpts.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTargetGroup

> TargetGroup CreateTargetGroup(ctx).CreateTargetGroupOpts(createTargetGroupOpts).Execute()

Create Target Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	createTargetGroupOpts := *openapiclient.NewCreateTargetGroupOpts("Name_example", openapiclient.protocol("HTTP"), int32(123), openapiclient.regionName("eu-west-3")) // CreateTargetGroupOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.CreateTargetGroup(context.Background()).CreateTargetGroupOpts(createTargetGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CreateTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTargetGroup`: TargetGroup
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.CreateTargetGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTargetGroupOpts** | [**CreateTargetGroupOpts**](CreateTargetGroupOpts.md) |  | 

### Return type

[**TargetGroup**](TargetGroup.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoScalingGroup

> DeleteAutoScalingGroup(ctx, autoScalingGroupId).Execute()

Delete Auto Scaling Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteAutoScalingGroup(context.Background(), autoScalingGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoScalingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredential

> DeleteCredential(ctx, instanceId, type_, username).Execute()

Delete Instance credential for a given type and username

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	type_ := openapiclient.credentialType("OPERATING_SYSTEM") // CredentialType | Type of credential
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteCredential(context.Background(), instanceId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**type_** | [**CredentialType**](.md) | Type of credential | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentials

> DeleteCredentials(ctx, instanceId).Execute()

Delete all credentials associated with a specific Instance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteCredentials(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadBalancerListener

> DeleteLoadBalancerListener(ctx, loadBalancerId, listenerId).Execute()

Delete load balancer listener



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	listenerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Listener ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteLoadBalancerListener(context.Background(), loadBalancerId, listenerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteLoadBalancerListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 
**listenerId** | **string** | Listener ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationSetting

> DeleteNotificationSetting(ctx, instanceId, notificationSettingId).Execute()

Delete a notification setting setting

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	notificationSettingId := "3a042956-0689-45dc-8322-8b8325464182" // string | Notification Setting ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteNotificationSetting(context.Background(), instanceId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**notificationSettingId** | **string** | Notification Setting ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityGroup

> DeleteSecurityGroup(ctx, securityGroupId).Execute()

Delete Security Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	securityGroupId := "166cc685-8273-11f0-b4f8-2687367a8a5e" // string | Security Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteSecurityGroup(context.Background(), securityGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | Security Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSnapshot

> DeleteSnapshot(ctx, instanceId, snapshotId).Execute()

Delete instance snapshot

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	snapshotId := "54eee4c2-aef2-4cda-9fd5-74bfbf438abf" // string | Snapshot's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteSnapshot(context.Background(), instanceId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**snapshotId** | **string** | Snapshot&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTargetGroup

> DeleteTargetGroup(ctx, targetGroupId).Execute()

Delete Target Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	targetGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Target Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteTargetGroup(context.Background(), targetGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | Target Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeregisterAutoScalingGroupTargetGroup

> AutoScalingGroupDetails DeregisterAutoScalingGroupTargetGroup(ctx, autoScalingGroupId).TargetGroupIdOpts(targetGroupIdOpts).Execute()

Deregister Target Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID
	targetGroupIdOpts := *openapiclient.NewTargetGroupIdOpts("TargetGroupId_example") // TargetGroupIdOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.DeregisterAutoScalingGroupTargetGroup(context.Background(), autoScalingGroupId).TargetGroupIdOpts(targetGroupIdOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeregisterAutoScalingGroupTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeregisterAutoScalingGroupTargetGroup`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.DeregisterAutoScalingGroupTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeregisterAutoScalingGroupTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetGroupIdOpts** | [**TargetGroupIdOpts**](TargetGroupIdOpts.md) |  | 

### Return type

[**AutoScalingGroupDetails**](AutoScalingGroupDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeregisterTargets

> DeregisterTargets(ctx, targetGroupId).RequestBody(requestBody).Execute()

Deregister Targets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	targetGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Target Group ID
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeregisterTargets(context.Background(), targetGroupId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeregisterTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | Target Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeregisterTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachIso

> DetachIso(ctx, instanceId).Execute()

Detach ISO from a specific Instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DetachIso(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DetachIso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachIsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachSecurityGroups

> DetachSecurityGroups(ctx, instanceId).DetachSecurityGroupsOpts(detachSecurityGroupsOpts).Execute()

Detach security groups from instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	detachSecurityGroupsOpts := *openapiclient.NewDetachSecurityGroupsOpts([]string{"SecurityGroupIds_example"}) // DetachSecurityGroupsOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DetachSecurityGroups(context.Background(), instanceId).DetachSecurityGroupsOpts(detachSecurityGroupsOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DetachSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detachSecurityGroupsOpts** | [**DetachSecurityGroupsOpts**](DetachSecurityGroupsOpts.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoScalingGroup

> AutoScalingGroupDetails GetAutoScalingGroup(ctx, autoScalingGroupId).Execute()

Get Auto Scaling Group details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetAutoScalingGroup(context.Background(), autoScalingGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoScalingGroup`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetAutoScalingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoScalingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoScalingGroupDetails**](AutoScalingGroupDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoScalingGroupInstanceList

> GetAutoScalingGroupInstanceListResult GetAutoScalingGroupInstanceList(ctx, autoScalingGroupId).Limit(limit).Offset(offset).Execute()

Get list of instances belonging to an Auto Scaling Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetAutoScalingGroupInstanceList(context.Background(), autoScalingGroupId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetAutoScalingGroupInstanceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoScalingGroupInstanceList`: GetAutoScalingGroupInstanceListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetAutoScalingGroupInstanceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoScalingGroupInstanceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetAutoScalingGroupInstanceListResult**](GetAutoScalingGroupInstanceListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoScalingGroupList

> GetAutoScalingGroupListResult GetAutoScalingGroupList(ctx).Limit(limit).Offset(offset).Id(id).InstanceId(instanceId).Type_(type_).Region(region).Reference(reference).State(state).Execute()

Get Auto Scaling Group list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	id := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" // string |  (optional)
	instanceId := "6762182e-7ae9-4d0b-b3b7-bea5a49b3f94" // string |  (optional)
	type_ := "type__example" // string | The Auto Scaling Group's type (optional)
	region := openapiclient.regionName("eu-west-3") // RegionName | The region in which the Auto Scaling Group was created (optional)
	reference := "reference_example" // string | The reference used to identify identifies the Auto Scaling Group (optional)
	state := "state_example" // string | The Auto Scaling Group's current state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetAutoScalingGroupList(context.Background()).Limit(limit).Offset(offset).Id(id).InstanceId(instanceId).Type_(type_).Region(region).Reference(reference).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetAutoScalingGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoScalingGroupList`: GetAutoScalingGroupListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetAutoScalingGroupList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoScalingGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **id** | **string** |  | 
 **instanceId** | **string** |  | 
 **type_** | **string** | The Auto Scaling Group&#39;s type | 
 **region** | [**RegionName**](RegionName.md) | The region in which the Auto Scaling Group was created | 
 **reference** | **string** | The reference used to identify identifies the Auto Scaling Group | 
 **state** | **string** | The Auto Scaling Group&#39;s current state | 

### Return type

[**GetAutoScalingGroupListResult**](GetAutoScalingGroupListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionsMetrics

> GetConnectionsMetricsResult GetConnectionsMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get connections metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from`
	granularity := openapiclient.metricGranularity("5m") // MetricGranularity | The interval for each metric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetConnectionsMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetConnectionsMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionsMetrics`: GetConnectionsMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetConnectionsMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricGranularity**](MetricGranularity.md) | The interval for each metric | 

### Return type

[**GetConnectionsMetricsResult**](GetConnectionsMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionsPerSecondMetrics

> GetConnectionsPerSecondMetricsResult GetConnectionsPerSecondMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get connections per second metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from`
	granularity := openapiclient.metricGranularity("5m") // MetricGranularity | The interval for each metric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetConnectionsPerSecondMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetConnectionsPerSecondMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionsPerSecondMetrics`: GetConnectionsPerSecondMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetConnectionsPerSecondMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsPerSecondMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricGranularity**](MetricGranularity.md) | The interval for each metric | 

### Return type

[**GetConnectionsPerSecondMetricsResult**](GetConnectionsPerSecondMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsoleAccess

> GetConsoleAccessResult GetConsoleAccess(ctx, instanceId).Execute()

Get console access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetConsoleAccess(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetConsoleAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConsoleAccess`: GetConsoleAccessResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetConsoleAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsoleAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetConsoleAccessResult**](GetConsoleAccessResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredential

> GetCredentialResult GetCredential(ctx, instanceId, type_, username).Execute()

Get Instance credentials by type and username.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	type_ := openapiclient.credentialType("OPERATING_SYSTEM") // CredentialType | Type of credential
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetCredential(context.Background(), instanceId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredential`: GetCredentialResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**type_** | [**CredentialType**](.md) | Type of credential | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetCredentialResult**](GetCredentialResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentialList

> GetCredentialListResult GetCredentialList(ctx, instanceId).Execute()

List credentials stored for a specific Instance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetCredentialList(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetCredentialList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialList`: GetCredentialListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCredentialListResult**](GetCredentialListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentialListByType

> GetCredentialListByTypeResult GetCredentialListByType(ctx, instanceId, type_).Execute()

Get credentials by type for a specific Instance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	type_ := openapiclient.credentialType("OPERATING_SYSTEM") // CredentialType | Type of credential

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetCredentialListByType(context.Background(), instanceId, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetCredentialListByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialListByType`: GetCredentialListByTypeResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetCredentialListByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**type_** | [**CredentialType**](.md) | Type of credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialListByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCredentialListByTypeResult**](GetCredentialListByTypeResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTransferredMetrics

> GetDataTransferredMetricsResult GetDataTransferredMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get load balancer data transferred metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from`
	granularity := openapiclient.metricGranularity("5m") // MetricGranularity | The interval for each metric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetDataTransferredMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetDataTransferredMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTransferredMetrics`: GetDataTransferredMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetDataTransferredMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTransferredMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricGranularity**](MetricGranularity.md) | The interval for each metric | 

### Return type

[**GetDataTransferredMetricsResult**](GetDataTransferredMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTransferredPerSecondMetrics

> GetDataTransferredPerSecondMetricsResult GetDataTransferredPerSecondMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get load balancer data transferred per second metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from`
	granularity := openapiclient.metricGranularity("5m") // MetricGranularity | The interval for each metric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetDataTransferredPerSecondMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetDataTransferredPerSecondMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTransferredPerSecondMetrics`: GetDataTransferredPerSecondMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetDataTransferredPerSecondMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTransferredPerSecondMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricGranularity**](MetricGranularity.md) | The interval for each metric | 

### Return type

[**GetDataTransferredPerSecondMetricsResult**](GetDataTransferredPerSecondMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenses

> GetExpensesResult GetExpenses(ctx, equipmentId).From(from).To(to).Execute()

Get costs for a given month.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	equipmentId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Equipment's UUID
	from := time.Now() // string | Start date of the period to get costs for. It must be the first day of the month
	to := time.Now() // string | End date of the period to get costs for. This date needs to be exactly one month after the 'from' date. If this value is not passed, it will be calculated based on 'from' parameter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetExpenses(context.Background(), equipmentId).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetExpenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpenses`: GetExpensesResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetExpenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**equipmentId** | **string** | Equipment&#39;s UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | Start date of the period to get costs for. It must be the first day of the month | 
 **to** | **string** | End date of the period to get costs for. This date needs to be exactly one month after the &#39;from&#39; date. If this value is not passed, it will be calculated based on &#39;from&#39; parameter | 

### Return type

[**GetExpensesResult**](GetExpensesResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageList

> GetImageListResult GetImageList(ctx).Limit(limit).Offset(offset).Custom(custom).Standard(standard).State(state).MarketAppId(marketAppId).StorageType(storageType).Name(name).Flavour(flavour).Region(region).Execute()

List all available Images

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	custom := true // bool | Filters the list to include only custom images. (optional)
	standard := true // bool | Filters the list to include only standard images. (optional)
	state := openapiclient.imageState("CREATING") // ImageState | Filters the list by the state of custom images. (optional)
	marketAppId := openapiclient.marketAppId("CPANEL_30") // MarketAppId | Filters the list by the market app of standard images. (optional)
	storageType := openapiclient.storageType("LOCAL") // StorageType | Filters the list by the market app of standard images. (optional)
	name := "Ubuntu 20.04 LTS (x86_64)" // string | Filters the list by the name of images. (optional)
	flavour := openapiclient.flavour("ubuntu") // Flavour | Filters the list by the flavour of standard images. (optional)
	region := "eu-west-3" // string | Available regions can be found using the List Regions endpoint. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetImageList(context.Background()).Limit(limit).Offset(offset).Custom(custom).Standard(standard).State(state).MarketAppId(marketAppId).StorageType(storageType).Name(name).Flavour(flavour).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetImageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageList`: GetImageListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetImageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **custom** | **bool** | Filters the list to include only custom images. | 
 **standard** | **bool** | Filters the list to include only standard images. | 
 **state** | [**ImageState**](ImageState.md) | Filters the list by the state of custom images. | 
 **marketAppId** | [**MarketAppId**](MarketAppId.md) | Filters the list by the market app of standard images. | 
 **storageType** | [**StorageType**](StorageType.md) | Filters the list by the market app of standard images. | 
 **name** | **string** | Filters the list by the name of images. | 
 **flavour** | [**Flavour**](Flavour.md) | Filters the list by the flavour of standard images. | 
 **region** | **string** | Available regions can be found using the List Regions endpoint. | 

### Return type

[**GetImageListResult**](GetImageListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> InstanceDetails GetInstance(ctx, instanceId).Execute()

Get instance details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance`: InstanceDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstanceDetails**](InstanceDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceCpuMetrics

> GetCpuMetricsResult GetInstanceCpuMetrics(ctx, instanceId).From(from).To(to).Granularity(granularity).Execute()

Get instance CPU metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	from := time.Now() // string | The start of the interval to get the metric
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from`
	granularity := openapiclient.metricGranularity("5m") // MetricGranularity | The interval for each metric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetInstanceCpuMetrics(context.Background(), instanceId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetInstanceCpuMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceCpuMetrics`: GetCpuMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetInstanceCpuMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceCpuMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricGranularity**](MetricGranularity.md) | The interval for each metric | 

### Return type

[**GetCpuMetricsResult**](GetCpuMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceDataTrafficMetrics

> GetDataTrafficMetricsResult GetInstanceDataTrafficMetrics(ctx, instanceId).From(from).To(to).Granularity(granularity).Aggregation(aggregation).Execute()

Get data traffic metrics for a specific Instance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	from := time.Now() // string | The start of the interval to get the metric
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from`
	granularity := openapiclient.dataTrafficMetricsGranularity("DAY") // DataTrafficMetricsGranularity | The interval for each metric
	aggregation := openapiclient.metricAggregation("SUM") // MetricAggregation | The metric aggregation function

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetInstanceDataTrafficMetrics(context.Background(), instanceId).From(from).To(to).Granularity(granularity).Aggregation(aggregation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetInstanceDataTrafficMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceDataTrafficMetrics`: GetDataTrafficMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetInstanceDataTrafficMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceDataTrafficMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**DataTrafficMetricsGranularity**](DataTrafficMetricsGranularity.md) | The interval for each metric | 
 **aggregation** | [**MetricAggregation**](MetricAggregation.md) | The metric aggregation function | 

### Return type

[**GetDataTrafficMetricsResult**](GetDataTrafficMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceIP

> IpDetails GetInstanceIP(ctx, instanceId, ip).Execute()

Get IP details for a specific Instance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	ip := "10.0.0.1" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetInstanceIP(context.Background(), instanceId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetInstanceIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceIP`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetInstanceIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IpDetails**](IpDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceIPList

> GetIPListResult GetInstanceIPList(ctx, instanceId).Version(version).NullRouted(nullRouted).Ips(ips).Execute()

List IP addresses associated with a specific Instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	version := openapiclient.ipVersion(4) // IpVersion |  (optional)
	nullRouted := true // bool |  (optional)
	ips := "ips_example" // string | A list of IPs separated by `|` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetInstanceIPList(context.Background(), instanceId).Version(version).NullRouted(nullRouted).Ips(ips).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetInstanceIPList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceIPList`: GetIPListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetInstanceIPList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceIPListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**IpVersion**](IpVersion.md) |  | 
 **nullRouted** | **bool** |  | 
 **ips** | **string** | A list of IPs separated by &#x60;|&#x60; | 

### Return type

[**GetIPListResult**](GetIPListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceList

> InstanceList GetInstanceList(ctx).Limit(limit).Offset(offset).Ip(ip).Reference(reference).Id(id).ContractType(contractType).ContractState(contractState).ImageId(imageId).State(state).Region(region).Type_(type_).SecurityGroupIds(securityGroupIds).Execute()

Get instance list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	ip := "10.0.0.1" // string |  (optional)
	reference := "my-webserver" // string |  (optional)
	id := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" // string |  (optional)
	contractType := openapiclient.contractType("HOURLY") // ContractType |  (optional)
	contractState := openapiclient.contractState("ACTIVE") // ContractState |  (optional)
	imageId := "UBUNTU_22_04_64BIT" // string | Available Images can be obtained using `/v1/images`. (optional)
	state := openapiclient.state("CREATING") // State | The instance's current state(s), separated by commas. (optional)
	region := openapiclient.regionName("eu-west-3") // RegionName | Available regions can be obtained using `/v1/regions` (optional)
	type_ := openapiclient.typeName("lsw.m3.large") // TypeName | Available instance types for your region can be obtained using `/v1/instanceTypes`. (optional)
	securityGroupIds := "166cc685-8273-11f0-b4f8-2687367a8a5e,7e59b33d-05f3-4078-b251-c7831ae8fe14" // string | Filter instances by security group UUIDs, separated by commas. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetInstanceList(context.Background()).Limit(limit).Offset(offset).Ip(ip).Reference(reference).Id(id).ContractType(contractType).ContractState(contractState).ImageId(imageId).State(state).Region(region).Type_(type_).SecurityGroupIds(securityGroupIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetInstanceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceList`: InstanceList
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetInstanceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **ip** | **string** |  | 
 **reference** | **string** |  | 
 **id** | **string** |  | 
 **contractType** | [**ContractType**](ContractType.md) |  | 
 **contractState** | [**ContractState**](ContractState.md) |  | 
 **imageId** | **string** | Available Images can be obtained using &#x60;/v1/images&#x60;. | 
 **state** | [**State**](State.md) | The instance&#39;s current state(s), separated by commas. | 
 **region** | [**RegionName**](RegionName.md) | Available regions can be obtained using &#x60;/v1/regions&#x60; | 
 **type_** | [**TypeName**](TypeName.md) | Available instance types for your region can be obtained using &#x60;/v1/instanceTypes&#x60;. | 
 **securityGroupIds** | **string** | Filter instances by security group UUIDs, separated by commas. | 

### Return type

[**InstanceList**](InstanceList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceSecurityGroups

> InstanceSecurityGroups GetInstanceSecurityGroups(ctx, instanceId).Limit(limit).Offset(offset).Execute()

Get Instance Security Groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetInstanceSecurityGroups(context.Background(), instanceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetInstanceSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceSecurityGroups`: InstanceSecurityGroups
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetInstanceSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**InstanceSecurityGroups**](InstanceSecurityGroups.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceTypeList

> InstanceTypes GetInstanceTypeList(ctx).Region(region).Limit(limit).Offset(offset).Execute()

List instance types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	region := openapiclient.regionName("eu-west-3") // RegionName | 
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetInstanceTypeList(context.Background()).Region(region).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetInstanceTypeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceTypeList`: InstanceTypes
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetInstanceTypeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceTypeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | [**RegionName**](RegionName.md) |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**InstanceTypes**](InstanceTypes.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIsoList

> GetIsoListResult GetIsoList(ctx).Limit(limit).Offset(offset).Execute()

List available ISOs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetIsoList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetIsoList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIsoList`: GetIsoListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetIsoList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIsoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetIsoListResult**](GetIsoListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancer

> LoadBalancerDetails GetLoadBalancer(ctx, loadBalancerId).Execute()

Get load balancer details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetLoadBalancer(context.Background(), loadBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancer`: LoadBalancerDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoadBalancerDetails**](LoadBalancerDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerCpuMetrics

> GetCpuMetricsResult GetLoadBalancerCpuMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get load balancer CPU metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from`
	granularity := openapiclient.metricGranularity("5m") // MetricGranularity | The interval for each metric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetLoadBalancerCpuMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetLoadBalancerCpuMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancerCpuMetrics`: GetCpuMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetLoadBalancerCpuMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerCpuMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricGranularity**](MetricGranularity.md) | The interval for each metric | 

### Return type

[**GetCpuMetricsResult**](GetCpuMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerDataTrafficMetrics

> GetDataTrafficMetricsResult GetLoadBalancerDataTrafficMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Aggregation(aggregation).Execute()

Get data traffic metrics for a specific Load Balancer

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from`
	granularity := openapiclient.dataTrafficMetricsGranularity("DAY") // DataTrafficMetricsGranularity | The interval for each metric
	aggregation := openapiclient.metricAggregation("SUM") // MetricAggregation | The metric aggregation function

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetLoadBalancerDataTrafficMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Aggregation(aggregation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetLoadBalancerDataTrafficMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancerDataTrafficMetrics`: GetDataTrafficMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetLoadBalancerDataTrafficMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerDataTrafficMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**DataTrafficMetricsGranularity**](DataTrafficMetricsGranularity.md) | The interval for each metric | 
 **aggregation** | [**MetricAggregation**](MetricAggregation.md) | The metric aggregation function | 

### Return type

[**GetDataTrafficMetricsResult**](GetDataTrafficMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerIP

> IpDetails GetLoadBalancerIP(ctx, loadBalancerId, ip).Execute()

Get IP details for a specific Load Balancer

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	ip := "10.0.0.1" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetLoadBalancerIP(context.Background(), loadBalancerId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetLoadBalancerIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancerIP`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetLoadBalancerIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IpDetails**](IpDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerIPList

> GetIPListResult GetLoadBalancerIPList(ctx, loadBalancerId).Version(version).NullRouted(nullRouted).Ips(ips).Execute()

List IP addresses associated with a specific Load Balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	version := openapiclient.ipVersion(4) // IpVersion |  (optional)
	nullRouted := true // bool |  (optional)
	ips := "ips_example" // string | A list of IPs separated by `|` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetLoadBalancerIPList(context.Background(), loadBalancerId).Version(version).NullRouted(nullRouted).Ips(ips).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetLoadBalancerIPList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancerIPList`: GetIPListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetLoadBalancerIPList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerIPListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**IpVersion**](IpVersion.md) |  | 
 **nullRouted** | **bool** |  | 
 **ips** | **string** | A list of IPs separated by &#x60;|&#x60; | 

### Return type

[**GetIPListResult**](GetIPListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerList

> LoadBalancers GetLoadBalancerList(ctx).Limit(limit).Offset(offset).Ip(ip).Reference(reference).Id(id).ContractState(contractState).ContractType(contractType).State(state).Region(region).Type_(type_).Execute()

Get load balancer list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	ip := "10.0.0.1" // string |  (optional)
	reference := "my-lb" // string |  (optional)
	id := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" // string |  (optional)
	contractState := openapiclient.contractState("ACTIVE") // ContractState |  (optional)
	contractType := "HOURLY" // string |  (optional)
	state := "RUNNING" // string |  (optional)
	region := openapiclient.regionName("eu-west-3") // RegionName | Available regions can be found using the List Regions endpoint. (optional)
	type_ := "lsw.c3.xlarge" // string | Available load balancer types can be found using the List Load Balancer Types endpoint. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetLoadBalancerList(context.Background()).Limit(limit).Offset(offset).Ip(ip).Reference(reference).Id(id).ContractState(contractState).ContractType(contractType).State(state).Region(region).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetLoadBalancerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancerList`: LoadBalancers
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetLoadBalancerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **ip** | **string** |  | 
 **reference** | **string** |  | 
 **id** | **string** |  | 
 **contractState** | [**ContractState**](ContractState.md) |  | 
 **contractType** | **string** |  | 
 **state** | **string** |  | 
 **region** | [**RegionName**](RegionName.md) | Available regions can be found using the List Regions endpoint. | 
 **type_** | **string** | Available load balancer types can be found using the List Load Balancer Types endpoint. | 

### Return type

[**LoadBalancers**](LoadBalancers.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerListener

> LoadBalancerListenerDetails GetLoadBalancerListener(ctx, loadBalancerId, listenerId).Execute()

Get listener details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	listenerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Listener ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetLoadBalancerListener(context.Background(), loadBalancerId, listenerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetLoadBalancerListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancerListener`: LoadBalancerListenerDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetLoadBalancerListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 
**listenerId** | **string** | Listener ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LoadBalancerListenerDetails**](LoadBalancerListenerDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerListenerList

> GetLoadBalancerListenerListResult GetLoadBalancerListenerList(ctx, loadBalancerId).Limit(limit).Offset(offset).Execute()

Get listener list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetLoadBalancerListenerList(context.Background(), loadBalancerId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetLoadBalancerListenerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancerListenerList`: GetLoadBalancerListenerListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetLoadBalancerListenerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerListenerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetLoadBalancerListenerListResult**](GetLoadBalancerListenerListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketAppList

> GetMarketAppListResult GetMarketAppList(ctx).Execute()

Get marketplace apps



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetMarketAppList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetMarketAppList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketAppList`: GetMarketAppListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetMarketAppList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketAppListRequest struct via the builder pattern


### Return type

[**GetMarketAppListResult**](GetMarketAppListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationSetting

> NotificationSetting GetNotificationSetting(ctx, instanceId, notificationSettingId).Execute()

Get details of a notification Setting

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	notificationSettingId := "3a042956-0689-45dc-8322-8b8325464182" // string | Notification Setting ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetNotificationSetting(context.Background(), instanceId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationSetting`: NotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**notificationSettingId** | **string** | Notification Setting ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NotificationSetting**](NotificationSetting.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationSettingList

> GetNotificationSettingListResult GetNotificationSettingList(ctx, instanceId).Limit(limit).Offset(offset).Execute()

List the notification settings of a customer

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetNotificationSettingList(context.Background(), instanceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetNotificationSettingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationSettingList`: GetNotificationSettingListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetNotificationSettingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationSettingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetNotificationSettingListResult**](GetNotificationSettingListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegionList

> GetRegionListResult GetRegionList(ctx).Limit(limit).Offset(offset).Execute()

List regions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetRegionList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetRegionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegionList`: GetRegionListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetRegionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetRegionListResult**](GetRegionListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReinstallImageList

> GetReinstallImageListResult GetReinstallImageList(ctx, instanceId).Limit(limit).Offset(offset).Custom(custom).Standard(standard).Execute()

List images available for reinstall

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	custom := true // bool | Filters the list to include only custom images. (optional)
	standard := true // bool | Filters the list to include only standard images. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetReinstallImageList(context.Background(), instanceId).Limit(limit).Offset(offset).Custom(custom).Standard(standard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetReinstallImageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReinstallImageList`: GetReinstallImageListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetReinstallImageList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReinstallImageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **custom** | **bool** | Filters the list to include only custom images. | 
 **standard** | **bool** | Filters the list to include only standard images. | 

### Return type

[**GetReinstallImageListResult**](GetReinstallImageListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestsMetrics

> GetRequestsMetricsResult GetRequestsMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get load balancer requests metrics. Not available for listeners with TCP protocol

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from`
	granularity := openapiclient.metricGranularity("5m") // MetricGranularity | The interval for each metric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetRequestsMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetRequestsMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequestsMetrics`: GetRequestsMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetRequestsMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestsMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricGranularity**](MetricGranularity.md) | The interval for each metric | 

### Return type

[**GetRequestsMetricsResult**](GetRequestsMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestsPerSecondMetrics

> GetRequestsPerSecondMetricsResult GetRequestsPerSecondMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get load balancer requests per second metrics. Not available for listeners with TCP protocol

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from`
	granularity := openapiclient.metricGranularity("5m") // MetricGranularity | The interval for each metric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetRequestsPerSecondMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetRequestsPerSecondMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequestsPerSecondMetrics`: GetRequestsPerSecondMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetRequestsPerSecondMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestsPerSecondMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricGranularity**](MetricGranularity.md) | The interval for each metric | 

### Return type

[**GetRequestsPerSecondMetricsResult**](GetRequestsPerSecondMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResponseCodesMetrics

> GetResponseCodesMetricsResult GetResponseCodesMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get response codes metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from`
	granularity := openapiclient.metricGranularity("5m") // MetricGranularity | The interval for each metric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetResponseCodesMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetResponseCodesMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResponseCodesMetrics`: GetResponseCodesMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetResponseCodesMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResponseCodesMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricGranularity**](MetricGranularity.md) | The interval for each metric | 

### Return type

[**GetResponseCodesMetricsResult**](GetResponseCodesMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResponseCodesPerSecondMetrics

> GetResponseCodesPerSecondMetricsResult GetResponseCodesPerSecondMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get response codes per second metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from`
	granularity := openapiclient.metricGranularity("5m") // MetricGranularity | The interval for each metric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetResponseCodesPerSecondMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetResponseCodesPerSecondMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResponseCodesPerSecondMetrics`: GetResponseCodesPerSecondMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetResponseCodesPerSecondMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResponseCodesPerSecondMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricGranularity**](MetricGranularity.md) | The interval for each metric | 

### Return type

[**GetResponseCodesPerSecondMetricsResult**](GetResponseCodesPerSecondMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityGroup

> SecurityGroup GetSecurityGroup(ctx, securityGroupId).Execute()

Get Security Group details

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	securityGroupId := "166cc685-8273-11f0-b4f8-2687367a8a5e" // string | Security Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetSecurityGroup(context.Background(), securityGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityGroup`: SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | Security Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityGroup**](SecurityGroup.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityGroupFirewallRules

> GetSecurityGroupFirewallRulesResult GetSecurityGroupFirewallRules(ctx, securityGroupId).Limit(limit).Offset(offset).Execute()

Get Security Group Firewall Rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	securityGroupId := "166cc685-8273-11f0-b4f8-2687367a8a5e" // string | Security Group ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetSecurityGroupFirewallRules(context.Background(), securityGroupId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetSecurityGroupFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityGroupFirewallRules`: GetSecurityGroupFirewallRulesResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetSecurityGroupFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | Security Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityGroupFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetSecurityGroupFirewallRulesResult**](GetSecurityGroupFirewallRulesResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityGroupList

> SecurityGroupList GetSecurityGroupList(ctx).Limit(limit).Offset(offset).State(state).Execute()

Get Security Group list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	state := openapiclient.securityGroupState("CREATING") // SecurityGroupState | Filter by state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetSecurityGroupList(context.Background()).Limit(limit).Offset(offset).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetSecurityGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityGroupList`: SecurityGroupList
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetSecurityGroupList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **state** | [**SecurityGroupState**](SecurityGroupState.md) | Filter by state | 

### Return type

[**SecurityGroupList**](SecurityGroupList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshot

> Snapshot GetSnapshot(ctx, instanceId, snapshotId).Execute()

Get snapshot detail

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	snapshotId := "54eee4c2-aef2-4cda-9fd5-74bfbf438abf" // string | Snapshot's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetSnapshot(context.Background(), instanceId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshot`: Snapshot
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**snapshotId** | **string** | Snapshot&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshotList

> GetSnapshotListResult GetSnapshotList(ctx, instanceId).Limit(limit).Offset(offset).Execute()

List snapshots



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetSnapshotList(context.Background(), instanceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetSnapshotList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshotList`: GetSnapshotListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetSnapshotList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetSnapshotListResult**](GetSnapshotListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTargetGroup

> TargetGroup GetTargetGroup(ctx, targetGroupId).Execute()

Get Target Group details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	targetGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Target Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetTargetGroup(context.Background(), targetGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTargetGroup`: TargetGroup
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | Target Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TargetGroup**](TargetGroup.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTargetGroupList

> GetTargetGroupListResult GetTargetGroupList(ctx).Limit(limit).Offset(offset).Id(id).Name(name).Protocol(protocol).Port(port).Region(region).Execute()

Get Target Group list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	id := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" // string |  (optional)
	name := "Foo bar" // string |  (optional)
	protocol := openapiclient.protocol("HTTP") // Protocol |  (optional)
	port := int32(80) // int32 |  (optional)
	region := openapiclient.regionName("eu-west-3") // RegionName |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetTargetGroupList(context.Background()).Limit(limit).Offset(offset).Id(id).Name(name).Protocol(protocol).Port(port).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetTargetGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTargetGroupList`: GetTargetGroupListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetTargetGroupList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **id** | **string** |  | 
 **name** | **string** |  | 
 **protocol** | [**Protocol**](Protocol.md) |  | 
 **port** | **int32** |  | 
 **region** | [**RegionName**](RegionName.md) |  | 

### Return type

[**GetTargetGroupListResult**](GetTargetGroupListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTargetList

> GetTargetListResult GetTargetList(ctx, targetGroupId).Limit(limit).Offset(offset).Execute()

Get Targets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	targetGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Target Group ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetTargetList(context.Background(), targetGroupId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetTargetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTargetList`: GetTargetListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetTargetList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | Target Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetTargetListResult**](GetTargetListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateInstanceTypeList

> InstanceTypes GetUpdateInstanceTypeList(ctx, instanceId).Limit(limit).Offset(offset).Execute()

List available instance types for update



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetUpdateInstanceTypeList(context.Background(), instanceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetUpdateInstanceTypeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpdateInstanceTypeList`: InstanceTypes
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetUpdateInstanceTypeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpdateInstanceTypeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**InstanceTypes**](InstanceTypes.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchInstance

> InstanceDetails LaunchInstance(ctx).LaunchInstanceOpts(launchInstanceOpts).Execute()

Launch instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	launchInstanceOpts := *openapiclient.NewLaunchInstanceOpts(openapiclient.regionName("eu-west-3"), openapiclient.typeName("lsw.m3.large"), "UBUNTU_22_04_64BIT", openapiclient.contractType("HOURLY"), int32(123), openapiclient.storageType("LOCAL")) // LaunchInstanceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.LaunchInstance(context.Background()).LaunchInstanceOpts(launchInstanceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.LaunchInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchInstance`: InstanceDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.LaunchInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLaunchInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **launchInstanceOpts** | [**LaunchInstanceOpts**](LaunchInstanceOpts.md) |  | 

### Return type

[**InstanceDetails**](InstanceDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchLoadBalancer

> LoadBalancerDetails LaunchLoadBalancer(ctx).LaunchLoadBalancerOpts(launchLoadBalancerOpts).Execute()

Launch Load balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	launchLoadBalancerOpts := *openapiclient.NewLaunchLoadBalancerOpts(openapiclient.regionName("eu-west-3"), openapiclient.typeName("lsw.m3.large"), openapiclient.contractType("HOURLY")) // LaunchLoadBalancerOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.LaunchLoadBalancer(context.Background()).LaunchLoadBalancerOpts(launchLoadBalancerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.LaunchLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchLoadBalancer`: LoadBalancerDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.LaunchLoadBalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLaunchLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **launchLoadBalancerOpts** | [**LaunchLoadBalancerOpts**](LaunchLoadBalancerOpts.md) |  | 

### Return type

[**LoadBalancerDetails**](LoadBalancerDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NullRouteInstanceIP

> IpDetails NullRouteInstanceIP(ctx, instanceId, ip).NullRouteIPOpts(nullRouteIPOpts).Execute()

Null route IP address for a specific resource Instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	ip := "10.0.0.1" // string | 
	nullRouteIPOpts := *openapiclient.NewNullRouteIPOpts() // NullRouteIPOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.NullRouteInstanceIP(context.Background(), instanceId, ip).NullRouteIPOpts(nullRouteIPOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.NullRouteInstanceIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NullRouteInstanceIP`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.NullRouteInstanceIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNullRouteInstanceIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nullRouteIPOpts** | [**NullRouteIPOpts**](NullRouteIPOpts.md) |  | 

### Return type

[**IpDetails**](IpDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NullRouteLoadBalancerIP

> IpDetails NullRouteLoadBalancerIP(ctx, loadBalancerId, ip).NullRouteIPOpts(nullRouteIPOpts).Execute()

Null route IP address for a specific resource Load Balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	ip := "10.0.0.1" // string | 
	nullRouteIPOpts := *openapiclient.NewNullRouteIPOpts() // NullRouteIPOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.NullRouteLoadBalancerIP(context.Background(), loadBalancerId, ip).NullRouteIPOpts(nullRouteIPOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.NullRouteLoadBalancerIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NullRouteLoadBalancerIP`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.NullRouteLoadBalancerIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNullRouteLoadBalancerIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nullRouteIPOpts** | [**NullRouteIPOpts**](NullRouteIPOpts.md) |  | 

### Return type

[**IpDetails**](IpDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootInstance

> RebootInstance(ctx, instanceId).Execute()

Reboot a specific Instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.RebootInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RebootInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootLoadBalancer

> RebootLoadBalancer(ctx, loadBalancerId).Execute()

Reboot a specific Load Balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.RebootLoadBalancer(context.Background(), loadBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RebootLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterAutoScalingGroupTargetGroup

> AutoScalingGroupDetails RegisterAutoScalingGroupTargetGroup(ctx, autoScalingGroupId).TargetGroupIdOpts(targetGroupIdOpts).Execute()

Register Target Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID
	targetGroupIdOpts := *openapiclient.NewTargetGroupIdOpts("TargetGroupId_example") // TargetGroupIdOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.RegisterAutoScalingGroupTargetGroup(context.Background(), autoScalingGroupId).TargetGroupIdOpts(targetGroupIdOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RegisterAutoScalingGroupTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAutoScalingGroupTargetGroup`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.RegisterAutoScalingGroupTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAutoScalingGroupTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetGroupIdOpts** | [**TargetGroupIdOpts**](TargetGroupIdOpts.md) |  | 

### Return type

[**AutoScalingGroupDetails**](AutoScalingGroupDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterTargets

> RegisterTargets(ctx, targetGroupId).RequestBody(requestBody).Execute()

Register Targets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	targetGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Target Group ID
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.RegisterTargets(context.Background(), targetGroupId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RegisterTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | Target Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReinstallInstance

> ReinstallInstance(ctx, instanceId).ReinstallResourceOpts(reinstallResourceOpts).Execute()

Reinstall an Instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	reinstallResourceOpts := *openapiclient.NewReinstallResourceOpts("ImageId_example") // ReinstallResourceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.ReinstallInstance(context.Background(), instanceId).ReinstallResourceOpts(reinstallResourceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.ReinstallInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReinstallInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reinstallResourceOpts** | [**ReinstallResourceOpts**](ReinstallResourceOpts.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFromPrivateNetwork

> RemoveFromPrivateNetwork(ctx, instanceId).Execute()

Remove instance from Private Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.RemoveFromPrivateNetwork(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RemoveFromPrivateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFromPrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveInstanceIPNullRoute

> IpDetails RemoveInstanceIPNullRoute(ctx, instanceId, ip).Execute()

Remove an IP null route for a specific Instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	ip := "10.0.0.1" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.RemoveInstanceIPNullRoute(context.Background(), instanceId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RemoveInstanceIPNullRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveInstanceIPNullRoute`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.RemoveInstanceIPNullRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveInstanceIPNullRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IpDetails**](IpDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLoadBalancerIPNullRoute

> IpDetails RemoveLoadBalancerIPNullRoute(ctx, loadBalancerId, ip).Execute()

Remove an IP null route for a specific Load Balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	ip := "10.0.0.1" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.RemoveLoadBalancerIPNullRoute(context.Background(), loadBalancerId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RemoveLoadBalancerIPNullRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveLoadBalancerIPNullRoute`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.RemoveLoadBalancerIPNullRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLoadBalancerIPNullRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IpDetails**](IpDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> ResetPassword(ctx, instanceId).Execute()

Reset the password for a specific Instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.ResetPassword(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.ResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreSnapshot

> RestoreSnapshot(ctx, instanceId, snapshotId).Execute()

Restore instance snapshot

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	snapshotId := "54eee4c2-aef2-4cda-9fd5-74bfbf438abf" // string | Snapshot's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.RestoreSnapshot(context.Background(), instanceId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RestoreSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**snapshotId** | **string** | Snapshot&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeSecurityGroupFirewallRules

> RevokeSecurityGroupFirewallRules(ctx, securityGroupId).RevokeFirewallRulesOpts(revokeFirewallRulesOpts).Execute()

Remove firewall rules from security group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	securityGroupId := "166cc685-8273-11f0-b4f8-2687367a8a5e" // string | Security Group ID
	revokeFirewallRulesOpts := *openapiclient.NewRevokeFirewallRulesOpts([]string{"Rules_example"}) // RevokeFirewallRulesOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.RevokeSecurityGroupFirewallRules(context.Background(), securityGroupId).RevokeFirewallRulesOpts(revokeFirewallRulesOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RevokeSecurityGroupFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | Security Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeSecurityGroupFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revokeFirewallRulesOpts** | [**RevokeFirewallRulesOpts**](RevokeFirewallRulesOpts.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartInstance

> StartInstance(ctx, instanceId).Execute()

Start a specific resource Instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.StartInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.StartInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartLoadBalancer

> StartLoadBalancer(ctx, loadBalancerId).Execute()

Start a specific resource Load Balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.StartLoadBalancer(context.Background(), loadBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.StartLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopInstance

> StopInstance(ctx, instanceId).Execute()

Stop a specific Instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.StopInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.StopInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopLoadBalancer

> StopLoadBalancer(ctx, loadBalancerId).Execute()

Stop a specific Load Balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.StopLoadBalancer(context.Background(), loadBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.StopLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreCredential

> StoreCredentialResult StoreCredential(ctx, instanceId).StoreCredentialOpts(storeCredentialOpts).Execute()

Store credentials for a specific Instance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	storeCredentialOpts := *openapiclient.NewStoreCredentialOpts(openapiclient.credentialType("OPERATING_SYSTEM"), "Username_example", "Password_example") // StoreCredentialOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.StoreCredential(context.Background(), instanceId).StoreCredentialOpts(storeCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.StoreCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreCredential`: StoreCredentialResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.StoreCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeCredentialOpts** | [**StoreCredentialOpts**](StoreCredentialOpts.md) |  | 

### Return type

[**StoreCredentialResult**](StoreCredentialResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateInstance

> TerminateInstance(ctx, instanceId).TerminateInstanceOpts(terminateInstanceOpts).Execute()

Terminate instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	terminateInstanceOpts := *openapiclient.NewTerminateInstanceOpts() // TerminateInstanceOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.TerminateInstance(context.Background(), instanceId).TerminateInstanceOpts(terminateInstanceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.TerminateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terminateInstanceOpts** | [**TerminateInstanceOpts**](TerminateInstanceOpts.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateLoadBalancer

> TerminateLoadBalancer(ctx, loadBalancerId).TerminateLoadBalancerOpts(terminateLoadBalancerOpts).Execute()

Delete load balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	terminateLoadBalancerOpts := *openapiclient.NewTerminateLoadBalancerOpts("ReasonCode_example") // TerminateLoadBalancerOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.TerminateLoadBalancer(context.Background(), loadBalancerId).TerminateLoadBalancerOpts(terminateLoadBalancerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.TerminateLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terminateLoadBalancerOpts** | [**TerminateLoadBalancerOpts**](TerminateLoadBalancerOpts.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoScalingGroup

> AutoScalingGroupDetails UpdateAutoScalingGroup(ctx, autoScalingGroupId).UpdateAutoScalingGroupOpts(updateAutoScalingGroupOpts).Execute()

Update Auto Scaling Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID
	updateAutoScalingGroupOpts := *openapiclient.NewUpdateAutoScalingGroupOpts() // UpdateAutoScalingGroupOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateAutoScalingGroup(context.Background(), autoScalingGroupId).UpdateAutoScalingGroupOpts(updateAutoScalingGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoScalingGroup`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateAutoScalingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoScalingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAutoScalingGroupOpts** | [**UpdateAutoScalingGroupOpts**](UpdateAutoScalingGroupOpts.md) |  | 

### Return type

[**AutoScalingGroupDetails**](AutoScalingGroupDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> UpdateCredentialResult UpdateCredential(ctx, instanceId, type_, username).UpdateCredentialOpts(updateCredentialOpts).Execute()

Update credentials for a given type and username

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	type_ := openapiclient.credentialType("OPERATING_SYSTEM") // CredentialType | Type of credential
	username := "root" // string | Username
	updateCredentialOpts := *openapiclient.NewUpdateCredentialOpts("Password_example") // UpdateCredentialOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateCredential(context.Background(), instanceId, type_, username).UpdateCredentialOpts(updateCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCredential`: UpdateCredentialResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**type_** | [**CredentialType**](.md) | Type of credential | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateCredentialOpts** | [**UpdateCredentialOpts**](UpdateCredentialOpts.md) |  | 

### Return type

[**UpdateCredentialResult**](UpdateCredentialResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateImage

> ImageDetails UpdateImage(ctx, imageId).UpdateImageOpts(updateImageOpts).Execute()

Update Custom Image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	imageId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Image's ID
	updateImageOpts := *openapiclient.NewUpdateImageOpts("Name_example") // UpdateImageOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateImage(context.Background(), imageId).UpdateImageOpts(updateImageOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateImage`: ImageDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateImageOpts** | [**UpdateImageOpts**](UpdateImageOpts.md) |  | 

### Return type

[**ImageDetails**](ImageDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstance

> InstanceDetails UpdateInstance(ctx, instanceId).UpdateInstanceOpts(updateInstanceOpts).Execute()

Update instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	updateInstanceOpts := *openapiclient.NewUpdateInstanceOpts() // UpdateInstanceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateInstance(context.Background(), instanceId).UpdateInstanceOpts(updateInstanceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstance`: InstanceDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInstanceOpts** | [**UpdateInstanceOpts**](UpdateInstanceOpts.md) |  | 

### Return type

[**InstanceDetails**](InstanceDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstanceIP

> IpDetails UpdateInstanceIP(ctx, instanceId, ip).UpdateIPOpts(updateIPOpts).Execute()

Update the IP address for a specific Instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	ip := "10.0.0.1" // string | 
	updateIPOpts := *openapiclient.NewUpdateIPOpts("ReverseLookup_example") // UpdateIPOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateInstanceIP(context.Background(), instanceId, ip).UpdateIPOpts(updateIPOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateInstanceIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstanceIP`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateInstanceIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateIPOpts** | [**UpdateIPOpts**](UpdateIPOpts.md) |  | 

### Return type

[**IpDetails**](IpDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancer

> LoadBalancerDetails UpdateLoadBalancer(ctx, loadBalancerId).UpdateLoadBalancerOpts(updateLoadBalancerOpts).Execute()

Update load balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	updateLoadBalancerOpts := *openapiclient.NewUpdateLoadBalancerOpts() // UpdateLoadBalancerOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateLoadBalancer(context.Background(), loadBalancerId).UpdateLoadBalancerOpts(updateLoadBalancerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoadBalancer`: LoadBalancerDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLoadBalancerOpts** | [**UpdateLoadBalancerOpts**](UpdateLoadBalancerOpts.md) |  | 

### Return type

[**LoadBalancerDetails**](LoadBalancerDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancerIP

> IpDetails UpdateLoadBalancerIP(ctx, loadBalancerId, ip).UpdateIPOpts(updateIPOpts).Execute()

Update the IP address for a specific Load Balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	ip := "10.0.0.1" // string | 
	updateIPOpts := *openapiclient.NewUpdateIPOpts("ReverseLookup_example") // UpdateIPOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateLoadBalancerIP(context.Background(), loadBalancerId, ip).UpdateIPOpts(updateIPOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateLoadBalancerIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoadBalancerIP`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateLoadBalancerIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateIPOpts** | [**UpdateIPOpts**](UpdateIPOpts.md) |  | 

### Return type

[**IpDetails**](IpDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancerListener

> LoadBalancerListener UpdateLoadBalancerListener(ctx, loadBalancerId, listenerId).LoadBalancerListenerOpts(loadBalancerListenerOpts).Execute()

Update a listener

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	listenerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Listener ID
	loadBalancerListenerOpts := *openapiclient.NewLoadBalancerListenerOpts() // LoadBalancerListenerOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateLoadBalancerListener(context.Background(), loadBalancerId, listenerId).LoadBalancerListenerOpts(loadBalancerListenerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateLoadBalancerListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoadBalancerListener`: LoadBalancerListener
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateLoadBalancerListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 
**listenerId** | **string** | Listener ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **loadBalancerListenerOpts** | [**LoadBalancerListenerOpts**](LoadBalancerListenerOpts.md) |  | 

### Return type

[**LoadBalancerListener**](LoadBalancerListener.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationSetting

> NotificationSetting UpdateNotificationSetting(ctx, instanceId, notificationSettingId).UpdateNotificationSettingOpts(updateNotificationSettingOpts).Execute()

Update Notification Setting details

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	notificationSettingId := "3a042956-0689-45dc-8322-8b8325464182" // string | Notification Setting ID
	updateNotificationSettingOpts := *openapiclient.NewUpdateNotificationSettingOpts() // UpdateNotificationSettingOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateNotificationSetting(context.Background(), instanceId, notificationSettingId).UpdateNotificationSettingOpts(updateNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotificationSetting`: NotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**notificationSettingId** | **string** | Notification Setting ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNotificationSettingOpts** | [**UpdateNotificationSettingOpts**](UpdateNotificationSettingOpts.md) |  | 

### Return type

[**NotificationSetting**](NotificationSetting.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityGroup

> SecurityGroup UpdateSecurityGroup(ctx, securityGroupId).UpdateSecurityGroupOpts(updateSecurityGroupOpts).Execute()

Update Security Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	securityGroupId := "166cc685-8273-11f0-b4f8-2687367a8a5e" // string | Security Group ID
	updateSecurityGroupOpts := *openapiclient.NewUpdateSecurityGroupOpts("Name_example") // UpdateSecurityGroupOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateSecurityGroup(context.Background(), securityGroupId).UpdateSecurityGroupOpts(updateSecurityGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSecurityGroup`: SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | Security Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSecurityGroupOpts** | [**UpdateSecurityGroupOpts**](UpdateSecurityGroupOpts.md) |  | 

### Return type

[**SecurityGroup**](SecurityGroup.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTargetGroup

> TargetGroup UpdateTargetGroup(ctx, targetGroupId).UpdateTargetGroupOpts(updateTargetGroupOpts).Execute()

Update Target Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	targetGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Target Group ID
	updateTargetGroupOpts := *openapiclient.NewUpdateTargetGroupOpts() // UpdateTargetGroupOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateTargetGroup(context.Background(), targetGroupId).UpdateTargetGroupOpts(updateTargetGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTargetGroup`: TargetGroup
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | Target Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTargetGroupOpts** | [**UpdateTargetGroupOpts**](UpdateTargetGroupOpts.md) |  | 

### Return type

[**TargetGroup**](TargetGroup.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

