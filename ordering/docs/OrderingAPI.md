# \OrderingAPI

All URIs are relative to *https://api.leaseweb.com/ordering/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDedicatedServer**](OrderingAPI.md#GetDedicatedServer) | **Get** /products/dedicatedServers/{dedicatedServerId} | Get a single dedicated server and its price.
[**GetDedicatedServerList**](OrderingAPI.md#GetDedicatedServerList) | **Get** /products/dedicatedServers | List available dedicated server configurations.
[**GetVps**](OrderingAPI.md#GetVps) | **Get** /products/vps/{vpsId} | Get a single VPS and its price.
[**GetVpsList**](OrderingAPI.md#GetVpsList) | **Get** /products/vps | Get a list of available VPS products and their prices.
[**OrderDedicatedServer**](OrderingAPI.md#OrderDedicatedServer) | **Post** /products/dedicatedServers/{dedicatedServerId}/order | Dedicated Server ordering.
[**OrderVps**](OrderingAPI.md#OrderVps) | **Post** /products/vps/{vpsId}/order | VPS ordering.



## GetDedicatedServer

> DedicatedServer GetDedicatedServer(ctx, dedicatedServerId).Location(location).ConnectedToAggregationPool(connectedToAggregationPool).Execute()

Get a single dedicated server and its price.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ordering"
)

func main() {
	dedicatedServerId := "27f38ef06cf3bf62e21e75401feac6cd" // string | The ID of a dedicated server
	location := "location_example" // string | 
	connectedToAggregationPool := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderingAPI.GetDedicatedServer(context.Background(), dedicatedServerId).Location(location).ConnectedToAggregationPool(connectedToAggregationPool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderingAPI.GetDedicatedServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDedicatedServer`: DedicatedServer
	fmt.Fprintf(os.Stdout, "Response from `OrderingAPI.GetDedicatedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dedicatedServerId** | **string** | The ID of a dedicated server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDedicatedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **location** | **string** |  | 
 **connectedToAggregationPool** | **bool** |  | [default to false]

### Return type

[**DedicatedServer**](DedicatedServer.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDedicatedServerList

> DedicatedServerList GetDedicatedServerList(ctx).Location(location).Ram(ram).DiskSize(diskSize).DiskAmount(diskAmount).Limit(limit).Offset(offset).Execute()

List available dedicated server configurations.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ordering"
)

func main() {
	location := "location_example" // string |  (optional)
	ram := "ram_example" // string |  (optional)
	diskSize := "diskSize_example" // string |  (optional)
	diskAmount := "diskAmount_example" // string |  (optional)
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderingAPI.GetDedicatedServerList(context.Background()).Location(location).Ram(ram).DiskSize(diskSize).DiskAmount(diskAmount).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderingAPI.GetDedicatedServerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDedicatedServerList`: DedicatedServerList
	fmt.Fprintf(os.Stdout, "Response from `OrderingAPI.GetDedicatedServerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDedicatedServerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **string** |  | 
 **ram** | **string** |  | 
 **diskSize** | **string** |  | 
 **diskAmount** | **string** |  | 
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**DedicatedServerList**](DedicatedServerList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVps

> Vps GetVps(ctx, vpsId).Location(location).DiskUpgrade(diskUpgrade).OperatingSystem(operatingSystem).ControlPanel(controlPanel).ContractTerm(contractTerm).BillingCycle(billingCycle).ServiceLevelAgreement(serviceLevelAgreement).Execute()

Get a single VPS and its price.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ordering"
)

func main() {
	vpsId := "VPS02_1" // string | The VPS ID
	location := "location_example" // string | 
	diskUpgrade := "diskUpgrade_example" // string |  (optional)
	operatingSystem := "operatingSystem_example" // string |  (optional)
	controlPanel := "controlPanel_example" // string |  (optional)
	contractTerm := openapiclient.contractTerm("1_MONTH") // ContractTerm |  (optional) (default to "1_YEAR")
	billingCycle := openapiclient.billingCycle("1_MONTH") // BillingCycle |  (optional) (default to "1_MONTH")
	serviceLevelAgreement := openapiclient.serviceLevelAgreement("Basic") // ServiceLevelAgreement |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderingAPI.GetVps(context.Background(), vpsId).Location(location).DiskUpgrade(diskUpgrade).OperatingSystem(operatingSystem).ControlPanel(controlPanel).ContractTerm(contractTerm).BillingCycle(billingCycle).ServiceLevelAgreement(serviceLevelAgreement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderingAPI.GetVps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVps`: Vps
	fmt.Fprintf(os.Stdout, "Response from `OrderingAPI.GetVps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpsId** | **string** | The VPS ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **location** | **string** |  | 
 **diskUpgrade** | **string** |  | 
 **operatingSystem** | **string** |  | 
 **controlPanel** | **string** |  | 
 **contractTerm** | [**ContractTerm**](ContractTerm.md) |  | [default to &quot;1_YEAR&quot;]
 **billingCycle** | [**BillingCycle**](BillingCycle.md) |  | [default to &quot;1_MONTH&quot;]
 **serviceLevelAgreement** | [**ServiceLevelAgreement**](ServiceLevelAgreement.md) |  | 

### Return type

[**Vps**](Vps.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpsList

> VpsList GetVpsList(ctx).Location(location).Limit(limit).Offset(offset).Execute()

Get a list of available VPS products and their prices.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ordering"
)

func main() {
	location := "location_example" // string |  (optional)
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderingAPI.GetVpsList(context.Background()).Location(location).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderingAPI.GetVpsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpsList`: VpsList
	fmt.Fprintf(os.Stdout, "Response from `OrderingAPI.GetVpsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVpsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **string** |  | 
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**VpsList**](VpsList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderDedicatedServer

> Order OrderDedicatedServer(ctx, dedicatedServerId).OrderDedicatedServerOpts(orderDedicatedServerOpts).Execute()

Dedicated Server ordering.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ordering"
)

func main() {
	dedicatedServerId := "27f38ef06cf3bf62e21e75401feac6cd" // string | The ID of a dedicated server
	orderDedicatedServerOpts := *openapiclient.NewOrderDedicatedServerOpts("AMS-01") // OrderDedicatedServerOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderingAPI.OrderDedicatedServer(context.Background(), dedicatedServerId).OrderDedicatedServerOpts(orderDedicatedServerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderingAPI.OrderDedicatedServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderDedicatedServer`: Order
	fmt.Fprintf(os.Stdout, "Response from `OrderingAPI.OrderDedicatedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dedicatedServerId** | **string** | The ID of a dedicated server | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderDedicatedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderDedicatedServerOpts** | [**OrderDedicatedServerOpts**](OrderDedicatedServerOpts.md) |  | 

### Return type

[**Order**](Order.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderVps

> Order OrderVps(ctx, vpsId).OrderVpsOpts(orderVpsOpts).Execute()

VPS ordering.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ordering"
)

func main() {
	vpsId := "VPS02_1" // string | The VPS ID
	orderVpsOpts := *openapiclient.NewOrderVpsOpts("AMS-01") // OrderVpsOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderingAPI.OrderVps(context.Background(), vpsId).OrderVpsOpts(orderVpsOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderingAPI.OrderVps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderVps`: Order
	fmt.Fprintf(os.Stdout, "Response from `OrderingAPI.OrderVps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpsId** | **string** | The VPS ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderVpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderVpsOpts** | [**OrderVpsOpts**](OrderVpsOpts.md) |  | 

### Return type

[**Order**](Order.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

