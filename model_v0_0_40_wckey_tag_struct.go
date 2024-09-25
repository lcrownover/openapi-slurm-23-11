/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.05.3&openapi/slurmdbd&openapi/slurmctld
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the V0040WckeyTagStruct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040WckeyTagStruct{}

// V0040WckeyTagStruct struct for V0040WckeyTagStruct
type V0040WckeyTagStruct struct {
	// WCKey name
	Wckey string `json:"wckey"`
	// Active flags
	Flags []string `json:"flags"`
}

type _V0040WckeyTagStruct V0040WckeyTagStruct

// NewV0040WckeyTagStruct instantiates a new V0040WckeyTagStruct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040WckeyTagStruct(wckey string, flags []string) *V0040WckeyTagStruct {
	this := V0040WckeyTagStruct{}
	this.Wckey = wckey
	this.Flags = flags
	return &this
}

// NewV0040WckeyTagStructWithDefaults instantiates a new V0040WckeyTagStruct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040WckeyTagStructWithDefaults() *V0040WckeyTagStruct {
	this := V0040WckeyTagStruct{}
	return &this
}

// GetWckey returns the Wckey field value
func (o *V0040WckeyTagStruct) GetWckey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Wckey
}

// GetWckeyOk returns a tuple with the Wckey field value
// and a boolean to check if the value has been set.
func (o *V0040WckeyTagStruct) GetWckeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wckey, true
}

// SetWckey sets field value
func (o *V0040WckeyTagStruct) SetWckey(v string) {
	o.Wckey = v
}

// GetFlags returns the Flags field value
func (o *V0040WckeyTagStruct) GetFlags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *V0040WckeyTagStruct) GetFlagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flags, true
}

// SetFlags sets field value
func (o *V0040WckeyTagStruct) SetFlags(v []string) {
	o.Flags = v
}

func (o V0040WckeyTagStruct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040WckeyTagStruct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wckey"] = o.Wckey
	toSerialize["flags"] = o.Flags
	return toSerialize, nil
}

func (o *V0040WckeyTagStruct) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wckey",
		"flags",
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

	varV0040WckeyTagStruct := _V0040WckeyTagStruct{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040WckeyTagStruct)

	if err != nil {
		return err
	}

	*o = V0040WckeyTagStruct(varV0040WckeyTagStruct)

	return err
}

type NullableV0040WckeyTagStruct struct {
	value *V0040WckeyTagStruct
	isSet bool
}

func (v NullableV0040WckeyTagStruct) Get() *V0040WckeyTagStruct {
	return v.value
}

func (v *NullableV0040WckeyTagStruct) Set(val *V0040WckeyTagStruct) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040WckeyTagStruct) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040WckeyTagStruct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040WckeyTagStruct(val *V0040WckeyTagStruct) *NullableV0040WckeyTagStruct {
	return &NullableV0040WckeyTagStruct{value: val, isSet: true}
}

func (v NullableV0040WckeyTagStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040WckeyTagStruct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


