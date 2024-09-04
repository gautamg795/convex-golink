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

// checks if the RequestLoadLoadOneArgs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestLoadLoadOneArgs{}

// RequestLoadLoadOneArgs struct for RequestLoadLoadOneArgs
type RequestLoadLoadOneArgs struct {
	NormalizedId string `json:"normalizedId"`
	Token string `json:"token"`
}

type _RequestLoadLoadOneArgs RequestLoadLoadOneArgs

// NewRequestLoadLoadOneArgs instantiates a new RequestLoadLoadOneArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestLoadLoadOneArgs(normalizedId string, token string) *RequestLoadLoadOneArgs {
	this := RequestLoadLoadOneArgs{}
	this.NormalizedId = normalizedId
	this.Token = token
	return &this
}

// NewRequestLoadLoadOneArgsWithDefaults instantiates a new RequestLoadLoadOneArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestLoadLoadOneArgsWithDefaults() *RequestLoadLoadOneArgs {
	this := RequestLoadLoadOneArgs{}
	return &this
}

// GetNormalizedId returns the NormalizedId field value
func (o *RequestLoadLoadOneArgs) GetNormalizedId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NormalizedId
}

// GetNormalizedIdOk returns a tuple with the NormalizedId field value
// and a boolean to check if the value has been set.
func (o *RequestLoadLoadOneArgs) GetNormalizedIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NormalizedId, true
}

// SetNormalizedId sets field value
func (o *RequestLoadLoadOneArgs) SetNormalizedId(v string) {
	o.NormalizedId = v
}

// GetToken returns the Token field value
func (o *RequestLoadLoadOneArgs) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *RequestLoadLoadOneArgs) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *RequestLoadLoadOneArgs) SetToken(v string) {
	o.Token = v
}

func (o RequestLoadLoadOneArgs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestLoadLoadOneArgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["normalizedId"] = o.NormalizedId
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *RequestLoadLoadOneArgs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"normalizedId",
		"token",
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

	varRequestLoadLoadOneArgs := _RequestLoadLoadOneArgs{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRequestLoadLoadOneArgs)

	if err != nil {
		return err
	}

	*o = RequestLoadLoadOneArgs(varRequestLoadLoadOneArgs)

	return err
}

type NullableRequestLoadLoadOneArgs struct {
	value *RequestLoadLoadOneArgs
	isSet bool
}

func (v NullableRequestLoadLoadOneArgs) Get() *RequestLoadLoadOneArgs {
	return v.value
}

func (v *NullableRequestLoadLoadOneArgs) Set(val *RequestLoadLoadOneArgs) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestLoadLoadOneArgs) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestLoadLoadOneArgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestLoadLoadOneArgs(val *RequestLoadLoadOneArgs) *NullableRequestLoadLoadOneArgs {
	return &NullableRequestLoadLoadOneArgs{value: val, isSet: true}
}

func (v NullableRequestLoadLoadOneArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestLoadLoadOneArgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


