# \VirtualNetworksApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignPort**](VirtualNetworksApi.md#AssignPort) | **Post** /ports/{id}/assign | Assign a port to virtual network
[**BondPort**](VirtualNetworksApi.md#BondPort) | **Post** /ports/{id}/bond | Enabling bonding
[**CreateVirtualNetwork**](VirtualNetworksApi.md#CreateVirtualNetwork) | **Post** /projects/{id}/virtual-networks | Create an virtual network
[**DeleteVirtualNetwork**](VirtualNetworksApi.md#DeleteVirtualNetwork) | **Delete** /virtual-networks/{id} | Delete a virtual network
[**DisbondPort**](VirtualNetworksApi.md#DisbondPort) | **Post** /ports/{id}/disbond | Disabling bonding
[**FindVirtualNetworks**](VirtualNetworksApi.md#FindVirtualNetworks) | **Get** /projects/{id}/virtual-networks | Retrieve all virtual networks
[**UnassignPort**](VirtualNetworksApi.md#UnassignPort) | **Post** /ports/{id}/unassign | Unassign a port


# **AssignPort**
> Port AssignPort($id, $vnid)

Assign a port to virtual network

Assign a port for a hardware to virtual network.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Port UUID | 
 **vnid** | **string**| Virtual Network ID | 

### Return type

[**Port**](Port.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BondPort**
> Port BondPort($id, $bulkEnable)

Enabling bonding

Enabling bonding for one or all ports


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Port UUID | 
 **bulkEnable** | **bool**| enable both ports | [optional] 

### Return type

[**Port**](Port.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVirtualNetwork**
> VirtualNetwork CreateVirtualNetwork($id, $virtualNetwork)

Create an virtual network

Creates an virtual network.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **virtualNetwork** | [**VirtualNetworkCreateInput**](VirtualNetworkCreateInput.md)| Virtual Network to create | 

### Return type

[**VirtualNetwork**](VirtualNetwork.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVirtualNetwork**
> VirtualNetwork DeleteVirtualNetwork($id)

Delete a virtual network

Deletes a virtual network.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Virtual Network UUID | 

### Return type

[**VirtualNetwork**](VirtualNetwork.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisbondPort**
> Port DisbondPort($id, $bulkDisable)

Disabling bonding

Disabling bonding for one or all ports


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Port UUID | 
 **bulkDisable** | **bool**| disable both ports | [optional] 

### Return type

[**Port**](Port.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindVirtualNetworks**
> VirtualNetworkList FindVirtualNetworks($id, $include)

Retrieve all virtual networks

Provides a list of virtual networks for a single project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**VirtualNetworkList**](VirtualNetworkList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnassignPort**
> Port UnassignPort($id, $vnid)

Unassign a port

Unassign a port for a hardware.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Port UUID | 
 **vnid** | **string**| Virtual Network ID | 

### Return type

[**Port**](Port.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

