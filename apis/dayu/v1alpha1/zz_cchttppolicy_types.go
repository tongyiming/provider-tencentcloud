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

type CcHTTPPolicyObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`
}

type CcHTTPPolicyParameters struct {

	// Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
	// +kubebuilder:validation:Optional
	Frequency *float64 `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// Ip of the CC self-define http policy, only valid when `resource_type` is `bgp-multip`. The num of list items can only be set one.
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// ID of the resource that the CC self-define http policy works for.
	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`

	// Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
	// +kubebuilder:validation:Required
	ResourceType *string `json:"resourceType" tf:"resource_type,omitempty"`

	// Rule list of the CC self-define http policy,  only valid when `smode` is `matching`.
	// +kubebuilder:validation:Optional
	RuleList []RuleListParameters `json:"ruleList,omitempty" tf:"rule_list,omitempty"`

	// Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set one.
	// +kubebuilder:validation:Optional
	Smode *string `json:"smode,omitempty" tf:"smode,omitempty"`

	// Indicate the CC self-define http policy takes effect or not.
	// +kubebuilder:validation:Optional
	Switch *bool `json:"switch,omitempty" tf:"switch,omitempty"`
}

type RuleListObservation struct {
}

type RuleListParameters struct {

	// Operator of the rule. Valid values: `include`, `not_include`, `equal`.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Key of the rule. Valid values: `host`, `cgi`, `ua`, `referer`.
	// +kubebuilder:validation:Optional
	Skey *string `json:"skey,omitempty" tf:"skey,omitempty"`

	// Rule value, then length should be less than 31 bytes.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// CcHTTPPolicySpec defines the desired state of CcHTTPPolicy
type CcHTTPPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CcHTTPPolicyParameters `json:"forProvider"`
}

// CcHTTPPolicyStatus defines the observed state of CcHTTPPolicy.
type CcHTTPPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CcHTTPPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CcHTTPPolicy is the Schema for the CcHTTPPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type CcHTTPPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CcHTTPPolicySpec   `json:"spec"`
	Status            CcHTTPPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CcHTTPPolicyList contains a list of CcHTTPPolicys
type CcHTTPPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CcHTTPPolicy `json:"items"`
}

// Repository type metadata.
var (
	CcHTTPPolicy_Kind             = "CcHTTPPolicy"
	CcHTTPPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CcHTTPPolicy_Kind}.String()
	CcHTTPPolicy_KindAPIVersion   = CcHTTPPolicy_Kind + "." + CRDGroupVersion.String()
	CcHTTPPolicy_GroupVersionKind = CRDGroupVersion.WithKind(CcHTTPPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&CcHTTPPolicy{}, &CcHTTPPolicyList{})
}
