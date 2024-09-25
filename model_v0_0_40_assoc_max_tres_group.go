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

// checks if the V0040AssocMaxTresGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040AssocMaxTresGroup{}

// V0040AssocMaxTresGroup struct for V0040AssocMaxTresGroup
type V0040AssocMaxTresGroup struct {
	Minutes []V0040Tres `json:"minutes,omitempty"`
	Active []V0040Tres `json:"active,omitempty"`
}

// NewV0040AssocMaxTresGroup instantiates a new V0040AssocMaxTresGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040AssocMaxTresGroup() *V0040AssocMaxTresGroup {
	this := V0040AssocMaxTresGroup{}
	return &this
}

// NewV0040AssocMaxTresGroupWithDefaults instantiates a new V0040AssocMaxTresGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AssocMaxTresGroupWithDefaults() *V0040AssocMaxTresGroup {
	this := V0040AssocMaxTresGroup{}
	return &this
}

// GetMinutes returns the Minutes field value if set, zero value otherwise.
func (o *V0040AssocMaxTresGroup) GetMinutes() []V0040Tres {
	if o == nil || IsNil(o.Minutes) {
		var ret []V0040Tres
		return ret
	}
	return o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMaxTresGroup) GetMinutesOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Minutes) {
		return nil, false
	}
	return o.Minutes, true
}

// HasMinutes returns a boolean if a field has been set.
func (o *V0040AssocMaxTresGroup) HasMinutes() bool {
	if o != nil && !IsNil(o.Minutes) {
		return true
	}

	return false
}

// SetMinutes gets a reference to the given []V0040Tres and assigns it to the Minutes field.
func (o *V0040AssocMaxTresGroup) SetMinutes(v []V0040Tres) {
	o.Minutes = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *V0040AssocMaxTresGroup) GetActive() []V0040Tres {
	if o == nil || IsNil(o.Active) {
		var ret []V0040Tres
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMaxTresGroup) GetActiveOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *V0040AssocMaxTresGroup) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given []V0040Tres and assigns it to the Active field.
func (o *V0040AssocMaxTresGroup) SetActive(v []V0040Tres) {
	o.Active = v
}

func (o V0040AssocMaxTresGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040AssocMaxTresGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Minutes) {
		toSerialize["minutes"] = o.Minutes
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableV0040AssocMaxTresGroup struct {
	value *V0040AssocMaxTresGroup
	isSet bool
}

func (v NullableV0040AssocMaxTresGroup) Get() *V0040AssocMaxTresGroup {
	return v.value
}

func (v *NullableV0040AssocMaxTresGroup) Set(val *V0040AssocMaxTresGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040AssocMaxTresGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040AssocMaxTresGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040AssocMaxTresGroup(val *V0040AssocMaxTresGroup) *NullableV0040AssocMaxTresGroup {
	return &NullableV0040AssocMaxTresGroup{value: val, isSet: true}
}

func (v NullableV0040AssocMaxTresGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040AssocMaxTresGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


