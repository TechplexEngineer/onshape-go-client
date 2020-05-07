/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.113
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package onshape

import (
	"encoding/json"
)

// BTExternalReference1936AllOf struct for BTExternalReference1936AllOf
type BTExternalReference1936AllOf struct {
	BtType *string `json:"btType,omitempty"`
	DocumentVersionId *string `json:"documentVersionId,omitempty"`
}

// NewBTExternalReference1936AllOf instantiates a new BTExternalReference1936AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExternalReference1936AllOf() *BTExternalReference1936AllOf {
	this := BTExternalReference1936AllOf{}
	return &this
}

// NewBTExternalReference1936AllOfWithDefaults instantiates a new BTExternalReference1936AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExternalReference1936AllOfWithDefaults() *BTExternalReference1936AllOf {
	this := BTExternalReference1936AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExternalReference1936AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalReference1936AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExternalReference1936AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExternalReference1936AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentVersionId returns the DocumentVersionId field value if set, zero value otherwise.
func (o *BTExternalReference1936AllOf) GetDocumentVersionId() string {
	if o == nil || o.DocumentVersionId == nil {
		var ret string
		return ret
	}
	return *o.DocumentVersionId
}

// GetDocumentVersionIdOk returns a tuple with the DocumentVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExternalReference1936AllOf) GetDocumentVersionIdOk() (*string, bool) {
	if o == nil || o.DocumentVersionId == nil {
		return nil, false
	}
	return o.DocumentVersionId, true
}

// HasDocumentVersionId returns a boolean if a field has been set.
func (o *BTExternalReference1936AllOf) HasDocumentVersionId() bool {
	if o != nil && o.DocumentVersionId != nil {
		return true
	}

	return false
}

// SetDocumentVersionId gets a reference to the given string and assigns it to the DocumentVersionId field.
func (o *BTExternalReference1936AllOf) SetDocumentVersionId(v string) {
	o.DocumentVersionId = &v
}

func (o BTExternalReference1936AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DocumentVersionId != nil {
		toSerialize["documentVersionId"] = o.DocumentVersionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTExternalReference1936AllOf struct {
	value *BTExternalReference1936AllOf
	isSet bool
}

func (v NullableBTExternalReference1936AllOf) Get() *BTExternalReference1936AllOf {
	return v.value
}

func (v *NullableBTExternalReference1936AllOf) Set(val *BTExternalReference1936AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExternalReference1936AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExternalReference1936AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExternalReference1936AllOf(val *BTExternalReference1936AllOf) *NullableBTExternalReference1936AllOf {
	return &NullableBTExternalReference1936AllOf{value: val, isSet: true}
}

func (v NullableBTExternalReference1936AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExternalReference1936AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
