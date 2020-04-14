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

// IncidentTodo Represents an Incident Todo Item.
type IncidentTodo struct {
	Attributes *IncidentTodoAttributes `json:"attributes,omitempty"`
	// Incident Todo ID.
	Id            *string                    `json:"id,omitempty"`
	Relationships *IncidentTodoRelationships `json:"relationships,omitempty"`
	// JSONAPI Model Type.
	Type *string `json:"type,omitempty"`
}

// NewIncidentTodo instantiates a new IncidentTodo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentTodo() *IncidentTodo {
	this := IncidentTodo{}
	return &this
}

// NewIncidentTodoWithDefaults instantiates a new IncidentTodo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentTodoWithDefaults() *IncidentTodo {
	this := IncidentTodo{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentTodo) GetAttributes() IncidentTodoAttributes {
	if o == nil || o.Attributes == nil {
		var ret IncidentTodoAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTodo) GetAttributesOk() (*IncidentTodoAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentTodo) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given IncidentTodoAttributes and assigns it to the Attributes field.
func (o *IncidentTodo) SetAttributes(v IncidentTodoAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IncidentTodo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTodo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IncidentTodo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IncidentTodo) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IncidentTodo) GetRelationships() IncidentTodoRelationships {
	if o == nil || o.Relationships == nil {
		var ret IncidentTodoRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTodo) GetRelationshipsOk() (*IncidentTodoRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentTodo) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given IncidentTodoRelationships and assigns it to the Relationships field.
func (o *IncidentTodo) SetRelationships(v IncidentTodoRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IncidentTodo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTodo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IncidentTodo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IncidentTodo) SetType(v string) {
	o.Type = &v
}

func (o IncidentTodo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentTodo struct {
	value *IncidentTodo
	isSet bool
}

func (v NullableIncidentTodo) Get() *IncidentTodo {
	return v.value
}

func (v *NullableIncidentTodo) Set(val *IncidentTodo) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentTodo) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentTodo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentTodo(val *IncidentTodo) *NullableIncidentTodo {
	return &NullableIncidentTodo{value: val, isSet: true}
}

func (v NullableIncidentTodo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentTodo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}