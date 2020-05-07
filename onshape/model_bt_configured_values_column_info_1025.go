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

// BTConfiguredValuesColumnInfo1025 struct for BTConfiguredValuesColumnInfo1025
type BTConfiguredValuesColumnInfo1025 struct {
	BTTableColumnInfo1222
	BtType *string `json:"btType,omitempty"`
	ParentId *string `json:"parentId,omitempty"`
	ParentName *string `json:"parentName,omitempty"`
	ParentType *string `json:"parentType,omitempty"`
}

// NewBTConfiguredValuesColumnInfo1025 instantiates a new BTConfiguredValuesColumnInfo1025 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfiguredValuesColumnInfo1025() *BTConfiguredValuesColumnInfo1025 {
	this := BTConfiguredValuesColumnInfo1025{}
	return &this
}

// NewBTConfiguredValuesColumnInfo1025WithDefaults instantiates a new BTConfiguredValuesColumnInfo1025 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfiguredValuesColumnInfo1025WithDefaults() *BTConfiguredValuesColumnInfo1025 {
	this := BTConfiguredValuesColumnInfo1025{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfiguredValuesColumnInfo1025) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredValuesColumnInfo1025) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConfiguredValuesColumnInfo1025) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConfiguredValuesColumnInfo1025) SetBtType(v string) {
	o.BtType = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTConfiguredValuesColumnInfo1025) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredValuesColumnInfo1025) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTConfiguredValuesColumnInfo1025) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTConfiguredValuesColumnInfo1025) SetParentId(v string) {
	o.ParentId = &v
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *BTConfiguredValuesColumnInfo1025) GetParentName() string {
	if o == nil || o.ParentName == nil {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredValuesColumnInfo1025) GetParentNameOk() (*string, bool) {
	if o == nil || o.ParentName == nil {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *BTConfiguredValuesColumnInfo1025) HasParentName() bool {
	if o != nil && o.ParentName != nil {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *BTConfiguredValuesColumnInfo1025) SetParentName(v string) {
	o.ParentName = &v
}

// GetParentType returns the ParentType field value if set, zero value otherwise.
func (o *BTConfiguredValuesColumnInfo1025) GetParentType() string {
	if o == nil || o.ParentType == nil {
		var ret string
		return ret
	}
	return *o.ParentType
}

// GetParentTypeOk returns a tuple with the ParentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredValuesColumnInfo1025) GetParentTypeOk() (*string, bool) {
	if o == nil || o.ParentType == nil {
		return nil, false
	}
	return o.ParentType, true
}

// HasParentType returns a boolean if a field has been set.
func (o *BTConfiguredValuesColumnInfo1025) HasParentType() bool {
	if o != nil && o.ParentType != nil {
		return true
	}

	return false
}

// SetParentType gets a reference to the given string and assigns it to the ParentType field.
func (o *BTConfiguredValuesColumnInfo1025) SetParentType(v string) {
	o.ParentType = &v
}

func (o BTConfiguredValuesColumnInfo1025) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTTableColumnInfo1222, errBTTableColumnInfo1222 := json.Marshal(o.BTTableColumnInfo1222)
	if errBTTableColumnInfo1222 != nil {
		return []byte{}, errBTTableColumnInfo1222
	}
	errBTTableColumnInfo1222 = json.Unmarshal([]byte(serializedBTTableColumnInfo1222), &toSerialize)
	if errBTTableColumnInfo1222 != nil {
		return []byte{}, errBTTableColumnInfo1222
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.ParentName != nil {
		toSerialize["parentName"] = o.ParentName
	}
	if o.ParentType != nil {
		toSerialize["parentType"] = o.ParentType
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfiguredValuesColumnInfo1025 struct {
	value *BTConfiguredValuesColumnInfo1025
	isSet bool
}

func (v NullableBTConfiguredValuesColumnInfo1025) Get() *BTConfiguredValuesColumnInfo1025 {
	return v.value
}

func (v *NullableBTConfiguredValuesColumnInfo1025) Set(val *BTConfiguredValuesColumnInfo1025) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfiguredValuesColumnInfo1025) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfiguredValuesColumnInfo1025) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfiguredValuesColumnInfo1025(val *BTConfiguredValuesColumnInfo1025) *NullableBTConfiguredValuesColumnInfo1025 {
	return &NullableBTConfiguredValuesColumnInfo1025{value: val, isSet: true}
}

func (v NullableBTConfiguredValuesColumnInfo1025) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfiguredValuesColumnInfo1025) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
