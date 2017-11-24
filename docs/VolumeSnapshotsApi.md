# \VolumeSnapshotsApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolumeSnapshot**](VolumeSnapshotsApi.md#CreateVolumeSnapshot) | **Post** /storage/{id}/snapshots | Create a volume snapshot
[**DeleteVolumeSnapshot**](VolumeSnapshotsApi.md#DeleteVolumeSnapshot) | **Delete** /storage/{volume_id}/snapshots/{id} | Delete volume snapshot
[**FindVolumeSnapshots**](VolumeSnapshotsApi.md#FindVolumeSnapshots) | **Get** /storage/{id}/snapshots | Retrieve all volume snapshot


# **CreateVolumeSnapshot**
> CreateVolumeSnapshot($id)

Create a volume snapshot

Creates a new snapshot of your volume.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Volume UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVolumeSnapshot**
> DeleteVolumeSnapshot($volumeId, $id)

Delete volume snapshot

Delete volume snapshot.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeId** | **string**| Volume UUID | 
 **id** | **string**| Snapshot UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindVolumeSnapshots**
> VolumeSnapshotList FindVolumeSnapshots($id, $include)

Retrieve all volume snapshot

Returns a list of the current volume’s snapshots.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Volume UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**VolumeSnapshotList**](VolumeSnapshotList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

