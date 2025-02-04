# \NetworkequipmentAPI

All URIs are relative to *https://api.leaseweb.com/bareMetals/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredential**](NetworkequipmentAPI.md#CreateCredential) | **Post** /networkEquipments/{networkEquipmentId}/credentials | Create new network equipment credentials
[**DeleteCredential**](NetworkequipmentAPI.md#DeleteCredential) | **Delete** /networkEquipments/{networkEquipmentId}/credentials/{type}/{username} | Delete network equipment credentials
[**GetCredential**](NetworkequipmentAPI.md#GetCredential) | **Get** /networkEquipments/{networkEquipmentId}/credentials/{type}/{username} | Show network equipment credentials
[**GetCredentialList**](NetworkequipmentAPI.md#GetCredentialList) | **Get** /networkEquipments/{networkEquipmentId}/credentials | List network equipment credentials
[**GetCredentialListByType**](NetworkequipmentAPI.md#GetCredentialListByType) | **Get** /networkEquipments/{networkEquipmentId}/credentials/{type} | List network equipment credentials by type
[**GetIp**](NetworkequipmentAPI.md#GetIp) | **Get** /networkEquipments/{networkEquipmentId}/ips/{ip} | Show a network equipment IP
[**GetIpList**](NetworkequipmentAPI.md#GetIpList) | **Get** /networkEquipments/{networkEquipmentId}/ips | List IPs
[**GetNetworkEquipment**](NetworkequipmentAPI.md#GetNetworkEquipment) | **Get** /networkEquipments/{networkEquipmentId} | Get network equipment
[**GetNetworkEquipmentList**](NetworkequipmentAPI.md#GetNetworkEquipmentList) | **Get** /networkEquipments | List network equipment
[**GetNullRouteHistory**](NetworkequipmentAPI.md#GetNullRouteHistory) | **Get** /networkEquipments/{networkEquipmentId}/nullRouteHistory | Show null route history
[**GetPowerStatus**](NetworkequipmentAPI.md#GetPowerStatus) | **Get** /networkEquipments/{networkEquipmentId}/powerInfo | Show power status
[**NullIpRoute**](NetworkequipmentAPI.md#NullIpRoute) | **Post** /networkEquipments/{networkEquipmentId}/ips/{ip}/null | Null route an IP
[**PowerCycle**](NetworkequipmentAPI.md#PowerCycle) | **Post** /networkEquipments/{networkEquipmentId}/powerCycle | Power cycle a network equipment
[**PowerOff**](NetworkequipmentAPI.md#PowerOff) | **Post** /networkEquipments/{networkEquipmentId}/powerOff | Power off network equipment
[**PowerOn**](NetworkequipmentAPI.md#PowerOn) | **Post** /networkEquipments/{networkEquipmentId}/powerOn | Power on network equipment
[**UnNullIpRoute**](NetworkequipmentAPI.md#UnNullIpRoute) | **Post** /networkEquipments/{networkEquipmentId}/ips/{ip}/unnull | Remove a null route
[**UpdateCredential**](NetworkequipmentAPI.md#UpdateCredential) | **Put** /networkEquipments/{networkEquipmentId}/credentials/{type}/{username} | Update network equipment credentials
[**UpdateIp**](NetworkequipmentAPI.md#UpdateIp) | **Put** /networkEquipments/{networkEquipmentId}/ips/{ip} | Update an IP
[**UpdateReference**](NetworkequipmentAPI.md#UpdateReference) | **Put** /networkEquipments/{networkEquipmentId} | Update network equipment



## CreateCredential

> Credential CreateCredential(ctx, networkEquipmentId).CreateCredentialOpts(createCredentialOpts).Execute()

Create new network equipment credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	createCredentialOpts := *openapiclient.NewCreateCredentialOpts("Password_example", openapiclient.credentialType("OPERATING_SYSTEM"), "Username_example") // CreateCredentialOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.CreateCredential(context.Background(), networkEquipmentId).CreateCredentialOpts(createCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.CreateCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.CreateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

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


## DeleteCredential

> DeleteCredential(ctx, networkEquipmentId, type_, username).Execute()

Delete network equipment credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	type_ := openapiclient.credentialType("OPERATING_SYSTEM") // CredentialType | The type of the credential.
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkequipmentAPI.DeleteCredential(context.Background(), networkEquipmentId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.DeleteCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
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


## GetCredential

> Credential GetCredential(ctx, networkEquipmentId, type_, username).Execute()

Show network equipment credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	type_ := openapiclient.credentialType("OPERATING_SYSTEM") // CredentialType | The type of the credential.
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.GetCredential(context.Background(), networkEquipmentId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.GetCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.GetCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
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

> CredentialList GetCredentialList(ctx, networkEquipmentId).Limit(limit).Offset(offset).Execute()

List network equipment credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.GetCredentialList(context.Background(), networkEquipmentId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.GetCredentialList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialList`: CredentialList
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.GetCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

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

> CredentialList GetCredentialListByType(ctx, networkEquipmentId, type_).Limit(limit).Offset(offset).Execute()

List network equipment credentials by type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	type_ := openapiclient.credentialType("OPERATING_SYSTEM") // CredentialType | The type of the credential.
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.GetCredentialListByType(context.Background(), networkEquipmentId, type_).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.GetCredentialListByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialListByType`: CredentialList
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.GetCredentialListByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
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


## GetIp

> Ip GetIp(ctx, networkEquipmentId, ip).Execute()

Show a network equipment IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.GetIp(context.Background(), networkEquipmentId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.GetIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIp`: Ip
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.GetIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
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

> IpList GetIpList(ctx, networkEquipmentId).NetworkType(networkType).Version(version).NullRouted(nullRouted).Ips(ips).Limit(limit).Offset(offset).Execute()

List IPs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	networkType := openapiclient.networkType("INTERNAL") // NetworkType | Filter the collection of ip addresses by network type (optional)
	version := "version_example" // string | Filter the collection by ip version (optional)
	nullRouted := "nullRouted_example" // string | Filter Ips by Nulled-Status (optional)
	ips := "ips_example" // string | Filter the collection of Ips for the comma separated list of Ips (optional)
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.GetIpList(context.Background(), networkEquipmentId).NetworkType(networkType).Version(version).NullRouted(nullRouted).Ips(ips).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.GetIpList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpList`: IpList
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.GetIpList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

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


## GetNetworkEquipment

> NetworkEquipment GetNetworkEquipment(ctx, networkEquipmentId).Execute()

Get network equipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.GetNetworkEquipment(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.GetNetworkEquipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipment`: NetworkEquipment
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.GetNetworkEquipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEquipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkEquipment**](NetworkEquipment.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEquipmentList

> GetNetworkEquipmentListResult GetNetworkEquipmentList(ctx).Limit(limit).Offset(offset).Reference(reference).Ip(ip).MacAddress(macAddress).Site(site).PrivateRackId(privateRackId).PrivateNetworkCapable(privateNetworkCapable).PrivateNetworkEnabled(privateNetworkEnabled).Execute()

List network equipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	reference := "my-switch" // string | Filter the list of network equipment by reference. (optional)
	ip := "127.0.0.4" // string | Filter the list of network equipment by ip address. (optional)
	macAddress := "aa:bb:cc:dd:ee:ff" // string | Filter the list of network equipment by mac address. (optional)
	site := "FRA-10" // string | Filter the list of network equipment by site (location). (optional)
	privateRackId := "123" // string | Filter the list of network equipment by dedicated rack id. (optional)
	privateNetworkCapable := "privateNetworkCapable_example" // string | Filter the list for private network capable network equipment (optional)
	privateNetworkEnabled := "privateNetworkEnabled_example" // string | Filter the list for private network enabled network equipment (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.GetNetworkEquipmentList(context.Background()).Limit(limit).Offset(offset).Reference(reference).Ip(ip).MacAddress(macAddress).Site(site).PrivateRackId(privateRackId).PrivateNetworkCapable(privateNetworkCapable).PrivateNetworkEnabled(privateNetworkEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.GetNetworkEquipmentList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentList`: GetNetworkEquipmentListResult
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.GetNetworkEquipmentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEquipmentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **reference** | **string** | Filter the list of network equipment by reference. | 
 **ip** | **string** | Filter the list of network equipment by ip address. | 
 **macAddress** | **string** | Filter the list of network equipment by mac address. | 
 **site** | **string** | Filter the list of network equipment by site (location). | 
 **privateRackId** | **string** | Filter the list of network equipment by dedicated rack id. | 
 **privateNetworkCapable** | **string** | Filter the list for private network capable network equipment | 
 **privateNetworkEnabled** | **string** | Filter the list for private network enabled network equipment | 

### Return type

[**GetNetworkEquipmentListResult**](GetNetworkEquipmentListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNullRouteHistory

> NullRouteList GetNullRouteHistory(ctx, networkEquipmentId).Limit(limit).Offset(offset).Execute()

Show null route history



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.GetNullRouteHistory(context.Background(), networkEquipmentId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.GetNullRouteHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNullRouteHistory`: NullRouteList
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.GetNullRouteHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

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


## GetPowerStatus

> GetPowerStatusResult GetPowerStatus(ctx, networkEquipmentId).Execute()

Show power status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.GetPowerStatus(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.GetPowerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPowerStatus`: GetPowerStatusResult
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.GetPowerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

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


## NullIpRoute

> Ip NullIpRoute(ctx, networkEquipmentId, ip).Execute()

Null route an IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.NullIpRoute(context.Background(), networkEquipmentId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.NullIpRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NullIpRoute`: Ip
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.NullIpRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
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


## PowerCycle

> PowerCycle(ctx, networkEquipmentId).Execute()

Power cycle a network equipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkequipmentAPI.PowerCycle(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.PowerCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

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

> PowerOff(ctx, networkEquipmentId).Execute()

Power off network equipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkequipmentAPI.PowerOff(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.PowerOff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

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

> PowerOn(ctx, networkEquipmentId).Execute()

Power on network equipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkequipmentAPI.PowerOn(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.PowerOn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

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


## UnNullIpRoute

> Ip UnNullIpRoute(ctx, networkEquipmentId, ip).Execute()

Remove a null route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.UnNullIpRoute(context.Background(), networkEquipmentId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.UnNullIpRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnNullIpRoute`: Ip
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.UnNullIpRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnNullIpRouteRequest struct via the builder pattern


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


## UpdateCredential

> Credential UpdateCredential(ctx, networkEquipmentId, type_, username).UpdateCredentialOpts(updateCredentialOpts).Execute()

Update network equipment credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	type_ := openapiclient.credentialType("OPERATING_SYSTEM") // CredentialType | The type of the credential.
	username := "root" // string | Username
	updateCredentialOpts := *openapiclient.NewUpdateCredentialOpts("Password_example") // UpdateCredentialOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.UpdateCredential(context.Background(), networkEquipmentId, type_, username).UpdateCredentialOpts(updateCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.UpdateCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.UpdateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
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


## UpdateIp

> Ip UpdateIp(ctx, networkEquipmentId, ip).UpdateIpOpts(updateIpOpts).Execute()

Update an IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	ip := "127.0.0.6" // string | The IP Address
	updateIpOpts := *openapiclient.NewUpdateIpOpts() // UpdateIpOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkequipmentAPI.UpdateIp(context.Background(), networkEquipmentId, ip).UpdateIpOpts(updateIpOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.UpdateIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIp`: Ip
	fmt.Fprintf(os.Stdout, "Response from `NetworkequipmentAPI.UpdateIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateIpOpts** | [**UpdateIpOpts**](UpdateIpOpts.md) |  | 

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

> UpdateReference(ctx, networkEquipmentId).UpdateReferenceOpts(updateReferenceOpts).Execute()

Update network equipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/networkequipment"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	updateReferenceOpts := *openapiclient.NewUpdateReferenceOpts("Reference_example") // UpdateReferenceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkequipmentAPI.UpdateReference(context.Background(), networkEquipmentId).UpdateReferenceOpts(updateReferenceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkequipmentAPI.UpdateReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

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

