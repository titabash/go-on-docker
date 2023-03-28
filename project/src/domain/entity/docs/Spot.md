# Spot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Instagram** | Pointer to [**Instagram**](Instagram.md) |  | [optional] 
**GoogleMap** | Pointer to [**GoogleMap**](GoogleMap.md) |  | [optional] 
**Proprietary** | Pointer to [**Proprietary**](Proprietary.md) |  | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Categories** | Pointer to [**[]Category**](Category.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewSpot

`func NewSpot(id string, createdAt time.Time, updatedAt time.Time, ) *Spot`

NewSpot instantiates a new Spot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotWithDefaults

`func NewSpotWithDefaults() *Spot`

NewSpotWithDefaults instantiates a new Spot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Spot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Spot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Spot) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Spot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Spot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Spot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Spot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Spot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Spot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Spot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Spot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProvider

`func (o *Spot) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Spot) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Spot) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Spot) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLocation

`func (o *Spot) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Spot) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Spot) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Spot) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetInstagram

`func (o *Spot) GetInstagram() Instagram`

GetInstagram returns the Instagram field if non-nil, zero value otherwise.

### GetInstagramOk

`func (o *Spot) GetInstagramOk() (*Instagram, bool)`

GetInstagramOk returns a tuple with the Instagram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagram

`func (o *Spot) SetInstagram(v Instagram)`

SetInstagram sets Instagram field to given value.

### HasInstagram

`func (o *Spot) HasInstagram() bool`

HasInstagram returns a boolean if a field has been set.

### GetGoogleMap

`func (o *Spot) GetGoogleMap() GoogleMap`

GetGoogleMap returns the GoogleMap field if non-nil, zero value otherwise.

### GetGoogleMapOk

`func (o *Spot) GetGoogleMapOk() (*GoogleMap, bool)`

GetGoogleMapOk returns a tuple with the GoogleMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleMap

`func (o *Spot) SetGoogleMap(v GoogleMap)`

SetGoogleMap sets GoogleMap field to given value.

### HasGoogleMap

`func (o *Spot) HasGoogleMap() bool`

HasGoogleMap returns a boolean if a field has been set.

### GetProprietary

`func (o *Spot) GetProprietary() Proprietary`

GetProprietary returns the Proprietary field if non-nil, zero value otherwise.

### GetProprietaryOk

`func (o *Spot) GetProprietaryOk() (*Proprietary, bool)`

GetProprietaryOk returns a tuple with the Proprietary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProprietary

`func (o *Spot) SetProprietary(v Proprietary)`

SetProprietary sets Proprietary field to given value.

### HasProprietary

`func (o *Spot) HasProprietary() bool`

HasProprietary returns a boolean if a field has been set.

### GetAddress

`func (o *Spot) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Spot) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Spot) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Spot) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCategories

`func (o *Spot) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Spot) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Spot) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Spot) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Spot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Spot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Spot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Spot) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Spot) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Spot) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


