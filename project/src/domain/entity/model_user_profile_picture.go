/*
Spot

A sample API that uses a petstore as an example to demonstrate features in the OpenAPI 3.0 specification

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package entity

import (
	"encoding/json"
)

// checks if the UserProfilePicture type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProfilePicture{}

// UserProfilePicture struct for UserProfilePicture
type UserProfilePicture struct {
	Url *string `json:"url,omitempty" firestore:"url"`
	Path *string `json:"path,omitempty" firestore:"path"`
}

// NewUserProfilePicture instantiates a new UserProfilePicture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfilePicture() *UserProfilePicture {
	this := UserProfilePicture{}
	return &this
}

// NewUserProfilePictureWithDefaults instantiates a new UserProfilePicture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfilePictureWithDefaults() *UserProfilePicture {
	this := UserProfilePicture{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UserProfilePicture) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfilePicture) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UserProfilePicture) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UserProfilePicture) SetUrl(v string) {
	o.Url = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *UserProfilePicture) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfilePicture) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *UserProfilePicture) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *UserProfilePicture) SetPath(v string) {
	o.Path = &v
}

func (o UserProfilePicture) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProfilePicture) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableUserProfilePicture struct {
	value *UserProfilePicture
	isSet bool
}

func (v NullableUserProfilePicture) Get() *UserProfilePicture {
	return v.value
}

func (v *NullableUserProfilePicture) Set(val *UserProfilePicture) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfilePicture) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfilePicture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfilePicture(val *UserProfilePicture) *NullableUserProfilePicture {
	return &NullableUserProfilePicture{value: val, isSet: true}
}

func (v NullableUserProfilePicture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfilePicture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

