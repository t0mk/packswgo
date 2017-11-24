# \InstanceBandwidthApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindInstanceBandwidth**](InstanceBandwidthApi.md#FindInstanceBandwidth) | **Get** /devices/{id}/bandwidth | Retrieve an instance bandwidth


# **FindInstanceBandwidth**
> FindInstanceBandwidth($id, $from, $until)

Retrieve an instance bandwidth

Retrieve an instance bandwidth for a given period of time.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Device UUID | 
 **from** | **string**| Timestamp from range | 
 **until** | **string**| Timestamp to range | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

