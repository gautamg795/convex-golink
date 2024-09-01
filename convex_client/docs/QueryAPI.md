# \QueryAPI

All URIs are relative to *https://tacit-grouse-50.convex.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRunLoadLoadAllPost**](QueryAPI.md#ApiRunLoadLoadAllPost) | **Post** /api/run/load/loadAll | Calls a query at the path load.js:loadAll
[**ApiRunLoadLoadOnePost**](QueryAPI.md#ApiRunLoadLoadOnePost) | **Post** /api/run/load/loadOne | Calls a query at the path load.js:loadOne
[**ApiRunStatsLoadStatsPost**](QueryAPI.md#ApiRunStatsLoadStatsPost) | **Post** /api/run/stats/loadStats | Calls a query at the path stats.js:loadStats



## ApiRunLoadLoadAllPost

> ResponseLoadLoadAll ApiRunLoadLoadAllPost(ctx).RequestLoadLoadAll(requestLoadLoadAll).Execute()

Calls a query at the path load.js:loadAll

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	requestLoadLoadAll := *openapiclient.NewRequestLoadLoadAll(*openapiclient.NewRequestClearDefaultArgs("Token_example")) // RequestLoadLoadAll | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryAPI.ApiRunLoadLoadAllPost(context.Background()).RequestLoadLoadAll(requestLoadLoadAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryAPI.ApiRunLoadLoadAllPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiRunLoadLoadAllPost`: ResponseLoadLoadAll
	fmt.Fprintf(os.Stdout, "Response from `QueryAPI.ApiRunLoadLoadAllPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRunLoadLoadAllPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestLoadLoadAll** | [**RequestLoadLoadAll**](RequestLoadLoadAll.md) |  | 

### Return type

[**ResponseLoadLoadAll**](ResponseLoadLoadAll.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRunLoadLoadOnePost

> ResponseLoadLoadOne ApiRunLoadLoadOnePost(ctx).RequestLoadLoadOne(requestLoadLoadOne).Execute()

Calls a query at the path load.js:loadOne

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	requestLoadLoadOne := *openapiclient.NewRequestLoadLoadOne(*openapiclient.NewRequestClearDeleteOneArgs("NormalizedId_example", "Token_example")) // RequestLoadLoadOne | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryAPI.ApiRunLoadLoadOnePost(context.Background()).RequestLoadLoadOne(requestLoadLoadOne).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryAPI.ApiRunLoadLoadOnePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiRunLoadLoadOnePost`: ResponseLoadLoadOne
	fmt.Fprintf(os.Stdout, "Response from `QueryAPI.ApiRunLoadLoadOnePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRunLoadLoadOnePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestLoadLoadOne** | [**RequestLoadLoadOne**](RequestLoadLoadOne.md) |  | 

### Return type

[**ResponseLoadLoadOne**](ResponseLoadLoadOne.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRunStatsLoadStatsPost

> ResponseStatsLoadStats ApiRunStatsLoadStatsPost(ctx).RequestStatsLoadStats(requestStatsLoadStats).Execute()

Calls a query at the path stats.js:loadStats

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	requestStatsLoadStats := *openapiclient.NewRequestStatsLoadStats(*openapiclient.NewRequestClearDefaultArgs("Token_example")) // RequestStatsLoadStats | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryAPI.ApiRunStatsLoadStatsPost(context.Background()).RequestStatsLoadStats(requestStatsLoadStats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryAPI.ApiRunStatsLoadStatsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiRunStatsLoadStatsPost`: ResponseStatsLoadStats
	fmt.Fprintf(os.Stdout, "Response from `QueryAPI.ApiRunStatsLoadStatsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRunStatsLoadStatsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestStatsLoadStats** | [**RequestStatsLoadStats**](RequestStatsLoadStats.md) |  | 

### Return type

[**ResponseStatsLoadStats**](ResponseStatsLoadStats.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

