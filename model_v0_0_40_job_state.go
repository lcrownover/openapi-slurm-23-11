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

// checks if the V0040JobState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040JobState{}

// V0040JobState struct for V0040JobState
type V0040JobState struct {
	// Current state
	Current []string `json:"current,omitempty"`
	// Reason for previous Pending or Failed state
	Reason *string `json:"reason,omitempty"`
}

// NewV0040JobState instantiates a new V0040JobState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040JobState() *V0040JobState {
	this := V0040JobState{}
	return &this
}

// NewV0040JobStateWithDefaults instantiates a new V0040JobState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040JobStateWithDefaults() *V0040JobState {
	this := V0040JobState{}
	return &this
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *V0040JobState) GetCurrent() []string {
	if o == nil || IsNil(o.Current) {
		var ret []string
		return ret
	}
	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobState) GetCurrentOk() ([]string, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *V0040JobState) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given []string and assigns it to the Current field.
func (o *V0040JobState) SetCurrent(v []string) {
	o.Current = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *V0040JobState) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobState) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *V0040JobState) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *V0040JobState) SetReason(v string) {
	o.Reason = &v
}

func (o V0040JobState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040JobState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableV0040JobState struct {
	value *V0040JobState
	isSet bool
}

func (v NullableV0040JobState) Get() *V0040JobState {
	return v.value
}

func (v *NullableV0040JobState) Set(val *V0040JobState) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040JobState) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040JobState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040JobState(val *V0040JobState) *NullableV0040JobState {
	return &NullableV0040JobState{value: val, isSet: true}
}

func (v NullableV0040JobState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040JobState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


