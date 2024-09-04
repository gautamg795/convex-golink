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

// checks if the RequestStoreDefaultArgsLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestStoreDefaultArgsLink{}

// RequestStoreDefaultArgsLink struct for RequestStoreDefaultArgsLink
type RequestStoreDefaultArgsLink struct {
	Created float32 `json:"created"`
	LastEdit float32 `json:"lastEdit"`
	Long string `json:"long"`
	NormalizedId string `json:"normalizedId"`
	Owner string `json:"owner"`
	Short string `json:"short"`
}

type _RequestStoreDefaultArgsLink RequestStoreDefaultArgsLink

// NewRequestStoreDefaultArgsLink instantiates a new RequestStoreDefaultArgsLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestStoreDefaultArgsLink(created float32, lastEdit float32, long string, normalizedId string, owner string, short string) *RequestStoreDefaultArgsLink {
	this := RequestStoreDefaultArgsLink{}
	this.Created = created
	this.LastEdit = lastEdit
	this.Long = long
	this.NormalizedId = normalizedId
	this.Owner = owner
	this.Short = short
	return &this
}

// NewRequestStoreDefaultArgsLinkWithDefaults instantiates a new RequestStoreDefaultArgsLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestStoreDefaultArgsLinkWithDefaults() *RequestStoreDefaultArgsLink {
	this := RequestStoreDefaultArgsLink{}
	return &this
}

// GetCreated returns the Created field value
func (o *RequestStoreDefaultArgsLink) GetCreated() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RequestStoreDefaultArgsLink) GetCreatedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RequestStoreDefaultArgsLink) SetCreated(v float32) {
	o.Created = v
}

// GetLastEdit returns the LastEdit field value
func (o *RequestStoreDefaultArgsLink) GetLastEdit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LastEdit
}

// GetLastEditOk returns a tuple with the LastEdit field value
// and a boolean to check if the value has been set.
func (o *RequestStoreDefaultArgsLink) GetLastEditOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEdit, true
}

// SetLastEdit sets field value
func (o *RequestStoreDefaultArgsLink) SetLastEdit(v float32) {
	o.LastEdit = v
}

// GetLong returns the Long field value
func (o *RequestStoreDefaultArgsLink) GetLong() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Long
}

// GetLongOk returns a tuple with the Long field value
// and a boolean to check if the value has been set.
func (o *RequestStoreDefaultArgsLink) GetLongOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Long, true
}

// SetLong sets field value
func (o *RequestStoreDefaultArgsLink) SetLong(v string) {
	o.Long = v
}

// GetNormalizedId returns the NormalizedId field value
func (o *RequestStoreDefaultArgsLink) GetNormalizedId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NormalizedId
}

// GetNormalizedIdOk returns a tuple with the NormalizedId field value
// and a boolean to check if the value has been set.
func (o *RequestStoreDefaultArgsLink) GetNormalizedIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NormalizedId, true
}

// SetNormalizedId sets field value
func (o *RequestStoreDefaultArgsLink) SetNormalizedId(v string) {
	o.NormalizedId = v
}

// GetOwner returns the Owner field value
func (o *RequestStoreDefaultArgsLink) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *RequestStoreDefaultArgsLink) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *RequestStoreDefaultArgsLink) SetOwner(v string) {
	o.Owner = v
}

// GetShort returns the Short field value
func (o *RequestStoreDefaultArgsLink) GetShort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Short
}

// GetShortOk returns a tuple with the Short field value
// and a boolean to check if the value has been set.
func (o *RequestStoreDefaultArgsLink) GetShortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Short, true
}

// SetShort sets field value
func (o *RequestStoreDefaultArgsLink) SetShort(v string) {
	o.Short = v
}

func (o RequestStoreDefaultArgsLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestStoreDefaultArgsLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["lastEdit"] = o.LastEdit
	toSerialize["long"] = o.Long
	toSerialize["normalizedId"] = o.NormalizedId
	toSerialize["owner"] = o.Owner
	toSerialize["short"] = o.Short
	return toSerialize, nil
}

func (o *RequestStoreDefaultArgsLink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"lastEdit",
		"long",
		"normalizedId",
		"owner",
		"short",
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

	varRequestStoreDefaultArgsLink := _RequestStoreDefaultArgsLink{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRequestStoreDefaultArgsLink)

	if err != nil {
		return err
	}

	*o = RequestStoreDefaultArgsLink(varRequestStoreDefaultArgsLink)

	return err
}

type NullableRequestStoreDefaultArgsLink struct {
	value *RequestStoreDefaultArgsLink
	isSet bool
}

func (v NullableRequestStoreDefaultArgsLink) Get() *RequestStoreDefaultArgsLink {
	return v.value
}

func (v *NullableRequestStoreDefaultArgsLink) Set(val *RequestStoreDefaultArgsLink) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestStoreDefaultArgsLink) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestStoreDefaultArgsLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestStoreDefaultArgsLink(val *RequestStoreDefaultArgsLink) *NullableRequestStoreDefaultArgsLink {
	return &NullableRequestStoreDefaultArgsLink{value: val, isSet: true}
}

func (v NullableRequestStoreDefaultArgsLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestStoreDefaultArgsLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


