package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	stablev1alpha1 "github.com/pkbhowmick/k8s-crd/pkg/apis/stable.example.com/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kubeapiClientset "github.com/pkbhowmick/k8s-crd/pkg/client/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func int32Ptr(i int32) *int32 {
	return &i
}

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubeapiClientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	kubeapiClient := clientset.StableV1alpha1().KubeApis("default")
	kubeapi := &stablev1alpha1.KubeApi{
		ObjectMeta: metav1.ObjectMeta{
			Name: "go-api-server",
		},
		Spec: stablev1alpha1.KubeApiSpec{
			Replicas:    int32Ptr(2),
			HostUrl:     "api.github.local",
			ServiceType: "NodePort",
			Container: stablev1alpha1.ContainerSpec{
				Image:         "pkbhowmick/go-rest-api:2.0.1",
				ContainerPort: 8080,
			},
		},
	}
	fmt.Println("Creating KubeApi Resource...")
	result, err := kubeapiClient.Create(context.TODO(), kubeapi, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created KubeApi object: %q\n", result.GetObjectMeta().GetName())
}
