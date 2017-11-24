# \SSHKeysApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSHKey**](SSHKeysApi.md#CreateSSHKey) | **Post** /ssh-keys | Create a ssh key for the current user
[**CreateSSHKey_0**](SSHKeysApi.md#CreateSSHKey_0) | **Post** /projects/{id}/ssh-keys | Create a ssh key for the given project
[**DeleteSSHKey**](SSHKeysApi.md#DeleteSSHKey) | **Delete** /ssh-keys/{id} | Delete the ssh key
[**FindSSHKeyById**](SSHKeysApi.md#FindSSHKeyById) | **Get** /ssh-keys/{id} | Retrieve a ssh key
[**FindSSHKeys**](SSHKeysApi.md#FindSSHKeys) | **Get** /ssh-keys | Retrieve all ssk keys
[**FindSSHKeys_0**](SSHKeysApi.md#FindSSHKeys_0) | **Get** /projects/{id}/ssh-keys | Retrieve a project&#39;s ssk keys
[**UpdateSSHKey**](SSHKeysApi.md#UpdateSSHKey) | **Put** /ssh-keys/{id} | Update the ssh key


# **CreateSSHKey**
> SshKey CreateSSHKey($sshKey)

Create a ssh key for the current user

Creates a ssh key.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshKey** | [**SshKeyInput**](SshKeyInput.md)| ssh key to create | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSSHKey_0**
> SshKey CreateSSHKey_0($id, $sshKey)

Create a ssh key for the given project

Creates a ssh key.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **sshKey** | [**SshKeyInput**](SshKeyInput.md)| ssh key to create | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSSHKey**
> DeleteSSHKey($id)

Delete the ssh key

Deletes the ssh key.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ssh key UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindSSHKeyById**
> SshKey FindSSHKeyById($id, $include)

Retrieve a ssh key

Returns a single ssh key if the user has access


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| SSH Key UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindSSHKeys**
> SshKeyList FindSSHKeys($include)

Retrieve all ssk keys

Returns a collection of the user’s ssh keys.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**SshKeyList**](SSHKeyList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindSSHKeys_0**
> SshKeyList FindSSHKeys_0($id, $include)

Retrieve a project's ssk keys

Returns a collection of the project's ssh keys.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**SshKeyList**](SSHKeyList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSSHKey**
> SshKey UpdateSSHKey($id, $sshKey)

Update the ssh key

Updates the ssh key.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| SSH Key UUID | 
 **sshKey** | [**SshKeyInput**](SshKeyInput.md)| ssh key to update | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

