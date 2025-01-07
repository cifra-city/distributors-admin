# DistributorCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Resource ID | [optional] 
**Type** | **string** |  | 
**Attributes** | [**DistributorUpdateDataAttributes**](DistributorUpdateDataAttributes.md) |  | 

## Methods

### NewDistributorCreateData

`func NewDistributorCreateData(type_ string, attributes DistributorUpdateDataAttributes, ) *DistributorCreateData`

NewDistributorCreateData instantiates a new DistributorCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributorCreateDataWithDefaults

`func NewDistributorCreateDataWithDefaults() *DistributorCreateData`

NewDistributorCreateDataWithDefaults instantiates a new DistributorCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DistributorCreateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DistributorCreateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DistributorCreateData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DistributorCreateData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DistributorCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DistributorCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DistributorCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *DistributorCreateData) GetAttributes() DistributorUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DistributorCreateData) GetAttributesOk() (*DistributorUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DistributorCreateData) SetAttributes(v DistributorUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


