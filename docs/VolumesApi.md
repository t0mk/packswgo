# \VolumesApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneVolume**](VolumesApi.md#CloneVolume) | **Post** /storage/{id}/clone | Clone volume/snapshot
[**CreateVolume**](VolumesApi.md#CreateVolume) | **Post** /projects/{id}/storage | Create a volume
[**DeleteVolume**](VolumesApi.md#DeleteVolume) | **Delete** /storage/{id} | Delete the volume
[**FindVolumeById**](VolumesApi.md#FindVolumeById) | **Get** /storage/{id} | Retrieve a volume
[**FindVolumes**](VolumesApi.md#FindVolumes) | **Get** /projects/{id}/storage | Retrieve all volumes
[**RestoreVolume**](VolumesApi.md#RestoreVolume) | **Post** /storage/{id}/restore | Restore volume
[**UpdateVolume**](VolumesApi.md#UpdateVolume) | **Put** /storage/{id} | Update the volume


# **CloneVolume**
> Volume CloneVolume($id, $snapshotTimestamp)

Clone volume/snapshot

Clone your volume or snapshot into a new volume. To clone the volume, send an empty body. To promote a volume snapshot into a new volume, include the snapshot_timestamp attribute in the request body.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Volume UUID | 
 **snapshotTimestamp** | **string**| snapshot timestamp | [optional] 

### Return type

[**Volume**](Volume.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVolume**
> Volume CreateVolume($id, $volume)

Create a volume

Creates a new volume in our datacenter.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **volume** | [**VolumeCreateInput**](VolumeCreateInput.md)| Volume to create | 

### Return type

[**Volume**](Volume.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVolume**
> DeleteVolume($id)

Delete the volume

Deletes the volume.


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

# **FindVolumeById**
> Volume FindVolumeById($id, $include)

Retrieve a volume

Returns a single volume if the user has access


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Volume UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**Volume**](Volume.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindVolumes**
> VolumeList FindVolumes($id, $include, $page, $perPage)

Retrieve all volumes

Returns a list of the current projects’s volumes.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **include** | **string**| related attributes to include | [optional] 
 **page** | **int32**| page to display, default to 1, max 100_000 | [optional] 
 **perPage** | **int32**| items per page, default to 10, max 1_000 | [optional] 

### Return type

[**VolumeList**](VolumeList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreVolume**
> Volume RestoreVolume($id, $restorePoint)

Restore volume

Restore the volume to the given snapshot.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Volume UUID | 
 **restorePoint** | **string**| restore point | 

### Return type

[**Volume**](Volume.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVolume**
> Volume UpdateVolume($id, $volume)

Update the volume

Updates the volume.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Volume UUID | 
 **volume** | [**VolumeUpdateInput**](VolumeUpdateInput.md)| Volume to update | 

### Return type

[**Volume**](Volume.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

