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

// checks if the V0040PartitionInfoPriority type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040PartitionInfoPriority{}

// V0040PartitionInfoPriority struct for V0040PartitionInfoPriority
type V0040PartitionInfoPriority struct {
	// PriorityJobFactor
	JobFactor *int32 `json:"job_factor,omitempty"`
	// PriorityTier
	Tier *int32 `json:"tier,omitempty"`
}

// NewV0040PartitionInfoPriority instantiates a new V0040PartitionInfoPriority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040PartitionInfoPriority() *V0040PartitionInfoPriority {
	this := V0040PartitionInfoPriority{}
	return &this
}

// NewV0040PartitionInfoPriorityWithDefaults instantiates a new V0040PartitionInfoPriority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040PartitionInfoPriorityWithDefaults() *V0040PartitionInfoPriority {
	this := V0040PartitionInfoPriority{}
	return &this
}

// GetJobFactor returns the JobFactor field value if set, zero value otherwise.
func (o *V0040PartitionInfoPriority) GetJobFactor() int32 {
	if o == nil || IsNil(o.JobFactor) {
		var ret int32
		return ret
	}
	return *o.JobFactor
}

// GetJobFactorOk returns a tuple with the JobFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoPriority) GetJobFactorOk() (*int32, bool) {
	if o == nil || IsNil(o.JobFactor) {
		return nil, false
	}
	return o.JobFactor, true
}

// HasJobFactor returns a boolean if a field has been set.
func (o *V0040PartitionInfoPriority) HasJobFactor() bool {
	if o != nil && !IsNil(o.JobFactor) {
		return true
	}

	return false
}

// SetJobFactor gets a reference to the given int32 and assigns it to the JobFactor field.
func (o *V0040PartitionInfoPriority) SetJobFactor(v int32) {
	o.JobFactor = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *V0040PartitionInfoPriority) GetTier() int32 {
	if o == nil || IsNil(o.Tier) {
		var ret int32
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoPriority) GetTierOk() (*int32, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *V0040PartitionInfoPriority) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given int32 and assigns it to the Tier field.
func (o *V0040PartitionInfoPriority) SetTier(v int32) {
	o.Tier = &v
}

func (o V0040PartitionInfoPriority) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040PartitionInfoPriority) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobFactor) {
		toSerialize["job_factor"] = o.JobFactor
	}
	if !IsNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	return toSerialize, nil
}

type NullableV0040PartitionInfoPriority struct {
	value *V0040PartitionInfoPriority
	isSet bool
}

func (v NullableV0040PartitionInfoPriority) Get() *V0040PartitionInfoPriority {
	return v.value
}

func (v *NullableV0040PartitionInfoPriority) Set(val *V0040PartitionInfoPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040PartitionInfoPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040PartitionInfoPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040PartitionInfoPriority(val *V0040PartitionInfoPriority) *NullableV0040PartitionInfoPriority {
	return &NullableV0040PartitionInfoPriority{value: val, isSet: true}
}

func (v NullableV0040PartitionInfoPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040PartitionInfoPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


