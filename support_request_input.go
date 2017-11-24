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

type SupportRequestInput struct {

	Subject string `json:"subject"`

	Message string `json:"message"`

	ProjectId string `json:"project_id,omitempty"`

	DeviceId string `json:"device_id,omitempty"`
}