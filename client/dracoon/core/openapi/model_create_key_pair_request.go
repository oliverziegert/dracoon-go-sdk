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

// checks if the CreateKeyPairRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateKeyPairRequest{}

// CreateKeyPairRequest Request model for creating a key pair
type CreateKeyPairRequest struct {
	PrivateKeyContainer PrivateKeyContainer `json:"privateKeyContainer"`
	PublicKeyContainer  PublicKeyContainer  `json:"publicKeyContainer"`
	PreviousPrivateKey  PrivateKeyContainer `json:"previousPrivateKey"`
}

type _CreateKeyPairRequest CreateKeyPairRequest

// NewCreateKeyPairRequest instantiates a new CreateKeyPairRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKeyPairRequest(privateKeyContainer PrivateKeyContainer, publicKeyContainer PublicKeyContainer, previousPrivateKey PrivateKeyContainer) *CreateKeyPairRequest {
	this := CreateKeyPairRequest{}
	this.PrivateKeyContainer = privateKeyContainer
	this.PublicKeyContainer = publicKeyContainer
	this.PreviousPrivateKey = previousPrivateKey
	return &this
}

// NewCreateKeyPairRequestWithDefaults instantiates a new CreateKeyPairRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKeyPairRequestWithDefaults() *CreateKeyPairRequest {
	this := CreateKeyPairRequest{}
	return &this
}

// GetPrivateKeyContainer returns the PrivateKeyContainer field value
func (o *CreateKeyPairRequest) GetPrivateKeyContainer() PrivateKeyContainer {
	if o == nil {
		var ret PrivateKeyContainer
		return ret
	}

	return o.PrivateKeyContainer
}

// GetPrivateKeyContainerOk returns a tuple with the PrivateKeyContainer field value
// and a boolean to check if the value has been set.
func (o *CreateKeyPairRequest) GetPrivateKeyContainerOk() (*PrivateKeyContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKeyContainer, true
}

// SetPrivateKeyContainer sets field value
func (o *CreateKeyPairRequest) SetPrivateKeyContainer(v PrivateKeyContainer) {
	o.PrivateKeyContainer = v
}

// GetPublicKeyContainer returns the PublicKeyContainer field value
func (o *CreateKeyPairRequest) GetPublicKeyContainer() PublicKeyContainer {
	if o == nil {
		var ret PublicKeyContainer
		return ret
	}

	return o.PublicKeyContainer
}

// GetPublicKeyContainerOk returns a tuple with the PublicKeyContainer field value
// and a boolean to check if the value has been set.
func (o *CreateKeyPairRequest) GetPublicKeyContainerOk() (*PublicKeyContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKeyContainer, true
}

// SetPublicKeyContainer sets field value
func (o *CreateKeyPairRequest) SetPublicKeyContainer(v PublicKeyContainer) {
	o.PublicKeyContainer = v
}

// GetPreviousPrivateKey returns the PreviousPrivateKey field value
func (o *CreateKeyPairRequest) GetPreviousPrivateKey() PrivateKeyContainer {
	if o == nil {
		var ret PrivateKeyContainer
		return ret
	}

	return o.PreviousPrivateKey
}

// GetPreviousPrivateKeyOk returns a tuple with the PreviousPrivateKey field value
// and a boolean to check if the value has been set.
func (o *CreateKeyPairRequest) GetPreviousPrivateKeyOk() (*PrivateKeyContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousPrivateKey, true
}

// SetPreviousPrivateKey sets field value
func (o *CreateKeyPairRequest) SetPreviousPrivateKey(v PrivateKeyContainer) {
	o.PreviousPrivateKey = v
}

func (o CreateKeyPairRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateKeyPairRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["privateKeyContainer"] = o.PrivateKeyContainer
	toSerialize["publicKeyContainer"] = o.PublicKeyContainer
	toSerialize["previousPrivateKey"] = o.PreviousPrivateKey
	return toSerialize, nil
}

func (o *CreateKeyPairRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"privateKeyContainer",
		"publicKeyContainer",
		"previousPrivateKey",
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

	varCreateKeyPairRequest := _CreateKeyPairRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateKeyPairRequest)

	if err != nil {
		return err
	}

	*o = CreateKeyPairRequest(varCreateKeyPairRequest)

	return err
}

type NullableCreateKeyPairRequest struct {
	value *CreateKeyPairRequest
	isSet bool
}

func (v NullableCreateKeyPairRequest) Get() *CreateKeyPairRequest {
	return v.value
}

func (v *NullableCreateKeyPairRequest) Set(val *CreateKeyPairRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKeyPairRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKeyPairRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKeyPairRequest(val *CreateKeyPairRequest) *NullableCreateKeyPairRequest {
	return &NullableCreateKeyPairRequest{value: val, isSet: true}
}

func (v NullableCreateKeyPairRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKeyPairRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
