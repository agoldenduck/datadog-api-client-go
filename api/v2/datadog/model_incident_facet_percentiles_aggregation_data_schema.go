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

// IncidentFacetPercentilesAggregationDataSchema A `percentiles` aggregation for facets.
type IncidentFacetPercentilesAggregationDataSchema struct {
	// The 1% distrubtion value.
	P1 NullableFloat64 `json:"p1,omitempty"`
	// The 25% distrubtion value.
	P25 NullableFloat64 `json:"p25,omitempty"`
	// The 5% distrubtion value.
	P5 NullableFloat64 `json:"p5,omitempty"`
	// The 50% distrubtion value.
	P50 NullableFloat64 `json:"p50,omitempty"`
	// The 75% distrubtion value.
	P75 NullableFloat64 `json:"p75,omitempty"`
	// The 95% distrubtion value.
	P95 NullableFloat64 `json:"p95,omitempty"`
	// The 99% distrubtion value.
	P99 NullableFloat64 `json:"p99,omitempty"`
}

// NewIncidentFacetPercentilesAggregationDataSchema instantiates a new IncidentFacetPercentilesAggregationDataSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentFacetPercentilesAggregationDataSchema() *IncidentFacetPercentilesAggregationDataSchema {
	this := IncidentFacetPercentilesAggregationDataSchema{}
	return &this
}

// NewIncidentFacetPercentilesAggregationDataSchemaWithDefaults instantiates a new IncidentFacetPercentilesAggregationDataSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentFacetPercentilesAggregationDataSchemaWithDefaults() *IncidentFacetPercentilesAggregationDataSchema {
	this := IncidentFacetPercentilesAggregationDataSchema{}
	return &this
}

// GetP1 returns the P1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP1() float64 {
	if o == nil || o.P1.Get() == nil {
		var ret float64
		return ret
	}
	return *o.P1.Get()
}

// GetP1Ok returns a tuple with the P1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP1Ok() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.P1.Get(), o.P1.IsSet()
}

// HasP1 returns a boolean if a field has been set.
func (o *IncidentFacetPercentilesAggregationDataSchema) HasP1() bool {
	if o != nil && o.P1.IsSet() {
		return true
	}

	return false
}

// SetP1 gets a reference to the given NullableFloat64 and assigns it to the P1 field.
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP1(v float64) {
	o.P1.Set(&v)
}

// SetP1Nil sets the value for P1 to be an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP1Nil() {
	o.P1.Set(nil)
}

// UnsetP1 ensures that no value is present for P1, not even an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) UnsetP1() {
	o.P1.Unset()
}

// GetP25 returns the P25 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP25() float64 {
	if o == nil || o.P25.Get() == nil {
		var ret float64
		return ret
	}
	return *o.P25.Get()
}

// GetP25Ok returns a tuple with the P25 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP25Ok() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.P25.Get(), o.P25.IsSet()
}

// HasP25 returns a boolean if a field has been set.
func (o *IncidentFacetPercentilesAggregationDataSchema) HasP25() bool {
	if o != nil && o.P25.IsSet() {
		return true
	}

	return false
}

// SetP25 gets a reference to the given NullableFloat64 and assigns it to the P25 field.
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP25(v float64) {
	o.P25.Set(&v)
}

// SetP25Nil sets the value for P25 to be an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP25Nil() {
	o.P25.Set(nil)
}

// UnsetP25 ensures that no value is present for P25, not even an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) UnsetP25() {
	o.P25.Unset()
}

// GetP5 returns the P5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP5() float64 {
	if o == nil || o.P5.Get() == nil {
		var ret float64
		return ret
	}
	return *o.P5.Get()
}

// GetP5Ok returns a tuple with the P5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP5Ok() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.P5.Get(), o.P5.IsSet()
}

// HasP5 returns a boolean if a field has been set.
func (o *IncidentFacetPercentilesAggregationDataSchema) HasP5() bool {
	if o != nil && o.P5.IsSet() {
		return true
	}

	return false
}

// SetP5 gets a reference to the given NullableFloat64 and assigns it to the P5 field.
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP5(v float64) {
	o.P5.Set(&v)
}

// SetP5Nil sets the value for P5 to be an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP5Nil() {
	o.P5.Set(nil)
}

// UnsetP5 ensures that no value is present for P5, not even an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) UnsetP5() {
	o.P5.Unset()
}

// GetP50 returns the P50 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP50() float64 {
	if o == nil || o.P50.Get() == nil {
		var ret float64
		return ret
	}
	return *o.P50.Get()
}

// GetP50Ok returns a tuple with the P50 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP50Ok() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.P50.Get(), o.P50.IsSet()
}

// HasP50 returns a boolean if a field has been set.
func (o *IncidentFacetPercentilesAggregationDataSchema) HasP50() bool {
	if o != nil && o.P50.IsSet() {
		return true
	}

	return false
}

// SetP50 gets a reference to the given NullableFloat64 and assigns it to the P50 field.
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP50(v float64) {
	o.P50.Set(&v)
}

// SetP50Nil sets the value for P50 to be an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP50Nil() {
	o.P50.Set(nil)
}

// UnsetP50 ensures that no value is present for P50, not even an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) UnsetP50() {
	o.P50.Unset()
}

// GetP75 returns the P75 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP75() float64 {
	if o == nil || o.P75.Get() == nil {
		var ret float64
		return ret
	}
	return *o.P75.Get()
}

// GetP75Ok returns a tuple with the P75 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP75Ok() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.P75.Get(), o.P75.IsSet()
}

// HasP75 returns a boolean if a field has been set.
func (o *IncidentFacetPercentilesAggregationDataSchema) HasP75() bool {
	if o != nil && o.P75.IsSet() {
		return true
	}

	return false
}

// SetP75 gets a reference to the given NullableFloat64 and assigns it to the P75 field.
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP75(v float64) {
	o.P75.Set(&v)
}

// SetP75Nil sets the value for P75 to be an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP75Nil() {
	o.P75.Set(nil)
}

// UnsetP75 ensures that no value is present for P75, not even an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) UnsetP75() {
	o.P75.Unset()
}

// GetP95 returns the P95 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP95() float64 {
	if o == nil || o.P95.Get() == nil {
		var ret float64
		return ret
	}
	return *o.P95.Get()
}

// GetP95Ok returns a tuple with the P95 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP95Ok() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.P95.Get(), o.P95.IsSet()
}

// HasP95 returns a boolean if a field has been set.
func (o *IncidentFacetPercentilesAggregationDataSchema) HasP95() bool {
	if o != nil && o.P95.IsSet() {
		return true
	}

	return false
}

// SetP95 gets a reference to the given NullableFloat64 and assigns it to the P95 field.
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP95(v float64) {
	o.P95.Set(&v)
}

// SetP95Nil sets the value for P95 to be an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP95Nil() {
	o.P95.Set(nil)
}

// UnsetP95 ensures that no value is present for P95, not even an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) UnsetP95() {
	o.P95.Unset()
}

// GetP99 returns the P99 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP99() float64 {
	if o == nil || o.P99.Get() == nil {
		var ret float64
		return ret
	}
	return *o.P99.Get()
}

// GetP99Ok returns a tuple with the P99 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentFacetPercentilesAggregationDataSchema) GetP99Ok() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.P99.Get(), o.P99.IsSet()
}

// HasP99 returns a boolean if a field has been set.
func (o *IncidentFacetPercentilesAggregationDataSchema) HasP99() bool {
	if o != nil && o.P99.IsSet() {
		return true
	}

	return false
}

// SetP99 gets a reference to the given NullableFloat64 and assigns it to the P99 field.
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP99(v float64) {
	o.P99.Set(&v)
}

// SetP99Nil sets the value for P99 to be an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) SetP99Nil() {
	o.P99.Set(nil)
}

// UnsetP99 ensures that no value is present for P99, not even an explicit nil
func (o *IncidentFacetPercentilesAggregationDataSchema) UnsetP99() {
	o.P99.Unset()
}

func (o IncidentFacetPercentilesAggregationDataSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.P1.IsSet() {
		toSerialize["p1"] = o.P1.Get()
	}
	if o.P25.IsSet() {
		toSerialize["p25"] = o.P25.Get()
	}
	if o.P5.IsSet() {
		toSerialize["p5"] = o.P5.Get()
	}
	if o.P50.IsSet() {
		toSerialize["p50"] = o.P50.Get()
	}
	if o.P75.IsSet() {
		toSerialize["p75"] = o.P75.Get()
	}
	if o.P95.IsSet() {
		toSerialize["p95"] = o.P95.Get()
	}
	if o.P99.IsSet() {
		toSerialize["p99"] = o.P99.Get()
	}
	return json.Marshal(toSerialize)
}

// AsIncidentFacetSchemaDataOneOf wraps this instance of IncidentFacetPercentilesAggregationDataSchema in IncidentFacetSchemaDataOneOf
func (s *IncidentFacetPercentilesAggregationDataSchema) AsIncidentFacetSchemaDataOneOf() IncidentFacetSchemaDataOneOf {
	return IncidentFacetSchemaDataOneOf{IncidentFacetSchemaDataOneOfInterface: s}
}

type NullableIncidentFacetPercentilesAggregationDataSchema struct {
	value *IncidentFacetPercentilesAggregationDataSchema
	isSet bool
}

func (v NullableIncidentFacetPercentilesAggregationDataSchema) Get() *IncidentFacetPercentilesAggregationDataSchema {
	return v.value
}

func (v *NullableIncidentFacetPercentilesAggregationDataSchema) Set(val *IncidentFacetPercentilesAggregationDataSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentFacetPercentilesAggregationDataSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentFacetPercentilesAggregationDataSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentFacetPercentilesAggregationDataSchema(val *IncidentFacetPercentilesAggregationDataSchema) *NullableIncidentFacetPercentilesAggregationDataSchema {
	return &NullableIncidentFacetPercentilesAggregationDataSchema{value: val, isSet: true}
}

func (v NullableIncidentFacetPercentilesAggregationDataSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentFacetPercentilesAggregationDataSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}