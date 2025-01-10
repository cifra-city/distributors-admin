# DistributorEmployeeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | employee ID | 
**Type** | **string** |  | 
**Attributes** | [**DistributorEmployeeDataAttributes**](DistributorEmployeeDataAttributes.md) |  | 

## Methods

### NewDistributorEmployeeData

`func NewDistributorEmployeeData(id string, type_ string, attributes DistributorEmployeeDataAttributes, ) *DistributorEmployeeData`

NewDistributorEmployeeData instantiates a new DistributorEmployeeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributorEmployeeDataWithDefaults

`func NewDistributorEmployeeDataWithDefaults() *DistributorEmployeeData`

NewDistributorEmployeeDataWithDefaults instantiates a new DistributorEmployeeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DistributorEmployeeData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DistributorEmployeeData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DistributorEmployeeData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *DistributorEmployeeData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DistributorEmployeeData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DistributorEmployeeData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *DistributorEmployeeData) GetAttributes() DistributorEmployeeDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DistributorEmployeeData) GetAttributesOk() (*DistributorEmployeeDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DistributorEmployeeData) SetAttributes(v DistributorEmployeeDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


