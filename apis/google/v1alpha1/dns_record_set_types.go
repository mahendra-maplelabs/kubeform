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

type DnsRecordSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsRecordSetSpec   `json:"spec,omitempty"`
	Status            DnsRecordSetStatus `json:"status,omitempty"`
}

type DnsRecordSetSpec struct {
	ManagedZone string   `json:"managed_zone"`
	Name        string   `json:"name"`
	Rrdatas     []string `json:"rrdatas"`
	Ttl         int      `json:"ttl"`
	Type        string   `json:"type"`
}

type DnsRecordSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsRecordSetList is a list of DnsRecordSets
type DnsRecordSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsRecordSet CRD objects
	Items []DnsRecordSet `json:"items,omitempty"`
}