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

type DnatObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DnatParameters struct {

	// Description of the NAT forward.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Network address of the EIP.
	// +kubebuilder:validation:Required
	ElasticIP *string `json:"elasticIp" tf:"elastic_ip,omitempty"`

	// Port of the EIP.
	// +kubebuilder:validation:Required
	ElasticPort *string `json:"elasticPort" tf:"elastic_port,omitempty"`

	// ID of the NAT gateway.
	// +crossplane:generate:reference:type=NatGateway
	// +kubebuilder:validation:Optional
	NATID *string `json:"natId,omitempty" tf:"nat_id,omitempty"`

	// +kubebuilder:validation:Optional
	NATIDRef *v1.Reference `json:"natidRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NATIDSelector *v1.Selector `json:"natidSelector,omitempty" tf:"-"`

	// Network address of the backend service.
	// +kubebuilder:validation:Required
	PrivateIP *string `json:"privateIp" tf:"private_ip,omitempty"`

	// Port of intranet.
	// +kubebuilder:validation:Required
	PrivatePort *string `json:"privatePort" tf:"private_port,omitempty"`

	// Type of the network protocol. Valid value: `TCP` and `UDP`.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// ID of the VPC.
	// +crossplane:generate:reference:type=VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcidRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcidSelector,omitempty" tf:"-"`
}

// DnatSpec defines the desired state of Dnat
type DnatSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DnatParameters `json:"forProvider"`
}

// DnatStatus defines the observed state of Dnat.
type DnatStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DnatObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Dnat is the Schema for the Dnats API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Dnat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnatSpec   `json:"spec"`
	Status            DnatStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DnatList contains a list of Dnats
type DnatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dnat `json:"items"`
}

// Repository type metadata.
var (
	Dnat_Kind             = "Dnat"
	Dnat_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dnat_Kind}.String()
	Dnat_KindAPIVersion   = Dnat_Kind + "." + CRDGroupVersion.String()
	Dnat_GroupVersionKind = CRDGroupVersion.WithKind(Dnat_Kind)
)

func init() {
	SchemeBuilder.Register(&Dnat{}, &DnatList{})
}