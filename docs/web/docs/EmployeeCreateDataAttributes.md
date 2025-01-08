# EmployeeCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User ID | 
**Role** | **string** | User role | 
**CreatedAt** | **time.Time** | Creation date | 

## Methods

### NewEmployeeCreateDataAttributes

`func NewEmployeeCreateDataAttributes(userId string, role string, createdAt time.Time, ) *EmployeeCreateDataAttributes`

NewEmployeeCreateDataAttributes instantiates a new EmployeeCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeCreateDataAttributesWithDefaults

`func NewEmployeeCreateDataAttributesWithDefaults() *EmployeeCreateDataAttributes`

NewEmployeeCreateDataAttributesWithDefaults instantiates a new EmployeeCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *EmployeeCreateDataAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EmployeeCreateDataAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EmployeeCreateDataAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetRole

`func (o *EmployeeCreateDataAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EmployeeCreateDataAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EmployeeCreateDataAttributes) SetRole(v string)`

SetRole sets Role field to given value.


### GetCreatedAt

`func (o *EmployeeCreateDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmployeeCreateDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmployeeCreateDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


