/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// AwsLogsListServicesResponse struct for AwsLogsListServicesResponse
type AwsLogsListServicesResponse struct {
	// Key value in returned object.
	Id *string `json:"id,omitempty"`
	// Name of service available for configuration with Datadog logs.
	Label *string `json:"label,omitempty"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AwsLogsListServicesResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AwsLogsListServicesResponse) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AwsLogsListServicesResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AwsLogsListServicesResponse) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AwsLogsListServicesResponse) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AwsLogsListServicesResponse) GetLabelOk() (string, bool) {
	if o == nil || o.Label == nil {
		var ret string
		return ret, false
	}
	return *o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AwsLogsListServicesResponse) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AwsLogsListServicesResponse) SetLabel(v string) {
	o.Label = &v
}

type NullableAwsLogsListServicesResponse struct {
	Value        AwsLogsListServicesResponse
	ExplicitNull bool
}

func (v NullableAwsLogsListServicesResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAwsLogsListServicesResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}