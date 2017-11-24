# \ProjectsApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProject**](ProjectsApi.md#CreateProject) | **Post** /projects | Create a project
[**DeleteProject**](ProjectsApi.md#DeleteProject) | **Delete** /projects/{id} | Delete the project
[**FindProjectById**](ProjectsApi.md#FindProjectById) | **Get** /projects/{id} | Retrieve a project
[**FindProjects**](ProjectsApi.md#FindProjects) | **Get** /projects | Retrieve all projects
[**UpdateProject**](ProjectsApi.md#UpdateProject) | **Put** /projects/{id} | Update the project


# **CreateProject**
> Project CreateProject($project)

Create a project

Creates a new projects and adds the current user as a member.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**ProjectCreateInput**](ProjectCreateInput.md)| Project to create | 

### Return type

[**Project**](Project.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> DeleteProject($id)

Delete the project

Deletes the project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindProjectById**
> Project FindProjectById($id, $include)

Retrieve a project

Returns a single project if the user has access


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**Project**](Project.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindProjects**
> ProjectList FindProjects($include, $page, $perPage)

Retrieve all projects

Returns a collection of projects that the current user is a member of.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string**| related attributes to include | [optional] 
 **page** | **int32**| page to display, default to 1, max 100_000 | [optional] 
 **perPage** | **int32**| items per page, default to 10, max 1_000 | [optional] 

### Return type

[**ProjectList**](ProjectList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProject**
> Project UpdateProject($id, $project)

Update the project

Updates the project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **project** | [**ProjectUpdateInput**](ProjectUpdateInput.md)| Project to update | 

### Return type

[**Project**](Project.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

