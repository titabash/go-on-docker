# GoogleMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaceId** | Pointer to **string** |  | [optional] 
**Types** | Pointer to **[]string** |  | [optional] 
**Reviews** | Pointer to [**[]Review**](Review.md) |  | [optional] 
**Rating** | Pointer to **float64** |  | [optional] 
**Photos** | Pointer to [**[]Photo**](Photo.md) |  | [optional] 

## Methods

### NewGoogleMap

`func NewGoogleMap() *GoogleMap`

NewGoogleMap instantiates a new GoogleMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleMapWithDefaults

`func NewGoogleMapWithDefaults() *GoogleMap`

NewGoogleMapWithDefaults instantiates a new GoogleMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaceId

`func (o *GoogleMap) GetPlaceId() string`

GetPlaceId returns the PlaceId field if non-nil, zero value otherwise.

### GetPlaceIdOk

`func (o *GoogleMap) GetPlaceIdOk() (*string, bool)`

GetPlaceIdOk returns a tuple with the PlaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceId

`func (o *GoogleMap) SetPlaceId(v string)`

SetPlaceId sets PlaceId field to given value.

### HasPlaceId

`func (o *GoogleMap) HasPlaceId() bool`

HasPlaceId returns a boolean if a field has been set.

### GetTypes

`func (o *GoogleMap) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *GoogleMap) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *GoogleMap) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *GoogleMap) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetReviews

`func (o *GoogleMap) GetReviews() []Review`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *GoogleMap) GetReviewsOk() (*[]Review, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *GoogleMap) SetReviews(v []Review)`

SetReviews sets Reviews field to given value.

### HasReviews

`func (o *GoogleMap) HasReviews() bool`

HasReviews returns a boolean if a field has been set.

### GetRating

`func (o *GoogleMap) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *GoogleMap) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *GoogleMap) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *GoogleMap) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetPhotos

`func (o *GoogleMap) GetPhotos() []Photo`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *GoogleMap) GetPhotosOk() (*[]Photo, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *GoogleMap) SetPhotos(v []Photo)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *GoogleMap) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


