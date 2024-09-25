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

// checks if the V0040ClusterRecAssociations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040ClusterRecAssociations{}

// V0040ClusterRecAssociations struct for V0040ClusterRecAssociations
type V0040ClusterRecAssociations struct {
	Root *V0040AssocShort `json:"root,omitempty"`
}

// NewV0040ClusterRecAssociations instantiates a new V0040ClusterRecAssociations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040ClusterRecAssociations() *V0040ClusterRecAssociations {
	this := V0040ClusterRecAssociations{}
	return &this
}

// NewV0040ClusterRecAssociationsWithDefaults instantiates a new V0040ClusterRecAssociations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040ClusterRecAssociationsWithDefaults() *V0040ClusterRecAssociations {
	this := V0040ClusterRecAssociations{}
	return &this
}

// GetRoot returns the Root field value if set, zero value otherwise.
func (o *V0040ClusterRecAssociations) GetRoot() V0040AssocShort {
	if o == nil || IsNil(o.Root) {
		var ret V0040AssocShort
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ClusterRecAssociations) GetRootOk() (*V0040AssocShort, bool) {
	if o == nil || IsNil(o.Root) {
		return nil, false
	}
	return o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *V0040ClusterRecAssociations) HasRoot() bool {
	if o != nil && !IsNil(o.Root) {
		return true
	}

	return false
}

// SetRoot gets a reference to the given V0040AssocShort and assigns it to the Root field.
func (o *V0040ClusterRecAssociations) SetRoot(v V0040AssocShort) {
	o.Root = &v
}

func (o V0040ClusterRecAssociations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040ClusterRecAssociations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Root) {
		toSerialize["root"] = o.Root
	}
	return toSerialize, nil
}

type NullableV0040ClusterRecAssociations struct {
	value *V0040ClusterRecAssociations
	isSet bool
}

func (v NullableV0040ClusterRecAssociations) Get() *V0040ClusterRecAssociations {
	return v.value
}

func (v *NullableV0040ClusterRecAssociations) Set(val *V0040ClusterRecAssociations) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040ClusterRecAssociations) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040ClusterRecAssociations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040ClusterRecAssociations(val *V0040ClusterRecAssociations) *NullableV0040ClusterRecAssociations {
	return &NullableV0040ClusterRecAssociations{value: val, isSet: true}
}

func (v NullableV0040ClusterRecAssociations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040ClusterRecAssociations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


