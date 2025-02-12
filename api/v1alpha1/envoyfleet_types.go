/*
MIT License

Copyright (c) 2021 Kubeshop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EnvoyFleetSpec defines the desired state of EnvoyFleet
type EnvoyFleetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Service describes Envoy K8s service settings
	Service *ServiceConfig `json:"service"`

	// Envoy image tag
	Image string `json:"image"`
	// Node Selector is used to schedule the Envoy pod(s) to the specificly labeled nodes, optional
	// This is the map of "key: value" labels (e.g. "disktype": "ssd")
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// Affinity is used to schedule Envoy pod(s) to specific nodes, optional
	// +optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`
	// Tolerations allow pod to be scheduled to the nodes that has specific toleration labels, optional
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
	// Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request.
	// Value must be non-negative integer. The value zero indicates stop immediately via
	// the kill signal (no opportunity to shut down).
	// If this value is nil, the default grace period will be used instead.
	// The grace period is the duration in seconds after the processes running in the pod are sent
	// a termination signal and the time when the processes are forcibly halted with a kill signal.
	// Set this value longer than the expected cleanup time for your process.
	// Defaults to 30 seconds.
	// +optional
	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty"`
	// Additional Envoy Deployment annotations, optional
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// Resources allow to set CPU and Memory resource requests and limits, optional
	// +optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// Size field specifies the number of Envoy Pods being deployed. Optional, default value is 1.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default:=1
	Size *int32 `json:"size,omitempty"`

	// Access logging settings for the Envoy
	AccessLog *AccessLoggingConfig `json:"accesslog,omitempty"`
}

type ServiceConfig struct {

	// Kubernetes service type: NodePort, ClusterIP or LoadBalancer
	// +kubebuilder:validation:Enum=NodePort;ClusterIP;LoadBalancer
	Type corev1.ServiceType `json:"type"`

	// Kubernetes Service ports
	Ports []corev1.ServicePort `json:"ports"`

	// Service's annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Static ip address for the LoadBalancer type if available
	// +optional
	LoadBalancerIP string `json:"loadBalancerIP,omitempty"`
}

// AccessLoggingConfig defines the access logs Envoy logging settings
type AccessLoggingConfig struct {
	// Stdout logging format - text for unstructured and json for the structured type of logging
	// +kubebuilder:validation:Enum=json;text
	Format string `json:"format"`

	// Logging format template for the unstructured text type.
	// See https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage for the usage.
	// Uses Kusk Gateway defaults if not specified.
	// +optional
	TextTemplate string `json:"text_template,omitempty"`

	// Logging format template for the structured json type.
	// See https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage for the usage.
	// Uses Kusk Gateway defaults if not specified.
	// +optional
	JsonTemplate map[string]string `json:"json_template,omitempty"`
}

// EnvoyFleetStatus defines the observed state of EnvoyFleet
type EnvoyFleetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// State indicates Envoy Fleet state
	State string `json:"state,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="size",type="integer",JSONPath=".spec.size"

// EnvoyFleet is the Schema for the envoyfleet API
type EnvoyFleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EnvoyFleetSpec   `json:"spec"`
	Status EnvoyFleetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EnvoyFleetList contains a list of EnvoyFleet
type EnvoyFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvoyFleet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EnvoyFleet{}, &EnvoyFleetList{})
}

// EnvoyFleetID is used to bind other CR configurations to the deployed Envoy Fleet
// Consists of EnvoyFleet CR name and namespace
type EnvoyFleetID struct {

	//+kubebuilder:validation:Pattern:="^[a-z0-9-]{1,62}$"
	// deployed Envoy Fleet CR name
	Name string `json:"name"`

	//+kubebuilder:validation:Pattern:="^[a-z0-9-]{1,62}$"
	// deployed Envoy Fleet CR namespace
	Namespace string `json:"namespace"`
}

func (e EnvoyFleetID) String() string {
	return string(e.Name + "." + e.Namespace)
}
