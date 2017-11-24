# \InstancesBatchApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceBatch**](InstancesBatchApi.md#CreateDeviceBatch) | **Post** /projects/{id}/devices/batch | Create a devices batch


# **CreateDeviceBatch**
> Batch CreateDeviceBatch($id, $batch)

Create a devices batch

Creates a new batch of instances.  Type-specific options (such as operating_system for baremetal devices) should be included in the main data structure alongside hostname and plan.  The features attribute allows you to optionally specify what features your server should have. For example, if you require a server with a TPM chip, you may specify { \"features\": { \"tpm\": \"required\" } } (or { \"features\": [\"tpm\"] } in shorthand). The request will fail if there are no available servers matching your criteria. Alternatively, if you do not require a certain feature, but would prefer to be assigned a server with that feature if there are any available, you may specify that feature with a preferred value (see the example request below). The request will not fail if we have no servers with that feature in our inventory.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **batch** | [**InstancesBatchCreateInput**](InstancesBatchCreateInput.md)| Batches to create | 

### Return type

[**Batch**](Batch.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

