# \InternetGatewaysApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInternetGateway**](InternetGatewaysApi.md#CreateInternetGateway) | **Post** /virtual-networks/{id}/internet-gateways | Create an internet gateway


# **CreateInternetGateway**
> InternetGateway CreateInternetGateway($id, $length)

Create an internet gateway

Creates an internet gateway.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Virtual Network UUID | 
 **length** | **string**| IP Reservation length | 

### Return type

[**InternetGateway**](InternetGateway.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

