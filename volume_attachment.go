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

type VolumeAttachment struct {

	Id string `json:"id,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	Volume Href `json:"volume,omitempty"`

	Device Href `json:"device,omitempty"`

	Href string `json:"href,omitempty"`
}
