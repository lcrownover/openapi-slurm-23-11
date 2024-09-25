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

// checks if the V0040AssocSharesObjWrapFairshare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040AssocSharesObjWrapFairshare{}

// V0040AssocSharesObjWrapFairshare struct for V0040AssocSharesObjWrapFairshare
type V0040AssocSharesObjWrapFairshare struct {
	// Fairshare factor
	Factor *float64 `json:"factor,omitempty"`
	// Fairshare factor at this level; stored on an assoc as a long double, but that is not needed for display in sshare
	Level *float64 `json:"level,omitempty"`
}

// NewV0040AssocSharesObjWrapFairshare instantiates a new V0040AssocSharesObjWrapFairshare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040AssocSharesObjWrapFairshare() *V0040AssocSharesObjWrapFairshare {
	this := V0040AssocSharesObjWrapFairshare{}
	return &this
}

// NewV0040AssocSharesObjWrapFairshareWithDefaults instantiates a new V0040AssocSharesObjWrapFairshare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AssocSharesObjWrapFairshareWithDefaults() *V0040AssocSharesObjWrapFairshare {
	this := V0040AssocSharesObjWrapFairshare{}
	return &this
}

// GetFactor returns the Factor field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrapFairshare) GetFactor() float64 {
	if o == nil || IsNil(o.Factor) {
		var ret float64
		return ret
	}
	return *o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrapFairshare) GetFactorOk() (*float64, bool) {
	if o == nil || IsNil(o.Factor) {
		return nil, false
	}
	return o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrapFairshare) HasFactor() bool {
	if o != nil && !IsNil(o.Factor) {
		return true
	}

	return false
}

// SetFactor gets a reference to the given float64 and assigns it to the Factor field.
func (o *V0040AssocSharesObjWrapFairshare) SetFactor(v float64) {
	o.Factor = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *V0040AssocSharesObjWrapFairshare) GetLevel() float64 {
	if o == nil || IsNil(o.Level) {
		var ret float64
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocSharesObjWrapFairshare) GetLevelOk() (*float64, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *V0040AssocSharesObjWrapFairshare) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given float64 and assigns it to the Level field.
func (o *V0040AssocSharesObjWrapFairshare) SetLevel(v float64) {
	o.Level = &v
}

func (o V0040AssocSharesObjWrapFairshare) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040AssocSharesObjWrapFairshare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Factor) {
		toSerialize["factor"] = o.Factor
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	return toSerialize, nil
}

type NullableV0040AssocSharesObjWrapFairshare struct {
	value *V0040AssocSharesObjWrapFairshare
	isSet bool
}

func (v NullableV0040AssocSharesObjWrapFairshare) Get() *V0040AssocSharesObjWrapFairshare {
	return v.value
}

func (v *NullableV0040AssocSharesObjWrapFairshare) Set(val *V0040AssocSharesObjWrapFairshare) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040AssocSharesObjWrapFairshare) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040AssocSharesObjWrapFairshare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040AssocSharesObjWrapFairshare(val *V0040AssocSharesObjWrapFairshare) *NullableV0040AssocSharesObjWrapFairshare {
	return &NullableV0040AssocSharesObjWrapFairshare{value: val, isSet: true}
}

func (v NullableV0040AssocSharesObjWrapFairshare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040AssocSharesObjWrapFairshare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


