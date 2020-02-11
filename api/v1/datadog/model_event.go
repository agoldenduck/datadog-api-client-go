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

// Event struct for Event
type Event struct {
	// An arbitrary string to use for aggregation. Limited to 100 characters. If you specify a key, all events using that key are grouped together in the Event Stream.
	AggregationKey *string         `json:"aggregation_key,omitempty"`
	AlertType      *EventAlertType `json:"alert_type,omitempty"`
	// POSIX timestamp of the event. Must be sent as an integer (i.e. no quotes). Limited to events no older than 1 year, 24 days (389 days)
	DateHappened *int64 `json:"date_happened,omitempty"`
	// A list of device names to post the event with.
	DeviceName *[]string `json:"device_name,omitempty"`
	// Host name to associate with the event. Any tags associated with the host are also applied to this event.
	Host     *string        `json:"host,omitempty"`
	Id       *int64         `json:"id,omitempty"`
	Payload  *string        `json:"payload,omitempty"`
	Priority *EventPriority `json:"priority,omitempty"`
	// ID of the parent event. Must be sent as an integer (i.e. no quotes).
	RelatedEventId *int64 `json:"related_event_id,omitempty"`
	// The type of event being posted. Options: nagios, hudson, jenkins, my_apps, chef, puppet, git, bitbucket, ... [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value)
	SourceTypeName *string `json:"source_type_name,omitempty"`
	// A list of tags to apply to the event.
	Tags *[]string `json:"tags,omitempty"`
	// The body of the event. Limited to 4000 characters. The text supports markdown. Use msg_text with the Datadog Ruby library
	Text string `json:"text"`
	// The event title. Limited to 100 characters. Use msg_title with the Datadog Ruby library.
	Title string  `json:"title"`
	Url   *string `json:"url,omitempty"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(text string, title string) *Event {
	this := Event{}
	this.Text = text
	this.Title = title
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetAggregationKey returns the AggregationKey field value if set, zero value otherwise.
func (o *Event) GetAggregationKey() string {
	if o == nil || o.AggregationKey == nil {
		var ret string
		return ret
	}
	return *o.AggregationKey
}

// GetAggregationKeyOk returns a tuple with the AggregationKey field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetAggregationKeyOk() (string, bool) {
	if o == nil || o.AggregationKey == nil {
		var ret string
		return ret, false
	}
	return *o.AggregationKey, true
}

// HasAggregationKey returns a boolean if a field has been set.
func (o *Event) HasAggregationKey() bool {
	if o != nil && o.AggregationKey != nil {
		return true
	}

	return false
}

// SetAggregationKey gets a reference to the given string and assigns it to the AggregationKey field.
func (o *Event) SetAggregationKey(v string) {
	o.AggregationKey = &v
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *Event) GetAlertType() EventAlertType {
	if o == nil || o.AlertType == nil {
		var ret EventAlertType
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetAlertTypeOk() (EventAlertType, bool) {
	if o == nil || o.AlertType == nil {
		var ret EventAlertType
		return ret, false
	}
	return *o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *Event) HasAlertType() bool {
	if o != nil && o.AlertType != nil {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given EventAlertType and assigns it to the AlertType field.
func (o *Event) SetAlertType(v EventAlertType) {
	o.AlertType = &v
}

// GetDateHappened returns the DateHappened field value if set, zero value otherwise.
func (o *Event) GetDateHappened() int64 {
	if o == nil || o.DateHappened == nil {
		var ret int64
		return ret
	}
	return *o.DateHappened
}

// GetDateHappenedOk returns a tuple with the DateHappened field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetDateHappenedOk() (int64, bool) {
	if o == nil || o.DateHappened == nil {
		var ret int64
		return ret, false
	}
	return *o.DateHappened, true
}

// HasDateHappened returns a boolean if a field has been set.
func (o *Event) HasDateHappened() bool {
	if o != nil && o.DateHappened != nil {
		return true
	}

	return false
}

// SetDateHappened gets a reference to the given int64 and assigns it to the DateHappened field.
func (o *Event) SetDateHappened(v int64) {
	o.DateHappened = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *Event) GetDeviceName() []string {
	if o == nil || o.DeviceName == nil {
		var ret []string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetDeviceNameOk() ([]string, bool) {
	if o == nil || o.DeviceName == nil {
		var ret []string
		return ret, false
	}
	return *o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *Event) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given []string and assigns it to the DeviceName field.
func (o *Event) SetDeviceName(v []string) {
	o.DeviceName = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *Event) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetHostOk() (string, bool) {
	if o == nil || o.Host == nil {
		var ret string
		return ret, false
	}
	return *o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Event) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *Event) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Event) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetIdOk() (int64, bool) {
	if o == nil || o.Id == nil {
		var ret int64
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Event) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Event) SetId(v int64) {
	o.Id = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *Event) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetPayloadOk() (string, bool) {
	if o == nil || o.Payload == nil {
		var ret string
		return ret, false
	}
	return *o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *Event) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *Event) SetPayload(v string) {
	o.Payload = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *Event) GetPriority() EventPriority {
	if o == nil || o.Priority == nil {
		var ret EventPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetPriorityOk() (EventPriority, bool) {
	if o == nil || o.Priority == nil {
		var ret EventPriority
		return ret, false
	}
	return *o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *Event) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given EventPriority and assigns it to the Priority field.
func (o *Event) SetPriority(v EventPriority) {
	o.Priority = &v
}

// GetRelatedEventId returns the RelatedEventId field value if set, zero value otherwise.
func (o *Event) GetRelatedEventId() int64 {
	if o == nil || o.RelatedEventId == nil {
		var ret int64
		return ret
	}
	return *o.RelatedEventId
}

// GetRelatedEventIdOk returns a tuple with the RelatedEventId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetRelatedEventIdOk() (int64, bool) {
	if o == nil || o.RelatedEventId == nil {
		var ret int64
		return ret, false
	}
	return *o.RelatedEventId, true
}

// HasRelatedEventId returns a boolean if a field has been set.
func (o *Event) HasRelatedEventId() bool {
	if o != nil && o.RelatedEventId != nil {
		return true
	}

	return false
}

// SetRelatedEventId gets a reference to the given int64 and assigns it to the RelatedEventId field.
func (o *Event) SetRelatedEventId(v int64) {
	o.RelatedEventId = &v
}

// GetSourceTypeName returns the SourceTypeName field value if set, zero value otherwise.
func (o *Event) GetSourceTypeName() string {
	if o == nil || o.SourceTypeName == nil {
		var ret string
		return ret
	}
	return *o.SourceTypeName
}

// GetSourceTypeNameOk returns a tuple with the SourceTypeName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSourceTypeNameOk() (string, bool) {
	if o == nil || o.SourceTypeName == nil {
		var ret string
		return ret, false
	}
	return *o.SourceTypeName, true
}

// HasSourceTypeName returns a boolean if a field has been set.
func (o *Event) HasSourceTypeName() bool {
	if o != nil && o.SourceTypeName != nil {
		return true
	}

	return false
}

// SetSourceTypeName gets a reference to the given string and assigns it to the SourceTypeName field.
func (o *Event) SetSourceTypeName(v string) {
	o.SourceTypeName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Event) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Event) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Event) SetTags(v []string) {
	o.Tags = &v
}

// GetText returns the Text field value
func (o *Event) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// SetText sets field value
func (o *Event) SetText(v string) {
	o.Text = v
}

// GetTitle returns the Title field value
func (o *Event) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *Event) SetTitle(v string) {
	o.Title = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Event) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetUrlOk() (string, bool) {
	if o == nil || o.Url == nil {
		var ret string
		return ret, false
	}
	return *o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Event) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Event) SetUrl(v string) {
	o.Url = &v
}

type NullableEvent struct {
	Value        Event
	ExplicitNull bool
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
