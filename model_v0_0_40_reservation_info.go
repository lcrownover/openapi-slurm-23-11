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

// checks if the V0040ReservationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040ReservationInfo{}

// V0040ReservationInfo struct for V0040ReservationInfo
type V0040ReservationInfo struct {
	// Comma separated list of permitted accounts
	Accounts *string `json:"accounts,omitempty"`
	// BurstBuffer
	BurstBuffer *string `json:"burst_buffer,omitempty"`
	// CoreCnt
	CoreCount *int32 `json:"core_count,omitempty"`
	CoreSpecializations []V0040ReservationCoreSpec `json:"core_specializations,omitempty"`
	EndTime *V0040Uint64NoVal `json:"end_time,omitempty"`
	// Features
	Features *string `json:"features,omitempty"`
	// Flags associated with the reservation
	Flags []string `json:"flags,omitempty"`
	// Groups
	Groups *string `json:"groups,omitempty"`
	// Licenses
	Licenses *string `json:"licenses,omitempty"`
	// MaxStartDelay in seconds
	MaxStartDelay *int32 `json:"max_start_delay,omitempty"`
	// ReservationName
	Name *string `json:"name,omitempty"`
	// NodeCnt
	NodeCount *int32 `json:"node_count,omitempty"`
	// Nodes
	NodeList *string `json:"node_list,omitempty"`
	// PartitionName
	Partition *string `json:"partition,omitempty"`
	PurgeCompleted *V0040ReservationInfoPurgeCompleted `json:"purge_completed,omitempty"`
	StartTime *V0040Uint64NoVal `json:"start_time,omitempty"`
	Watts *V0040Uint32NoVal `json:"watts,omitempty"`
	// Comma separated list of required TRES
	Tres *string `json:"tres,omitempty"`
	// Comma separated list of permitted users
	Users *string `json:"users,omitempty"`
}

// NewV0040ReservationInfo instantiates a new V0040ReservationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040ReservationInfo() *V0040ReservationInfo {
	this := V0040ReservationInfo{}
	return &this
}

// NewV0040ReservationInfoWithDefaults instantiates a new V0040ReservationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040ReservationInfoWithDefaults() *V0040ReservationInfo {
	this := V0040ReservationInfo{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetAccounts() string {
	if o == nil || IsNil(o.Accounts) {
		var ret string
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetAccountsOk() (*string, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given string and assigns it to the Accounts field.
func (o *V0040ReservationInfo) SetAccounts(v string) {
	o.Accounts = &v
}

// GetBurstBuffer returns the BurstBuffer field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetBurstBuffer() string {
	if o == nil || IsNil(o.BurstBuffer) {
		var ret string
		return ret
	}
	return *o.BurstBuffer
}

// GetBurstBufferOk returns a tuple with the BurstBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetBurstBufferOk() (*string, bool) {
	if o == nil || IsNil(o.BurstBuffer) {
		return nil, false
	}
	return o.BurstBuffer, true
}

// HasBurstBuffer returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasBurstBuffer() bool {
	if o != nil && !IsNil(o.BurstBuffer) {
		return true
	}

	return false
}

// SetBurstBuffer gets a reference to the given string and assigns it to the BurstBuffer field.
func (o *V0040ReservationInfo) SetBurstBuffer(v string) {
	o.BurstBuffer = &v
}

// GetCoreCount returns the CoreCount field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetCoreCount() int32 {
	if o == nil || IsNil(o.CoreCount) {
		var ret int32
		return ret
	}
	return *o.CoreCount
}

// GetCoreCountOk returns a tuple with the CoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetCoreCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CoreCount) {
		return nil, false
	}
	return o.CoreCount, true
}

// HasCoreCount returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasCoreCount() bool {
	if o != nil && !IsNil(o.CoreCount) {
		return true
	}

	return false
}

// SetCoreCount gets a reference to the given int32 and assigns it to the CoreCount field.
func (o *V0040ReservationInfo) SetCoreCount(v int32) {
	o.CoreCount = &v
}

// GetCoreSpecializations returns the CoreSpecializations field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetCoreSpecializations() []V0040ReservationCoreSpec {
	if o == nil || IsNil(o.CoreSpecializations) {
		var ret []V0040ReservationCoreSpec
		return ret
	}
	return o.CoreSpecializations
}

// GetCoreSpecializationsOk returns a tuple with the CoreSpecializations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetCoreSpecializationsOk() ([]V0040ReservationCoreSpec, bool) {
	if o == nil || IsNil(o.CoreSpecializations) {
		return nil, false
	}
	return o.CoreSpecializations, true
}

// HasCoreSpecializations returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasCoreSpecializations() bool {
	if o != nil && !IsNil(o.CoreSpecializations) {
		return true
	}

	return false
}

// SetCoreSpecializations gets a reference to the given []V0040ReservationCoreSpec and assigns it to the CoreSpecializations field.
func (o *V0040ReservationInfo) SetCoreSpecializations(v []V0040ReservationCoreSpec) {
	o.CoreSpecializations = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetEndTime() V0040Uint64NoVal {
	if o == nil || IsNil(o.EndTime) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetEndTimeOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given V0040Uint64NoVal and assigns it to the EndTime field.
func (o *V0040ReservationInfo) SetEndTime(v V0040Uint64NoVal) {
	o.EndTime = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetFeatures() string {
	if o == nil || IsNil(o.Features) {
		var ret string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given string and assigns it to the Features field.
func (o *V0040ReservationInfo) SetFeatures(v string) {
	o.Features = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0040ReservationInfo) SetFlags(v []string) {
	o.Flags = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetGroups() string {
	if o == nil || IsNil(o.Groups) {
		var ret string
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetGroupsOk() (*string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given string and assigns it to the Groups field.
func (o *V0040ReservationInfo) SetGroups(v string) {
	o.Groups = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetLicenses() string {
	if o == nil || IsNil(o.Licenses) {
		var ret string
		return ret
	}
	return *o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetLicensesOk() (*string, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given string and assigns it to the Licenses field.
func (o *V0040ReservationInfo) SetLicenses(v string) {
	o.Licenses = &v
}

// GetMaxStartDelay returns the MaxStartDelay field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetMaxStartDelay() int32 {
	if o == nil || IsNil(o.MaxStartDelay) {
		var ret int32
		return ret
	}
	return *o.MaxStartDelay
}

// GetMaxStartDelayOk returns a tuple with the MaxStartDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetMaxStartDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxStartDelay) {
		return nil, false
	}
	return o.MaxStartDelay, true
}

// HasMaxStartDelay returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasMaxStartDelay() bool {
	if o != nil && !IsNil(o.MaxStartDelay) {
		return true
	}

	return false
}

// SetMaxStartDelay gets a reference to the given int32 and assigns it to the MaxStartDelay field.
func (o *V0040ReservationInfo) SetMaxStartDelay(v int32) {
	o.MaxStartDelay = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0040ReservationInfo) SetName(v string) {
	o.Name = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetNodeCount() int32 {
	if o == nil || IsNil(o.NodeCount) {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetNodeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *V0040ReservationInfo) SetNodeCount(v int32) {
	o.NodeCount = &v
}

// GetNodeList returns the NodeList field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetNodeList() string {
	if o == nil || IsNil(o.NodeList) {
		var ret string
		return ret
	}
	return *o.NodeList
}

// GetNodeListOk returns a tuple with the NodeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetNodeListOk() (*string, bool) {
	if o == nil || IsNil(o.NodeList) {
		return nil, false
	}
	return o.NodeList, true
}

// HasNodeList returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasNodeList() bool {
	if o != nil && !IsNil(o.NodeList) {
		return true
	}

	return false
}

// SetNodeList gets a reference to the given string and assigns it to the NodeList field.
func (o *V0040ReservationInfo) SetNodeList(v string) {
	o.NodeList = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetPartition() string {
	if o == nil || IsNil(o.Partition) {
		var ret string
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetPartitionOk() (*string, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given string and assigns it to the Partition field.
func (o *V0040ReservationInfo) SetPartition(v string) {
	o.Partition = &v
}

// GetPurgeCompleted returns the PurgeCompleted field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetPurgeCompleted() V0040ReservationInfoPurgeCompleted {
	if o == nil || IsNil(o.PurgeCompleted) {
		var ret V0040ReservationInfoPurgeCompleted
		return ret
	}
	return *o.PurgeCompleted
}

// GetPurgeCompletedOk returns a tuple with the PurgeCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetPurgeCompletedOk() (*V0040ReservationInfoPurgeCompleted, bool) {
	if o == nil || IsNil(o.PurgeCompleted) {
		return nil, false
	}
	return o.PurgeCompleted, true
}

// HasPurgeCompleted returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasPurgeCompleted() bool {
	if o != nil && !IsNil(o.PurgeCompleted) {
		return true
	}

	return false
}

// SetPurgeCompleted gets a reference to the given V0040ReservationInfoPurgeCompleted and assigns it to the PurgeCompleted field.
func (o *V0040ReservationInfo) SetPurgeCompleted(v V0040ReservationInfoPurgeCompleted) {
	o.PurgeCompleted = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetStartTime() V0040Uint64NoVal {
	if o == nil || IsNil(o.StartTime) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetStartTimeOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given V0040Uint64NoVal and assigns it to the StartTime field.
func (o *V0040ReservationInfo) SetStartTime(v V0040Uint64NoVal) {
	o.StartTime = &v
}

// GetWatts returns the Watts field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetWatts() V0040Uint32NoVal {
	if o == nil || IsNil(o.Watts) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Watts
}

// GetWattsOk returns a tuple with the Watts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetWattsOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Watts) {
		return nil, false
	}
	return o.Watts, true
}

// HasWatts returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasWatts() bool {
	if o != nil && !IsNil(o.Watts) {
		return true
	}

	return false
}

// SetWatts gets a reference to the given V0040Uint32NoVal and assigns it to the Watts field.
func (o *V0040ReservationInfo) SetWatts(v V0040Uint32NoVal) {
	o.Watts = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetTres() string {
	if o == nil || IsNil(o.Tres) {
		var ret string
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetTresOk() (*string, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given string and assigns it to the Tres field.
func (o *V0040ReservationInfo) SetTres(v string) {
	o.Tres = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *V0040ReservationInfo) GetUsers() string {
	if o == nil || IsNil(o.Users) {
		var ret string
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfo) GetUsersOk() (*string, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *V0040ReservationInfo) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given string and assigns it to the Users field.
func (o *V0040ReservationInfo) SetUsers(v string) {
	o.Users = &v
}

func (o V0040ReservationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040ReservationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.BurstBuffer) {
		toSerialize["burst_buffer"] = o.BurstBuffer
	}
	if !IsNil(o.CoreCount) {
		toSerialize["core_count"] = o.CoreCount
	}
	if !IsNil(o.CoreSpecializations) {
		toSerialize["core_specializations"] = o.CoreSpecializations
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	if !IsNil(o.MaxStartDelay) {
		toSerialize["max_start_delay"] = o.MaxStartDelay
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NodeCount) {
		toSerialize["node_count"] = o.NodeCount
	}
	if !IsNil(o.NodeList) {
		toSerialize["node_list"] = o.NodeList
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.PurgeCompleted) {
		toSerialize["purge_completed"] = o.PurgeCompleted
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Watts) {
		toSerialize["watts"] = o.Watts
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableV0040ReservationInfo struct {
	value *V0040ReservationInfo
	isSet bool
}

func (v NullableV0040ReservationInfo) Get() *V0040ReservationInfo {
	return v.value
}

func (v *NullableV0040ReservationInfo) Set(val *V0040ReservationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040ReservationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040ReservationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040ReservationInfo(val *V0040ReservationInfo) *NullableV0040ReservationInfo {
	return &NullableV0040ReservationInfo{value: val, isSet: true}
}

func (v NullableV0040ReservationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040ReservationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


