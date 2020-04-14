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
// BtAppElementBasicInfo struct for BtAppElementBasicInfo
type BtAppElementBasicInfo struct {
	ChangeId string `json:"changeId,omitempty"`
	ErrorCode int32 `json:"errorCode,omitempty"`
	ErrorDescription string `json:"errorDescription,omitempty"`
	ErrorValue string `json:"errorValue,omitempty"`
	ParentChangeId string `json:"parentChangeId,omitempty"`
}
