# \BatchesApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindBatchById**](BatchesApi.md#FindBatchById) | **Get** /batches/{id} | Retrieve a Batch
[**FindBatchesByProject**](BatchesApi.md#FindBatchesByProject) | **Get** /projects/{id}/batches | Retrieve all batches by project


# **FindBatchById**
> Batch FindBatchById($id, $include)

Retrieve a Batch

Returns a Batch


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Batch UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**Batch**](Batch.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindBatchesByProject**
> BatchesList FindBatchesByProject($id, $include)

Retrieve all batches by project

Returns all batches for the given project


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**BatchesList**](BatchesList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

