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

type Notification struct {

	Id string `json:"id,omitempty"`

	Type_ string `json:"type,omitempty"`

	Body string `json:"body,omitempty"`

	Severity string `json:"severity,omitempty"`

	Read bool `json:"read,omitempty"`

	Context string `json:"context,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	User Href `json:"user,omitempty"`

	Href string `json:"href,omitempty"`
}