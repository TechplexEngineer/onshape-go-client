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
// BtpExpressionIs242 struct for BtpExpressionIs242
type BtpExpressionIs242 struct {
	Atomic bool `json:"atomic,omitempty"`
	BtType string `json:"btType,omitempty"`
	DocumentationType string `json:"documentationType,omitempty"`
	EndSourceLocation int32 `json:"endSourceLocation,omitempty"`
	NodeId string `json:"nodeId,omitempty"`
	ShortDescriptor string `json:"shortDescriptor,omitempty"`
	SpaceAfter BtpSpace10 `json:"spaceAfter,omitempty"`
	SpaceBefore BtpSpace10 `json:"spaceBefore,omitempty"`
	SpaceDefault bool `json:"spaceDefault,omitempty"`
	StartSourceLocation int32 `json:"startSourceLocation,omitempty"`
	TypeName BtpTypeName290 `json:"typeName,omitempty"`
	Operand BtpExpression9 `json:"operand,omitempty"`
}
