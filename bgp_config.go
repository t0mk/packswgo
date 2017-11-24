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

type BgpConfig struct {

	Id string `json:"id,omitempty"`

	Status string `json:"status,omitempty"`

	DeploymentType string `json:"deployment_type,omitempty"`

	Asn int32 `json:"asn,omitempty"`

	RouteObject string `json:"route_object,omitempty"`

	Md5 string `json:"md5,omitempty"`

	MaxPrefix int32 `json:"max_prefix,omitempty"`

	Project Href `json:"project,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	RequestedAt time.Time `json:"requested_at,omitempty"`

	Session interface{} `json:"session,omitempty"`

	Href string `json:"href,omitempty"`
}
