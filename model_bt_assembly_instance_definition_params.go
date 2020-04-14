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
// BtAssemblyInstanceDefinitionParams struct for BtAssemblyInstanceDefinitionParams
type BtAssemblyInstanceDefinitionParams struct {
	Configuration string `json:"configuration,omitempty"`
	DocumentId string `json:"documentId,omitempty"`
	ElementId string `json:"elementId,omitempty"`
	FeatureId string `json:"featureId,omitempty"`
	IsAssembly bool `json:"isAssembly,omitempty"`
	IsHidden bool `json:"isHidden,omitempty"`
	IsSuppressed bool `json:"isSuppressed,omitempty"`
	IsWholePartStudio bool `json:"isWholePartStudio,omitempty"`
	MicroversionId string `json:"microversionId,omitempty"`
	PartId string `json:"partId,omitempty"`
	PartNumber string `json:"partNumber,omitempty"`
	Revision string `json:"revision,omitempty"`
	VersionId string `json:"versionId,omitempty"`
}
