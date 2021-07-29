/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type InstallationTemplatePartitionSchemePartition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstallationTemplatePartitionSchemePartitionSpec   `json:"spec,omitempty"`
	Status            InstallationTemplatePartitionSchemePartitionStatus `json:"status,omitempty"`
}

type InstallationTemplatePartitionSchemePartitionSpec struct {
	State *InstallationTemplatePartitionSchemePartitionSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstallationTemplatePartitionSchemePartitionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type InstallationTemplatePartitionSchemePartitionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Partition filesystem
	Filesystem *string `json:"filesystem" tf:"filesystem"`
	// partition mount point
	Mountpoint *string `json:"mountpoint" tf:"mountpoint"`
	// step or order. specifies the creation order of the partition on the disk
	Order *int64 `json:"order" tf:"order"`
	// raid partition type
	// +optional
	Raid *string `json:"raid,omitempty" tf:"raid"`
	// name of this partitioning scheme
	SchemeName *string `json:"schemeName" tf:"scheme_name"`
	// size of partition in MB, 0 => rest of the space
	Size *int64 `json:"size" tf:"size"`
	// Template name
	TemplateName *string `json:"templateName" tf:"template_name"`
	// partition type
	Type *string `json:"type" tf:"type"`
	// The volume name needed for proxmox distribution
	// +optional
	VolumeName *string `json:"volumeName,omitempty" tf:"volume_name"`
}

type InstallationTemplatePartitionSchemePartitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// InstallationTemplatePartitionSchemePartitionList is a list of InstallationTemplatePartitionSchemePartitions
type InstallationTemplatePartitionSchemePartitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InstallationTemplatePartitionSchemePartition CRD objects
	Items []InstallationTemplatePartitionSchemePartition `json:"items,omitempty"`
}
