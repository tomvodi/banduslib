/*
 * Set and Tune API
 *
 * API for managing sets and tunes
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apimodel

type ImportFile struct {
	Name string `json:"name"`

	Result ParseResult `json:"result,omitempty"`

	// if import was successful, the array of imported tunes
	Tunes []*ImportTune `json:"tunes,omitempty"`
}
