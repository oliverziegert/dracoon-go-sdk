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

// checks if the SystemInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInfo{}

// SystemInfo System information (default language and authentication methods)
type SystemInfo struct {
	// System default language  cf. [RFC 5646](https://tools.ietf.org/html/rfc5646)
	LanguageDefault string `json:"languageDefault"`
	// &#128679; Deprecated since v4.42.0  Defines if login fields should be hidden
	// Deprecated
	HideLoginInputFields bool `json:"hideLoginInputFields"`
	// &#128640; Since v4.14.0  List of S3 Hosts for CSP header
	S3Hosts []string `json:"s3Hosts"`
	// &#128640; Since v4.15.0  Determines whether S3 direct upload is enforced or not
	S3EnforceDirectUpload bool `json:"s3EnforceDirectUpload"`
	// &#128640; Since v4.21.0  Defines if S3 is used as storage backend
	UseS3Storage bool `json:"useS3Storage"`
	// &#128679; Deprecated since v4.13.0  Authentication methods:  * `sql`  * `active_directory`  * `openid`  use `authData` instead
	// Deprecated
	AuthMethods []AuthMethod `json:"authMethods"`
}

type _SystemInfo SystemInfo

// NewSystemInfo instantiates a new SystemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInfo(languageDefault string, hideLoginInputFields bool, s3Hosts []string, s3EnforceDirectUpload bool, useS3Storage bool, authMethods []AuthMethod) *SystemInfo {
	this := SystemInfo{}
	this.LanguageDefault = languageDefault
	this.HideLoginInputFields = hideLoginInputFields
	this.S3Hosts = s3Hosts
	this.S3EnforceDirectUpload = s3EnforceDirectUpload
	this.UseS3Storage = useS3Storage
	this.AuthMethods = authMethods
	return &this
}

// NewSystemInfoWithDefaults instantiates a new SystemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInfoWithDefaults() *SystemInfo {
	this := SystemInfo{}
	return &this
}

// GetLanguageDefault returns the LanguageDefault field value
func (o *SystemInfo) GetLanguageDefault() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LanguageDefault
}

// GetLanguageDefaultOk returns a tuple with the LanguageDefault field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetLanguageDefaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LanguageDefault, true
}

// SetLanguageDefault sets field value
func (o *SystemInfo) SetLanguageDefault(v string) {
	o.LanguageDefault = v
}

// GetHideLoginInputFields returns the HideLoginInputFields field value
// Deprecated
func (o *SystemInfo) GetHideLoginInputFields() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HideLoginInputFields
}

// GetHideLoginInputFieldsOk returns a tuple with the HideLoginInputFields field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *SystemInfo) GetHideLoginInputFieldsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HideLoginInputFields, true
}

// SetHideLoginInputFields sets field value
// Deprecated
func (o *SystemInfo) SetHideLoginInputFields(v bool) {
	o.HideLoginInputFields = v
}

// GetS3Hosts returns the S3Hosts field value
func (o *SystemInfo) GetS3Hosts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.S3Hosts
}

// GetS3HostsOk returns a tuple with the S3Hosts field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetS3HostsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3Hosts, true
}

// SetS3Hosts sets field value
func (o *SystemInfo) SetS3Hosts(v []string) {
	o.S3Hosts = v
}

// GetS3EnforceDirectUpload returns the S3EnforceDirectUpload field value
func (o *SystemInfo) GetS3EnforceDirectUpload() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.S3EnforceDirectUpload
}

// GetS3EnforceDirectUploadOk returns a tuple with the S3EnforceDirectUpload field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetS3EnforceDirectUploadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.S3EnforceDirectUpload, true
}

// SetS3EnforceDirectUpload sets field value
func (o *SystemInfo) SetS3EnforceDirectUpload(v bool) {
	o.S3EnforceDirectUpload = v
}

// GetUseS3Storage returns the UseS3Storage field value
func (o *SystemInfo) GetUseS3Storage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseS3Storage
}

// GetUseS3StorageOk returns a tuple with the UseS3Storage field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetUseS3StorageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseS3Storage, true
}

// SetUseS3Storage sets field value
func (o *SystemInfo) SetUseS3Storage(v bool) {
	o.UseS3Storage = v
}

// GetAuthMethods returns the AuthMethods field value
// Deprecated
func (o *SystemInfo) GetAuthMethods() []AuthMethod {
	if o == nil {
		var ret []AuthMethod
		return ret
	}

	return o.AuthMethods
}

// GetAuthMethodsOk returns a tuple with the AuthMethods field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *SystemInfo) GetAuthMethodsOk() ([]AuthMethod, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthMethods, true
}

// SetAuthMethods sets field value
// Deprecated
func (o *SystemInfo) SetAuthMethods(v []AuthMethod) {
	o.AuthMethods = v
}

func (o SystemInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["languageDefault"] = o.LanguageDefault
	toSerialize["hideLoginInputFields"] = o.HideLoginInputFields
	toSerialize["s3Hosts"] = o.S3Hosts
	toSerialize["s3EnforceDirectUpload"] = o.S3EnforceDirectUpload
	toSerialize["useS3Storage"] = o.UseS3Storage
	toSerialize["authMethods"] = o.AuthMethods
	return toSerialize, nil
}

func (o *SystemInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"languageDefault",
		"hideLoginInputFields",
		"s3Hosts",
		"s3EnforceDirectUpload",
		"useS3Storage",
		"authMethods",
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

	varSystemInfo := _SystemInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSystemInfo)

	if err != nil {
		return err
	}

	*o = SystemInfo(varSystemInfo)

	return err
}

type NullableSystemInfo struct {
	value *SystemInfo
	isSet bool
}

func (v NullableSystemInfo) Get() *SystemInfo {
	return v.value
}

func (v *NullableSystemInfo) Set(val *SystemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInfo(val *SystemInfo) *NullableSystemInfo {
	return &NullableSystemInfo{value: val, isSet: true}
}

func (v NullableSystemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
