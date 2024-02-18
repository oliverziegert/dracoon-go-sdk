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

// checks if the LastAdminGroupRoomList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LastAdminGroupRoomList{}

// LastAdminGroupRoomList List of (last admin group) rooms
type LastAdminGroupRoomList struct {
	// List of last admin rooms
	Items []LastAdminGroupRoom `json:"items"`
}

type _LastAdminGroupRoomList LastAdminGroupRoomList

// NewLastAdminGroupRoomList instantiates a new LastAdminGroupRoomList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastAdminGroupRoomList(items []LastAdminGroupRoom) *LastAdminGroupRoomList {
	this := LastAdminGroupRoomList{}
	this.Items = items
	return &this
}

// NewLastAdminGroupRoomListWithDefaults instantiates a new LastAdminGroupRoomList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastAdminGroupRoomListWithDefaults() *LastAdminGroupRoomList {
	this := LastAdminGroupRoomList{}
	return &this
}

// GetItems returns the Items field value
func (o *LastAdminGroupRoomList) GetItems() []LastAdminGroupRoom {
	if o == nil {
		var ret []LastAdminGroupRoom
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *LastAdminGroupRoomList) GetItemsOk() ([]LastAdminGroupRoom, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *LastAdminGroupRoomList) SetItems(v []LastAdminGroupRoom) {
	o.Items = v
}

func (o LastAdminGroupRoomList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LastAdminGroupRoomList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *LastAdminGroupRoomList) UnmarshalJSON(data []byte) (err error) {
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

	varLastAdminGroupRoomList := _LastAdminGroupRoomList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLastAdminGroupRoomList)

	if err != nil {
		return err
	}

	*o = LastAdminGroupRoomList(varLastAdminGroupRoomList)

	return err
}

type NullableLastAdminGroupRoomList struct {
	value *LastAdminGroupRoomList
	isSet bool
}

func (v NullableLastAdminGroupRoomList) Get() *LastAdminGroupRoomList {
	return v.value
}

func (v *NullableLastAdminGroupRoomList) Set(val *LastAdminGroupRoomList) {
	v.value = val
	v.isSet = true
}

func (v NullableLastAdminGroupRoomList) IsSet() bool {
	return v.isSet
}

func (v *NullableLastAdminGroupRoomList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastAdminGroupRoomList(val *LastAdminGroupRoomList) *NullableLastAdminGroupRoomList {
	return &NullableLastAdminGroupRoomList{value: val, isSet: true}
}

func (v NullableLastAdminGroupRoomList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastAdminGroupRoomList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
