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

// ServiceListPayload Represents the JSON API Payload of a List of Service Items.
type ServiceListPayload struct {
	// The service payloads.
	Data []Service `json:"data"`
	// The User relationships.
	Included *[]UserJSONAPIResponse `json:"included,omitempty"`
	Meta     *QueryMetadata         `json:"meta,omitempty"`
}

// NewServiceListPayload instantiates a new ServiceListPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceListPayload(data []Service) *ServiceListPayload {
	this := ServiceListPayload{}
	this.Data = data
	return &this
}

// NewServiceListPayloadWithDefaults instantiates a new ServiceListPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceListPayloadWithDefaults() *ServiceListPayload {
	this := ServiceListPayload{}
	return &this
}

// GetData returns the Data field value
func (o *ServiceListPayload) GetData() []Service {
	if o == nil {
		var ret []Service
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ServiceListPayload) GetDataOk() (*[]Service, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ServiceListPayload) SetData(v []Service) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ServiceListPayload) GetIncluded() []UserJSONAPIResponse {
	if o == nil || o.Included == nil {
		var ret []UserJSONAPIResponse
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceListPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ServiceListPayload) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []UserJSONAPIResponse and assigns it to the Included field.
func (o *ServiceListPayload) SetIncluded(v []UserJSONAPIResponse) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ServiceListPayload) GetMeta() QueryMetadata {
	if o == nil || o.Meta == nil {
		var ret QueryMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceListPayload) GetMetaOk() (*QueryMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ServiceListPayload) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given QueryMetadata and assigns it to the Meta field.
func (o *ServiceListPayload) SetMeta(v QueryMetadata) {
	o.Meta = &v
}

func (o ServiceListPayload) MarshalJSON() ([]byte, error) {
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

type NullableServiceListPayload struct {
	value *ServiceListPayload
	isSet bool
}

func (v NullableServiceListPayload) Get() *ServiceListPayload {
	return v.value
}

func (v *NullableServiceListPayload) Set(val *ServiceListPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceListPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceListPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceListPayload(val *ServiceListPayload) *NullableServiceListPayload {
	return &NullableServiceListPayload{value: val, isSet: true}
}

func (v NullableServiceListPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceListPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}