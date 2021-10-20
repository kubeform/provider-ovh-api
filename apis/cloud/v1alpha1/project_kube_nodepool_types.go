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

type ProjectKubeNodepool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectKubeNodepoolSpec   `json:"spec,omitempty"`
	Status            ProjectKubeNodepoolStatus `json:"status,omitempty"`
}

type ProjectKubeNodepoolSpec struct {
	State *ProjectKubeNodepoolSpecResource `json:"state,omitempty" tf:"-"`

	Resource ProjectKubeNodepoolSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ProjectKubeNodepoolSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Enable anti affinity groups for nodes in the pool
	// +optional
	AntiAffinity *bool `json:"antiAffinity,omitempty" tf:"anti_affinity"`
	// Enable auto-scaling for the pool
	// +optional
	Autoscale *bool `json:"autoscale,omitempty" tf:"autoscale"`
	// Number of nodes which are actually ready in the pool
	// +optional
	AvailableNodes *int64 `json:"availableNodes,omitempty" tf:"available_nodes"`
	// Creation date
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// Number of nodes present in the pool
	// +optional
	CurrentNodes *int64 `json:"currentNodes,omitempty" tf:"current_nodes"`
	// Number of nodes you desire in the pool
	// +optional
	DesiredNodes *int64 `json:"desiredNodes,omitempty" tf:"desired_nodes"`
	// Flavor name
	// +optional
	Flavor *string `json:"flavor,omitempty" tf:"flavor"`
	// Flavor name
	FlavorName *string `json:"flavorName" tf:"flavor_name"`
	// Kube ID
	KubeID *string `json:"kubeID" tf:"kube_id"`
	// Number of nodes you desire in the pool
	// +optional
	MaxNodes *int64 `json:"maxNodes,omitempty" tf:"max_nodes"`
	// Number of nodes you desire in the pool
	// +optional
	MinNodes *int64 `json:"minNodes,omitempty" tf:"min_nodes"`
	// Enable monthly billing on all nodes in the pool
	// +optional
	MonthlyBilled *bool `json:"monthlyBilled,omitempty" tf:"monthly_billed"`
	// NodePool resource name
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Project id
	// +optional
	ProjectID *string `json:"projectID,omitempty" tf:"project_id"`
	// Service name
	ServiceName *string `json:"serviceName" tf:"service_name"`
	// Status describing the state between number of nodes wanted and available ones
	// +optional
	SizeStatus *string `json:"sizeStatus,omitempty" tf:"size_status"`
	// Current status
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Number of nodes with latest version installed in the pool
	// +optional
	UpToDateNodes *int64 `json:"upToDateNodes,omitempty" tf:"up_to_date_nodes"`
	// Last update date
	// +optional
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at"`
}

type ProjectKubeNodepoolStatus struct {
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

// ProjectKubeNodepoolList is a list of ProjectKubeNodepools
type ProjectKubeNodepoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProjectKubeNodepool CRD objects
	Items []ProjectKubeNodepool `json:"items,omitempty"`
}
