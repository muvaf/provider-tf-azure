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

type PostgresqlFirewallRuleObservation struct {
}

type PostgresqlFirewallRuleParameters struct {

	// +kubebuilder:validation:Required
	EndIPAddress *string `json:"endIpAddress" tf:"end_ip_address"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	ServerName *string `json:"serverName" tf:"server_name"`

	// +kubebuilder:validation:Required
	StartIPAddress *string `json:"startIpAddress" tf:"start_ip_address"`
}

// PostgresqlFirewallRuleSpec defines the desired state of PostgresqlFirewallRule
type PostgresqlFirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PostgresqlFirewallRuleParameters `json:"forProvider"`
}

// PostgresqlFirewallRuleStatus defines the observed state of PostgresqlFirewallRule.
type PostgresqlFirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PostgresqlFirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlFirewallRule is the Schema for the PostgresqlFirewallRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PostgresqlFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlFirewallRuleSpec   `json:"spec"`
	Status            PostgresqlFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlFirewallRuleList contains a list of PostgresqlFirewallRules
type PostgresqlFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlFirewallRule `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlFirewallRuleKind             = "PostgresqlFirewallRule"
	PostgresqlFirewallRuleGroupKind        = schema.GroupKind{Group: Group, Kind: PostgresqlFirewallRuleKind}.String()
	PostgresqlFirewallRuleKindAPIVersion   = PostgresqlFirewallRuleKind + "." + GroupVersion.String()
	PostgresqlFirewallRuleGroupVersionKind = GroupVersion.WithKind(PostgresqlFirewallRuleKind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlFirewallRule{}, &PostgresqlFirewallRuleList{})
}
