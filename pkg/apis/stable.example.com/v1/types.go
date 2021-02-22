// +k8s:deepcopy-gen=package,register
// groupName=stable.example.com

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/code-generator"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interface=k8s.io/apimachinery/pkg/runtime.Object
type KubeApi struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec KubeApiSpec 	`json:"spec"`
}

type KubeApiSpec struct {
	Replicas 	*int32  		`json:"replicas"`
	HostUrl  	string			`json:"hostUrl"`
	ServiceType string  		`json:"serviceType"`
	Container 	ContainerSpec   `json:"container"`
}

type ContainerSpec struct {
	Image		    string		`json:"image"`
	ContainerPort	int			`json:"containerPort"`
}

// +k8s:deepcopy-gen:interfaces=k87s.io/apimachinery/pkg/runtime.Object
type KubeApiList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ListMeta   `json:"metadata"`

	Items []KubeApi   `json:"items"`
}
