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

// checks if the KeyValueEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyValueEntry{}

// KeyValueEntry Key-value pair
type KeyValueEntry struct {
	// Entry key
	Key string `json:"key"`
	// Entry value
	Value string `json:"value"`
}

type _KeyValueEntry KeyValueEntry

// NewKeyValueEntry instantiates a new KeyValueEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyValueEntry(key string, value string) *KeyValueEntry {
	this := KeyValueEntry{}
	this.Key = key
	this.Value = value
	return &this
}

// NewKeyValueEntryWithDefaults instantiates a new KeyValueEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyValueEntryWithDefaults() *KeyValueEntry {
	this := KeyValueEntry{}
	return &this
}

// GetKey returns the Key field value
func (o *KeyValueEntry) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *KeyValueEntry) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *KeyValueEntry) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *KeyValueEntry) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *KeyValueEntry) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *KeyValueEntry) SetValue(v string) {
	o.Value = v
}

func (o KeyValueEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyValueEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *KeyValueEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
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

	varKeyValueEntry := _KeyValueEntry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKeyValueEntry)

	if err != nil {
		return err
	}

	*o = KeyValueEntry(varKeyValueEntry)

	return err
}

type NullableKeyValueEntry struct {
	value *KeyValueEntry
	isSet bool
}

func (v NullableKeyValueEntry) Get() *KeyValueEntry {
	return v.value
}

func (v *NullableKeyValueEntry) Set(val *KeyValueEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyValueEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyValueEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyValueEntry(val *KeyValueEntry) *NullableKeyValueEntry {
	return &NullableKeyValueEntry{value: val, isSet: true}
}

func (v NullableKeyValueEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyValueEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
