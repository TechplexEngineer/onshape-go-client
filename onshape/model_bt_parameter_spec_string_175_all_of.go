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

// BTParameterSpecString175AllOf struct for BTParameterSpecString175AllOf
type BTParameterSpecString175AllOf struct {
	BtType *string `json:"btType,omitempty"`
	FormatConditions *[]BTStringFormatCondition683 `json:"formatConditions,omitempty"`
}

// NewBTParameterSpecString175AllOf instantiates a new BTParameterSpecString175AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecString175AllOf() *BTParameterSpecString175AllOf {
	this := BTParameterSpecString175AllOf{}
	return &this
}

// NewBTParameterSpecString175AllOfWithDefaults instantiates a new BTParameterSpecString175AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecString175AllOfWithDefaults() *BTParameterSpecString175AllOf {
	this := BTParameterSpecString175AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecString175AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecString175AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecString175AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetFormatConditions returns the FormatConditions field value if set, zero value otherwise.
func (o *BTParameterSpecString175AllOf) GetFormatConditions() []BTStringFormatCondition683 {
	if o == nil || o.FormatConditions == nil {
		var ret []BTStringFormatCondition683
		return ret
	}
	return *o.FormatConditions
}

// GetFormatConditionsOk returns a tuple with the FormatConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175AllOf) GetFormatConditionsOk() (*[]BTStringFormatCondition683, bool) {
	if o == nil || o.FormatConditions == nil {
		return nil, false
	}
	return o.FormatConditions, true
}

// HasFormatConditions returns a boolean if a field has been set.
func (o *BTParameterSpecString175AllOf) HasFormatConditions() bool {
	if o != nil && o.FormatConditions != nil {
		return true
	}

	return false
}

// SetFormatConditions gets a reference to the given []BTStringFormatCondition683 and assigns it to the FormatConditions field.
func (o *BTParameterSpecString175AllOf) SetFormatConditions(v []BTStringFormatCondition683) {
	o.FormatConditions = &v
}

func (o BTParameterSpecString175AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FormatConditions != nil {
		toSerialize["formatConditions"] = o.FormatConditions
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecString175AllOf struct {
	value *BTParameterSpecString175AllOf
	isSet bool
}

func (v NullableBTParameterSpecString175AllOf) Get() *BTParameterSpecString175AllOf {
	return v.value
}

func (v *NullableBTParameterSpecString175AllOf) Set(val *BTParameterSpecString175AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecString175AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecString175AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecString175AllOf(val *BTParameterSpecString175AllOf) *NullableBTParameterSpecString175AllOf {
	return &NullableBTParameterSpecString175AllOf{value: val, isSet: true}
}

func (v NullableBTParameterSpecString175AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecString175AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
