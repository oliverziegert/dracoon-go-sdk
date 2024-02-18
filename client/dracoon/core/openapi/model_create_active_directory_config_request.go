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

// checks if the CreateActiveDirectoryConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateActiveDirectoryConfigRequest{}

// CreateActiveDirectoryConfigRequest Request model for creating an Active Directory configuration
type CreateActiveDirectoryConfigRequest struct {
	// Unique name for an Active Directory configuration
	Alias string `json:"alias"`
	// IPv4 or IPv6 address or host name
	ServerIp string `json:"serverIp"`
	// Port
	ServerPort int32 `json:"serverPort"`
	// Distinguished Name (DN) of Active Directory administrative account
	ServerAdminName string `json:"serverAdminName"`
	// Password of Active Directory administrative account
	ServerAdminPassword string `json:"serverAdminPassword"`
	// Search scope of Active Directory; only users below this node can log on.
	LdapUsersDomain string `json:"ldapUsersDomain"`
	// Name of Active Directory attribute that is used as login name.
	UserFilter string `json:"userFilter"`
	// Determines if a DRACOON account is automatically created for a new user  who successfully logs on with his / her AD / IDP account.
	UserImport *bool `json:"userImport,omitempty"`
	// Determines whether LDAPS should be used instead of plain LDAP.
	UseLdaps *bool `json:"useLdaps,omitempty"`
	// If `userImport` is set to `true`,  the user must be member of this Active Directory group to receive a newly created DRACOON account.
	AdExportGroup *string `json:"adExportGroup,omitempty"`
	// User group that is assigned to users who are created by automatic import.  Reset with `0`
	SdsImportGroup *int64 `json:"sdsImportGroup,omitempty"`
	// DEPRECATED, will be ignored  Determines whether a room is created for each user that is created by automatic import (like a home folder).  Room's name will equal the user's login name.
	CreateHomeFolder *bool `json:"createHomeFolder,omitempty"`
	// DEPRECATED, will be ignored  ID of the room in which the individual rooms for users will be created.
	HomeFolderParent *int64 `json:"homeFolderParent,omitempty"`
	// SSL finger print of Active Directory server.  Mandatory for LDAPS connections.  Format: `Algorithm/Fingerprint`
	SslFingerPrint *string `json:"sslFingerPrint,omitempty"`
}

type _CreateActiveDirectoryConfigRequest CreateActiveDirectoryConfigRequest

// NewCreateActiveDirectoryConfigRequest instantiates a new CreateActiveDirectoryConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateActiveDirectoryConfigRequest(alias string, serverIp string, serverPort int32, serverAdminName string, serverAdminPassword string, ldapUsersDomain string, userFilter string) *CreateActiveDirectoryConfigRequest {
	this := CreateActiveDirectoryConfigRequest{}
	this.Alias = alias
	this.ServerIp = serverIp
	this.ServerPort = serverPort
	this.ServerAdminName = serverAdminName
	this.ServerAdminPassword = serverAdminPassword
	this.LdapUsersDomain = ldapUsersDomain
	this.UserFilter = userFilter
	var userImport bool = false
	this.UserImport = &userImport
	var useLdaps bool = false
	this.UseLdaps = &useLdaps
	var createHomeFolder bool = false
	this.CreateHomeFolder = &createHomeFolder
	return &this
}

// NewCreateActiveDirectoryConfigRequestWithDefaults instantiates a new CreateActiveDirectoryConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateActiveDirectoryConfigRequestWithDefaults() *CreateActiveDirectoryConfigRequest {
	this := CreateActiveDirectoryConfigRequest{}
	var userImport bool = false
	this.UserImport = &userImport
	var useLdaps bool = false
	this.UseLdaps = &useLdaps
	var createHomeFolder bool = false
	this.CreateHomeFolder = &createHomeFolder
	return &this
}

// GetAlias returns the Alias field value
func (o *CreateActiveDirectoryConfigRequest) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *CreateActiveDirectoryConfigRequest) SetAlias(v string) {
	o.Alias = v
}

// GetServerIp returns the ServerIp field value
func (o *CreateActiveDirectoryConfigRequest) GetServerIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerIp
}

// GetServerIpOk returns a tuple with the ServerIp field value
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetServerIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerIp, true
}

// SetServerIp sets field value
func (o *CreateActiveDirectoryConfigRequest) SetServerIp(v string) {
	o.ServerIp = v
}

// GetServerPort returns the ServerPort field value
func (o *CreateActiveDirectoryConfigRequest) GetServerPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetServerPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerPort, true
}

// SetServerPort sets field value
func (o *CreateActiveDirectoryConfigRequest) SetServerPort(v int32) {
	o.ServerPort = v
}

// GetServerAdminName returns the ServerAdminName field value
func (o *CreateActiveDirectoryConfigRequest) GetServerAdminName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerAdminName
}

// GetServerAdminNameOk returns a tuple with the ServerAdminName field value
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetServerAdminNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerAdminName, true
}

// SetServerAdminName sets field value
func (o *CreateActiveDirectoryConfigRequest) SetServerAdminName(v string) {
	o.ServerAdminName = v
}

// GetServerAdminPassword returns the ServerAdminPassword field value
func (o *CreateActiveDirectoryConfigRequest) GetServerAdminPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerAdminPassword
}

// GetServerAdminPasswordOk returns a tuple with the ServerAdminPassword field value
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetServerAdminPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerAdminPassword, true
}

// SetServerAdminPassword sets field value
func (o *CreateActiveDirectoryConfigRequest) SetServerAdminPassword(v string) {
	o.ServerAdminPassword = v
}

// GetLdapUsersDomain returns the LdapUsersDomain field value
func (o *CreateActiveDirectoryConfigRequest) GetLdapUsersDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LdapUsersDomain
}

// GetLdapUsersDomainOk returns a tuple with the LdapUsersDomain field value
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetLdapUsersDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LdapUsersDomain, true
}

// SetLdapUsersDomain sets field value
func (o *CreateActiveDirectoryConfigRequest) SetLdapUsersDomain(v string) {
	o.LdapUsersDomain = v
}

// GetUserFilter returns the UserFilter field value
func (o *CreateActiveDirectoryConfigRequest) GetUserFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserFilter
}

// GetUserFilterOk returns a tuple with the UserFilter field value
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetUserFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserFilter, true
}

// SetUserFilter sets field value
func (o *CreateActiveDirectoryConfigRequest) SetUserFilter(v string) {
	o.UserFilter = v
}

// GetUserImport returns the UserImport field value if set, zero value otherwise.
func (o *CreateActiveDirectoryConfigRequest) GetUserImport() bool {
	if o == nil || IsNil(o.UserImport) {
		var ret bool
		return ret
	}
	return *o.UserImport
}

// GetUserImportOk returns a tuple with the UserImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetUserImportOk() (*bool, bool) {
	if o == nil || IsNil(o.UserImport) {
		return nil, false
	}
	return o.UserImport, true
}

// HasUserImport returns a boolean if a field has been set.
func (o *CreateActiveDirectoryConfigRequest) HasUserImport() bool {
	if o != nil && !IsNil(o.UserImport) {
		return true
	}

	return false
}

// SetUserImport gets a reference to the given bool and assigns it to the UserImport field.
func (o *CreateActiveDirectoryConfigRequest) SetUserImport(v bool) {
	o.UserImport = &v
}

// GetUseLdaps returns the UseLdaps field value if set, zero value otherwise.
func (o *CreateActiveDirectoryConfigRequest) GetUseLdaps() bool {
	if o == nil || IsNil(o.UseLdaps) {
		var ret bool
		return ret
	}
	return *o.UseLdaps
}

// GetUseLdapsOk returns a tuple with the UseLdaps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetUseLdapsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseLdaps) {
		return nil, false
	}
	return o.UseLdaps, true
}

// HasUseLdaps returns a boolean if a field has been set.
func (o *CreateActiveDirectoryConfigRequest) HasUseLdaps() bool {
	if o != nil && !IsNil(o.UseLdaps) {
		return true
	}

	return false
}

// SetUseLdaps gets a reference to the given bool and assigns it to the UseLdaps field.
func (o *CreateActiveDirectoryConfigRequest) SetUseLdaps(v bool) {
	o.UseLdaps = &v
}

// GetAdExportGroup returns the AdExportGroup field value if set, zero value otherwise.
func (o *CreateActiveDirectoryConfigRequest) GetAdExportGroup() string {
	if o == nil || IsNil(o.AdExportGroup) {
		var ret string
		return ret
	}
	return *o.AdExportGroup
}

// GetAdExportGroupOk returns a tuple with the AdExportGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetAdExportGroupOk() (*string, bool) {
	if o == nil || IsNil(o.AdExportGroup) {
		return nil, false
	}
	return o.AdExportGroup, true
}

// HasAdExportGroup returns a boolean if a field has been set.
func (o *CreateActiveDirectoryConfigRequest) HasAdExportGroup() bool {
	if o != nil && !IsNil(o.AdExportGroup) {
		return true
	}

	return false
}

// SetAdExportGroup gets a reference to the given string and assigns it to the AdExportGroup field.
func (o *CreateActiveDirectoryConfigRequest) SetAdExportGroup(v string) {
	o.AdExportGroup = &v
}

// GetSdsImportGroup returns the SdsImportGroup field value if set, zero value otherwise.
func (o *CreateActiveDirectoryConfigRequest) GetSdsImportGroup() int64 {
	if o == nil || IsNil(o.SdsImportGroup) {
		var ret int64
		return ret
	}
	return *o.SdsImportGroup
}

// GetSdsImportGroupOk returns a tuple with the SdsImportGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetSdsImportGroupOk() (*int64, bool) {
	if o == nil || IsNil(o.SdsImportGroup) {
		return nil, false
	}
	return o.SdsImportGroup, true
}

// HasSdsImportGroup returns a boolean if a field has been set.
func (o *CreateActiveDirectoryConfigRequest) HasSdsImportGroup() bool {
	if o != nil && !IsNil(o.SdsImportGroup) {
		return true
	}

	return false
}

// SetSdsImportGroup gets a reference to the given int64 and assigns it to the SdsImportGroup field.
func (o *CreateActiveDirectoryConfigRequest) SetSdsImportGroup(v int64) {
	o.SdsImportGroup = &v
}

// GetCreateHomeFolder returns the CreateHomeFolder field value if set, zero value otherwise.
func (o *CreateActiveDirectoryConfigRequest) GetCreateHomeFolder() bool {
	if o == nil || IsNil(o.CreateHomeFolder) {
		var ret bool
		return ret
	}
	return *o.CreateHomeFolder
}

// GetCreateHomeFolderOk returns a tuple with the CreateHomeFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetCreateHomeFolderOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateHomeFolder) {
		return nil, false
	}
	return o.CreateHomeFolder, true
}

// HasCreateHomeFolder returns a boolean if a field has been set.
func (o *CreateActiveDirectoryConfigRequest) HasCreateHomeFolder() bool {
	if o != nil && !IsNil(o.CreateHomeFolder) {
		return true
	}

	return false
}

// SetCreateHomeFolder gets a reference to the given bool and assigns it to the CreateHomeFolder field.
func (o *CreateActiveDirectoryConfigRequest) SetCreateHomeFolder(v bool) {
	o.CreateHomeFolder = &v
}

// GetHomeFolderParent returns the HomeFolderParent field value if set, zero value otherwise.
func (o *CreateActiveDirectoryConfigRequest) GetHomeFolderParent() int64 {
	if o == nil || IsNil(o.HomeFolderParent) {
		var ret int64
		return ret
	}
	return *o.HomeFolderParent
}

// GetHomeFolderParentOk returns a tuple with the HomeFolderParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetHomeFolderParentOk() (*int64, bool) {
	if o == nil || IsNil(o.HomeFolderParent) {
		return nil, false
	}
	return o.HomeFolderParent, true
}

// HasHomeFolderParent returns a boolean if a field has been set.
func (o *CreateActiveDirectoryConfigRequest) HasHomeFolderParent() bool {
	if o != nil && !IsNil(o.HomeFolderParent) {
		return true
	}

	return false
}

// SetHomeFolderParent gets a reference to the given int64 and assigns it to the HomeFolderParent field.
func (o *CreateActiveDirectoryConfigRequest) SetHomeFolderParent(v int64) {
	o.HomeFolderParent = &v
}

// GetSslFingerPrint returns the SslFingerPrint field value if set, zero value otherwise.
func (o *CreateActiveDirectoryConfigRequest) GetSslFingerPrint() string {
	if o == nil || IsNil(o.SslFingerPrint) {
		var ret string
		return ret
	}
	return *o.SslFingerPrint
}

// GetSslFingerPrintOk returns a tuple with the SslFingerPrint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateActiveDirectoryConfigRequest) GetSslFingerPrintOk() (*string, bool) {
	if o == nil || IsNil(o.SslFingerPrint) {
		return nil, false
	}
	return o.SslFingerPrint, true
}

// HasSslFingerPrint returns a boolean if a field has been set.
func (o *CreateActiveDirectoryConfigRequest) HasSslFingerPrint() bool {
	if o != nil && !IsNil(o.SslFingerPrint) {
		return true
	}

	return false
}

// SetSslFingerPrint gets a reference to the given string and assigns it to the SslFingerPrint field.
func (o *CreateActiveDirectoryConfigRequest) SetSslFingerPrint(v string) {
	o.SslFingerPrint = &v
}

func (o CreateActiveDirectoryConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateActiveDirectoryConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alias"] = o.Alias
	toSerialize["serverIp"] = o.ServerIp
	toSerialize["serverPort"] = o.ServerPort
	toSerialize["serverAdminName"] = o.ServerAdminName
	toSerialize["serverAdminPassword"] = o.ServerAdminPassword
	toSerialize["ldapUsersDomain"] = o.LdapUsersDomain
	toSerialize["userFilter"] = o.UserFilter
	if !IsNil(o.UserImport) {
		toSerialize["userImport"] = o.UserImport
	}
	if !IsNil(o.UseLdaps) {
		toSerialize["useLdaps"] = o.UseLdaps
	}
	if !IsNil(o.AdExportGroup) {
		toSerialize["adExportGroup"] = o.AdExportGroup
	}
	if !IsNil(o.SdsImportGroup) {
		toSerialize["sdsImportGroup"] = o.SdsImportGroup
	}
	if !IsNil(o.CreateHomeFolder) {
		toSerialize["createHomeFolder"] = o.CreateHomeFolder
	}
	if !IsNil(o.HomeFolderParent) {
		toSerialize["homeFolderParent"] = o.HomeFolderParent
	}
	if !IsNil(o.SslFingerPrint) {
		toSerialize["sslFingerPrint"] = o.SslFingerPrint
	}
	return toSerialize, nil
}

func (o *CreateActiveDirectoryConfigRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alias",
		"serverIp",
		"serverPort",
		"serverAdminName",
		"serverAdminPassword",
		"ldapUsersDomain",
		"userFilter",
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

	varCreateActiveDirectoryConfigRequest := _CreateActiveDirectoryConfigRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateActiveDirectoryConfigRequest)

	if err != nil {
		return err
	}

	*o = CreateActiveDirectoryConfigRequest(varCreateActiveDirectoryConfigRequest)

	return err
}

type NullableCreateActiveDirectoryConfigRequest struct {
	value *CreateActiveDirectoryConfigRequest
	isSet bool
}

func (v NullableCreateActiveDirectoryConfigRequest) Get() *CreateActiveDirectoryConfigRequest {
	return v.value
}

func (v *NullableCreateActiveDirectoryConfigRequest) Set(val *CreateActiveDirectoryConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateActiveDirectoryConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateActiveDirectoryConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateActiveDirectoryConfigRequest(val *CreateActiveDirectoryConfigRequest) *NullableCreateActiveDirectoryConfigRequest {
	return &NullableCreateActiveDirectoryConfigRequest{value: val, isSet: true}
}

func (v NullableCreateActiveDirectoryConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateActiveDirectoryConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
