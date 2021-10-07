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

type PostgresqlConfigurationObservation struct {
}

type PostgresqlConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	ServerName *string `json:"serverName" tf:"server_name"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value"`
}

// PostgresqlConfigurationSpec defines the desired state of PostgresqlConfiguration
type PostgresqlConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PostgresqlConfigurationParameters `json:"forProvider"`
}

// PostgresqlConfigurationStatus defines the observed state of PostgresqlConfiguration.
type PostgresqlConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PostgresqlConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlConfiguration is the Schema for the PostgresqlConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PostgresqlConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlConfigurationSpec   `json:"spec"`
	Status            PostgresqlConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlConfigurationList contains a list of PostgresqlConfigurations
type PostgresqlConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlConfiguration `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlConfigurationKind             = "PostgresqlConfiguration"
	PostgresqlConfigurationGroupKind        = schema.GroupKind{Group: Group, Kind: PostgresqlConfigurationKind}.String()
	PostgresqlConfigurationKindAPIVersion   = PostgresqlConfigurationKind + "." + GroupVersion.String()
	PostgresqlConfigurationGroupVersionKind = GroupVersion.WithKind(PostgresqlConfigurationKind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlConfiguration{}, &PostgresqlConfigurationList{})
}
