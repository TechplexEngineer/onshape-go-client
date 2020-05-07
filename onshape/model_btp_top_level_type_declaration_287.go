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

// BTPTopLevelTypeDeclaration287 struct for BTPTopLevelTypeDeclaration287
type BTPTopLevelTypeDeclaration287 struct {
	BTPTopLevelNode286
	BtType *string `json:"btType,omitempty"`
	Name *BTPIdentifier8 `json:"name,omitempty"`
	SpaceAfterVersion *BTPSpace10 `json:"spaceAfterVersion,omitempty"`
	Version *BTPLiteralNumber258 `json:"version,omitempty"`
}

// NewBTPTopLevelTypeDeclaration287 instantiates a new BTPTopLevelTypeDeclaration287 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPTopLevelTypeDeclaration287() *BTPTopLevelTypeDeclaration287 {
	this := BTPTopLevelTypeDeclaration287{}
	return &this
}

// NewBTPTopLevelTypeDeclaration287WithDefaults instantiates a new BTPTopLevelTypeDeclaration287 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPTopLevelTypeDeclaration287WithDefaults() *BTPTopLevelTypeDeclaration287 {
	this := BTPTopLevelTypeDeclaration287{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPTopLevelTypeDeclaration287) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelTypeDeclaration287) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPTopLevelTypeDeclaration287) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPTopLevelTypeDeclaration287) SetBtType(v string) {
	o.BtType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPTopLevelTypeDeclaration287) GetName() BTPIdentifier8 {
	if o == nil || o.Name == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelTypeDeclaration287) GetNameOk() (*BTPIdentifier8, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPTopLevelTypeDeclaration287) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given BTPIdentifier8 and assigns it to the Name field.
func (o *BTPTopLevelTypeDeclaration287) SetName(v BTPIdentifier8) {
	o.Name = &v
}

// GetSpaceAfterVersion returns the SpaceAfterVersion field value if set, zero value otherwise.
func (o *BTPTopLevelTypeDeclaration287) GetSpaceAfterVersion() BTPSpace10 {
	if o == nil || o.SpaceAfterVersion == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterVersion
}

// GetSpaceAfterVersionOk returns a tuple with the SpaceAfterVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelTypeDeclaration287) GetSpaceAfterVersionOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterVersion == nil {
		return nil, false
	}
	return o.SpaceAfterVersion, true
}

// HasSpaceAfterVersion returns a boolean if a field has been set.
func (o *BTPTopLevelTypeDeclaration287) HasSpaceAfterVersion() bool {
	if o != nil && o.SpaceAfterVersion != nil {
		return true
	}

	return false
}

// SetSpaceAfterVersion gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterVersion field.
func (o *BTPTopLevelTypeDeclaration287) SetSpaceAfterVersion(v BTPSpace10) {
	o.SpaceAfterVersion = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTPTopLevelTypeDeclaration287) GetVersion() BTPLiteralNumber258 {
	if o == nil || o.Version == nil {
		var ret BTPLiteralNumber258
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelTypeDeclaration287) GetVersionOk() (*BTPLiteralNumber258, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTPTopLevelTypeDeclaration287) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given BTPLiteralNumber258 and assigns it to the Version field.
func (o *BTPTopLevelTypeDeclaration287) SetVersion(v BTPLiteralNumber258) {
	o.Version = &v
}

func (o BTPTopLevelTypeDeclaration287) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPTopLevelNode286, errBTPTopLevelNode286 := json.Marshal(o.BTPTopLevelNode286)
	if errBTPTopLevelNode286 != nil {
		return []byte{}, errBTPTopLevelNode286
	}
	errBTPTopLevelNode286 = json.Unmarshal([]byte(serializedBTPTopLevelNode286), &toSerialize)
	if errBTPTopLevelNode286 != nil {
		return []byte{}, errBTPTopLevelNode286
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SpaceAfterVersion != nil {
		toSerialize["spaceAfterVersion"] = o.SpaceAfterVersion
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableBTPTopLevelTypeDeclaration287 struct {
	value *BTPTopLevelTypeDeclaration287
	isSet bool
}

func (v NullableBTPTopLevelTypeDeclaration287) Get() *BTPTopLevelTypeDeclaration287 {
	return v.value
}

func (v *NullableBTPTopLevelTypeDeclaration287) Set(val *BTPTopLevelTypeDeclaration287) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPTopLevelTypeDeclaration287) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPTopLevelTypeDeclaration287) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPTopLevelTypeDeclaration287(val *BTPTopLevelTypeDeclaration287) *NullableBTPTopLevelTypeDeclaration287 {
	return &NullableBTPTopLevelTypeDeclaration287{value: val, isSet: true}
}

func (v NullableBTPTopLevelTypeDeclaration287) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPTopLevelTypeDeclaration287) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
