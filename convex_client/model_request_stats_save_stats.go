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

// checks if the RequestStatsSaveStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestStatsSaveStats{}

// RequestStatsSaveStats struct for RequestStatsSaveStats
type RequestStatsSaveStats struct {
	Args RequestStatsSaveStatsArgs `json:"args"`
}

type _RequestStatsSaveStats RequestStatsSaveStats

// NewRequestStatsSaveStats instantiates a new RequestStatsSaveStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestStatsSaveStats(args RequestStatsSaveStatsArgs) *RequestStatsSaveStats {
	this := RequestStatsSaveStats{}
	this.Args = args
	return &this
}

// NewRequestStatsSaveStatsWithDefaults instantiates a new RequestStatsSaveStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestStatsSaveStatsWithDefaults() *RequestStatsSaveStats {
	this := RequestStatsSaveStats{}
	return &this
}

// GetArgs returns the Args field value
func (o *RequestStatsSaveStats) GetArgs() RequestStatsSaveStatsArgs {
	if o == nil {
		var ret RequestStatsSaveStatsArgs
		return ret
	}

	return o.Args
}

// GetArgsOk returns a tuple with the Args field value
// and a boolean to check if the value has been set.
func (o *RequestStatsSaveStats) GetArgsOk() (*RequestStatsSaveStatsArgs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Args, true
}

// SetArgs sets field value
func (o *RequestStatsSaveStats) SetArgs(v RequestStatsSaveStatsArgs) {
	o.Args = v
}

func (o RequestStatsSaveStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestStatsSaveStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["args"] = o.Args
	return toSerialize, nil
}

func (o *RequestStatsSaveStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"args",
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

	varRequestStatsSaveStats := _RequestStatsSaveStats{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRequestStatsSaveStats)

	if err != nil {
		return err
	}

	*o = RequestStatsSaveStats(varRequestStatsSaveStats)

	return err
}

type NullableRequestStatsSaveStats struct {
	value *RequestStatsSaveStats
	isSet bool
}

func (v NullableRequestStatsSaveStats) Get() *RequestStatsSaveStats {
	return v.value
}

func (v *NullableRequestStatsSaveStats) Set(val *RequestStatsSaveStats) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestStatsSaveStats) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestStatsSaveStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestStatsSaveStats(val *RequestStatsSaveStats) *NullableRequestStatsSaveStats {
	return &NullableRequestStatsSaveStats{value: val, isSet: true}
}

func (v NullableRequestStatsSaveStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestStatsSaveStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


