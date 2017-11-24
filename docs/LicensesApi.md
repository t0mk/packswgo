# \LicensesApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLicense**](LicensesApi.md#CreateLicense) | **Post** /projects/{id}/licenses | Create a License
[**DeleteLicense**](LicensesApi.md#DeleteLicense) | **Delete** /licenses/{id} | Delete the license
[**FindLicenseById**](LicensesApi.md#FindLicenseById) | **Get** /licenses/{id} | Retrieve a license
[**FindLicenses**](LicensesApi.md#FindLicenses) | **Get** /projects/{id}/licenses | Retrieve all licenses
[**UpdateLicense**](LicensesApi.md#UpdateLicense) | **Put** /licenses/{id} | Update the license


# **CreateLicense**
> License CreateLicense($id, $license)

Create a License

Creates a new license for the given project


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **license** | [**LicenseCreateInput**](LicenseCreateInput.md)| License to create | 

### Return type

[**License**](License.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLicense**
> DeleteLicense($id)

Delete the license

Deletes a license.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| License UUID | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindLicenseById**
> License FindLicenseById($id, $include)

Retrieve a license

Returns a license


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| License UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**License**](License.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindLicenses**
> LicenseList FindLicenses($id, $include)

Retrieve all licenses

Provides a collection of licenses for a given project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Project UUID | 
 **include** | **string**| related attributes to include | [optional] 

### Return type

[**LicenseList**](LicenseList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLicense**
> License UpdateLicense($id, $license)

Update the license

Updates the license.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| License UUID | 
 **license** | [**LicenseUpdateInput**](LicenseUpdateInput.md)| License to update | 

### Return type

[**License**](License.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

