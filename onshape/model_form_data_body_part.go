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

// FormDataBodyPart struct for FormDataBodyPart
type FormDataBodyPart struct {
	ContentDisposition *ContentDisposition `json:"contentDisposition,omitempty"`
	Entity *map[string]interface{} `json:"entity,omitempty"`
	FormDataContentDisposition *FormDataContentDisposition `json:"formDataContentDisposition,omitempty"`
	Headers *map[string][]string `json:"headers,omitempty"`
	MediaType *BodyPartMediaType `json:"mediaType,omitempty"`
	MessageBodyWorkers *map[string]interface{} `json:"messageBodyWorkers,omitempty"`
	Name *string `json:"name,omitempty"`
	ParameterizedHeaders *map[string][]ParameterizedHeader `json:"parameterizedHeaders,omitempty"`
	Parent *MultiPart `json:"parent,omitempty"`
	Providers *map[string]interface{} `json:"providers,omitempty"`
	Simple *bool `json:"simple,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewFormDataBodyPart instantiates a new FormDataBodyPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormDataBodyPart() *FormDataBodyPart {
	this := FormDataBodyPart{}
	return &this
}

// NewFormDataBodyPartWithDefaults instantiates a new FormDataBodyPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormDataBodyPartWithDefaults() *FormDataBodyPart {
	this := FormDataBodyPart{}
	return &this
}

// GetContentDisposition returns the ContentDisposition field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetContentDisposition() ContentDisposition {
	if o == nil || o.ContentDisposition == nil {
		var ret ContentDisposition
		return ret
	}
	return *o.ContentDisposition
}

// GetContentDispositionOk returns a tuple with the ContentDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetContentDispositionOk() (*ContentDisposition, bool) {
	if o == nil || o.ContentDisposition == nil {
		return nil, false
	}
	return o.ContentDisposition, true
}

// HasContentDisposition returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasContentDisposition() bool {
	if o != nil && o.ContentDisposition != nil {
		return true
	}

	return false
}

// SetContentDisposition gets a reference to the given ContentDisposition and assigns it to the ContentDisposition field.
func (o *FormDataBodyPart) SetContentDisposition(v ContentDisposition) {
	o.ContentDisposition = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetEntity() map[string]interface{} {
	if o == nil || o.Entity == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetEntityOk() (*map[string]interface{}, bool) {
	if o == nil || o.Entity == nil {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasEntity() bool {
	if o != nil && o.Entity != nil {
		return true
	}

	return false
}

// SetEntity gets a reference to the given map[string]interface{} and assigns it to the Entity field.
func (o *FormDataBodyPart) SetEntity(v map[string]interface{}) {
	o.Entity = &v
}

// GetFormDataContentDisposition returns the FormDataContentDisposition field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetFormDataContentDisposition() FormDataContentDisposition {
	if o == nil || o.FormDataContentDisposition == nil {
		var ret FormDataContentDisposition
		return ret
	}
	return *o.FormDataContentDisposition
}

// GetFormDataContentDispositionOk returns a tuple with the FormDataContentDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetFormDataContentDispositionOk() (*FormDataContentDisposition, bool) {
	if o == nil || o.FormDataContentDisposition == nil {
		return nil, false
	}
	return o.FormDataContentDisposition, true
}

// HasFormDataContentDisposition returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasFormDataContentDisposition() bool {
	if o != nil && o.FormDataContentDisposition != nil {
		return true
	}

	return false
}

// SetFormDataContentDisposition gets a reference to the given FormDataContentDisposition and assigns it to the FormDataContentDisposition field.
func (o *FormDataBodyPart) SetFormDataContentDisposition(v FormDataContentDisposition) {
	o.FormDataContentDisposition = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetHeaders() map[string][]string {
	if o == nil || o.Headers == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetHeadersOk() (*map[string][]string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string][]string and assigns it to the Headers field.
func (o *FormDataBodyPart) SetHeaders(v map[string][]string) {
	o.Headers = &v
}

// GetMediaType returns the MediaType field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetMediaType() BodyPartMediaType {
	if o == nil || o.MediaType == nil {
		var ret BodyPartMediaType
		return ret
	}
	return *o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetMediaTypeOk() (*BodyPartMediaType, bool) {
	if o == nil || o.MediaType == nil {
		return nil, false
	}
	return o.MediaType, true
}

// HasMediaType returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasMediaType() bool {
	if o != nil && o.MediaType != nil {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given BodyPartMediaType and assigns it to the MediaType field.
func (o *FormDataBodyPart) SetMediaType(v BodyPartMediaType) {
	o.MediaType = &v
}

// GetMessageBodyWorkers returns the MessageBodyWorkers field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetMessageBodyWorkers() map[string]interface{} {
	if o == nil || o.MessageBodyWorkers == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.MessageBodyWorkers
}

// GetMessageBodyWorkersOk returns a tuple with the MessageBodyWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetMessageBodyWorkersOk() (*map[string]interface{}, bool) {
	if o == nil || o.MessageBodyWorkers == nil {
		return nil, false
	}
	return o.MessageBodyWorkers, true
}

// HasMessageBodyWorkers returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasMessageBodyWorkers() bool {
	if o != nil && o.MessageBodyWorkers != nil {
		return true
	}

	return false
}

// SetMessageBodyWorkers gets a reference to the given map[string]interface{} and assigns it to the MessageBodyWorkers field.
func (o *FormDataBodyPart) SetMessageBodyWorkers(v map[string]interface{}) {
	o.MessageBodyWorkers = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FormDataBodyPart) SetName(v string) {
	o.Name = &v
}

// GetParameterizedHeaders returns the ParameterizedHeaders field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetParameterizedHeaders() map[string][]ParameterizedHeader {
	if o == nil || o.ParameterizedHeaders == nil {
		var ret map[string][]ParameterizedHeader
		return ret
	}
	return *o.ParameterizedHeaders
}

// GetParameterizedHeadersOk returns a tuple with the ParameterizedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetParameterizedHeadersOk() (*map[string][]ParameterizedHeader, bool) {
	if o == nil || o.ParameterizedHeaders == nil {
		return nil, false
	}
	return o.ParameterizedHeaders, true
}

// HasParameterizedHeaders returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasParameterizedHeaders() bool {
	if o != nil && o.ParameterizedHeaders != nil {
		return true
	}

	return false
}

// SetParameterizedHeaders gets a reference to the given map[string][]ParameterizedHeader and assigns it to the ParameterizedHeaders field.
func (o *FormDataBodyPart) SetParameterizedHeaders(v map[string][]ParameterizedHeader) {
	o.ParameterizedHeaders = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetParent() MultiPart {
	if o == nil || o.Parent == nil {
		var ret MultiPart
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetParentOk() (*MultiPart, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given MultiPart and assigns it to the Parent field.
func (o *FormDataBodyPart) SetParent(v MultiPart) {
	o.Parent = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetProviders() map[string]interface{} {
	if o == nil || o.Providers == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetProvidersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Providers == nil {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasProviders() bool {
	if o != nil && o.Providers != nil {
		return true
	}

	return false
}

// SetProviders gets a reference to the given map[string]interface{} and assigns it to the Providers field.
func (o *FormDataBodyPart) SetProviders(v map[string]interface{}) {
	o.Providers = &v
}

// GetSimple returns the Simple field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetSimple() bool {
	if o == nil || o.Simple == nil {
		var ret bool
		return ret
	}
	return *o.Simple
}

// GetSimpleOk returns a tuple with the Simple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetSimpleOk() (*bool, bool) {
	if o == nil || o.Simple == nil {
		return nil, false
	}
	return o.Simple, true
}

// HasSimple returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasSimple() bool {
	if o != nil && o.Simple != nil {
		return true
	}

	return false
}

// SetSimple gets a reference to the given bool and assigns it to the Simple field.
func (o *FormDataBodyPart) SetSimple(v bool) {
	o.Simple = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FormDataBodyPart) SetValue(v string) {
	o.Value = &v
}

func (o FormDataBodyPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentDisposition != nil {
		toSerialize["contentDisposition"] = o.ContentDisposition
	}
	if o.Entity != nil {
		toSerialize["entity"] = o.Entity
	}
	if o.FormDataContentDisposition != nil {
		toSerialize["formDataContentDisposition"] = o.FormDataContentDisposition
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.MediaType != nil {
		toSerialize["mediaType"] = o.MediaType
	}
	if o.MessageBodyWorkers != nil {
		toSerialize["messageBodyWorkers"] = o.MessageBodyWorkers
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ParameterizedHeaders != nil {
		toSerialize["parameterizedHeaders"] = o.ParameterizedHeaders
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}
	if o.Simple != nil {
		toSerialize["simple"] = o.Simple
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableFormDataBodyPart struct {
	value *FormDataBodyPart
	isSet bool
}

func (v NullableFormDataBodyPart) Get() *FormDataBodyPart {
	return v.value
}

func (v *NullableFormDataBodyPart) Set(val *FormDataBodyPart) {
	v.value = val
	v.isSet = true
}

func (v NullableFormDataBodyPart) IsSet() bool {
	return v.isSet
}

func (v *NullableFormDataBodyPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormDataBodyPart(val *FormDataBodyPart) *NullableFormDataBodyPart {
	return &NullableFormDataBodyPart{value: val, isSet: true}
}

func (v NullableFormDataBodyPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormDataBodyPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
