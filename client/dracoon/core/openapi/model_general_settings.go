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

// checks if the GeneralSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneralSettings{}

// GeneralSettings General settings
type GeneralSettings struct {
	// Allow sending of share passwords via SMS
	SharePasswordSmsEnabled *bool `json:"sharePasswordSmsEnabled,omitempty"`
	// Activation status of client-side encryption.  Can only be enabled once; disabling is not possible.
	CryptoEnabled *bool `json:"cryptoEnabled,omitempty"`
	// Enable email notification button
	EmailNotificationButtonEnabled *bool `json:"emailNotificationButtonEnabled,omitempty"`
	// Each user has to confirm the EULA at first login.
	EulaEnabled *bool `json:"eulaEnabled,omitempty"`
	// Defines if S3 is used as storage backend
	UseS3Storage *bool `json:"useS3Storage,omitempty"`
	// &#128640; Since v4.9.0  Defines if S3 tags are enabled
	S3TagsEnabled *bool `json:"s3TagsEnabled,omitempty"`
	// &#128679; Deprecated since v4.42.0  Defines if login fields should be hidden
	// Deprecated
	HideLoginInputFields  *bool                  `json:"hideLoginInputFields,omitempty"`
	AuthTokenRestrictions *AuthTokenRestrictions `json:"authTokenRestrictions,omitempty"`
	// &#128679; Deprecated since v4.12.0  Determines if the media server is enabled
	// Deprecated
	MediaServerEnabled *bool `json:"mediaServerEnabled,omitempty"`
	// &#128679; Deprecated since v4.14.0  Allow weak password  * A weak password has to fulfill the following criteria:     * is at least 8 characters long     * contains letters and numbers  * A strong password has to fulfill the following criteria in addition:     * contains at least one special character     * contains upper and lower case characters  Please use `GET /system/config/policies/passwords` API to get configured password policies.
	// Deprecated
	WeakPasswordEnabled *bool `json:"weakPasswordEnabled,omitempty"`
}

// NewGeneralSettings instantiates a new GeneralSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneralSettings() *GeneralSettings {
	this := GeneralSettings{}
	return &this
}

// NewGeneralSettingsWithDefaults instantiates a new GeneralSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralSettingsWithDefaults() *GeneralSettings {
	this := GeneralSettings{}
	return &this
}

// GetSharePasswordSmsEnabled returns the SharePasswordSmsEnabled field value if set, zero value otherwise.
func (o *GeneralSettings) GetSharePasswordSmsEnabled() bool {
	if o == nil || IsNil(o.SharePasswordSmsEnabled) {
		var ret bool
		return ret
	}
	return *o.SharePasswordSmsEnabled
}

// GetSharePasswordSmsEnabledOk returns a tuple with the SharePasswordSmsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettings) GetSharePasswordSmsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SharePasswordSmsEnabled) {
		return nil, false
	}
	return o.SharePasswordSmsEnabled, true
}

// HasSharePasswordSmsEnabled returns a boolean if a field has been set.
func (o *GeneralSettings) HasSharePasswordSmsEnabled() bool {
	if o != nil && !IsNil(o.SharePasswordSmsEnabled) {
		return true
	}

	return false
}

// SetSharePasswordSmsEnabled gets a reference to the given bool and assigns it to the SharePasswordSmsEnabled field.
func (o *GeneralSettings) SetSharePasswordSmsEnabled(v bool) {
	o.SharePasswordSmsEnabled = &v
}

// GetCryptoEnabled returns the CryptoEnabled field value if set, zero value otherwise.
func (o *GeneralSettings) GetCryptoEnabled() bool {
	if o == nil || IsNil(o.CryptoEnabled) {
		var ret bool
		return ret
	}
	return *o.CryptoEnabled
}

// GetCryptoEnabledOk returns a tuple with the CryptoEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettings) GetCryptoEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CryptoEnabled) {
		return nil, false
	}
	return o.CryptoEnabled, true
}

// HasCryptoEnabled returns a boolean if a field has been set.
func (o *GeneralSettings) HasCryptoEnabled() bool {
	if o != nil && !IsNil(o.CryptoEnabled) {
		return true
	}

	return false
}

// SetCryptoEnabled gets a reference to the given bool and assigns it to the CryptoEnabled field.
func (o *GeneralSettings) SetCryptoEnabled(v bool) {
	o.CryptoEnabled = &v
}

// GetEmailNotificationButtonEnabled returns the EmailNotificationButtonEnabled field value if set, zero value otherwise.
func (o *GeneralSettings) GetEmailNotificationButtonEnabled() bool {
	if o == nil || IsNil(o.EmailNotificationButtonEnabled) {
		var ret bool
		return ret
	}
	return *o.EmailNotificationButtonEnabled
}

// GetEmailNotificationButtonEnabledOk returns a tuple with the EmailNotificationButtonEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettings) GetEmailNotificationButtonEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailNotificationButtonEnabled) {
		return nil, false
	}
	return o.EmailNotificationButtonEnabled, true
}

// HasEmailNotificationButtonEnabled returns a boolean if a field has been set.
func (o *GeneralSettings) HasEmailNotificationButtonEnabled() bool {
	if o != nil && !IsNil(o.EmailNotificationButtonEnabled) {
		return true
	}

	return false
}

// SetEmailNotificationButtonEnabled gets a reference to the given bool and assigns it to the EmailNotificationButtonEnabled field.
func (o *GeneralSettings) SetEmailNotificationButtonEnabled(v bool) {
	o.EmailNotificationButtonEnabled = &v
}

// GetEulaEnabled returns the EulaEnabled field value if set, zero value otherwise.
func (o *GeneralSettings) GetEulaEnabled() bool {
	if o == nil || IsNil(o.EulaEnabled) {
		var ret bool
		return ret
	}
	return *o.EulaEnabled
}

// GetEulaEnabledOk returns a tuple with the EulaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettings) GetEulaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EulaEnabled) {
		return nil, false
	}
	return o.EulaEnabled, true
}

// HasEulaEnabled returns a boolean if a field has been set.
func (o *GeneralSettings) HasEulaEnabled() bool {
	if o != nil && !IsNil(o.EulaEnabled) {
		return true
	}

	return false
}

// SetEulaEnabled gets a reference to the given bool and assigns it to the EulaEnabled field.
func (o *GeneralSettings) SetEulaEnabled(v bool) {
	o.EulaEnabled = &v
}

// GetUseS3Storage returns the UseS3Storage field value if set, zero value otherwise.
func (o *GeneralSettings) GetUseS3Storage() bool {
	if o == nil || IsNil(o.UseS3Storage) {
		var ret bool
		return ret
	}
	return *o.UseS3Storage
}

// GetUseS3StorageOk returns a tuple with the UseS3Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettings) GetUseS3StorageOk() (*bool, bool) {
	if o == nil || IsNil(o.UseS3Storage) {
		return nil, false
	}
	return o.UseS3Storage, true
}

// HasUseS3Storage returns a boolean if a field has been set.
func (o *GeneralSettings) HasUseS3Storage() bool {
	if o != nil && !IsNil(o.UseS3Storage) {
		return true
	}

	return false
}

// SetUseS3Storage gets a reference to the given bool and assigns it to the UseS3Storage field.
func (o *GeneralSettings) SetUseS3Storage(v bool) {
	o.UseS3Storage = &v
}

// GetS3TagsEnabled returns the S3TagsEnabled field value if set, zero value otherwise.
func (o *GeneralSettings) GetS3TagsEnabled() bool {
	if o == nil || IsNil(o.S3TagsEnabled) {
		var ret bool
		return ret
	}
	return *o.S3TagsEnabled
}

// GetS3TagsEnabledOk returns a tuple with the S3TagsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettings) GetS3TagsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.S3TagsEnabled) {
		return nil, false
	}
	return o.S3TagsEnabled, true
}

// HasS3TagsEnabled returns a boolean if a field has been set.
func (o *GeneralSettings) HasS3TagsEnabled() bool {
	if o != nil && !IsNil(o.S3TagsEnabled) {
		return true
	}

	return false
}

// SetS3TagsEnabled gets a reference to the given bool and assigns it to the S3TagsEnabled field.
func (o *GeneralSettings) SetS3TagsEnabled(v bool) {
	o.S3TagsEnabled = &v
}

// GetHideLoginInputFields returns the HideLoginInputFields field value if set, zero value otherwise.
// Deprecated
func (o *GeneralSettings) GetHideLoginInputFields() bool {
	if o == nil || IsNil(o.HideLoginInputFields) {
		var ret bool
		return ret
	}
	return *o.HideLoginInputFields
}

// GetHideLoginInputFieldsOk returns a tuple with the HideLoginInputFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GeneralSettings) GetHideLoginInputFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.HideLoginInputFields) {
		return nil, false
	}
	return o.HideLoginInputFields, true
}

// HasHideLoginInputFields returns a boolean if a field has been set.
func (o *GeneralSettings) HasHideLoginInputFields() bool {
	if o != nil && !IsNil(o.HideLoginInputFields) {
		return true
	}

	return false
}

// SetHideLoginInputFields gets a reference to the given bool and assigns it to the HideLoginInputFields field.
// Deprecated
func (o *GeneralSettings) SetHideLoginInputFields(v bool) {
	o.HideLoginInputFields = &v
}

// GetAuthTokenRestrictions returns the AuthTokenRestrictions field value if set, zero value otherwise.
func (o *GeneralSettings) GetAuthTokenRestrictions() AuthTokenRestrictions {
	if o == nil || IsNil(o.AuthTokenRestrictions) {
		var ret AuthTokenRestrictions
		return ret
	}
	return *o.AuthTokenRestrictions
}

// GetAuthTokenRestrictionsOk returns a tuple with the AuthTokenRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettings) GetAuthTokenRestrictionsOk() (*AuthTokenRestrictions, bool) {
	if o == nil || IsNil(o.AuthTokenRestrictions) {
		return nil, false
	}
	return o.AuthTokenRestrictions, true
}

// HasAuthTokenRestrictions returns a boolean if a field has been set.
func (o *GeneralSettings) HasAuthTokenRestrictions() bool {
	if o != nil && !IsNil(o.AuthTokenRestrictions) {
		return true
	}

	return false
}

// SetAuthTokenRestrictions gets a reference to the given AuthTokenRestrictions and assigns it to the AuthTokenRestrictions field.
func (o *GeneralSettings) SetAuthTokenRestrictions(v AuthTokenRestrictions) {
	o.AuthTokenRestrictions = &v
}

// GetMediaServerEnabled returns the MediaServerEnabled field value if set, zero value otherwise.
// Deprecated
func (o *GeneralSettings) GetMediaServerEnabled() bool {
	if o == nil || IsNil(o.MediaServerEnabled) {
		var ret bool
		return ret
	}
	return *o.MediaServerEnabled
}

// GetMediaServerEnabledOk returns a tuple with the MediaServerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GeneralSettings) GetMediaServerEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MediaServerEnabled) {
		return nil, false
	}
	return o.MediaServerEnabled, true
}

// HasMediaServerEnabled returns a boolean if a field has been set.
func (o *GeneralSettings) HasMediaServerEnabled() bool {
	if o != nil && !IsNil(o.MediaServerEnabled) {
		return true
	}

	return false
}

// SetMediaServerEnabled gets a reference to the given bool and assigns it to the MediaServerEnabled field.
// Deprecated
func (o *GeneralSettings) SetMediaServerEnabled(v bool) {
	o.MediaServerEnabled = &v
}

// GetWeakPasswordEnabled returns the WeakPasswordEnabled field value if set, zero value otherwise.
// Deprecated
func (o *GeneralSettings) GetWeakPasswordEnabled() bool {
	if o == nil || IsNil(o.WeakPasswordEnabled) {
		var ret bool
		return ret
	}
	return *o.WeakPasswordEnabled
}

// GetWeakPasswordEnabledOk returns a tuple with the WeakPasswordEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GeneralSettings) GetWeakPasswordEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WeakPasswordEnabled) {
		return nil, false
	}
	return o.WeakPasswordEnabled, true
}

// HasWeakPasswordEnabled returns a boolean if a field has been set.
func (o *GeneralSettings) HasWeakPasswordEnabled() bool {
	if o != nil && !IsNil(o.WeakPasswordEnabled) {
		return true
	}

	return false
}

// SetWeakPasswordEnabled gets a reference to the given bool and assigns it to the WeakPasswordEnabled field.
// Deprecated
func (o *GeneralSettings) SetWeakPasswordEnabled(v bool) {
	o.WeakPasswordEnabled = &v
}

func (o GeneralSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeneralSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SharePasswordSmsEnabled) {
		toSerialize["sharePasswordSmsEnabled"] = o.SharePasswordSmsEnabled
	}
	if !IsNil(o.CryptoEnabled) {
		toSerialize["cryptoEnabled"] = o.CryptoEnabled
	}
	if !IsNil(o.EmailNotificationButtonEnabled) {
		toSerialize["emailNotificationButtonEnabled"] = o.EmailNotificationButtonEnabled
	}
	if !IsNil(o.EulaEnabled) {
		toSerialize["eulaEnabled"] = o.EulaEnabled
	}
	if !IsNil(o.UseS3Storage) {
		toSerialize["useS3Storage"] = o.UseS3Storage
	}
	if !IsNil(o.S3TagsEnabled) {
		toSerialize["s3TagsEnabled"] = o.S3TagsEnabled
	}
	if !IsNil(o.HideLoginInputFields) {
		toSerialize["hideLoginInputFields"] = o.HideLoginInputFields
	}
	if !IsNil(o.AuthTokenRestrictions) {
		toSerialize["authTokenRestrictions"] = o.AuthTokenRestrictions
	}
	if !IsNil(o.MediaServerEnabled) {
		toSerialize["mediaServerEnabled"] = o.MediaServerEnabled
	}
	if !IsNil(o.WeakPasswordEnabled) {
		toSerialize["weakPasswordEnabled"] = o.WeakPasswordEnabled
	}
	return toSerialize, nil
}

type NullableGeneralSettings struct {
	value *GeneralSettings
	isSet bool
}

func (v NullableGeneralSettings) Get() *GeneralSettings {
	return v.value
}

func (v *NullableGeneralSettings) Set(val *GeneralSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralSettings(val *GeneralSettings) *NullableGeneralSettings {
	return &NullableGeneralSettings{value: val, isSet: true}
}

func (v NullableGeneralSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
