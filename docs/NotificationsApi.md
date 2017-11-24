# \NotificationsApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindNotificationById**](NotificationsApi.md#FindNotificationById) | **Get** /notifications/{id} | Retrieve a notification
[**FindNotifications**](NotificationsApi.md#FindNotifications) | **Get** /notifications | Retrieve all notifications
[**UpdateNotification**](NotificationsApi.md#UpdateNotification) | **Put** /notifications/{id} | Update the notification


# **FindNotificationById**
> Notification FindNotificationById($id, $include)

Retrieve a notification

Returns a single notification if the user has access


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Notification UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**Notification**](Notification.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindNotifications**
> NotificationList FindNotifications($all, $since, $include, $page, $perPage)

Retrieve all notifications

Returns a collection of the current user’s notification.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **bool**| Whether to include read notification in the listing (default false) | [optional] 
 **since** | **string**| Only include notifications updated since the given datetime. Must be a UTC datetime in ISO8601 format. (default Time.now) | [optional] 
 **include** | **string**| related attributes to include | [optional] 
 **page** | **int32**| page to display, default to 1, max 100_000 | [optional] 
 **perPage** | **int32**| items per page, default to 10, max 1_000 | [optional] 

### Return type

[**NotificationList**](NotificationList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotification**
> Notification UpdateNotification($id)

Update the notification

Updates a single notification. Currently, the only supported operating is marking a notification as read.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Notification UUID | 

### Return type

[**Notification**](Notification.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

