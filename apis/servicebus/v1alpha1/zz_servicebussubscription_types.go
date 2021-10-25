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

type ServicebusSubscriptionObservation struct {
}

type ServicebusSubscriptionParameters struct {

	// +kubebuilder:validation:Optional
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty"`

	// +kubebuilder:validation:Optional
	DeadLetteringOnFilterEvaluationError *bool `json:"deadLetteringOnFilterEvaluationError,omitempty" tf:"dead_lettering_on_filter_evaluation_error,omitempty"`

	// +kubebuilder:validation:Optional
	DeadLetteringOnMessageExpiration *bool `json:"deadLetteringOnMessageExpiration,omitempty" tf:"dead_lettering_on_message_expiration,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultMessageTTL *string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty"`

	// +kubebuilder:validation:Optional
	ForwardDeadLetteredMessagesTo *string `json:"forwardDeadLetteredMessagesTo,omitempty" tf:"forward_dead_lettered_messages_to,omitempty"`

	// +kubebuilder:validation:Optional
	ForwardTo *string `json:"forwardTo,omitempty" tf:"forward_to,omitempty"`

	// +kubebuilder:validation:Optional
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// +kubebuilder:validation:Required
	MaxDeliveryCount *int64 `json:"maxDeliveryCount" tf:"max_delivery_count,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NamespaceName *string `json:"namespaceName" tf:"namespace_name,omitempty"`

	// +kubebuilder:validation:Optional
	RequiresSession *bool `json:"requiresSession,omitempty" tf:"requires_session,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Required
	TopicName *string `json:"topicName" tf:"topic_name,omitempty"`
}

// ServicebusSubscriptionSpec defines the desired state of ServicebusSubscription
type ServicebusSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServicebusSubscriptionParameters `json:"forProvider"`
}

// ServicebusSubscriptionStatus defines the observed state of ServicebusSubscription.
type ServicebusSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServicebusSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicebusSubscription is the Schema for the ServicebusSubscriptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ServicebusSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusSubscriptionSpec   `json:"spec"`
	Status            ServicebusSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicebusSubscriptionList contains a list of ServicebusSubscriptions
type ServicebusSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicebusSubscription `json:"items"`
}

// Repository type metadata.
var (
	ServicebusSubscriptionKind             = "ServicebusSubscription"
	ServicebusSubscriptionGroupKind        = schema.GroupKind{Group: Group, Kind: ServicebusSubscriptionKind}.String()
	ServicebusSubscriptionKindAPIVersion   = ServicebusSubscriptionKind + "." + GroupVersion.String()
	ServicebusSubscriptionGroupVersionKind = GroupVersion.WithKind(ServicebusSubscriptionKind)
)

func init() {
	SchemeBuilder.Register(&ServicebusSubscription{}, &ServicebusSubscriptionList{})
}