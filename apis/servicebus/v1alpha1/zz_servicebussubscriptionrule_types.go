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

type CorrelationFilterObservation struct {
}

type CorrelationFilterParameters struct {

	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// +kubebuilder:validation:Optional
	CorrelationID *string `json:"correlationId,omitempty" tf:"correlation_id,omitempty"`

	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// +kubebuilder:validation:Optional
	MessageID *string `json:"messageId,omitempty" tf:"message_id,omitempty"`

	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// +kubebuilder:validation:Optional
	ReplyTo *string `json:"replyTo,omitempty" tf:"reply_to,omitempty"`

	// +kubebuilder:validation:Optional
	ReplyToSessionID *string `json:"replyToSessionId,omitempty" tf:"reply_to_session_id,omitempty"`

	// +kubebuilder:validation:Optional
	SessionID *string `json:"sessionId,omitempty" tf:"session_id,omitempty"`

	// +kubebuilder:validation:Optional
	To *string `json:"to,omitempty" tf:"to,omitempty"`
}

type ServicebusSubscriptionRuleObservation struct {
}

type ServicebusSubscriptionRuleParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	CorrelationFilter []CorrelationFilterParameters `json:"correlationFilter,omitempty" tf:"correlation_filter,omitempty"`

	// +kubebuilder:validation:Required
	FilterType *string `json:"filterType" tf:"filter_type,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NamespaceName *string `json:"namespaceName" tf:"namespace_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	SQLFilter *string `json:"sqlFilter,omitempty" tf:"sql_filter,omitempty"`

	// +kubebuilder:validation:Required
	SubscriptionName *string `json:"subscriptionName" tf:"subscription_name,omitempty"`

	// +kubebuilder:validation:Required
	TopicName *string `json:"topicName" tf:"topic_name,omitempty"`
}

// ServicebusSubscriptionRuleSpec defines the desired state of ServicebusSubscriptionRule
type ServicebusSubscriptionRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServicebusSubscriptionRuleParameters `json:"forProvider"`
}

// ServicebusSubscriptionRuleStatus defines the observed state of ServicebusSubscriptionRule.
type ServicebusSubscriptionRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServicebusSubscriptionRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicebusSubscriptionRule is the Schema for the ServicebusSubscriptionRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ServicebusSubscriptionRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusSubscriptionRuleSpec   `json:"spec"`
	Status            ServicebusSubscriptionRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicebusSubscriptionRuleList contains a list of ServicebusSubscriptionRules
type ServicebusSubscriptionRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicebusSubscriptionRule `json:"items"`
}

// Repository type metadata.
var (
	ServicebusSubscriptionRuleKind             = "ServicebusSubscriptionRule"
	ServicebusSubscriptionRuleGroupKind        = schema.GroupKind{Group: Group, Kind: ServicebusSubscriptionRuleKind}.String()
	ServicebusSubscriptionRuleKindAPIVersion   = ServicebusSubscriptionRuleKind + "." + GroupVersion.String()
	ServicebusSubscriptionRuleGroupVersionKind = GroupVersion.WithKind(ServicebusSubscriptionRuleKind)
)

func init() {
	SchemeBuilder.Register(&ServicebusSubscriptionRule{}, &ServicebusSubscriptionRuleList{})
}