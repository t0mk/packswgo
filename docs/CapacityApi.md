# \CapacityApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCapacity**](CapacityApi.md#CheckCapacity) | **Post** /capacity | Check capacity
[**FindCapacity**](CapacityApi.md#FindCapacity) | **Get** /capacity | View capacity


# **CheckCapacity**
> CheckCapacity($facility)

Check capacity

Validates if a deploy can be fulfilled.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **facility** | [**CapacityInput**](CapacityInput.md)| Facility to create | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindCapacity**
> CapacityList FindCapacity()

View capacity

Returns a list of facilities and plans with their current capacity.


### Parameters
This endpoint does not need any parameter.

### Return type

[**CapacityList**](CapacityList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

