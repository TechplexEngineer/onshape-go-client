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
// BtmFeatureQueryWithOccurrence157 struct for BtmFeatureQueryWithOccurrence157
type BtmFeatureQueryWithOccurrence157 struct {
	BtType string `json:"btType,omitempty"`
	DeterministicIdList BtmIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds []string `json:"deterministicIds,omitempty"`
	ImportMicroversion string `json:"importMicroversion,omitempty"`
	NodeId string `json:"nodeId,omitempty"`
	Path []string `json:"path,omitempty"`
	Query BtmIndividualQueryBase139 `json:"query,omitempty"`
	QueryString string `json:"queryString,omitempty"`
	FeatureIdWithOccurrence string `json:"featureIdWithOccurrence,omitempty"`
	PartStudioMateConnectorQuery bool `json:"partStudioMateConnectorQuery,omitempty"`
	FeatureId string `json:"featureId,omitempty"`
	QueryData string `json:"queryData,omitempty"`
}
