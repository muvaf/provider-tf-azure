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

type AdvancedThreatProtectionObservation struct {
}

type AdvancedThreatProtectionParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	TargetResourceID *string `json:"targetResourceId" tf:"target_resource_id,omitempty"`
}

// AdvancedThreatProtectionSpec defines the desired state of AdvancedThreatProtection
type AdvancedThreatProtectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AdvancedThreatProtectionParameters `json:"forProvider"`
}

// AdvancedThreatProtectionStatus defines the observed state of AdvancedThreatProtection.
type AdvancedThreatProtectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AdvancedThreatProtectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AdvancedThreatProtection is the Schema for the AdvancedThreatProtections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AdvancedThreatProtection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AdvancedThreatProtectionSpec   `json:"spec"`
	Status            AdvancedThreatProtectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AdvancedThreatProtectionList contains a list of AdvancedThreatProtections
type AdvancedThreatProtectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdvancedThreatProtection `json:"items"`
}

// Repository type metadata.
var (
	AdvancedThreatProtectionKind             = "AdvancedThreatProtection"
	AdvancedThreatProtectionGroupKind        = schema.GroupKind{Group: Group, Kind: AdvancedThreatProtectionKind}.String()
	AdvancedThreatProtectionKindAPIVersion   = AdvancedThreatProtectionKind + "." + GroupVersion.String()
	AdvancedThreatProtectionGroupVersionKind = GroupVersion.WithKind(AdvancedThreatProtectionKind)
)

func init() {
	SchemeBuilder.Register(&AdvancedThreatProtection{}, &AdvancedThreatProtectionList{})
}