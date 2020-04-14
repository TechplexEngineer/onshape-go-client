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
// RequestOptions struct for RequestOptions
type RequestOptions struct {
	ApiKey string `json:"apiKey,omitempty"`
	ConnectTimeout int32 `json:"connectTimeout,omitempty"`
	IdempotencyKey string `json:"idempotencyKey,omitempty"`
	ReadTimeout int32 `json:"readTimeout,omitempty"`
	StripeAccount string `json:"stripeAccount,omitempty"`
	StripeVersion string `json:"stripeVersion,omitempty"`
}
