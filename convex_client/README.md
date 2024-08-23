# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.0
- Package version: 1.0.0
- Generator version: 7.8.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://tacit-grouse-50.convex.cloud*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*MutationAPI* | [**ApiRunClearDefaultPost**](docs/MutationAPI.md#apiruncleardefaultpost) | **Post** /api/run/clear/default | Calls a mutation at the path clear.js:default
*MutationAPI* | [**ApiRunStatsSaveStatsPost**](docs/MutationAPI.md#apirunstatssavestatspost) | **Post** /api/run/stats/saveStats | Calls a mutation at the path stats.js:saveStats
*MutationAPI* | [**ApiRunStoreDefaultPost**](docs/MutationAPI.md#apirunstoredefaultpost) | **Post** /api/run/store/default | Calls a mutation at the path store.js:default
*QueryAPI* | [**ApiRunLoadLoadAllPost**](docs/QueryAPI.md#apirunloadloadallpost) | **Post** /api/run/load/loadAll | Calls a query at the path load.js:loadAll
*QueryAPI* | [**ApiRunLoadLoadOnePost**](docs/QueryAPI.md#apirunloadloadonepost) | **Post** /api/run/load/loadOne | Calls a query at the path load.js:loadOne
*QueryAPI* | [**ApiRunStatsLoadStatsPost**](docs/QueryAPI.md#apirunstatsloadstatspost) | **Post** /api/run/stats/loadStats | Calls a query at the path stats.js:loadStats


## Documentation For Models

 - [RequestClearDefault](docs/RequestClearDefault.md)
 - [RequestClearDefaultArgs](docs/RequestClearDefaultArgs.md)
 - [RequestLoadLoadAll](docs/RequestLoadLoadAll.md)
 - [RequestLoadLoadOne](docs/RequestLoadLoadOne.md)
 - [RequestLoadLoadOneArgs](docs/RequestLoadLoadOneArgs.md)
 - [RequestStatsLoadStats](docs/RequestStatsLoadStats.md)
 - [RequestStatsSaveStats](docs/RequestStatsSaveStats.md)
 - [RequestStatsSaveStatsArgs](docs/RequestStatsSaveStatsArgs.md)
 - [RequestStoreDefault](docs/RequestStoreDefault.md)
 - [RequestStoreDefaultArgs](docs/RequestStoreDefaultArgs.md)
 - [RequestStoreDefaultArgsLink](docs/RequestStoreDefaultArgsLink.md)
 - [ResponseClearDefault](docs/ResponseClearDefault.md)
 - [ResponseLoadLoadAll](docs/ResponseLoadLoadAll.md)
 - [ResponseLoadLoadAllValueInner](docs/ResponseLoadLoadAllValueInner.md)
 - [ResponseLoadLoadOne](docs/ResponseLoadLoadOne.md)
 - [ResponseLoadLoadOneValue](docs/ResponseLoadLoadOneValue.md)
 - [ResponseStatsLoadStats](docs/ResponseStatsLoadStats.md)
 - [ResponseStatsSaveStats](docs/ResponseStatsSaveStats.md)
 - [ResponseStoreDefault](docs/ResponseStoreDefault.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextAccessToken, "BEARER_TOKEN_STRING")
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


