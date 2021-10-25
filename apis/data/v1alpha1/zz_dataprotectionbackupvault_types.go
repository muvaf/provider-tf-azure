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

type DataProtectionBackupVaultIdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type DataProtectionBackupVaultIdentityParameters struct {

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataProtectionBackupVaultObservation struct {
}

type DataProtectionBackupVaultParameters struct {

	// +kubebuilder:validation:Required
	DatastoreType *string `json:"datastoreType" tf:"datastore_type,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []DataProtectionBackupVaultIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Redundancy *string `json:"redundancy" tf:"redundancy,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DataProtectionBackupVaultSpec defines the desired state of DataProtectionBackupVault
type DataProtectionBackupVaultSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataProtectionBackupVaultParameters `json:"forProvider"`
}

// DataProtectionBackupVaultStatus defines the observed state of DataProtectionBackupVault.
type DataProtectionBackupVaultStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataProtectionBackupVaultObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataProtectionBackupVault is the Schema for the DataProtectionBackupVaults API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataProtectionBackupVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataProtectionBackupVaultSpec   `json:"spec"`
	Status            DataProtectionBackupVaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataProtectionBackupVaultList contains a list of DataProtectionBackupVaults
type DataProtectionBackupVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataProtectionBackupVault `json:"items"`
}

// Repository type metadata.
var (
	DataProtectionBackupVaultKind             = "DataProtectionBackupVault"
	DataProtectionBackupVaultGroupKind        = schema.GroupKind{Group: Group, Kind: DataProtectionBackupVaultKind}.String()
	DataProtectionBackupVaultKindAPIVersion   = DataProtectionBackupVaultKind + "." + GroupVersion.String()
	DataProtectionBackupVaultGroupVersionKind = GroupVersion.WithKind(DataProtectionBackupVaultKind)
)

func init() {
	SchemeBuilder.Register(&DataProtectionBackupVault{}, &DataProtectionBackupVaultList{})
}