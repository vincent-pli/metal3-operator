package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BaremetalSpec defines the desired state of Baremetal
type BaremetalSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Registry Registry `json:"registry,omitempty"`
}

// Registry defines image overrides of knative images.
// The default value is used as a default format to override for all knative deployments.
// The override values are specific to each knative deployment.
// +k8s:openapi-gen=true
type Registry struct {
	// A map of a container name or arg key to the full image location of the individual knative container.
	// +optional
	Override map[string]string `json:"override,omitempty"`
}

// BaremetalStatus defines the observed state of Baremetal
type BaremetalStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Baremetal is the Schema for the baremetals API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=baremetals,scope=Namespaced
type Baremetal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BaremetalSpec   `json:"spec,omitempty"`
	Status BaremetalStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BaremetalList contains a list of Baremetal
type BaremetalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Baremetal `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Baremetal{}, &BaremetalList{})
}
