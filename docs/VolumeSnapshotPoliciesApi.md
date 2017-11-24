# \VolumeSnapshotPoliciesApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolumeSnapshotPolicy**](VolumeSnapshotPoliciesApi.md#CreateVolumeSnapshotPolicy) | **Post** /storage/{id}/snapshot-policies | Create a volume snapshot policy
[**DeleteVolumeSnapshotPolicy**](VolumeSnapshotPoliciesApi.md#DeleteVolumeSnapshotPolicy) | **Delete** /storage/snapshot-policies/{id} | Delete the volume snapshot policy
[**UpdateVolumeSnapshotPolicy**](VolumeSnapshotPoliciesApi.md#UpdateVolumeSnapshotPolicy) | **Put** /storage/snapshot-policies/{id} | Update the volume snapshot policy


# **CreateVolumeSnapshotPolicy**
> SnapshotPolicy CreateVolumeSnapshotPolicy($id, $snapshotFrequency, $snapshotCount)

Create a volume snapshot policy

Creates a new snapshot policy of your volume.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Volume UUID | 
 **snapshotFrequency** | **string**| Snapshot frequency | 
 **snapshotCount** | **int32**| Snapshot count | [optional] 

### Return type

[**SnapshotPolicy**](SnapshotPolicy.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVolumeSnapshotPolicy**
> DeleteVolumeSnapshotPolicy($id)

Delete the volume snapshot policy

Deletes the volume snapshot policy.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Snapshot Policy UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVolumeSnapshotPolicy**
> SnapshotPolicy UpdateVolumeSnapshotPolicy($id, $snapshotFrequency, $snapshotCount)

Update the volume snapshot policy

Updates the volume snapshot policy.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Snapshot Policy UUID | 
 **snapshotFrequency** | **string**| Snapshot frequency | 
 **snapshotCount** | **int32**| Snapshot count | [optional] 

### Return type

[**SnapshotPolicy**](SnapshotPolicy.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

