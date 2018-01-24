/*
 * Aha.io API
 *
 * Articles that matter on social publishing platform
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aha

import (
	"time"
)

type FeatureMeta struct {
	Id string `json:"id,omitempty"`

	ReferenceNum string `json:"reference_num,omitempty"`

	Name string `json:"name,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	Url string `json:"url,omitempty"`

	Resource string `json:"resource,omitempty"`
}
