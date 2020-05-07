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

// BTParameterSpecReference2789 struct for BTParameterSpecReference2789
type BTParameterSpecReference2789 struct {
	BTParameterSpec6
	BtType *string `json:"btType,omitempty"`
}

// NewBTParameterSpecReference2789 instantiates a new BTParameterSpecReference2789 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecReference2789() *BTParameterSpecReference2789 {
	this := BTParameterSpecReference2789{}
	return &this
}

// NewBTParameterSpecReference2789WithDefaults instantiates a new BTParameterSpecReference2789 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecReference2789WithDefaults() *BTParameterSpecReference2789 {
	this := BTParameterSpecReference2789{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecReference2789) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReference2789) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecReference2789) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecReference2789) SetBtType(v string) {
	o.BtType = &v
}

func (o BTParameterSpecReference2789) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTParameterSpec6, errBTParameterSpec6 := json.Marshal(o.BTParameterSpec6)
	if errBTParameterSpec6 != nil {
		return []byte{}, errBTParameterSpec6
	}
	errBTParameterSpec6 = json.Unmarshal([]byte(serializedBTParameterSpec6), &toSerialize)
	if errBTParameterSpec6 != nil {
		return []byte{}, errBTParameterSpec6
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecReference2789 struct {
	value *BTParameterSpecReference2789
	isSet bool
}

func (v NullableBTParameterSpecReference2789) Get() *BTParameterSpecReference2789 {
	return v.value
}

func (v *NullableBTParameterSpecReference2789) Set(val *BTParameterSpecReference2789) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecReference2789) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecReference2789) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecReference2789(val *BTParameterSpecReference2789) *NullableBTParameterSpecReference2789 {
	return &NullableBTParameterSpecReference2789{value: val, isSet: true}
}

func (v NullableBTParameterSpecReference2789) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecReference2789) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
