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

type Plan struct {

	Id string `json:"id,omitempty"`

	Slug string `json:"slug,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Line string `json:"line,omitempty"`

	Specs interface{} `json:"specs,omitempty"`

	Pricing interface{} `json:"pricing,omitempty"`

	AvailableIn []Href `json:"available_in,omitempty"`
}