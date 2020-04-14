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
// BtEllipseDescription866 struct for BtEllipseDescription866
type BtEllipseDescription866 struct {
	BtType string `json:"btType,omitempty"`
	Type string `json:"type,omitempty"`
	Normal BtVector3d389 `json:"normal,omitempty"`
	MinorRadius float64 `json:"minorRadius,omitempty"`
	MajorAxis BtVector3d389 `json:"majorAxis,omitempty"`
	MajorRadius float64 `json:"majorRadius,omitempty"`
	Origin BtVector3d389 `json:"origin,omitempty"`
}
