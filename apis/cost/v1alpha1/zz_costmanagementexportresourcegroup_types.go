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

type CostManagementExportResourceGroupObservation struct {
}

type CostManagementExportResourceGroupParameters struct {

	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// +kubebuilder:validation:Required
	DeliveryInfo []DeliveryInfoParameters `json:"deliveryInfo" tf:"delivery_info,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Query []QueryParameters `json:"query" tf:"query,omitempty"`

	// +kubebuilder:validation:Required
	RecurrencePeriodEnd *string `json:"recurrencePeriodEnd" tf:"recurrence_period_end,omitempty"`

	// +kubebuilder:validation:Required
	RecurrencePeriodStart *string `json:"recurrencePeriodStart" tf:"recurrence_period_start,omitempty"`

	// +kubebuilder:validation:Required
	RecurrenceType *string `json:"recurrenceType" tf:"recurrence_type,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupID *string `json:"resourceGroupId" tf:"resource_group_id,omitempty"`
}

type DeliveryInfoObservation struct {
}

type DeliveryInfoParameters struct {

	// +kubebuilder:validation:Required
	ContainerName *string `json:"containerName" tf:"container_name,omitempty"`

	// +kubebuilder:validation:Required
	RootFolderPath *string `json:"rootFolderPath" tf:"root_folder_path,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`
}

type QueryObservation struct {
}

type QueryParameters struct {

	// +kubebuilder:validation:Required
	TimeFrame *string `json:"timeFrame" tf:"time_frame,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// CostManagementExportResourceGroupSpec defines the desired state of CostManagementExportResourceGroup
type CostManagementExportResourceGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CostManagementExportResourceGroupParameters `json:"forProvider"`
}

// CostManagementExportResourceGroupStatus defines the observed state of CostManagementExportResourceGroup.
type CostManagementExportResourceGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CostManagementExportResourceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CostManagementExportResourceGroup is the Schema for the CostManagementExportResourceGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CostManagementExportResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CostManagementExportResourceGroupSpec   `json:"spec"`
	Status            CostManagementExportResourceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CostManagementExportResourceGroupList contains a list of CostManagementExportResourceGroups
type CostManagementExportResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CostManagementExportResourceGroup `json:"items"`
}

// Repository type metadata.
var (
	CostManagementExportResourceGroupKind             = "CostManagementExportResourceGroup"
	CostManagementExportResourceGroupGroupKind        = schema.GroupKind{Group: Group, Kind: CostManagementExportResourceGroupKind}.String()
	CostManagementExportResourceGroupKindAPIVersion   = CostManagementExportResourceGroupKind + "." + GroupVersion.String()
	CostManagementExportResourceGroupGroupVersionKind = GroupVersion.WithKind(CostManagementExportResourceGroupKind)
)

func init() {
	SchemeBuilder.Register(&CostManagementExportResourceGroup{}, &CostManagementExportResourceGroupList{})
}