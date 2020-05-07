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

// BTPLiteralUndefined260 struct for BTPLiteralUndefined260
type BTPLiteralUndefined260 struct {
	BTPLiteral253
	BtType *string `json:"btType,omitempty"`
}

// NewBTPLiteralUndefined260 instantiates a new BTPLiteralUndefined260 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPLiteralUndefined260() *BTPLiteralUndefined260 {
	this := BTPLiteralUndefined260{}
	return &this
}

// NewBTPLiteralUndefined260WithDefaults instantiates a new BTPLiteralUndefined260 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPLiteralUndefined260WithDefaults() *BTPLiteralUndefined260 {
	this := BTPLiteralUndefined260{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPLiteralUndefined260) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralUndefined260) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPLiteralUndefined260) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPLiteralUndefined260) SetBtType(v string) {
	o.BtType = &v
}

func (o BTPLiteralUndefined260) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPLiteral253, errBTPLiteral253 := json.Marshal(o.BTPLiteral253)
	if errBTPLiteral253 != nil {
		return []byte{}, errBTPLiteral253
	}
	errBTPLiteral253 = json.Unmarshal([]byte(serializedBTPLiteral253), &toSerialize)
	if errBTPLiteral253 != nil {
		return []byte{}, errBTPLiteral253
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTPLiteralUndefined260 struct {
	value *BTPLiteralUndefined260
	isSet bool
}

func (v NullableBTPLiteralUndefined260) Get() *BTPLiteralUndefined260 {
	return v.value
}

func (v *NullableBTPLiteralUndefined260) Set(val *BTPLiteralUndefined260) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPLiteralUndefined260) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPLiteralUndefined260) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPLiteralUndefined260(val *BTPLiteralUndefined260) *NullableBTPLiteralUndefined260 {
	return &NullableBTPLiteralUndefined260{value: val, isSet: true}
}

func (v NullableBTPLiteralUndefined260) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPLiteralUndefined260) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
