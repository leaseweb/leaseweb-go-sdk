# Go API client for ordering

This document outlines the ordering API. The API is only available for customers with invoice payment method.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Generator version: 7.10.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import ordering "github.com/leaseweb/leaseweb-go-sdk/ordering"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `ordering.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), ordering.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `ordering.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), ordering.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `ordering.ContextOperationServerIndices` and `ordering.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), ordering.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), ordering.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.leaseweb.com/ordering/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OrderingAPI* | [**GetDedicatedServer**](docs/OrderingAPI.md#getdedicatedserver) | **Get** /products/dedicatedServers/{dedicatedServerId} | Get a single dedicated server and its price.
*OrderingAPI* | [**GetDedicatedServerList**](docs/OrderingAPI.md#getdedicatedserverlist) | **Get** /products/dedicatedServers | List available dedicated server configurations.
*OrderingAPI* | [**OrderDedicatedServer**](docs/OrderingAPI.md#orderdedicatedserver) | **Post** /products/dedicatedServers/{dedicatedServerId}/order | Dedicated Server ordering.


## Documentation For Models

 - [ContractTerm](docs/ContractTerm.md)
 - [Cpu](docs/Cpu.md)
 - [DedicatedServer](docs/DedicatedServer.md)
 - [DedicatedServerDetail](docs/DedicatedServerDetail.md)
 - [DedicatedServerList](docs/DedicatedServerList.md)
 - [DedicatedServerOrder](docs/DedicatedServerOrder.md)
 - [DedicatedServerPrice](docs/DedicatedServerPrice.md)
 - [DedicatedServerPricePrice](docs/DedicatedServerPricePrice.md)
 - [DedicatedServerPricePriceDiscounts](docs/DedicatedServerPricePriceDiscounts.md)
 - [Discount](docs/Discount.md)
 - [ErrorResult](docs/ErrorResult.md)
 - [Hdd](docs/Hdd.md)
 - [Metadata](docs/Metadata.md)
 - [OrderDedicatedServerOpts](docs/OrderDedicatedServerOpts.md)
 - [Ram](docs/Ram.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), ordering.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

### X-LSW-Auth

- **Type**: API key
- **API key parameter name**: X-LSW-Auth
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-LSW-Auth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		ordering.ContextAPIKeys,
		map[string]ordering.APIKey{
			"X-LSW-Auth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



