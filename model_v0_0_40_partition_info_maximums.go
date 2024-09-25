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

// checks if the V0040PartitionInfoMaximums type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040PartitionInfoMaximums{}

// V0040PartitionInfoMaximums struct for V0040PartitionInfoMaximums
type V0040PartitionInfoMaximums struct {
	CpusPerNode *V0040Uint32NoVal `json:"cpus_per_node,omitempty"`
	CpusPerSocket *V0040Uint32NoVal `json:"cpus_per_socket,omitempty"`
	// MaxMemPerCPU or MaxMemPerNode
	MemoryPerCpu *int64 `json:"memory_per_cpu,omitempty"`
	PartitionMemoryPerCpu *V0040Uint64NoVal `json:"partition_memory_per_cpu,omitempty"`
	PartitionMemoryPerNode *V0040Uint64NoVal `json:"partition_memory_per_node,omitempty"`
	Nodes *V0040Uint32NoVal `json:"nodes,omitempty"`
	// OverSubscribe
	Shares *int32 `json:"shares,omitempty"`
	Oversubscribe *V0040PartitionInfoMaximumsOversubscribe `json:"oversubscribe,omitempty"`
	Time *V0040Uint32NoVal `json:"time,omitempty"`
	OverTimeLimit *V0040Uint16NoVal `json:"over_time_limit,omitempty"`
}

// NewV0040PartitionInfoMaximums instantiates a new V0040PartitionInfoMaximums object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040PartitionInfoMaximums() *V0040PartitionInfoMaximums {
	this := V0040PartitionInfoMaximums{}
	return &this
}

// NewV0040PartitionInfoMaximumsWithDefaults instantiates a new V0040PartitionInfoMaximums object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040PartitionInfoMaximumsWithDefaults() *V0040PartitionInfoMaximums {
	this := V0040PartitionInfoMaximums{}
	return &this
}

// GetCpusPerNode returns the CpusPerNode field value if set, zero value otherwise.
func (o *V0040PartitionInfoMaximums) GetCpusPerNode() V0040Uint32NoVal {
	if o == nil || IsNil(o.CpusPerNode) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.CpusPerNode
}

// GetCpusPerNodeOk returns a tuple with the CpusPerNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoMaximums) GetCpusPerNodeOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.CpusPerNode) {
		return nil, false
	}
	return o.CpusPerNode, true
}

// HasCpusPerNode returns a boolean if a field has been set.
func (o *V0040PartitionInfoMaximums) HasCpusPerNode() bool {
	if o != nil && !IsNil(o.CpusPerNode) {
		return true
	}

	return false
}

// SetCpusPerNode gets a reference to the given V0040Uint32NoVal and assigns it to the CpusPerNode field.
func (o *V0040PartitionInfoMaximums) SetCpusPerNode(v V0040Uint32NoVal) {
	o.CpusPerNode = &v
}

// GetCpusPerSocket returns the CpusPerSocket field value if set, zero value otherwise.
func (o *V0040PartitionInfoMaximums) GetCpusPerSocket() V0040Uint32NoVal {
	if o == nil || IsNil(o.CpusPerSocket) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.CpusPerSocket
}

// GetCpusPerSocketOk returns a tuple with the CpusPerSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoMaximums) GetCpusPerSocketOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.CpusPerSocket) {
		return nil, false
	}
	return o.CpusPerSocket, true
}

// HasCpusPerSocket returns a boolean if a field has been set.
func (o *V0040PartitionInfoMaximums) HasCpusPerSocket() bool {
	if o != nil && !IsNil(o.CpusPerSocket) {
		return true
	}

	return false
}

// SetCpusPerSocket gets a reference to the given V0040Uint32NoVal and assigns it to the CpusPerSocket field.
func (o *V0040PartitionInfoMaximums) SetCpusPerSocket(v V0040Uint32NoVal) {
	o.CpusPerSocket = &v
}

// GetMemoryPerCpu returns the MemoryPerCpu field value if set, zero value otherwise.
func (o *V0040PartitionInfoMaximums) GetMemoryPerCpu() int64 {
	if o == nil || IsNil(o.MemoryPerCpu) {
		var ret int64
		return ret
	}
	return *o.MemoryPerCpu
}

// GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoMaximums) GetMemoryPerCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.MemoryPerCpu) {
		return nil, false
	}
	return o.MemoryPerCpu, true
}

// HasMemoryPerCpu returns a boolean if a field has been set.
func (o *V0040PartitionInfoMaximums) HasMemoryPerCpu() bool {
	if o != nil && !IsNil(o.MemoryPerCpu) {
		return true
	}

	return false
}

// SetMemoryPerCpu gets a reference to the given int64 and assigns it to the MemoryPerCpu field.
func (o *V0040PartitionInfoMaximums) SetMemoryPerCpu(v int64) {
	o.MemoryPerCpu = &v
}

// GetPartitionMemoryPerCpu returns the PartitionMemoryPerCpu field value if set, zero value otherwise.
func (o *V0040PartitionInfoMaximums) GetPartitionMemoryPerCpu() V0040Uint64NoVal {
	if o == nil || IsNil(o.PartitionMemoryPerCpu) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.PartitionMemoryPerCpu
}

// GetPartitionMemoryPerCpuOk returns a tuple with the PartitionMemoryPerCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoMaximums) GetPartitionMemoryPerCpuOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.PartitionMemoryPerCpu) {
		return nil, false
	}
	return o.PartitionMemoryPerCpu, true
}

// HasPartitionMemoryPerCpu returns a boolean if a field has been set.
func (o *V0040PartitionInfoMaximums) HasPartitionMemoryPerCpu() bool {
	if o != nil && !IsNil(o.PartitionMemoryPerCpu) {
		return true
	}

	return false
}

// SetPartitionMemoryPerCpu gets a reference to the given V0040Uint64NoVal and assigns it to the PartitionMemoryPerCpu field.
func (o *V0040PartitionInfoMaximums) SetPartitionMemoryPerCpu(v V0040Uint64NoVal) {
	o.PartitionMemoryPerCpu = &v
}

// GetPartitionMemoryPerNode returns the PartitionMemoryPerNode field value if set, zero value otherwise.
func (o *V0040PartitionInfoMaximums) GetPartitionMemoryPerNode() V0040Uint64NoVal {
	if o == nil || IsNil(o.PartitionMemoryPerNode) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.PartitionMemoryPerNode
}

// GetPartitionMemoryPerNodeOk returns a tuple with the PartitionMemoryPerNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoMaximums) GetPartitionMemoryPerNodeOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.PartitionMemoryPerNode) {
		return nil, false
	}
	return o.PartitionMemoryPerNode, true
}

// HasPartitionMemoryPerNode returns a boolean if a field has been set.
func (o *V0040PartitionInfoMaximums) HasPartitionMemoryPerNode() bool {
	if o != nil && !IsNil(o.PartitionMemoryPerNode) {
		return true
	}

	return false
}

// SetPartitionMemoryPerNode gets a reference to the given V0040Uint64NoVal and assigns it to the PartitionMemoryPerNode field.
func (o *V0040PartitionInfoMaximums) SetPartitionMemoryPerNode(v V0040Uint64NoVal) {
	o.PartitionMemoryPerNode = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *V0040PartitionInfoMaximums) GetNodes() V0040Uint32NoVal {
	if o == nil || IsNil(o.Nodes) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoMaximums) GetNodesOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *V0040PartitionInfoMaximums) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given V0040Uint32NoVal and assigns it to the Nodes field.
func (o *V0040PartitionInfoMaximums) SetNodes(v V0040Uint32NoVal) {
	o.Nodes = &v
}

// GetShares returns the Shares field value if set, zero value otherwise.
func (o *V0040PartitionInfoMaximums) GetShares() int32 {
	if o == nil || IsNil(o.Shares) {
		var ret int32
		return ret
	}
	return *o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoMaximums) GetSharesOk() (*int32, bool) {
	if o == nil || IsNil(o.Shares) {
		return nil, false
	}
	return o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *V0040PartitionInfoMaximums) HasShares() bool {
	if o != nil && !IsNil(o.Shares) {
		return true
	}

	return false
}

// SetShares gets a reference to the given int32 and assigns it to the Shares field.
func (o *V0040PartitionInfoMaximums) SetShares(v int32) {
	o.Shares = &v
}

// GetOversubscribe returns the Oversubscribe field value if set, zero value otherwise.
func (o *V0040PartitionInfoMaximums) GetOversubscribe() V0040PartitionInfoMaximumsOversubscribe {
	if o == nil || IsNil(o.Oversubscribe) {
		var ret V0040PartitionInfoMaximumsOversubscribe
		return ret
	}
	return *o.Oversubscribe
}

// GetOversubscribeOk returns a tuple with the Oversubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoMaximums) GetOversubscribeOk() (*V0040PartitionInfoMaximumsOversubscribe, bool) {
	if o == nil || IsNil(o.Oversubscribe) {
		return nil, false
	}
	return o.Oversubscribe, true
}

// HasOversubscribe returns a boolean if a field has been set.
func (o *V0040PartitionInfoMaximums) HasOversubscribe() bool {
	if o != nil && !IsNil(o.Oversubscribe) {
		return true
	}

	return false
}

// SetOversubscribe gets a reference to the given V0040PartitionInfoMaximumsOversubscribe and assigns it to the Oversubscribe field.
func (o *V0040PartitionInfoMaximums) SetOversubscribe(v V0040PartitionInfoMaximumsOversubscribe) {
	o.Oversubscribe = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V0040PartitionInfoMaximums) GetTime() V0040Uint32NoVal {
	if o == nil || IsNil(o.Time) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoMaximums) GetTimeOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V0040PartitionInfoMaximums) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given V0040Uint32NoVal and assigns it to the Time field.
func (o *V0040PartitionInfoMaximums) SetTime(v V0040Uint32NoVal) {
	o.Time = &v
}

// GetOverTimeLimit returns the OverTimeLimit field value if set, zero value otherwise.
func (o *V0040PartitionInfoMaximums) GetOverTimeLimit() V0040Uint16NoVal {
	if o == nil || IsNil(o.OverTimeLimit) {
		var ret V0040Uint16NoVal
		return ret
	}
	return *o.OverTimeLimit
}

// GetOverTimeLimitOk returns a tuple with the OverTimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoMaximums) GetOverTimeLimitOk() (*V0040Uint16NoVal, bool) {
	if o == nil || IsNil(o.OverTimeLimit) {
		return nil, false
	}
	return o.OverTimeLimit, true
}

// HasOverTimeLimit returns a boolean if a field has been set.
func (o *V0040PartitionInfoMaximums) HasOverTimeLimit() bool {
	if o != nil && !IsNil(o.OverTimeLimit) {
		return true
	}

	return false
}

// SetOverTimeLimit gets a reference to the given V0040Uint16NoVal and assigns it to the OverTimeLimit field.
func (o *V0040PartitionInfoMaximums) SetOverTimeLimit(v V0040Uint16NoVal) {
	o.OverTimeLimit = &v
}

func (o V0040PartitionInfoMaximums) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040PartitionInfoMaximums) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpusPerNode) {
		toSerialize["cpus_per_node"] = o.CpusPerNode
	}
	if !IsNil(o.CpusPerSocket) {
		toSerialize["cpus_per_socket"] = o.CpusPerSocket
	}
	if !IsNil(o.MemoryPerCpu) {
		toSerialize["memory_per_cpu"] = o.MemoryPerCpu
	}
	if !IsNil(o.PartitionMemoryPerCpu) {
		toSerialize["partition_memory_per_cpu"] = o.PartitionMemoryPerCpu
	}
	if !IsNil(o.PartitionMemoryPerNode) {
		toSerialize["partition_memory_per_node"] = o.PartitionMemoryPerNode
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.Shares) {
		toSerialize["shares"] = o.Shares
	}
	if !IsNil(o.Oversubscribe) {
		toSerialize["oversubscribe"] = o.Oversubscribe
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.OverTimeLimit) {
		toSerialize["over_time_limit"] = o.OverTimeLimit
	}
	return toSerialize, nil
}

type NullableV0040PartitionInfoMaximums struct {
	value *V0040PartitionInfoMaximums
	isSet bool
}

func (v NullableV0040PartitionInfoMaximums) Get() *V0040PartitionInfoMaximums {
	return v.value
}

func (v *NullableV0040PartitionInfoMaximums) Set(val *V0040PartitionInfoMaximums) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040PartitionInfoMaximums) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040PartitionInfoMaximums) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040PartitionInfoMaximums(val *V0040PartitionInfoMaximums) *NullableV0040PartitionInfoMaximums {
	return &NullableV0040PartitionInfoMaximums{value: val, isSet: true}
}

func (v NullableV0040PartitionInfoMaximums) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040PartitionInfoMaximums) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


