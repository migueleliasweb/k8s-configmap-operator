/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Docs:
// 		https://book.kubebuilder.io/reference/generating-crd.html#validation
// 		https://book.kubebuilder.io/reference/markers/crd-validation.html

type DataSource struct {
	ConfigMap      DataSourceSelector `json:"configmap,omitempty"`
	CustomResource DataSourceSelector `json:"customresource,omitempty"`
}

type DataSourceSelector struct {
	LabelSelector    metav1.LabelSelector `json:"labelselector,omitempty"`
	NamespaceInclude []string             `json:"nsinclude,omitempty"`
	NamespaceExclude []string             `json:"nsexclude,omitempty"`
}

// +kubebuilder:validation:Enum=Webhook;RolloutRestart
type ReloadConfig string

// +kubebuilder:validation:Enum=Merge;Append;AppendFileMount
// TODO: consider Webhook?
type Strategy string

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DistributedConfigSpec defines the desired state of DistributedConfig
type DistributedConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Strategy string `json:"strategy"`

	// Foo is an example field of DistributedConfig. Edit distributedconfig_types.go to remove/update
	DataSource DataSource `json:"datasource"`

	ReloadConfig ReloadConfig `json:"reloadconfig"`
}

// DistributedConfigStatus defines the observed state of DistributedConfig
type DistributedConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DistributedConfig is the Schema for the distributedconfigs API
type DistributedConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DistributedConfigSpec   `json:"spec,omitempty"`
	Status DistributedConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DistributedConfigList contains a list of DistributedConfig
type DistributedConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DistributedConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DistributedConfig{}, &DistributedConfigList{})
}
