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

type L4RuleV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type L4RuleV2Parameters struct {

	// Business of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
	// +kubebuilder:validation:Required
	Business *string `json:"business" tf:"business,omitempty"`

	// Resource id.
	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`

	// A list of layer 4 rules. Each element contains the following attributes:
	// +kubebuilder:validation:Required
	Rules []RulesParameters `json:"rules" tf:"rules,omitempty"`

	// Resource vpn.
	// +kubebuilder:validation:Required
	VPN *string `json:"vpn" tf:"vpn,omitempty"`

	// The virtual port of the layer 4 rule.
	// +kubebuilder:validation:Required
	VirtualPort *float64 `json:"virtualPort" tf:"virtual_port,omitempty"`
}

type RulesObservation struct {
}

type RulesParameters struct {

	// session hold switch.
	// +kubebuilder:validation:Required
	KeepEnable *bool `json:"keepEnable" tf:"keep_enable,omitempty"`

	// The keeptime of the layer 4 rule.
	// +kubebuilder:validation:Required
	Keeptime *float64 `json:"keeptime" tf:"keeptime,omitempty"`

	// LB type of the rule, `1` for weight cycling and `2` for IP hash.
	// +kubebuilder:validation:Required
	LBType *float64 `json:"lbType" tf:"lb_type,omitempty"`

	// Protocol of the rule.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Corresponding regional information.
	// +kubebuilder:validation:Required
	Region *float64 `json:"region" tf:"region,omitempty"`

	// Remove the watermark state.
	// +kubebuilder:validation:Required
	RemoveSwitch *bool `json:"removeSwitch" tf:"remove_switch,omitempty"`

	// Name of the rule.
	// +kubebuilder:validation:Required
	RuleName *string `json:"ruleName" tf:"rule_name,omitempty"`

	// Source list of the rule.
	// +kubebuilder:validation:Required
	SourceList []RulesSourceListParameters `json:"sourceList" tf:"source_list,omitempty"`

	// The source port of the layer 4 rule.
	// +kubebuilder:validation:Required
	SourcePort *float64 `json:"sourcePort" tf:"source_port,omitempty"`

	// Source type, `1` for source of host, `2` for source of IP.
	// +kubebuilder:validation:Required
	SourceType *float64 `json:"sourceType" tf:"source_type,omitempty"`

	// The virtual port of the layer 4 rule.
	// +kubebuilder:validation:Required
	VirtualPort *float64 `json:"virtualPort" tf:"virtual_port,omitempty"`
}

type RulesSourceListObservation struct {
}

type RulesSourceListParameters struct {

	// Source IP or domain.
	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`

	// Weight of the source.
	// +kubebuilder:validation:Required
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

// L4RuleV2Spec defines the desired state of L4RuleV2
type L4RuleV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     L4RuleV2Parameters `json:"forProvider"`
}

// L4RuleV2Status defines the observed state of L4RuleV2.
type L4RuleV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        L4RuleV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// L4RuleV2 is the Schema for the L4RuleV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type L4RuleV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              L4RuleV2Spec   `json:"spec"`
	Status            L4RuleV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// L4RuleV2List contains a list of L4RuleV2s
type L4RuleV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []L4RuleV2 `json:"items"`
}

// Repository type metadata.
var (
	L4RuleV2_Kind             = "L4RuleV2"
	L4RuleV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: L4RuleV2_Kind}.String()
	L4RuleV2_KindAPIVersion   = L4RuleV2_Kind + "." + CRDGroupVersion.String()
	L4RuleV2_GroupVersionKind = CRDGroupVersion.WithKind(L4RuleV2_Kind)
)

func init() {
	SchemeBuilder.Register(&L4RuleV2{}, &L4RuleV2List{})
}
