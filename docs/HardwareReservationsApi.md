# \HardwareReservationsApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindHardwareReservationById**](HardwareReservationsApi.md#FindHardwareReservationById) | **Get** /hardware-reservations/{id} | Retrieve a hardware reservation
[**FindHardwareReservations**](HardwareReservationsApi.md#FindHardwareReservations) | **Get** /projects/{id}/hardware-reservations | Retrieve all hardware reservations


# **FindHardwareReservationById**
> Device FindHardwareReservationById($id, $include)

Retrieve a hardware reservation

Returns a single hardware reservation


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| HardwareReservation UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**Device**](Device.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindHardwareReservations**
> HardwareReservationList FindHardwareReservations($id, $include, $page, $perPage)

Retrieve all hardware reservations

Provides a collection of hardware reservations for a given project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **include** | **string**| related attributes to include | [optional] 
 **page** | **int32**| page to display, default to 1, max 100_000 | [optional] 
 **perPage** | **int32**| items per page, default to 10, max 1_000 | [optional] 

### Return type

[**HardwareReservationList**](HardwareReservationList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

