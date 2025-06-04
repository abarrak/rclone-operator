package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RcloneSpec defines the desired state of Rclone.
type RcloneSpec struct {
	// Configuration for the rclone task.
	// +kubebuilder:validation:Required
	Configuration RcloneConfiguration `json:"configuration,omitempty"`
	// Type specifies the type of rclone operation.
	Type string `json:"type,omitempty"`
}

// RcloneStatus defines the observed state of Rclone.
type RcloneStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
	Active     string             `json:"active"`
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

type RcloneConfiguration struct {
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
