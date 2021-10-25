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

type ConsumptionBudgetSubscriptionFilterObservation struct {
}

type ConsumptionBudgetSubscriptionFilterParameters struct {

	// +kubebuilder:validation:Optional
	Dimension []FilterDimensionParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// +kubebuilder:validation:Optional
	Not []FilterNotParameters `json:"not,omitempty" tf:"not,omitempty"`

	// +kubebuilder:validation:Optional
	Tag []ConsumptionBudgetSubscriptionFilterTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type ConsumptionBudgetSubscriptionFilterTagObservation struct {
}

type ConsumptionBudgetSubscriptionFilterTagParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConsumptionBudgetSubscriptionNotificationObservation struct {
}

type ConsumptionBudgetSubscriptionNotificationParameters struct {

	// +kubebuilder:validation:Optional
	ContactEmails []*string `json:"contactEmails,omitempty" tf:"contact_emails,omitempty"`

	// +kubebuilder:validation:Optional
	ContactGroups []*string `json:"contactGroups,omitempty" tf:"contact_groups,omitempty"`

	// +kubebuilder:validation:Optional
	ContactRoles []*string `json:"contactRoles,omitempty" tf:"contact_roles,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Threshold *int64 `json:"threshold" tf:"threshold,omitempty"`
}

type ConsumptionBudgetSubscriptionObservation struct {
}

type ConsumptionBudgetSubscriptionParameters struct {

	// +kubebuilder:validation:Required
	Amount *float64 `json:"amount" tf:"amount,omitempty"`

	// +kubebuilder:validation:Optional
	Filter []ConsumptionBudgetSubscriptionFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Notification []ConsumptionBudgetSubscriptionNotificationParameters `json:"notification" tf:"notification,omitempty"`

	// +kubebuilder:validation:Required
	SubscriptionID *string `json:"subscriptionId" tf:"subscription_id,omitempty"`

	// +kubebuilder:validation:Optional
	TimeGrain *string `json:"timeGrain,omitempty" tf:"time_grain,omitempty"`

	// +kubebuilder:validation:Required
	TimePeriod []ConsumptionBudgetSubscriptionTimePeriodParameters `json:"timePeriod" tf:"time_period,omitempty"`
}

type ConsumptionBudgetSubscriptionTimePeriodObservation struct {
}

type ConsumptionBudgetSubscriptionTimePeriodParameters struct {

	// +kubebuilder:validation:Optional
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// +kubebuilder:validation:Required
	StartDate *string `json:"startDate" tf:"start_date,omitempty"`
}

type FilterDimensionObservation struct {
}

type FilterDimensionParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type FilterNotDimensionObservation struct {
}

type FilterNotDimensionParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type FilterNotObservation struct {
}

type FilterNotParameters struct {

	// +kubebuilder:validation:Optional
	Dimension []FilterNotDimensionParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// +kubebuilder:validation:Optional
	Tag []NotTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type NotTagObservation struct {
}

type NotTagParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

// ConsumptionBudgetSubscriptionSpec defines the desired state of ConsumptionBudgetSubscription
type ConsumptionBudgetSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConsumptionBudgetSubscriptionParameters `json:"forProvider"`
}

// ConsumptionBudgetSubscriptionStatus defines the observed state of ConsumptionBudgetSubscription.
type ConsumptionBudgetSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConsumptionBudgetSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConsumptionBudgetSubscription is the Schema for the ConsumptionBudgetSubscriptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ConsumptionBudgetSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConsumptionBudgetSubscriptionSpec   `json:"spec"`
	Status            ConsumptionBudgetSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConsumptionBudgetSubscriptionList contains a list of ConsumptionBudgetSubscriptions
type ConsumptionBudgetSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConsumptionBudgetSubscription `json:"items"`
}

// Repository type metadata.
var (
	ConsumptionBudgetSubscriptionKind             = "ConsumptionBudgetSubscription"
	ConsumptionBudgetSubscriptionGroupKind        = schema.GroupKind{Group: Group, Kind: ConsumptionBudgetSubscriptionKind}.String()
	ConsumptionBudgetSubscriptionKindAPIVersion   = ConsumptionBudgetSubscriptionKind + "." + GroupVersion.String()
	ConsumptionBudgetSubscriptionGroupVersionKind = GroupVersion.WithKind(ConsumptionBudgetSubscriptionKind)
)

func init() {
	SchemeBuilder.Register(&ConsumptionBudgetSubscription{}, &ConsumptionBudgetSubscriptionList{})
}