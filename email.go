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

type Email struct {

	Id string `json:"id,omitempty"`

	Address string `json:"address,omitempty"`

	Default_ bool `json:"default,omitempty"`

	Href string `json:"href,omitempty"`
}
