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

type ExpressRouteCircuitObservation struct {
	ServiceProviderProvisioningState *string `json:"serviceProviderProvisioningState,omitempty" tf:"service_provider_provisioning_state,omitempty"`
}

type ExpressRouteCircuitParameters struct {

	// +kubebuilder:validation:Optional
	AllowClassicOperations *bool `json:"allowClassicOperations,omitempty" tf:"allow_classic_operations,omitempty"`

	// +kubebuilder:validation:Optional
	BandwidthInGbps *float64 `json:"bandwidthInGbps,omitempty" tf:"bandwidth_in_gbps,omitempty"`

	// +kubebuilder:validation:Optional
	BandwidthInMbps *int64 `json:"bandwidthInMbps,omitempty" tf:"bandwidth_in_mbps,omitempty"`

	// +kubebuilder:validation:Optional
	ExpressRoutePortID *string `json:"expressRoutePortId,omitempty" tf:"express_route_port_id,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PeeringLocation *string `json:"peeringLocation,omitempty" tf:"peering_location,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceProviderName *string `json:"serviceProviderName,omitempty" tf:"service_provider_name,omitempty"`

	// +kubebuilder:validation:Required
	Sku []SkuParameters `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SkuObservation struct {
}

type SkuParameters struct {

	// +kubebuilder:validation:Required
	Family *string `json:"family" tf:"family,omitempty"`

	// +kubebuilder:validation:Required
	Tier *string `json:"tier" tf:"tier,omitempty"`
}

// ExpressRouteCircuitSpec defines the desired state of ExpressRouteCircuit
type ExpressRouteCircuitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpressRouteCircuitParameters `json:"forProvider"`
}

// ExpressRouteCircuitStatus defines the observed state of ExpressRouteCircuit.
type ExpressRouteCircuitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpressRouteCircuitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuit is the Schema for the ExpressRouteCircuits API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ExpressRouteCircuit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitSpec   `json:"spec"`
	Status            ExpressRouteCircuitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitList contains a list of ExpressRouteCircuits
type ExpressRouteCircuitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRouteCircuit `json:"items"`
}

// Repository type metadata.
var (
	ExpressRouteCircuitKind             = "ExpressRouteCircuit"
	ExpressRouteCircuitGroupKind        = schema.GroupKind{Group: Group, Kind: ExpressRouteCircuitKind}.String()
	ExpressRouteCircuitKindAPIVersion   = ExpressRouteCircuitKind + "." + GroupVersion.String()
	ExpressRouteCircuitGroupVersionKind = GroupVersion.WithKind(ExpressRouteCircuitKind)
)

func init() {
	SchemeBuilder.Register(&ExpressRouteCircuit{}, &ExpressRouteCircuitList{})
}