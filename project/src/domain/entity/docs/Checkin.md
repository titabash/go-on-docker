# Checkin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**SpotId** | **string** |  | 
**Photos** | [**[]Photo**](Photo.md) |  | 
**Address** | [**Address**](Address.md) |  | 
**Location** | [**Location**](Location.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCheckin

`func NewCheckin(id string, userId string, spotId string, photos []Photo, address Address, location Location, createdAt time.Time, updatedAt time.Time, ) *Checkin`

NewCheckin instantiates a new Checkin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckinWithDefaults

`func NewCheckinWithDefaults() *Checkin`

NewCheckinWithDefaults instantiates a new Checkin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Checkin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Checkin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Checkin) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *Checkin) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Checkin) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Checkin) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetSpotId

`func (o *Checkin) GetSpotId() string`

GetSpotId returns the SpotId field if non-nil, zero value otherwise.

### GetSpotIdOk

`func (o *Checkin) GetSpotIdOk() (*string, bool)`

GetSpotIdOk returns a tuple with the SpotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotId

`func (o *Checkin) SetSpotId(v string)`

SetSpotId sets SpotId field to given value.


### GetPhotos

`func (o *Checkin) GetPhotos() []Photo`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *Checkin) GetPhotosOk() (*[]Photo, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *Checkin) SetPhotos(v []Photo)`

SetPhotos sets Photos field to given value.


### GetAddress

`func (o *Checkin) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Checkin) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Checkin) SetAddress(v Address)`

SetAddress sets Address field to given value.


### GetLocation

`func (o *Checkin) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Checkin) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Checkin) SetLocation(v Location)`

SetLocation sets Location field to given value.


### GetCreatedAt

`func (o *Checkin) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Checkin) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Checkin) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Checkin) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Checkin) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Checkin) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


