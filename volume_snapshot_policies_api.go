/* 
 * Packet API
 *
 * This is the API for Packet. Interact with your devices, user account, and projects.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: help@packet.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package packswgo

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type VolumeSnapshotPoliciesApi struct {
	Configuration *Configuration
}

func NewVolumeSnapshotPoliciesApi() *VolumeSnapshotPoliciesApi {
	configuration := NewConfiguration()
	return &VolumeSnapshotPoliciesApi{
		Configuration: configuration,
	}
}

func NewVolumeSnapshotPoliciesApiWithBasePath(basePath string) *VolumeSnapshotPoliciesApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &VolumeSnapshotPoliciesApi{
		Configuration: configuration,
	}
}

/**
 * Create a volume snapshot policy
 * Creates a new snapshot policy of your volume.
 *
 * @param id Volume UUID
 * @param snapshotFrequency Snapshot frequency
 * @param snapshotCount Snapshot count
 * @return *SnapshotPolicy
 */
func (a VolumeSnapshotPoliciesApi) CreateVolumeSnapshotPolicy(id string, snapshotFrequency string, snapshotCount int32) (*SnapshotPolicy, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/storage/{id}/snapshot-policies"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(x_auth_token)' required
	// set key with prefix in header
	localVarHeaderParams["X-Auth-Token"] = a.Configuration.GetAPIKeyWithPrefix("X-Auth-Token")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("snapshot_count", a.Configuration.APIClient.ParameterToString(snapshotCount, ""))
	localVarQueryParams.Add("snapshot_frequency", a.Configuration.APIClient.ParameterToString(snapshotFrequency, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SnapshotPolicy)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateVolumeSnapshotPolicy", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Delete the volume snapshot policy
 * Deletes the volume snapshot policy.
 *
 * @param id Snapshot Policy UUID
 * @return void
 */
func (a VolumeSnapshotPoliciesApi) DeleteVolumeSnapshotPolicy(id string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/storage/snapshot-policies/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(x_auth_token)' required
	// set key with prefix in header
	localVarHeaderParams["X-Auth-Token"] = a.Configuration.GetAPIKeyWithPrefix("X-Auth-Token")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteVolumeSnapshotPolicy", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

/**
 * Update the volume snapshot policy
 * Updates the volume snapshot policy.
 *
 * @param id Snapshot Policy UUID
 * @param snapshotFrequency Snapshot frequency
 * @param snapshotCount Snapshot count
 * @return *SnapshotPolicy
 */
func (a VolumeSnapshotPoliciesApi) UpdateVolumeSnapshotPolicy(id string, snapshotFrequency string, snapshotCount int32) (*SnapshotPolicy, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/storage/snapshot-policies/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(x_auth_token)' required
	// set key with prefix in header
	localVarHeaderParams["X-Auth-Token"] = a.Configuration.GetAPIKeyWithPrefix("X-Auth-Token")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("snapshot_count", a.Configuration.APIClient.ParameterToString(snapshotCount, ""))
	localVarQueryParams.Add("snapshot_frequency", a.Configuration.APIClient.ParameterToString(snapshotFrequency, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SnapshotPolicy)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "UpdateVolumeSnapshotPolicy", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

