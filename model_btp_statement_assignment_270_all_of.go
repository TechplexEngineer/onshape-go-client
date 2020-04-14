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
// BtpStatementAssignment270AllOf struct for BtpStatementAssignment270AllOf
type BtpStatementAssignment270AllOf struct {
	Operator string `json:"operator,omitempty"`
	Lvalue BtplValue249 `json:"lvalue,omitempty"`
	Rvalue BtpExpression9 `json:"rvalue,omitempty"`
	BtType string `json:"btType,omitempty"`
}
