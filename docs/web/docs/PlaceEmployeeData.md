# PlaceEmployeeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | employee ID | 
**Type** | **string** |  | 
**Attributes** | [**PlaceEmployeeDataAttributes**](PlaceEmployeeDataAttributes.md) |  | 

## Methods

### NewPlaceEmployeeData

`func NewPlaceEmployeeData(id string, type_ string, attributes PlaceEmployeeDataAttributes, ) *PlaceEmployeeData`

NewPlaceEmployeeData instantiates a new PlaceEmployeeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceEmployeeDataWithDefaults

`func NewPlaceEmployeeDataWithDefaults() *PlaceEmployeeData`

NewPlaceEmployeeDataWithDefaults instantiates a new PlaceEmployeeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlaceEmployeeData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaceEmployeeData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaceEmployeeData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PlaceEmployeeData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlaceEmployeeData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlaceEmployeeData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PlaceEmployeeData) GetAttributes() PlaceEmployeeDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PlaceEmployeeData) GetAttributesOk() (*PlaceEmployeeDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PlaceEmployeeData) SetAttributes(v PlaceEmployeeDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


