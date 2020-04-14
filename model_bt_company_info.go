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
// BtCompanyInfo struct for BtCompanyInfo
type BtCompanyInfo struct {
	Address BtAddressInfo `json:"address,omitempty"`
	Admin bool `json:"admin,omitempty"`
	Description string `json:"description,omitempty"`
	DomainPrefix string `json:"domainPrefix,omitempty"`
	EnterpriseBaseUrl string `json:"enterpriseBaseUrl,omitempty"`
	Href string `json:"href,omitempty"`
	Id string `json:"id,omitempty"`
	Image string `json:"image,omitempty"`
	Name string `json:"name,omitempty"`
	NoPublicDocuments bool `json:"noPublicDocuments,omitempty"`
	OwnerId string `json:"ownerId,omitempty"`
	Purchase BtPurchaseInfo `json:"purchase,omitempty"`
	SecondaryDomainPrefixes []string `json:"secondaryDomainPrefixes,omitempty"`
	State int32 `json:"state,omitempty"`
	Type int32 `json:"type,omitempty"`
	ViewRef string `json:"viewRef,omitempty"`
}
