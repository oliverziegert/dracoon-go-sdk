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

// checks if the CustomerSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerSettingsResponse{}

// CustomerSettingsResponse Customer settings
type CustomerSettingsResponse struct {
	// Homerooms active
	HomeRoomsActive bool `json:"homeRoomsActive"`
	// Homeroom Parent ID
	HomeRoomParentId *int64 `json:"homeRoomParentId,omitempty"`
	// Homeroom Parent Name
	HomeRoomParentName *string `json:"homeRoomParentName,omitempty"`
	// Homeroom Quota in bytes
	HomeRoomQuota *int64 `json:"homeRoomQuota,omitempty"`
}

type _CustomerSettingsResponse CustomerSettingsResponse

// NewCustomerSettingsResponse instantiates a new CustomerSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSettingsResponse(homeRoomsActive bool) *CustomerSettingsResponse {
	this := CustomerSettingsResponse{}
	this.HomeRoomsActive = homeRoomsActive
	return &this
}

// NewCustomerSettingsResponseWithDefaults instantiates a new CustomerSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSettingsResponseWithDefaults() *CustomerSettingsResponse {
	this := CustomerSettingsResponse{}
	return &this
}

// GetHomeRoomsActive returns the HomeRoomsActive field value
func (o *CustomerSettingsResponse) GetHomeRoomsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HomeRoomsActive
}

// GetHomeRoomsActiveOk returns a tuple with the HomeRoomsActive field value
// and a boolean to check if the value has been set.
func (o *CustomerSettingsResponse) GetHomeRoomsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomeRoomsActive, true
}

// SetHomeRoomsActive sets field value
func (o *CustomerSettingsResponse) SetHomeRoomsActive(v bool) {
	o.HomeRoomsActive = v
}

// GetHomeRoomParentId returns the HomeRoomParentId field value if set, zero value otherwise.
func (o *CustomerSettingsResponse) GetHomeRoomParentId() int64 {
	if o == nil || IsNil(o.HomeRoomParentId) {
		var ret int64
		return ret
	}
	return *o.HomeRoomParentId
}

// GetHomeRoomParentIdOk returns a tuple with the HomeRoomParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSettingsResponse) GetHomeRoomParentIdOk() (*int64, bool) {
	if o == nil || IsNil(o.HomeRoomParentId) {
		return nil, false
	}
	return o.HomeRoomParentId, true
}

// HasHomeRoomParentId returns a boolean if a field has been set.
func (o *CustomerSettingsResponse) HasHomeRoomParentId() bool {
	if o != nil && !IsNil(o.HomeRoomParentId) {
		return true
	}

	return false
}

// SetHomeRoomParentId gets a reference to the given int64 and assigns it to the HomeRoomParentId field.
func (o *CustomerSettingsResponse) SetHomeRoomParentId(v int64) {
	o.HomeRoomParentId = &v
}

// GetHomeRoomParentName returns the HomeRoomParentName field value if set, zero value otherwise.
func (o *CustomerSettingsResponse) GetHomeRoomParentName() string {
	if o == nil || IsNil(o.HomeRoomParentName) {
		var ret string
		return ret
	}
	return *o.HomeRoomParentName
}

// GetHomeRoomParentNameOk returns a tuple with the HomeRoomParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSettingsResponse) GetHomeRoomParentNameOk() (*string, bool) {
	if o == nil || IsNil(o.HomeRoomParentName) {
		return nil, false
	}
	return o.HomeRoomParentName, true
}

// HasHomeRoomParentName returns a boolean if a field has been set.
func (o *CustomerSettingsResponse) HasHomeRoomParentName() bool {
	if o != nil && !IsNil(o.HomeRoomParentName) {
		return true
	}

	return false
}

// SetHomeRoomParentName gets a reference to the given string and assigns it to the HomeRoomParentName field.
func (o *CustomerSettingsResponse) SetHomeRoomParentName(v string) {
	o.HomeRoomParentName = &v
}

// GetHomeRoomQuota returns the HomeRoomQuota field value if set, zero value otherwise.
func (o *CustomerSettingsResponse) GetHomeRoomQuota() int64 {
	if o == nil || IsNil(o.HomeRoomQuota) {
		var ret int64
		return ret
	}
	return *o.HomeRoomQuota
}

// GetHomeRoomQuotaOk returns a tuple with the HomeRoomQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSettingsResponse) GetHomeRoomQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.HomeRoomQuota) {
		return nil, false
	}
	return o.HomeRoomQuota, true
}

// HasHomeRoomQuota returns a boolean if a field has been set.
func (o *CustomerSettingsResponse) HasHomeRoomQuota() bool {
	if o != nil && !IsNil(o.HomeRoomQuota) {
		return true
	}

	return false
}

// SetHomeRoomQuota gets a reference to the given int64 and assigns it to the HomeRoomQuota field.
func (o *CustomerSettingsResponse) SetHomeRoomQuota(v int64) {
	o.HomeRoomQuota = &v
}

func (o CustomerSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["homeRoomsActive"] = o.HomeRoomsActive
	if !IsNil(o.HomeRoomParentId) {
		toSerialize["homeRoomParentId"] = o.HomeRoomParentId
	}
	if !IsNil(o.HomeRoomParentName) {
		toSerialize["homeRoomParentName"] = o.HomeRoomParentName
	}
	if !IsNil(o.HomeRoomQuota) {
		toSerialize["homeRoomQuota"] = o.HomeRoomQuota
	}
	return toSerialize, nil
}

func (o *CustomerSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"homeRoomsActive",
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

	varCustomerSettingsResponse := _CustomerSettingsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomerSettingsResponse)

	if err != nil {
		return err
	}

	*o = CustomerSettingsResponse(varCustomerSettingsResponse)

	return err
}

type NullableCustomerSettingsResponse struct {
	value *CustomerSettingsResponse
	isSet bool
}

func (v NullableCustomerSettingsResponse) Get() *CustomerSettingsResponse {
	return v.value
}

func (v *NullableCustomerSettingsResponse) Set(val *CustomerSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSettingsResponse(val *CustomerSettingsResponse) *NullableCustomerSettingsResponse {
	return &NullableCustomerSettingsResponse{value: val, isSet: true}
}

func (v NullableCustomerSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
