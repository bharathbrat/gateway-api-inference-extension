/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha2

import (
	apiv1alpha2 "sigs.k8s.io/gateway-api-inference-extension/api/v1alpha2"
)

// PoolObjectReferenceApplyConfiguration represents a declarative configuration of the PoolObjectReference type for use
// with apply.
type PoolObjectReferenceApplyConfiguration struct {
	Group *apiv1alpha2.Group      `json:"group,omitempty"`
	Kind  *apiv1alpha2.Kind       `json:"kind,omitempty"`
	Name  *apiv1alpha2.ObjectName `json:"name,omitempty"`
}

// PoolObjectReferenceApplyConfiguration constructs a declarative configuration of the PoolObjectReference type for use with
// apply.
func PoolObjectReference() *PoolObjectReferenceApplyConfiguration {
	return &PoolObjectReferenceApplyConfiguration{}
}

// WithGroup sets the Group field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Group field is set to the value of the last call.
func (b *PoolObjectReferenceApplyConfiguration) WithGroup(value apiv1alpha2.Group) *PoolObjectReferenceApplyConfiguration {
	b.Group = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *PoolObjectReferenceApplyConfiguration) WithKind(value apiv1alpha2.Kind) *PoolObjectReferenceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PoolObjectReferenceApplyConfiguration) WithName(value apiv1alpha2.ObjectName) *PoolObjectReferenceApplyConfiguration {
	b.Name = &value
	return b
}
