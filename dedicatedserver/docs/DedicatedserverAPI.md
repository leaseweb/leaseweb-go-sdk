# \DedicatedserverAPI

All URIs are relative to *https://api.leaseweb.com/bareMetals/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToPrivateNetwork**](DedicatedserverAPI.md#AddToPrivateNetwork) | **Put** /servers/{serverId}/privateNetworks/{privateNetworkId} | Add a server to private network
[**CancelActiveJob**](DedicatedserverAPI.md#CancelActiveJob) | **Post** /servers/{serverId}/cancelActiveJob | Cancel active job
[**CloseNetworkInterface**](DedicatedserverAPI.md#CloseNetworkInterface) | **Post** /servers/{serverId}/networkInterfaces/{networkTypeURL}/close | Close network interface
[**CloseNetworkInterfaces**](DedicatedserverAPI.md#CloseNetworkInterfaces) | **Post** /servers/{serverId}/networkInterfaces/close | Close all network interfaces
[**CreateBandwidthNotificationSetting**](DedicatedserverAPI.md#CreateBandwidthNotificationSetting) | **Post** /servers/{serverId}/notificationSettings/bandwidth | Create a bandwidth notification setting
[**CreateCredential**](DedicatedserverAPI.md#CreateCredential) | **Post** /servers/{serverId}/credentials | Create new server credentials
[**CreateDataTrafficNotificationSetting**](DedicatedserverAPI.md#CreateDataTrafficNotificationSetting) | **Post** /servers/{serverId}/notificationSettings/datatraffic | Create a data traffic notification setting
[**CreateDhcpReservation**](DedicatedserverAPI.md#CreateDhcpReservation) | **Post** /servers/{serverId}/leases | Create a DHCP reservation
[**DeleteBandwidthNotificationSetting**](DedicatedserverAPI.md#DeleteBandwidthNotificationSetting) | **Delete** /servers/{serverId}/notificationSettings/bandwidth/{notificationSettingId} | Delete a bandwidth notification setting
[**DeleteCredential**](DedicatedserverAPI.md#DeleteCredential) | **Delete** /servers/{serverId}/credentials/{type}/{username} | Delete server credentials
[**DeleteDataTrafficNotificationSetting**](DedicatedserverAPI.md#DeleteDataTrafficNotificationSetting) | **Delete** /servers/{serverId}/notificationSettings/datatraffic/{notificationSettingId} | Delete a data traffic notification setting
[**DeleteDhcpReservation**](DedicatedserverAPI.md#DeleteDhcpReservation) | **Delete** /servers/{serverId}/leases | Delete a DHCP reservation
[**DeleteFromPrivateNetwork**](DedicatedserverAPI.md#DeleteFromPrivateNetwork) | **Delete** /servers/{serverId}/privateNetworks/{privateNetworkId} | Delete a server from a private network
[**EnableRescueMode**](DedicatedserverAPI.md#EnableRescueMode) | **Post** /servers/{serverId}/rescueMode | Launch rescue mode
[**ExpireActiveJob**](DedicatedserverAPI.md#ExpireActiveJob) | **Post** /servers/{serverId}/expireActiveJob | Expire active job
[**GetBandwidthMetrics**](DedicatedserverAPI.md#GetBandwidthMetrics) | **Get** /servers/{serverId}/metrics/bandwidth | Show bandwidth metrics
[**GetBandwidthNotificationSetting**](DedicatedserverAPI.md#GetBandwidthNotificationSetting) | **Get** /servers/{serverId}/notificationSettings/bandwidth/{notificationSettingId} | Show a bandwidth notification setting
[**GetBandwidthNotificationSettingList**](DedicatedserverAPI.md#GetBandwidthNotificationSettingList) | **Get** /servers/{serverId}/notificationSettings/bandwidth | List bandwidth notification settings
[**GetControlPanelList**](DedicatedserverAPI.md#GetControlPanelList) | **Get** /controlPanels | List control panels
[**GetControlPanelListByOperatingSystemId**](DedicatedserverAPI.md#GetControlPanelListByOperatingSystemId) | **Get** /operatingSystems/{operatingSystemId}/controlPanels | List control panels by Operating System
[**GetCredential**](DedicatedserverAPI.md#GetCredential) | **Get** /servers/{serverId}/credentials/{type}/{username} | Show server credentials
[**GetCredentialList**](DedicatedserverAPI.md#GetCredentialList) | **Get** /servers/{serverId}/credentials | List server credentials
[**GetCredentialListByType**](DedicatedserverAPI.md#GetCredentialListByType) | **Get** /servers/{serverId}/credentials/{type} | List server credentials by type
[**GetDataTrafficMetrics**](DedicatedserverAPI.md#GetDataTrafficMetrics) | **Get** /servers/{serverId}/metrics/datatraffic | Show datatraffic metrics
[**GetDataTrafficNotificationSetting**](DedicatedserverAPI.md#GetDataTrafficNotificationSetting) | **Get** /servers/{serverId}/notificationSettings/datatraffic/{notificationSettingId} | Show a data traffic notification setting
[**GetDataTrafficNotificationSettingList**](DedicatedserverAPI.md#GetDataTrafficNotificationSettingList) | **Get** /servers/{serverId}/notificationSettings/datatraffic | List data traffic notification settings
[**GetDdosNotificationSetting**](DedicatedserverAPI.md#GetDdosNotificationSetting) | **Get** /servers/{serverId}/notificationSettings/ddos | Inspect DDoS notification settings
[**GetDhcpReservationList**](DedicatedserverAPI.md#GetDhcpReservationList) | **Get** /servers/{serverId}/leases | List DHCP reservations
[**GetHardware**](DedicatedserverAPI.md#GetHardware) | **Get** /servers/{serverId}/hardwareInfo | Show hardware information
[**GetIp**](DedicatedserverAPI.md#GetIp) | **Get** /servers/{serverId}/ips/{ip} | Show a server IP
[**GetIpList**](DedicatedserverAPI.md#GetIpList) | **Get** /servers/{serverId}/ips | List IPs
[**GetJob**](DedicatedserverAPI.md#GetJob) | **Get** /servers/{serverId}/jobs/{jobId} | Show a job
[**GetJobList**](DedicatedserverAPI.md#GetJobList) | **Get** /servers/{serverId}/jobs | List jobs
[**GetNetworkInterface**](DedicatedserverAPI.md#GetNetworkInterface) | **Get** /servers/{serverId}/networkInterfaces/{networkTypeURL} | Show a network interface
[**GetNetworkInterfaceList**](DedicatedserverAPI.md#GetNetworkInterfaceList) | **Get** /servers/{serverId}/networkInterfaces | List network interfaces
[**GetNullRouteHistory**](DedicatedserverAPI.md#GetNullRouteHistory) | **Get** /servers/{serverId}/nullRouteHistory | Show null route history
[**GetOperatingSystem**](DedicatedserverAPI.md#GetOperatingSystem) | **Get** /operatingSystems/{operatingSystemId} | Show an operating system
[**GetOperatingSystemList**](DedicatedserverAPI.md#GetOperatingSystemList) | **Get** /operatingSystems | List Operating Systems
[**GetPowerStatus**](DedicatedserverAPI.md#GetPowerStatus) | **Get** /servers/{serverId}/powerInfo | Show power status
[**GetRescueImageList**](DedicatedserverAPI.md#GetRescueImageList) | **Get** /rescueImages | Rescue Images
[**GetServer**](DedicatedserverAPI.md#GetServer) | **Get** /servers/{serverId} | Get server
[**GetServerList**](DedicatedserverAPI.md#GetServerList) | **Get** /servers | List servers
[**InstallOperatingSystem**](DedicatedserverAPI.md#InstallOperatingSystem) | **Post** /servers/{serverId}/install | Launch installation
[**IpmiReset**](DedicatedserverAPI.md#IpmiReset) | **Post** /servers/{serverId}/ipmiReset | Launch IPMI reset
[**NullIpRoute**](DedicatedserverAPI.md#NullIpRoute) | **Post** /servers/{serverId}/ips/{ip}/null | Null route an IP
[**OpenNetworkInterface**](DedicatedserverAPI.md#OpenNetworkInterface) | **Post** /servers/{serverId}/networkInterfaces/{networkTypeURL}/open | Open network interface
[**OpenNetworkInterfaces**](DedicatedserverAPI.md#OpenNetworkInterfaces) | **Post** /servers/{serverId}/networkInterfaces/open | Open all network interfaces
[**PowerCycle**](DedicatedserverAPI.md#PowerCycle) | **Post** /servers/{serverId}/powerCycle | Power cycle a server
[**PowerOff**](DedicatedserverAPI.md#PowerOff) | **Post** /servers/{serverId}/powerOff | Power off server
[**PowerOn**](DedicatedserverAPI.md#PowerOn) | **Post** /servers/{serverId}/powerOn | Power on server
[**RemoveNullIpRoute**](DedicatedserverAPI.md#RemoveNullIpRoute) | **Post** /servers/{serverId}/ips/{ip}/unnull | Remove a null route
[**RetryJob**](DedicatedserverAPI.md#RetryJob) | **Post** /servers/{serverId}/jobs/{jobId}/retry | Retry a job
[**ScanHardware**](DedicatedserverAPI.md#ScanHardware) | **Post** /servers/{serverId}/hardwareScan | Launch hardware scan
[**UpdateBandwidthNotificationSetting**](DedicatedserverAPI.md#UpdateBandwidthNotificationSetting) | **Put** /servers/{serverId}/notificationSettings/bandwidth/{notificationSettingId} | Update a bandwidth notification setting
[**UpdateCredential**](DedicatedserverAPI.md#UpdateCredential) | **Put** /servers/{serverId}/credentials/{type}/{username} | Update server credentials
[**UpdateDataTrafficNotificationSetting**](DedicatedserverAPI.md#UpdateDataTrafficNotificationSetting) | **Put** /servers/{serverId}/notificationSettings/datatraffic/{notificationSettingId} | Update a data traffic notification setting
[**UpdateDdosNotificationSetting**](DedicatedserverAPI.md#UpdateDdosNotificationSetting) | **Put** /servers/{serverId}/notificationSettings/ddos | Update DDoS notification settings
[**UpdateIpProfile**](DedicatedserverAPI.md#UpdateIpProfile) | **Put** /servers/{serverId}/ips/{ip} | Update an IP
[**UpdateReference**](DedicatedserverAPI.md#UpdateReference) | **Put** /servers/{serverId} | Update server



## AddToPrivateNetwork

> AddToPrivateNetwork(ctx, serverId, privateNetworkId).AddToPrivateNetworkOpts(addToPrivateNetworkOpts).Execute()

Add a server to private network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	privateNetworkId := "892" // string | The ID of a Private Network
	addToPrivateNetworkOpts := *openapiclient.NewAddToPrivateNetworkOpts(openapiclient.linkSpeed(100)) // AddToPrivateNetworkOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.AddToPrivateNetwork(context.Background(), serverId, privateNetworkId).AddToPrivateNetworkOpts(addToPrivateNetworkOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.AddToPrivateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**privateNetworkId** | **string** | The ID of a Private Network | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddToPrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addToPrivateNetworkOpts** | [**AddToPrivateNetworkOpts**](AddToPrivateNetworkOpts.md) |  | 

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


## CancelActiveJob

> Job CancelActiveJob(ctx, serverId).Execute()

Cancel active job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.CancelActiveJob(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.CancelActiveJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelActiveJob`: Job
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.CancelActiveJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelActiveJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Job**](Job.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloseNetworkInterface

> CloseNetworkInterface(ctx, serverId, networkTypeURL).Execute()

Close network interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	networkTypeURL := openapiclient.networkTypeURL("internal") // NetworkTypeURL | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.CloseNetworkInterface(context.Background(), serverId, networkTypeURL).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.CloseNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**networkTypeURL** | [**NetworkTypeURL**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseNetworkInterfaceRequest struct via the builder pattern


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


## CloseNetworkInterfaces

> CloseNetworkInterfaces(ctx, serverId).Execute()

Close all network interfaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.CloseNetworkInterfaces(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.CloseNetworkInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseNetworkInterfacesRequest struct via the builder pattern


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


## CreateBandwidthNotificationSetting

> BandwidthNotificationSetting CreateBandwidthNotificationSetting(ctx, serverId).BandwidthNotificationSettingOpts(bandwidthNotificationSettingOpts).Execute()

Create a bandwidth notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	bandwidthNotificationSettingOpts := *openapiclient.NewBandwidthNotificationSettingOpts("Frequency_example", "Threshold_example", "Unit_example") // BandwidthNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.CreateBandwidthNotificationSetting(context.Background(), serverId).BandwidthNotificationSettingOpts(bandwidthNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.CreateBandwidthNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBandwidthNotificationSetting`: BandwidthNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.CreateBandwidthNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBandwidthNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bandwidthNotificationSettingOpts** | [**BandwidthNotificationSettingOpts**](BandwidthNotificationSettingOpts.md) |  | 

### Return type

[**BandwidthNotificationSetting**](BandwidthNotificationSetting.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredential

> Credential CreateCredential(ctx, serverId).CreateCredentialOpts(createCredentialOpts).Execute()

Create new server credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	createCredentialOpts := *openapiclient.NewCreateCredentialOpts("Password_example", openapiclient.credentialType("OPERATING_SYSTEM"), "Username_example") // CreateCredentialOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.CreateCredential(context.Background(), serverId).CreateCredentialOpts(createCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.CreateCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.CreateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCredentialOpts** | [**CreateCredentialOpts**](CreateCredentialOpts.md) |  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataTrafficNotificationSetting

> DataTrafficNotificationSetting CreateDataTrafficNotificationSetting(ctx, serverId).DataTrafficNotificationSettingOpts(dataTrafficNotificationSettingOpts).Execute()

Create a data traffic notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	dataTrafficNotificationSettingOpts := *openapiclient.NewDataTrafficNotificationSettingOpts("Frequency_example", "Threshold_example", "Unit_example") // DataTrafficNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.CreateDataTrafficNotificationSetting(context.Background(), serverId).DataTrafficNotificationSettingOpts(dataTrafficNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.CreateDataTrafficNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataTrafficNotificationSetting`: DataTrafficNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.CreateDataTrafficNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataTrafficNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataTrafficNotificationSettingOpts** | [**DataTrafficNotificationSettingOpts**](DataTrafficNotificationSettingOpts.md) |  | 

### Return type

[**DataTrafficNotificationSetting**](DataTrafficNotificationSetting.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDhcpReservation

> CreateDhcpReservation(ctx, serverId).CreateDhcpReservationOpts(createDhcpReservationOpts).Execute()

Create a DHCP reservation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	createDhcpReservationOpts := *openapiclient.NewCreateDhcpReservationOpts("Bootfile_example") // CreateDhcpReservationOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.CreateDhcpReservation(context.Background(), serverId).CreateDhcpReservationOpts(createDhcpReservationOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.CreateDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDhcpReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDhcpReservationOpts** | [**CreateDhcpReservationOpts**](CreateDhcpReservationOpts.md) |  | 

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


## DeleteBandwidthNotificationSetting

> DeleteBandwidthNotificationSetting(ctx, serverId, notificationSettingId).Execute()

Delete a bandwidth notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.DeleteBandwidthNotificationSetting(context.Background(), serverId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.DeleteBandwidthNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBandwidthNotificationSettingRequest struct via the builder pattern


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

> DeleteCredential(ctx, serverId, type_, username).Execute()

Delete server credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	type_ := openapiclient.credentialType("OPERATING_SYSTEM") // CredentialType | The type of the credential.
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.DeleteCredential(context.Background(), serverId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.DeleteCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**type_** | [**CredentialType**](.md) | The type of the credential. | 
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


## DeleteDataTrafficNotificationSetting

> DeleteDataTrafficNotificationSetting(ctx, serverId, notificationSettingId).Execute()

Delete a data traffic notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.DeleteDataTrafficNotificationSetting(context.Background(), serverId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.DeleteDataTrafficNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataTrafficNotificationSettingRequest struct via the builder pattern


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


## DeleteDhcpReservation

> DeleteDhcpReservation(ctx, serverId).Execute()

Delete a DHCP reservation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.DeleteDhcpReservation(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.DeleteDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDhcpReservationRequest struct via the builder pattern


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


## DeleteFromPrivateNetwork

> DeleteFromPrivateNetwork(ctx, serverId, privateNetworkId).Execute()

Delete a server from a private network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	privateNetworkId := "892" // string | The ID of a Private Network

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.DeleteFromPrivateNetwork(context.Background(), serverId, privateNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.DeleteFromPrivateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**privateNetworkId** | **string** | The ID of a Private Network | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFromPrivateNetworkRequest struct via the builder pattern


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


## EnableRescueMode

> Job EnableRescueMode(ctx, serverId).EnableRescueModeOpts(enableRescueModeOpts).Execute()

Launch rescue mode



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	enableRescueModeOpts := *openapiclient.NewEnableRescueModeOpts("RescueImageId_example") // EnableRescueModeOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.EnableRescueMode(context.Background(), serverId).EnableRescueModeOpts(enableRescueModeOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.EnableRescueMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableRescueMode`: Job
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.EnableRescueMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableRescueModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableRescueModeOpts** | [**EnableRescueModeOpts**](EnableRescueModeOpts.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpireActiveJob

> Job ExpireActiveJob(ctx, serverId).Execute()

Expire active job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.ExpireActiveJob(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.ExpireActiveJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpireActiveJob`: Job
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.ExpireActiveJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpireActiveJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Job**](Job.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBandwidthMetrics

> Metrics GetBandwidthMetrics(ctx, serverId).From(from).To(to).Aggregation(aggregation).Granularity(granularity).Execute()

Show bandwidth metrics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	from := time.Now() // time.Time | Start of date interval in ISO-8601 format. The returned data will include everything up from - and including - the specified date time.
	to := time.Now() // time.Time | End of date interval in ISO-8601 format. The returned data will include everything up until - but not including - the specified date time.
	aggregation := "aggregation_example" // string | Aggregate each metric using the given aggregation function. When the aggregation type `95TH` is specified the granularity parameter should be omitted from the request.
	granularity := "granularity_example" // string | Specify the preferred interval for each metric. If granularity is omitted from the request, only one metric is returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetBandwidthMetrics(context.Background(), serverId).From(from).To(to).Aggregation(aggregation).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetBandwidthMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBandwidthMetrics`: Metrics
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetBandwidthMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBandwidthMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Start of date interval in ISO-8601 format. The returned data will include everything up from - and including - the specified date time. | 
 **to** | **time.Time** | End of date interval in ISO-8601 format. The returned data will include everything up until - but not including - the specified date time. | 
 **aggregation** | **string** | Aggregate each metric using the given aggregation function. When the aggregation type &#x60;95TH&#x60; is specified the granularity parameter should be omitted from the request. | 
 **granularity** | **string** | Specify the preferred interval for each metric. If granularity is omitted from the request, only one metric is returned. | 

### Return type

[**Metrics**](Metrics.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBandwidthNotificationSetting

> BandwidthNotificationSetting GetBandwidthNotificationSetting(ctx, serverId, notificationSettingId).Execute()

Show a bandwidth notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetBandwidthNotificationSetting(context.Background(), serverId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetBandwidthNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBandwidthNotificationSetting`: BandwidthNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetBandwidthNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBandwidthNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BandwidthNotificationSetting**](BandwidthNotificationSetting.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBandwidthNotificationSettingList

> GetBandwidthNotificationSettingListResult GetBandwidthNotificationSettingList(ctx, serverId).Limit(limit).Offset(offset).Execute()

List bandwidth notification settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetBandwidthNotificationSettingList(context.Background(), serverId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetBandwidthNotificationSettingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBandwidthNotificationSettingList`: GetBandwidthNotificationSettingListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetBandwidthNotificationSettingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBandwidthNotificationSettingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetBandwidthNotificationSettingListResult**](GetBandwidthNotificationSettingListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControlPanelList

> ControlPanelList GetControlPanelList(ctx).Limit(limit).Offset(offset).Execute()

List control panels



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetControlPanelList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetControlPanelList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetControlPanelList`: ControlPanelList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetControlPanelList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetControlPanelListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**ControlPanelList**](ControlPanelList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControlPanelListByOperatingSystemId

> ControlPanelList GetControlPanelListByOperatingSystemId(ctx, operatingSystemId).Limit(limit).Offset(offset).Execute()

List control panels by Operating System



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	operatingSystemId := "UBUNTU_22_04_64BIT" // string | Operating system ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetControlPanelListByOperatingSystemId(context.Background(), operatingSystemId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetControlPanelListByOperatingSystemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetControlPanelListByOperatingSystemId`: ControlPanelList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetControlPanelListByOperatingSystemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatingSystemId** | **string** | Operating system ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetControlPanelListByOperatingSystemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**ControlPanelList**](ControlPanelList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredential

> Credential GetCredential(ctx, serverId, type_, username).Execute()

Show server credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	type_ := openapiclient.credentialType("OPERATING_SYSTEM") // CredentialType | The type of the credential.
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetCredential(context.Background(), serverId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**type_** | [**CredentialType**](.md) | The type of the credential. | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Credential**](Credential.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentialList

> CredentialList GetCredentialList(ctx, serverId).Limit(limit).Offset(offset).Execute()

List server credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetCredentialList(context.Background(), serverId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetCredentialList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialList`: CredentialList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**CredentialList**](CredentialList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentialListByType

> CredentialList GetCredentialListByType(ctx, serverId, type_).Limit(limit).Offset(offset).Execute()

List server credentials by type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	type_ := openapiclient.credentialType("OPERATING_SYSTEM") // CredentialType | The type of the credential.
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetCredentialListByType(context.Background(), serverId, type_).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetCredentialListByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialListByType`: CredentialList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetCredentialListByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**type_** | [**CredentialType**](.md) | The type of the credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialListByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**CredentialList**](CredentialList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTrafficMetrics

> Metrics GetDataTrafficMetrics(ctx, serverId).From(from).To(to).Aggregation(aggregation).Granularity(granularity).Execute()

Show datatraffic metrics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	from := time.Now() // time.Time | Start of date interval in ISO-8601 format. The returned data will include everything up from - and including - the specified date time.
	to := time.Now() // time.Time | End of date interval in ISO-8601 format. The returned data will include everything up until - but not including - the specified date time.
	aggregation := "aggregation_example" // string | Aggregate each metric using the given aggregation function.
	granularity := "granularity_example" // string | Specify the preferred interval for each metric. If granularity is omitted from the request, only one metric is returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetDataTrafficMetrics(context.Background(), serverId).From(from).To(to).Aggregation(aggregation).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetDataTrafficMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTrafficMetrics`: Metrics
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetDataTrafficMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTrafficMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Start of date interval in ISO-8601 format. The returned data will include everything up from - and including - the specified date time. | 
 **to** | **time.Time** | End of date interval in ISO-8601 format. The returned data will include everything up until - but not including - the specified date time. | 
 **aggregation** | **string** | Aggregate each metric using the given aggregation function. | 
 **granularity** | **string** | Specify the preferred interval for each metric. If granularity is omitted from the request, only one metric is returned. | 

### Return type

[**Metrics**](Metrics.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTrafficNotificationSetting

> DataTrafficNotificationSetting GetDataTrafficNotificationSetting(ctx, serverId, notificationSettingId).Execute()

Show a data traffic notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetDataTrafficNotificationSetting(context.Background(), serverId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetDataTrafficNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTrafficNotificationSetting`: DataTrafficNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetDataTrafficNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTrafficNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataTrafficNotificationSetting**](DataTrafficNotificationSetting.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTrafficNotificationSettingList

> GetDataTrafficNotificationSettingListResult GetDataTrafficNotificationSettingList(ctx, serverId).Limit(limit).Offset(offset).Execute()

List data traffic notification settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetDataTrafficNotificationSettingList(context.Background(), serverId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetDataTrafficNotificationSettingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTrafficNotificationSettingList`: GetDataTrafficNotificationSettingListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetDataTrafficNotificationSettingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTrafficNotificationSettingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetDataTrafficNotificationSettingListResult**](GetDataTrafficNotificationSettingListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDdosNotificationSetting

> GetDdosNotificationSettingResult GetDdosNotificationSetting(ctx, serverId).Execute()

Inspect DDoS notification settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetDdosNotificationSetting(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetDdosNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDdosNotificationSetting`: GetDdosNotificationSettingResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetDdosNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDdosNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDdosNotificationSettingResult**](GetDdosNotificationSettingResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDhcpReservationList

> GetDhcpReservationListResult GetDhcpReservationList(ctx, serverId).Execute()

List DHCP reservations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetDhcpReservationList(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetDhcpReservationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDhcpReservationList`: GetDhcpReservationListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetDhcpReservationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDhcpReservationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDhcpReservationListResult**](GetDhcpReservationListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHardware

> GetHardwareResult GetHardware(ctx, serverId).Execute()

Show hardware information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetHardware(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetHardware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHardware`: GetHardwareResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetHardware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHardwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetHardwareResult**](GetHardwareResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIp

> Ip GetIp(ctx, serverId, ip).Execute()

Show a server IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetIp(context.Background(), serverId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIp`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Ip**](Ip.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpList

> IpList GetIpList(ctx, serverId).NetworkType(networkType).Version(version).NullRouted(nullRouted).Ips(ips).Limit(limit).Offset(offset).Execute()

List IPs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	networkType := openapiclient.networkType("INTERNAL") // NetworkType | Filter the collection of ip addresses by network type (optional)
	version := "version_example" // string | Filter the collection by ip version (optional)
	nullRouted := "nullRouted_example" // string | Filter Ips by Nulled-Status (optional)
	ips := "ips_example" // string | Filter the collection of Ips for the comma separated list of Ips (optional)
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetIpList(context.Background(), serverId).NetworkType(networkType).Version(version).NullRouted(nullRouted).Ips(ips).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetIpList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpList`: IpList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetIpList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkType** | [**NetworkType**](NetworkType.md) | Filter the collection of ip addresses by network type | 
 **version** | **string** | Filter the collection by ip version | 
 **nullRouted** | **string** | Filter Ips by Nulled-Status | 
 **ips** | **string** | Filter the collection of Ips for the comma separated list of Ips | 
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**IpList**](IpList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJob

> CurrentJob GetJob(ctx, serverId, jobId).Execute()

Show a job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	jobId := "3a867358-5b4b-44ee-88ac-4274603ef641" // string | The ID of a Job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetJob(context.Background(), serverId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJob`: CurrentJob
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**jobId** | **string** | The ID of a Job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CurrentJob**](CurrentJob.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobList

> JobList GetJobList(ctx, serverId).Limit(limit).Offset(offset).Type_(type_).Status(status).IsRunning(isRunning).Execute()

List jobs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	type_ := "install" // string | Filter the list of jobs by type. (optional)
	status := "CANCELED" // string | Filter the list of jobs by status. (optional)
	isRunning := "true" // string | Filter the list for running jobs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetJobList(context.Background(), serverId).Limit(limit).Offset(offset).Type_(type_).Status(status).IsRunning(isRunning).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetJobList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobList`: JobList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetJobList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **type_** | **string** | Filter the list of jobs by type. | 
 **status** | **string** | Filter the list of jobs by status. | 
 **isRunning** | **string** | Filter the list for running jobs | 

### Return type

[**JobList**](JobList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkInterface

> OperationNetworkInterface GetNetworkInterface(ctx, serverId, networkTypeURL).Execute()

Show a network interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	networkTypeURL := openapiclient.networkTypeURL("internal") // NetworkTypeURL | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetNetworkInterface(context.Background(), serverId, networkTypeURL).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkInterface`: OperationNetworkInterface
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**networkTypeURL** | [**NetworkTypeURL**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationNetworkInterface**](OperationNetworkInterface.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkInterfaceList

> GetNetworkInterfaceListResult GetNetworkInterfaceList(ctx, serverId).Execute()

List network interfaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetNetworkInterfaceList(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetNetworkInterfaceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkInterfaceList`: GetNetworkInterfaceListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetNetworkInterfaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInterfaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkInterfaceListResult**](GetNetworkInterfaceListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNullRouteHistory

> NullRouteList GetNullRouteHistory(ctx, serverId).Limit(limit).Offset(offset).Execute()

Show null route history



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetNullRouteHistory(context.Background(), serverId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetNullRouteHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNullRouteHistory`: NullRouteList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetNullRouteHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNullRouteHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**NullRouteList**](NullRouteList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperatingSystem

> GetOperatingSystemResult GetOperatingSystem(ctx, operatingSystemId).ControlPanelId(controlPanelId).Execute()

Show an operating system



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	operatingSystemId := "UBUNTU_22_04_64BIT" // string | Operating system ID
	controlPanelId := "controlPanelId_example" // string | The Control Panel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetOperatingSystem(context.Background(), operatingSystemId).ControlPanelId(controlPanelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetOperatingSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperatingSystem`: GetOperatingSystemResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetOperatingSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatingSystemId** | **string** | Operating system ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatingSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **controlPanelId** | **string** | The Control Panel ID | 

### Return type

[**GetOperatingSystemResult**](GetOperatingSystemResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperatingSystemList

> GetOperatingSystemListResult GetOperatingSystemList(ctx).Limit(limit).Offset(offset).ControlPanelId(controlPanelId).Execute()

List Operating Systems



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	controlPanelId := "controlPanelId_example" // string | Filter operating systems by control panel id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetOperatingSystemList(context.Background()).Limit(limit).Offset(offset).ControlPanelId(controlPanelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetOperatingSystemList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperatingSystemList`: GetOperatingSystemListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetOperatingSystemList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatingSystemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **controlPanelId** | **string** | Filter operating systems by control panel id | 

### Return type

[**GetOperatingSystemListResult**](GetOperatingSystemListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPowerStatus

> GetPowerStatusResult GetPowerStatus(ctx, serverId).Execute()

Show power status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetPowerStatus(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetPowerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPowerStatus`: GetPowerStatusResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetPowerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPowerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPowerStatusResult**](GetPowerStatusResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRescueImageList

> GetRescueImageListResult GetRescueImageList(ctx).Limit(limit).Offset(offset).Execute()

Rescue Images



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetRescueImageList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetRescueImageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRescueImageList`: GetRescueImageListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetRescueImageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRescueImageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetRescueImageListResult**](GetRescueImageListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServer

> Server GetServer(ctx, serverId).Execute()

Get server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServer`: Server
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Server**](Server.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerList

> GetServerListResult GetServerList(ctx).Limit(limit).Offset(offset).Reference(reference).Ip(ip).MacAddress(macAddress).Site(site).PrivateRackId(privateRackId).PrivateNetworkCapable(privateNetworkCapable).PrivateNetworkEnabled(privateNetworkEnabled).Execute()

List servers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	reference := "my-db" // string | Filter the list of servers by reference. (optional)
	ip := "127.0.0.4" // string | Filter the list of servers by ip address. (optional)
	macAddress := "aa:bb:cc:dd:ee:ff" // string | Filter the list of servers by mac address. (optional)
	site := "FRA-10" // string | Filter the list of servers by site (location). (optional)
	privateRackId := "123" // string | Filter the list of servers by dedicated rack id. (optional)
	privateNetworkCapable := "privateNetworkCapable_example" // string | Filter the list for private network capable servers (optional)
	privateNetworkEnabled := "privateNetworkEnabled_example" // string | Filter the list for private network enabled servers (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetServerList(context.Background()).Limit(limit).Offset(offset).Reference(reference).Ip(ip).MacAddress(macAddress).Site(site).PrivateRackId(privateRackId).PrivateNetworkCapable(privateNetworkCapable).PrivateNetworkEnabled(privateNetworkEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetServerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerList`: GetServerListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetServerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **reference** | **string** | Filter the list of servers by reference. | 
 **ip** | **string** | Filter the list of servers by ip address. | 
 **macAddress** | **string** | Filter the list of servers by mac address. | 
 **site** | **string** | Filter the list of servers by site (location). | 
 **privateRackId** | **string** | Filter the list of servers by dedicated rack id. | 
 **privateNetworkCapable** | **string** | Filter the list for private network capable servers | 
 **privateNetworkEnabled** | **string** | Filter the list for private network enabled servers | 

### Return type

[**GetServerListResult**](GetServerListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallOperatingSystem

> ServerJob InstallOperatingSystem(ctx, serverId).InstallOperatingSystemOpts(installOperatingSystemOpts).Execute()

Launch installation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	installOperatingSystemOpts := *openapiclient.NewInstallOperatingSystemOpts("OperatingSystemId_example") // InstallOperatingSystemOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.InstallOperatingSystem(context.Background(), serverId).InstallOperatingSystemOpts(installOperatingSystemOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.InstallOperatingSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstallOperatingSystem`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.InstallOperatingSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallOperatingSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **installOperatingSystemOpts** | [**InstallOperatingSystemOpts**](InstallOperatingSystemOpts.md) |  | 

### Return type

[**ServerJob**](ServerJob.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpmiReset

> Job IpmiReset(ctx, serverId).IpmiResetOpts(ipmiResetOpts).Execute()

Launch IPMI reset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ipmiResetOpts := *openapiclient.NewIpmiResetOpts() // IpmiResetOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.IpmiReset(context.Background(), serverId).IpmiResetOpts(ipmiResetOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.IpmiReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IpmiReset`: Job
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.IpmiReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpmiResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipmiResetOpts** | [**IpmiResetOpts**](IpmiResetOpts.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NullIpRoute

> Ip NullIpRoute(ctx, serverId, ip).Execute()

Null route an IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.NullIpRoute(context.Background(), serverId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.NullIpRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NullIpRoute`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.NullIpRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiNullIpRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Ip**](Ip.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenNetworkInterface

> OpenNetworkInterface(ctx, serverId, networkTypeURL).Execute()

Open network interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	networkTypeURL := openapiclient.networkTypeURL("internal") // NetworkTypeURL | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.OpenNetworkInterface(context.Background(), serverId, networkTypeURL).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.OpenNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**networkTypeURL** | [**NetworkTypeURL**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenNetworkInterfaceRequest struct via the builder pattern


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


## OpenNetworkInterfaces

> OpenNetworkInterfaces(ctx, serverId).Execute()

Open all network interfaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.OpenNetworkInterfaces(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.OpenNetworkInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenNetworkInterfacesRequest struct via the builder pattern


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


## PowerCycle

> PowerCycle(ctx, serverId).Execute()

Power cycle a server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.PowerCycle(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.PowerCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiPowerCycleRequest struct via the builder pattern


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


## PowerOff

> PowerOff(ctx, serverId).Execute()

Power off server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.PowerOff(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.PowerOff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiPowerOffRequest struct via the builder pattern


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


## PowerOn

> PowerOn(ctx, serverId).Execute()

Power on server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.PowerOn(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.PowerOn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiPowerOnRequest struct via the builder pattern


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


## RemoveNullIpRoute

> Ip RemoveNullIpRoute(ctx, serverId, ip).Execute()

Remove a null route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.RemoveNullIpRoute(context.Background(), serverId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.RemoveNullIpRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveNullIpRoute`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.RemoveNullIpRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNullIpRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Ip**](Ip.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryJob

> CurrentJob RetryJob(ctx, serverId, jobId).Execute()

Retry a job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	jobId := "3a867358-5b4b-44ee-88ac-4274603ef641" // string | The ID of a Job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.RetryJob(context.Background(), serverId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.RetryJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryJob`: CurrentJob
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.RetryJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**jobId** | **string** | The ID of a Job | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CurrentJob**](CurrentJob.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScanHardware

> Job ScanHardware(ctx, serverId).ScanHardwareOpts(scanHardwareOpts).Execute()

Launch hardware scan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	scanHardwareOpts := *openapiclient.NewScanHardwareOpts() // ScanHardwareOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.ScanHardware(context.Background(), serverId).ScanHardwareOpts(scanHardwareOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.ScanHardware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScanHardware`: Job
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.ScanHardware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiScanHardwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scanHardwareOpts** | [**ScanHardwareOpts**](ScanHardwareOpts.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBandwidthNotificationSetting

> BandwidthNotificationSetting UpdateBandwidthNotificationSetting(ctx, serverId, notificationSettingId).BandwidthNotificationSettingOpts(bandwidthNotificationSettingOpts).Execute()

Update a bandwidth notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting
	bandwidthNotificationSettingOpts := *openapiclient.NewBandwidthNotificationSettingOpts("Frequency_example", "Threshold_example", "Unit_example") // BandwidthNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.UpdateBandwidthNotificationSetting(context.Background(), serverId, notificationSettingId).BandwidthNotificationSettingOpts(bandwidthNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.UpdateBandwidthNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBandwidthNotificationSetting`: BandwidthNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.UpdateBandwidthNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBandwidthNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bandwidthNotificationSettingOpts** | [**BandwidthNotificationSettingOpts**](BandwidthNotificationSettingOpts.md) |  | 

### Return type

[**BandwidthNotificationSetting**](BandwidthNotificationSetting.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> Credential UpdateCredential(ctx, serverId, type_, username).UpdateCredentialOpts(updateCredentialOpts).Execute()

Update server credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	type_ := openapiclient.credentialType("OPERATING_SYSTEM") // CredentialType | The type of the credential.
	username := "root" // string | Username
	updateCredentialOpts := *openapiclient.NewUpdateCredentialOpts("Password_example") // UpdateCredentialOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.UpdateCredential(context.Background(), serverId, type_, username).UpdateCredentialOpts(updateCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.UpdateCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.UpdateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**type_** | [**CredentialType**](.md) | The type of the credential. | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateCredentialOpts** | [**UpdateCredentialOpts**](UpdateCredentialOpts.md) |  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataTrafficNotificationSetting

> DataTrafficNotificationSettingOpts UpdateDataTrafficNotificationSetting(ctx, serverId, notificationSettingId).DataTrafficNotificationSettingOpts(dataTrafficNotificationSettingOpts).Execute()

Update a data traffic notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting
	dataTrafficNotificationSettingOpts := *openapiclient.NewDataTrafficNotificationSettingOpts("Frequency_example", "Threshold_example", "Unit_example") // DataTrafficNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.UpdateDataTrafficNotificationSetting(context.Background(), serverId, notificationSettingId).DataTrafficNotificationSettingOpts(dataTrafficNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.UpdateDataTrafficNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataTrafficNotificationSetting`: DataTrafficNotificationSettingOpts
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.UpdateDataTrafficNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataTrafficNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataTrafficNotificationSettingOpts** | [**DataTrafficNotificationSettingOpts**](DataTrafficNotificationSettingOpts.md) |  | 

### Return type

[**DataTrafficNotificationSettingOpts**](DataTrafficNotificationSettingOpts.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDdosNotificationSetting

> UpdateDdosNotificationSetting(ctx, serverId).UpdateDdosNotificationSettingOpts(updateDdosNotificationSettingOpts).Execute()

Update DDoS notification settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	updateDdosNotificationSettingOpts := *openapiclient.NewUpdateDdosNotificationSettingOpts() // UpdateDdosNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.UpdateDdosNotificationSetting(context.Background(), serverId).UpdateDdosNotificationSettingOpts(updateDdosNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.UpdateDdosNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDdosNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDdosNotificationSettingOpts** | [**UpdateDdosNotificationSettingOpts**](UpdateDdosNotificationSettingOpts.md) |  | 

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


## UpdateIpProfile

> Ip UpdateIpProfile(ctx, serverId, ip).UpdateIpProfileOpts(updateIpProfileOpts).Execute()

Update an IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ip := "127.0.0.6" // string | The IP Address
	updateIpProfileOpts := *openapiclient.NewUpdateIpProfileOpts() // UpdateIpProfileOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.UpdateIpProfile(context.Background(), serverId, ip).UpdateIpProfileOpts(updateIpProfileOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.UpdateIpProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIpProfile`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.UpdateIpProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIpProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateIpProfileOpts** | [**UpdateIpProfileOpts**](UpdateIpProfileOpts.md) |  | 

### Return type

[**Ip**](Ip.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReference

> UpdateReference(ctx, serverId).UpdateReferenceOpts(updateReferenceOpts).Execute()

Update server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	updateReferenceOpts := *openapiclient.NewUpdateReferenceOpts("Reference_example") // UpdateReferenceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.UpdateReference(context.Background(), serverId).UpdateReferenceOpts(updateReferenceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.UpdateReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateReferenceOpts** | [**UpdateReferenceOpts**](UpdateReferenceOpts.md) |  | 

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

