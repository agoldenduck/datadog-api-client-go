/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SyntheticsAssertionOperator Assertion operator to apply.
type SyntheticsAssertionOperator string

// List of SyntheticsAssertionOperator
const (
	SYNTHETICSASSERTIONOPERATOR_CONTAINS             SyntheticsAssertionOperator = "contains"
	SYNTHETICSASSERTIONOPERATOR_DOES_NOT_CONTAIN     SyntheticsAssertionOperator = "doesNotContain"
	SYNTHETICSASSERTIONOPERATOR_IS                   SyntheticsAssertionOperator = "is"
	SYNTHETICSASSERTIONOPERATOR_IS_NOT               SyntheticsAssertionOperator = "isNot"
	SYNTHETICSASSERTIONOPERATOR_LESS_THAN            SyntheticsAssertionOperator = "lessThan"
	SYNTHETICSASSERTIONOPERATOR_MATCHES              SyntheticsAssertionOperator = "matches"
	SYNTHETICSASSERTIONOPERATOR_DOES_NOT_MATCH       SyntheticsAssertionOperator = "doesNotMatch"
	SYNTHETICSASSERTIONOPERATOR_VALIDATES            SyntheticsAssertionOperator = "validates"
	SYNTHETICSASSERTIONOPERATOR_IS_IN_MORE_DAYS_THAN SyntheticsAssertionOperator = "isInMoreThan"
)

func (v *SyntheticsAssertionOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SyntheticsAssertionOperator(value)
	for _, existing := range []SyntheticsAssertionOperator{"contains", "doesNotContain", "is", "isNot", "lessThan", "matches", "doesNotMatch", "validates", "isInMoreThan"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SyntheticsAssertionOperator", *v)
}

// Ptr returns reference to SyntheticsAssertionOperator value
func (v SyntheticsAssertionOperator) Ptr() *SyntheticsAssertionOperator {
	return &v
}

type NullableSyntheticsAssertionOperator struct {
	value *SyntheticsAssertionOperator
	isSet bool
}

func (v NullableSyntheticsAssertionOperator) Get() *SyntheticsAssertionOperator {
	return v.value
}

func (v *NullableSyntheticsAssertionOperator) Set(val *SyntheticsAssertionOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsAssertionOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsAssertionOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsAssertionOperator(val *SyntheticsAssertionOperator) *NullableSyntheticsAssertionOperator {
	return &NullableSyntheticsAssertionOperator{value: val, isSet: true}
}

func (v NullableSyntheticsAssertionOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsAssertionOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
