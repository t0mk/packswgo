# \VolumeAttachmentsApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolumeAttachment**](VolumeAttachmentsApi.md#CreateVolumeAttachment) | **Post** /storage/{id}/attachments | Attach your volume
[**DeleteVolumeAttachment**](VolumeAttachmentsApi.md#DeleteVolumeAttachment) | **Delete** /storage/attachments/{id} | Detach volume
[**FindVolumeAttachmentById**](VolumeAttachmentsApi.md#FindVolumeAttachmentById) | **Get** /storage/attachments/{id} | Retrieve an attachment
[**FindVolumeAttachments**](VolumeAttachmentsApi.md#FindVolumeAttachments) | **Get** /storage/{id}/attachments | Retrieve all volume attachment


# **CreateVolumeAttachment**
> VolumeAttachment CreateVolumeAttachment($id, $attachment)

Attach your volume

Attach your volume to a device.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Volume UUID | 
 **attachment** | [**VolumeAttachmentInput**](VolumeAttachmentInput.md)| Device to attach | 

### Return type

[**VolumeAttachment**](VolumeAttachment.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVolumeAttachment**
> DeleteVolumeAttachment($id)

Detach volume

Detach volume.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Attachment UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindVolumeAttachmentById**
> VolumeAttachment FindVolumeAttachmentById($id, $include)

Retrieve an attachment

Returns a single attachment if the user has access


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Attachment UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**VolumeAttachment**](VolumeAttachment.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindVolumeAttachments**
> VolumeAttachmentList FindVolumeAttachments($id, $include)

Retrieve all volume attachment

Returns a list of the current volume’s attachments.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Volume UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**VolumeAttachmentList**](VolumeAttachmentList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

