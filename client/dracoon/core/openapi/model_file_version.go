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

// checks if the FileVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileVersion{}

// FileVersion List of File Versions
type FileVersion struct {
	// Node ID
	Id int64 `json:"id"`
	// Reference ID. Identical across all versions of a file
	ReferenceId int64 `json:"referenceId"`
	// Name
	Name string `json:"name"`
	// Parent node ID (room or folder)
	ParentId *int64 `json:"parentId,omitempty"`
	Deleted  *bool  `json:"deleted,omitempty"`
}

type _FileVersion FileVersion

// NewFileVersion instantiates a new FileVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileVersion(id int64, referenceId int64, name string) *FileVersion {
	this := FileVersion{}
	this.Id = id
	this.ReferenceId = referenceId
	this.Name = name
	return &this
}

// NewFileVersionWithDefaults instantiates a new FileVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileVersionWithDefaults() *FileVersion {
	this := FileVersion{}
	return &this
}

// GetId returns the Id field value
func (o *FileVersion) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileVersion) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileVersion) SetId(v int64) {
	o.Id = v
}

// GetReferenceId returns the ReferenceId field value
func (o *FileVersion) GetReferenceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *FileVersion) GetReferenceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *FileVersion) SetReferenceId(v int64) {
	o.ReferenceId = v
}

// GetName returns the Name field value
func (o *FileVersion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FileVersion) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FileVersion) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *FileVersion) GetParentId() int64 {
	if o == nil || IsNil(o.ParentId) {
		var ret int64
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersion) GetParentIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *FileVersion) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int64 and assigns it to the ParentId field.
func (o *FileVersion) SetParentId(v int64) {
	o.ParentId = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *FileVersion) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersion) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *FileVersion) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *FileVersion) SetDeleted(v bool) {
	o.Deleted = &v
}

func (o FileVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["referenceId"] = o.ReferenceId
	toSerialize["name"] = o.Name
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	return toSerialize, nil
}

func (o *FileVersion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"referenceId",
		"name",
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

	varFileVersion := _FileVersion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileVersion)

	if err != nil {
		return err
	}

	*o = FileVersion(varFileVersion)

	return err
}

type NullableFileVersion struct {
	value *FileVersion
	isSet bool
}

func (v NullableFileVersion) Get() *FileVersion {
	return v.value
}

func (v *NullableFileVersion) Set(val *FileVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableFileVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableFileVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileVersion(val *FileVersion) *NullableFileVersion {
	return &NullableFileVersion{value: val, isSet: true}
}

func (v NullableFileVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
