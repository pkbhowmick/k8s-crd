package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +groupName=stable.example.com
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeapis,singular=kubeapi,shortName=kapi,categories={}
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name=Age,type=date
// +kubebuilder:printcolumn:JSONPath=".spec.hostUrl",name=HostUrl,type=string
// +kubebuilder:subresources:status
// +kubebuilder:storageversion

type KubeApi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeApiSpec   `json:"spec"`
	Status KubeApiStatus `json:"status,omitempty"`
}

type KubeApiSpec struct {
	Version     string        `json:"version,omitempty"`
	Replicas    *int32        `json:"replicas,omitempty"`
	HostUrl     string        `json:"hostUrl"`
	ServiceType string        `json:"serviceType,omitempty"`
	Container   ContainerSpec `json:"container"`
}

type ContainerSpec struct {
	Image         string `json:"image"`
	ContainerPort int32  `json:"containerPort"`
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
