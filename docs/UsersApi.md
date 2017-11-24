# \UsersApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindCurrentUser**](UsersApi.md#FindCurrentUser) | **Get** /user | Retrieve the current user
[**FindUserById**](UsersApi.md#FindUserById) | **Get** /users/{id} | Retrieve a user
[**FindUsers**](UsersApi.md#FindUsers) | **Get** /users | Retrieve all users
[**UpdateCurrentUser**](UsersApi.md#UpdateCurrentUser) | **Put** /user | Update the current user


# **FindCurrentUser**
> User FindCurrentUser($include)

Retrieve the current user

Returns the user object for the currently logged-in user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**User**](User.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUserById**
> User FindUserById($id, $include)

Retrieve a user

Returns a single user if the user has access


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| User UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**User**](User.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUsers**
> UserList FindUsers($include, $page, $perPage)

Retrieve all users

Returns a list of users that the are accessible to the current user (all users in the current user’s projects, essentially).


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string**| related attributes to include | [optional] 
 **page** | **int32**| page to display, default to 1, max 100_000 | [optional] 
 **perPage** | **int32**| items per page, default to 10, max 1_000 | [optional] 

### Return type

[**UserList**](UserList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCurrentUser**
> User UpdateCurrentUser($user)

Update the current user

Updates the currently logged-in user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**UserUpdateInput**](UserUpdateInput.md)| User to update | 

### Return type

[**User**](User.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

