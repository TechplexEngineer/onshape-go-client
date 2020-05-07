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

// BTTableBaseCrossHighlightData2609 struct for BTTableBaseCrossHighlightData2609
type BTTableBaseCrossHighlightData2609 struct {
	BtType *string `json:"btType,omitempty"`
}

// NewBTTableBaseCrossHighlightData2609 instantiates a new BTTableBaseCrossHighlightData2609 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableBaseCrossHighlightData2609() *BTTableBaseCrossHighlightData2609 {
	this := BTTableBaseCrossHighlightData2609{}
	return &this
}

// NewBTTableBaseCrossHighlightData2609WithDefaults instantiates a new BTTableBaseCrossHighlightData2609 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableBaseCrossHighlightData2609WithDefaults() *BTTableBaseCrossHighlightData2609 {
	this := BTTableBaseCrossHighlightData2609{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableBaseCrossHighlightData2609) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableBaseCrossHighlightData2609) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableBaseCrossHighlightData2609) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableBaseCrossHighlightData2609) SetBtType(v string) {
	o.BtType = &v
}

func (o BTTableBaseCrossHighlightData2609) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableBaseCrossHighlightData2609 struct {
	value *BTTableBaseCrossHighlightData2609
	isSet bool
}

func (v NullableBTTableBaseCrossHighlightData2609) Get() *BTTableBaseCrossHighlightData2609 {
	return v.value
}

func (v *NullableBTTableBaseCrossHighlightData2609) Set(val *BTTableBaseCrossHighlightData2609) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableBaseCrossHighlightData2609) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableBaseCrossHighlightData2609) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableBaseCrossHighlightData2609(val *BTTableBaseCrossHighlightData2609) *NullableBTTableBaseCrossHighlightData2609 {
	return &NullableBTTableBaseCrossHighlightData2609{value: val, isSet: true}
}

func (v NullableBTTableBaseCrossHighlightData2609) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableBaseCrossHighlightData2609) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
