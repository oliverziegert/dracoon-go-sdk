/*
DRACOON API

REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>

API version: 5.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the UpdateFolderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFolderRequest{}

// UpdateFolderRequest Request model for updating folder's metadata
type UpdateFolderRequest struct {
	// Folder name
	Name *string `json:"name,omitempty"`
	// User notes  Use empty string to remove.
	Notes *string `json:"notes,omitempty"`
	// &#128640; Since v4.30.0  Classification ID:  * `1` - public  * `2` - internal  * `3` - confidential  * `4` - strictly confidential    Provided (or default) classification is taken from room  when file gets uploaded without any classification.
	Classification *int32 `json:"classification,omitempty"`
	// &#128640; Since v4.22.0  Time the node was created on external file system  (default: current server datetime in UTC format)
	TimestampCreation *time.Time `json:"timestampCreation,omitempty"`
	// &#128640; Since v4.22.0  Time the content of a node was last modified on external file system  (default: current server datetime in UTC format)
	TimestampModification *time.Time `json:"timestampModification,omitempty"`
}

// NewUpdateFolderRequest instantiates a new UpdateFolderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFolderRequest() *UpdateFolderRequest {
	this := UpdateFolderRequest{}
	return &this
}

// NewUpdateFolderRequestWithDefaults instantiates a new UpdateFolderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFolderRequestWithDefaults() *UpdateFolderRequest {
	this := UpdateFolderRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateFolderRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFolderRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateFolderRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateFolderRequest) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *UpdateFolderRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFolderRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *UpdateFolderRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *UpdateFolderRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetClassification returns the Classification field value if set, zero value otherwise.
func (o *UpdateFolderRequest) GetClassification() int32 {
	if o == nil || IsNil(o.Classification) {
		var ret int32
		return ret
	}
	return *o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFolderRequest) GetClassificationOk() (*int32, bool) {
	if o == nil || IsNil(o.Classification) {
		return nil, false
	}
	return o.Classification, true
}

// HasClassification returns a boolean if a field has been set.
func (o *UpdateFolderRequest) HasClassification() bool {
	if o != nil && !IsNil(o.Classification) {
		return true
	}

	return false
}

// SetClassification gets a reference to the given int32 and assigns it to the Classification field.
func (o *UpdateFolderRequest) SetClassification(v int32) {
	o.Classification = &v
}

// GetTimestampCreation returns the TimestampCreation field value if set, zero value otherwise.
func (o *UpdateFolderRequest) GetTimestampCreation() time.Time {
	if o == nil || IsNil(o.TimestampCreation) {
		var ret time.Time
		return ret
	}
	return *o.TimestampCreation
}

// GetTimestampCreationOk returns a tuple with the TimestampCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFolderRequest) GetTimestampCreationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampCreation) {
		return nil, false
	}
	return o.TimestampCreation, true
}

// HasTimestampCreation returns a boolean if a field has been set.
func (o *UpdateFolderRequest) HasTimestampCreation() bool {
	if o != nil && !IsNil(o.TimestampCreation) {
		return true
	}

	return false
}

// SetTimestampCreation gets a reference to the given time.Time and assigns it to the TimestampCreation field.
func (o *UpdateFolderRequest) SetTimestampCreation(v time.Time) {
	o.TimestampCreation = &v
}

// GetTimestampModification returns the TimestampModification field value if set, zero value otherwise.
func (o *UpdateFolderRequest) GetTimestampModification() time.Time {
	if o == nil || IsNil(o.TimestampModification) {
		var ret time.Time
		return ret
	}
	return *o.TimestampModification
}

// GetTimestampModificationOk returns a tuple with the TimestampModification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFolderRequest) GetTimestampModificationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampModification) {
		return nil, false
	}
	return o.TimestampModification, true
}

// HasTimestampModification returns a boolean if a field has been set.
func (o *UpdateFolderRequest) HasTimestampModification() bool {
	if o != nil && !IsNil(o.TimestampModification) {
		return true
	}

	return false
}

// SetTimestampModification gets a reference to the given time.Time and assigns it to the TimestampModification field.
func (o *UpdateFolderRequest) SetTimestampModification(v time.Time) {
	o.TimestampModification = &v
}

func (o UpdateFolderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFolderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Classification) {
		toSerialize["classification"] = o.Classification
	}
	if !IsNil(o.TimestampCreation) {
		toSerialize["timestampCreation"] = o.TimestampCreation
	}
	if !IsNil(o.TimestampModification) {
		toSerialize["timestampModification"] = o.TimestampModification
	}
	return toSerialize, nil
}

type NullableUpdateFolderRequest struct {
	value *UpdateFolderRequest
	isSet bool
}

func (v NullableUpdateFolderRequest) Get() *UpdateFolderRequest {
	return v.value
}

func (v *NullableUpdateFolderRequest) Set(val *UpdateFolderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFolderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFolderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFolderRequest(val *UpdateFolderRequest) *NullableUpdateFolderRequest {
	return &NullableUpdateFolderRequest{value: val, isSet: true}
}

func (v NullableUpdateFolderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFolderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
