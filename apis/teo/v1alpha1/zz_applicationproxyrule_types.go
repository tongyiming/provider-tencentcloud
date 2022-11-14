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

type ApplicationProxyRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`
}

type ApplicationProxyRuleParameters struct {

	// Passes the client IP. Default value is OFF.When Proto is TCP, valid values:- `TOA`: Pass the client IP via TOA.- `PPV1`: Pass the client IP via Proxy Protocol V1.- `PPV2`: Pass the client IP via Proxy Protocol V2.- `OFF`: Do not pass the client IP.When Proto=UDP, valid values:- `PPV2`: Pass the client IP via Proxy Protocol V2.- `OFF`: Do not pass the client IP.
	// +kubebuilder:validation:Optional
	ForwardClientIP *string `json:"forwardClientIp,omitempty" tf:"forward_client_ip,omitempty"`

	// Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.
	// +kubebuilder:validation:Required
	OriginPort *string `json:"originPort" tf:"origin_port,omitempty"`

	// Origin server type.- `custom`: Specified origins.- `origins`: An origin group.
	// +kubebuilder:validation:Required
	OriginType *string `json:"originType" tf:"origin_type,omitempty"`

	// Origin server information.When `OriginType` is custom, this field value indicates multiple origin servers in either of the following formats:- `IP`:Port- Domain name:Port.When `OriginType` is origins, it indicates the origin group ID.
	// +kubebuilder:validation:Required
	OriginValue []*string `json:"originValue" tf:"origin_value,omitempty"`

	// Valid values:- port number: `80` means port 80.- port range: `81-90` means port range 81-90.
	// +kubebuilder:validation:Required
	Port []*string `json:"port" tf:"port,omitempty"`

	// Protocol. Valid values: `TCP`, `UDP`.
	// +kubebuilder:validation:Required
	Proto *string `json:"proto" tf:"proto,omitempty"`

	// Proxy ID.
	// +kubebuilder:validation:Required
	ProxyID *string `json:"proxyId" tf:"proxy_id,omitempty"`

	// Specifies whether to enable session persistence. Default value is false.
	// +kubebuilder:validation:Optional
	SessionPersist *bool `json:"sessionPersist,omitempty" tf:"session_persist,omitempty"`

	// Status of this application proxy rule. Valid values to set is `online` and `offline`.- `online`: Enable.- `offline`: Disable.- `progress`: Deploying.- `stopping`: Disabling.- `fail`: Deployment/Disabling failed.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Site ID.
	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// ApplicationProxyRuleSpec defines the desired state of ApplicationProxyRule
type ApplicationProxyRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationProxyRuleParameters `json:"forProvider"`
}

// ApplicationProxyRuleStatus defines the observed state of ApplicationProxyRule.
type ApplicationProxyRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationProxyRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationProxyRule is the Schema for the ApplicationProxyRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type ApplicationProxyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationProxyRuleSpec   `json:"spec"`
	Status            ApplicationProxyRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationProxyRuleList contains a list of ApplicationProxyRules
type ApplicationProxyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationProxyRule `json:"items"`
}

// Repository type metadata.
var (
	ApplicationProxyRule_Kind             = "ApplicationProxyRule"
	ApplicationProxyRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationProxyRule_Kind}.String()
	ApplicationProxyRule_KindAPIVersion   = ApplicationProxyRule_Kind + "." + CRDGroupVersion.String()
	ApplicationProxyRule_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationProxyRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationProxyRule{}, &ApplicationProxyRuleList{})
}
