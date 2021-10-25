/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SynapseRoleAssignmentObservation struct {
}

type SynapseRoleAssignmentParameters struct {

	// +kubebuilder:validation:Required
	PrincipalID *string `json:"principalId" tf:"principal_id,omitempty"`

	// +kubebuilder:validation:Required
	RoleName *string `json:"roleName" tf:"role_name,omitempty"`

	// +kubebuilder:validation:Optional
	SynapseSparkPoolID *string `json:"synapseSparkPoolId,omitempty" tf:"synapse_spark_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`
}

// SynapseRoleAssignmentSpec defines the desired state of SynapseRoleAssignment
type SynapseRoleAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SynapseRoleAssignmentParameters `json:"forProvider"`
}

// SynapseRoleAssignmentStatus defines the observed state of SynapseRoleAssignment.
type SynapseRoleAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SynapseRoleAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SynapseRoleAssignment is the Schema for the SynapseRoleAssignments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SynapseRoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SynapseRoleAssignmentSpec   `json:"spec"`
	Status            SynapseRoleAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SynapseRoleAssignmentList contains a list of SynapseRoleAssignments
type SynapseRoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SynapseRoleAssignment `json:"items"`
}

// Repository type metadata.
var (
	SynapseRoleAssignmentKind             = "SynapseRoleAssignment"
	SynapseRoleAssignmentGroupKind        = schema.GroupKind{Group: Group, Kind: SynapseRoleAssignmentKind}.String()
	SynapseRoleAssignmentKindAPIVersion   = SynapseRoleAssignmentKind + "." + GroupVersion.String()
	SynapseRoleAssignmentGroupVersionKind = GroupVersion.WithKind(SynapseRoleAssignmentKind)
)

func init() {
	SchemeBuilder.Register(&SynapseRoleAssignment{}, &SynapseRoleAssignmentList{})
}