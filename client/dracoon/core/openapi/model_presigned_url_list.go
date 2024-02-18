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

// checks if the PresignedUrlList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PresignedUrlList{}

// PresignedUrlList List of generated presigned URLs
type PresignedUrlList struct {
	// List of S3 presigned URLs
	Urls []PresignedUrl `json:"urls"`
}

type _PresignedUrlList PresignedUrlList

// NewPresignedUrlList instantiates a new PresignedUrlList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresignedUrlList(urls []PresignedUrl) *PresignedUrlList {
	this := PresignedUrlList{}
	this.Urls = urls
	return &this
}

// NewPresignedUrlListWithDefaults instantiates a new PresignedUrlList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresignedUrlListWithDefaults() *PresignedUrlList {
	this := PresignedUrlList{}
	return &this
}

// GetUrls returns the Urls field value
func (o *PresignedUrlList) GetUrls() []PresignedUrl {
	if o == nil {
		var ret []PresignedUrl
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *PresignedUrlList) GetUrlsOk() ([]PresignedUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.Urls, true
}

// SetUrls sets field value
func (o *PresignedUrlList) SetUrls(v []PresignedUrl) {
	o.Urls = v
}

func (o PresignedUrlList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PresignedUrlList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["urls"] = o.Urls
	return toSerialize, nil
}

func (o *PresignedUrlList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"urls",
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

	varPresignedUrlList := _PresignedUrlList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPresignedUrlList)

	if err != nil {
		return err
	}

	*o = PresignedUrlList(varPresignedUrlList)

	return err
}

type NullablePresignedUrlList struct {
	value *PresignedUrlList
	isSet bool
}

func (v NullablePresignedUrlList) Get() *PresignedUrlList {
	return v.value
}

func (v *NullablePresignedUrlList) Set(val *PresignedUrlList) {
	v.value = val
	v.isSet = true
}

func (v NullablePresignedUrlList) IsSet() bool {
	return v.isSet
}

func (v *NullablePresignedUrlList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresignedUrlList(val *PresignedUrlList) *NullablePresignedUrlList {
	return &NullablePresignedUrlList{value: val, isSet: true}
}

func (v NullablePresignedUrlList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresignedUrlList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
