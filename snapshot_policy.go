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

type SnapshotPolicy struct {

	Id string `json:"id,omitempty"`

	SnapshotCount int32 `json:"snapshot_count,omitempty"`

	SnapshotFrequency string `json:"snapshot_frequency,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Volume Href `json:"volume,omitempty"`

	Href string `json:"href,omitempty"`
}