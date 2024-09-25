/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.05.3&openapi/slurmdbd&openapi/slurmctld
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_slurm_23_11

import (
	"encoding/json"
)

// checks if the V0040JobInfoPower type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040JobInfoPower{}

// V0040JobInfoPower struct for V0040JobInfoPower
type V0040JobInfoPower struct {
	// removed field
	// Deprecated
	Flags []interface{} `json:"flags,omitempty"`
}

// NewV0040JobInfoPower instantiates a new V0040JobInfoPower object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040JobInfoPower() *V0040JobInfoPower {
	this := V0040JobInfoPower{}
	return &this
}

// NewV0040JobInfoPowerWithDefaults instantiates a new V0040JobInfoPower object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040JobInfoPowerWithDefaults() *V0040JobInfoPower {
	this := V0040JobInfoPower{}
	return &this
}

// GetFlags returns the Flags field value if set, zero value otherwise.
// Deprecated
func (o *V0040JobInfoPower) GetFlags() []interface{} {
	if o == nil || IsNil(o.Flags) {
		var ret []interface{}
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *V0040JobInfoPower) GetFlagsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0040JobInfoPower) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []interface{} and assigns it to the Flags field.
// Deprecated
func (o *V0040JobInfoPower) SetFlags(v []interface{}) {
	o.Flags = v
}

func (o V0040JobInfoPower) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040JobInfoPower) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

type NullableV0040JobInfoPower struct {
	value *V0040JobInfoPower
	isSet bool
}

func (v NullableV0040JobInfoPower) Get() *V0040JobInfoPower {
	return v.value
}

func (v *NullableV0040JobInfoPower) Set(val *V0040JobInfoPower) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040JobInfoPower) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040JobInfoPower) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040JobInfoPower(val *V0040JobInfoPower) *NullableV0040JobInfoPower {
	return &NullableV0040JobInfoPower{value: val, isSet: true}
}

func (v NullableV0040JobInfoPower) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040JobInfoPower) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


