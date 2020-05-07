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

// BTConfiguredDimensionColumnInfo2168AllOf struct for BTConfiguredDimensionColumnInfo2168AllOf
type BTConfiguredDimensionColumnInfo2168AllOf struct {
	BtType *string `json:"btType,omitempty"`
	DimensionId *string `json:"dimensionId,omitempty"`
	ParameterId *string `json:"parameterId,omitempty"`
}

// NewBTConfiguredDimensionColumnInfo2168AllOf instantiates a new BTConfiguredDimensionColumnInfo2168AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfiguredDimensionColumnInfo2168AllOf() *BTConfiguredDimensionColumnInfo2168AllOf {
	this := BTConfiguredDimensionColumnInfo2168AllOf{}
	return &this
}

// NewBTConfiguredDimensionColumnInfo2168AllOfWithDefaults instantiates a new BTConfiguredDimensionColumnInfo2168AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfiguredDimensionColumnInfo2168AllOfWithDefaults() *BTConfiguredDimensionColumnInfo2168AllOf {
	this := BTConfiguredDimensionColumnInfo2168AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfiguredDimensionColumnInfo2168AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredDimensionColumnInfo2168AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConfiguredDimensionColumnInfo2168AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConfiguredDimensionColumnInfo2168AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetDimensionId returns the DimensionId field value if set, zero value otherwise.
func (o *BTConfiguredDimensionColumnInfo2168AllOf) GetDimensionId() string {
	if o == nil || o.DimensionId == nil {
		var ret string
		return ret
	}
	return *o.DimensionId
}

// GetDimensionIdOk returns a tuple with the DimensionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredDimensionColumnInfo2168AllOf) GetDimensionIdOk() (*string, bool) {
	if o == nil || o.DimensionId == nil {
		return nil, false
	}
	return o.DimensionId, true
}

// HasDimensionId returns a boolean if a field has been set.
func (o *BTConfiguredDimensionColumnInfo2168AllOf) HasDimensionId() bool {
	if o != nil && o.DimensionId != nil {
		return true
	}

	return false
}

// SetDimensionId gets a reference to the given string and assigns it to the DimensionId field.
func (o *BTConfiguredDimensionColumnInfo2168AllOf) SetDimensionId(v string) {
	o.DimensionId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTConfiguredDimensionColumnInfo2168AllOf) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredDimensionColumnInfo2168AllOf) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTConfiguredDimensionColumnInfo2168AllOf) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTConfiguredDimensionColumnInfo2168AllOf) SetParameterId(v string) {
	o.ParameterId = &v
}

func (o BTConfiguredDimensionColumnInfo2168AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DimensionId != nil {
		toSerialize["dimensionId"] = o.DimensionId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfiguredDimensionColumnInfo2168AllOf struct {
	value *BTConfiguredDimensionColumnInfo2168AllOf
	isSet bool
}

func (v NullableBTConfiguredDimensionColumnInfo2168AllOf) Get() *BTConfiguredDimensionColumnInfo2168AllOf {
	return v.value
}

func (v *NullableBTConfiguredDimensionColumnInfo2168AllOf) Set(val *BTConfiguredDimensionColumnInfo2168AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfiguredDimensionColumnInfo2168AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfiguredDimensionColumnInfo2168AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfiguredDimensionColumnInfo2168AllOf(val *BTConfiguredDimensionColumnInfo2168AllOf) *NullableBTConfiguredDimensionColumnInfo2168AllOf {
	return &NullableBTConfiguredDimensionColumnInfo2168AllOf{value: val, isSet: true}
}

func (v NullableBTConfiguredDimensionColumnInfo2168AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfiguredDimensionColumnInfo2168AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
