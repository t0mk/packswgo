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

type ProjectCreateInput struct {
	Name string `json:"name"`

	PaymentMethodId string `json:"payment_method_id,omitempty"`
}
