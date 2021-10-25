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

type NatGatewayPublicIpAssociationObservation struct {
}

type NatGatewayPublicIpAssociationParameters struct {

	// +kubebuilder:validation:Required
	NatGatewayID *string `json:"natGatewayId" tf:"nat_gateway_id,omitempty"`

	// +kubebuilder:validation:Required
	PublicIPAddressID *string `json:"publicIpAddressId" tf:"public_ip_address_id,omitempty"`
}

// NatGatewayPublicIpAssociationSpec defines the desired state of NatGatewayPublicIpAssociation
type NatGatewayPublicIpAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NatGatewayPublicIpAssociationParameters `json:"forProvider"`
}

// NatGatewayPublicIpAssociationStatus defines the observed state of NatGatewayPublicIpAssociation.
type NatGatewayPublicIpAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NatGatewayPublicIpAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NatGatewayPublicIpAssociation is the Schema for the NatGatewayPublicIpAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NatGatewayPublicIpAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NatGatewayPublicIpAssociationSpec   `json:"spec"`
	Status            NatGatewayPublicIpAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NatGatewayPublicIpAssociationList contains a list of NatGatewayPublicIpAssociations
type NatGatewayPublicIpAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NatGatewayPublicIpAssociation `json:"items"`
}

// Repository type metadata.
var (
	NatGatewayPublicIpAssociationKind             = "NatGatewayPublicIpAssociation"
	NatGatewayPublicIpAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: NatGatewayPublicIpAssociationKind}.String()
	NatGatewayPublicIpAssociationKindAPIVersion   = NatGatewayPublicIpAssociationKind + "." + GroupVersion.String()
	NatGatewayPublicIpAssociationGroupVersionKind = GroupVersion.WithKind(NatGatewayPublicIpAssociationKind)
)

func init() {
	SchemeBuilder.Register(&NatGatewayPublicIpAssociation{}, &NatGatewayPublicIpAssociationList{})
}