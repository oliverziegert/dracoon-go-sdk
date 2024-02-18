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

// checks if the SubscribedNodeList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscribedNodeList{}

// SubscribedNodeList List of subscribed nodes
type SubscribedNodeList struct {
	Range Range `json:"range"`
	// List of subscribed nodes
	Items []SubscribedNode `json:"items"`
}

type _SubscribedNodeList SubscribedNodeList

// NewSubscribedNodeList instantiates a new SubscribedNodeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribedNodeList(range_ Range, items []SubscribedNode) *SubscribedNodeList {
	this := SubscribedNodeList{}
	this.Range = range_
	this.Items = items
	return &this
}

// NewSubscribedNodeListWithDefaults instantiates a new SubscribedNodeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribedNodeListWithDefaults() *SubscribedNodeList {
	this := SubscribedNodeList{}
	return &this
}

// GetRange returns the Range field value
func (o *SubscribedNodeList) GetRange() Range {
	if o == nil {
		var ret Range
		return ret
	}

	return o.Range
}

// GetRangeOk returns a tuple with the Range field value
// and a boolean to check if the value has been set.
func (o *SubscribedNodeList) GetRangeOk() (*Range, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Range, true
}

// SetRange sets field value
func (o *SubscribedNodeList) SetRange(v Range) {
	o.Range = v
}

// GetItems returns the Items field value
func (o *SubscribedNodeList) GetItems() []SubscribedNode {
	if o == nil {
		var ret []SubscribedNode
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SubscribedNodeList) GetItemsOk() ([]SubscribedNode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SubscribedNodeList) SetItems(v []SubscribedNode) {
	o.Items = v
}

func (o SubscribedNodeList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribedNodeList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["range"] = o.Range
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *SubscribedNodeList) UnmarshalJSON(data []byte) (err error) {
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

	varSubscribedNodeList := _SubscribedNodeList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscribedNodeList)

	if err != nil {
		return err
	}

	*o = SubscribedNodeList(varSubscribedNodeList)

	return err
}

type NullableSubscribedNodeList struct {
	value *SubscribedNodeList
	isSet bool
}

func (v NullableSubscribedNodeList) Get() *SubscribedNodeList {
	return v.value
}

func (v *NullableSubscribedNodeList) Set(val *SubscribedNodeList) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribedNodeList) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribedNodeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribedNodeList(val *SubscribedNodeList) *NullableSubscribedNodeList {
	return &NullableSubscribedNodeList{value: val, isSet: true}
}

func (v NullableSubscribedNodeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribedNodeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
