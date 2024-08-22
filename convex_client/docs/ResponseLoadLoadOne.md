# ResponseLoadLoadOne

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ErrorData** | Pointer to **map[string]interface{}** |  | [optional] 
**Value** | Pointer to [**NullableResponseLoadLoadOneValue**](ResponseLoadLoadOneValue.md) |  | [optional] 

## Methods

### NewResponseLoadLoadOne

`func NewResponseLoadLoadOne(status string, ) *ResponseLoadLoadOne`

NewResponseLoadLoadOne instantiates a new ResponseLoadLoadOne object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseLoadLoadOneWithDefaults

`func NewResponseLoadLoadOneWithDefaults() *ResponseLoadLoadOne`

NewResponseLoadLoadOneWithDefaults instantiates a new ResponseLoadLoadOne object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResponseLoadLoadOne) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseLoadLoadOne) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseLoadLoadOne) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorMessage

`func (o *ResponseLoadLoadOne) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ResponseLoadLoadOne) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ResponseLoadLoadOne) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ResponseLoadLoadOne) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorData

`func (o *ResponseLoadLoadOne) GetErrorData() map[string]interface{}`

GetErrorData returns the ErrorData field if non-nil, zero value otherwise.

### GetErrorDataOk

`func (o *ResponseLoadLoadOne) GetErrorDataOk() (*map[string]interface{}, bool)`

GetErrorDataOk returns a tuple with the ErrorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorData

`func (o *ResponseLoadLoadOne) SetErrorData(v map[string]interface{})`

SetErrorData sets ErrorData field to given value.

### HasErrorData

`func (o *ResponseLoadLoadOne) HasErrorData() bool`

HasErrorData returns a boolean if a field has been set.

### GetValue

`func (o *ResponseLoadLoadOne) GetValue() ResponseLoadLoadOneValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResponseLoadLoadOne) GetValueOk() (*ResponseLoadLoadOneValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResponseLoadLoadOne) SetValue(v ResponseLoadLoadOneValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResponseLoadLoadOne) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ResponseLoadLoadOne) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ResponseLoadLoadOne) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


