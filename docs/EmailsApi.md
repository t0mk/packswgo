# \EmailsApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmail**](EmailsApi.md#CreateEmail) | **Post** /emails | Create an email
[**DeleteEmail**](EmailsApi.md#DeleteEmail) | **Delete** /emails/{id} | Delete the email
[**FindEmailById**](EmailsApi.md#FindEmailById) | **Get** /emails/{id} | Retrieve an email
[**UpdateEmail**](EmailsApi.md#UpdateEmail) | **Put** /emails/{id} | Update the email


# **CreateEmail**
> Email CreateEmail($email)

Create an email

Add a new email address to the current user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | [**EmailInput**](EmailInput.md)| Email to create | 

### Return type

[**Email**](Email.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEmail**
> DeleteEmail($id)

Delete the email

Deletes the email.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Email UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindEmailById**
> Email FindEmailById($id, $include)

Retrieve an email

Provides one of the user’s emails.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Email UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**Email**](Email.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEmail**
> Email UpdateEmail($id, $email)

Update the email

Updates the email.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Email UUID | 
 **email** | [**EmailInput**](EmailInput.md)| email to update | 

### Return type

[**Email**](Email.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

