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
	"bytes"
	"fmt"
)

// checks if the V0040OpenapiTresResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiTresResp{}

// V0040OpenapiTresResp struct for V0040OpenapiTresResp
type V0040OpenapiTresResp struct {
	TRES []V0040Tres `json:"TRES"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiTresResp V0040OpenapiTresResp

// NewV0040OpenapiTresResp instantiates a new V0040OpenapiTresResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiTresResp(tRES []V0040Tres) *V0040OpenapiTresResp {
	this := V0040OpenapiTresResp{}
	this.TRES = tRES
	return &this
}

// NewV0040OpenapiTresRespWithDefaults instantiates a new V0040OpenapiTresResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiTresRespWithDefaults() *V0040OpenapiTresResp {
	this := V0040OpenapiTresResp{}
	return &this
}

// GetTRES returns the TRES field value
func (o *V0040OpenapiTresResp) GetTRES() []V0040Tres {
	if o == nil {
		var ret []V0040Tres
		return ret
	}

	return o.TRES
}

// GetTRESOk returns a tuple with the TRES field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiTresResp) GetTRESOk() ([]V0040Tres, bool) {
	if o == nil {
		return nil, false
	}
	return o.TRES, true
}

// SetTRES sets field value
func (o *V0040OpenapiTresResp) SetTRES(v []V0040Tres) {
	o.TRES = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiTresResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiTresResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiTresResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiTresResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiTresResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiTresResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiTresResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiTresResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiTresResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiTresResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiTresResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiTresResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiTresResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiTresResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["TRES"] = o.TRES
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *V0040OpenapiTresResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"TRES",
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

	varV0040OpenapiTresResp := _V0040OpenapiTresResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiTresResp)

	if err != nil {
		return err
	}

	*o = V0040OpenapiTresResp(varV0040OpenapiTresResp)

	return err
}

type NullableV0040OpenapiTresResp struct {
	value *V0040OpenapiTresResp
	isSet bool
}

func (v NullableV0040OpenapiTresResp) Get() *V0040OpenapiTresResp {
	return v.value
}

func (v *NullableV0040OpenapiTresResp) Set(val *V0040OpenapiTresResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiTresResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiTresResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiTresResp(val *V0040OpenapiTresResp) *NullableV0040OpenapiTresResp {
	return &NullableV0040OpenapiTresResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiTresResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiTresResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


