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

type DevTestLinuxVirtualMachineObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	UniqueIdentifier *string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier,omitempty"`
}

type DevTestLinuxVirtualMachineParameters struct {

	// +kubebuilder:validation:Optional
	AllowClaim *bool `json:"allowClaim,omitempty" tf:"allow_claim,omitempty"`

	// +kubebuilder:validation:Optional
	DisallowPublicIPAddress *bool `json:"disallowPublicIpAddress,omitempty" tf:"disallow_public_ip_address,omitempty"`

	// +kubebuilder:validation:Required
	GalleryImageReference []GalleryImageReferenceParameters `json:"galleryImageReference" tf:"gallery_image_reference,omitempty"`

	// +kubebuilder:validation:Optional
	InboundNatRule []InboundNatRuleParameters `json:"inboundNatRule,omitempty" tf:"inbound_nat_rule,omitempty"`

	// +kubebuilder:validation:Required
	LabName *string `json:"labName" tf:"lab_name,omitempty"`

	// +kubebuilder:validation:Required
	LabSubnetName *string `json:"labSubnetName" tf:"lab_subnet_name,omitempty"`

	// +kubebuilder:validation:Required
	LabVirtualNetworkID *string `json:"labVirtualNetworkId" tf:"lab_virtual_network_id,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	SSHKey *string `json:"sshKey,omitempty" tf:"ssh_key,omitempty"`

	// +kubebuilder:validation:Required
	Size *string `json:"size" tf:"size,omitempty"`

	// +kubebuilder:validation:Required
	StorageType *string `json:"storageType" tf:"storage_type,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type GalleryImageReferenceObservation struct {
}

type GalleryImageReferenceParameters struct {

	// +kubebuilder:validation:Required
	Offer *string `json:"offer" tf:"offer,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`

	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type InboundNatRuleObservation struct {
	FrontendPort *int64 `json:"frontendPort,omitempty" tf:"frontend_port,omitempty"`
}

type InboundNatRuleParameters struct {

	// +kubebuilder:validation:Required
	BackendPort *int64 `json:"backendPort" tf:"backend_port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

// DevTestLinuxVirtualMachineSpec defines the desired state of DevTestLinuxVirtualMachine
type DevTestLinuxVirtualMachineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DevTestLinuxVirtualMachineParameters `json:"forProvider"`
}

// DevTestLinuxVirtualMachineStatus defines the observed state of DevTestLinuxVirtualMachine.
type DevTestLinuxVirtualMachineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DevTestLinuxVirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DevTestLinuxVirtualMachine is the Schema for the DevTestLinuxVirtualMachines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DevTestLinuxVirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevTestLinuxVirtualMachineSpec   `json:"spec"`
	Status            DevTestLinuxVirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DevTestLinuxVirtualMachineList contains a list of DevTestLinuxVirtualMachines
type DevTestLinuxVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevTestLinuxVirtualMachine `json:"items"`
}

// Repository type metadata.
var (
	DevTestLinuxVirtualMachineKind             = "DevTestLinuxVirtualMachine"
	DevTestLinuxVirtualMachineGroupKind        = schema.GroupKind{Group: Group, Kind: DevTestLinuxVirtualMachineKind}.String()
	DevTestLinuxVirtualMachineKindAPIVersion   = DevTestLinuxVirtualMachineKind + "." + GroupVersion.String()
	DevTestLinuxVirtualMachineGroupVersionKind = GroupVersion.WithKind(DevTestLinuxVirtualMachineKind)
)

func init() {
	SchemeBuilder.Register(&DevTestLinuxVirtualMachine{}, &DevTestLinuxVirtualMachineList{})
}