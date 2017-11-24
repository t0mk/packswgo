# \BgpConfigsApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindBgpConfigByProject**](BgpConfigsApi.md#FindBgpConfigByProject) | **Get** /projects/{id}/bgp-config | Retrieve a bgp config
[**RequestBgpConfig**](BgpConfigsApi.md#RequestBgpConfig) | **Post** /projects/{id}/bgp-configs | Requesting bgp config


# **FindBgpConfigByProject**
> BgpConfig FindBgpConfigByProject($id, $include)

Retrieve a bgp config

Returns a bgp config


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**BgpConfig**](BgpConfig.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestBgpConfig**
> RequestBgpConfig($id, $bgpConfigRequest)

Requesting bgp config

Requests to enable bgp configuration for a project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **bgpConfigRequest** | [**BgpConfigRequestInput**](BgpConfigRequestInput.md)| BGP config Request to create | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

