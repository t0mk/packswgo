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

type VolumeUpdateInput struct {

	Description string `json:"description,omitempty"`

	Size int32 `json:"size,omitempty"`

	Locked bool `json:"locked,omitempty"`

	// hourly, daily, monthly, yearly
	BillingCycle string `json:"billing_cycle,omitempty"`
}