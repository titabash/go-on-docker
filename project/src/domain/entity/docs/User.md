# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FirebaseAuthId** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** |  | 
**AccountName** | **string** |  | 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**ProfilePicture** | Pointer to [**UserProfilePicture**](UserProfilePicture.md) |  | [optional] 
**Age** | Pointer to **int32** |  | [optional] 
**Birthday** | Pointer to **time.Time** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**SpouseId** | Pointer to **NullableString** |  | [optional] 
**ChildrenIds** | Pointer to **[]string** |  | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**HomeLocation** | Pointer to [**Location**](Location.md) |  | [optional] 
**CurrentLocation** | Pointer to [**Location**](Location.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**FollowerIds** | Pointer to **[]string** |  | [optional] 
**FollowingIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUser

`func NewUser(id string, displayName string, accountName string, createdAt time.Time, updatedAt time.Time, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetFirebaseAuthId

`func (o *User) GetFirebaseAuthId() string`

GetFirebaseAuthId returns the FirebaseAuthId field if non-nil, zero value otherwise.

### GetFirebaseAuthIdOk

`func (o *User) GetFirebaseAuthIdOk() (*string, bool)`

GetFirebaseAuthIdOk returns a tuple with the FirebaseAuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirebaseAuthId

`func (o *User) SetFirebaseAuthId(v string)`

SetFirebaseAuthId sets FirebaseAuthId field to given value.

### HasFirebaseAuthId

`func (o *User) HasFirebaseAuthId() bool`

HasFirebaseAuthId returns a boolean if a field has been set.

### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetAccountName

`func (o *User) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *User) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *User) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetProfilePicture

`func (o *User) GetProfilePicture() UserProfilePicture`

GetProfilePicture returns the ProfilePicture field if non-nil, zero value otherwise.

### GetProfilePictureOk

`func (o *User) GetProfilePictureOk() (*UserProfilePicture, bool)`

GetProfilePictureOk returns a tuple with the ProfilePicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicture

`func (o *User) SetProfilePicture(v UserProfilePicture)`

SetProfilePicture sets ProfilePicture field to given value.

### HasProfilePicture

`func (o *User) HasProfilePicture() bool`

HasProfilePicture returns a boolean if a field has been set.

### GetAge

`func (o *User) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *User) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *User) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *User) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetBirthday

`func (o *User) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *User) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *User) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *User) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetGender

`func (o *User) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *User) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *User) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *User) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetSpouseId

`func (o *User) GetSpouseId() string`

GetSpouseId returns the SpouseId field if non-nil, zero value otherwise.

### GetSpouseIdOk

`func (o *User) GetSpouseIdOk() (*string, bool)`

GetSpouseIdOk returns a tuple with the SpouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouseId

`func (o *User) SetSpouseId(v string)`

SetSpouseId sets SpouseId field to given value.

### HasSpouseId

`func (o *User) HasSpouseId() bool`

HasSpouseId returns a boolean if a field has been set.

### SetSpouseIdNil

`func (o *User) SetSpouseIdNil(b bool)`

 SetSpouseIdNil sets the value for SpouseId to be an explicit nil

### UnsetSpouseId
`func (o *User) UnsetSpouseId()`

UnsetSpouseId ensures that no value is present for SpouseId, not even an explicit nil
### GetChildrenIds

`func (o *User) GetChildrenIds() []string`

GetChildrenIds returns the ChildrenIds field if non-nil, zero value otherwise.

### GetChildrenIdsOk

`func (o *User) GetChildrenIdsOk() (*[]string, bool)`

GetChildrenIdsOk returns a tuple with the ChildrenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenIds

`func (o *User) SetChildrenIds(v []string)`

SetChildrenIds sets ChildrenIds field to given value.

### HasChildrenIds

`func (o *User) HasChildrenIds() bool`

HasChildrenIds returns a boolean if a field has been set.

### GetAddress

`func (o *User) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *User) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *User) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *User) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetHomeLocation

`func (o *User) GetHomeLocation() Location`

GetHomeLocation returns the HomeLocation field if non-nil, zero value otherwise.

### GetHomeLocationOk

`func (o *User) GetHomeLocationOk() (*Location, bool)`

GetHomeLocationOk returns a tuple with the HomeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeLocation

`func (o *User) SetHomeLocation(v Location)`

SetHomeLocation sets HomeLocation field to given value.

### HasHomeLocation

`func (o *User) HasHomeLocation() bool`

HasHomeLocation returns a boolean if a field has been set.

### GetCurrentLocation

`func (o *User) GetCurrentLocation() Location`

GetCurrentLocation returns the CurrentLocation field if non-nil, zero value otherwise.

### GetCurrentLocationOk

`func (o *User) GetCurrentLocationOk() (*Location, bool)`

GetCurrentLocationOk returns a tuple with the CurrentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLocation

`func (o *User) SetCurrentLocation(v Location)`

SetCurrentLocation sets CurrentLocation field to given value.

### HasCurrentLocation

`func (o *User) HasCurrentLocation() bool`

HasCurrentLocation returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *User) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFollowerIds

`func (o *User) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *User) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *User) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *User) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetFollowingIds

`func (o *User) GetFollowingIds() []string`

GetFollowingIds returns the FollowingIds field if non-nil, zero value otherwise.

### GetFollowingIdsOk

`func (o *User) GetFollowingIdsOk() (*[]string, bool)`

GetFollowingIdsOk returns a tuple with the FollowingIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingIds

`func (o *User) SetFollowingIds(v []string)`

SetFollowingIds sets FollowingIds field to given value.

### HasFollowingIds

`func (o *User) HasFollowingIds() bool`

HasFollowingIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


