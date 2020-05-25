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

import (
	"fmt"
)

// WidgetServiceSummaryDisplayFormat Number of columns to display.
type WidgetServiceSummaryDisplayFormat string

// List of WidgetServiceSummaryDisplayFormat
const (
	WIDGETSERVICESUMMARYDISPLAYFORMAT_ONE_COLUMN WidgetServiceSummaryDisplayFormat = "one_column"
	WIDGETSERVICESUMMARYDISPLAYFORMAT_TWO_COLUMN WidgetServiceSummaryDisplayFormat = "two_column"
)

func (v *WidgetServiceSummaryDisplayFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WidgetServiceSummaryDisplayFormat(value)
	for _, existing := range []WidgetServiceSummaryDisplayFormat{"one_column", "two_column"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WidgetServiceSummaryDisplayFormat", *v)
}

// Ptr returns reference to WidgetServiceSummaryDisplayFormat value
func (v WidgetServiceSummaryDisplayFormat) Ptr() *WidgetServiceSummaryDisplayFormat {
	return &v
}

type NullableWidgetServiceSummaryDisplayFormat struct {
	value *WidgetServiceSummaryDisplayFormat
	isSet bool
}

func (v NullableWidgetServiceSummaryDisplayFormat) Get() *WidgetServiceSummaryDisplayFormat {
	return v.value
}

func (v *NullableWidgetServiceSummaryDisplayFormat) Set(val *WidgetServiceSummaryDisplayFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetServiceSummaryDisplayFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetServiceSummaryDisplayFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetServiceSummaryDisplayFormat(val *WidgetServiceSummaryDisplayFormat) *NullableWidgetServiceSummaryDisplayFormat {
	return &NullableWidgetServiceSummaryDisplayFormat{value: val, isSet: true}
}

func (v NullableWidgetServiceSummaryDisplayFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetServiceSummaryDisplayFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
