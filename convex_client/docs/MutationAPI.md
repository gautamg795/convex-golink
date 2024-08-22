# \MutationAPI

All URIs are relative to *https://tacit-grouse-50.convex.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRunClearDefaultPost**](MutationAPI.md#ApiRunClearDefaultPost) | **Post** /api/run/clear/default | Calls a mutation at the path clear.js:default
[**ApiRunStatsSaveStatsPost**](MutationAPI.md#ApiRunStatsSaveStatsPost) | **Post** /api/run/stats/saveStats | Calls a mutation at the path stats.js:saveStats
[**ApiRunStoreDefaultPost**](MutationAPI.md#ApiRunStoreDefaultPost) | **Post** /api/run/store/default | Calls a mutation at the path store.js:default



## ApiRunClearDefaultPost

> ResponseClearDefault ApiRunClearDefaultPost(ctx).RequestClearDefault(requestClearDefault).Execute()

Calls a mutation at the path clear.js:default

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
	requestClearDefault := *openapiclient.NewRequestClearDefault(*openapiclient.NewRequestClearDefaultArgs("Token_example")) // RequestClearDefault | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutationAPI.ApiRunClearDefaultPost(context.Background()).RequestClearDefault(requestClearDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutationAPI.ApiRunClearDefaultPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiRunClearDefaultPost`: ResponseClearDefault
	fmt.Fprintf(os.Stdout, "Response from `MutationAPI.ApiRunClearDefaultPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRunClearDefaultPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestClearDefault** | [**RequestClearDefault**](RequestClearDefault.md) |  | 

### Return type

[**ResponseClearDefault**](ResponseClearDefault.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRunStatsSaveStatsPost

> ResponseStatsSaveStats ApiRunStatsSaveStatsPost(ctx).RequestStatsSaveStats(requestStatsSaveStats).Execute()

Calls a mutation at the path stats.js:saveStats

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
	requestStatsSaveStats := *openapiclient.NewRequestStatsSaveStats(*openapiclient.NewRequestStatsSaveStatsArgs(map[string]interface{}(123), "Token_example")) // RequestStatsSaveStats | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutationAPI.ApiRunStatsSaveStatsPost(context.Background()).RequestStatsSaveStats(requestStatsSaveStats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutationAPI.ApiRunStatsSaveStatsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiRunStatsSaveStatsPost`: ResponseStatsSaveStats
	fmt.Fprintf(os.Stdout, "Response from `MutationAPI.ApiRunStatsSaveStatsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRunStatsSaveStatsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestStatsSaveStats** | [**RequestStatsSaveStats**](RequestStatsSaveStats.md) |  | 

### Return type

[**ResponseStatsSaveStats**](ResponseStatsSaveStats.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRunStoreDefaultPost

> ResponseStoreDefault ApiRunStoreDefaultPost(ctx).RequestStoreDefault(requestStoreDefault).Execute()

Calls a mutation at the path store.js:default

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
	requestStoreDefault := *openapiclient.NewRequestStoreDefault(*openapiclient.NewRequestStoreDefaultArgs(*openapiclient.NewResponseLoadLoadAllValueInner(float32(123), float32(123), "Long_example", "NormalizedId_example", "Owner_example", "Short_example"), "Token_example")) // RequestStoreDefault | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutationAPI.ApiRunStoreDefaultPost(context.Background()).RequestStoreDefault(requestStoreDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutationAPI.ApiRunStoreDefaultPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiRunStoreDefaultPost`: ResponseStoreDefault
	fmt.Fprintf(os.Stdout, "Response from `MutationAPI.ApiRunStoreDefaultPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRunStoreDefaultPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestStoreDefault** | [**RequestStoreDefault**](RequestStoreDefault.md) |  | 

### Return type

[**ResponseStoreDefault**](ResponseStoreDefault.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

