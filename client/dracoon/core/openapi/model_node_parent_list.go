/*
DRACOON API

REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>

API version: 5.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the NodeParentList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeParentList{}

// NodeParentList List of parent nodes
type NodeParentList struct {
	// List of node parents
	Items []NodeParent `json:"items,omitempty"`
}

// NewNodeParentList instantiates a new NodeParentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeParentList() *NodeParentList {
	this := NodeParentList{}
	return &this
}

// NewNodeParentListWithDefaults instantiates a new NodeParentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeParentListWithDefaults() *NodeParentList {
	this := NodeParentList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *NodeParentList) GetItems() []NodeParent {
	if o == nil || IsNil(o.Items) {
		var ret []NodeParent
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeParentList) GetItemsOk() ([]NodeParent, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *NodeParentList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []NodeParent and assigns it to the Items field.
func (o *NodeParentList) SetItems(v []NodeParent) {
	o.Items = v
}

func (o NodeParentList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeParentList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableNodeParentList struct {
	value *NodeParentList
	isSet bool
}

func (v NullableNodeParentList) Get() *NodeParentList {
	return v.value
}

func (v *NullableNodeParentList) Set(val *NodeParentList) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeParentList) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeParentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeParentList(val *NodeParentList) *NullableNodeParentList {
	return &NullableNodeParentList{value: val, isSet: true}
}

func (v NullableNodeParentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeParentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
