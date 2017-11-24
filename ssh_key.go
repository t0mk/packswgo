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
	"time"
)

type SshKey struct {

	Id string `json:"id,omitempty"`

	Label string `json:"label,omitempty"`

	Key string `json:"key,omitempty"`

	Fingerprint string `json:"fingerprint,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Entity Href `json:"entity,omitempty"`

	Href string `json:"href,omitempty"`
}