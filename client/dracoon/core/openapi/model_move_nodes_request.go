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

// checks if the MoveNodesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoveNodesRequest{}

// MoveNodesRequest Request model for moving nodes
type MoveNodesRequest struct {
	// List of nodes to be moved
	Items []MoveNode `json:"items,omitempty"`
	// Node conflict resolution strategy:  * `autorename`  * `overwrite`  * `fail`
	ResolutionStrategy *string `json:"resolutionStrategy,omitempty"`
	// Preserve Download Share Links and point them to the new node.
	KeepShareLinks *bool `json:"keepShareLinks,omitempty"`
	// &#128679; Deprecated since v4.5.0  Node IDs  Please use `items` instead.
	// Deprecated
	NodeIds []int64 `json:"nodeIds,omitempty"`
}

// NewMoveNodesRequest instantiates a new MoveNodesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveNodesRequest() *MoveNodesRequest {
	this := MoveNodesRequest{}
	var resolutionStrategy string = "autorename"
	this.ResolutionStrategy = &resolutionStrategy
	var keepShareLinks bool = false
	this.KeepShareLinks = &keepShareLinks
	return &this
}

// NewMoveNodesRequestWithDefaults instantiates a new MoveNodesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveNodesRequestWithDefaults() *MoveNodesRequest {
	this := MoveNodesRequest{}
	var resolutionStrategy string = "autorename"
	this.ResolutionStrategy = &resolutionStrategy
	var keepShareLinks bool = false
	this.KeepShareLinks = &keepShareLinks
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *MoveNodesRequest) GetItems() []MoveNode {
	if o == nil || IsNil(o.Items) {
		var ret []MoveNode
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveNodesRequest) GetItemsOk() ([]MoveNode, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MoveNodesRequest) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MoveNode and assigns it to the Items field.
func (o *MoveNodesRequest) SetItems(v []MoveNode) {
	o.Items = v
}

// GetResolutionStrategy returns the ResolutionStrategy field value if set, zero value otherwise.
func (o *MoveNodesRequest) GetResolutionStrategy() string {
	if o == nil || IsNil(o.ResolutionStrategy) {
		var ret string
		return ret
	}
	return *o.ResolutionStrategy
}

// GetResolutionStrategyOk returns a tuple with the ResolutionStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveNodesRequest) GetResolutionStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.ResolutionStrategy) {
		return nil, false
	}
	return o.ResolutionStrategy, true
}

// HasResolutionStrategy returns a boolean if a field has been set.
func (o *MoveNodesRequest) HasResolutionStrategy() bool {
	if o != nil && !IsNil(o.ResolutionStrategy) {
		return true
	}

	return false
}

// SetResolutionStrategy gets a reference to the given string and assigns it to the ResolutionStrategy field.
func (o *MoveNodesRequest) SetResolutionStrategy(v string) {
	o.ResolutionStrategy = &v
}

// GetKeepShareLinks returns the KeepShareLinks field value if set, zero value otherwise.
func (o *MoveNodesRequest) GetKeepShareLinks() bool {
	if o == nil || IsNil(o.KeepShareLinks) {
		var ret bool
		return ret
	}
	return *o.KeepShareLinks
}

// GetKeepShareLinksOk returns a tuple with the KeepShareLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveNodesRequest) GetKeepShareLinksOk() (*bool, bool) {
	if o == nil || IsNil(o.KeepShareLinks) {
		return nil, false
	}
	return o.KeepShareLinks, true
}

// HasKeepShareLinks returns a boolean if a field has been set.
func (o *MoveNodesRequest) HasKeepShareLinks() bool {
	if o != nil && !IsNil(o.KeepShareLinks) {
		return true
	}

	return false
}

// SetKeepShareLinks gets a reference to the given bool and assigns it to the KeepShareLinks field.
func (o *MoveNodesRequest) SetKeepShareLinks(v bool) {
	o.KeepShareLinks = &v
}

// GetNodeIds returns the NodeIds field value if set, zero value otherwise.
// Deprecated
func (o *MoveNodesRequest) GetNodeIds() []int64 {
	if o == nil || IsNil(o.NodeIds) {
		var ret []int64
		return ret
	}
	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *MoveNodesRequest) GetNodeIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.NodeIds) {
		return nil, false
	}
	return o.NodeIds, true
}

// HasNodeIds returns a boolean if a field has been set.
func (o *MoveNodesRequest) HasNodeIds() bool {
	if o != nil && !IsNil(o.NodeIds) {
		return true
	}

	return false
}

// SetNodeIds gets a reference to the given []int64 and assigns it to the NodeIds field.
// Deprecated
func (o *MoveNodesRequest) SetNodeIds(v []int64) {
	o.NodeIds = v
}

func (o MoveNodesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoveNodesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResolutionStrategy) {
		toSerialize["resolutionStrategy"] = o.ResolutionStrategy
	}
	if !IsNil(o.KeepShareLinks) {
		toSerialize["keepShareLinks"] = o.KeepShareLinks
	}
	if !IsNil(o.NodeIds) {
		toSerialize["nodeIds"] = o.NodeIds
	}
	return toSerialize, nil
}

type NullableMoveNodesRequest struct {
	value *MoveNodesRequest
	isSet bool
}

func (v NullableMoveNodesRequest) Get() *MoveNodesRequest {
	return v.value
}

func (v *NullableMoveNodesRequest) Set(val *MoveNodesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveNodesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveNodesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveNodesRequest(val *MoveNodesRequest) *NullableMoveNodesRequest {
	return &NullableMoveNodesRequest{value: val, isSet: true}
}

func (v NullableMoveNodesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveNodesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
