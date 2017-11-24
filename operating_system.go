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

type OperatingSystem struct {

	Id string `json:"id,omitempty"`

	Slug string `json:"slug,omitempty"`

	Name string `json:"name,omitempty"`

	Distro string `json:"distro,omitempty"`

	Version string `json:"version,omitempty"`

	ProvisionableOn []string `json:"provisionable_on,omitempty"`
}