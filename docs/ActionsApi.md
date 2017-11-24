# \ActionsApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PerformAction**](ActionsApi.md#PerformAction) | **Post** /devices/{id}/actions | Perform an action


# **PerformAction**
> PerformAction($id, $type_)

Perform an action

Performs an action for the given device.  Possible actions include: power_on, power_off, reboot, and rescue (reboot the device into rescue OS.)


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Device UUID | 
 **type_** | **string**| Action to perform | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

