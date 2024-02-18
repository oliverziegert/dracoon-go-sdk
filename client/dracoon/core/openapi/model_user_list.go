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

// checks if the UserList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserList{}

// UserList List of users
type UserList struct {
	Range Range `json:"range"`
	// List of users
	Items []UserItem `json:"items"`
}

type _UserList UserList

// NewUserList instantiates a new UserList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserList(range_ Range, items []UserItem) *UserList {
	this := UserList{}
	this.Range = range_
	this.Items = items
	return &this
}

// NewUserListWithDefaults instantiates a new UserList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserListWithDefaults() *UserList {
	this := UserList{}
	return &this
}

// GetRange returns the Range field value
func (o *UserList) GetRange() Range {
	if o == nil {
		var ret Range
		return ret
	}

	return o.Range
}

// GetRangeOk returns a tuple with the Range field value
// and a boolean to check if the value has been set.
func (o *UserList) GetRangeOk() (*Range, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Range, true
}

// SetRange sets field value
func (o *UserList) SetRange(v Range) {
	o.Range = v
}

// GetItems returns the Items field value
func (o *UserList) GetItems() []UserItem {
	if o == nil {
		var ret []UserItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *UserList) GetItemsOk() ([]UserItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *UserList) SetItems(v []UserItem) {
	o.Items = v
}

func (o UserList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["range"] = o.Range
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *UserList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"range",
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

	varUserList := _UserList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserList)

	if err != nil {
		return err
	}

	*o = UserList(varUserList)

	return err
}

type NullableUserList struct {
	value *UserList
	isSet bool
}

func (v NullableUserList) Get() *UserList {
	return v.value
}

func (v *NullableUserList) Set(val *UserList) {
	v.value = val
	v.isSet = true
}

func (v NullableUserList) IsSet() bool {
	return v.isSet
}

func (v *NullableUserList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserList(val *UserList) *NullableUserList {
	return &NullableUserList{value: val, isSet: true}
}

func (v NullableUserList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
