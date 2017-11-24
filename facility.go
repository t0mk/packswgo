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

type Facility struct {

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Code string `json:"code,omitempty"`

	Features []string `json:"features,omitempty"`

	Address Address `json:"address,omitempty"`
}
