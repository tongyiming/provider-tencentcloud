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

type TmpCvmAgentObservation struct {
	AgentID *string `json:"agentId,omitempty" tf:"agent_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TmpCvmAgentParameters struct {

	// Instance id.
	// +crossplane:generate:reference:type=TmpInstance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Agent name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// TmpCvmAgentSpec defines the desired state of TmpCvmAgent
type TmpCvmAgentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TmpCvmAgentParameters `json:"forProvider"`
}

// TmpCvmAgentStatus defines the observed state of TmpCvmAgent.
type TmpCvmAgentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TmpCvmAgentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TmpCvmAgent is the Schema for the TmpCvmAgents API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type TmpCvmAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TmpCvmAgentSpec   `json:"spec"`
	Status            TmpCvmAgentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TmpCvmAgentList contains a list of TmpCvmAgents
type TmpCvmAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TmpCvmAgent `json:"items"`
}

// Repository type metadata.
var (
	TmpCvmAgent_Kind             = "TmpCvmAgent"
	TmpCvmAgent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TmpCvmAgent_Kind}.String()
	TmpCvmAgent_KindAPIVersion   = TmpCvmAgent_Kind + "." + CRDGroupVersion.String()
	TmpCvmAgent_GroupVersionKind = CRDGroupVersion.WithKind(TmpCvmAgent_Kind)
)

func init() {
	SchemeBuilder.Register(&TmpCvmAgent{}, &TmpCvmAgentList{})
}
