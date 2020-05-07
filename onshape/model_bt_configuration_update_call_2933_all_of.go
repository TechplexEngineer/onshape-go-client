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

// BTConfigurationUpdateCall2933AllOf struct for BTConfigurationUpdateCall2933AllOf
type BTConfigurationUpdateCall2933AllOf struct {
	BtType *string `json:"btType,omitempty"`
	ConfigurationParameters *[]BTMConfigurationParameter819 `json:"configurationParameters,omitempty"`
	CurrentConfiguration *[]BTMParameter1 `json:"currentConfiguration,omitempty"`
}

// NewBTConfigurationUpdateCall2933AllOf instantiates a new BTConfigurationUpdateCall2933AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfigurationUpdateCall2933AllOf() *BTConfigurationUpdateCall2933AllOf {
	this := BTConfigurationUpdateCall2933AllOf{}
	return &this
}

// NewBTConfigurationUpdateCall2933AllOfWithDefaults instantiates a new BTConfigurationUpdateCall2933AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfigurationUpdateCall2933AllOfWithDefaults() *BTConfigurationUpdateCall2933AllOf {
	this := BTConfigurationUpdateCall2933AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfigurationUpdateCall2933AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationUpdateCall2933AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConfigurationUpdateCall2933AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConfigurationUpdateCall2933AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigurationParameters returns the ConfigurationParameters field value if set, zero value otherwise.
func (o *BTConfigurationUpdateCall2933AllOf) GetConfigurationParameters() []BTMConfigurationParameter819 {
	if o == nil || o.ConfigurationParameters == nil {
		var ret []BTMConfigurationParameter819
		return ret
	}
	return *o.ConfigurationParameters
}

// GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationUpdateCall2933AllOf) GetConfigurationParametersOk() (*[]BTMConfigurationParameter819, bool) {
	if o == nil || o.ConfigurationParameters == nil {
		return nil, false
	}
	return o.ConfigurationParameters, true
}

// HasConfigurationParameters returns a boolean if a field has been set.
func (o *BTConfigurationUpdateCall2933AllOf) HasConfigurationParameters() bool {
	if o != nil && o.ConfigurationParameters != nil {
		return true
	}

	return false
}

// SetConfigurationParameters gets a reference to the given []BTMConfigurationParameter819 and assigns it to the ConfigurationParameters field.
func (o *BTConfigurationUpdateCall2933AllOf) SetConfigurationParameters(v []BTMConfigurationParameter819) {
	o.ConfigurationParameters = &v
}

// GetCurrentConfiguration returns the CurrentConfiguration field value if set, zero value otherwise.
func (o *BTConfigurationUpdateCall2933AllOf) GetCurrentConfiguration() []BTMParameter1 {
	if o == nil || o.CurrentConfiguration == nil {
		var ret []BTMParameter1
		return ret
	}
	return *o.CurrentConfiguration
}

// GetCurrentConfigurationOk returns a tuple with the CurrentConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationUpdateCall2933AllOf) GetCurrentConfigurationOk() (*[]BTMParameter1, bool) {
	if o == nil || o.CurrentConfiguration == nil {
		return nil, false
	}
	return o.CurrentConfiguration, true
}

// HasCurrentConfiguration returns a boolean if a field has been set.
func (o *BTConfigurationUpdateCall2933AllOf) HasCurrentConfiguration() bool {
	if o != nil && o.CurrentConfiguration != nil {
		return true
	}

	return false
}

// SetCurrentConfiguration gets a reference to the given []BTMParameter1 and assigns it to the CurrentConfiguration field.
func (o *BTConfigurationUpdateCall2933AllOf) SetCurrentConfiguration(v []BTMParameter1) {
	o.CurrentConfiguration = &v
}

func (o BTConfigurationUpdateCall2933AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ConfigurationParameters != nil {
		toSerialize["configurationParameters"] = o.ConfigurationParameters
	}
	if o.CurrentConfiguration != nil {
		toSerialize["currentConfiguration"] = o.CurrentConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfigurationUpdateCall2933AllOf struct {
	value *BTConfigurationUpdateCall2933AllOf
	isSet bool
}

func (v NullableBTConfigurationUpdateCall2933AllOf) Get() *BTConfigurationUpdateCall2933AllOf {
	return v.value
}

func (v *NullableBTConfigurationUpdateCall2933AllOf) Set(val *BTConfigurationUpdateCall2933AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfigurationUpdateCall2933AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfigurationUpdateCall2933AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfigurationUpdateCall2933AllOf(val *BTConfigurationUpdateCall2933AllOf) *NullableBTConfigurationUpdateCall2933AllOf {
	return &NullableBTConfigurationUpdateCall2933AllOf{value: val, isSet: true}
}

func (v NullableBTConfigurationUpdateCall2933AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfigurationUpdateCall2933AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
