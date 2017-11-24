# \InvitationsApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptInvitation**](InvitationsApi.md#AcceptInvitation) | **Put** /invitations/{id} | Accept an invitation
[**CreateInvitation**](InvitationsApi.md#CreateInvitation) | **Post** /projects/{id}/invitations | Create an invitation
[**DeclineInvitation**](InvitationsApi.md#DeclineInvitation) | **Delete** /invitations/{id} | Decline an invitation
[**FindInvitationById**](InvitationsApi.md#FindInvitationById) | **Get** /invitations/{id} | View an invitation


# **AcceptInvitation**
> Membership AcceptInvitation($id)

Accept an invitation

Accept an invitation.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Invitation UUID | 

### Return type

[**Membership**](Membership.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInvitation**
> Invitation CreateInvitation($id, $invitation)

Create an invitation

In order to add a user to a project, they must first be invited.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **invitation** | [**InvitationInput**](InvitationInput.md)| Invitation to create | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeclineInvitation**
> DeclineInvitation($id)

Decline an invitation

Decline an invitation.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Invitation UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindInvitationById**
> Invitation FindInvitationById($id, $include)

View an invitation

Returns a single invitation.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Invitation UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**Invitation**](Invitation.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

