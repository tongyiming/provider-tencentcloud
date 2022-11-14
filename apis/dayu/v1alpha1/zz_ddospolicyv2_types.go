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

type AclsObservation struct {
}

type AclsParameters struct {

	// Action, optional values: drop, transmit, forward.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// The destination port ends, and the value range is 0~65535.
	// +kubebuilder:validation:Required
	DPortEnd *float64 `json:"dPortEnd" tf:"d_port_end,omitempty"`

	// The destination port starts, and the value range is 0~65535.
	// +kubebuilder:validation:Required
	DPortStart *float64 `json:"dPortStart" tf:"d_port_start,omitempty"`

	// Protocol type, desirable values tcp, udp, all.
	// +kubebuilder:validation:Required
	ForwardProtocol *string `json:"forwardProtocol" tf:"forward_protocol,omitempty"`

	// Policy priority, the lower the number, the higher the level, the higher the rule matches, taking a value of 1-1000.Note: This field may return null, indicating that a valid value could not be retrieved.
	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// The source port ends, and the acceptable value ranges from 0 to 65535.
	// +kubebuilder:validation:Required
	SPortEnd *float64 `json:"sPortEnd" tf:"s_port_end,omitempty"`

	// The source port starts, and the value range is 0~65535.
	// +kubebuilder:validation:Required
	SPortStart *float64 `json:"sPortStart" tf:"s_port_start,omitempty"`
}

type BlackWhiteIpsObservation struct {
}

type BlackWhiteIpsParameters struct {

	// Ip of resource instance.
	// +kubebuilder:validation:Required
	IP *string `json:"ip" tf:"ip,omitempty"`

	// IP type, value [`black`(blacklist IP), `white` (whitelist IP)].
	// +kubebuilder:validation:Required
	IPType *string `json:"ipType" tf:"ip_type,omitempty"`
}

type DdosConnectLimitObservation struct {
}

type DdosConnectLimitParameters struct {

	// Based on connection suppression trigger threshold, value range [0,4294967295].
	// +kubebuilder:validation:Required
	BadConnThreshold *float64 `json:"badConnThreshold" tf:"bad_conn_threshold,omitempty"`

	// Abnormal connection detection condition, connection timeout, value range [0,65535].
	// +kubebuilder:validation:Required
	ConnTimeout *float64 `json:"connTimeout" tf:"conn_timeout,omitempty"`

	// Concurrent connection control based on destination IP+ destination port.
	// +kubebuilder:validation:Required
	DstConnLimit *float64 `json:"dstConnLimit" tf:"dst_conn_limit,omitempty"`

	// Limit on the number of news per second based on the destination IP.
	// +kubebuilder:validation:Required
	DstNewLimit *float64 `json:"dstNewLimit" tf:"dst_new_limit,omitempty"`

	// Abnormal connection detection conditions, empty connection guard switch, value range[0,1].
	// +kubebuilder:validation:Required
	NullConnEnable *float64 `json:"nullConnEnable" tf:"null_conn_enable,omitempty"`

	// Concurrent connection control based on source IP + destination IP.
	// +kubebuilder:validation:Required
	SdConnLimit *float64 `json:"sdConnLimit" tf:"sd_conn_limit,omitempty"`

	// The limit on the number of news per second based on source IP + destination IP.
	// +kubebuilder:validation:Required
	SdNewLimit *float64 `json:"sdNewLimit" tf:"sd_new_limit,omitempty"`

	// Anomaly connection detection condition, syn threshold, value range [0,100].
	// +kubebuilder:validation:Required
	SynLimit *float64 `json:"synLimit" tf:"syn_limit,omitempty"`

	// Anomalous connection detection condition, percentage of syn ack, value range [0,100].
	// +kubebuilder:validation:Required
	SynRate *float64 `json:"synRate" tf:"syn_rate,omitempty"`
}

type DdosGeoIPBlockConfigObservation struct {
}

type DdosGeoIPBlockConfigParameters struct {

	// Block action, take the value [`drop`, `trans`].
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// When the RegionType is customized, the AreaList must be filled in, and a maximum of 128 must be filled in.
	// +kubebuilder:validation:Required
	AreaList []*float64 `json:"areaList" tf:"area_list,omitempty"`

	// Zone type, value [oversea (overseas),china (domestic),customized (custom region)].
	// +kubebuilder:validation:Required
	RegionType *string `json:"regionType" tf:"region_type,omitempty"`
}

type DdosPolicyV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DdosPolicyV2PacketFiltersObservation struct {
}

type DdosPolicyV2PacketFiltersParameters struct {

	// Action, take the value [drop,transmit,drop_black (discard and black out),drop_rst (Interception),drop_black_rst (intercept and block),forward].
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// The end destination port, take the value 1~65535, which must be greater than or equal to the starting destination port.
	// +kubebuilder:validation:Required
	DPortEnd *float64 `json:"dPortEnd" tf:"d_port_end,omitempty"`

	// From the destination port, take the value 0~65535.
	// +kubebuilder:validation:Required
	DPortStart *float64 `json:"dPortStart" tf:"d_port_start,omitempty"`

	// Detection depth from the detection position, value [0,1500].
	// +kubebuilder:validation:Required
	Depth *float64 `json:"depth" tf:"depth,omitempty"`

	// Second detection depth starting from the second detection position, value [0,1500].
	// +kubebuilder:validation:Required
	Depth2 *float64 `json:"depth2" tf:"depth2,omitempty"`

	// Whether to include the detected value, take the value [0 (included),1 (not included)].
	// +kubebuilder:validation:Required
	IsNot *float64 `json:"isNot" tf:"is_not,omitempty"`

	// Whether the second detection contains the detected value, the value [0 (included),1 (not included)].
	// +kubebuilder:validation:Required
	IsNot2 *float64 `json:"isNot2" tf:"is_not2,omitempty"`

	// Detect position, take the value [begin_l3 (IP header),begin_l4 (TCP/UDP header),begin_l5 (T load), no_match (mismatch)].
	// +kubebuilder:validation:Required
	MatchBegin *string `json:"matchBegin" tf:"match_begin,omitempty"`

	// The second detection position. take the value [begin_l3 (IP header),begin_l4 (TCP/UDP header),begin_l5 (T load), no_match (mismatch)].
	// +kubebuilder:validation:Required
	MatchBegin2 *string `json:"matchBegin2" tf:"match_begin2,omitempty"`

	// When there is a second detection condition, the and/or relationship with the first detection condition, takes the value [And (and relationship),none (fill in this value when there is no second detection condition)].
	// +kubebuilder:validation:Required
	MatchLogic *string `json:"matchLogic" tf:"match_logic,omitempty"`

	// Detection type, value [sunday (keyword),pcre (regular expression)].
	// +kubebuilder:validation:Required
	MatchType *string `json:"matchType" tf:"match_type,omitempty"`

	// The second type of detection, takes the value [sunday (keyword),pcre (regular expression)].
	// +kubebuilder:validation:Required
	MatchType2 *string `json:"matchType2" tf:"match_type2,omitempty"`

	// Offset from detection position, value range [0, Depth].
	// +kubebuilder:validation:Required
	Offset *float64 `json:"offset" tf:"offset,omitempty"`

	// Offset from the second detection position, value range [0,Depth2].
	// +kubebuilder:validation:Required
	Offset2 *float64 `json:"offset2" tf:"offset2,omitempty"`

	// The maximum message length, taken from 1 to 1500, must be greater than or equal to the minimum message length.
	// +kubebuilder:validation:Required
	PktlenMax *float64 `json:"pktlenMax" tf:"pktlen_max,omitempty"`

	// Minimum message length, 1-1500.
	// +kubebuilder:validation:Required
	PktlenMin *float64 `json:"pktlenMin" tf:"pktlen_min,omitempty"`

	// Protocol, value [tcp udp icmp all].
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// End source port, take the value 1~65535, must be greater than or equal to the starting source port.
	// +kubebuilder:validation:Required
	SPortEnd *float64 `json:"sPortEnd" tf:"s_port_end,omitempty"`

	// Start the source port, take the value 0~65535.
	// +kubebuilder:validation:Required
	SPortStart *float64 `json:"sPortStart" tf:"s_port_start,omitempty"`

	// Detect values, key strings or regular expressions, take the value [When the detection type is sunday, please fill in the string or hexadecimal bytecode, for example 13233 corresponds to the hexadecimal bytecode of the string `123`;When the detection type is pcre, please fill in the regular expression string;].
	// +kubebuilder:validation:Required
	Str *string `json:"str" tf:"str,omitempty"`

	// The second detection value, the key string or regular expression, takes the value [When the detection type is sunday, please fill in the string or hexadecimal bytecode, for example 13233 corresponds to the hexadecimal bytecode of the string `123`;When the detection type is pcre, please fill in the regular expression string;].
	// +kubebuilder:validation:Required
	Str2 *string `json:"str2" tf:"str2,omitempty"`
}

type DdosPolicyV2Parameters struct {

	// Port ACL policy for DDoS protection.
	// +kubebuilder:validation:Optional
	Acls []AclsParameters `json:"acls,omitempty" tf:"acls,omitempty"`

	// DDoS-protected IP blacklist and whitelist.
	// +kubebuilder:validation:Optional
	BlackWhiteIps []BlackWhiteIpsParameters `json:"blackWhiteIps,omitempty" tf:"black_white_ips,omitempty"`

	// Business of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.
	// +kubebuilder:validation:Optional
	Business *string `json:"business,omitempty" tf:"business,omitempty"`

	// AI protection switch, take the value [`on`, `off`].
	// +kubebuilder:validation:Optional
	DdosAI *string `json:"ddosAi,omitempty" tf:"ddos_ai,omitempty"`

	// DDoS connection suppression options.
	// +kubebuilder:validation:Optional
	DdosConnectLimit []DdosConnectLimitParameters `json:"ddosConnectLimit,omitempty" tf:"ddos_connect_limit,omitempty"`

	// DDoS-protected area block configuration.
	// +kubebuilder:validation:Optional
	DdosGeoIPBlockConfig []DdosGeoIPBlockConfigParameters `json:"ddosGeoIpBlockConfig,omitempty" tf:"ddos_geo_ip_block_config,omitempty"`

	// Protection class, value [`low`, `middle`, `high`].
	// +kubebuilder:validation:Optional
	DdosLevel *string `json:"ddosLevel,omitempty" tf:"ddos_level,omitempty"`

	// Access speed limit configuration for DDoS protection.
	// +kubebuilder:validation:Optional
	DdosSpeedLimitConfig []DdosSpeedLimitConfigParameters `json:"ddosSpeedLimitConfig,omitempty" tf:"ddos_speed_limit_config,omitempty"`

	// DDoS cleaning threshold, value[0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000]; When the value is set to 0, it means that the default value is adopted.
	// +kubebuilder:validation:Optional
	DdosThreshold *float64 `json:"ddosThreshold,omitempty" tf:"ddos_threshold,omitempty"`

	// Feature filtering rules for DDoS protection.
	// +kubebuilder:validation:Optional
	PacketFilters []DdosPolicyV2PacketFiltersParameters `json:"packetFilters,omitempty" tf:"packet_filters,omitempty"`

	// Protocol block configuration for DDoS protection.
	// +kubebuilder:validation:Optional
	ProtocolBlockConfig []ProtocolBlockConfigParameters `json:"protocolBlockConfig,omitempty" tf:"protocol_block_config,omitempty"`

	// The ID of the resource instance.
	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`
}

type DdosSpeedLimitConfigObservation struct {
}

type DdosSpeedLimitConfigParameters struct {

	// Bandwidth bps.
	// +kubebuilder:validation:Required
	Bandwidth *float64 `json:"bandwidth" tf:"bandwidth,omitempty"`

	// List of port ranges, up to 8, multiple; Separated, the range is represented with -; this port range must be filled in; fill in the style 1:0-65535, style 2:80; 443; 1000-2000.
	// +kubebuilder:validation:Required
	DstPortList *string `json:"dstPortList" tf:"dst_port_list,omitempty"`

	// Speed limit mode, take the value [1 (speed limit based on source IP),2 (speed limit based on destination port)].
	// +kubebuilder:validation:Required
	Mode *float64 `json:"mode" tf:"mode,omitempty"`

	// Packet rate pps.
	// +kubebuilder:validation:Required
	PacketRate *float64 `json:"packetRate" tf:"packet_rate,omitempty"`

	// IP protocol numbers, take the value[ ALL (all protocols),TCP (tcp protocol),UDP (udp protocol),SMP (smp protocol),1; 2-100 (custom protocol number range, up to 8)].
	// +kubebuilder:validation:Required
	ProtocolList *string `json:"protocolList" tf:"protocol_list,omitempty"`
}

type ProtocolBlockConfigObservation struct {
}

type ProtocolBlockConfigParameters struct {

	// ICMP block, value [0 (block off), 1 (block on)].
	// +kubebuilder:validation:Required
	DropIcmp *float64 `json:"dropIcmp" tf:"drop_icmp,omitempty"`

	// Other block, value [0 (block off), 1 (block on)].
	// +kubebuilder:validation:Required
	DropOther *float64 `json:"dropOther" tf:"drop_other,omitempty"`

	// TCP block, value [0 (block off), 1 (block on)].
	// +kubebuilder:validation:Required
	DropTCP *float64 `json:"dropTcp" tf:"drop_tcp,omitempty"`

	// UDP block, value [0 (block off), 1 (block on)].
	// +kubebuilder:validation:Required
	DropUDP *float64 `json:"dropUdp" tf:"drop_udp,omitempty"`
}

// DdosPolicyV2Spec defines the desired state of DdosPolicyV2
type DdosPolicyV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DdosPolicyV2Parameters `json:"forProvider"`
}

// DdosPolicyV2Status defines the observed state of DdosPolicyV2.
type DdosPolicyV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DdosPolicyV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DdosPolicyV2 is the Schema for the DdosPolicyV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type DdosPolicyV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DdosPolicyV2Spec   `json:"spec"`
	Status            DdosPolicyV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DdosPolicyV2List contains a list of DdosPolicyV2s
type DdosPolicyV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DdosPolicyV2 `json:"items"`
}

// Repository type metadata.
var (
	DdosPolicyV2_Kind             = "DdosPolicyV2"
	DdosPolicyV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DdosPolicyV2_Kind}.String()
	DdosPolicyV2_KindAPIVersion   = DdosPolicyV2_Kind + "." + CRDGroupVersion.String()
	DdosPolicyV2_GroupVersionKind = CRDGroupVersion.WithKind(DdosPolicyV2_Kind)
)

func init() {
	SchemeBuilder.Register(&DdosPolicyV2{}, &DdosPolicyV2List{})
}
