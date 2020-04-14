/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.111
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// BtMassPropertiesBulkInfo struct for BtMassPropertiesBulkInfo
type BtMassPropertiesBulkInfo struct {
	Bodies map[string]BtMassPropertiesInfoNull `json:"bodies,omitempty"`
	MicroversionId string `json:"microversionId,omitempty"`
}
