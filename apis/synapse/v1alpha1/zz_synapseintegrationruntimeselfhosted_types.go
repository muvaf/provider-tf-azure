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

type SynapseIntegrationRuntimeSelfHostedObservation struct {
	AuthorizationKeyPrimary *string `json:"authorizationKeyPrimary,omitempty" tf:"authorization_key_primary,omitempty"`

	AuthorizationKeySecondary *string `json:"authorizationKeySecondary,omitempty" tf:"authorization_key_secondary,omitempty"`
}

type SynapseIntegrationRuntimeSelfHostedParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	SynapseWorkspaceID *string `json:"synapseWorkspaceId" tf:"synapse_workspace_id,omitempty"`
}

// SynapseIntegrationRuntimeSelfHostedSpec defines the desired state of SynapseIntegrationRuntimeSelfHosted
type SynapseIntegrationRuntimeSelfHostedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SynapseIntegrationRuntimeSelfHostedParameters `json:"forProvider"`
}

// SynapseIntegrationRuntimeSelfHostedStatus defines the observed state of SynapseIntegrationRuntimeSelfHosted.
type SynapseIntegrationRuntimeSelfHostedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SynapseIntegrationRuntimeSelfHostedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SynapseIntegrationRuntimeSelfHosted is the Schema for the SynapseIntegrationRuntimeSelfHosteds API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SynapseIntegrationRuntimeSelfHosted struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SynapseIntegrationRuntimeSelfHostedSpec   `json:"spec"`
	Status            SynapseIntegrationRuntimeSelfHostedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SynapseIntegrationRuntimeSelfHostedList contains a list of SynapseIntegrationRuntimeSelfHosteds
type SynapseIntegrationRuntimeSelfHostedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SynapseIntegrationRuntimeSelfHosted `json:"items"`
}

// Repository type metadata.
var (
	SynapseIntegrationRuntimeSelfHostedKind             = "SynapseIntegrationRuntimeSelfHosted"
	SynapseIntegrationRuntimeSelfHostedGroupKind        = schema.GroupKind{Group: Group, Kind: SynapseIntegrationRuntimeSelfHostedKind}.String()
	SynapseIntegrationRuntimeSelfHostedKindAPIVersion   = SynapseIntegrationRuntimeSelfHostedKind + "." + GroupVersion.String()
	SynapseIntegrationRuntimeSelfHostedGroupVersionKind = GroupVersion.WithKind(SynapseIntegrationRuntimeSelfHostedKind)
)

func init() {
	SchemeBuilder.Register(&SynapseIntegrationRuntimeSelfHosted{}, &SynapseIntegrationRuntimeSelfHostedList{})
}