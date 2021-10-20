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

type Service struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec,omitempty"`
	Status            ServiceStatus `json:"status,omitempty"`
}

type ServiceSpecOrderDetails struct {
	// description
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// expiration date
	// +optional
	Domain *string `json:"domain,omitempty" tf:"domain"`
	// order detail id
	// +optional
	OrderDetailID *int64 `json:"orderDetailID,omitempty" tf:"order_detail_id"`
	// quantity
	// +optional
	Quantity *string `json:"quantity,omitempty" tf:"quantity"`
}

type ServiceSpecOrder struct {
	// date
	// +optional
	Date *string `json:"date,omitempty" tf:"date"`
	// Information about a Bill entry
	// +optional
	Details []ServiceSpecOrderDetails `json:"details,omitempty" tf:"details"`
	// expiration date
	// +optional
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date"`
	// order id
	// +optional
	OrderID *int64 `json:"orderID,omitempty" tf:"order_id"`
}

type ServiceSpecPlanConfiguration struct {
	// Identifier of the resource
	Label *string `json:"label" tf:"label"`
	// Path to the resource in API.OVH.COM
	Value *string `json:"value" tf:"value"`
}

type ServiceSpecPlan struct {
	// Catalog name
	// +optional
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name"`
	// Representation of a configuration item for personalizing product
	// +optional
	Configuration []ServiceSpecPlanConfiguration `json:"configuration,omitempty" tf:"configuration"`
	// duration
	Duration *string `json:"duration" tf:"duration"`
	// Plan code
	PlanCode *string `json:"planCode" tf:"plan_code"`
	// Pricing model identifier
	PricingMode *string `json:"pricingMode" tf:"pricing_mode"`
}

type ServiceSpecPlanOptionConfiguration struct {
	// Identifier of the resource
	Label *string `json:"label" tf:"label"`
	// Path to the resource in API.OVH.COM
	Value *string `json:"value" tf:"value"`
}

type ServiceSpecPlanOption struct {
	// Catalog name
	// +optional
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name"`
	// Representation of a configuration item for personalizing product
	// +optional
	Configuration []ServiceSpecPlanOptionConfiguration `json:"configuration,omitempty" tf:"configuration"`
	// duration
	Duration *string `json:"duration" tf:"duration"`
	// Plan code
	PlanCode *string `json:"planCode" tf:"plan_code"`
	// Pricing model identifier
	PricingMode *string `json:"pricingMode" tf:"pricing_mode"`
}

type ServiceSpecRoutedTo struct {
	// Service where ip is routed to
	// +optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name"`
}

type ServiceSpec struct {
	State *ServiceSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ServiceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CanBeTerminated *bool `json:"canBeTerminated,omitempty" tf:"can_be_terminated"`
	// +optional
	Country *string `json:"country,omitempty" tf:"country"`
	// Custom description on your ip
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Ip *string `json:"ip,omitempty" tf:"ip"`
	// Details about an Order
	// +optional
	Order []ServiceSpecOrder `json:"order,omitempty" tf:"order"`
	// +optional
	OrganisationID *string `json:"organisationID,omitempty" tf:"organisation_id"`
	// Ovh Subsidiary
	OvhSubsidiary *string `json:"ovhSubsidiary" tf:"ovh_subsidiary"`
	// Ovh payment mode
	PaymentMean *string `json:"paymentMean" tf:"payment_mean"`
	// Product Plan to order
	Plan *ServiceSpecPlan `json:"plan" tf:"plan"`
	// Product Plan to order
	// +optional
	PlanOption []ServiceSpecPlanOption `json:"planOption,omitempty" tf:"plan_option"`
	// Routage information
	// +optional
	RoutedTo []ServiceSpecRoutedTo `json:"routedTo,omitempty" tf:"routed_to"`
	// +optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name"`
	// Possible values for ip type
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ServiceStatus struct {
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

// ServiceList is a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Service CRD objects
	Items []Service `json:"items,omitempty"`
}
