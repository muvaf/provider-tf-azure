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

type NatIPConfigurationObservation struct {
}

type NatIPConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Primary *bool `json:"primary" tf:"primary,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIPAddressVersion *string `json:"privateIpAddressVersion,omitempty" tf:"private_ip_address_version,omitempty"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`
}

type PrivateLinkServiceObservation struct {
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`
}

type PrivateLinkServiceParameters struct {

	// +kubebuilder:validation:Optional
	AutoApprovalSubscriptionIds []*string `json:"autoApprovalSubscriptionIds,omitempty" tf:"auto_approval_subscription_ids,omitempty"`

	// +kubebuilder:validation:Optional
	EnableProxyProtocol *bool `json:"enableProxyProtocol,omitempty" tf:"enable_proxy_protocol,omitempty"`

	// +kubebuilder:validation:Required
	LoadBalancerFrontendIPConfigurationIds []*string `json:"loadBalancerFrontendIpConfigurationIds" tf:"load_balancer_frontend_ip_configuration_ids,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NatIPConfiguration []NatIPConfigurationParameters `json:"natIpConfiguration" tf:"nat_ip_configuration,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VisibilitySubscriptionIds []*string `json:"visibilitySubscriptionIds,omitempty" tf:"visibility_subscription_ids,omitempty"`
}

// PrivateLinkServiceSpec defines the desired state of PrivateLinkService
type PrivateLinkServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateLinkServiceParameters `json:"forProvider"`
}

// PrivateLinkServiceStatus defines the observed state of PrivateLinkService.
type PrivateLinkServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateLinkServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateLinkService is the Schema for the PrivateLinkServices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateLinkService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateLinkServiceSpec   `json:"spec"`
	Status            PrivateLinkServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateLinkServiceList contains a list of PrivateLinkServices
type PrivateLinkServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateLinkService `json:"items"`
}

// Repository type metadata.
var (
	PrivateLinkServiceKind             = "PrivateLinkService"
	PrivateLinkServiceGroupKind        = schema.GroupKind{Group: Group, Kind: PrivateLinkServiceKind}.String()
	PrivateLinkServiceKindAPIVersion   = PrivateLinkServiceKind + "." + GroupVersion.String()
	PrivateLinkServiceGroupVersionKind = GroupVersion.WithKind(PrivateLinkServiceKind)
)

func init() {
	SchemeBuilder.Register(&PrivateLinkService{}, &PrivateLinkServiceList{})
}