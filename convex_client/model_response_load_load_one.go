/*
Convex App - OpenAPI 3.0

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResponseLoadLoadOne type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseLoadLoadOne{}

// ResponseLoadLoadOne struct for ResponseLoadLoadOne
type ResponseLoadLoadOne struct {
	Status string `json:"status"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	ErrorData map[string]interface{} `json:"errorData,omitempty"`
	Value NullableResponseLoadLoadOneValue `json:"value,omitempty"`
}

type _ResponseLoadLoadOne ResponseLoadLoadOne

// NewResponseLoadLoadOne instantiates a new ResponseLoadLoadOne object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseLoadLoadOne(status string) *ResponseLoadLoadOne {
	this := ResponseLoadLoadOne{}
	this.Status = status
	return &this
}

// NewResponseLoadLoadOneWithDefaults instantiates a new ResponseLoadLoadOne object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseLoadLoadOneWithDefaults() *ResponseLoadLoadOne {
	this := ResponseLoadLoadOne{}
	return &this
}

// GetStatus returns the Status field value
func (o *ResponseLoadLoadOne) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResponseLoadLoadOne) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResponseLoadLoadOne) SetStatus(v string) {
	o.Status = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ResponseLoadLoadOne) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseLoadLoadOne) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ResponseLoadLoadOne) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ResponseLoadLoadOne) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorData returns the ErrorData field value if set, zero value otherwise.
func (o *ResponseLoadLoadOne) GetErrorData() map[string]interface{} {
	if o == nil || IsNil(o.ErrorData) {
		var ret map[string]interface{}
		return ret
	}
	return o.ErrorData
}

// GetErrorDataOk returns a tuple with the ErrorData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseLoadLoadOne) GetErrorDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ErrorData) {
		return map[string]interface{}{}, false
	}
	return o.ErrorData, true
}

// HasErrorData returns a boolean if a field has been set.
func (o *ResponseLoadLoadOne) HasErrorData() bool {
	if o != nil && !IsNil(o.ErrorData) {
		return true
	}

	return false
}

// SetErrorData gets a reference to the given map[string]interface{} and assigns it to the ErrorData field.
func (o *ResponseLoadLoadOne) SetErrorData(v map[string]interface{}) {
	o.ErrorData = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseLoadLoadOne) GetValue() ResponseLoadLoadOneValue {
	if o == nil || IsNil(o.Value.Get()) {
		var ret ResponseLoadLoadOneValue
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseLoadLoadOne) GetValueOk() (*ResponseLoadLoadOneValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ResponseLoadLoadOne) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableResponseLoadLoadOneValue and assigns it to the Value field.
func (o *ResponseLoadLoadOne) SetValue(v ResponseLoadLoadOneValue) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ResponseLoadLoadOne) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ResponseLoadLoadOne) UnsetValue() {
	o.Value.Unset()
}

func (o ResponseLoadLoadOne) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseLoadLoadOne) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.ErrorData) {
		toSerialize["errorData"] = o.ErrorData
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

func (o *ResponseLoadLoadOne) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varResponseLoadLoadOne := _ResponseLoadLoadOne{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseLoadLoadOne)

	if err != nil {
		return err
	}

	*o = ResponseLoadLoadOne(varResponseLoadLoadOne)

	return err
}

type NullableResponseLoadLoadOne struct {
	value *ResponseLoadLoadOne
	isSet bool
}

func (v NullableResponseLoadLoadOne) Get() *ResponseLoadLoadOne {
	return v.value
}

func (v *NullableResponseLoadLoadOne) Set(val *ResponseLoadLoadOne) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseLoadLoadOne) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseLoadLoadOne) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseLoadLoadOne(val *ResponseLoadLoadOne) *NullableResponseLoadLoadOne {
	return &NullableResponseLoadLoadOne{value: val, isSet: true}
}

func (v NullableResponseLoadLoadOne) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseLoadLoadOne) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

