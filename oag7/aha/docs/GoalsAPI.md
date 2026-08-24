# \GoalsAPI

All URIs are relative to *https://secure.aha.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProductGoal**](GoalsAPI.md#CreateProductGoal) | **Post** /products/{product_id}/goals | Create goal in product
[**GetGoal**](GoalsAPI.md#GetGoal) | **Get** /goals/{goal_id} | Get goal
[**ListGoals**](GoalsAPI.md#ListGoals) | **Get** /goals | List goals
[**ListProductGoals**](GoalsAPI.md#ListProductGoals) | **Get** /products/{product_id}/goals | List product goals
[**UpdateGoal**](GoalsAPI.md#UpdateGoal) | **Put** /goals/{goal_id} | Update goal



## CreateProductGoal

> GoalResponse CreateProductGoal(ctx, productId).GoalCreateRequest(goalCreateRequest).Execute()

Create goal in product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	productId := "productId_example" // string | Product ID or reference prefix
	goalCreateRequest := *openapiclient.NewGoalCreateRequest(*openapiclient.NewGoalCreate("Name_example")) // GoalCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoalsAPI.CreateProductGoal(context.Background(), productId).GoalCreateRequest(goalCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.CreateProductGoal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProductGoal`: GoalResponse
	fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.CreateProductGoal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product ID or reference prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductGoalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **goalCreateRequest** | [**GoalCreateRequest**](GoalCreateRequest.md) |  | 

### Return type

[**GoalResponse**](GoalResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGoal

> GoalResponse GetGoal(ctx, goalId).Execute()

Get goal



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	goalId := "goalId_example" // string | Goal ID or reference number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoalsAPI.GetGoal(context.Background(), goalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.GetGoal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoal`: GoalResponse
	fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.GetGoal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**goalId** | **string** | Goal ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GoalResponse**](GoalResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGoals

> GoalsResponse ListGoals(ctx).Q(q).UpdatedSince(updatedSince).Page(page).PerPage(perPage).Execute()

List goals



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	q := "q_example" // string | Sub-string to match against goal name (optional)
	updatedSince := time.Now() // time.Time | UTC timestamp (ISO8601). Only goals updated after this time. (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoalsAPI.ListGoals(context.Background()).Q(q).UpdatedSince(updatedSince).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.ListGoals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGoals`: GoalsResponse
	fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.ListGoals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGoalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Sub-string to match against goal name | 
 **updatedSince** | **time.Time** | UTC timestamp (ISO8601). Only goals updated after this time. | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**GoalsResponse**](GoalsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductGoals

> GoalsResponse ListProductGoals(ctx, productId).Page(page).PerPage(perPage).Execute()

List product goals



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	productId := "productId_example" // string | Product ID or reference prefix
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoalsAPI.ListProductGoals(context.Background(), productId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.ListProductGoals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductGoals`: GoalsResponse
	fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.ListProductGoals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product ID or reference prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProductGoalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**GoalsResponse**](GoalsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGoal

> GoalResponse UpdateGoal(ctx, goalId).GoalUpdateRequest(goalUpdateRequest).Execute()

Update goal



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grokify/go-aha/v3/oag7/aha"
)

func main() {
	goalId := "goalId_example" // string | Goal ID or reference number
	goalUpdateRequest := *openapiclient.NewGoalUpdateRequest(*openapiclient.NewGoalUpdate()) // GoalUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoalsAPI.UpdateGoal(context.Background(), goalId).GoalUpdateRequest(goalUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.UpdateGoal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGoal`: GoalResponse
	fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.UpdateGoal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**goalId** | **string** | Goal ID or reference number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGoalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **goalUpdateRequest** | [**GoalUpdateRequest**](GoalUpdateRequest.md) |  | 

### Return type

[**GoalResponse**](GoalResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

