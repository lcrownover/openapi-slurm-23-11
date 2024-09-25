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

// checks if the V0040Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040Account{}

// V0040Account struct for V0040Account
type V0040Account struct {
	Associations []V0040AssocShort `json:"associations,omitempty"`
	Coordinators []V0040Coord `json:"coordinators,omitempty"`
	// Arbitrary string describing the account
	Description string `json:"description"`
	// Account name
	Name string `json:"name"`
	// Organization to which the account belongs
	Organization string `json:"organization"`
	// Flags associated with the account
	Flags []string `json:"flags,omitempty"`
}

type _V0040Account V0040Account

// NewV0040Account instantiates a new V0040Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040Account(description string, name string, organization string) *V0040Account {
	this := V0040Account{}
	this.Description = description
	this.Name = name
	this.Organization = organization
	return &this
}

// NewV0040AccountWithDefaults instantiates a new V0040Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AccountWithDefaults() *V0040Account {
	this := V0040Account{}
	return &this
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *V0040Account) GetAssociations() []V0040AssocShort {
	if o == nil || IsNil(o.Associations) {
		var ret []V0040AssocShort
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040Account) GetAssociationsOk() ([]V0040AssocShort, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *V0040Account) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []V0040AssocShort and assigns it to the Associations field.
func (o *V0040Account) SetAssociations(v []V0040AssocShort) {
	o.Associations = v
}

// GetCoordinators returns the Coordinators field value if set, zero value otherwise.
func (o *V0040Account) GetCoordinators() []V0040Coord {
	if o == nil || IsNil(o.Coordinators) {
		var ret []V0040Coord
		return ret
	}
	return o.Coordinators
}

// GetCoordinatorsOk returns a tuple with the Coordinators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040Account) GetCoordinatorsOk() ([]V0040Coord, bool) {
	if o == nil || IsNil(o.Coordinators) {
		return nil, false
	}
	return o.Coordinators, true
}

// HasCoordinators returns a boolean if a field has been set.
func (o *V0040Account) HasCoordinators() bool {
	if o != nil && !IsNil(o.Coordinators) {
		return true
	}

	return false
}

// SetCoordinators gets a reference to the given []V0040Coord and assigns it to the Coordinators field.
func (o *V0040Account) SetCoordinators(v []V0040Coord) {
	o.Coordinators = v
}

// GetDescription returns the Description field value
func (o *V0040Account) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *V0040Account) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *V0040Account) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *V0040Account) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V0040Account) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V0040Account) SetName(v string) {
	o.Name = v
}

// GetOrganization returns the Organization field value
func (o *V0040Account) GetOrganization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *V0040Account) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *V0040Account) SetOrganization(v string) {
	o.Organization = v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0040Account) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040Account) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0040Account) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0040Account) SetFlags(v []string) {
	o.Flags = v
}

func (o V0040Account) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.Coordinators) {
		toSerialize["coordinators"] = o.Coordinators
	}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["organization"] = o.Organization
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

func (o *V0040Account) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"name",
		"organization",
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

	varV0040Account := _V0040Account{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040Account)

	if err != nil {
		return err
	}

	*o = V0040Account(varV0040Account)

	return err
}

type NullableV0040Account struct {
	value *V0040Account
	isSet bool
}

func (v NullableV0040Account) Get() *V0040Account {
	return v.value
}

func (v *NullableV0040Account) Set(val *V0040Account) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040Account) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040Account) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040Account(val *V0040Account) *NullableV0040Account {
	return &NullableV0040Account{value: val, isSet: true}
}

func (v NullableV0040Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040Account) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


