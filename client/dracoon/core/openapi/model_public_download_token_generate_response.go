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

// checks if the PublicDownloadTokenGenerateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicDownloadTokenGenerateResponse{}

// PublicDownloadTokenGenerateResponse Download URL
type PublicDownloadTokenGenerateResponse struct {
	// Download URL
	DownloadUrl *string `json:"downloadUrl,omitempty"`
}

// NewPublicDownloadTokenGenerateResponse instantiates a new PublicDownloadTokenGenerateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicDownloadTokenGenerateResponse() *PublicDownloadTokenGenerateResponse {
	this := PublicDownloadTokenGenerateResponse{}
	return &this
}

// NewPublicDownloadTokenGenerateResponseWithDefaults instantiates a new PublicDownloadTokenGenerateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicDownloadTokenGenerateResponseWithDefaults() *PublicDownloadTokenGenerateResponse {
	this := PublicDownloadTokenGenerateResponse{}
	return &this
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *PublicDownloadTokenGenerateResponse) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicDownloadTokenGenerateResponse) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *PublicDownloadTokenGenerateResponse) HasDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *PublicDownloadTokenGenerateResponse) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

func (o PublicDownloadTokenGenerateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicDownloadTokenGenerateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DownloadUrl) {
		toSerialize["downloadUrl"] = o.DownloadUrl
	}
	return toSerialize, nil
}

type NullablePublicDownloadTokenGenerateResponse struct {
	value *PublicDownloadTokenGenerateResponse
	isSet bool
}

func (v NullablePublicDownloadTokenGenerateResponse) Get() *PublicDownloadTokenGenerateResponse {
	return v.value
}

func (v *NullablePublicDownloadTokenGenerateResponse) Set(val *PublicDownloadTokenGenerateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicDownloadTokenGenerateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicDownloadTokenGenerateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicDownloadTokenGenerateResponse(val *PublicDownloadTokenGenerateResponse) *NullablePublicDownloadTokenGenerateResponse {
	return &NullablePublicDownloadTokenGenerateResponse{value: val, isSet: true}
}

func (v NullablePublicDownloadTokenGenerateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicDownloadTokenGenerateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
