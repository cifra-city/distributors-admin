# EmployeeUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | User email | [optional] 
**Password** | Pointer to **string** | New user password | [optional] 
**DistributorId** | Pointer to **string** | Distributor ID | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation date | [optional] 

## Methods

### NewEmployeeUpdateDataAttributes

`func NewEmployeeUpdateDataAttributes() *EmployeeUpdateDataAttributes`

NewEmployeeUpdateDataAttributes instantiates a new EmployeeUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeUpdateDataAttributesWithDefaults

`func NewEmployeeUpdateDataAttributesWithDefaults() *EmployeeUpdateDataAttributes`

NewEmployeeUpdateDataAttributesWithDefaults instantiates a new EmployeeUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmployeeUpdateDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmployeeUpdateDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmployeeUpdateDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmployeeUpdateDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *EmployeeUpdateDataAttributes) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EmployeeUpdateDataAttributes) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EmployeeUpdateDataAttributes) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EmployeeUpdateDataAttributes) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDistributorId

`func (o *EmployeeUpdateDataAttributes) GetDistributorId() string`

GetDistributorId returns the DistributorId field if non-nil, zero value otherwise.

### GetDistributorIdOk

`func (o *EmployeeUpdateDataAttributes) GetDistributorIdOk() (*string, bool)`

GetDistributorIdOk returns a tuple with the DistributorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributorId

`func (o *EmployeeUpdateDataAttributes) SetDistributorId(v string)`

SetDistributorId sets DistributorId field to given value.

### HasDistributorId

`func (o *EmployeeUpdateDataAttributes) HasDistributorId() bool`

HasDistributorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EmployeeUpdateDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmployeeUpdateDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmployeeUpdateDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EmployeeUpdateDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


