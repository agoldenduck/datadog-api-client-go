/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// IncidentSelectionListPayload Represents the JSON API Payload of a List of Incident Selection Items.
type IncidentSelectionListPayload struct {
	// The Incident Selection payloads.
	Data []IncidentSelection `json:"data"`
	// The User relationships.
	Included *[]UserJSONAPIResponse `json:"included,omitempty"`
	Meta     *QueryMetadata         `json:"meta,omitempty"`
}

// NewIncidentSelectionListPayload instantiates a new IncidentSelectionListPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentSelectionListPayload(data []IncidentSelection) *IncidentSelectionListPayload {
	this := IncidentSelectionListPayload{}
	this.Data = data
	return &this
}

// NewIncidentSelectionListPayloadWithDefaults instantiates a new IncidentSelectionListPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentSelectionListPayloadWithDefaults() *IncidentSelectionListPayload {
	this := IncidentSelectionListPayload{}
	return &this
}

// GetData returns the Data field value
func (o *IncidentSelectionListPayload) GetData() []IncidentSelection {
	if o == nil {
		var ret []IncidentSelection
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentSelectionListPayload) GetDataOk() (*[]IncidentSelection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IncidentSelectionListPayload) SetData(v []IncidentSelection) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentSelectionListPayload) GetIncluded() []UserJSONAPIResponse {
	if o == nil || o.Included == nil {
		var ret []UserJSONAPIResponse
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSelectionListPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentSelectionListPayload) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []UserJSONAPIResponse and assigns it to the Included field.
func (o *IncidentSelectionListPayload) SetIncluded(v []UserJSONAPIResponse) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IncidentSelectionListPayload) GetMeta() QueryMetadata {
	if o == nil || o.Meta == nil {
		var ret QueryMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSelectionListPayload) GetMetaOk() (*QueryMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IncidentSelectionListPayload) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given QueryMetadata and assigns it to the Meta field.
func (o *IncidentSelectionListPayload) SetMeta(v QueryMetadata) {
	o.Meta = &v
}

func (o IncidentSelectionListPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentSelectionListPayload struct {
	value *IncidentSelectionListPayload
	isSet bool
}

func (v NullableIncidentSelectionListPayload) Get() *IncidentSelectionListPayload {
	return v.value
}

func (v *NullableIncidentSelectionListPayload) Set(val *IncidentSelectionListPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentSelectionListPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentSelectionListPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentSelectionListPayload(val *IncidentSelectionListPayload) *NullableIncidentSelectionListPayload {
	return &NullableIncidentSelectionListPayload{value: val, isSet: true}
}

func (v NullableIncidentSelectionListPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentSelectionListPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}