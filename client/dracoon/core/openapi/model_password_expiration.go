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

// checks if the PasswordExpiration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordExpiration{}

// PasswordExpiration Password expiration information
type PasswordExpiration struct {
	// Determines whether password expiration is enabled
	Enabled bool `json:"enabled"`
	// Maximum allowed password age (in days)
	MaxPasswordAge *int32 `json:"maxPasswordAge,omitempty"`
}

type _PasswordExpiration PasswordExpiration

// NewPasswordExpiration instantiates a new PasswordExpiration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordExpiration(enabled bool) *PasswordExpiration {
	this := PasswordExpiration{}
	this.Enabled = enabled
	return &this
}

// NewPasswordExpirationWithDefaults instantiates a new PasswordExpiration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordExpirationWithDefaults() *PasswordExpiration {
	this := PasswordExpiration{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *PasswordExpiration) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PasswordExpiration) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PasswordExpiration) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMaxPasswordAge returns the MaxPasswordAge field value if set, zero value otherwise.
func (o *PasswordExpiration) GetMaxPasswordAge() int32 {
	if o == nil || IsNil(o.MaxPasswordAge) {
		var ret int32
		return ret
	}
	return *o.MaxPasswordAge
}

// GetMaxPasswordAgeOk returns a tuple with the MaxPasswordAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordExpiration) GetMaxPasswordAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPasswordAge) {
		return nil, false
	}
	return o.MaxPasswordAge, true
}

// HasMaxPasswordAge returns a boolean if a field has been set.
func (o *PasswordExpiration) HasMaxPasswordAge() bool {
	if o != nil && !IsNil(o.MaxPasswordAge) {
		return true
	}

	return false
}

// SetMaxPasswordAge gets a reference to the given int32 and assigns it to the MaxPasswordAge field.
func (o *PasswordExpiration) SetMaxPasswordAge(v int32) {
	o.MaxPasswordAge = &v
}

func (o PasswordExpiration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordExpiration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.MaxPasswordAge) {
		toSerialize["maxPasswordAge"] = o.MaxPasswordAge
	}
	return toSerialize, nil
}

func (o *PasswordExpiration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
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

	varPasswordExpiration := _PasswordExpiration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPasswordExpiration)

	if err != nil {
		return err
	}

	*o = PasswordExpiration(varPasswordExpiration)

	return err
}

type NullablePasswordExpiration struct {
	value *PasswordExpiration
	isSet bool
}

func (v NullablePasswordExpiration) Get() *PasswordExpiration {
	return v.value
}

func (v *NullablePasswordExpiration) Set(val *PasswordExpiration) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordExpiration) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordExpiration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordExpiration(val *PasswordExpiration) *NullablePasswordExpiration {
	return &NullablePasswordExpiration{value: val, isSet: true}
}

func (v NullablePasswordExpiration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordExpiration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
