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

// checks if the V0040PartitionInfoQos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040PartitionInfoQos{}

// V0040PartitionInfoQos struct for V0040PartitionInfoQos
type V0040PartitionInfoQos struct {
	// AllowQOS
	Allowed *string `json:"allowed,omitempty"`
	// DenyQOS
	Deny *string `json:"deny,omitempty"`
	// QOS
	Assigned *string `json:"assigned,omitempty"`
}

// NewV0040PartitionInfoQos instantiates a new V0040PartitionInfoQos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040PartitionInfoQos() *V0040PartitionInfoQos {
	this := V0040PartitionInfoQos{}
	return &this
}

// NewV0040PartitionInfoQosWithDefaults instantiates a new V0040PartitionInfoQos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040PartitionInfoQosWithDefaults() *V0040PartitionInfoQos {
	this := V0040PartitionInfoQos{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *V0040PartitionInfoQos) GetAllowed() string {
	if o == nil || IsNil(o.Allowed) {
		var ret string
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoQos) GetAllowedOk() (*string, bool) {
	if o == nil || IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *V0040PartitionInfoQos) HasAllowed() bool {
	if o != nil && !IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given string and assigns it to the Allowed field.
func (o *V0040PartitionInfoQos) SetAllowed(v string) {
	o.Allowed = &v
}

// GetDeny returns the Deny field value if set, zero value otherwise.
func (o *V0040PartitionInfoQos) GetDeny() string {
	if o == nil || IsNil(o.Deny) {
		var ret string
		return ret
	}
	return *o.Deny
}

// GetDenyOk returns a tuple with the Deny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoQos) GetDenyOk() (*string, bool) {
	if o == nil || IsNil(o.Deny) {
		return nil, false
	}
	return o.Deny, true
}

// HasDeny returns a boolean if a field has been set.
func (o *V0040PartitionInfoQos) HasDeny() bool {
	if o != nil && !IsNil(o.Deny) {
		return true
	}

	return false
}

// SetDeny gets a reference to the given string and assigns it to the Deny field.
func (o *V0040PartitionInfoQos) SetDeny(v string) {
	o.Deny = &v
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *V0040PartitionInfoQos) GetAssigned() string {
	if o == nil || IsNil(o.Assigned) {
		var ret string
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoQos) GetAssignedOk() (*string, bool) {
	if o == nil || IsNil(o.Assigned) {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *V0040PartitionInfoQos) HasAssigned() bool {
	if o != nil && !IsNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given string and assigns it to the Assigned field.
func (o *V0040PartitionInfoQos) SetAssigned(v string) {
	o.Assigned = &v
}

func (o V0040PartitionInfoQos) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040PartitionInfoQos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	if !IsNil(o.Deny) {
		toSerialize["deny"] = o.Deny
	}
	if !IsNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	return toSerialize, nil
}

type NullableV0040PartitionInfoQos struct {
	value *V0040PartitionInfoQos
	isSet bool
}

func (v NullableV0040PartitionInfoQos) Get() *V0040PartitionInfoQos {
	return v.value
}

func (v *NullableV0040PartitionInfoQos) Set(val *V0040PartitionInfoQos) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040PartitionInfoQos) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040PartitionInfoQos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040PartitionInfoQos(val *V0040PartitionInfoQos) *NullableV0040PartitionInfoQos {
	return &NullableV0040PartitionInfoQos{value: val, isSet: true}
}

func (v NullableV0040PartitionInfoQos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040PartitionInfoQos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


