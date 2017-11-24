# \VPNApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindCurrentUserVpnConfig**](VPNApi.md#FindCurrentUserVpnConfig) | **Get** /user/vpn | Retrieve the client vpn config for current user
[**TurnOffCurrentUserVpn**](VPNApi.md#TurnOffCurrentUserVpn) | **Delete** /user/vpn | Turn off vpn for the current user
[**TurnOnCurrentUserVpn**](VPNApi.md#TurnOnCurrentUserVpn) | **Post** /user/vpn | Turn on vpn for the current user


# **FindCurrentUserVpnConfig**
> VpnConfig FindCurrentUserVpnConfig($code)

Retrieve the client vpn config for current user

Returns the client vpn config for the currently logged-in user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string**| Facility code | 

### Return type

[**VpnConfig**](VPNConfig.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TurnOffCurrentUserVpn**
> TurnOffCurrentUserVpn()

Turn off vpn for the current user

Turns off vpn for the currently logged-in user.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TurnOnCurrentUserVpn**
> TurnOnCurrentUserVpn()

Turn on vpn for the current user

Turns on vpn for the currently logged-in user.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

