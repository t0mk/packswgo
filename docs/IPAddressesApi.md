# \IPAddressesApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIPAssignment**](IPAddressesApi.md#CreateIPAssignment) | **Post** /devices/{id}/ips | Create a ip assignment
[**DeleteIPAddress**](IPAddressesApi.md#DeleteIPAddress) | **Delete** /ips/{id} | Unassign an ip address
[**FindIPAddressById**](IPAddressesApi.md#FindIPAddressById) | **Get** /ips/{id} | Retrieve an ip address
[**FindIPAssignments**](IPAddressesApi.md#FindIPAssignments) | **Get** /devices/{id}/ips | Retrieve all ip assignments
[**FindIPAvailabilities**](IPAddressesApi.md#FindIPAvailabilities) | **Get** /ips/{id}/available | Retrieve all available subnets of a particular reservation
[**FindIPReservations**](IPAddressesApi.md#FindIPReservations) | **Get** /projects/{id}/ips | Retrieve all ip reservations
[**RequestIPReservation**](IPAddressesApi.md#RequestIPReservation) | **Post** /projects/{id}/ips | Requesting ip reservations


# **CreateIPAssignment**
> IpAssignment CreateIPAssignment($id, $ipAssignment)

Create a ip assignment

Creates an ip assignment for a device.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Device UUID | 
 **ipAssignment** | [**IpAssignmentInput**](IpAssignmentInput.md)| IPAssignment to create | 

### Return type

[**IpAssignment**](IPAssignment.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPAddress**
> DeleteIPAddress($id)

Unassign an ip address

Unassign an IP address record. This will remove the relationship between an IP and the device and will make the IP address available to be assigned to another device.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| IP Address UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindIPAddressById**
> IpAssignment FindIPAddressById($id, $include)

Retrieve an ip address

Returns a single ip address if the user has access.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| IP Address UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**IpAssignment**](IPAssignment.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindIPAssignments**
> IpAssignmentList FindIPAssignments($id, $include)

Retrieve all ip assignments

Returns all ip assignments for a device.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Device UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**IpAssignmentList**](IPAssignmentList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindIPAvailabilities**
> IpAvailabilitiesList FindIPAvailabilities($id, $cidr)

Retrieve all available subnets of a particular reservation

Provides a list of IP resevations for a single project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| IP Reservation UUID | 
 **cidr** | **string**| Size of subnets in bits | 

### Return type

[**IpAvailabilitiesList**](IPAvailabilitiesList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindIPReservations**
> IpReservationList FindIPReservations($id, $include)

Retrieve all ip reservations

Provides a list of IP resevations for a single project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**IpReservationList**](IPReservationList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestIPReservation**
> IpReservationRequest RequestIPReservation($id, $ipReservationRequest)

Requesting ip reservations

Request more IP space for a project in order to have additional IP addresses to assign to devices.  If the request is within the max quota, ip reservation will be created, else, the request will be submitted.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **ipReservationRequest** | [**IpReservationRequestInput**](IpReservationRequestInput.md)| IP Reservation Request to create | 

### Return type

[**IpReservationRequest**](IPReservationRequest.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

