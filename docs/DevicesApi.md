# \DevicesApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDevice**](DevicesApi.md#CreateDevice) | **Post** /projects/{id}/devices | Create a device
[**DeleteDevice**](DevicesApi.md#DeleteDevice) | **Delete** /devices/{id} | Delete the device
[**FindDeviceById**](DevicesApi.md#FindDeviceById) | **Get** /devices/{id} | Retrieve a device
[**FindDevices**](DevicesApi.md#FindDevices) | **Get** /projects/{id}/devices | Retrieve all devices
[**UpdateDevice**](DevicesApi.md#UpdateDevice) | **Put** /devices/{id} | Update the device


# **CreateDevice**
> Device CreateDevice($id, $device)

Create a device

Creates a new device and provisions it in our datacenter.  Type-specific options (such as operating_system for baremetal devices) should be included in the main data structure alongside hostname and plan.  The features attribute allows you to optionally specify what features your server should have. For example, if you require a server with a TPM chip, you may specify { \"features\": { \"tpm\": \"required\" } } (or { \"features\": [\"tpm\"] } in shorthand). The request will fail if there are no available servers matching your criteria. Alternatively, if you do not require a certain feature, but would prefer to be assigned a server with that feature if there are any available, you may specify that feature with a preferred value (see the example request below). The request will not fail if we have no servers with that feature in our inventory.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **device** | [**DeviceCreateInput**](DeviceCreateInput.md)| Device to create | 

### Return type

[**Device**](Device.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDevice**
> DeleteDevice($id)

Delete the device

Deletes a device and deprovisions it in our datacenter.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Device UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindDeviceById**
> Device FindDeviceById($id, $include)

Retrieve a device

Type-specific options (such as facility for baremetal devices) will be included as part of the main data structure.  state value can be one of: active inactive queued or provisioning


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Device UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**Device**](Device.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindDevices**
> DeviceList FindDevices($id, $include, $page, $perPage)

Retrieve all devices

Provides a collection of devices for a given project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **include** | **string**| related attributes to include | [optional] 
 **page** | **int32**| page to display, default to 1, max 100_000 | [optional] 
 **perPage** | **int32**| items per page, default to 10, max 1_000 | [optional] 

### Return type

[**DeviceList**](DeviceList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDevice**
> Device UpdateDevice($id, $device)

Update the device

Updates the device.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Device UUID | 
 **device** | [**DeviceUpdateInput**](DeviceUpdateInput.md)| Facility to update | 

### Return type

[**Device**](Device.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

