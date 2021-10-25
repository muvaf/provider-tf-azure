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

type EventhubNamespaceAuthorizationRuleObservation struct {
}

type EventhubNamespaceAuthorizationRuleParameters struct {

	// +kubebuilder:validation:Optional
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`

	// +kubebuilder:validation:Optional
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NamespaceName *string `json:"namespaceName" tf:"namespace_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`
}

// EventhubNamespaceAuthorizationRuleSpec defines the desired state of EventhubNamespaceAuthorizationRule
type EventhubNamespaceAuthorizationRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventhubNamespaceAuthorizationRuleParameters `json:"forProvider"`
}

// EventhubNamespaceAuthorizationRuleStatus defines the observed state of EventhubNamespaceAuthorizationRule.
type EventhubNamespaceAuthorizationRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventhubNamespaceAuthorizationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubNamespaceAuthorizationRule is the Schema for the EventhubNamespaceAuthorizationRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EventhubNamespaceAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubNamespaceAuthorizationRuleSpec   `json:"spec"`
	Status            EventhubNamespaceAuthorizationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubNamespaceAuthorizationRuleList contains a list of EventhubNamespaceAuthorizationRules
type EventhubNamespaceAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventhubNamespaceAuthorizationRule `json:"items"`
}

// Repository type metadata.
var (
	EventhubNamespaceAuthorizationRuleKind             = "EventhubNamespaceAuthorizationRule"
	EventhubNamespaceAuthorizationRuleGroupKind        = schema.GroupKind{Group: Group, Kind: EventhubNamespaceAuthorizationRuleKind}.String()
	EventhubNamespaceAuthorizationRuleKindAPIVersion   = EventhubNamespaceAuthorizationRuleKind + "." + GroupVersion.String()
	EventhubNamespaceAuthorizationRuleGroupVersionKind = GroupVersion.WithKind(EventhubNamespaceAuthorizationRuleKind)
)

func init() {
	SchemeBuilder.Register(&EventhubNamespaceAuthorizationRule{}, &EventhubNamespaceAuthorizationRuleList{})
}