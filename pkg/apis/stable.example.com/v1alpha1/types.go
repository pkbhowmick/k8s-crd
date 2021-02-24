package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +groupName=stable.example.com
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeapis,singular=kubeapi,shortName=kapi,categories={}
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name=Age,type=date
// +kubebuilder:printcolumn:JSONPath=".spec.hostUrl",name=HostUrl,type=string
// +kubebuilder:subresources:status
// +kubebuilder:storageversion

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KubeApi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeApiSpec   `json:"spec"`
	Status KubeApiStatus `json:"status,omitempty"`
}

// KubeApiSpec Defines KubeApi Object spec
type KubeApiSpec struct {
	// +optional
	Version string `json:"version,omitempty"`

	// +kubebuilder:default=1
	// +optional
	Replicas *int32 `json:"replicas"`

	HostUrl string `json:"hostUrl"`

	// +kubebuilder:default=ClusterIP
	// +optional
	ServiceType string `json:"serviceType"`

	Container ContainerSpec `json:"container"`
}

type ContainerSpec struct {
	// Container image of the Api Server
	Image string `json:"image"`

	// Container port of Api Server
	ContainerPort int32 `json:"containerPort"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KubeApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []KubeApi `json:"items"`
}

type KubeApiStatus struct {
	Phase string `json:"phase,omitempty"`
}
