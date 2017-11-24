# \BgpSessionsApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBgpSession**](BgpSessionsApi.md#CreateBgpSession) | **Post** /devices/{id}/bgp/sessions | Create a BGP session
[**DeleteBgpSession**](BgpSessionsApi.md#DeleteBgpSession) | **Delete** /bgp/sessions/{id} | Delete the BGP session
[**FindBgpSessionById**](BgpSessionsApi.md#FindBgpSessionById) | **Get** /bgp/sessions/{id} | Retrieve a BGP session
[**FindBgpSessions**](BgpSessionsApi.md#FindBgpSessions) | **Get** /devices/{id}/bgp/sessions | Retrieve all BGP sessions


# **CreateBgpSession**
> BgpSession CreateBgpSession($id, $addressFamily)

Create a BGP session

Creates a BGP session.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Device UUID | 
 **addressFamily** | **string**| address family for BGP session | 

### Return type

[**BgpSession**](BgpSession.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBgpSession**
> DeleteBgpSession($id)

Delete the BGP session

Deletes the BGP session.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| BGP session UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindBgpSessionById**
> BgpSession FindBgpSessionById($id, $include)

Retrieve a BGP session

Returns a BGP session


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| BGP session UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**BgpSession**](BgpSession.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindBgpSessions**
> BgpSessionList FindBgpSessions($id)

Retrieve all BGP sessions

Provides a listing of available BGP sessions for the device.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Device UUID | 

### Return type

[**BgpSessionList**](BgpSessionList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

