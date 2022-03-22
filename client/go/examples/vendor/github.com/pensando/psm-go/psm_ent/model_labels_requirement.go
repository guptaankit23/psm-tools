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

// LabelsRequirement Requirement defines a single matching condition for a selector.
type LabelsRequirement struct {
	// The label key that the condition applies to.
	Key *string `json:"key,omitempty"`
	// Condition checked for the key.
	Operator *string `json:"operator,omitempty"`
	// Values contains one or more values corresponding to the label key. \"equals\" and \"notEquals\" operators need a single Value. \"in\" and \"notIn\" operators can have one or more values.
	Values *[]string `json:"values,omitempty"`
}

// NewLabelsRequirement instantiates a new LabelsRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelsRequirement() *LabelsRequirement {
	this := LabelsRequirement{}
	var operator string = "equals"
	this.Operator = &operator
	return &this
}

// NewLabelsRequirementWithDefaults instantiates a new LabelsRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelsRequirementWithDefaults() *LabelsRequirement {
	this := LabelsRequirement{}
	var operator string = "equals"
	this.Operator = &operator
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *LabelsRequirement) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsRequirement) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *LabelsRequirement) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *LabelsRequirement) SetKey(v string) {
	o.Key = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *LabelsRequirement) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsRequirement) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *LabelsRequirement) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *LabelsRequirement) SetOperator(v string) {
	o.Operator = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *LabelsRequirement) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsRequirement) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *LabelsRequirement) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *LabelsRequirement) SetValues(v []string) {
	o.Values = &v
}

func (o LabelsRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableLabelsRequirement struct {
	value *LabelsRequirement
	isSet bool
}

func (v NullableLabelsRequirement) Get() *LabelsRequirement {
	return v.value
}

func (v *NullableLabelsRequirement) Set(val *LabelsRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelsRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelsRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelsRequirement(val *LabelsRequirement) *NullableLabelsRequirement {
	return &NullableLabelsRequirement{value: val, isSet: true}
}

func (v NullableLabelsRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelsRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


