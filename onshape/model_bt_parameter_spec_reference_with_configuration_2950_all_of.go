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

// BTParameterSpecReferenceWithConfiguration2950AllOf struct for BTParameterSpecReferenceWithConfiguration2950AllOf
type BTParameterSpecReferenceWithConfiguration2950AllOf struct {
	AllowAssemblies *bool `json:"allowAssemblies,omitempty"`
	AllowedInsertableTypes *[]string `json:"allowedInsertableTypes,omitempty"`
	BtType *string `json:"btType,omitempty"`
	MaxNumberOfPicks *int32 `json:"maxNumberOfPicks,omitempty"`
}

// NewBTParameterSpecReferenceWithConfiguration2950AllOf instantiates a new BTParameterSpecReferenceWithConfiguration2950AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecReferenceWithConfiguration2950AllOf() *BTParameterSpecReferenceWithConfiguration2950AllOf {
	this := BTParameterSpecReferenceWithConfiguration2950AllOf{}
	return &this
}

// NewBTParameterSpecReferenceWithConfiguration2950AllOfWithDefaults instantiates a new BTParameterSpecReferenceWithConfiguration2950AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecReferenceWithConfiguration2950AllOfWithDefaults() *BTParameterSpecReferenceWithConfiguration2950AllOf {
	this := BTParameterSpecReferenceWithConfiguration2950AllOf{}
	return &this
}

// GetAllowAssemblies returns the AllowAssemblies field value if set, zero value otherwise.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) GetAllowAssemblies() bool {
	if o == nil || o.AllowAssemblies == nil {
		var ret bool
		return ret
	}
	return *o.AllowAssemblies
}

// GetAllowAssembliesOk returns a tuple with the AllowAssemblies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) GetAllowAssembliesOk() (*bool, bool) {
	if o == nil || o.AllowAssemblies == nil {
		return nil, false
	}
	return o.AllowAssemblies, true
}

// HasAllowAssemblies returns a boolean if a field has been set.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) HasAllowAssemblies() bool {
	if o != nil && o.AllowAssemblies != nil {
		return true
	}

	return false
}

// SetAllowAssemblies gets a reference to the given bool and assigns it to the AllowAssemblies field.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) SetAllowAssemblies(v bool) {
	o.AllowAssemblies = &v
}

// GetAllowedInsertableTypes returns the AllowedInsertableTypes field value if set, zero value otherwise.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) GetAllowedInsertableTypes() []string {
	if o == nil || o.AllowedInsertableTypes == nil {
		var ret []string
		return ret
	}
	return *o.AllowedInsertableTypes
}

// GetAllowedInsertableTypesOk returns a tuple with the AllowedInsertableTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) GetAllowedInsertableTypesOk() (*[]string, bool) {
	if o == nil || o.AllowedInsertableTypes == nil {
		return nil, false
	}
	return o.AllowedInsertableTypes, true
}

// HasAllowedInsertableTypes returns a boolean if a field has been set.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) HasAllowedInsertableTypes() bool {
	if o != nil && o.AllowedInsertableTypes != nil {
		return true
	}

	return false
}

// SetAllowedInsertableTypes gets a reference to the given []string and assigns it to the AllowedInsertableTypes field.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) SetAllowedInsertableTypes(v []string) {
	o.AllowedInsertableTypes = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetMaxNumberOfPicks returns the MaxNumberOfPicks field value if set, zero value otherwise.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) GetMaxNumberOfPicks() int32 {
	if o == nil || o.MaxNumberOfPicks == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumberOfPicks
}

// GetMaxNumberOfPicksOk returns a tuple with the MaxNumberOfPicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) GetMaxNumberOfPicksOk() (*int32, bool) {
	if o == nil || o.MaxNumberOfPicks == nil {
		return nil, false
	}
	return o.MaxNumberOfPicks, true
}

// HasMaxNumberOfPicks returns a boolean if a field has been set.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) HasMaxNumberOfPicks() bool {
	if o != nil && o.MaxNumberOfPicks != nil {
		return true
	}

	return false
}

// SetMaxNumberOfPicks gets a reference to the given int32 and assigns it to the MaxNumberOfPicks field.
func (o *BTParameterSpecReferenceWithConfiguration2950AllOf) SetMaxNumberOfPicks(v int32) {
	o.MaxNumberOfPicks = &v
}

func (o BTParameterSpecReferenceWithConfiguration2950AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowAssemblies != nil {
		toSerialize["allowAssemblies"] = o.AllowAssemblies
	}
	if o.AllowedInsertableTypes != nil {
		toSerialize["allowedInsertableTypes"] = o.AllowedInsertableTypes
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.MaxNumberOfPicks != nil {
		toSerialize["maxNumberOfPicks"] = o.MaxNumberOfPicks
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecReferenceWithConfiguration2950AllOf struct {
	value *BTParameterSpecReferenceWithConfiguration2950AllOf
	isSet bool
}

func (v NullableBTParameterSpecReferenceWithConfiguration2950AllOf) Get() *BTParameterSpecReferenceWithConfiguration2950AllOf {
	return v.value
}

func (v *NullableBTParameterSpecReferenceWithConfiguration2950AllOf) Set(val *BTParameterSpecReferenceWithConfiguration2950AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecReferenceWithConfiguration2950AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecReferenceWithConfiguration2950AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecReferenceWithConfiguration2950AllOf(val *BTParameterSpecReferenceWithConfiguration2950AllOf) *NullableBTParameterSpecReferenceWithConfiguration2950AllOf {
	return &NullableBTParameterSpecReferenceWithConfiguration2950AllOf{value: val, isSet: true}
}

func (v NullableBTParameterSpecReferenceWithConfiguration2950AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecReferenceWithConfiguration2950AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
