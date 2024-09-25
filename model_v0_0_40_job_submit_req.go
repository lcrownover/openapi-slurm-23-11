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

// checks if the V0040JobSubmitReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040JobSubmitReq{}

// V0040JobSubmitReq struct for V0040JobSubmitReq
type V0040JobSubmitReq struct {
	// Batch job script; must be specified in first component of jobs or in job if this field is not populated
	Script *string `json:"script,omitempty"`
	Jobs []V0040JobDescMsg `json:"jobs,omitempty"`
	Job *V0040JobDescMsg `json:"job,omitempty"`
}

// NewV0040JobSubmitReq instantiates a new V0040JobSubmitReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040JobSubmitReq() *V0040JobSubmitReq {
	this := V0040JobSubmitReq{}
	return &this
}

// NewV0040JobSubmitReqWithDefaults instantiates a new V0040JobSubmitReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040JobSubmitReqWithDefaults() *V0040JobSubmitReq {
	this := V0040JobSubmitReq{}
	return &this
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *V0040JobSubmitReq) GetScript() string {
	if o == nil || IsNil(o.Script) {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobSubmitReq) GetScriptOk() (*string, bool) {
	if o == nil || IsNil(o.Script) {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *V0040JobSubmitReq) HasScript() bool {
	if o != nil && !IsNil(o.Script) {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *V0040JobSubmitReq) SetScript(v string) {
	o.Script = &v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *V0040JobSubmitReq) GetJobs() []V0040JobDescMsg {
	if o == nil || IsNil(o.Jobs) {
		var ret []V0040JobDescMsg
		return ret
	}
	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobSubmitReq) GetJobsOk() ([]V0040JobDescMsg, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *V0040JobSubmitReq) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []V0040JobDescMsg and assigns it to the Jobs field.
func (o *V0040JobSubmitReq) SetJobs(v []V0040JobDescMsg) {
	o.Jobs = v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0040JobSubmitReq) GetJob() V0040JobDescMsg {
	if o == nil || IsNil(o.Job) {
		var ret V0040JobDescMsg
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobSubmitReq) GetJobOk() (*V0040JobDescMsg, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0040JobSubmitReq) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given V0040JobDescMsg and assigns it to the Job field.
func (o *V0040JobSubmitReq) SetJob(v V0040JobDescMsg) {
	o.Job = &v
}

func (o V0040JobSubmitReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040JobSubmitReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Script) {
		toSerialize["script"] = o.Script
	}
	if !IsNil(o.Jobs) {
		toSerialize["jobs"] = o.Jobs
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableV0040JobSubmitReq struct {
	value *V0040JobSubmitReq
	isSet bool
}

func (v NullableV0040JobSubmitReq) Get() *V0040JobSubmitReq {
	return v.value
}

func (v *NullableV0040JobSubmitReq) Set(val *V0040JobSubmitReq) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040JobSubmitReq) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040JobSubmitReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040JobSubmitReq(val *V0040JobSubmitReq) *NullableV0040JobSubmitReq {
	return &NullableV0040JobSubmitReq{value: val, isSet: true}
}

func (v NullableV0040JobSubmitReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040JobSubmitReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


