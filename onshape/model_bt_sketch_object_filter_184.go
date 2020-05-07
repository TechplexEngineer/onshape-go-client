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

// BTSketchObjectFilter184 struct for BTSketchObjectFilter184
type BTSketchObjectFilter184 struct {
	BTQueryFilter183
	BtType *string `json:"btType,omitempty"`
	IsSketchObject *bool `json:"isSketchObject,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewBTSketchObjectFilter184 instantiates a new BTSketchObjectFilter184 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchObjectFilter184() *BTSketchObjectFilter184 {
	this := BTSketchObjectFilter184{}
	return &this
}

// NewBTSketchObjectFilter184WithDefaults instantiates a new BTSketchObjectFilter184 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchObjectFilter184WithDefaults() *BTSketchObjectFilter184 {
	this := BTSketchObjectFilter184{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchObjectFilter184) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchObjectFilter184) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSketchObjectFilter184) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSketchObjectFilter184) SetBtType(v string) {
	o.BtType = &v
}

// GetIsSketchObject returns the IsSketchObject field value if set, zero value otherwise.
func (o *BTSketchObjectFilter184) GetIsSketchObject() bool {
	if o == nil || o.IsSketchObject == nil {
		var ret bool
		return ret
	}
	return *o.IsSketchObject
}

// GetIsSketchObjectOk returns a tuple with the IsSketchObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchObjectFilter184) GetIsSketchObjectOk() (*bool, bool) {
	if o == nil || o.IsSketchObject == nil {
		return nil, false
	}
	return o.IsSketchObject, true
}

// HasIsSketchObject returns a boolean if a field has been set.
func (o *BTSketchObjectFilter184) HasIsSketchObject() bool {
	if o != nil && o.IsSketchObject != nil {
		return true
	}

	return false
}

// SetIsSketchObject gets a reference to the given bool and assigns it to the IsSketchObject field.
func (o *BTSketchObjectFilter184) SetIsSketchObject(v bool) {
	o.IsSketchObject = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *BTSketchObjectFilter184) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchObjectFilter184) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *BTSketchObjectFilter184) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *BTSketchObjectFilter184) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o BTSketchObjectFilter184) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTQueryFilter183, errBTQueryFilter183 := json.Marshal(o.BTQueryFilter183)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	errBTQueryFilter183 = json.Unmarshal([]byte(serializedBTQueryFilter183), &toSerialize)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsSketchObject != nil {
		toSerialize["isSketchObject"] = o.IsSketchObject
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableBTSketchObjectFilter184 struct {
	value *BTSketchObjectFilter184
	isSet bool
}

func (v NullableBTSketchObjectFilter184) Get() *BTSketchObjectFilter184 {
	return v.value
}

func (v *NullableBTSketchObjectFilter184) Set(val *BTSketchObjectFilter184) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchObjectFilter184) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchObjectFilter184) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchObjectFilter184(val *BTSketchObjectFilter184) *NullableBTSketchObjectFilter184 {
	return &NullableBTSketchObjectFilter184{value: val, isSet: true}
}

func (v NullableBTSketchObjectFilter184) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchObjectFilter184) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
