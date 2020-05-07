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

// BTListResponseBTMetadataPropertySummaryInfo struct for BTListResponseBTMetadataPropertySummaryInfo
type BTListResponseBTMetadataPropertySummaryInfo struct {
	Href *string `json:"href,omitempty"`
	Items *[]BTMetadataPropertySummaryInfo `json:"items,omitempty"`
	Next *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
}

// NewBTListResponseBTMetadataPropertySummaryInfo instantiates a new BTListResponseBTMetadataPropertySummaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTListResponseBTMetadataPropertySummaryInfo() *BTListResponseBTMetadataPropertySummaryInfo {
	this := BTListResponseBTMetadataPropertySummaryInfo{}
	return &this
}

// NewBTListResponseBTMetadataPropertySummaryInfoWithDefaults instantiates a new BTListResponseBTMetadataPropertySummaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTListResponseBTMetadataPropertySummaryInfoWithDefaults() *BTListResponseBTMetadataPropertySummaryInfo {
	this := BTListResponseBTMetadataPropertySummaryInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTListResponseBTMetadataPropertySummaryInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTListResponseBTMetadataPropertySummaryInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTListResponseBTMetadataPropertySummaryInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTListResponseBTMetadataPropertySummaryInfo) SetHref(v string) {
	o.Href = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BTListResponseBTMetadataPropertySummaryInfo) GetItems() []BTMetadataPropertySummaryInfo {
	if o == nil || o.Items == nil {
		var ret []BTMetadataPropertySummaryInfo
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTListResponseBTMetadataPropertySummaryInfo) GetItemsOk() (*[]BTMetadataPropertySummaryInfo, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BTListResponseBTMetadataPropertySummaryInfo) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BTMetadataPropertySummaryInfo and assigns it to the Items field.
func (o *BTListResponseBTMetadataPropertySummaryInfo) SetItems(v []BTMetadataPropertySummaryInfo) {
	o.Items = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *BTListResponseBTMetadataPropertySummaryInfo) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTListResponseBTMetadataPropertySummaryInfo) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *BTListResponseBTMetadataPropertySummaryInfo) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *BTListResponseBTMetadataPropertySummaryInfo) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *BTListResponseBTMetadataPropertySummaryInfo) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTListResponseBTMetadataPropertySummaryInfo) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *BTListResponseBTMetadataPropertySummaryInfo) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *BTListResponseBTMetadataPropertySummaryInfo) SetPrevious(v string) {
	o.Previous = &v
}

func (o BTListResponseBTMetadataPropertySummaryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	return json.Marshal(toSerialize)
}

type NullableBTListResponseBTMetadataPropertySummaryInfo struct {
	value *BTListResponseBTMetadataPropertySummaryInfo
	isSet bool
}

func (v NullableBTListResponseBTMetadataPropertySummaryInfo) Get() *BTListResponseBTMetadataPropertySummaryInfo {
	return v.value
}

func (v *NullableBTListResponseBTMetadataPropertySummaryInfo) Set(val *BTListResponseBTMetadataPropertySummaryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTListResponseBTMetadataPropertySummaryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTListResponseBTMetadataPropertySummaryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTListResponseBTMetadataPropertySummaryInfo(val *BTListResponseBTMetadataPropertySummaryInfo) *NullableBTListResponseBTMetadataPropertySummaryInfo {
	return &NullableBTListResponseBTMetadataPropertySummaryInfo{value: val, isSet: true}
}

func (v NullableBTListResponseBTMetadataPropertySummaryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTListResponseBTMetadataPropertySummaryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
