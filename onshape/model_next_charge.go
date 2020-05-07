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

// NextCharge struct for NextCharge
type NextCharge struct {
	Amount *int64 `json:"amount,omitempty"`
	CurrentPeriodEnd *JSONTime `json:"currentPeriodEnd,omitempty"`
	Interval *string `json:"interval,omitempty"`
	Total *int64 `json:"total,omitempty"`
}

// NewNextCharge instantiates a new NextCharge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNextCharge() *NextCharge {
	this := NextCharge{}
	return &this
}

// NewNextChargeWithDefaults instantiates a new NextCharge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNextChargeWithDefaults() *NextCharge {
	this := NextCharge{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *NextCharge) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextCharge) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *NextCharge) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *NextCharge) SetAmount(v int64) {
	o.Amount = &v
}

// GetCurrentPeriodEnd returns the CurrentPeriodEnd field value if set, zero value otherwise.
func (o *NextCharge) GetCurrentPeriodEnd() JSONTime {
	if o == nil || o.CurrentPeriodEnd == nil {
		var ret JSONTime
		return ret
	}
	return *o.CurrentPeriodEnd
}

// GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextCharge) GetCurrentPeriodEndOk() (*JSONTime, bool) {
	if o == nil || o.CurrentPeriodEnd == nil {
		return nil, false
	}
	return o.CurrentPeriodEnd, true
}

// HasCurrentPeriodEnd returns a boolean if a field has been set.
func (o *NextCharge) HasCurrentPeriodEnd() bool {
	if o != nil && o.CurrentPeriodEnd != nil {
		return true
	}

	return false
}

// SetCurrentPeriodEnd gets a reference to the given JSONTime and assigns it to the CurrentPeriodEnd field.
func (o *NextCharge) SetCurrentPeriodEnd(v JSONTime) {
	o.CurrentPeriodEnd = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *NextCharge) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextCharge) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *NextCharge) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *NextCharge) SetInterval(v string) {
	o.Interval = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *NextCharge) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextCharge) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *NextCharge) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *NextCharge) SetTotal(v int64) {
	o.Total = &v
}

func (o NextCharge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.CurrentPeriodEnd != nil {
		toSerialize["currentPeriodEnd"] = o.CurrentPeriodEnd
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableNextCharge struct {
	value *NextCharge
	isSet bool
}

func (v NullableNextCharge) Get() *NextCharge {
	return v.value
}

func (v *NullableNextCharge) Set(val *NextCharge) {
	v.value = val
	v.isSet = true
}

func (v NullableNextCharge) IsSet() bool {
	return v.isSet
}

func (v *NullableNextCharge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNextCharge(val *NextCharge) *NullableNextCharge {
	return &NullableNextCharge{value: val, isSet: true}
}

func (v NullableNextCharge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNextCharge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
