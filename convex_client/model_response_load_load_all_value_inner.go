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

// checks if the ResponseLoadLoadAllValueInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseLoadLoadAllValueInner{}

// ResponseLoadLoadAllValueInner struct for ResponseLoadLoadAllValueInner
type ResponseLoadLoadAllValueInner struct {
	CreationTime float32 `json:"_creationTime"`
	// ID from table \"links\"
	Id string `json:"_id"`
	Created float32 `json:"created"`
	LastEdit float32 `json:"lastEdit"`
	Long string `json:"long"`
	NormalizedId string `json:"normalizedId"`
	Owner string `json:"owner"`
	Short string `json:"short"`
}

type _ResponseLoadLoadAllValueInner ResponseLoadLoadAllValueInner

// NewResponseLoadLoadAllValueInner instantiates a new ResponseLoadLoadAllValueInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseLoadLoadAllValueInner(creationTime float32, id string, created float32, lastEdit float32, long string, normalizedId string, owner string, short string) *ResponseLoadLoadAllValueInner {
	this := ResponseLoadLoadAllValueInner{}
	this.CreationTime = creationTime
	this.Id = id
	this.Created = created
	this.LastEdit = lastEdit
	this.Long = long
	this.NormalizedId = normalizedId
	this.Owner = owner
	this.Short = short
	return &this
}

// NewResponseLoadLoadAllValueInnerWithDefaults instantiates a new ResponseLoadLoadAllValueInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseLoadLoadAllValueInnerWithDefaults() *ResponseLoadLoadAllValueInner {
	this := ResponseLoadLoadAllValueInner{}
	return &this
}

// GetCreationTime returns the CreationTime field value
func (o *ResponseLoadLoadAllValueInner) GetCreationTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *ResponseLoadLoadAllValueInner) GetCreationTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *ResponseLoadLoadAllValueInner) SetCreationTime(v float32) {
	o.CreationTime = v
}

// GetId returns the Id field value
func (o *ResponseLoadLoadAllValueInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseLoadLoadAllValueInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseLoadLoadAllValueInner) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *ResponseLoadLoadAllValueInner) GetCreated() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ResponseLoadLoadAllValueInner) GetCreatedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ResponseLoadLoadAllValueInner) SetCreated(v float32) {
	o.Created = v
}

// GetLastEdit returns the LastEdit field value
func (o *ResponseLoadLoadAllValueInner) GetLastEdit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LastEdit
}

// GetLastEditOk returns a tuple with the LastEdit field value
// and a boolean to check if the value has been set.
func (o *ResponseLoadLoadAllValueInner) GetLastEditOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEdit, true
}

// SetLastEdit sets field value
func (o *ResponseLoadLoadAllValueInner) SetLastEdit(v float32) {
	o.LastEdit = v
}

// GetLong returns the Long field value
func (o *ResponseLoadLoadAllValueInner) GetLong() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Long
}

// GetLongOk returns a tuple with the Long field value
// and a boolean to check if the value has been set.
func (o *ResponseLoadLoadAllValueInner) GetLongOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Long, true
}

// SetLong sets field value
func (o *ResponseLoadLoadAllValueInner) SetLong(v string) {
	o.Long = v
}

// GetNormalizedId returns the NormalizedId field value
func (o *ResponseLoadLoadAllValueInner) GetNormalizedId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NormalizedId
}

// GetNormalizedIdOk returns a tuple with the NormalizedId field value
// and a boolean to check if the value has been set.
func (o *ResponseLoadLoadAllValueInner) GetNormalizedIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NormalizedId, true
}

// SetNormalizedId sets field value
func (o *ResponseLoadLoadAllValueInner) SetNormalizedId(v string) {
	o.NormalizedId = v
}

// GetOwner returns the Owner field value
func (o *ResponseLoadLoadAllValueInner) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *ResponseLoadLoadAllValueInner) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *ResponseLoadLoadAllValueInner) SetOwner(v string) {
	o.Owner = v
}

// GetShort returns the Short field value
func (o *ResponseLoadLoadAllValueInner) GetShort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Short
}

// GetShortOk returns a tuple with the Short field value
// and a boolean to check if the value has been set.
func (o *ResponseLoadLoadAllValueInner) GetShortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Short, true
}

// SetShort sets field value
func (o *ResponseLoadLoadAllValueInner) SetShort(v string) {
	o.Short = v
}

func (o ResponseLoadLoadAllValueInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseLoadLoadAllValueInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_creationTime"] = o.CreationTime
	toSerialize["_id"] = o.Id
	toSerialize["created"] = o.Created
	toSerialize["lastEdit"] = o.LastEdit
	toSerialize["long"] = o.Long
	toSerialize["normalizedId"] = o.NormalizedId
	toSerialize["owner"] = o.Owner
	toSerialize["short"] = o.Short
	return toSerialize, nil
}

func (o *ResponseLoadLoadAllValueInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_creationTime",
		"_id",
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

	varResponseLoadLoadAllValueInner := _ResponseLoadLoadAllValueInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseLoadLoadAllValueInner)

	if err != nil {
		return err
	}

	*o = ResponseLoadLoadAllValueInner(varResponseLoadLoadAllValueInner)

	return err
}

type NullableResponseLoadLoadAllValueInner struct {
	value *ResponseLoadLoadAllValueInner
	isSet bool
}

func (v NullableResponseLoadLoadAllValueInner) Get() *ResponseLoadLoadAllValueInner {
	return v.value
}

func (v *NullableResponseLoadLoadAllValueInner) Set(val *ResponseLoadLoadAllValueInner) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseLoadLoadAllValueInner) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseLoadLoadAllValueInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseLoadLoadAllValueInner(val *ResponseLoadLoadAllValueInner) *NullableResponseLoadLoadAllValueInner {
	return &NullableResponseLoadLoadAllValueInner{value: val, isSet: true}
}

func (v NullableResponseLoadLoadAllValueInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseLoadLoadAllValueInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


