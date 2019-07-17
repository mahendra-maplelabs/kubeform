package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SesDomainIdentityVerification struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesDomainIdentityVerificationSpec   `json:"spec,omitempty"`
	Status            SesDomainIdentityVerificationStatus `json:"status,omitempty"`
}

type SesDomainIdentityVerificationSpec struct {
	Domain string `json:"domain"`
}

type SesDomainIdentityVerificationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SesDomainIdentityVerificationList is a list of SesDomainIdentityVerifications
type SesDomainIdentityVerificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SesDomainIdentityVerification CRD objects
	Items []SesDomainIdentityVerification `json:"items,omitempty"`
}