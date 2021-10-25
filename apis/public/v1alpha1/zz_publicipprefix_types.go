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

type PublicIpPrefixObservation struct {
	IPPrefix *string `json:"ipPrefix,omitempty" tf:"ip_prefix,omitempty"`
}

type PublicIpPrefixParameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PrefixLength *int64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

// PublicIpPrefixSpec defines the desired state of PublicIpPrefix
type PublicIpPrefixSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicIpPrefixParameters `json:"forProvider"`
}

// PublicIpPrefixStatus defines the observed state of PublicIpPrefix.
type PublicIpPrefixStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicIpPrefixObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PublicIpPrefix is the Schema for the PublicIpPrefixs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PublicIpPrefix struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicIpPrefixSpec   `json:"spec"`
	Status            PublicIpPrefixStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicIpPrefixList contains a list of PublicIpPrefixs
type PublicIpPrefixList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicIpPrefix `json:"items"`
}

// Repository type metadata.
var (
	PublicIpPrefixKind             = "PublicIpPrefix"
	PublicIpPrefixGroupKind        = schema.GroupKind{Group: Group, Kind: PublicIpPrefixKind}.String()
	PublicIpPrefixKindAPIVersion   = PublicIpPrefixKind + "." + GroupVersion.String()
	PublicIpPrefixGroupVersionKind = GroupVersion.WithKind(PublicIpPrefixKind)
)

func init() {
	SchemeBuilder.Register(&PublicIpPrefix{}, &PublicIpPrefixList{})
}