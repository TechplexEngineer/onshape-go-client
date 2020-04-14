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
// BtmIndividualCreatedByQuery137 struct for BtmIndividualCreatedByQuery137
type BtmIndividualCreatedByQuery137 struct {
	BtType string `json:"btType,omitempty"`
	DeterministicIdList BtmIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds []string `json:"deterministicIds,omitempty"`
	ImportMicroversion string `json:"importMicroversion,omitempty"`
	NodeId string `json:"nodeId,omitempty"`
	Query BtmIndividualQueryBase139 `json:"query,omitempty"`
	QueryString string `json:"queryString,omitempty"`
	PersistentQuery BtpStatement269 `json:"persistentQuery,omitempty"`
	VariableName BtmIndividualQuery138 `json:"variableName,omitempty"`
	QueryStatement BtpStatement269 `json:"queryStatement,omitempty"`
	FeatureId string `json:"featureId,omitempty"`
	BodyType string `json:"bodyType,omitempty"`
	EntityType string `json:"entityType,omitempty"`
	FilterConstruction bool `json:"filterConstruction,omitempty"`
}
