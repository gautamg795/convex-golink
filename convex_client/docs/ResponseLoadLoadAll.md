# ResponseLoadLoadAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ErrorData** | Pointer to **map[string]interface{}** |  | [optional] 
**Value** | Pointer to [**[]ResponseLoadLoadAllValueInner**](ResponseLoadLoadAllValueInner.md) |  | [optional] 

## Methods

### NewResponseLoadLoadAll

`func NewResponseLoadLoadAll(status string, ) *ResponseLoadLoadAll`

NewResponseLoadLoadAll instantiates a new ResponseLoadLoadAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseLoadLoadAllWithDefaults

`func NewResponseLoadLoadAllWithDefaults() *ResponseLoadLoadAll`

NewResponseLoadLoadAllWithDefaults instantiates a new ResponseLoadLoadAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResponseLoadLoadAll) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseLoadLoadAll) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseLoadLoadAll) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorMessage

`func (o *ResponseLoadLoadAll) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ResponseLoadLoadAll) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ResponseLoadLoadAll) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ResponseLoadLoadAll) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorData

`func (o *ResponseLoadLoadAll) GetErrorData() map[string]interface{}`

GetErrorData returns the ErrorData field if non-nil, zero value otherwise.

### GetErrorDataOk

`func (o *ResponseLoadLoadAll) GetErrorDataOk() (*map[string]interface{}, bool)`

GetErrorDataOk returns a tuple with the ErrorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorData

`func (o *ResponseLoadLoadAll) SetErrorData(v map[string]interface{})`

SetErrorData sets ErrorData field to given value.

### HasErrorData

`func (o *ResponseLoadLoadAll) HasErrorData() bool`

HasErrorData returns a boolean if a field has been set.

### GetValue

`func (o *ResponseLoadLoadAll) GetValue() []ResponseLoadLoadAllValueInner`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResponseLoadLoadAll) GetValueOk() (*[]ResponseLoadLoadAllValueInner, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResponseLoadLoadAll) SetValue(v []ResponseLoadLoadAllValueInner)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResponseLoadLoadAll) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


