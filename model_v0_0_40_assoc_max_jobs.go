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

// checks if the V0040AssocMaxJobs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040AssocMaxJobs{}

// V0040AssocMaxJobs struct for V0040AssocMaxJobs
type V0040AssocMaxJobs struct {
	Per *V0040AssocMaxJobsPer `json:"per,omitempty"`
	Active *V0040Uint32NoVal `json:"active,omitempty"`
	Accruing *V0040Uint32NoVal `json:"accruing,omitempty"`
	Total *V0040Uint32NoVal `json:"total,omitempty"`
}

// NewV0040AssocMaxJobs instantiates a new V0040AssocMaxJobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040AssocMaxJobs() *V0040AssocMaxJobs {
	this := V0040AssocMaxJobs{}
	return &this
}

// NewV0040AssocMaxJobsWithDefaults instantiates a new V0040AssocMaxJobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AssocMaxJobsWithDefaults() *V0040AssocMaxJobs {
	this := V0040AssocMaxJobs{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0040AssocMaxJobs) GetPer() V0040AssocMaxJobsPer {
	if o == nil || IsNil(o.Per) {
		var ret V0040AssocMaxJobsPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMaxJobs) GetPerOk() (*V0040AssocMaxJobsPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0040AssocMaxJobs) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0040AssocMaxJobsPer and assigns it to the Per field.
func (o *V0040AssocMaxJobs) SetPer(v V0040AssocMaxJobsPer) {
	o.Per = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *V0040AssocMaxJobs) GetActive() V0040Uint32NoVal {
	if o == nil || IsNil(o.Active) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMaxJobs) GetActiveOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *V0040AssocMaxJobs) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given V0040Uint32NoVal and assigns it to the Active field.
func (o *V0040AssocMaxJobs) SetActive(v V0040Uint32NoVal) {
	o.Active = &v
}

// GetAccruing returns the Accruing field value if set, zero value otherwise.
func (o *V0040AssocMaxJobs) GetAccruing() V0040Uint32NoVal {
	if o == nil || IsNil(o.Accruing) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Accruing
}

// GetAccruingOk returns a tuple with the Accruing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMaxJobs) GetAccruingOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Accruing) {
		return nil, false
	}
	return o.Accruing, true
}

// HasAccruing returns a boolean if a field has been set.
func (o *V0040AssocMaxJobs) HasAccruing() bool {
	if o != nil && !IsNil(o.Accruing) {
		return true
	}

	return false
}

// SetAccruing gets a reference to the given V0040Uint32NoVal and assigns it to the Accruing field.
func (o *V0040AssocMaxJobs) SetAccruing(v V0040Uint32NoVal) {
	o.Accruing = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0040AssocMaxJobs) GetTotal() V0040Uint32NoVal {
	if o == nil || IsNil(o.Total) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMaxJobs) GetTotalOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0040AssocMaxJobs) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given V0040Uint32NoVal and assigns it to the Total field.
func (o *V0040AssocMaxJobs) SetTotal(v V0040Uint32NoVal) {
	o.Total = &v
}

func (o V0040AssocMaxJobs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040AssocMaxJobs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Accruing) {
		toSerialize["accruing"] = o.Accruing
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableV0040AssocMaxJobs struct {
	value *V0040AssocMaxJobs
	isSet bool
}

func (v NullableV0040AssocMaxJobs) Get() *V0040AssocMaxJobs {
	return v.value
}

func (v *NullableV0040AssocMaxJobs) Set(val *V0040AssocMaxJobs) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040AssocMaxJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040AssocMaxJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040AssocMaxJobs(val *V0040AssocMaxJobs) *NullableV0040AssocMaxJobs {
	return &NullableV0040AssocMaxJobs{value: val, isSet: true}
}

func (v NullableV0040AssocMaxJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040AssocMaxJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


