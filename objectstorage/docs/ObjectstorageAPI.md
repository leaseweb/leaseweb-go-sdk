# \ObjectstorageAPI

All URIs are relative to *https://api.leaseweb.com/objectStorage/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectStorageAccessKey**](ObjectstorageAPI.md#CreateObjectStorageAccessKey) | **Post** /objectStorages/{objectStorageId}/users/{userId}/accessKeys | Create a user access key
[**CreateObjectStorageBucket**](ObjectstorageAPI.md#CreateObjectStorageBucket) | **Post** /objectStorages/{objectStorageId}/buckets | Create an object storage bucket
[**CreateObjectStorageGroup**](ObjectstorageAPI.md#CreateObjectStorageGroup) | **Post** /objectStorages/{objectStorageId}/groups | Create an object storage group
[**CreateObjectStorageUser**](ObjectstorageAPI.md#CreateObjectStorageUser) | **Post** /objectStorages/{objectStorageId}/users | Create an object storage user
[**DeleteObjectStorageAccessKey**](ObjectstorageAPI.md#DeleteObjectStorageAccessKey) | **Delete** /objectStorages/{objectStorageId}/users/{userId}/accessKeys/{accessKeyId} | Delete a user access key
[**DeleteObjectStorageBucket**](ObjectstorageAPI.md#DeleteObjectStorageBucket) | **Delete** /objectStorages/{objectStorageId}/buckets/{bucketName} | Delete an object storage bucket
[**DeleteObjectStorageGroup**](ObjectstorageAPI.md#DeleteObjectStorageGroup) | **Delete** /objectStorages/{objectStorageId}/groups/{groupId} | Delete an object storage group
[**DeleteObjectStorageUser**](ObjectstorageAPI.md#DeleteObjectStorageUser) | **Delete** /objectStorages/{objectStorageId}/users/{userId} | Delete an object storage user
[**GetObjectStorageAccessKeyList**](ObjectstorageAPI.md#GetObjectStorageAccessKeyList) | **Get** /objectStorages/{objectStorageId}/users/{userId}/accessKeys | List user access keys
[**GetObjectStorageBucketList**](ObjectstorageAPI.md#GetObjectStorageBucketList) | **Get** /objectStorages/{objectStorageId}/buckets | List object storage buckets
[**GetObjectStorageGroup**](ObjectstorageAPI.md#GetObjectStorageGroup) | **Get** /objectStorages/{objectStorageId}/groups/{groupId} | Get an object storage group
[**GetObjectStorageGroupList**](ObjectstorageAPI.md#GetObjectStorageGroupList) | **Get** /objectStorages/{objectStorageId}/groups | List object storage groups
[**GetObjectStorageList**](ObjectstorageAPI.md#GetObjectStorageList) | **Get** /objectStorages | List all Object Storages
[**GetObjectStorageUser**](ObjectstorageAPI.md#GetObjectStorageUser) | **Get** /objectStorages/{objectStorageId}/users/{userId} | Get an object storage user
[**GetObjectStorageUserList**](ObjectstorageAPI.md#GetObjectStorageUserList) | **Get** /objectStorages/{objectStorageId}/users | List object storage users
[**UpdateObjectStorageBucket**](ObjectstorageAPI.md#UpdateObjectStorageBucket) | **Put** /objectStorages/{objectStorageId}/buckets/{bucketName} | Update an object storage bucket
[**UpdateObjectStorageGroup**](ObjectstorageAPI.md#UpdateObjectStorageGroup) | **Put** /objectStorages/{objectStorageId}/groups/{groupId} | Update an object storage group
[**UpdateObjectStorageUser**](ObjectstorageAPI.md#UpdateObjectStorageUser) | **Put** /objectStorages/{objectStorageId}/users/{userId} | Update an object storage user



## CreateObjectStorageAccessKey

> CreateObjectStorageAccessKeyResult CreateObjectStorageAccessKey(ctx, objectStorageId, userId).CreateObjectStorageAccessKeyOpts(createObjectStorageAccessKeyOpts).Execute()

Create a user access key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	userId := "userId_example" // string | The user id.
	createObjectStorageAccessKeyOpts := *openapiclient.NewCreateObjectStorageAccessKeyOpts() // CreateObjectStorageAccessKeyOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.CreateObjectStorageAccessKey(context.Background(), objectStorageId, userId).CreateObjectStorageAccessKeyOpts(createObjectStorageAccessKeyOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.CreateObjectStorageAccessKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateObjectStorageAccessKey`: CreateObjectStorageAccessKeyResult
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.CreateObjectStorageAccessKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectStorageAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createObjectStorageAccessKeyOpts** | [**CreateObjectStorageAccessKeyOpts**](CreateObjectStorageAccessKeyOpts.md) |  | 

### Return type

[**CreateObjectStorageAccessKeyResult**](CreateObjectStorageAccessKeyResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateObjectStorageBucket

> Bucket CreateObjectStorageBucket(ctx, objectStorageId).CreateObjectStorageBucketOpts(createObjectStorageBucketOpts).Execute()

Create an object storage bucket



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	createObjectStorageBucketOpts := *openapiclient.NewCreateObjectStorageBucketOpts("Name_example", false) // CreateObjectStorageBucketOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.CreateObjectStorageBucket(context.Background(), objectStorageId).CreateObjectStorageBucketOpts(createObjectStorageBucketOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.CreateObjectStorageBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateObjectStorageBucket`: Bucket
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.CreateObjectStorageBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectStorageBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createObjectStorageBucketOpts** | [**CreateObjectStorageBucketOpts**](CreateObjectStorageBucketOpts.md) |  | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateObjectStorageGroup

> Group CreateObjectStorageGroup(ctx, objectStorageId).CreateObjectStorageGroupOpts(createObjectStorageGroupOpts).Execute()

Create an object storage group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	createObjectStorageGroupOpts := *openapiclient.NewCreateObjectStorageGroupOpts("DisplayName_example", "UniqueName_example") // CreateObjectStorageGroupOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.CreateObjectStorageGroup(context.Background(), objectStorageId).CreateObjectStorageGroupOpts(createObjectStorageGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.CreateObjectStorageGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateObjectStorageGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.CreateObjectStorageGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectStorageGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createObjectStorageGroupOpts** | [**CreateObjectStorageGroupOpts**](CreateObjectStorageGroupOpts.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateObjectStorageUser

> User CreateObjectStorageUser(ctx, objectStorageId).CreateObjectStorageUserOpts(createObjectStorageUserOpts).Execute()

Create an object storage user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	createObjectStorageUserOpts := *openapiclient.NewCreateObjectStorageUserOpts("FullName_example", "UniqueName_example", []string{"Groups_example"}) // CreateObjectStorageUserOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.CreateObjectStorageUser(context.Background(), objectStorageId).CreateObjectStorageUserOpts(createObjectStorageUserOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.CreateObjectStorageUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateObjectStorageUser`: User
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.CreateObjectStorageUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectStorageUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createObjectStorageUserOpts** | [**CreateObjectStorageUserOpts**](CreateObjectStorageUserOpts.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObjectStorageAccessKey

> DeleteObjectStorageAccessKey(ctx, objectStorageId, userId, accessKeyId).Execute()

Delete a user access key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	userId := "userId_example" // string | The user id.
	accessKeyId := "accessKeyId_example" // string | The access key id, UUID format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectstorageAPI.DeleteObjectStorageAccessKey(context.Background(), objectStorageId, userId, accessKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.DeleteObjectStorageAccessKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 
**userId** | **string** | The user id. | 
**accessKeyId** | **string** | The access key id, UUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageAccessKeyRequest struct via the builder pattern


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


## DeleteObjectStorageBucket

> DeleteObjectStorageBucket(ctx, objectStorageId, bucketName).Execute()

Delete an object storage bucket



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	bucketName := "bucketName_example" // string | The bucket name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectstorageAPI.DeleteObjectStorageBucket(context.Background(), objectStorageId, bucketName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.DeleteObjectStorageBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 
**bucketName** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageBucketRequest struct via the builder pattern


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


## DeleteObjectStorageGroup

> DeleteObjectStorageGroup(ctx, objectStorageId, groupId).Execute()

Delete an object storage group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	groupId := "groupId_example" // string | The group id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectstorageAPI.DeleteObjectStorageGroup(context.Background(), objectStorageId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.DeleteObjectStorageGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 
**groupId** | **string** | The group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageGroupRequest struct via the builder pattern


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


## DeleteObjectStorageUser

> DeleteObjectStorageUser(ctx, objectStorageId, userId).Execute()

Delete an object storage user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	userId := "userId_example" // string | The user id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectstorageAPI.DeleteObjectStorageUser(context.Background(), objectStorageId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.DeleteObjectStorageUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageUserRequest struct via the builder pattern


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


## GetObjectStorageAccessKeyList

> GetObjectStorageAccessKeyListResult GetObjectStorageAccessKeyList(ctx, objectStorageId, userId).Execute()

List user access keys



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	userId := "userId_example" // string | The user id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.GetObjectStorageAccessKeyList(context.Background(), objectStorageId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.GetObjectStorageAccessKeyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageAccessKeyList`: GetObjectStorageAccessKeyListResult
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.GetObjectStorageAccessKeyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageAccessKeyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetObjectStorageAccessKeyListResult**](GetObjectStorageAccessKeyListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageBucketList

> GetObjectStorageBucketListResult GetObjectStorageBucketList(ctx, objectStorageId).Execute()

List object storage buckets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.GetObjectStorageBucketList(context.Background(), objectStorageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.GetObjectStorageBucketList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageBucketList`: GetObjectStorageBucketListResult
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.GetObjectStorageBucketList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageBucketListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetObjectStorageBucketListResult**](GetObjectStorageBucketListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageGroup

> Group GetObjectStorageGroup(ctx, objectStorageId, groupId).Execute()

Get an object storage group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	groupId := "groupId_example" // string | The group id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.GetObjectStorageGroup(context.Background(), objectStorageId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.GetObjectStorageGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.GetObjectStorageGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 
**groupId** | **string** | The group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Group**](Group.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageGroupList

> GetObjectStorageGroupListResult GetObjectStorageGroupList(ctx, objectStorageId).Execute()

List object storage groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.GetObjectStorageGroupList(context.Background(), objectStorageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.GetObjectStorageGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageGroupList`: GetObjectStorageGroupListResult
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.GetObjectStorageGroupList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetObjectStorageGroupListResult**](GetObjectStorageGroupListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageList

> GetObjectStorageListResult GetObjectStorageList(ctx).Limit(limit).Offset(offset).Execute()

List all Object Storages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	limit := int32(56) // int32 | Limit the number of results returned. (optional) (default to 10)
	offset := int32(56) // int32 | Return results starting from the given offset. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.GetObjectStorageList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.GetObjectStorageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageList`: GetObjectStorageListResult
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.GetObjectStorageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | [default to 10]
 **offset** | **int32** | Return results starting from the given offset. | [default to 0]

### Return type

[**GetObjectStorageListResult**](GetObjectStorageListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageUser

> User GetObjectStorageUser(ctx, objectStorageId, userId).Execute()

Get an object storage user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	userId := "userId_example" // string | The user id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.GetObjectStorageUser(context.Background(), objectStorageId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.GetObjectStorageUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageUser`: User
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.GetObjectStorageUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**User**](User.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageUserList

> GetObjectStorageUserListResult GetObjectStorageUserList(ctx, objectStorageId).Execute()

List object storage users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.GetObjectStorageUserList(context.Background(), objectStorageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.GetObjectStorageUserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageUserList`: GetObjectStorageUserListResult
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.GetObjectStorageUserList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetObjectStorageUserListResult**](GetObjectStorageUserListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateObjectStorageBucket

> Bucket UpdateObjectStorageBucket(ctx, objectStorageId, bucketName).UpdateObjectStorageBucketOpts(updateObjectStorageBucketOpts).Execute()

Update an object storage bucket



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	bucketName := "bucketName_example" // string | The bucket name.
	updateObjectStorageBucketOpts := *openapiclient.NewUpdateObjectStorageBucketOpts(false) // UpdateObjectStorageBucketOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.UpdateObjectStorageBucket(context.Background(), objectStorageId, bucketName).UpdateObjectStorageBucketOpts(updateObjectStorageBucketOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.UpdateObjectStorageBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateObjectStorageBucket`: Bucket
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.UpdateObjectStorageBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 
**bucketName** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectStorageBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateObjectStorageBucketOpts** | [**UpdateObjectStorageBucketOpts**](UpdateObjectStorageBucketOpts.md) |  | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateObjectStorageGroup

> Group UpdateObjectStorageGroup(ctx, objectStorageId, groupId).UpdateObjectStorageGroupOpts(updateObjectStorageGroupOpts).Execute()

Update an object storage group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	groupId := "groupId_example" // string | The group id.
	updateObjectStorageGroupOpts := *openapiclient.NewUpdateObjectStorageGroupOpts("DisplayName_example") // UpdateObjectStorageGroupOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.UpdateObjectStorageGroup(context.Background(), objectStorageId, groupId).UpdateObjectStorageGroupOpts(updateObjectStorageGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.UpdateObjectStorageGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateObjectStorageGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.UpdateObjectStorageGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 
**groupId** | **string** | The group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectStorageGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateObjectStorageGroupOpts** | [**UpdateObjectStorageGroupOpts**](UpdateObjectStorageGroupOpts.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateObjectStorageUser

> User UpdateObjectStorageUser(ctx, objectStorageId, userId).UpdateObjectStorageUserOpts(updateObjectStorageUserOpts).Execute()

Update an object storage user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/objectstorage"
)

func main() {
	objectStorageId := "10242480" // string | The id of the ObjectStorage.
	userId := "userId_example" // string | The user id.
	updateObjectStorageUserOpts := *openapiclient.NewUpdateObjectStorageUserOpts("FullName_example", []string{"Groups_example"}) // UpdateObjectStorageUserOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstorageAPI.UpdateObjectStorageUser(context.Background(), objectStorageId, userId).UpdateObjectStorageUserOpts(updateObjectStorageUserOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstorageAPI.UpdateObjectStorageUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateObjectStorageUser`: User
	fmt.Fprintf(os.Stdout, "Response from `ObjectstorageAPI.UpdateObjectStorageUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | The id of the ObjectStorage. | 
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectStorageUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateObjectStorageUserOpts** | [**UpdateObjectStorageUserOpts**](UpdateObjectStorageUserOpts.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

