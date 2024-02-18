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

// checks if the ViolatedPasswordPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolatedPasswordPolicy{}

// ViolatedPasswordPolicy Violated password policy information
type ViolatedPasswordPolicy struct {
	// Name of the violated password policy
	Name *string `json:"name,omitempty"`
	// Message from password validator
	Message *string `json:"message,omitempty"`
}

// NewViolatedPasswordPolicy instantiates a new ViolatedPasswordPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolatedPasswordPolicy() *ViolatedPasswordPolicy {
	this := ViolatedPasswordPolicy{}
	return &this
}

// NewViolatedPasswordPolicyWithDefaults instantiates a new ViolatedPasswordPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolatedPasswordPolicyWithDefaults() *ViolatedPasswordPolicy {
	this := ViolatedPasswordPolicy{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ViolatedPasswordPolicy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatedPasswordPolicy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ViolatedPasswordPolicy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ViolatedPasswordPolicy) SetName(v string) {
	o.Name = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ViolatedPasswordPolicy) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolatedPasswordPolicy) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ViolatedPasswordPolicy) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ViolatedPasswordPolicy) SetMessage(v string) {
	o.Message = &v
}

func (o ViolatedPasswordPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViolatedPasswordPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableViolatedPasswordPolicy struct {
	value *ViolatedPasswordPolicy
	isSet bool
}

func (v NullableViolatedPasswordPolicy) Get() *ViolatedPasswordPolicy {
	return v.value
}

func (v *NullableViolatedPasswordPolicy) Set(val *ViolatedPasswordPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableViolatedPasswordPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableViolatedPasswordPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolatedPasswordPolicy(val *ViolatedPasswordPolicy) *NullableViolatedPasswordPolicy {
	return &NullableViolatedPasswordPolicy{value: val, isSet: true}
}

func (v NullableViolatedPasswordPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViolatedPasswordPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
