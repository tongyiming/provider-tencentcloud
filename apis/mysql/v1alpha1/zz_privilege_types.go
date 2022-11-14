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

type ColumnObservation struct {
}

type ColumnParameters struct {

	// Column name.
	// +kubebuilder:validation:Required
	ColumnName *string `json:"columnName" tf:"column_name,omitempty"`

	// Database name.
	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// Column privilege.available values for Privileges:SELECT,INSERT,UPDATE,REFERENCES.
	// +kubebuilder:validation:Required
	Privileges []*string `json:"privileges" tf:"privileges,omitempty"`

	// Table name.
	// +kubebuilder:validation:Required
	TableName *string `json:"tableName" tf:"table_name,omitempty"`
}

type DatabaseObservation struct {
}

type DatabaseParameters struct {

	// Database name.
	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// Database privilege.available values for Privileges:SELECT,INSERT,UPDATE,DELETE,CREATE,DROP,REFERENCES,INDEX,ALTER,CREATE TEMPORARY TABLES,LOCK TABLES,EXECUTE,CREATE VIEW,SHOW VIEW,CREATE ROUTINE,ALTER ROUTINE,EVENT,TRIGGER.
	// +kubebuilder:validation:Required
	Privileges []*string `json:"privileges" tf:"privileges,omitempty"`
}

type PrivilegeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivilegeParameters struct {

	// Account host, default is `%`.
	// +kubebuilder:validation:Optional
	AccountHost *string `json:"accountHost,omitempty" tf:"account_host,omitempty"`

	// Account name.the forbidden value is:root,mysql.sys,tencentroot.
	// +kubebuilder:validation:Required
	AccountName *string `json:"accountName" tf:"account_name,omitempty"`

	// Column privileges list.
	// +kubebuilder:validation:Optional
	Column []ColumnParameters `json:"column,omitempty" tf:"column,omitempty"`

	// Database privileges list.
	// +kubebuilder:validation:Optional
	Database []DatabaseParameters `json:"database,omitempty" tf:"database,omitempty"`

	// Global privileges. available values for Privileges:ALTER,ALTER ROUTINE,CREATE,CREATE ROUTINE,CREATE TEMPORARY TABLES,CREATE USER,CREATE VIEW,DELETE,DROP,EVENT,EXECUTE,INDEX,INSERT,LOCK TABLES,PROCESS,REFERENCES,RELOAD,REPLICATION CLIENT,REPLICATION SLAVE,SELECT,SHOW DATABASES,SHOW VIEW,TRIGGER,UPDATE.
	// +kubebuilder:validation:Required
	Global []*string `json:"global" tf:"global,omitempty"`

	// Instance ID.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	MySQLID *string `json:"mysqlId,omitempty" tf:"mysql_id,omitempty"`

	// +kubebuilder:validation:Optional
	MySQLIDRef *v1.Reference `json:"mySqlidRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	MySQLIDSelector *v1.Selector `json:"mySqlidSelector,omitempty" tf:"-"`

	// Table privileges list.
	// +kubebuilder:validation:Optional
	Table []TableParameters `json:"table,omitempty" tf:"table,omitempty"`
}

type TableObservation struct {
}

type TableParameters struct {

	// Database name.
	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// Table privilege.available values for Privileges:SELECT,INSERT,UPDATE,DELETE,CREATE,DROP,REFERENCES,INDEX,ALTER,CREATE VIEW,SHOW VIEW,TRIGGER.
	// +kubebuilder:validation:Required
	Privileges []*string `json:"privileges" tf:"privileges,omitempty"`

	// Table name.
	// +kubebuilder:validation:Required
	TableName *string `json:"tableName" tf:"table_name,omitempty"`
}

// PrivilegeSpec defines the desired state of Privilege
type PrivilegeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivilegeParameters `json:"forProvider"`
}

// PrivilegeStatus defines the observed state of Privilege.
type PrivilegeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivilegeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Privilege is the Schema for the Privileges API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Privilege struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivilegeSpec   `json:"spec"`
	Status            PrivilegeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivilegeList contains a list of Privileges
type PrivilegeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Privilege `json:"items"`
}

// Repository type metadata.
var (
	Privilege_Kind             = "Privilege"
	Privilege_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Privilege_Kind}.String()
	Privilege_KindAPIVersion   = Privilege_Kind + "." + CRDGroupVersion.String()
	Privilege_GroupVersionKind = CRDGroupVersion.WithKind(Privilege_Kind)
)

func init() {
	SchemeBuilder.Register(&Privilege{}, &PrivilegeList{})
}
