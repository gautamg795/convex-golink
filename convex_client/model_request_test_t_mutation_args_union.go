/*
Convex App - OpenAPI 3.0

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// RequestTestTMutationArgsUnion - struct for RequestTestTMutationArgsUnion
type RequestTestTMutationArgsUnion struct {
	String *string
}

// stringAsRequestTestTMutationArgsUnion is a convenience function that returns string wrapped in RequestTestTMutationArgsUnion
func StringAsRequestTestTMutationArgsUnion(v *string) RequestTestTMutationArgsUnion {
	return RequestTestTMutationArgsUnion{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestTestTMutationArgsUnion) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RequestTestTMutationArgsUnion)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestTestTMutationArgsUnion)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestTestTMutationArgsUnion) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestTestTMutationArgsUnion) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableRequestTestTMutationArgsUnion struct {
	value *RequestTestTMutationArgsUnion
	isSet bool
}

func (v NullableRequestTestTMutationArgsUnion) Get() *RequestTestTMutationArgsUnion {
	return v.value
}

func (v *NullableRequestTestTMutationArgsUnion) Set(val *RequestTestTMutationArgsUnion) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTestTMutationArgsUnion) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTestTMutationArgsUnion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTestTMutationArgsUnion(val *RequestTestTMutationArgsUnion) *NullableRequestTestTMutationArgsUnion {
	return &NullableRequestTestTMutationArgsUnion{value: val, isSet: true}
}

func (v NullableRequestTestTMutationArgsUnion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTestTMutationArgsUnion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

