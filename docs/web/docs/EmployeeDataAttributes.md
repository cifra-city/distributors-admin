# EmployeeDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | User email | [optional] 
**Password** | Pointer to **string** | New user password | [optional] 
**DistributorId** | **string** | Distributor ID | 
**CreatedAt** | **time.Time** | Creation date | 

## Methods

### NewEmployeeDataAttributes

`func NewEmployeeDataAttributes(distributorId string, createdAt time.Time, ) *EmployeeDataAttributes`

NewEmployeeDataAttributes instantiates a new EmployeeDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeDataAttributesWithDefaults

`func NewEmployeeDataAttributesWithDefaults() *EmployeeDataAttributes`

NewEmployeeDataAttributesWithDefaults instantiates a new EmployeeDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmployeeDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmployeeDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmployeeDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmployeeDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *EmployeeDataAttributes) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EmployeeDataAttributes) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EmployeeDataAttributes) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EmployeeDataAttributes) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

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


