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

type PropagatedRouteTableObservation struct {
}

type PropagatedRouteTableParameters struct {

	// +kubebuilder:validation:Optional
	Labels []*string `json:"labels,omitempty" tf:"labels"`

	// +kubebuilder:validation:Optional
	RouteTableIds []*string `json:"routeTableIds,omitempty" tf:"route_table_ids"`
}

type RoutingObservation struct {
}

type RoutingParameters struct {

	// +kubebuilder:validation:Optional
	AssociatedRouteTableID *string `json:"associatedRouteTableId,omitempty" tf:"associated_route_table_id"`

	// +kubebuilder:validation:Optional
	PropagatedRouteTable []PropagatedRouteTableParameters `json:"propagatedRouteTable,omitempty" tf:"propagated_route_table"`

	// +kubebuilder:validation:Optional
	StaticVnetRoute []StaticVnetRouteParameters `json:"staticVnetRoute,omitempty" tf:"static_vnet_route"`
}

type StaticVnetRouteObservation struct {
}

type StaticVnetRouteParameters struct {

	// +kubebuilder:validation:Optional
	AddressPrefixes []*string `json:"addressPrefixes,omitempty" tf:"address_prefixes"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	NextHopIPAddress *string `json:"nextHopIpAddress,omitempty" tf:"next_hop_ip_address"`
}

type VirtualHubConnectionObservation struct {
}

type VirtualHubConnectionParameters struct {

	// +kubebuilder:validation:Optional
	HubToVitualNetworkTrafficAllowed *bool `json:"hubToVitualNetworkTrafficAllowed,omitempty" tf:"hub_to_vitual_network_traffic_allowed"`

	// +kubebuilder:validation:Optional
	InternetSecurityEnabled *bool `json:"internetSecurityEnabled,omitempty" tf:"internet_security_enabled"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	RemoteVirtualNetworkID *string `json:"remoteVirtualNetworkId" tf:"remote_virtual_network_id"`

	// +kubebuilder:validation:Optional
	Routing []RoutingParameters `json:"routing,omitempty" tf:"routing"`

	// +kubebuilder:validation:Required
	VirtualHubID *string `json:"virtualHubId" tf:"virtual_hub_id"`

	// +kubebuilder:validation:Optional
	VitualNetworkToHubGatewaysTrafficAllowed *bool `json:"vitualNetworkToHubGatewaysTrafficAllowed,omitempty" tf:"vitual_network_to_hub_gateways_traffic_allowed"`
}

// VirtualHubConnectionSpec defines the desired state of VirtualHubConnection
type VirtualHubConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualHubConnectionParameters `json:"forProvider"`
}

// VirtualHubConnectionStatus defines the observed state of VirtualHubConnection.
type VirtualHubConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualHubConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubConnection is the Schema for the VirtualHubConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualHubConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualHubConnectionSpec   `json:"spec"`
	Status            VirtualHubConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubConnectionList contains a list of VirtualHubConnections
type VirtualHubConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualHubConnection `json:"items"`
}

// Repository type metadata.
var (
	VirtualHubConnectionKind             = "VirtualHubConnection"
	VirtualHubConnectionGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualHubConnectionKind}.String()
	VirtualHubConnectionKindAPIVersion   = VirtualHubConnectionKind + "." + GroupVersion.String()
	VirtualHubConnectionGroupVersionKind = GroupVersion.WithKind(VirtualHubConnectionKind)
)

func init() {
	SchemeBuilder.Register(&VirtualHubConnection{}, &VirtualHubConnectionList{})
}
