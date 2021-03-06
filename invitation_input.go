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

type InvitationInput struct {

	Invitee string `json:"invitee"`

	Message string `json:"message,omitempty"`

	Roles []string `json:"roles,omitempty"`
}
