# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**GoogleMapTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCategory

`func NewCategory(id string, name string, ) *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Category) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Category) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Category) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Category) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Category) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Category) SetName(v string)`

SetName sets Name field to given value.


### GetGoogleMapTypes

`func (o *Category) GetGoogleMapTypes() []string`

GetGoogleMapTypes returns the GoogleMapTypes field if non-nil, zero value otherwise.

### GetGoogleMapTypesOk

`func (o *Category) GetGoogleMapTypesOk() (*[]string, bool)`

GetGoogleMapTypesOk returns a tuple with the GoogleMapTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleMapTypes

`func (o *Category) SetGoogleMapTypes(v []string)`

SetGoogleMapTypes sets GoogleMapTypes field to given value.

### HasGoogleMapTypes

`func (o *Category) HasGoogleMapTypes() bool`

HasGoogleMapTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


