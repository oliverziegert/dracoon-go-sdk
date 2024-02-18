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

// checks if the NotificationChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationChannel{}

// NotificationChannel Notification channel information
type NotificationChannel struct {
	// Channel ID
	Id int32 `json:"id"`
	// Name
	Name string `json:"name"`
	// Determines whether channel is enabled
	IsEnabled bool `json:"isEnabled"`
	// Channel type (only `EMAIL` available at the moment)
	Type string `json:"type"`
	// Channel frequency (aggregation window size in minutes)
	Frequency int64 `json:"frequency"`
}

type _NotificationChannel NotificationChannel

// NewNotificationChannel instantiates a new NotificationChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationChannel(id int32, name string, isEnabled bool, type_ string, frequency int64) *NotificationChannel {
	this := NotificationChannel{}
	this.Id = id
	this.Name = name
	this.IsEnabled = isEnabled
	this.Type = type_
	this.Frequency = frequency
	return &this
}

// NewNotificationChannelWithDefaults instantiates a new NotificationChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationChannelWithDefaults() *NotificationChannel {
	this := NotificationChannel{}
	return &this
}

// GetId returns the Id field value
func (o *NotificationChannel) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NotificationChannel) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NotificationChannel) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *NotificationChannel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotificationChannel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NotificationChannel) SetName(v string) {
	o.Name = v
}

// GetIsEnabled returns the IsEnabled field value
func (o *NotificationChannel) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *NotificationChannel) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *NotificationChannel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetType returns the Type field value
func (o *NotificationChannel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationChannel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificationChannel) SetType(v string) {
	o.Type = v
}

// GetFrequency returns the Frequency field value
func (o *NotificationChannel) GetFrequency() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *NotificationChannel) GetFrequencyOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *NotificationChannel) SetFrequency(v int64) {
	o.Frequency = v
}

func (o NotificationChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["isEnabled"] = o.IsEnabled
	toSerialize["type"] = o.Type
	toSerialize["frequency"] = o.Frequency
	return toSerialize, nil
}

func (o *NotificationChannel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"isEnabled",
		"type",
		"frequency",
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

	varNotificationChannel := _NotificationChannel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationChannel)

	if err != nil {
		return err
	}

	*o = NotificationChannel(varNotificationChannel)

	return err
}

type NullableNotificationChannel struct {
	value *NotificationChannel
	isSet bool
}

func (v NullableNotificationChannel) Get() *NotificationChannel {
	return v.value
}

func (v *NullableNotificationChannel) Set(val *NotificationChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationChannel(val *NotificationChannel) *NullableNotificationChannel {
	return &NullableNotificationChannel{value: val, isSet: true}
}

func (v NullableNotificationChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
