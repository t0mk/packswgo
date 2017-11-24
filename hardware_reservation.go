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

type HardwareReservation struct {

	Id string `json:"id,omitempty"`

	Facility Facility `json:"facility,omitempty"`

	Plan Plan `json:"plan,omitempty"`

	Href string `json:"href,omitempty"`

	Project Project `json:"project,omitempty"`

	Device Device `json:"device,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
}
