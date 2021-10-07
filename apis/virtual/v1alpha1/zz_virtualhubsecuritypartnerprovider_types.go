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

type VirtualHubSecurityPartnerProviderObservation struct {
}

type VirtualHubSecurityPartnerProviderParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	SecurityProviderName *string `json:"securityProviderName" tf:"security_provider_name"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	VirtualHubID *string `json:"virtualHubId,omitempty" tf:"virtual_hub_id"`
}

// VirtualHubSecurityPartnerProviderSpec defines the desired state of VirtualHubSecurityPartnerProvider
type VirtualHubSecurityPartnerProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualHubSecurityPartnerProviderParameters `json:"forProvider"`
}

// VirtualHubSecurityPartnerProviderStatus defines the observed state of VirtualHubSecurityPartnerProvider.
type VirtualHubSecurityPartnerProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualHubSecurityPartnerProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubSecurityPartnerProvider is the Schema for the VirtualHubSecurityPartnerProviders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualHubSecurityPartnerProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualHubSecurityPartnerProviderSpec   `json:"spec"`
	Status            VirtualHubSecurityPartnerProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubSecurityPartnerProviderList contains a list of VirtualHubSecurityPartnerProviders
type VirtualHubSecurityPartnerProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualHubSecurityPartnerProvider `json:"items"`
}

// Repository type metadata.
var (
	VirtualHubSecurityPartnerProviderKind             = "VirtualHubSecurityPartnerProvider"
	VirtualHubSecurityPartnerProviderGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualHubSecurityPartnerProviderKind}.String()
	VirtualHubSecurityPartnerProviderKindAPIVersion   = VirtualHubSecurityPartnerProviderKind + "." + GroupVersion.String()
	VirtualHubSecurityPartnerProviderGroupVersionKind = GroupVersion.WithKind(VirtualHubSecurityPartnerProviderKind)
)

func init() {
	SchemeBuilder.Register(&VirtualHubSecurityPartnerProvider{}, &VirtualHubSecurityPartnerProviderList{})
}
