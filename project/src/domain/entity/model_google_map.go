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

// checks if the GoogleMap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleMap{}

// GoogleMap struct for GoogleMap
type GoogleMap struct {
	PlaceId *string `json:"place_id,omitempty" firestore:"place_id"`
	Types []string `json:"types,omitempty" firestore:"types"`
	Reviews []Review `json:"reviews,omitempty" firestore:"reviews"`
	Rating *float64 `json:"rating,omitempty" firestore:"rating"`
	Photos []Photo `json:"photos,omitempty" firestore:"photos"`
}

// NewGoogleMap instantiates a new GoogleMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleMap() *GoogleMap {
	this := GoogleMap{}
	return &this
}

// NewGoogleMapWithDefaults instantiates a new GoogleMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleMapWithDefaults() *GoogleMap {
	this := GoogleMap{}
	return &this
}

// GetPlaceId returns the PlaceId field value if set, zero value otherwise.
func (o *GoogleMap) GetPlaceId() string {
	if o == nil || IsNil(o.PlaceId) {
		var ret string
		return ret
	}
	return *o.PlaceId
}

// GetPlaceIdOk returns a tuple with the PlaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleMap) GetPlaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaceId) {
		return nil, false
	}
	return o.PlaceId, true
}

// HasPlaceId returns a boolean if a field has been set.
func (o *GoogleMap) HasPlaceId() bool {
	if o != nil && !IsNil(o.PlaceId) {
		return true
	}

	return false
}

// SetPlaceId gets a reference to the given string and assigns it to the PlaceId field.
func (o *GoogleMap) SetPlaceId(v string) {
	o.PlaceId = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *GoogleMap) GetTypes() []string {
	if o == nil || IsNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleMap) GetTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *GoogleMap) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *GoogleMap) SetTypes(v []string) {
	o.Types = v
}

// GetReviews returns the Reviews field value if set, zero value otherwise.
func (o *GoogleMap) GetReviews() []Review {
	if o == nil || IsNil(o.Reviews) {
		var ret []Review
		return ret
	}
	return o.Reviews
}

// GetReviewsOk returns a tuple with the Reviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleMap) GetReviewsOk() ([]Review, bool) {
	if o == nil || IsNil(o.Reviews) {
		return nil, false
	}
	return o.Reviews, true
}

// HasReviews returns a boolean if a field has been set.
func (o *GoogleMap) HasReviews() bool {
	if o != nil && !IsNil(o.Reviews) {
		return true
	}

	return false
}

// SetReviews gets a reference to the given []Review and assigns it to the Reviews field.
func (o *GoogleMap) SetReviews(v []Review) {
	o.Reviews = v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *GoogleMap) GetRating() float64 {
	if o == nil || IsNil(o.Rating) {
		var ret float64
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleMap) GetRatingOk() (*float64, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *GoogleMap) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given float64 and assigns it to the Rating field.
func (o *GoogleMap) SetRating(v float64) {
	o.Rating = &v
}

// GetPhotos returns the Photos field value if set, zero value otherwise.
func (o *GoogleMap) GetPhotos() []Photo {
	if o == nil || IsNil(o.Photos) {
		var ret []Photo
		return ret
	}
	return o.Photos
}

// GetPhotosOk returns a tuple with the Photos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleMap) GetPhotosOk() ([]Photo, bool) {
	if o == nil || IsNil(o.Photos) {
		return nil, false
	}
	return o.Photos, true
}

// HasPhotos returns a boolean if a field has been set.
func (o *GoogleMap) HasPhotos() bool {
	if o != nil && !IsNil(o.Photos) {
		return true
	}

	return false
}

// SetPhotos gets a reference to the given []Photo and assigns it to the Photos field.
func (o *GoogleMap) SetPhotos(v []Photo) {
	o.Photos = v
}

func (o GoogleMap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleMap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlaceId) {
		toSerialize["place_id"] = o.PlaceId
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.Reviews) {
		toSerialize["reviews"] = o.Reviews
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.Photos) {
		toSerialize["photos"] = o.Photos
	}
	return toSerialize, nil
}

type NullableGoogleMap struct {
	value *GoogleMap
	isSet bool
}

func (v NullableGoogleMap) Get() *GoogleMap {
	return v.value
}

func (v *NullableGoogleMap) Set(val *GoogleMap) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleMap) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleMap(val *GoogleMap) *NullableGoogleMap {
	return &NullableGoogleMap{value: val, isSet: true}
}

func (v NullableGoogleMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


