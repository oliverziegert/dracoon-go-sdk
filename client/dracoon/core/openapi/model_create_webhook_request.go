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

// checks if the CreateWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWebhookRequest{}

// CreateWebhookRequest Request model for creating a webhook
type CreateWebhookRequest struct {
	// Name
	Name string `json:"name"`
	// List of names of event types
	EventTypeNames []string `json:"eventTypeNames"`
	// URL (must begin with the `HTTPS` scheme)
	Url string `json:"url"`
	// Secret; used for event message signatures
	Secret *string `json:"secret,omitempty"`
	// Is enabled
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// If set to true, an example event is being created
	TriggerExampleEvent *bool `json:"triggerExampleEvent,omitempty"`
}

type _CreateWebhookRequest CreateWebhookRequest

// NewCreateWebhookRequest instantiates a new CreateWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebhookRequest(name string, eventTypeNames []string, url string) *CreateWebhookRequest {
	this := CreateWebhookRequest{}
	this.Name = name
	this.EventTypeNames = eventTypeNames
	this.Url = url
	return &this
}

// NewCreateWebhookRequestWithDefaults instantiates a new CreateWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebhookRequestWithDefaults() *CreateWebhookRequest {
	this := CreateWebhookRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateWebhookRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWebhookRequest) SetName(v string) {
	o.Name = v
}

// GetEventTypeNames returns the EventTypeNames field value
func (o *CreateWebhookRequest) GetEventTypeNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EventTypeNames
}

// GetEventTypeNamesOk returns a tuple with the EventTypeNames field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetEventTypeNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventTypeNames, true
}

// SetEventTypeNames sets field value
func (o *CreateWebhookRequest) SetEventTypeNames(v []string) {
	o.EventTypeNames = v
}

// GetUrl returns the Url field value
func (o *CreateWebhookRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CreateWebhookRequest) SetUrl(v string) {
	o.Url = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreateWebhookRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *CreateWebhookRequest) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetTriggerExampleEvent returns the TriggerExampleEvent field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetTriggerExampleEvent() bool {
	if o == nil || IsNil(o.TriggerExampleEvent) {
		var ret bool
		return ret
	}
	return *o.TriggerExampleEvent
}

// GetTriggerExampleEventOk returns a tuple with the TriggerExampleEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetTriggerExampleEventOk() (*bool, bool) {
	if o == nil || IsNil(o.TriggerExampleEvent) {
		return nil, false
	}
	return o.TriggerExampleEvent, true
}

// HasTriggerExampleEvent returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasTriggerExampleEvent() bool {
	if o != nil && !IsNil(o.TriggerExampleEvent) {
		return true
	}

	return false
}

// SetTriggerExampleEvent gets a reference to the given bool and assigns it to the TriggerExampleEvent field.
func (o *CreateWebhookRequest) SetTriggerExampleEvent(v bool) {
	o.TriggerExampleEvent = &v
}

func (o CreateWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["eventTypeNames"] = o.EventTypeNames
	toSerialize["url"] = o.Url
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !IsNil(o.TriggerExampleEvent) {
		toSerialize["triggerExampleEvent"] = o.TriggerExampleEvent
	}
	return toSerialize, nil
}

func (o *CreateWebhookRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"eventTypeNames",
		"url",
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

	varCreateWebhookRequest := _CreateWebhookRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateWebhookRequest)

	if err != nil {
		return err
	}

	*o = CreateWebhookRequest(varCreateWebhookRequest)

	return err
}

type NullableCreateWebhookRequest struct {
	value *CreateWebhookRequest
	isSet bool
}

func (v NullableCreateWebhookRequest) Get() *CreateWebhookRequest {
	return v.value
}

func (v *NullableCreateWebhookRequest) Set(val *CreateWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebhookRequest(val *CreateWebhookRequest) *NullableCreateWebhookRequest {
	return &NullableCreateWebhookRequest{value: val, isSet: true}
}

func (v NullableCreateWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
