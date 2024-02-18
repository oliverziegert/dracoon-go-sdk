/*
DRACOON API

REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>

API version: 5.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the EventTypeList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventTypeList{}

// EventTypeList List of event types
type EventTypeList struct {
	// List of event types
	Items []EventType `json:"items"`
}

type _EventTypeList EventTypeList

// NewEventTypeList instantiates a new EventTypeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTypeList(items []EventType) *EventTypeList {
	this := EventTypeList{}
	this.Items = items
	return &this
}

// NewEventTypeListWithDefaults instantiates a new EventTypeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypeListWithDefaults() *EventTypeList {
	this := EventTypeList{}
	return &this
}

// GetItems returns the Items field value
func (o *EventTypeList) GetItems() []EventType {
	if o == nil {
		var ret []EventType
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *EventTypeList) GetItemsOk() ([]EventType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *EventTypeList) SetItems(v []EventType) {
	o.Items = v
}

func (o EventTypeList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventTypeList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *EventTypeList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEventTypeList := _EventTypeList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventTypeList)

	if err != nil {
		return err
	}

	*o = EventTypeList(varEventTypeList)

	return err
}

type NullableEventTypeList struct {
	value *EventTypeList
	isSet bool
}

func (v NullableEventTypeList) Get() *EventTypeList {
	return v.value
}

func (v *NullableEventTypeList) Set(val *EventTypeList) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeList) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeList(val *EventTypeList) *NullableEventTypeList {
	return &NullableEventTypeList{value: val, isSet: true}
}

func (v NullableEventTypeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
