/*
Copyright 2024.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	Version string `json:"version"`
	// +kubebuilder:validation:Enum=Calico;Kindnet;Flannel;Canel;Wave;Cilium
	CNI        string `json:"CNI"`
	MasterSize int    `json:"masterSize"`
	WorkerSize int    `json:"workerSize"`
}

type Node struct {
	HostName  string
	IPAddress string
}

type MasterNode struct {
	Node
}

type WorkerNode struct {
	Node
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	Phase     string             `json:"phase"`
	Master    []MasterNode       `json:"master"`
	Worker    []WorkerNode       `json:"worker"`
	Condition []ClusterCondition `json:"condition"`
}

type ClusterConditionType string

type ClusterCondition struct {
	Type          ClusterConditionType   `json:"type"`
	Status        corev1.ConditionStatus `json:"status"`
	LastProbeTime metav1.Time            `json:"lastProbeTime,omitempty" protobuf:"bytes,3,opt,name=lastProbeTime"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.version",name=VERSION,type=string
// +kubebuilder:printcolumn:JSONPath=".spec.CNI",name=CNI,type=string
// +kubebuilder:printcolumn:JSONPath=".spec.masterSize",name=MASTER,type=integer
// +kubebuilder:printcolumn:JSONPath=".spec.workerSize",name=WORK,type=integer

// Cluster is the Schema for the clusters API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterList contains a list of Cluster
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
