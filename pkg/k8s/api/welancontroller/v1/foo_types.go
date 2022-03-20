/*
controller-gen  crd  paths=./pkg/k8s/api/welancontroller/v1/...  output:stdout

*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ---------- +genclient is for code-generator/clientset to generate clientset code
// ---------- +k8s:deepcopy-gen is for code-generator/deepcopy-gen to generate deepcopy code

// ------- +kubebuilder:resource:categories={welan},singular="welanpoint",path="welanpoints",scope="Namespaced",shortName={cep,welanep}

// +kubebuilder:resource:categories={welan},singular="foo",path="foos",scope="Namespaced",shortName={fo,foshort}
// +kubebuilder:printcolumn:JSONPath=".status.state",description="the state of cr",name="Status",type=string
// +kubebuilder:printcolumn:JSONPath=".spec.deploymentName",description="my name",name="name",type=string
// +kubebuilder:printcolumn:JSONPath=".spec.conflicted",description="ip conflicted",name="conflicted flag",type=boolean
// +kubebuilder:printcolumn:JSONPath=".spec.length",description="subnet length",name="length",type=integer
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type=date
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// this will show in for the type description: Foo is a specification for a Foo resource
type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FooSpec   `json:"spec"`
	Status FooStatus `json:"status"`
}

// FooSpec is the spec for a Foo resource
type FooSpec struct {

	// this will show in for the field description: deployment name
	DeploymentName string `json:"deploymentName"`

	Replicas *int32 `json:"replicas"`

	// this will show in for the field description: topic name
	//
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:MinLength=10
	// +kubebuilder:validation:Required
	Topic string `json:"topic,omitempty"`

	//  --------- for controller-gen to generate crd yaml

	//  this will show in for the field description: ipv4 address , not must-require the field
	//
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`
	IPv4 string `json:"ipv4,omitempty"`

	//  this will show in for the field description: whether ip is conflicted
	//
	// +kubebuilder:default=true
	Conflicted bool `json:"conflicted"`

	// this will show in for the field description: subnet length
	//
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16
	SubnetLength int64 `json:"length,omitempty"`
}

// FooStatus is the status for a Foo resource
type FooStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`

	// this will show in for the field description
	// validation for the value, who must be one of them
	//
	// +kubebuilder:validation:Enum=creating;pending;invalid
	State string `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type FooList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Foo `json:"items"`
}
