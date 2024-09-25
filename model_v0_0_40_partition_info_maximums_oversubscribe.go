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

// checks if the V0040PartitionInfoMaximumsOversubscribe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040PartitionInfoMaximumsOversubscribe{}

// V0040PartitionInfoMaximumsOversubscribe struct for V0040PartitionInfoMaximumsOversubscribe
type V0040PartitionInfoMaximumsOversubscribe struct {
	// Maximum number of jobs allowed to oversubscribe resources
	Jobs *int32 `json:"jobs,omitempty"`
	// Flags applicable to the OverSubscribe setting
	Flags []string `json:"flags,omitempty"`
}

// NewV0040PartitionInfoMaximumsOversubscribe instantiates a new V0040PartitionInfoMaximumsOversubscribe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040PartitionInfoMaximumsOversubscribe() *V0040PartitionInfoMaximumsOversubscribe {
	this := V0040PartitionInfoMaximumsOversubscribe{}
	return &this
}

// NewV0040PartitionInfoMaximumsOversubscribeWithDefaults instantiates a new V0040PartitionInfoMaximumsOversubscribe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040PartitionInfoMaximumsOversubscribeWithDefaults() *V0040PartitionInfoMaximumsOversubscribe {
	this := V0040PartitionInfoMaximumsOversubscribe{}
	return &this
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *V0040PartitionInfoMaximumsOversubscribe) GetJobs() int32 {
	if o == nil || IsNil(o.Jobs) {
		var ret int32
		return ret
	}
	return *o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoMaximumsOversubscribe) GetJobsOk() (*int32, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *V0040PartitionInfoMaximumsOversubscribe) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given int32 and assigns it to the Jobs field.
func (o *V0040PartitionInfoMaximumsOversubscribe) SetJobs(v int32) {
	o.Jobs = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0040PartitionInfoMaximumsOversubscribe) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoMaximumsOversubscribe) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0040PartitionInfoMaximumsOversubscribe) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0040PartitionInfoMaximumsOversubscribe) SetFlags(v []string) {
	o.Flags = v
}

func (o V0040PartitionInfoMaximumsOversubscribe) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040PartitionInfoMaximumsOversubscribe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Jobs) {
		toSerialize["jobs"] = o.Jobs
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

type NullableV0040PartitionInfoMaximumsOversubscribe struct {
	value *V0040PartitionInfoMaximumsOversubscribe
	isSet bool
}

func (v NullableV0040PartitionInfoMaximumsOversubscribe) Get() *V0040PartitionInfoMaximumsOversubscribe {
	return v.value
}

func (v *NullableV0040PartitionInfoMaximumsOversubscribe) Set(val *V0040PartitionInfoMaximumsOversubscribe) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040PartitionInfoMaximumsOversubscribe) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040PartitionInfoMaximumsOversubscribe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040PartitionInfoMaximumsOversubscribe(val *V0040PartitionInfoMaximumsOversubscribe) *NullableV0040PartitionInfoMaximumsOversubscribe {
	return &NullableV0040PartitionInfoMaximumsOversubscribe{value: val, isSet: true}
}

func (v NullableV0040PartitionInfoMaximumsOversubscribe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040PartitionInfoMaximumsOversubscribe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


