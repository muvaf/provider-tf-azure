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

type BatchAccountObservation struct {
	AccountEndpoint *string `json:"accountEndpoint,omitempty" tf:"account_endpoint,omitempty"`
}

type BatchAccountParameters struct {

	// +kubebuilder:validation:Optional
	KeyVaultReference []KeyVaultReferenceParameters `json:"keyVaultReference,omitempty" tf:"key_vault_reference,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PoolAllocationMode *string `json:"poolAllocationMode,omitempty" tf:"pool_allocation_mode,omitempty"`

	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type KeyVaultReferenceObservation struct {
}

type KeyVaultReferenceParameters struct {

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

// BatchAccountSpec defines the desired state of BatchAccount
type BatchAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BatchAccountParameters `json:"forProvider"`
}

// BatchAccountStatus defines the observed state of BatchAccount.
type BatchAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BatchAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BatchAccount is the Schema for the BatchAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BatchAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchAccountSpec   `json:"spec"`
	Status            BatchAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BatchAccountList contains a list of BatchAccounts
type BatchAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BatchAccount `json:"items"`
}

// Repository type metadata.
var (
	BatchAccountKind             = "BatchAccount"
	BatchAccountGroupKind        = schema.GroupKind{Group: Group, Kind: BatchAccountKind}.String()
	BatchAccountKindAPIVersion   = BatchAccountKind + "." + GroupVersion.String()
	BatchAccountGroupVersionKind = GroupVersion.WithKind(BatchAccountKind)
)

func init() {
	SchemeBuilder.Register(&BatchAccount{}, &BatchAccountList{})
}