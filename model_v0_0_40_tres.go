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
	"bytes"
	"fmt"
)

// checks if the V0040Tres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040Tres{}

// V0040Tres struct for V0040Tres
type V0040Tres struct {
	// TRES type (CPU, MEM, etc)
	Type string `json:"type"`
	// TRES name (if applicable)
	Name *string `json:"name,omitempty"`
	// ID used in database
	Id *int32 `json:"id,omitempty"`
	// TRES count (0 if listed generically)
	Count *int64 `json:"count,omitempty"`
}

type _V0040Tres V0040Tres

// NewV0040Tres instantiates a new V0040Tres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040Tres(type_ string) *V0040Tres {
	this := V0040Tres{}
	this.Type = type_
	return &this
}

// NewV0040TresWithDefaults instantiates a new V0040Tres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040TresWithDefaults() *V0040Tres {
	this := V0040Tres{}
	return &this
}

// GetType returns the Type field value
func (o *V0040Tres) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V0040Tres) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V0040Tres) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0040Tres) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040Tres) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0040Tres) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0040Tres) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0040Tres) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040Tres) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0040Tres) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *V0040Tres) SetId(v int32) {
	o.Id = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0040Tres) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040Tres) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0040Tres) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *V0040Tres) SetCount(v int64) {
	o.Count = &v
}

func (o V0040Tres) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040Tres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

func (o *V0040Tres) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV0040Tres := _V0040Tres{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040Tres)

	if err != nil {
		return err
	}

	*o = V0040Tres(varV0040Tres)

	return err
}

type NullableV0040Tres struct {
	value *V0040Tres
	isSet bool
}

func (v NullableV0040Tres) Get() *V0040Tres {
	return v.value
}

func (v *NullableV0040Tres) Set(val *V0040Tres) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040Tres) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040Tres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040Tres(val *V0040Tres) *NullableV0040Tres {
	return &NullableV0040Tres{value: val, isSet: true}
}

func (v NullableV0040Tres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040Tres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


