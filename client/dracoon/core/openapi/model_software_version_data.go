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
	"time"
)

// checks if the SoftwareVersionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftwareVersionData{}

// SoftwareVersionData Software version information
type SoftwareVersionData struct {
	// REST API version
	RestApiVersion string `json:"restApiVersion"`
	// DRACOON server version
	SdsServerVersion string `json:"sdsServerVersion"`
	// Build date
	BuildDate time.Time `json:"buildDate"`
	// Revision number
	ScmRevisionNumber string `json:"scmRevisionNumber"`
	// &#128640; Since v4.24.0  Determines if the DRACOON Core is deployed in the cloud environment
	IsDracoonCloud *bool `json:"isDracoonCloud,omitempty"`
}

type _SoftwareVersionData SoftwareVersionData

// NewSoftwareVersionData instantiates a new SoftwareVersionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareVersionData(restApiVersion string, sdsServerVersion string, buildDate time.Time, scmRevisionNumber string) *SoftwareVersionData {
	this := SoftwareVersionData{}
	this.RestApiVersion = restApiVersion
	this.SdsServerVersion = sdsServerVersion
	this.BuildDate = buildDate
	this.ScmRevisionNumber = scmRevisionNumber
	return &this
}

// NewSoftwareVersionDataWithDefaults instantiates a new SoftwareVersionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareVersionDataWithDefaults() *SoftwareVersionData {
	this := SoftwareVersionData{}
	return &this
}

// GetRestApiVersion returns the RestApiVersion field value
func (o *SoftwareVersionData) GetRestApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RestApiVersion
}

// GetRestApiVersionOk returns a tuple with the RestApiVersion field value
// and a boolean to check if the value has been set.
func (o *SoftwareVersionData) GetRestApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestApiVersion, true
}

// SetRestApiVersion sets field value
func (o *SoftwareVersionData) SetRestApiVersion(v string) {
	o.RestApiVersion = v
}

// GetSdsServerVersion returns the SdsServerVersion field value
func (o *SoftwareVersionData) GetSdsServerVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SdsServerVersion
}

// GetSdsServerVersionOk returns a tuple with the SdsServerVersion field value
// and a boolean to check if the value has been set.
func (o *SoftwareVersionData) GetSdsServerVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SdsServerVersion, true
}

// SetSdsServerVersion sets field value
func (o *SoftwareVersionData) SetSdsServerVersion(v string) {
	o.SdsServerVersion = v
}

// GetBuildDate returns the BuildDate field value
func (o *SoftwareVersionData) GetBuildDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.BuildDate
}

// GetBuildDateOk returns a tuple with the BuildDate field value
// and a boolean to check if the value has been set.
func (o *SoftwareVersionData) GetBuildDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildDate, true
}

// SetBuildDate sets field value
func (o *SoftwareVersionData) SetBuildDate(v time.Time) {
	o.BuildDate = v
}

// GetScmRevisionNumber returns the ScmRevisionNumber field value
func (o *SoftwareVersionData) GetScmRevisionNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScmRevisionNumber
}

// GetScmRevisionNumberOk returns a tuple with the ScmRevisionNumber field value
// and a boolean to check if the value has been set.
func (o *SoftwareVersionData) GetScmRevisionNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScmRevisionNumber, true
}

// SetScmRevisionNumber sets field value
func (o *SoftwareVersionData) SetScmRevisionNumber(v string) {
	o.ScmRevisionNumber = v
}

// GetIsDracoonCloud returns the IsDracoonCloud field value if set, zero value otherwise.
func (o *SoftwareVersionData) GetIsDracoonCloud() bool {
	if o == nil || IsNil(o.IsDracoonCloud) {
		var ret bool
		return ret
	}
	return *o.IsDracoonCloud
}

// GetIsDracoonCloudOk returns a tuple with the IsDracoonCloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareVersionData) GetIsDracoonCloudOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDracoonCloud) {
		return nil, false
	}
	return o.IsDracoonCloud, true
}

// HasIsDracoonCloud returns a boolean if a field has been set.
func (o *SoftwareVersionData) HasIsDracoonCloud() bool {
	if o != nil && !IsNil(o.IsDracoonCloud) {
		return true
	}

	return false
}

// SetIsDracoonCloud gets a reference to the given bool and assigns it to the IsDracoonCloud field.
func (o *SoftwareVersionData) SetIsDracoonCloud(v bool) {
	o.IsDracoonCloud = &v
}

func (o SoftwareVersionData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftwareVersionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["restApiVersion"] = o.RestApiVersion
	toSerialize["sdsServerVersion"] = o.SdsServerVersion
	toSerialize["buildDate"] = o.BuildDate
	toSerialize["scmRevisionNumber"] = o.ScmRevisionNumber
	if !IsNil(o.IsDracoonCloud) {
		toSerialize["isDracoonCloud"] = o.IsDracoonCloud
	}
	return toSerialize, nil
}

func (o *SoftwareVersionData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"restApiVersion",
		"sdsServerVersion",
		"buildDate",
		"scmRevisionNumber",
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

	varSoftwareVersionData := _SoftwareVersionData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSoftwareVersionData)

	if err != nil {
		return err
	}

	*o = SoftwareVersionData(varSoftwareVersionData)

	return err
}

type NullableSoftwareVersionData struct {
	value *SoftwareVersionData
	isSet bool
}

func (v NullableSoftwareVersionData) Get() *SoftwareVersionData {
	return v.value
}

func (v *NullableSoftwareVersionData) Set(val *SoftwareVersionData) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareVersionData) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareVersionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareVersionData(val *SoftwareVersionData) *NullableSoftwareVersionData {
	return &NullableSoftwareVersionData{value: val, isSet: true}
}

func (v NullableSoftwareVersionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareVersionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
