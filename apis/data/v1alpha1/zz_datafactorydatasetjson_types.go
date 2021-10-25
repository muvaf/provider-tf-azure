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

type DataFactoryDatasetJsonAzureBlobStorageLocationObservation struct {
}

type DataFactoryDatasetJsonAzureBlobStorageLocationParameters struct {

	// +kubebuilder:validation:Required
	Container *string `json:"container" tf:"container,omitempty"`

	// +kubebuilder:validation:Optional
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Filename *string `json:"filename" tf:"filename,omitempty"`

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type DataFactoryDatasetJsonHTTPServerLocationObservation struct {
}

type DataFactoryDatasetJsonHTTPServerLocationParameters struct {

	// +kubebuilder:validation:Optional
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Filename *string `json:"filename" tf:"filename,omitempty"`

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// +kubebuilder:validation:Required
	RelativeURL *string `json:"relativeUrl" tf:"relative_url,omitempty"`
}

type DataFactoryDatasetJsonObservation struct {
}

type DataFactoryDatasetJsonParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Optional
	AzureBlobStorageLocation []DataFactoryDatasetJsonAzureBlobStorageLocationParameters `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPServerLocation []DataFactoryDatasetJsonHTTPServerLocationParameters `json:"httpServerLocation,omitempty" tf:"http_server_location,omitempty"`

	// +kubebuilder:validation:Required
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	SchemaColumn []DataFactoryDatasetJsonSchemaColumnParameters `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataFactoryDatasetJsonSchemaColumnObservation struct {
}

type DataFactoryDatasetJsonSchemaColumnParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DataFactoryDatasetJsonSpec defines the desired state of DataFactoryDatasetJson
type DataFactoryDatasetJsonSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataFactoryDatasetJsonParameters `json:"forProvider"`
}

// DataFactoryDatasetJsonStatus defines the observed state of DataFactoryDatasetJson.
type DataFactoryDatasetJsonStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataFactoryDatasetJsonObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryDatasetJson is the Schema for the DataFactoryDatasetJsons API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataFactoryDatasetJson struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryDatasetJsonSpec   `json:"spec"`
	Status            DataFactoryDatasetJsonStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryDatasetJsonList contains a list of DataFactoryDatasetJsons
type DataFactoryDatasetJsonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataFactoryDatasetJson `json:"items"`
}

// Repository type metadata.
var (
	DataFactoryDatasetJsonKind             = "DataFactoryDatasetJson"
	DataFactoryDatasetJsonGroupKind        = schema.GroupKind{Group: Group, Kind: DataFactoryDatasetJsonKind}.String()
	DataFactoryDatasetJsonKindAPIVersion   = DataFactoryDatasetJsonKind + "." + GroupVersion.String()
	DataFactoryDatasetJsonGroupVersionKind = GroupVersion.WithKind(DataFactoryDatasetJsonKind)
)

func init() {
	SchemeBuilder.Register(&DataFactoryDatasetJson{}, &DataFactoryDatasetJsonList{})
}