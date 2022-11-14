/*
Copyright 2021 The Crossplane Authors.

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

type TopicObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TopicParameters struct {

	// Whether to enable automatic split. Default value: true.
	// +kubebuilder:validation:Optional
	AutoSplit *bool `json:"autoSplit,omitempty" tf:"auto_split,omitempty"`

	// Logset ID.
	// +crossplane:generate:reference:type=Logset
	// +kubebuilder:validation:Optional
	LogsetID *string `json:"logsetId,omitempty" tf:"logset_id,omitempty"`

	// +kubebuilder:validation:Optional
	LogsetIDRef *v1.Reference `json:"logsetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LogsetIDSelector *v1.Selector `json:"logsetIdSelector,omitempty" tf:"-"`

	// Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.
	// +kubebuilder:validation:Optional
	MaxSplitPartitions *float64 `json:"maxSplitPartitions,omitempty" tf:"max_split_partitions,omitempty"`

	// Number of log topic partitions. Default value: 1. Maximum value: 10.
	// +kubebuilder:validation:Optional
	PartitionCount *float64 `json:"partitionCount,omitempty" tf:"partition_count,omitempty"`

	// Lifecycle in days. Value range: 1~366. Default value: 30.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is passed in, please contact the customer service to add the log topic to the allowlist first..
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Log topic name.
	// +kubebuilder:validation:Required
	TopicName *string `json:"topicName" tf:"topic_name,omitempty"`
}

// TopicSpec defines the desired state of Topic
type TopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicParameters `json:"forProvider"`
}

// TopicStatus defines the observed state of Topic.
type TopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Topic is the Schema for the Topics API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicSpec   `json:"spec"`
	Status            TopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicList contains a list of Topics
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Topic `json:"items"`
}

// Repository type metadata.
var (
	Topic_Kind             = "Topic"
	Topic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Topic_Kind}.String()
	Topic_KindAPIVersion   = Topic_Kind + "." + CRDGroupVersion.String()
	Topic_GroupVersionKind = CRDGroupVersion.WithKind(Topic_Kind)
)

func init() {
	SchemeBuilder.Register(&Topic{}, &TopicList{})
}
