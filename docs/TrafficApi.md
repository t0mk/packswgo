# \TrafficApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindTraffic**](TrafficApi.md#FindTraffic) | **Get** /devices/{id}/traffic | Retrieve device traffic


# **FindTraffic**
> FindTraffic($id, $direction, $timeframe, $interval, $bucket)

Retrieve device traffic

Returns traffic for a specific device.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Device UUID | 
 **direction** | **string**| Traffic direction | 
 **timeframe** | [**Timeframe**](Timeframe.md)| Traffic timeframe | 
 **interval** | **string**| Traffic interval | [optional] 
 **bucket** | **string**| Traffic bucket | [optional] 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

