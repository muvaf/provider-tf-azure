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

type NetworkSecurityGroupObservation struct {
}

type NetworkSecurityGroupParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityRule []SecurityRuleParameters `json:"securityRule,omitempty" tf:"security_rule,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SecurityRuleObservation struct {
}

type SecurityRuleParameters struct {

	// +kubebuilder:validation:Required
	Access *string `json:"access" tf:"access,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationAddressPrefix *string `json:"destinationAddressPrefix,omitempty" tf:"destination_address_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationAddressPrefixes []*string `json:"destinationAddressPrefixes,omitempty" tf:"destination_address_prefixes,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationApplicationSecurityGroupIds []*string `json:"destinationApplicationSecurityGroupIds,omitempty" tf:"destination_application_security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationPortRanges []*string `json:"destinationPortRanges,omitempty" tf:"destination_port_ranges,omitempty"`

	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Priority *int64 `json:"priority" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	SourceAddressPrefix *string `json:"sourceAddressPrefix,omitempty" tf:"source_address_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	SourceAddressPrefixes []*string `json:"sourceAddressPrefixes,omitempty" tf:"source_address_prefixes,omitempty"`

	// +kubebuilder:validation:Optional
	SourceApplicationSecurityGroupIds []*string `json:"sourceApplicationSecurityGroupIds,omitempty" tf:"source_application_security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`

	// +kubebuilder:validation:Optional
	SourcePortRanges []*string `json:"sourcePortRanges,omitempty" tf:"source_port_ranges,omitempty"`
}

// NetworkSecurityGroupSpec defines the desired state of NetworkSecurityGroup
type NetworkSecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkSecurityGroupParameters `json:"forProvider"`
}

// NetworkSecurityGroupStatus defines the observed state of NetworkSecurityGroup.
type NetworkSecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkSecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkSecurityGroup is the Schema for the NetworkSecurityGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NetworkSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityGroupSpec   `json:"spec"`
	Status            NetworkSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkSecurityGroupList contains a list of NetworkSecurityGroups
type NetworkSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	NetworkSecurityGroupKind             = "NetworkSecurityGroup"
	NetworkSecurityGroupGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkSecurityGroupKind}.String()
	NetworkSecurityGroupKindAPIVersion   = NetworkSecurityGroupKind + "." + GroupVersion.String()
	NetworkSecurityGroupGroupVersionKind = GroupVersion.WithKind(NetworkSecurityGroupKind)
)

func init() {
	SchemeBuilder.Register(&NetworkSecurityGroup{}, &NetworkSecurityGroupList{})
}