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

// checks if the UserUserPublicKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUserPublicKey{}

// UserUserPublicKey Public key information
type UserUserPublicKey struct {
	// Unique identifier for the user
	Id                 *int64              `json:"id,omitempty"`
	PublicKeyContainer *PublicKeyContainer `json:"publicKeyContainer,omitempty"`
}

// NewUserUserPublicKey instantiates a new UserUserPublicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUserPublicKey() *UserUserPublicKey {
	this := UserUserPublicKey{}
	return &this
}

// NewUserUserPublicKeyWithDefaults instantiates a new UserUserPublicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUserPublicKeyWithDefaults() *UserUserPublicKey {
	this := UserUserPublicKey{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserUserPublicKey) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUserPublicKey) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserUserPublicKey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UserUserPublicKey) SetId(v int64) {
	o.Id = &v
}

// GetPublicKeyContainer returns the PublicKeyContainer field value if set, zero value otherwise.
func (o *UserUserPublicKey) GetPublicKeyContainer() PublicKeyContainer {
	if o == nil || IsNil(o.PublicKeyContainer) {
		var ret PublicKeyContainer
		return ret
	}
	return *o.PublicKeyContainer
}

// GetPublicKeyContainerOk returns a tuple with the PublicKeyContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUserPublicKey) GetPublicKeyContainerOk() (*PublicKeyContainer, bool) {
	if o == nil || IsNil(o.PublicKeyContainer) {
		return nil, false
	}
	return o.PublicKeyContainer, true
}

// HasPublicKeyContainer returns a boolean if a field has been set.
func (o *UserUserPublicKey) HasPublicKeyContainer() bool {
	if o != nil && !IsNil(o.PublicKeyContainer) {
		return true
	}

	return false
}

// SetPublicKeyContainer gets a reference to the given PublicKeyContainer and assigns it to the PublicKeyContainer field.
func (o *UserUserPublicKey) SetPublicKeyContainer(v PublicKeyContainer) {
	o.PublicKeyContainer = &v
}

func (o UserUserPublicKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUserPublicKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PublicKeyContainer) {
		toSerialize["publicKeyContainer"] = o.PublicKeyContainer
	}
	return toSerialize, nil
}

type NullableUserUserPublicKey struct {
	value *UserUserPublicKey
	isSet bool
}

func (v NullableUserUserPublicKey) Get() *UserUserPublicKey {
	return v.value
}

func (v *NullableUserUserPublicKey) Set(val *UserUserPublicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUserPublicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUserPublicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUserPublicKey(val *UserUserPublicKey) *NullableUserUserPublicKey {
	return &NullableUserUserPublicKey{value: val, isSet: true}
}

func (v NullableUserUserPublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUserPublicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
