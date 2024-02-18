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

// checks if the SyslogConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyslogConfig{}

// SyslogConfig Syslog settings
type SyslogConfig struct {
	// Is syslog enabled?
	Enabled *bool `json:"enabled,omitempty"`
	// Syslog server (IP or FQDN)
	Host *string `json:"host,omitempty"`
	// Syslog server port
	Port *int32 `json:"port,omitempty"`
	// Protocol to connect to syslog server
	Protocol *string `json:"protocol,omitempty"`
	// Determines whether user’s IP address is logged.
	LogIpEnabled *bool `json:"logIpEnabled,omitempty"`
}

// NewSyslogConfig instantiates a new SyslogConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogConfig() *SyslogConfig {
	this := SyslogConfig{}
	return &this
}

// NewSyslogConfigWithDefaults instantiates a new SyslogConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogConfigWithDefaults() *SyslogConfig {
	this := SyslogConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SyslogConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SyslogConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SyslogConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *SyslogConfig) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogConfig) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *SyslogConfig) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *SyslogConfig) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SyslogConfig) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogConfig) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SyslogConfig) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *SyslogConfig) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SyslogConfig) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogConfig) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SyslogConfig) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *SyslogConfig) SetProtocol(v string) {
	o.Protocol = &v
}

// GetLogIpEnabled returns the LogIpEnabled field value if set, zero value otherwise.
func (o *SyslogConfig) GetLogIpEnabled() bool {
	if o == nil || IsNil(o.LogIpEnabled) {
		var ret bool
		return ret
	}
	return *o.LogIpEnabled
}

// GetLogIpEnabledOk returns a tuple with the LogIpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogConfig) GetLogIpEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LogIpEnabled) {
		return nil, false
	}
	return o.LogIpEnabled, true
}

// HasLogIpEnabled returns a boolean if a field has been set.
func (o *SyslogConfig) HasLogIpEnabled() bool {
	if o != nil && !IsNil(o.LogIpEnabled) {
		return true
	}

	return false
}

// SetLogIpEnabled gets a reference to the given bool and assigns it to the LogIpEnabled field.
func (o *SyslogConfig) SetLogIpEnabled(v bool) {
	o.LogIpEnabled = &v
}

func (o SyslogConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyslogConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.LogIpEnabled) {
		toSerialize["logIpEnabled"] = o.LogIpEnabled
	}
	return toSerialize, nil
}

type NullableSyslogConfig struct {
	value *SyslogConfig
	isSet bool
}

func (v NullableSyslogConfig) Get() *SyslogConfig {
	return v.value
}

func (v *NullableSyslogConfig) Set(val *SyslogConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogConfig(val *SyslogConfig) *NullableSyslogConfig {
	return &NullableSyslogConfig{value: val, isSet: true}
}

func (v NullableSyslogConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
