# \SpotMarketPricesApi

All URIs are relative to *https://api.packet.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindSpotMarketPrices**](SpotMarketPricesApi.md#FindSpotMarketPrices) | **Get** /market/spot/prices | Get current spot market prices
[**FindSpotMarketPricesHistory**](SpotMarketPricesApi.md#FindSpotMarketPricesHistory) | **Get** /market/spot/prices/history | Get spot market prices for a given period of time


# **FindSpotMarketPrices**
> FindSpotMarketPrices($facility, $plan)

Get current spot market prices

Get Packet current spot market prices.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **facility** | **string**| Facility to check spot market prices | [optional] 
 **plan** | **string**| Plan to check spot market prices | [optional] 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindSpotMarketPricesHistory**
> FindSpotMarketPricesHistory($facility, $plan)

Get spot market prices for a given period of time

Get spot market prices for a given plan and facility in a fixed period of time


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **facility** | **string**| Facility to check spot market prices | 
 **plan** | **string**| Plan to check spot market prices | 

### Return type

void (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

