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

// checks if the RoomPoliciesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomPoliciesRequest{}

// RoomPoliciesRequest Room Policies
type RoomPoliciesRequest struct {
	// Default policy room expiration period in seconds.  All files in a room will have their expiration date set to this period after their respective upload.   0 means no default expiration policy is set.
	DefaultExpirationPeriod *int32 `json:"defaultExpirationPeriod,omitempty"`
	// Determines whether virus protection is enabled for room. To be set by room admins
	VirusProtectionEnabled *bool `json:"virusProtectionEnabled,omitempty"`
}

// NewRoomPoliciesRequest instantiates a new RoomPoliciesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomPoliciesRequest() *RoomPoliciesRequest {
	this := RoomPoliciesRequest{}
	return &this
}

// NewRoomPoliciesRequestWithDefaults instantiates a new RoomPoliciesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomPoliciesRequestWithDefaults() *RoomPoliciesRequest {
	this := RoomPoliciesRequest{}
	return &this
}

// GetDefaultExpirationPeriod returns the DefaultExpirationPeriod field value if set, zero value otherwise.
func (o *RoomPoliciesRequest) GetDefaultExpirationPeriod() int32 {
	if o == nil || IsNil(o.DefaultExpirationPeriod) {
		var ret int32
		return ret
	}
	return *o.DefaultExpirationPeriod
}

// GetDefaultExpirationPeriodOk returns a tuple with the DefaultExpirationPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomPoliciesRequest) GetDefaultExpirationPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultExpirationPeriod) {
		return nil, false
	}
	return o.DefaultExpirationPeriod, true
}

// HasDefaultExpirationPeriod returns a boolean if a field has been set.
func (o *RoomPoliciesRequest) HasDefaultExpirationPeriod() bool {
	if o != nil && !IsNil(o.DefaultExpirationPeriod) {
		return true
	}

	return false
}

// SetDefaultExpirationPeriod gets a reference to the given int32 and assigns it to the DefaultExpirationPeriod field.
func (o *RoomPoliciesRequest) SetDefaultExpirationPeriod(v int32) {
	o.DefaultExpirationPeriod = &v
}

// GetVirusProtectionEnabled returns the VirusProtectionEnabled field value if set, zero value otherwise.
func (o *RoomPoliciesRequest) GetVirusProtectionEnabled() bool {
	if o == nil || IsNil(o.VirusProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.VirusProtectionEnabled
}

// GetVirusProtectionEnabledOk returns a tuple with the VirusProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomPoliciesRequest) GetVirusProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.VirusProtectionEnabled) {
		return nil, false
	}
	return o.VirusProtectionEnabled, true
}

// HasVirusProtectionEnabled returns a boolean if a field has been set.
func (o *RoomPoliciesRequest) HasVirusProtectionEnabled() bool {
	if o != nil && !IsNil(o.VirusProtectionEnabled) {
		return true
	}

	return false
}

// SetVirusProtectionEnabled gets a reference to the given bool and assigns it to the VirusProtectionEnabled field.
func (o *RoomPoliciesRequest) SetVirusProtectionEnabled(v bool) {
	o.VirusProtectionEnabled = &v
}

func (o RoomPoliciesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomPoliciesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultExpirationPeriod) {
		toSerialize["defaultExpirationPeriod"] = o.DefaultExpirationPeriod
	}
	if !IsNil(o.VirusProtectionEnabled) {
		toSerialize["virusProtectionEnabled"] = o.VirusProtectionEnabled
	}
	return toSerialize, nil
}

type NullableRoomPoliciesRequest struct {
	value *RoomPoliciesRequest
	isSet bool
}

func (v NullableRoomPoliciesRequest) Get() *RoomPoliciesRequest {
	return v.value
}

func (v *NullableRoomPoliciesRequest) Set(val *RoomPoliciesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomPoliciesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomPoliciesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomPoliciesRequest(val *RoomPoliciesRequest) *NullableRoomPoliciesRequest {
	return &NullableRoomPoliciesRequest{value: val, isSet: true}
}

func (v NullableRoomPoliciesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomPoliciesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
