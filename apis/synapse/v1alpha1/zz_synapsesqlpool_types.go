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

type RestoreObservation struct {
}

type RestoreParameters struct {

	// +kubebuilder:validation:Required
	PointInTime *string `json:"pointInTime" tf:"point_in_time,omitempty"`

	// +kubebuilder:validation:Required
	SourceDatabaseID *string `json:"sourceDatabaseId" tf:"source_database_id,omitempty"`
}

type SynapseSqlPoolObservation struct {
}

type SynapseSqlPoolParameters struct {

	// +kubebuilder:validation:Optional
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// +kubebuilder:validation:Optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// +kubebuilder:validation:Optional
	DataEncrypted *bool `json:"dataEncrypted,omitempty" tf:"data_encrypted,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RecoveryDatabaseID *string `json:"recoveryDatabaseId,omitempty" tf:"recovery_database_id,omitempty"`

	// +kubebuilder:validation:Optional
	Restore []RestoreParameters `json:"restore,omitempty" tf:"restore,omitempty"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Required
	SynapseWorkspaceID *string `json:"synapseWorkspaceId" tf:"synapse_workspace_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SynapseSqlPoolSpec defines the desired state of SynapseSqlPool
type SynapseSqlPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SynapseSqlPoolParameters `json:"forProvider"`
}

// SynapseSqlPoolStatus defines the observed state of SynapseSqlPool.
type SynapseSqlPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SynapseSqlPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SynapseSqlPool is the Schema for the SynapseSqlPools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SynapseSqlPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SynapseSqlPoolSpec   `json:"spec"`
	Status            SynapseSqlPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SynapseSqlPoolList contains a list of SynapseSqlPools
type SynapseSqlPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SynapseSqlPool `json:"items"`
}

// Repository type metadata.
var (
	SynapseSqlPoolKind             = "SynapseSqlPool"
	SynapseSqlPoolGroupKind        = schema.GroupKind{Group: Group, Kind: SynapseSqlPoolKind}.String()
	SynapseSqlPoolKindAPIVersion   = SynapseSqlPoolKind + "." + GroupVersion.String()
	SynapseSqlPoolGroupVersionKind = GroupVersion.WithKind(SynapseSqlPoolKind)
)

func init() {
	SchemeBuilder.Register(&SynapseSqlPool{}, &SynapseSqlPoolList{})
}