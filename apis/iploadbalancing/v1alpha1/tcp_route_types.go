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

type TcpRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TcpRouteSpec   `json:"spec,omitempty"`
	Status            TcpRouteStatus `json:"status,omitempty"`
}

type TcpRouteSpecAction struct {
	// Farm ID for "farm" action type, empty for others
	// +optional
	Target *string `json:"target,omitempty" tf:"target"`
	// Action to trigger if all the rules of this route matches
	Type *string `json:"type" tf:"type"`
}

type TcpRouteSpecRules struct {
	// Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/route/availableRules" for a list of available rules
	// +optional
	Field *string `json:"field,omitempty" tf:"field"`
	// Matching operator. Not all operators are available for all fields. See "/availableRules"
	// +optional
	Match *string `json:"match,omitempty" tf:"match"`
	// Invert the matching operator effect
	// +optional
	Negate *bool `json:"negate,omitempty" tf:"negate"`
	// Value to match against this match. Interpretation if this field depends on the match and field
	// +optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern"`
	// Id of your rule
	// +optional
	RuleID *int64 `json:"ruleID,omitempty" tf:"rule_id"`
	// Name of sub-field, if applicable. This may be a Cookie or Header name for instance
	// +optional
	SubField *string `json:"subField,omitempty" tf:"sub_field"`
}

type TcpRouteSpec struct {
	State *TcpRouteSpecResource `json:"state,omitempty" tf:"-"`

	Resource TcpRouteSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TcpRouteSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Action triggered when all rules match
	Action *TcpRouteSpecAction `json:"action" tf:"action"`
	// Human readable name for your route, this field is for you
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// Route traffic for this frontend
	// +optional
	FrontendID *int64 `json:"frontendID,omitempty" tf:"frontend_id"`
	// List of rules to match to trigger action
	// +optional
	Rules []TcpRouteSpecRules `json:"rules,omitempty" tf:"rules"`
	// The internal name of your IP load balancing
	ServiceName *string `json:"serviceName" tf:"service_name"`
	// Route status. Routes in "ok" state are ready to operate
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated last. Only the first matching route will trigger an action
	// +optional
	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

type TcpRouteStatus struct {
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

// TcpRouteList is a list of TcpRoutes
type TcpRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TcpRoute CRD objects
	Items []TcpRoute `json:"items,omitempty"`
}
