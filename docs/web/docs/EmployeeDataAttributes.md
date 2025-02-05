# EmployeeDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User ID | 
**Role** | **string** | User role | 
**DistributorId** | **string** | Distributor ID | 
**CreatedAt** | **time.Time** | Creation date | 

## Methods

### NewEmployeeDataAttributes

`func NewEmployeeDataAttributes(userId string, role string, distributorId string, createdAt time.Time, ) *EmployeeDataAttributes`

NewEmployeeDataAttributes instantiates a new EmployeeDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeDataAttributesWithDefaults

`func NewEmployeeDataAttributesWithDefaults() *EmployeeDataAttributes`

NewEmployeeDataAttributesWithDefaults instantiates a new EmployeeDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *EmployeeDataAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EmployeeDataAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EmployeeDataAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetRole

`func (o *EmployeeDataAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EmployeeDataAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EmployeeDataAttributes) SetRole(v string)`

SetRole sets Role field to given value.


### GetDistributorId

`func (o *EmployeeDataAttributes) GetDistributorId() string`

GetDistributorId returns the DistributorId field if non-nil, zero value otherwise.

### GetDistributorIdOk

`func (o *EmployeeDataAttributes) GetDistributorIdOk() (*string, bool)`

GetDistributorIdOk returns a tuple with the DistributorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributorId

`func (o *EmployeeDataAttributes) SetDistributorId(v string)`

SetDistributorId sets DistributorId field to given value.


### GetCreatedAt

`func (o *EmployeeDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmployeeDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmployeeDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


