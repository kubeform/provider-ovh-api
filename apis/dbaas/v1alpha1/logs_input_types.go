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

type LogsInput struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogsInputSpec   `json:"spec,omitempty"`
	Status            LogsInputStatus `json:"status,omitempty"`
}

type LogsInputSpecConfigurationFlowgger struct {
	// Type of format to decode
	LogFormat *string `json:"logFormat" tf:"log_format"`
	// Indicates how messages are delimited
	LogFraming *string `json:"logFraming" tf:"log_framing"`
}

type LogsInputSpecConfigurationLogstash struct {
	// The filter section of logstash.conf
	// +optional
	FilterSection *string `json:"filterSection,omitempty" tf:"filter_section"`
	// The filter section of logstash.conf
	InputSection *string `json:"inputSection" tf:"input_section"`
	// The list of customs Grok patterns
	// +optional
	PatternSection *string `json:"patternSection,omitempty" tf:"pattern_section"`
}

type LogsInputSpecConfiguration struct {
	// Flowgger configuration
	// +optional
	Flowgger *LogsInputSpecConfigurationFlowgger `json:"flowgger,omitempty" tf:"flowgger"`
	// Logstash configuration
	// +optional
	Logstash *LogsInputSpecConfigurationLogstash `json:"logstash,omitempty" tf:"logstash"`
}

type LogsInputSpec struct {
	State *LogsInputSpecResource `json:"state,omitempty" tf:"-"`

	Resource LogsInputSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type LogsInputSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// IP blocks
	// +optional
	AllowedNetworks []string `json:"allowedNetworks,omitempty" tf:"allowed_networks"`
	// Input configuration
	Configuration *LogsInputSpecConfiguration `json:"configuration" tf:"configuration"`
	// Input creation
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// Input description
	Description *string `json:"description" tf:"description"`
	// Input engine ID
	EngineID *string `json:"engineID" tf:"engine_id"`
	// Port
	// +optional
	ExposedPort *string `json:"exposedPort,omitempty" tf:"exposed_port"`
	// Hostname
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// Input ID
	// +optional
	InputID *string `json:"inputID,omitempty" tf:"input_id"`
	// Indicate if input need to be restarted
	// +optional
	IsRestartRequired *bool `json:"isRestartRequired,omitempty" tf:"is_restart_required"`
	// Number of instance running
	// +optional
	NbInstance *int64 `json:"nbInstance,omitempty" tf:"nb_instance"`
	// Input IP address
	// +optional
	PublicAddress *string `json:"publicAddress,omitempty" tf:"public_address"`
	ServiceName   *string `json:"serviceName" tf:"service_name"`
	// Input SSL certificate
	// +optional
	SslCertificate *string `json:"-" sensitive:"true" tf:"ssl_certificate"`
	// init: configuration required, pending: ready to start, running: available
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Associated Graylog stream
	StreamID *string `json:"streamID" tf:"stream_id"`
	// Input title
	Title *string `json:"title" tf:"title"`
	// Input last update
	// +optional
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at"`
}

type LogsInputStatus struct {
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

// LogsInputList is a list of LogsInputs
type LogsInputList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogsInput CRD objects
	Items []LogsInput `json:"items,omitempty"`
}
