/*
 * Workload API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// SecurityDSCStatus struct for SecurityDSCStatus
type SecurityDSCStatus struct {
	// DSC ID for which the agent error or warning is issued.
	DscId *string `json:"dsc-id,omitempty"`
	// InfoStatus contains agent message the operation is failed or warning is issued.
	DscInfoStatus *string `json:"dsc-info-status,omitempty"`
	// PolicyEntriesConsumed shows number of policy entries consumed in the DSC by this policy.
	PolicyEntriesConsumed *string `json:"policy-entries-consumed,omitempty"`
	// RuleEntriesConsumed shows number of rule entries consumed in the DSE by this policy.
	RuleEntriesConsumed *string `json:"rule-entries-consumed,omitempty"`
}

// NewSecurityDSCStatus instantiates a new SecurityDSCStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityDSCStatus() *SecurityDSCStatus {
	this := SecurityDSCStatus{}
	return &this
}

// NewSecurityDSCStatusWithDefaults instantiates a new SecurityDSCStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityDSCStatusWithDefaults() *SecurityDSCStatus {
	this := SecurityDSCStatus{}
	return &this
}

// GetDscId returns the DscId field value if set, zero value otherwise.
func (o *SecurityDSCStatus) GetDscId() string {
	if o == nil || o.DscId == nil {
		var ret string
		return ret
	}
	return *o.DscId
}

// GetDscIdOk returns a tuple with the DscId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityDSCStatus) GetDscIdOk() (*string, bool) {
	if o == nil || o.DscId == nil {
		return nil, false
	}
	return o.DscId, true
}

// HasDscId returns a boolean if a field has been set.
func (o *SecurityDSCStatus) HasDscId() bool {
	if o != nil && o.DscId != nil {
		return true
	}

	return false
}

// SetDscId gets a reference to the given string and assigns it to the DscId field.
func (o *SecurityDSCStatus) SetDscId(v string) {
	o.DscId = &v
}

// GetDscInfoStatus returns the DscInfoStatus field value if set, zero value otherwise.
func (o *SecurityDSCStatus) GetDscInfoStatus() string {
	if o == nil || o.DscInfoStatus == nil {
		var ret string
		return ret
	}
	return *o.DscInfoStatus
}

// GetDscInfoStatusOk returns a tuple with the DscInfoStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityDSCStatus) GetDscInfoStatusOk() (*string, bool) {
	if o == nil || o.DscInfoStatus == nil {
		return nil, false
	}
	return o.DscInfoStatus, true
}

// HasDscInfoStatus returns a boolean if a field has been set.
func (o *SecurityDSCStatus) HasDscInfoStatus() bool {
	if o != nil && o.DscInfoStatus != nil {
		return true
	}

	return false
}

// SetDscInfoStatus gets a reference to the given string and assigns it to the DscInfoStatus field.
func (o *SecurityDSCStatus) SetDscInfoStatus(v string) {
	o.DscInfoStatus = &v
}

// GetPolicyEntriesConsumed returns the PolicyEntriesConsumed field value if set, zero value otherwise.
func (o *SecurityDSCStatus) GetPolicyEntriesConsumed() string {
	if o == nil || o.PolicyEntriesConsumed == nil {
		var ret string
		return ret
	}
	return *o.PolicyEntriesConsumed
}

// GetPolicyEntriesConsumedOk returns a tuple with the PolicyEntriesConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityDSCStatus) GetPolicyEntriesConsumedOk() (*string, bool) {
	if o == nil || o.PolicyEntriesConsumed == nil {
		return nil, false
	}
	return o.PolicyEntriesConsumed, true
}

// HasPolicyEntriesConsumed returns a boolean if a field has been set.
func (o *SecurityDSCStatus) HasPolicyEntriesConsumed() bool {
	if o != nil && o.PolicyEntriesConsumed != nil {
		return true
	}

	return false
}

// SetPolicyEntriesConsumed gets a reference to the given string and assigns it to the PolicyEntriesConsumed field.
func (o *SecurityDSCStatus) SetPolicyEntriesConsumed(v string) {
	o.PolicyEntriesConsumed = &v
}

// GetRuleEntriesConsumed returns the RuleEntriesConsumed field value if set, zero value otherwise.
func (o *SecurityDSCStatus) GetRuleEntriesConsumed() string {
	if o == nil || o.RuleEntriesConsumed == nil {
		var ret string
		return ret
	}
	return *o.RuleEntriesConsumed
}

// GetRuleEntriesConsumedOk returns a tuple with the RuleEntriesConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityDSCStatus) GetRuleEntriesConsumedOk() (*string, bool) {
	if o == nil || o.RuleEntriesConsumed == nil {
		return nil, false
	}
	return o.RuleEntriesConsumed, true
}

// HasRuleEntriesConsumed returns a boolean if a field has been set.
func (o *SecurityDSCStatus) HasRuleEntriesConsumed() bool {
	if o != nil && o.RuleEntriesConsumed != nil {
		return true
	}

	return false
}

// SetRuleEntriesConsumed gets a reference to the given string and assigns it to the RuleEntriesConsumed field.
func (o *SecurityDSCStatus) SetRuleEntriesConsumed(v string) {
	o.RuleEntriesConsumed = &v
}

func (o SecurityDSCStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DscId != nil {
		toSerialize["dsc-id"] = o.DscId
	}
	if o.DscInfoStatus != nil {
		toSerialize["dsc-info-status"] = o.DscInfoStatus
	}
	if o.PolicyEntriesConsumed != nil {
		toSerialize["policy-entries-consumed"] = o.PolicyEntriesConsumed
	}
	if o.RuleEntriesConsumed != nil {
		toSerialize["rule-entries-consumed"] = o.RuleEntriesConsumed
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityDSCStatus struct {
	value *SecurityDSCStatus
	isSet bool
}

func (v NullableSecurityDSCStatus) Get() *SecurityDSCStatus {
	return v.value
}

func (v *NullableSecurityDSCStatus) Set(val *SecurityDSCStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityDSCStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityDSCStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityDSCStatus(val *SecurityDSCStatus) *NullableSecurityDSCStatus {
	return &NullableSecurityDSCStatus{value: val, isSet: true}
}

func (v NullableSecurityDSCStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityDSCStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


