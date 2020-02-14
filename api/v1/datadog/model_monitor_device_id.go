/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// MonitorDeviceId the model 'MonitorDeviceId'
type MonitorDeviceId string

// List of MonitorDeviceID
const (
	MONITORDEVICEID_LAPTOP_LARGE MonitorDeviceId = "laptop_large"
	MONITORDEVICEID_TABLET       MonitorDeviceId = "tablet"
	MONITORDEVICEID_MOBILE_SMALL MonitorDeviceId = "mobile_small"
)

// Ptr returns reference to MonitorDeviceID value
func (v MonitorDeviceId) Ptr() *MonitorDeviceId {
	return &v
}

type NullableMonitorDeviceId struct {
	Value        MonitorDeviceId
	ExplicitNull bool
}

func (v NullableMonitorDeviceId) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableMonitorDeviceId) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}