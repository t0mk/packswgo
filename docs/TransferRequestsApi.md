# \TransferRequestsApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptTransferRequest**](TransferRequestsApi.md#AcceptTransferRequest) | **Put** /transfers/{id} | Accept a transfer request
[**CreateTransferRequest**](TransferRequestsApi.md#CreateTransferRequest) | **Post** /projects/{id}/transfers | Create a transfer request
[**DeclineTransferRequest**](TransferRequestsApi.md#DeclineTransferRequest) | **Delete** /transfers/{id} | Decline a transfer request
[**FindTransferRequestById**](TransferRequestsApi.md#FindTransferRequestById) | **Get** /transfers/{id} | View a transfer request


# **AcceptTransferRequest**
> AcceptTransferRequest($id)

Accept a transfer request

Accept a transfer request.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Transfer request UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTransferRequest**
> TransferRequest CreateTransferRequest($id, $transferRequest)

Create a transfer request

Project owners can transfer ownership of their projects to other members.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **transferRequest** | [**TransferRequestInput**](TransferRequestInput.md)| Transfer Request to create | 

### Return type

[**TransferRequest**](TransferRequest.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeclineTransferRequest**
> DeclineTransferRequest($id)

Decline a transfer request

Decline a transfer request.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Transfer request UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTransferRequestById**
> TransferRequest FindTransferRequestById($id, $include)

View a transfer request

Returns a single transfer request.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Transfer request UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**TransferRequest**](TransferRequest.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

