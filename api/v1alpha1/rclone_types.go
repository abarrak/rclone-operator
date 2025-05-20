/*
Copyright 2025.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RcloneSpec defines the desired state of Rclone.
type RcloneSpec struct {
	// +kubebuilder:validation:Required
	Config RcloneConfig `json:"config,omitempty"`
	Type   string       `json:"type,omitempty"`
}

// RcloneStatus defines the observed state of Rclone.
type RcloneStatus struct {
	// Represents the observations of a RcloneStatus's current state.
	// RcloneStatus.status.conditions.type are: "Available", "Progressing", and "Degraded"
	// RcloneStatus.status.conditions.status are one of True, False, Unknown.
	// RcloneStatus.status.conditions.reason the value should be a CamelCase string and producers of specific
	// condition types may define expected values and meanings for this field, and whether the values
	// are considered a guaranteed API.
	// RcloneStatus.status.conditions.Message is a human readable message indicating details about the transition.
	// For further information see: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`

	// PodName of the active rclone task.
	Active string `json:"active"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Rclone is the Schema for the rclones API.
type Rclone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RcloneSpec   `json:"spec,omitempty"`
	Status RcloneStatus `json:"status,omitempty"`
}

type RcloneConfig struct {
	// +kubebuilder:validation:Required
	Config string `json:"config,omitempty"`
	// +kubebuilder:validation:Required
	Command string `json:"command,omitempty"`
}

// +kubebuilder:object:root=true

// RcloneList contains a list of Rclone.
type RcloneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rclone `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Rclone{}, &RcloneList{})
}
