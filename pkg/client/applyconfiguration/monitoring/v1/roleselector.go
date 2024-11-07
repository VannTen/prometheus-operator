// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// RoleSelectorApplyConfiguration represents a declarative configuration of the RoleSelector type for use
// with apply.
type RoleSelectorApplyConfiguration struct {
	Role   *string                                     `json:"role,omitempty"`
	Labels []RoleSelectorRequirementApplyConfiguration `json:"labels,omitempty"`
	Fields []RoleSelectorRequirementApplyConfiguration `json:"fields,omitempty"`
}

// RoleSelectorApplyConfiguration constructs a declarative configuration of the RoleSelector type for use with
// apply.
func RoleSelector() *RoleSelectorApplyConfiguration {
	return &RoleSelectorApplyConfiguration{}
}

// WithRole sets the Role field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Role field is set to the value of the last call.
func (b *RoleSelectorApplyConfiguration) WithRole(value string) *RoleSelectorApplyConfiguration {
	b.Role = &value
	return b
}

// WithLabels adds the given value to the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Labels field.
func (b *RoleSelectorApplyConfiguration) WithLabels(values ...*RoleSelectorRequirementApplyConfiguration) *RoleSelectorApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithLabels")
		}
		b.Labels = append(b.Labels, *values[i])
	}
	return b
}

// WithFields adds the given value to the Fields field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Fields field.
func (b *RoleSelectorApplyConfiguration) WithFields(values ...*RoleSelectorRequirementApplyConfiguration) *RoleSelectorApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFields")
		}
		b.Fields = append(b.Fields, *values[i])
	}
	return b
}
