/*
 * Set and Tune API
 *
 * API for managing sets and tunes
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apimodel

type Tune struct {

	// Unique identifier for an object
	ID uint64 `json:"id"`

	Title string `json:"title"`

	Type string `json:"type,omitempty"`

	TimeSig string `json:"timeSig,omitempty"`

	Composer string `json:"composer,omitempty"`

	Arranger string `json:"arranger,omitempty"`
}
