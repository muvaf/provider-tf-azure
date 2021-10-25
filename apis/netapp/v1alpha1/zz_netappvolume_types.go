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

type DataProtectionReplicationObservation struct {
}

type DataProtectionReplicationParameters struct {

	// +kubebuilder:validation:Optional
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// +kubebuilder:validation:Required
	RemoteVolumeLocation *string `json:"remoteVolumeLocation" tf:"remote_volume_location,omitempty"`

	// +kubebuilder:validation:Required
	RemoteVolumeResourceID *string `json:"remoteVolumeResourceId" tf:"remote_volume_resource_id,omitempty"`

	// +kubebuilder:validation:Required
	ReplicationFrequency *string `json:"replicationFrequency" tf:"replication_frequency,omitempty"`
}

type ExportPolicyRuleObservation struct {
}

type ExportPolicyRuleParameters struct {

	// +kubebuilder:validation:Required
	AllowedClients []*string `json:"allowedClients" tf:"allowed_clients,omitempty"`

	// +kubebuilder:validation:Optional
	CifsEnabled *bool `json:"cifsEnabled,omitempty" tf:"cifs_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Nfsv3Enabled *bool `json:"nfsv3Enabled,omitempty" tf:"nfsv3_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Nfsv4Enabled *bool `json:"nfsv4Enabled,omitempty" tf:"nfsv4_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ProtocolsEnabled []*string `json:"protocolsEnabled,omitempty" tf:"protocols_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RootAccessEnabled *bool `json:"rootAccessEnabled,omitempty" tf:"root_access_enabled,omitempty"`

	// +kubebuilder:validation:Required
	RuleIndex *int64 `json:"ruleIndex" tf:"rule_index,omitempty"`

	// +kubebuilder:validation:Optional
	UnixReadOnly *bool `json:"unixReadOnly,omitempty" tf:"unix_read_only,omitempty"`

	// +kubebuilder:validation:Optional
	UnixReadWrite *bool `json:"unixReadWrite,omitempty" tf:"unix_read_write,omitempty"`
}

type NetappVolumeObservation struct {
	MountIPAddresses []*string `json:"mountIpAddresses,omitempty" tf:"mount_ip_addresses,omitempty"`
}

type NetappVolumeParameters struct {

	// +kubebuilder:validation:Required
	AccountName *string `json:"accountName" tf:"account_name,omitempty"`

	// +kubebuilder:validation:Optional
	CreateFromSnapshotResourceID *string `json:"createFromSnapshotResourceId,omitempty" tf:"create_from_snapshot_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	DataProtectionReplication []DataProtectionReplicationParameters `json:"dataProtectionReplication,omitempty" tf:"data_protection_replication,omitempty"`

	// +kubebuilder:validation:Optional
	ExportPolicyRule []ExportPolicyRuleParameters `json:"exportPolicyRule,omitempty" tf:"export_policy_rule,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PoolName *string `json:"poolName" tf:"pool_name,omitempty"`

	// +kubebuilder:validation:Optional
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityStyle *string `json:"securityStyle,omitempty" tf:"security_style,omitempty"`

	// +kubebuilder:validation:Required
	ServiceLevel *string `json:"serviceLevel" tf:"service_level,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotDirectoryVisible *bool `json:"snapshotDirectoryVisible,omitempty" tf:"snapshot_directory_visible,omitempty"`

	// +kubebuilder:validation:Required
	StorageQuotaInGb *int64 `json:"storageQuotaInGb" tf:"storage_quota_in_gb,omitempty"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VolumePath *string `json:"volumePath" tf:"volume_path,omitempty"`
}

// NetappVolumeSpec defines the desired state of NetappVolume
type NetappVolumeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetappVolumeParameters `json:"forProvider"`
}

// NetappVolumeStatus defines the observed state of NetappVolume.
type NetappVolumeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetappVolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetappVolume is the Schema for the NetappVolumes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NetappVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetappVolumeSpec   `json:"spec"`
	Status            NetappVolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetappVolumeList contains a list of NetappVolumes
type NetappVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetappVolume `json:"items"`
}

// Repository type metadata.
var (
	NetappVolumeKind             = "NetappVolume"
	NetappVolumeGroupKind        = schema.GroupKind{Group: Group, Kind: NetappVolumeKind}.String()
	NetappVolumeKindAPIVersion   = NetappVolumeKind + "." + GroupVersion.String()
	NetappVolumeGroupVersionKind = GroupVersion.WithKind(NetappVolumeKind)
)

func init() {
	SchemeBuilder.Register(&NetappVolume{}, &NetappVolumeList{})
}