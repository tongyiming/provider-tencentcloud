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

type AnnotationsObservation struct {
}

type AnnotationsParameters struct {

	// key.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// value.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type LabelsObservation struct {
}

type LabelsParameters struct {

	// key.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// value.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type TmpAlertRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TmpAlertRuleParameters struct {

	// Rule alarm duration.
	// +kubebuilder:validation:Optional
	Annotations []AnnotationsParameters `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Rule alarm duration.
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Rule expression, reference documentation: `https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/`.
	// +kubebuilder:validation:Required
	Expr *string `json:"expr" tf:"expr,omitempty"`

	// Instance id.
	// +crossplane:generate:reference:type=TmpInstance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Rule alarm duration.
	// +kubebuilder:validation:Optional
	Labels []LabelsParameters `json:"labels,omitempty" tf:"labels,omitempty"`

	// Alarm notification template id list.
	// +kubebuilder:validation:Required
	Receivers []*string `json:"receivers" tf:"receivers,omitempty"`

	// Rule name.
	// +kubebuilder:validation:Required
	RuleName *string `json:"ruleName" tf:"rule_name,omitempty"`

	// Rule state code.
	// +kubebuilder:validation:Optional
	RuleState *float64 `json:"ruleState,omitempty" tf:"rule_state,omitempty"`

	// Alarm Policy Template Classification.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// TmpAlertRuleSpec defines the desired state of TmpAlertRule
type TmpAlertRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TmpAlertRuleParameters `json:"forProvider"`
}

// TmpAlertRuleStatus defines the observed state of TmpAlertRule.
type TmpAlertRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TmpAlertRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TmpAlertRule is the Schema for the TmpAlertRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type TmpAlertRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TmpAlertRuleSpec   `json:"spec"`
	Status            TmpAlertRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TmpAlertRuleList contains a list of TmpAlertRules
type TmpAlertRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TmpAlertRule `json:"items"`
}

// Repository type metadata.
var (
	TmpAlertRule_Kind             = "TmpAlertRule"
	TmpAlertRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TmpAlertRule_Kind}.String()
	TmpAlertRule_KindAPIVersion   = TmpAlertRule_Kind + "." + CRDGroupVersion.String()
	TmpAlertRule_GroupVersionKind = CRDGroupVersion.WithKind(TmpAlertRule_Kind)
)

func init() {
	SchemeBuilder.Register(&TmpAlertRule{}, &TmpAlertRuleList{})
}