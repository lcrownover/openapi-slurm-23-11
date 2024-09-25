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

// checks if the V0040PartitionInfoAccounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040PartitionInfoAccounts{}

// V0040PartitionInfoAccounts struct for V0040PartitionInfoAccounts
type V0040PartitionInfoAccounts struct {
	// AllowAccounts
	Allowed *string `json:"allowed,omitempty"`
	// DenyAccounts
	Deny *string `json:"deny,omitempty"`
}

// NewV0040PartitionInfoAccounts instantiates a new V0040PartitionInfoAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040PartitionInfoAccounts() *V0040PartitionInfoAccounts {
	this := V0040PartitionInfoAccounts{}
	return &this
}

// NewV0040PartitionInfoAccountsWithDefaults instantiates a new V0040PartitionInfoAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040PartitionInfoAccountsWithDefaults() *V0040PartitionInfoAccounts {
	this := V0040PartitionInfoAccounts{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *V0040PartitionInfoAccounts) GetAllowed() string {
	if o == nil || IsNil(o.Allowed) {
		var ret string
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoAccounts) GetAllowedOk() (*string, bool) {
	if o == nil || IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *V0040PartitionInfoAccounts) HasAllowed() bool {
	if o != nil && !IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given string and assigns it to the Allowed field.
func (o *V0040PartitionInfoAccounts) SetAllowed(v string) {
	o.Allowed = &v
}

// GetDeny returns the Deny field value if set, zero value otherwise.
func (o *V0040PartitionInfoAccounts) GetDeny() string {
	if o == nil || IsNil(o.Deny) {
		var ret string
		return ret
	}
	return *o.Deny
}

// GetDenyOk returns a tuple with the Deny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoAccounts) GetDenyOk() (*string, bool) {
	if o == nil || IsNil(o.Deny) {
		return nil, false
	}
	return o.Deny, true
}

// HasDeny returns a boolean if a field has been set.
func (o *V0040PartitionInfoAccounts) HasDeny() bool {
	if o != nil && !IsNil(o.Deny) {
		return true
	}

	return false
}

// SetDeny gets a reference to the given string and assigns it to the Deny field.
func (o *V0040PartitionInfoAccounts) SetDeny(v string) {
	o.Deny = &v
}

func (o V0040PartitionInfoAccounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040PartitionInfoAccounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	if !IsNil(o.Deny) {
		toSerialize["deny"] = o.Deny
	}
	return toSerialize, nil
}

type NullableV0040PartitionInfoAccounts struct {
	value *V0040PartitionInfoAccounts
	isSet bool
}

func (v NullableV0040PartitionInfoAccounts) Get() *V0040PartitionInfoAccounts {
	return v.value
}

func (v *NullableV0040PartitionInfoAccounts) Set(val *V0040PartitionInfoAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040PartitionInfoAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040PartitionInfoAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040PartitionInfoAccounts(val *V0040PartitionInfoAccounts) *NullableV0040PartitionInfoAccounts {
	return &NullableV0040PartitionInfoAccounts{value: val, isSet: true}
}

func (v NullableV0040PartitionInfoAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040PartitionInfoAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


