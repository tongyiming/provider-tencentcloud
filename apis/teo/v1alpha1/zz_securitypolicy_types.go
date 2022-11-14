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

type ACLConfigObservation struct {
	UserRules []UserRulesObservation `json:"userRules,omitempty" tf:"user_rules,omitempty"`
}

type ACLConfigParameters struct {

	// - `on`: Enable.- `off`: Disable.
	// +kubebuilder:validation:Required
	Switch *string `json:"switch" tf:"switch,omitempty"`

	// Custom configuration.
	// +kubebuilder:validation:Optional
	UserRules []UserRulesParameters `json:"userRules,omitempty" tf:"user_rules,omitempty"`
}

type ACLDropPageDetailObservation struct {
}

type ACLDropPageDetailParameters struct {

	// File name or URL.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the custom error page. when set to 0, use system default error page.
	// +kubebuilder:validation:Optional
	PageID *float64 `json:"pageId,omitempty" tf:"page_id,omitempty"`

	// HTTP status code to use. Valid range: 100-600.
	// +kubebuilder:validation:Optional
	StatusCode *float64 `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// Type of the custom error page. Valid values: `file`, `url`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AIRuleObservation struct {
}

type AIRuleParameters struct {

	// Valid values:- `smart_status_close`: disabled.- `smart_status_open`: blocked.- `smart_status_observe`: observed.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type BotConfigObservation struct {
	ManagedRule []ManagedRuleObservation `json:"managedRule,omitempty" tf:"managed_rule,omitempty"`

	PortraitRule []PortraitRuleObservation `json:"portraitRule,omitempty" tf:"portrait_rule,omitempty"`
}

type BotConfigParameters struct {

	// Bot intelligent rule configuration.
	// +kubebuilder:validation:Optional
	IntelligenceRule []IntelligenceRuleParameters `json:"intelligenceRule,omitempty" tf:"intelligence_rule,omitempty"`

	// Preset rules.
	// +kubebuilder:validation:Optional
	ManagedRule []ManagedRuleParameters `json:"managedRule,omitempty" tf:"managed_rule,omitempty"`

	// Portrait rule.
	// +kubebuilder:validation:Optional
	PortraitRule []PortraitRuleParameters `json:"portraitRule,omitempty" tf:"portrait_rule,omitempty"`

	// - `on`: Enable.- `off`: Disable.
	// +kubebuilder:validation:Optional
	Switch *string `json:"switch,omitempty" tf:"switch,omitempty"`
}

type ConditionsObservation struct {
}

type ConditionsParameters struct {

	// Content to match.
	// +kubebuilder:validation:Required
	MatchContent *string `json:"matchContent" tf:"match_content,omitempty"`

	// Items to match. Valid values:- `host`: Host of the request.- `sip`: Client IP.- `ua`: User-Agent.- `cookie`: Session cookie.- `cgi`: CGI script.- `xff`: XFF extension header.- `url`: URL of the request.- `accept`: Accept encoding of the request.- `method`: HTTP method of the request.- `header`: HTTP header of the request.- `sip_proto`: Network protocol of the request.
	// +kubebuilder:validation:Required
	MatchFrom *string `json:"matchFrom" tf:"match_from,omitempty"`

	// Parameter for match item. For example, when match from header, match parameter can be set to a header key.
	// +kubebuilder:validation:Required
	MatchParam *string `json:"matchParam" tf:"match_param,omitempty"`

	// Valid values:- `equal`: string equal.- `not_equal`: string not equal.- `include`: string include.- `not_include`: string not include.- `match`: ip match.- `not_match`: ip not match.- `include_area`: area include.- `is_empty`: field existed but empty.- `not_exists`: field is not existed.- `regexp`: regex match.- `len_gt`: value greater than.- `len_lt`: value less than.- `len_eq`: value equal.- `match_prefix`: string prefix match.- `match_suffix`: string suffix match.- `wildcard`: wildcard match.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`
}

type ConfigObservation struct {
	ACLConfig []ACLConfigObservation `json:"aclConfig,omitempty" tf:"acl_config,omitempty"`

	BotConfig []BotConfigObservation `json:"botConfig,omitempty" tf:"bot_config,omitempty"`

	ExceptConfig []ExceptConfigObservation `json:"exceptConfig,omitempty" tf:"except_config,omitempty"`

	IPTableConfig []IPTableConfigObservation `json:"ipTableConfig,omitempty" tf:"ip_table_config,omitempty"`

	RateLimitConfig []RateLimitConfigObservation `json:"rateLimitConfig,omitempty" tf:"rate_limit_config,omitempty"`
}

type ConfigParameters struct {

	// ACL configuration.
	// +kubebuilder:validation:Optional
	ACLConfig []ACLConfigParameters `json:"aclConfig,omitempty" tf:"acl_config,omitempty"`

	// Bot Configuration.
	// +kubebuilder:validation:Optional
	BotConfig []BotConfigParameters `json:"botConfig,omitempty" tf:"bot_config,omitempty"`

	// Custom drop page configuration.
	// +kubebuilder:validation:Optional
	DropPageConfig []DropPageConfigParameters `json:"dropPageConfig,omitempty" tf:"drop_page_config,omitempty"`

	// Exception rule configuration.
	// +kubebuilder:validation:Optional
	ExceptConfig []ExceptConfigParameters `json:"exceptConfig,omitempty" tf:"except_config,omitempty"`

	// Basic access control.
	// +kubebuilder:validation:Optional
	IPTableConfig []IPTableConfigParameters `json:"ipTableConfig,omitempty" tf:"ip_table_config,omitempty"`

	// RateLimit Configuration.
	// +kubebuilder:validation:Optional
	RateLimitConfig []RateLimitConfigParameters `json:"rateLimitConfig,omitempty" tf:"rate_limit_config,omitempty"`

	// Main switch of 7-layer security.
	// +kubebuilder:validation:Optional
	SwitchConfig []SwitchConfigParameters `json:"switchConfig,omitempty" tf:"switch_config,omitempty"`

	// WAF (Web Application Firewall) Configuration.
	// +kubebuilder:validation:Optional
	WafConfig []WafConfigParameters `json:"wafConfig,omitempty" tf:"waf_config,omitempty"`
}

type DetailObservation struct {
}

type DetailParameters struct {

	// Action to take.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Template ID. Note: This field may return null, indicating that no valid value can be obtained.
	// +kubebuilder:validation:Optional
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// Template Name. Note: This field may return null, indicating that no valid value can be obtained.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Period.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Punish time.
	// +kubebuilder:validation:Optional
	PunishTime *float64 `json:"punishTime,omitempty" tf:"punish_time,omitempty"`

	// Threshold.
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type DropPageConfigObservation struct {
}

type DropPageConfigParameters struct {

	// Custom error page of ACL rules.
	// +kubebuilder:validation:Optional
	ACLDropPageDetail []ACLDropPageDetailParameters `json:"aclDropPageDetail,omitempty" tf:"acl_drop_page_detail,omitempty"`

	// - `on`: Enable.- `off`: Disable.
	// +kubebuilder:validation:Optional
	Switch *string `json:"switch,omitempty" tf:"switch,omitempty"`

	// Custom error page of WAF rules.
	// +kubebuilder:validation:Optional
	WafDropPageDetail []WafDropPageDetailParameters `json:"wafDropPageDetail,omitempty" tf:"waf_drop_page_detail,omitempty"`
}

type ExceptConfigObservation struct {
	ExceptUserRules []ExceptUserRulesObservation `json:"exceptUserRules,omitempty" tf:"except_user_rules,omitempty"`
}

type ExceptConfigParameters struct {

	// Exception rules.
	// +kubebuilder:validation:Optional
	ExceptUserRules []ExceptUserRulesParameters `json:"exceptUserRules,omitempty" tf:"except_user_rules,omitempty"`

	// - `on`: Enable.- `off`: Disable.
	// +kubebuilder:validation:Optional
	Switch *string `json:"switch,omitempty" tf:"switch,omitempty"`
}

type ExceptUserRuleConditionsObservation struct {
}

type ExceptUserRuleConditionsParameters struct {

	// Content to match.
	// +kubebuilder:validation:Optional
	MatchContent *string `json:"matchContent,omitempty" tf:"match_content,omitempty"`

	// Items to match. Valid values:- `host`: Host of the request.- `sip`: Client IP.- `ua`: User-Agent.- `cookie`: Session cookie.- `cgi`: CGI script.- `xff`: XFF extension header.- `url`: URL of the request.- `accept`: Accept encoding of the request.- `method`: HTTP method of the request.- `header`: HTTP header of the request.- `sip_proto`: Network protocol of the request.
	// +kubebuilder:validation:Optional
	MatchFrom *string `json:"matchFrom,omitempty" tf:"match_from,omitempty"`

	// Parameter for match item. For example, when match from header, match parameter can be set to a header key.
	// +kubebuilder:validation:Optional
	MatchParam *string `json:"matchParam,omitempty" tf:"match_param,omitempty"`

	// Valid values:- `equal`: string equal.- `not_equal`: string not equal.- `include`: string include.- `not_include`: string not include.- `match`: ip match.- `not_match`: ip not match.- `include_area`: area include.- `is_empty`: field existed but empty.- `not_exists`: field is not existed.- `regexp`: regex match.- `len_gt`: value greater than.- `len_lt`: value less than.- `len_eq`: value equal.- `match_prefix`: string prefix match.- `match_suffix`: string suffix match.- `wildcard`: wildcard match.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
}

type ExceptUserRuleScopeObservation struct {
}

type ExceptUserRuleScopeParameters struct {

	// Modules in which the rule take effect. Valid values: `waf`.
	// +kubebuilder:validation:Optional
	Modules []*string `json:"modules,omitempty" tf:"modules,omitempty"`
}

type ExceptUserRulesObservation struct {
	RuleID *float64 `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	RuleName *string `json:"ruleName,omitempty" tf:"rule_name,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type ExceptUserRulesParameters struct {

	// Action to take. Valid values: `skip`.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Conditions of the rule.
	// +kubebuilder:validation:Optional
	ExceptUserRuleConditions []ExceptUserRuleConditionsParameters `json:"exceptUserRuleConditions,omitempty" tf:"except_user_rule_conditions,omitempty"`

	// Scope of the rule in effect.
	// +kubebuilder:validation:Optional
	ExceptUserRuleScope []ExceptUserRuleScopeParameters `json:"exceptUserRuleScope,omitempty" tf:"except_user_rule_scope,omitempty"`

	// Priority of the rule. Valid value range: 0-100.
	// +kubebuilder:validation:Optional
	RulePriority *float64 `json:"rulePriority,omitempty" tf:"rule_priority,omitempty"`

	// Status of the rule. Valid values:- `on`: Enabled.- `off`: Disabled.
	// +kubebuilder:validation:Optional
	RuleStatus *string `json:"ruleStatus,omitempty" tf:"rule_status,omitempty"`
}

type IPTableConfigObservation struct {
	Rules []IPTableConfigRulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`
}

type IPTableConfigParameters struct {

	// Rules list.
	// +kubebuilder:validation:Optional
	Rules []IPTableConfigRulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// - `on`: Enable.- `off`: Disable.
	// +kubebuilder:validation:Optional
	Switch *string `json:"switch,omitempty" tf:"switch,omitempty"`
}

type IPTableConfigRulesObservation struct {
	RuleID *float64 `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type IPTableConfigRulesParameters struct {

	// Actions to take. Valid values: `drop`, `trans`, `monitor`.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Matching content.
	// +kubebuilder:validation:Optional
	MatchContent *string `json:"matchContent,omitempty" tf:"match_content,omitempty"`

	// Matching type. Valid values: `ip`, `area`.
	// +kubebuilder:validation:Optional
	MatchFrom *string `json:"matchFrom,omitempty" tf:"match_from,omitempty"`
}

type IntelligenceObservation struct {
}

type IntelligenceParameters struct {

	// Action to take. Valid values: `monitor`, `alg`.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// - `on`: Enable.- `off`: Disable.
	// +kubebuilder:validation:Optional
	Switch *string `json:"switch,omitempty" tf:"switch,omitempty"`
}

type IntelligenceRuleObservation struct {
}

type IntelligenceRuleParameters struct {

	// Configuration detail.
	// +kubebuilder:validation:Optional
	Items []ItemsParameters `json:"items,omitempty" tf:"items,omitempty"`

	// - `on`: Enable.- `off`: Disable.
	// +kubebuilder:validation:Optional
	Switch *string `json:"switch,omitempty" tf:"switch,omitempty"`
}

type ItemsObservation struct {
}

type ItemsParameters struct {

	// Action to take. Valid values: `trans`, `monitor`, `alg`, `captcha`, `drop`.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Bot label, valid values: `evil_bot`, `suspect_bot`, `good_bot`, `normal`.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`
}

type ManagedRuleObservation struct {
	RuleID *float64 `json:"ruleId,omitempty" tf:"rule_id,omitempty"`
}

type ManagedRuleParameters struct {

	// Action to take. Valid values: `drop`, `trans`, `monitor`, `alg`.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Rules to enable when action is `alg`. See details in data source `bot_managed_rules`.
	// +kubebuilder:validation:Optional
	AlgManagedIds []*float64 `json:"algManagedIds,omitempty" tf:"alg_managed_ids,omitempty"`

	// Rules to enable when action is `captcha`. See details in data source `bot_managed_rules`.
	// +kubebuilder:validation:Optional
	CapManagedIds []*float64 `json:"capManagedIds,omitempty" tf:"cap_managed_ids,omitempty"`

	// Rules to enable when action is `drop`. See details in data source `bot_managed_rules`.
	// +kubebuilder:validation:Optional
	DropManagedIds []*float64 `json:"dropManagedIds,omitempty" tf:"drop_managed_ids,omitempty"`

	// Rules to enable when action is `monitor`. See details in data source `bot_managed_rules`.
	// +kubebuilder:validation:Optional
	MonManagedIds []*float64 `json:"monManagedIds,omitempty" tf:"mon_managed_ids,omitempty"`

	// Name of the custom response page.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the custom response page.
	// +kubebuilder:validation:Optional
	PageID *float64 `json:"pageId,omitempty" tf:"page_id,omitempty"`

	// Punish time.
	// +kubebuilder:validation:Optional
	PunishTime *float64 `json:"punishTime,omitempty" tf:"punish_time,omitempty"`

	// Time unit of the punish time.
	// +kubebuilder:validation:Optional
	PunishTimeUnit *string `json:"punishTimeUnit,omitempty" tf:"punish_time_unit,omitempty"`

	// Redirect target URL, must be an sub-domain from one of the account&#39;s site.
	// +kubebuilder:validation:Optional
	RedirectURL *string `json:"redirectUrl,omitempty" tf:"redirect_url,omitempty"`

	// Response code to use when redirecting.
	// +kubebuilder:validation:Optional
	ResponseCode *float64 `json:"responseCode,omitempty" tf:"response_code,omitempty"`

	// Rules to enable when action is `trans`. See details in data source `bot_managed_rules`.
	// +kubebuilder:validation:Optional
	TransManagedIds []*float64 `json:"transManagedIds,omitempty" tf:"trans_managed_ids,omitempty"`
}

type PortraitRuleObservation struct {
	RuleID *float64 `json:"ruleId,omitempty" tf:"rule_id,omitempty"`
}

type PortraitRuleParameters struct {

	// Rules to enable when action is `alg`. See details in data source `bot_portrait_rules`.
	// +kubebuilder:validation:Optional
	AlgManagedIds []*float64 `json:"algManagedIds,omitempty" tf:"alg_managed_ids,omitempty"`

	// Rules to enable when action is `captcha`. See details in data source `bot_portrait_rules`.
	// +kubebuilder:validation:Optional
	CapManagedIds []*float64 `json:"capManagedIds,omitempty" tf:"cap_managed_ids,omitempty"`

	// Rules to enable when action is `drop`. See details in data source `bot_portrait_rules`.
	// +kubebuilder:validation:Optional
	DropManagedIds []*float64 `json:"dropManagedIds,omitempty" tf:"drop_managed_ids,omitempty"`

	// Rules to enable when action is `monitor`. See details in data source `bot_portrait_rules`.
	// +kubebuilder:validation:Optional
	MonManagedIds []*float64 `json:"monManagedIds,omitempty" tf:"mon_managed_ids,omitempty"`

	// - `on`: Enable.- `off`: Disable.
	// +kubebuilder:validation:Optional
	Switch *string `json:"switch,omitempty" tf:"switch,omitempty"`
}

type RateLimitConfigObservation struct {
	UserRules []RateLimitConfigUserRulesObservation `json:"userRules,omitempty" tf:"user_rules,omitempty"`
}

type RateLimitConfigParameters struct {

	// Intelligent client filter.
	// +kubebuilder:validation:Optional
	Intelligence []IntelligenceParameters `json:"intelligence,omitempty" tf:"intelligence,omitempty"`

	// - `on`: Enable.- `off`: Disable.
	// +kubebuilder:validation:Optional
	Switch *string `json:"switch,omitempty" tf:"switch,omitempty"`

	// Default Template. Note: This field may return null, indicating that no valid value can be obtained.
	// +kubebuilder:validation:Optional
	Template []TemplateParameters `json:"template,omitempty" tf:"template,omitempty"`

	// Custom configuration.
	// +kubebuilder:validation:Optional
	UserRules []RateLimitConfigUserRulesParameters `json:"userRules,omitempty" tf:"user_rules,omitempty"`
}

type RateLimitConfigUserRulesObservation struct {
	RuleID *float64 `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type RateLimitConfigUserRulesParameters struct {

	// Valid values: `monitor`, `drop`.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Conditions of the rule.
	// +kubebuilder:validation:Required
	Conditions []UserRulesConditionsParameters `json:"conditions" tf:"conditions,omitempty"`

	// Filter words.
	// +kubebuilder:validation:Optional
	FreqFields []*string `json:"freqFields,omitempty" tf:"freq_fields,omitempty"`

	// Period of the rate limit. Valid values: 10, 20, 30, 40, 50, 60 (in seconds).
	// +kubebuilder:validation:Required
	Period *float64 `json:"period" tf:"period,omitempty"`

	// Punish time, Valid value range: 0-2 days.
	// +kubebuilder:validation:Required
	PunishTime *float64 `json:"punishTime" tf:"punish_time,omitempty"`

	// Time unit of the punish time. Valid values: `second`, `minutes`, `hour`.
	// +kubebuilder:validation:Required
	PunishTimeUnit *string `json:"punishTimeUnit" tf:"punish_time_unit,omitempty"`

	// Rule Name.
	// +kubebuilder:validation:Required
	RuleName *string `json:"ruleName" tf:"rule_name,omitempty"`

	// Priority of the rule. Valid value range: 1-100.
	// +kubebuilder:validation:Required
	RulePriority *float64 `json:"rulePriority" tf:"rule_priority,omitempty"`

	// Status of the rule. Valid values: `on`, `off`, `hour`.
	// +kubebuilder:validation:Optional
	RuleStatus *string `json:"ruleStatus,omitempty" tf:"rule_status,omitempty"`

	// Threshold of the rate limit. Valid value range: 0-4294967294.
	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`
}

type SecurityPolicyObservation struct {
	Config []ConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecurityPolicyParameters struct {

	// Security policy configuration.
	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// Subdomain.
	// +kubebuilder:validation:Required
	Entity *string `json:"entity" tf:"entity,omitempty"`

	// Site ID.
	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

type SwitchConfigObservation struct {
}

type SwitchConfigParameters struct {

	// - `on`: Enable.- `off`: Disable.
	// +kubebuilder:validation:Optional
	WebSwitch *string `json:"webSwitch,omitempty" tf:"web_switch,omitempty"`
}

type TemplateObservation struct {
}

type TemplateParameters struct {

	// Detail of the template.
	// +kubebuilder:validation:Optional
	Detail []DetailParameters `json:"detail,omitempty" tf:"detail,omitempty"`

	// Template Name. Note: This field may return null, indicating that no valid value can be obtained.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type UserRulesConditionsObservation struct {
}

type UserRulesConditionsParameters struct {

	// Content to match.
	// +kubebuilder:validation:Required
	MatchContent *string `json:"matchContent" tf:"match_content,omitempty"`

	// Items to match. Valid values:- `host`: Host of the request.- `sip`: Client IP.- `ua`: User-Agent.- `cookie`: Session cookie.- `cgi`: CGI script.- `xff`: XFF extension header.- `url`: URL of the request.- `accept`: Accept encoding of the request.- `method`: HTTP method of the request.- `header`: HTTP header of the request.- `sip_proto`: Network protocol of the request.
	// +kubebuilder:validation:Required
	MatchFrom *string `json:"matchFrom" tf:"match_from,omitempty"`

	// Parameter for match item. For example, when match from header, match parameter can be set to a header key.
	// +kubebuilder:validation:Required
	MatchParam *string `json:"matchParam" tf:"match_param,omitempty"`

	// Valid values:- `equal`: string equal.- `not_equal`: string not equal.- `include`: string include.- `not_include`: string not include.- `match`: ip match.- `not_match`: ip not match.- `include_area`: area include.- `is_empty`: field existed but empty.- `not_exists`: field is not existed.- `regexp`: regex match.- `len_gt`: value greater than.- `len_lt`: value less than.- `len_eq`: value equal.- `match_prefix`: string prefix match.- `match_suffix`: string suffix match.- `wildcard`: wildcard match.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`
}

type UserRulesObservation struct {
	RuleID *float64 `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type UserRulesParameters struct {

	// Action to take. Valid values: `trans`, `drop`, `monitor`, `ban`, `redirect`, `page`, `alg`.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Conditions of the rule.
	// +kubebuilder:validation:Required
	Conditions []ConditionsParameters `json:"conditions" tf:"conditions,omitempty"`

	// Name of the custom response page.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the custom response page.
	// +kubebuilder:validation:Optional
	PageID *float64 `json:"pageId,omitempty" tf:"page_id,omitempty"`

	// Punish time, Valid value range: 0-2 days.
	// +kubebuilder:validation:Optional
	PunishTime *float64 `json:"punishTime,omitempty" tf:"punish_time,omitempty"`

	// Time unit of the punish time. Valid values: `second`, `minutes`, `hour`.
	// +kubebuilder:validation:Optional
	PunishTimeUnit *string `json:"punishTimeUnit,omitempty" tf:"punish_time_unit,omitempty"`

	// Redirect target URL, must be an sub-domain from one of the account&#39;s site.
	// +kubebuilder:validation:Optional
	RedirectURL *string `json:"redirectUrl,omitempty" tf:"redirect_url,omitempty"`

	// Response code to use when redirecting.
	// +kubebuilder:validation:Optional
	ResponseCode *float64 `json:"responseCode,omitempty" tf:"response_code,omitempty"`

	// Rule name.
	// +kubebuilder:validation:Required
	RuleName *string `json:"ruleName" tf:"rule_name,omitempty"`

	// Priority of the rule. Valid value range: 0-100.
	// +kubebuilder:validation:Required
	RulePriority *float64 `json:"rulePriority" tf:"rule_priority,omitempty"`

	// Status of the rule. Valid values: `on`, `off`.
	// +kubebuilder:validation:Required
	RuleStatus *string `json:"ruleStatus" tf:"rule_status,omitempty"`
}

type WafConfigObservation struct {
}

type WafConfigParameters struct {

	// AI based rules configuration.
	// +kubebuilder:validation:Optional
	AIRule []AIRuleParameters `json:"aiRule,omitempty" tf:"ai_rule,omitempty"`

	// Protection level. Valid values: `loose`, `normal`, `strict`, `stricter`, `custom`.
	// +kubebuilder:validation:Required
	Level *string `json:"level" tf:"level,omitempty"`

	// Protection mode. Valid values:- `block`: use block mode globally, you still can set a group of rules to use observe mode.- `observe`: use observe mode globally.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// Whether to enable WAF rules. Valid values:- `on`: Enable.- `off`: Disable.
	// +kubebuilder:validation:Required
	Switch *string `json:"switch" tf:"switch,omitempty"`

	// WAF Rules Configuration.
	// +kubebuilder:validation:Required
	WafRules []WafRulesParameters `json:"wafRules" tf:"waf_rules,omitempty"`
}

type WafDropPageDetailObservation struct {
}

type WafDropPageDetailParameters struct {

	// File name or URL.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the custom error page. when set to 0, use system default error page.
	// +kubebuilder:validation:Optional
	PageID *float64 `json:"pageId,omitempty" tf:"page_id,omitempty"`

	// HTTP status code to use. Valid range: 100-600.
	// +kubebuilder:validation:Optional
	StatusCode *float64 `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// Type of the custom error page. Valid values: `file`, `url`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type WafRulesObservation struct {
}

type WafRulesParameters struct {

	// Block mode rules list. See details in data source `waf_managed_rules`.
	// +kubebuilder:validation:Required
	BlockRuleIds []*float64 `json:"blockRuleIds" tf:"block_rule_ids,omitempty"`

	// Observe rules list. See details in data source `waf_managed_rules`.
	// +kubebuilder:validation:Optional
	ObserveRuleIds []*float64 `json:"observeRuleIds,omitempty" tf:"observe_rule_ids,omitempty"`

	// Whether to host the rules&#39; configuration.- `on`: Enable.- `off`: Disable.
	// +kubebuilder:validation:Required
	Switch *string `json:"switch" tf:"switch,omitempty"`
}

// SecurityPolicySpec defines the desired state of SecurityPolicy
type SecurityPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityPolicyParameters `json:"forProvider"`
}

// SecurityPolicyStatus defines the observed state of SecurityPolicy.
type SecurityPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityPolicy is the Schema for the SecurityPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type SecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityPolicySpec   `json:"spec"`
	Status            SecurityPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityPolicyList contains a list of SecurityPolicys
type SecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityPolicy `json:"items"`
}

// Repository type metadata.
var (
	SecurityPolicy_Kind             = "SecurityPolicy"
	SecurityPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityPolicy_Kind}.String()
	SecurityPolicy_KindAPIVersion   = SecurityPolicy_Kind + "." + CRDGroupVersion.String()
	SecurityPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SecurityPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityPolicy{}, &SecurityPolicyList{})
}
