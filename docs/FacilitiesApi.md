# \FacilitiesApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindFacilities**](FacilitiesApi.md#FindFacilities) | **Get** /facilities | Retrieve all facilities
[**FindFacilitiesByProject**](FacilitiesApi.md#FindFacilitiesByProject) | **Get** /projects/{id}/facilities | Retrieve all facilities visible by the project


# **FindFacilities**
> FacilityList FindFacilities()

Retrieve all facilities

Provides a listing of available datacenters where you can provision Packet devices.


### Parameters
This endpoint does not need any parameter.

### Return type

[**FacilityList**](FacilityList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindFacilitiesByProject**
> FacilityList FindFacilitiesByProject($id, $include, $page, $perPage)

Retrieve all facilities visible by the project

Returns a listing of available datacenters for the given project


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **include** | **string**| related attributes to include | [optional] 
 **page** | **int32**| page to display, default to 1, max 100_000 | [optional] 
 **perPage** | **int32**| items per page, default to 10, max 1_000 | [optional] 

### Return type

[**FacilityList**](FacilityList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

