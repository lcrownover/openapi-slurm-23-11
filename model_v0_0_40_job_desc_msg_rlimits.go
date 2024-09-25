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

// checks if the V0040JobDescMsgRlimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040JobDescMsgRlimits{}

// V0040JobDescMsgRlimits struct for V0040JobDescMsgRlimits
type V0040JobDescMsgRlimits struct {
	Cpu *V0040Uint64NoVal `json:"cpu,omitempty"`
	Fsize *V0040Uint64NoVal `json:"fsize,omitempty"`
	Data *V0040Uint64NoVal `json:"data,omitempty"`
	Stack *V0040Uint64NoVal `json:"stack,omitempty"`
	Core *V0040Uint64NoVal `json:"core,omitempty"`
	Rss *V0040Uint64NoVal `json:"rss,omitempty"`
	Nproc *V0040Uint64NoVal `json:"nproc,omitempty"`
	Nofile *V0040Uint64NoVal `json:"nofile,omitempty"`
	Memlock *V0040Uint64NoVal `json:"memlock,omitempty"`
	As *V0040Uint64NoVal `json:"as,omitempty"`
}

// NewV0040JobDescMsgRlimits instantiates a new V0040JobDescMsgRlimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040JobDescMsgRlimits() *V0040JobDescMsgRlimits {
	this := V0040JobDescMsgRlimits{}
	return &this
}

// NewV0040JobDescMsgRlimitsWithDefaults instantiates a new V0040JobDescMsgRlimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040JobDescMsgRlimitsWithDefaults() *V0040JobDescMsgRlimits {
	this := V0040JobDescMsgRlimits{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *V0040JobDescMsgRlimits) GetCpu() V0040Uint64NoVal {
	if o == nil || IsNil(o.Cpu) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobDescMsgRlimits) GetCpuOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *V0040JobDescMsgRlimits) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given V0040Uint64NoVal and assigns it to the Cpu field.
func (o *V0040JobDescMsgRlimits) SetCpu(v V0040Uint64NoVal) {
	o.Cpu = &v
}

// GetFsize returns the Fsize field value if set, zero value otherwise.
func (o *V0040JobDescMsgRlimits) GetFsize() V0040Uint64NoVal {
	if o == nil || IsNil(o.Fsize) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.Fsize
}

// GetFsizeOk returns a tuple with the Fsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobDescMsgRlimits) GetFsizeOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.Fsize) {
		return nil, false
	}
	return o.Fsize, true
}

// HasFsize returns a boolean if a field has been set.
func (o *V0040JobDescMsgRlimits) HasFsize() bool {
	if o != nil && !IsNil(o.Fsize) {
		return true
	}

	return false
}

// SetFsize gets a reference to the given V0040Uint64NoVal and assigns it to the Fsize field.
func (o *V0040JobDescMsgRlimits) SetFsize(v V0040Uint64NoVal) {
	o.Fsize = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *V0040JobDescMsgRlimits) GetData() V0040Uint64NoVal {
	if o == nil || IsNil(o.Data) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobDescMsgRlimits) GetDataOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *V0040JobDescMsgRlimits) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given V0040Uint64NoVal and assigns it to the Data field.
func (o *V0040JobDescMsgRlimits) SetData(v V0040Uint64NoVal) {
	o.Data = &v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *V0040JobDescMsgRlimits) GetStack() V0040Uint64NoVal {
	if o == nil || IsNil(o.Stack) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobDescMsgRlimits) GetStackOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.Stack) {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *V0040JobDescMsgRlimits) HasStack() bool {
	if o != nil && !IsNil(o.Stack) {
		return true
	}

	return false
}

// SetStack gets a reference to the given V0040Uint64NoVal and assigns it to the Stack field.
func (o *V0040JobDescMsgRlimits) SetStack(v V0040Uint64NoVal) {
	o.Stack = &v
}

// GetCore returns the Core field value if set, zero value otherwise.
func (o *V0040JobDescMsgRlimits) GetCore() V0040Uint64NoVal {
	if o == nil || IsNil(o.Core) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.Core
}

// GetCoreOk returns a tuple with the Core field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobDescMsgRlimits) GetCoreOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.Core) {
		return nil, false
	}
	return o.Core, true
}

// HasCore returns a boolean if a field has been set.
func (o *V0040JobDescMsgRlimits) HasCore() bool {
	if o != nil && !IsNil(o.Core) {
		return true
	}

	return false
}

// SetCore gets a reference to the given V0040Uint64NoVal and assigns it to the Core field.
func (o *V0040JobDescMsgRlimits) SetCore(v V0040Uint64NoVal) {
	o.Core = &v
}

// GetRss returns the Rss field value if set, zero value otherwise.
func (o *V0040JobDescMsgRlimits) GetRss() V0040Uint64NoVal {
	if o == nil || IsNil(o.Rss) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.Rss
}

// GetRssOk returns a tuple with the Rss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobDescMsgRlimits) GetRssOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.Rss) {
		return nil, false
	}
	return o.Rss, true
}

// HasRss returns a boolean if a field has been set.
func (o *V0040JobDescMsgRlimits) HasRss() bool {
	if o != nil && !IsNil(o.Rss) {
		return true
	}

	return false
}

// SetRss gets a reference to the given V0040Uint64NoVal and assigns it to the Rss field.
func (o *V0040JobDescMsgRlimits) SetRss(v V0040Uint64NoVal) {
	o.Rss = &v
}

// GetNproc returns the Nproc field value if set, zero value otherwise.
func (o *V0040JobDescMsgRlimits) GetNproc() V0040Uint64NoVal {
	if o == nil || IsNil(o.Nproc) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.Nproc
}

// GetNprocOk returns a tuple with the Nproc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobDescMsgRlimits) GetNprocOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.Nproc) {
		return nil, false
	}
	return o.Nproc, true
}

// HasNproc returns a boolean if a field has been set.
func (o *V0040JobDescMsgRlimits) HasNproc() bool {
	if o != nil && !IsNil(o.Nproc) {
		return true
	}

	return false
}

// SetNproc gets a reference to the given V0040Uint64NoVal and assigns it to the Nproc field.
func (o *V0040JobDescMsgRlimits) SetNproc(v V0040Uint64NoVal) {
	o.Nproc = &v
}

// GetNofile returns the Nofile field value if set, zero value otherwise.
func (o *V0040JobDescMsgRlimits) GetNofile() V0040Uint64NoVal {
	if o == nil || IsNil(o.Nofile) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.Nofile
}

// GetNofileOk returns a tuple with the Nofile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobDescMsgRlimits) GetNofileOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.Nofile) {
		return nil, false
	}
	return o.Nofile, true
}

// HasNofile returns a boolean if a field has been set.
func (o *V0040JobDescMsgRlimits) HasNofile() bool {
	if o != nil && !IsNil(o.Nofile) {
		return true
	}

	return false
}

// SetNofile gets a reference to the given V0040Uint64NoVal and assigns it to the Nofile field.
func (o *V0040JobDescMsgRlimits) SetNofile(v V0040Uint64NoVal) {
	o.Nofile = &v
}

// GetMemlock returns the Memlock field value if set, zero value otherwise.
func (o *V0040JobDescMsgRlimits) GetMemlock() V0040Uint64NoVal {
	if o == nil || IsNil(o.Memlock) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.Memlock
}

// GetMemlockOk returns a tuple with the Memlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobDescMsgRlimits) GetMemlockOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.Memlock) {
		return nil, false
	}
	return o.Memlock, true
}

// HasMemlock returns a boolean if a field has been set.
func (o *V0040JobDescMsgRlimits) HasMemlock() bool {
	if o != nil && !IsNil(o.Memlock) {
		return true
	}

	return false
}

// SetMemlock gets a reference to the given V0040Uint64NoVal and assigns it to the Memlock field.
func (o *V0040JobDescMsgRlimits) SetMemlock(v V0040Uint64NoVal) {
	o.Memlock = &v
}

// GetAs returns the As field value if set, zero value otherwise.
func (o *V0040JobDescMsgRlimits) GetAs() V0040Uint64NoVal {
	if o == nil || IsNil(o.As) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.As
}

// GetAsOk returns a tuple with the As field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobDescMsgRlimits) GetAsOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.As) {
		return nil, false
	}
	return o.As, true
}

// HasAs returns a boolean if a field has been set.
func (o *V0040JobDescMsgRlimits) HasAs() bool {
	if o != nil && !IsNil(o.As) {
		return true
	}

	return false
}

// SetAs gets a reference to the given V0040Uint64NoVal and assigns it to the As field.
func (o *V0040JobDescMsgRlimits) SetAs(v V0040Uint64NoVal) {
	o.As = &v
}

func (o V0040JobDescMsgRlimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040JobDescMsgRlimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Fsize) {
		toSerialize["fsize"] = o.Fsize
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Stack) {
		toSerialize["stack"] = o.Stack
	}
	if !IsNil(o.Core) {
		toSerialize["core"] = o.Core
	}
	if !IsNil(o.Rss) {
		toSerialize["rss"] = o.Rss
	}
	if !IsNil(o.Nproc) {
		toSerialize["nproc"] = o.Nproc
	}
	if !IsNil(o.Nofile) {
		toSerialize["nofile"] = o.Nofile
	}
	if !IsNil(o.Memlock) {
		toSerialize["memlock"] = o.Memlock
	}
	if !IsNil(o.As) {
		toSerialize["as"] = o.As
	}
	return toSerialize, nil
}

type NullableV0040JobDescMsgRlimits struct {
	value *V0040JobDescMsgRlimits
	isSet bool
}

func (v NullableV0040JobDescMsgRlimits) Get() *V0040JobDescMsgRlimits {
	return v.value
}

func (v *NullableV0040JobDescMsgRlimits) Set(val *V0040JobDescMsgRlimits) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040JobDescMsgRlimits) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040JobDescMsgRlimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040JobDescMsgRlimits(val *V0040JobDescMsgRlimits) *NullableV0040JobDescMsgRlimits {
	return &NullableV0040JobDescMsgRlimits{value: val, isSet: true}
}

func (v NullableV0040JobDescMsgRlimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040JobDescMsgRlimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


