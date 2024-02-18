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

// checks if the GeneralSettingsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneralSettingsInfo{}

// GeneralSettingsInfo General settings
type GeneralSettingsInfo struct {
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
	// &#128640; Since v4.10.0  Homerooms active
	HomeRoomsActive bool `json:"homeRoomsActive"`
	// &#128640; Since v4.10.0  Homeroom Parent ID
	HomeRoomParentId *int64 `json:"homeRoomParentId,omitempty"`
	// &#128640; Since v4.30.0  Subscription Plan
	SubscriptionPlan int32 `json:"subscriptionPlan"`
}

type _GeneralSettingsInfo GeneralSettingsInfo

// NewGeneralSettingsInfo instantiates a new GeneralSettingsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneralSettingsInfo(homeRoomsActive bool, subscriptionPlan int32) *GeneralSettingsInfo {
	this := GeneralSettingsInfo{}
	this.HomeRoomsActive = homeRoomsActive
	this.SubscriptionPlan = subscriptionPlan
	return &this
}

// NewGeneralSettingsInfoWithDefaults instantiates a new GeneralSettingsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralSettingsInfoWithDefaults() *GeneralSettingsInfo {
	this := GeneralSettingsInfo{}
	return &this
}

// GetSharePasswordSmsEnabled returns the SharePasswordSmsEnabled field value if set, zero value otherwise.
func (o *GeneralSettingsInfo) GetSharePasswordSmsEnabled() bool {
	if o == nil || IsNil(o.SharePasswordSmsEnabled) {
		var ret bool
		return ret
	}
	return *o.SharePasswordSmsEnabled
}

// GetSharePasswordSmsEnabledOk returns a tuple with the SharePasswordSmsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettingsInfo) GetSharePasswordSmsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SharePasswordSmsEnabled) {
		return nil, false
	}
	return o.SharePasswordSmsEnabled, true
}

// HasSharePasswordSmsEnabled returns a boolean if a field has been set.
func (o *GeneralSettingsInfo) HasSharePasswordSmsEnabled() bool {
	if o != nil && !IsNil(o.SharePasswordSmsEnabled) {
		return true
	}

	return false
}

// SetSharePasswordSmsEnabled gets a reference to the given bool and assigns it to the SharePasswordSmsEnabled field.
func (o *GeneralSettingsInfo) SetSharePasswordSmsEnabled(v bool) {
	o.SharePasswordSmsEnabled = &v
}

// GetCryptoEnabled returns the CryptoEnabled field value if set, zero value otherwise.
func (o *GeneralSettingsInfo) GetCryptoEnabled() bool {
	if o == nil || IsNil(o.CryptoEnabled) {
		var ret bool
		return ret
	}
	return *o.CryptoEnabled
}

// GetCryptoEnabledOk returns a tuple with the CryptoEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettingsInfo) GetCryptoEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CryptoEnabled) {
		return nil, false
	}
	return o.CryptoEnabled, true
}

// HasCryptoEnabled returns a boolean if a field has been set.
func (o *GeneralSettingsInfo) HasCryptoEnabled() bool {
	if o != nil && !IsNil(o.CryptoEnabled) {
		return true
	}

	return false
}

// SetCryptoEnabled gets a reference to the given bool and assigns it to the CryptoEnabled field.
func (o *GeneralSettingsInfo) SetCryptoEnabled(v bool) {
	o.CryptoEnabled = &v
}

// GetEmailNotificationButtonEnabled returns the EmailNotificationButtonEnabled field value if set, zero value otherwise.
func (o *GeneralSettingsInfo) GetEmailNotificationButtonEnabled() bool {
	if o == nil || IsNil(o.EmailNotificationButtonEnabled) {
		var ret bool
		return ret
	}
	return *o.EmailNotificationButtonEnabled
}

// GetEmailNotificationButtonEnabledOk returns a tuple with the EmailNotificationButtonEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettingsInfo) GetEmailNotificationButtonEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailNotificationButtonEnabled) {
		return nil, false
	}
	return o.EmailNotificationButtonEnabled, true
}

// HasEmailNotificationButtonEnabled returns a boolean if a field has been set.
func (o *GeneralSettingsInfo) HasEmailNotificationButtonEnabled() bool {
	if o != nil && !IsNil(o.EmailNotificationButtonEnabled) {
		return true
	}

	return false
}

// SetEmailNotificationButtonEnabled gets a reference to the given bool and assigns it to the EmailNotificationButtonEnabled field.
func (o *GeneralSettingsInfo) SetEmailNotificationButtonEnabled(v bool) {
	o.EmailNotificationButtonEnabled = &v
}

// GetEulaEnabled returns the EulaEnabled field value if set, zero value otherwise.
func (o *GeneralSettingsInfo) GetEulaEnabled() bool {
	if o == nil || IsNil(o.EulaEnabled) {
		var ret bool
		return ret
	}
	return *o.EulaEnabled
}

// GetEulaEnabledOk returns a tuple with the EulaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettingsInfo) GetEulaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EulaEnabled) {
		return nil, false
	}
	return o.EulaEnabled, true
}

// HasEulaEnabled returns a boolean if a field has been set.
func (o *GeneralSettingsInfo) HasEulaEnabled() bool {
	if o != nil && !IsNil(o.EulaEnabled) {
		return true
	}

	return false
}

// SetEulaEnabled gets a reference to the given bool and assigns it to the EulaEnabled field.
func (o *GeneralSettingsInfo) SetEulaEnabled(v bool) {
	o.EulaEnabled = &v
}

// GetUseS3Storage returns the UseS3Storage field value if set, zero value otherwise.
func (o *GeneralSettingsInfo) GetUseS3Storage() bool {
	if o == nil || IsNil(o.UseS3Storage) {
		var ret bool
		return ret
	}
	return *o.UseS3Storage
}

// GetUseS3StorageOk returns a tuple with the UseS3Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettingsInfo) GetUseS3StorageOk() (*bool, bool) {
	if o == nil || IsNil(o.UseS3Storage) {
		return nil, false
	}
	return o.UseS3Storage, true
}

// HasUseS3Storage returns a boolean if a field has been set.
func (o *GeneralSettingsInfo) HasUseS3Storage() bool {
	if o != nil && !IsNil(o.UseS3Storage) {
		return true
	}

	return false
}

// SetUseS3Storage gets a reference to the given bool and assigns it to the UseS3Storage field.
func (o *GeneralSettingsInfo) SetUseS3Storage(v bool) {
	o.UseS3Storage = &v
}

// GetS3TagsEnabled returns the S3TagsEnabled field value if set, zero value otherwise.
func (o *GeneralSettingsInfo) GetS3TagsEnabled() bool {
	if o == nil || IsNil(o.S3TagsEnabled) {
		var ret bool
		return ret
	}
	return *o.S3TagsEnabled
}

// GetS3TagsEnabledOk returns a tuple with the S3TagsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettingsInfo) GetS3TagsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.S3TagsEnabled) {
		return nil, false
	}
	return o.S3TagsEnabled, true
}

// HasS3TagsEnabled returns a boolean if a field has been set.
func (o *GeneralSettingsInfo) HasS3TagsEnabled() bool {
	if o != nil && !IsNil(o.S3TagsEnabled) {
		return true
	}

	return false
}

// SetS3TagsEnabled gets a reference to the given bool and assigns it to the S3TagsEnabled field.
func (o *GeneralSettingsInfo) SetS3TagsEnabled(v bool) {
	o.S3TagsEnabled = &v
}

// GetHideLoginInputFields returns the HideLoginInputFields field value if set, zero value otherwise.
// Deprecated
func (o *GeneralSettingsInfo) GetHideLoginInputFields() bool {
	if o == nil || IsNil(o.HideLoginInputFields) {
		var ret bool
		return ret
	}
	return *o.HideLoginInputFields
}

// GetHideLoginInputFieldsOk returns a tuple with the HideLoginInputFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GeneralSettingsInfo) GetHideLoginInputFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.HideLoginInputFields) {
		return nil, false
	}
	return o.HideLoginInputFields, true
}

// HasHideLoginInputFields returns a boolean if a field has been set.
func (o *GeneralSettingsInfo) HasHideLoginInputFields() bool {
	if o != nil && !IsNil(o.HideLoginInputFields) {
		return true
	}

	return false
}

// SetHideLoginInputFields gets a reference to the given bool and assigns it to the HideLoginInputFields field.
// Deprecated
func (o *GeneralSettingsInfo) SetHideLoginInputFields(v bool) {
	o.HideLoginInputFields = &v
}

// GetAuthTokenRestrictions returns the AuthTokenRestrictions field value if set, zero value otherwise.
func (o *GeneralSettingsInfo) GetAuthTokenRestrictions() AuthTokenRestrictions {
	if o == nil || IsNil(o.AuthTokenRestrictions) {
		var ret AuthTokenRestrictions
		return ret
	}
	return *o.AuthTokenRestrictions
}

// GetAuthTokenRestrictionsOk returns a tuple with the AuthTokenRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettingsInfo) GetAuthTokenRestrictionsOk() (*AuthTokenRestrictions, bool) {
	if o == nil || IsNil(o.AuthTokenRestrictions) {
		return nil, false
	}
	return o.AuthTokenRestrictions, true
}

// HasAuthTokenRestrictions returns a boolean if a field has been set.
func (o *GeneralSettingsInfo) HasAuthTokenRestrictions() bool {
	if o != nil && !IsNil(o.AuthTokenRestrictions) {
		return true
	}

	return false
}

// SetAuthTokenRestrictions gets a reference to the given AuthTokenRestrictions and assigns it to the AuthTokenRestrictions field.
func (o *GeneralSettingsInfo) SetAuthTokenRestrictions(v AuthTokenRestrictions) {
	o.AuthTokenRestrictions = &v
}

// GetMediaServerEnabled returns the MediaServerEnabled field value if set, zero value otherwise.
// Deprecated
func (o *GeneralSettingsInfo) GetMediaServerEnabled() bool {
	if o == nil || IsNil(o.MediaServerEnabled) {
		var ret bool
		return ret
	}
	return *o.MediaServerEnabled
}

// GetMediaServerEnabledOk returns a tuple with the MediaServerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GeneralSettingsInfo) GetMediaServerEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MediaServerEnabled) {
		return nil, false
	}
	return o.MediaServerEnabled, true
}

// HasMediaServerEnabled returns a boolean if a field has been set.
func (o *GeneralSettingsInfo) HasMediaServerEnabled() bool {
	if o != nil && !IsNil(o.MediaServerEnabled) {
		return true
	}

	return false
}

// SetMediaServerEnabled gets a reference to the given bool and assigns it to the MediaServerEnabled field.
// Deprecated
func (o *GeneralSettingsInfo) SetMediaServerEnabled(v bool) {
	o.MediaServerEnabled = &v
}

// GetWeakPasswordEnabled returns the WeakPasswordEnabled field value if set, zero value otherwise.
// Deprecated
func (o *GeneralSettingsInfo) GetWeakPasswordEnabled() bool {
	if o == nil || IsNil(o.WeakPasswordEnabled) {
		var ret bool
		return ret
	}
	return *o.WeakPasswordEnabled
}

// GetWeakPasswordEnabledOk returns a tuple with the WeakPasswordEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GeneralSettingsInfo) GetWeakPasswordEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WeakPasswordEnabled) {
		return nil, false
	}
	return o.WeakPasswordEnabled, true
}

// HasWeakPasswordEnabled returns a boolean if a field has been set.
func (o *GeneralSettingsInfo) HasWeakPasswordEnabled() bool {
	if o != nil && !IsNil(o.WeakPasswordEnabled) {
		return true
	}

	return false
}

// SetWeakPasswordEnabled gets a reference to the given bool and assigns it to the WeakPasswordEnabled field.
// Deprecated
func (o *GeneralSettingsInfo) SetWeakPasswordEnabled(v bool) {
	o.WeakPasswordEnabled = &v
}

// GetHomeRoomsActive returns the HomeRoomsActive field value
func (o *GeneralSettingsInfo) GetHomeRoomsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HomeRoomsActive
}

// GetHomeRoomsActiveOk returns a tuple with the HomeRoomsActive field value
// and a boolean to check if the value has been set.
func (o *GeneralSettingsInfo) GetHomeRoomsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomeRoomsActive, true
}

// SetHomeRoomsActive sets field value
func (o *GeneralSettingsInfo) SetHomeRoomsActive(v bool) {
	o.HomeRoomsActive = v
}

// GetHomeRoomParentId returns the HomeRoomParentId field value if set, zero value otherwise.
func (o *GeneralSettingsInfo) GetHomeRoomParentId() int64 {
	if o == nil || IsNil(o.HomeRoomParentId) {
		var ret int64
		return ret
	}
	return *o.HomeRoomParentId
}

// GetHomeRoomParentIdOk returns a tuple with the HomeRoomParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSettingsInfo) GetHomeRoomParentIdOk() (*int64, bool) {
	if o == nil || IsNil(o.HomeRoomParentId) {
		return nil, false
	}
	return o.HomeRoomParentId, true
}

// HasHomeRoomParentId returns a boolean if a field has been set.
func (o *GeneralSettingsInfo) HasHomeRoomParentId() bool {
	if o != nil && !IsNil(o.HomeRoomParentId) {
		return true
	}

	return false
}

// SetHomeRoomParentId gets a reference to the given int64 and assigns it to the HomeRoomParentId field.
func (o *GeneralSettingsInfo) SetHomeRoomParentId(v int64) {
	o.HomeRoomParentId = &v
}

// GetSubscriptionPlan returns the SubscriptionPlan field value
func (o *GeneralSettingsInfo) GetSubscriptionPlan() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SubscriptionPlan
}

// GetSubscriptionPlanOk returns a tuple with the SubscriptionPlan field value
// and a boolean to check if the value has been set.
func (o *GeneralSettingsInfo) GetSubscriptionPlanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionPlan, true
}

// SetSubscriptionPlan sets field value
func (o *GeneralSettingsInfo) SetSubscriptionPlan(v int32) {
	o.SubscriptionPlan = v
}

func (o GeneralSettingsInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeneralSettingsInfo) ToMap() (map[string]interface{}, error) {
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
	toSerialize["homeRoomsActive"] = o.HomeRoomsActive
	if !IsNil(o.HomeRoomParentId) {
		toSerialize["homeRoomParentId"] = o.HomeRoomParentId
	}
	toSerialize["subscriptionPlan"] = o.SubscriptionPlan
	return toSerialize, nil
}

func (o *GeneralSettingsInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"homeRoomsActive",
		"subscriptionPlan",
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

	varGeneralSettingsInfo := _GeneralSettingsInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGeneralSettingsInfo)

	if err != nil {
		return err
	}

	*o = GeneralSettingsInfo(varGeneralSettingsInfo)

	return err
}

type NullableGeneralSettingsInfo struct {
	value *GeneralSettingsInfo
	isSet bool
}

func (v NullableGeneralSettingsInfo) Get() *GeneralSettingsInfo {
	return v.value
}

func (v *NullableGeneralSettingsInfo) Set(val *GeneralSettingsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralSettingsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralSettingsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralSettingsInfo(val *GeneralSettingsInfo) *NullableGeneralSettingsInfo {
	return &NullableGeneralSettingsInfo{value: val, isSet: true}
}

func (v NullableGeneralSettingsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralSettingsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
