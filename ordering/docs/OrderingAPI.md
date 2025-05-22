# \OrderingAPI

All URIs are relative to *https://api.leaseweb.com/ordering/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDedicatedServer**](OrderingAPI.md#GetDedicatedServer) | **Get** /products/dedicatedServers/{dedicatedServerId} | Get a single dedicated server and its price.
[**GetDedicatedServerList**](OrderingAPI.md#GetDedicatedServerList) | **Get** /products/dedicatedServers | List available dedicated server configurations.
[**OrderDedicatedServer**](OrderingAPI.md#OrderDedicatedServer) | **Post** /products/dedicatedServers/{dedicatedServerId}/order | Dedicated Server ordering.



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

> DedicatedServerList GetDedicatedServerList(ctx).Location(location).Ram(ram).DiskSize(diskSize).DiskAmount(diskAmount).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderingAPI.GetDedicatedServerList(context.Background()).Location(location).Ram(ram).DiskSize(diskSize).DiskAmount(diskAmount).Execute()
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


## OrderDedicatedServer

> DedicatedServerOrder OrderDedicatedServer(ctx, dedicatedServerId).OrderDedicatedServerOpts(orderDedicatedServerOpts).Execute()

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
	// response from `OrderDedicatedServer`: DedicatedServerOrder
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

[**DedicatedServerOrder**](DedicatedServerOrder.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

