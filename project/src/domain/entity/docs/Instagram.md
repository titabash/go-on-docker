# Instagram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Types** | Pointer to **[]string** |  | [optional] 
**Reviews** | Pointer to [**[]Review**](Review.md) |  | [optional] 
**Rating** | Pointer to **float64** |  | [optional] 
**Photos** | Pointer to [**[]Photo**](Photo.md) |  | [optional] 

## Methods

### NewInstagram

`func NewInstagram() *Instagram`

NewInstagram instantiates a new Instagram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstagramWithDefaults

`func NewInstagramWithDefaults() *Instagram`

NewInstagramWithDefaults instantiates a new Instagram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypes

`func (o *Instagram) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *Instagram) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *Instagram) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *Instagram) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetReviews

`func (o *Instagram) GetReviews() []Review`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *Instagram) GetReviewsOk() (*[]Review, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *Instagram) SetReviews(v []Review)`

SetReviews sets Reviews field to given value.

### HasReviews

`func (o *Instagram) HasReviews() bool`

HasReviews returns a boolean if a field has been set.

### GetRating

`func (o *Instagram) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Instagram) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Instagram) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *Instagram) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetPhotos

`func (o *Instagram) GetPhotos() []Photo`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *Instagram) GetPhotosOk() (*[]Photo, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *Instagram) SetPhotos(v []Photo)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *Instagram) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


