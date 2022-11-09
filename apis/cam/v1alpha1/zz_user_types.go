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

type UserObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	UID *float64 `json:"uid,omitempty" tf:"uid,omitempty"`

	Uin *float64 `json:"uin,omitempty" tf:"uin,omitempty"`
}

type UserParameters struct {

	// Indicate whether the CAM user can login to the web console or not.
	// +kubebuilder:validation:Optional
	ConsoleLogin *bool `json:"consoleLogin,omitempty" tf:"console_login,omitempty"`

	// Country code of the phone number, for example: '86'.
	// +kubebuilder:validation:Optional
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// Email of the CAM user.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists; otherwise the user will be deleted directly. Default is false.
	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// Name of the CAM user.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Indicate whether the CAM user need to reset the password when first logins.
	// +kubebuilder:validation:Optional
	NeedResetPassword *bool `json:"needResetPassword,omitempty" tf:"need_reset_password,omitempty"`

	// The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes uppercase letters, lowercase letters, numbers and special characters. Only required when `console_login` is true. If not set, a random password will be automatically generated.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Phone number of the CAM user.
	// +kubebuilder:validation:Optional
	PhoneNum *string `json:"phoneNum,omitempty" tf:"phone_num,omitempty"`

	// Remark of the CAM user.
	// +kubebuilder:validation:Optional
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// A list of tags used to associate different resources.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Indicate whether to generate the API secret key or not.
	// +kubebuilder:validation:Optional
	UseAPI *bool `json:"useApi,omitempty" tf:"use_api,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// User is the Schema for the Users API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec   `json:"spec"`
	Status            UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}