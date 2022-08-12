/*
caraml/upi/v1/upi.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// V1NamedValue struct for V1NamedValue
type V1NamedValue struct {
	Name *string `json:"name,omitempty"`
	Type *V1NamedValueType `json:"type,omitempty"`
	DoubleValue *float64 `json:"doubleValue,omitempty"`
	IntegerValue *string `json:"integerValue,omitempty"`
	StringValue *string `json:"stringValue,omitempty"`
}

// NewV1NamedValue instantiates a new V1NamedValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NamedValue() *V1NamedValue {
	this := V1NamedValue{}
	var type_ V1NamedValueType = V1NAMEDVALUETYPE_UNSPECIFIED
	this.Type = &type_
	return &this
}

// NewV1NamedValueWithDefaults instantiates a new V1NamedValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NamedValueWithDefaults() *V1NamedValue {
	this := V1NamedValue{}
	var type_ V1NamedValueType = V1NAMEDVALUETYPE_UNSPECIFIED
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1NamedValue) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NamedValue) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1NamedValue) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1NamedValue) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1NamedValue) GetType() V1NamedValueType {
	if o == nil || o.Type == nil {
		var ret V1NamedValueType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NamedValue) GetTypeOk() (*V1NamedValueType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1NamedValue) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given V1NamedValueType and assigns it to the Type field.
func (o *V1NamedValue) SetType(v V1NamedValueType) {
	o.Type = &v
}

// GetDoubleValue returns the DoubleValue field value if set, zero value otherwise.
func (o *V1NamedValue) GetDoubleValue() float64 {
	if o == nil || o.DoubleValue == nil {
		var ret float64
		return ret
	}
	return *o.DoubleValue
}

// GetDoubleValueOk returns a tuple with the DoubleValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NamedValue) GetDoubleValueOk() (*float64, bool) {
	if o == nil || o.DoubleValue == nil {
		return nil, false
	}
	return o.DoubleValue, true
}

// HasDoubleValue returns a boolean if a field has been set.
func (o *V1NamedValue) HasDoubleValue() bool {
	if o != nil && o.DoubleValue != nil {
		return true
	}

	return false
}

// SetDoubleValue gets a reference to the given float64 and assigns it to the DoubleValue field.
func (o *V1NamedValue) SetDoubleValue(v float64) {
	o.DoubleValue = &v
}

// GetIntegerValue returns the IntegerValue field value if set, zero value otherwise.
func (o *V1NamedValue) GetIntegerValue() string {
	if o == nil || o.IntegerValue == nil {
		var ret string
		return ret
	}
	return *o.IntegerValue
}

// GetIntegerValueOk returns a tuple with the IntegerValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NamedValue) GetIntegerValueOk() (*string, bool) {
	if o == nil || o.IntegerValue == nil {
		return nil, false
	}
	return o.IntegerValue, true
}

// HasIntegerValue returns a boolean if a field has been set.
func (o *V1NamedValue) HasIntegerValue() bool {
	if o != nil && o.IntegerValue != nil {
		return true
	}

	return false
}

// SetIntegerValue gets a reference to the given string and assigns it to the IntegerValue field.
func (o *V1NamedValue) SetIntegerValue(v string) {
	o.IntegerValue = &v
}

// GetStringValue returns the StringValue field value if set, zero value otherwise.
func (o *V1NamedValue) GetStringValue() string {
	if o == nil || o.StringValue == nil {
		var ret string
		return ret
	}
	return *o.StringValue
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NamedValue) GetStringValueOk() (*string, bool) {
	if o == nil || o.StringValue == nil {
		return nil, false
	}
	return o.StringValue, true
}

// HasStringValue returns a boolean if a field has been set.
func (o *V1NamedValue) HasStringValue() bool {
	if o != nil && o.StringValue != nil {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given string and assigns it to the StringValue field.
func (o *V1NamedValue) SetStringValue(v string) {
	o.StringValue = &v
}

func (o V1NamedValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.DoubleValue != nil {
		toSerialize["doubleValue"] = o.DoubleValue
	}
	if o.IntegerValue != nil {
		toSerialize["integerValue"] = o.IntegerValue
	}
	if o.StringValue != nil {
		toSerialize["stringValue"] = o.StringValue
	}
	return json.Marshal(toSerialize)
}

type NullableV1NamedValue struct {
	value *V1NamedValue
	isSet bool
}

func (v NullableV1NamedValue) Get() *V1NamedValue {
	return v.value
}

func (v *NullableV1NamedValue) Set(val *V1NamedValue) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NamedValue) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NamedValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NamedValue(val *V1NamedValue) *NullableV1NamedValue {
	return &NullableV1NamedValue{value: val, isSet: true}
}

func (v NullableV1NamedValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NamedValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


