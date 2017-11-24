# \PlansApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindPlans**](PlansApi.md#FindPlans) | **Get** /plans | Retrieve all plans
[**FindPlansByProject**](PlansApi.md#FindPlansByProject) | **Get** /projects/{id}/plans | Retrieve all plans visible by the project


# **FindPlans**
> PlanList FindPlans($include)

Retrieve all plans

Provides a listing of available services plans available to provision your device on.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**PlanList**](PlanList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindPlansByProject**
> PlanList FindPlansByProject($id, $include, $page, $perPage)

Retrieve all plans visible by the project

Returns a listing of available services plans available for the given project


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **include** | **string**| related attributes to include | [optional] 
 **page** | **int32**| page to display, default to 1, max 100_000 | [optional] 
 **perPage** | **int32**| items per page, default to 10, max 1_000 | [optional] 

### Return type

[**PlanList**](PlanList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

