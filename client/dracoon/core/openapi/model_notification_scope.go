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

// checks if the NotificationScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationScope{}

// NotificationScope Notification scope information
type NotificationScope struct {
	// Scope ID
	Id int32 `json:"id"`
	// Name
	Name string `json:"name"`
}

type _NotificationScope NotificationScope

// NewNotificationScope instantiates a new NotificationScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationScope(id int32, name string) *NotificationScope {
	this := NotificationScope{}
	this.Id = id
	this.Name = name
	return &this
}

// NewNotificationScopeWithDefaults instantiates a new NotificationScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationScopeWithDefaults() *NotificationScope {
	this := NotificationScope{}
	return &this
}

// GetId returns the Id field value
func (o *NotificationScope) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NotificationScope) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NotificationScope) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *NotificationScope) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotificationScope) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NotificationScope) SetName(v string) {
	o.Name = v
}

func (o NotificationScope) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *NotificationScope) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varNotificationScope := _NotificationScope{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationScope)

	if err != nil {
		return err
	}

	*o = NotificationScope(varNotificationScope)

	return err
}

type NullableNotificationScope struct {
	value *NotificationScope
	isSet bool
}

func (v NullableNotificationScope) Get() *NotificationScope {
	return v.value
}

func (v *NullableNotificationScope) Set(val *NotificationScope) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationScope) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationScope(val *NotificationScope) *NullableNotificationScope {
	return &NullableNotificationScope{value: val, isSet: true}
}

func (v NullableNotificationScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
