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

type DataProtectionBackupPolicyPostgresqlObservation struct {
}

type DataProtectionBackupPolicyPostgresqlParameters struct {

	// +kubebuilder:validation:Required
	BackupRepeatingTimeIntervals []*string `json:"backupRepeatingTimeIntervals" tf:"backup_repeating_time_intervals,omitempty"`

	// +kubebuilder:validation:Required
	DefaultRetentionDuration *string `json:"defaultRetentionDuration" tf:"default_retention_duration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionRule []DataProtectionBackupPolicyPostgresqlRetentionRuleParameters `json:"retentionRule,omitempty" tf:"retention_rule,omitempty"`

	// +kubebuilder:validation:Required
	VaultName *string `json:"vaultName" tf:"vault_name,omitempty"`
}

type DataProtectionBackupPolicyPostgresqlRetentionRuleObservation struct {
}

type DataProtectionBackupPolicyPostgresqlRetentionRuleParameters struct {

	// +kubebuilder:validation:Required
	Criteria []RetentionRuleCriteriaParameters `json:"criteria" tf:"criteria,omitempty"`

	// +kubebuilder:validation:Required
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Priority *int64 `json:"priority" tf:"priority,omitempty"`
}

type RetentionRuleCriteriaObservation struct {
}

type RetentionRuleCriteriaParameters struct {

	// +kubebuilder:validation:Optional
	AbsoluteCriteria *string `json:"absoluteCriteria,omitempty" tf:"absolute_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	DaysOfWeek []*string `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// +kubebuilder:validation:Optional
	MonthsOfYear []*string `json:"monthsOfYear,omitempty" tf:"months_of_year,omitempty"`

	// +kubebuilder:validation:Optional
	ScheduledBackupTimes []*string `json:"scheduledBackupTimes,omitempty" tf:"scheduled_backup_times,omitempty"`

	// +kubebuilder:validation:Optional
	WeeksOfMonth []*string `json:"weeksOfMonth,omitempty" tf:"weeks_of_month,omitempty"`
}

// DataProtectionBackupPolicyPostgresqlSpec defines the desired state of DataProtectionBackupPolicyPostgresql
type DataProtectionBackupPolicyPostgresqlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataProtectionBackupPolicyPostgresqlParameters `json:"forProvider"`
}

// DataProtectionBackupPolicyPostgresqlStatus defines the observed state of DataProtectionBackupPolicyPostgresql.
type DataProtectionBackupPolicyPostgresqlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataProtectionBackupPolicyPostgresqlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataProtectionBackupPolicyPostgresql is the Schema for the DataProtectionBackupPolicyPostgresqls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataProtectionBackupPolicyPostgresql struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataProtectionBackupPolicyPostgresqlSpec   `json:"spec"`
	Status            DataProtectionBackupPolicyPostgresqlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataProtectionBackupPolicyPostgresqlList contains a list of DataProtectionBackupPolicyPostgresqls
type DataProtectionBackupPolicyPostgresqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataProtectionBackupPolicyPostgresql `json:"items"`
}

// Repository type metadata.
var (
	DataProtectionBackupPolicyPostgresqlKind             = "DataProtectionBackupPolicyPostgresql"
	DataProtectionBackupPolicyPostgresqlGroupKind        = schema.GroupKind{Group: Group, Kind: DataProtectionBackupPolicyPostgresqlKind}.String()
	DataProtectionBackupPolicyPostgresqlKindAPIVersion   = DataProtectionBackupPolicyPostgresqlKind + "." + GroupVersion.String()
	DataProtectionBackupPolicyPostgresqlGroupVersionKind = GroupVersion.WithKind(DataProtectionBackupPolicyPostgresqlKind)
)

func init() {
	SchemeBuilder.Register(&DataProtectionBackupPolicyPostgresql{}, &DataProtectionBackupPolicyPostgresqlList{})
}