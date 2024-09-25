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
)

// checks if the V0040PartitionInfoCpus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040PartitionInfoCpus{}

// V0040PartitionInfoCpus struct for V0040PartitionInfoCpus
type V0040PartitionInfoCpus struct {
	// CpuBind
	TaskBinding *int32 `json:"task_binding,omitempty"`
	// TotalCPUs
	Total *int32 `json:"total,omitempty"`
}

// NewV0040PartitionInfoCpus instantiates a new V0040PartitionInfoCpus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040PartitionInfoCpus() *V0040PartitionInfoCpus {
	this := V0040PartitionInfoCpus{}
	return &this
}

// NewV0040PartitionInfoCpusWithDefaults instantiates a new V0040PartitionInfoCpus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040PartitionInfoCpusWithDefaults() *V0040PartitionInfoCpus {
	this := V0040PartitionInfoCpus{}
	return &this
}

// GetTaskBinding returns the TaskBinding field value if set, zero value otherwise.
func (o *V0040PartitionInfoCpus) GetTaskBinding() int32 {
	if o == nil || IsNil(o.TaskBinding) {
		var ret int32
		return ret
	}
	return *o.TaskBinding
}

// GetTaskBindingOk returns a tuple with the TaskBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoCpus) GetTaskBindingOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskBinding) {
		return nil, false
	}
	return o.TaskBinding, true
}

// HasTaskBinding returns a boolean if a field has been set.
func (o *V0040PartitionInfoCpus) HasTaskBinding() bool {
	if o != nil && !IsNil(o.TaskBinding) {
		return true
	}

	return false
}

// SetTaskBinding gets a reference to the given int32 and assigns it to the TaskBinding field.
func (o *V0040PartitionInfoCpus) SetTaskBinding(v int32) {
	o.TaskBinding = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0040PartitionInfoCpus) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoCpus) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0040PartitionInfoCpus) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *V0040PartitionInfoCpus) SetTotal(v int32) {
	o.Total = &v
}

func (o V0040PartitionInfoCpus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040PartitionInfoCpus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaskBinding) {
		toSerialize["task_binding"] = o.TaskBinding
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableV0040PartitionInfoCpus struct {
	value *V0040PartitionInfoCpus
	isSet bool
}

func (v NullableV0040PartitionInfoCpus) Get() *V0040PartitionInfoCpus {
	return v.value
}

func (v *NullableV0040PartitionInfoCpus) Set(val *V0040PartitionInfoCpus) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040PartitionInfoCpus) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040PartitionInfoCpus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040PartitionInfoCpus(val *V0040PartitionInfoCpus) *NullableV0040PartitionInfoCpus {
	return &NullableV0040PartitionInfoCpus{value: val, isSet: true}
}

func (v NullableV0040PartitionInfoCpus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040PartitionInfoCpus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


