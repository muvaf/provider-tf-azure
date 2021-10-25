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

type DataFactoryPipelineObservation struct {
}

type DataFactoryPipelineParameters struct {

	// +kubebuilder:validation:Optional
	ActivitiesJSON *string `json:"activitiesJson,omitempty" tf:"activities_json,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Optional
	Concurrency *int64 `json:"concurrency,omitempty" tf:"concurrency,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// +kubebuilder:validation:Optional
	MoniterMetricsAfterDuration *string `json:"moniterMetricsAfterDuration,omitempty" tf:"moniter_metrics_after_duration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`
}

// DataFactoryPipelineSpec defines the desired state of DataFactoryPipeline
type DataFactoryPipelineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataFactoryPipelineParameters `json:"forProvider"`
}

// DataFactoryPipelineStatus defines the observed state of DataFactoryPipeline.
type DataFactoryPipelineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataFactoryPipelineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryPipeline is the Schema for the DataFactoryPipelines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataFactoryPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryPipelineSpec   `json:"spec"`
	Status            DataFactoryPipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryPipelineList contains a list of DataFactoryPipelines
type DataFactoryPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataFactoryPipeline `json:"items"`
}

// Repository type metadata.
var (
	DataFactoryPipelineKind             = "DataFactoryPipeline"
	DataFactoryPipelineGroupKind        = schema.GroupKind{Group: Group, Kind: DataFactoryPipelineKind}.String()
	DataFactoryPipelineKindAPIVersion   = DataFactoryPipelineKind + "." + GroupVersion.String()
	DataFactoryPipelineGroupVersionKind = GroupVersion.WithKind(DataFactoryPipelineKind)
)

func init() {
	SchemeBuilder.Register(&DataFactoryPipeline{}, &DataFactoryPipelineList{})
}