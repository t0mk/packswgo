# \MembershipsApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMembership**](MembershipsApi.md#DeleteMembership) | **Delete** /memberships/{id} | Delete the membership
[**FindMembershipById**](MembershipsApi.md#FindMembershipById) | **Get** /memberships/{id} | Retrieve a membership
[**UpdateMembership**](MembershipsApi.md#UpdateMembership) | **Put** /memberships/{id} | Update the membership


# **DeleteMembership**
> DeleteMembership($id)

Delete the membership

Deletes the membership.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Membership UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindMembershipById**
> Membership FindMembershipById($id, $include)

Retrieve a membership

Returns a single membership.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Membership UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**Membership**](Membership.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMembership**
> Membership UpdateMembership($id, $membership)

Update the membership

Updates the membership.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Membership UUID | 
 **membership** | [**MembershipInput**](MembershipInput.md)| Membership to update | 

### Return type

[**Membership**](Membership.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

