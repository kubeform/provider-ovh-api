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

type HttpFarmServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HttpFarmServerSpec   `json:"spec,omitempty"`
	Status            HttpFarmServerStatus `json:"status,omitempty"`
}

type HttpFarmServerSpec struct {
	State *HttpFarmServerSpecResource `json:"state,omitempty" tf:"-"`

	Resource HttpFarmServerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type HttpFarmServerSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Address *string `json:"address" tf:"address"`
	// +optional
	Backup *bool `json:"backup,omitempty" tf:"backup"`
	// +optional
	Chain *string `json:"chain,omitempty" tf:"chain"`
	// +optional
	Cookie *string `json:"cookie,omitempty" tf:"cookie"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	FarmID      *int64  `json:"farmID" tf:"farm_id"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	Probe *bool `json:"probe,omitempty" tf:"probe"`
	// +optional
	ProxyProtocolVersion *string `json:"proxyProtocolVersion,omitempty" tf:"proxy_protocol_version"`
	ServiceName          *string `json:"serviceName" tf:"service_name"`
	// +optional
	Ssl    *bool   `json:"ssl,omitempty" tf:"ssl"`
	Status *string `json:"status" tf:"status"`
	// +optional
	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

type HttpFarmServerStatus struct {
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

// HttpFarmServerList is a list of HttpFarmServers
type HttpFarmServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HttpFarmServer CRD objects
	Items []HttpFarmServer `json:"items,omitempty"`
}
