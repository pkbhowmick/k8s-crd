package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +groupName=stable.example.com
// +kubebuilder:validation:Optional
// +kubebuilder:validation:MaxItems=2
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeapis,singular:kubeapi,shortName=kapi,categories={api,all}
// +kubebuilder:subresources:status
// +kubebuilder:printcolumen:name="Version",type="string",JSONPath=".spec.version"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPATH=".spec.phase"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPATH=".metadata.creationTimestamp"
// +kubebuilder:storageversion

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KubeApi struct {
	// +kubebuilder:validation:Required
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec KubeApiSpec 	`json:"spec"`
	Status KubeApiStatus `json:"status"`
}

type KubeApiSpec struct {
	Version 	string 			`json:"version"`
	Replicas 	*int32 		    `json:"replicas,omitempty"`
	// +kubebuilder:validation:Required
	HostUrl  	string			`json:"hostUrl"`
	ServiceType string  		`json:"serviceType"`
	// +kubebuilder:validation:Required
	Container 	ContainerSpec   `json:"container"`
}

type ContainerSpec struct {
	// +kubebuilder:validation:Required
	Image		    string		    `json:"image"`
	// +kubebuilder:validation:Required
	ContainerPort	int32			`json:"containerPort"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KubeApiList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ListMeta   `json:"metadata"`

	Items []KubeApi   `json:"items"`
}

type KubeApiStatus struct {
	Phase string `json:"phase,omitempty"`
}

