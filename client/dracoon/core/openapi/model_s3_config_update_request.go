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

// checks if the S3ConfigUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3ConfigUpdateRequest{}

// S3ConfigUpdateRequest Request model for updating a S3 configuration
type S3ConfigUpdateRequest struct {
	// S3 object storage bucket URL
	BucketUrl *string `json:"bucketUrl,omitempty"`
	// Access Key ID
	AccessKey *string `json:"accessKey,omitempty"`
	// Secret Access Key
	SecretKey *string `json:"secretKey,omitempty"`
	// S3 region
	Region *string `json:"region,omitempty"`
	// &#128679; Deprecated since v4.24.0  S3 object storage endpoint URL  use `bucketUrl` instead
	// Deprecated
	EndpointUrl *string `json:"endpointUrl,omitempty"`
	// &#128679; Deprecated since v4.24.0  S3 bucket name  use `bucketUrl` instead
	// Deprecated
	BucketName *string `json:"bucketName,omitempty"`
}

// NewS3ConfigUpdateRequest instantiates a new S3ConfigUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3ConfigUpdateRequest() *S3ConfigUpdateRequest {
	this := S3ConfigUpdateRequest{}
	return &this
}

// NewS3ConfigUpdateRequestWithDefaults instantiates a new S3ConfigUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3ConfigUpdateRequestWithDefaults() *S3ConfigUpdateRequest {
	this := S3ConfigUpdateRequest{}
	return &this
}

// GetBucketUrl returns the BucketUrl field value if set, zero value otherwise.
func (o *S3ConfigUpdateRequest) GetBucketUrl() string {
	if o == nil || IsNil(o.BucketUrl) {
		var ret string
		return ret
	}
	return *o.BucketUrl
}

// GetBucketUrlOk returns a tuple with the BucketUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ConfigUpdateRequest) GetBucketUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BucketUrl) {
		return nil, false
	}
	return o.BucketUrl, true
}

// HasBucketUrl returns a boolean if a field has been set.
func (o *S3ConfigUpdateRequest) HasBucketUrl() bool {
	if o != nil && !IsNil(o.BucketUrl) {
		return true
	}

	return false
}

// SetBucketUrl gets a reference to the given string and assigns it to the BucketUrl field.
func (o *S3ConfigUpdateRequest) SetBucketUrl(v string) {
	o.BucketUrl = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *S3ConfigUpdateRequest) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ConfigUpdateRequest) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *S3ConfigUpdateRequest) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *S3ConfigUpdateRequest) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *S3ConfigUpdateRequest) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ConfigUpdateRequest) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *S3ConfigUpdateRequest) HasSecretKey() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *S3ConfigUpdateRequest) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *S3ConfigUpdateRequest) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ConfigUpdateRequest) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *S3ConfigUpdateRequest) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *S3ConfigUpdateRequest) SetRegion(v string) {
	o.Region = &v
}

// GetEndpointUrl returns the EndpointUrl field value if set, zero value otherwise.
// Deprecated
func (o *S3ConfigUpdateRequest) GetEndpointUrl() string {
	if o == nil || IsNil(o.EndpointUrl) {
		var ret string
		return ret
	}
	return *o.EndpointUrl
}

// GetEndpointUrlOk returns a tuple with the EndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *S3ConfigUpdateRequest) GetEndpointUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointUrl) {
		return nil, false
	}
	return o.EndpointUrl, true
}

// HasEndpointUrl returns a boolean if a field has been set.
func (o *S3ConfigUpdateRequest) HasEndpointUrl() bool {
	if o != nil && !IsNil(o.EndpointUrl) {
		return true
	}

	return false
}

// SetEndpointUrl gets a reference to the given string and assigns it to the EndpointUrl field.
// Deprecated
func (o *S3ConfigUpdateRequest) SetEndpointUrl(v string) {
	o.EndpointUrl = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
// Deprecated
func (o *S3ConfigUpdateRequest) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *S3ConfigUpdateRequest) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *S3ConfigUpdateRequest) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
// Deprecated
func (o *S3ConfigUpdateRequest) SetBucketName(v string) {
	o.BucketName = &v
}

func (o S3ConfigUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3ConfigUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BucketUrl) {
		toSerialize["bucketUrl"] = o.BucketUrl
	}
	if !IsNil(o.AccessKey) {
		toSerialize["accessKey"] = o.AccessKey
	}
	if !IsNil(o.SecretKey) {
		toSerialize["secretKey"] = o.SecretKey
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.EndpointUrl) {
		toSerialize["endpointUrl"] = o.EndpointUrl
	}
	if !IsNil(o.BucketName) {
		toSerialize["bucketName"] = o.BucketName
	}
	return toSerialize, nil
}

type NullableS3ConfigUpdateRequest struct {
	value *S3ConfigUpdateRequest
	isSet bool
}

func (v NullableS3ConfigUpdateRequest) Get() *S3ConfigUpdateRequest {
	return v.value
}

func (v *NullableS3ConfigUpdateRequest) Set(val *S3ConfigUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableS3ConfigUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableS3ConfigUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3ConfigUpdateRequest(val *S3ConfigUpdateRequest) *NullableS3ConfigUpdateRequest {
	return &NullableS3ConfigUpdateRequest{value: val, isSet: true}
}

func (v NullableS3ConfigUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3ConfigUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
