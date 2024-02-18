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

// checks if the CreateOAuthClientRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOAuthClientRequest{}

// CreateOAuthClientRequest Request model for creating an OAuth client
type CreateOAuthClientRequest struct {
	// Name, which is shown at the client configuration and authorization.
	ClientName string `json:"clientName"`
	// Authorized grant types  * `authorization_code`  * `implicit`  * `password`  * `client_credentials`  * `refresh_token`    cf. [RFC 6749](https://tools.ietf.org/html/rfc6749)
	GrantTypes []string `json:"grantTypes"`
	// ID of the OAuth client
	ClientId *string `json:"clientId,omitempty"`
	// Secret, which client uses at authentication.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// Determines whether client is a confidential or public client.
	ClientType *string `json:"clientType,omitempty"`
	// URIs, to which a user is redirected after authorization.
	RedirectUris []string `json:"redirectUris"`
	// Validity of the access token in seconds.
	AccessTokenValidity *int32 `json:"accessTokenValidity,omitempty"`
	// Validity of the refresh token in seconds.
	RefreshTokenValidity *int32 `json:"refreshTokenValidity,omitempty"`
	// &#128640; Since v4.22.0  Validity of the approval interval in seconds.
	ApprovalValidity *int32 `json:"approvalValidity,omitempty"`
}

type _CreateOAuthClientRequest CreateOAuthClientRequest

// NewCreateOAuthClientRequest instantiates a new CreateOAuthClientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOAuthClientRequest(clientName string, grantTypes []string, redirectUris []string) *CreateOAuthClientRequest {
	this := CreateOAuthClientRequest{}
	this.ClientName = clientName
	this.GrantTypes = grantTypes
	var clientType string = "confidential"
	this.ClientType = &clientType
	this.RedirectUris = redirectUris
	return &this
}

// NewCreateOAuthClientRequestWithDefaults instantiates a new CreateOAuthClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOAuthClientRequestWithDefaults() *CreateOAuthClientRequest {
	this := CreateOAuthClientRequest{}
	var clientType string = "confidential"
	this.ClientType = &clientType
	return &this
}

// GetClientName returns the ClientName field value
func (o *CreateOAuthClientRequest) GetClientName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetClientNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientName, true
}

// SetClientName sets field value
func (o *CreateOAuthClientRequest) SetClientName(v string) {
	o.ClientName = v
}

// GetGrantTypes returns the GrantTypes field value
func (o *CreateOAuthClientRequest) GetGrantTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetGrantTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// SetGrantTypes sets field value
func (o *CreateOAuthClientRequest) SetGrantTypes(v []string) {
	o.GrantTypes = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreateOAuthClientRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *CreateOAuthClientRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetClientType returns the ClientType field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetClientType() string {
	if o == nil || IsNil(o.ClientType) {
		var ret string
		return ret
	}
	return *o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetClientTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ClientType) {
		return nil, false
	}
	return o.ClientType, true
}

// HasClientType returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasClientType() bool {
	if o != nil && !IsNil(o.ClientType) {
		return true
	}

	return false
}

// SetClientType gets a reference to the given string and assigns it to the ClientType field.
func (o *CreateOAuthClientRequest) SetClientType(v string) {
	o.ClientType = &v
}

// GetRedirectUris returns the RedirectUris field value
func (o *CreateOAuthClientRequest) GetRedirectUris() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetRedirectUrisOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// SetRedirectUris sets field value
func (o *CreateOAuthClientRequest) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetAccessTokenValidity returns the AccessTokenValidity field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetAccessTokenValidity() int32 {
	if o == nil || IsNil(o.AccessTokenValidity) {
		var ret int32
		return ret
	}
	return *o.AccessTokenValidity
}

// GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetAccessTokenValidityOk() (*int32, bool) {
	if o == nil || IsNil(o.AccessTokenValidity) {
		return nil, false
	}
	return o.AccessTokenValidity, true
}

// HasAccessTokenValidity returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasAccessTokenValidity() bool {
	if o != nil && !IsNil(o.AccessTokenValidity) {
		return true
	}

	return false
}

// SetAccessTokenValidity gets a reference to the given int32 and assigns it to the AccessTokenValidity field.
func (o *CreateOAuthClientRequest) SetAccessTokenValidity(v int32) {
	o.AccessTokenValidity = &v
}

// GetRefreshTokenValidity returns the RefreshTokenValidity field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetRefreshTokenValidity() int32 {
	if o == nil || IsNil(o.RefreshTokenValidity) {
		var ret int32
		return ret
	}
	return *o.RefreshTokenValidity
}

// GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetRefreshTokenValidityOk() (*int32, bool) {
	if o == nil || IsNil(o.RefreshTokenValidity) {
		return nil, false
	}
	return o.RefreshTokenValidity, true
}

// HasRefreshTokenValidity returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasRefreshTokenValidity() bool {
	if o != nil && !IsNil(o.RefreshTokenValidity) {
		return true
	}

	return false
}

// SetRefreshTokenValidity gets a reference to the given int32 and assigns it to the RefreshTokenValidity field.
func (o *CreateOAuthClientRequest) SetRefreshTokenValidity(v int32) {
	o.RefreshTokenValidity = &v
}

// GetApprovalValidity returns the ApprovalValidity field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetApprovalValidity() int32 {
	if o == nil || IsNil(o.ApprovalValidity) {
		var ret int32
		return ret
	}
	return *o.ApprovalValidity
}

// GetApprovalValidityOk returns a tuple with the ApprovalValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetApprovalValidityOk() (*int32, bool) {
	if o == nil || IsNil(o.ApprovalValidity) {
		return nil, false
	}
	return o.ApprovalValidity, true
}

// HasApprovalValidity returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasApprovalValidity() bool {
	if o != nil && !IsNil(o.ApprovalValidity) {
		return true
	}

	return false
}

// SetApprovalValidity gets a reference to the given int32 and assigns it to the ApprovalValidity field.
func (o *CreateOAuthClientRequest) SetApprovalValidity(v int32) {
	o.ApprovalValidity = &v
}

func (o CreateOAuthClientRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOAuthClientRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientName"] = o.ClientName
	toSerialize["grantTypes"] = o.GrantTypes
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !IsNil(o.ClientType) {
		toSerialize["clientType"] = o.ClientType
	}
	toSerialize["redirectUris"] = o.RedirectUris
	if !IsNil(o.AccessTokenValidity) {
		toSerialize["accessTokenValidity"] = o.AccessTokenValidity
	}
	if !IsNil(o.RefreshTokenValidity) {
		toSerialize["refreshTokenValidity"] = o.RefreshTokenValidity
	}
	if !IsNil(o.ApprovalValidity) {
		toSerialize["approvalValidity"] = o.ApprovalValidity
	}
	return toSerialize, nil
}

func (o *CreateOAuthClientRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientName",
		"grantTypes",
		"redirectUris",
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

	varCreateOAuthClientRequest := _CreateOAuthClientRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOAuthClientRequest)

	if err != nil {
		return err
	}

	*o = CreateOAuthClientRequest(varCreateOAuthClientRequest)

	return err
}

type NullableCreateOAuthClientRequest struct {
	value *CreateOAuthClientRequest
	isSet bool
}

func (v NullableCreateOAuthClientRequest) Get() *CreateOAuthClientRequest {
	return v.value
}

func (v *NullableCreateOAuthClientRequest) Set(val *CreateOAuthClientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOAuthClientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOAuthClientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOAuthClientRequest(val *CreateOAuthClientRequest) *NullableCreateOAuthClientRequest {
	return &NullableCreateOAuthClientRequest{value: val, isSet: true}
}

func (v NullableCreateOAuthClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOAuthClientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
