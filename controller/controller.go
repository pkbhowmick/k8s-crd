package main

import (
	clientset "github.com/pkbhowmick/k8s-crd/pkg/client/clientset/versioned"
	"github.com/pkbhowmick/k8s-crd/pkg/client/clientset/versioned/scheme"
	informers "github.com/pkbhowmick/k8s-crd/pkg/client/informers/externalversions/stable.example.com/v1alpha1"
	listers "github.com/pkbhowmick/k8s-crd/pkg/client/listers/stable.example.com/v1alpha1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	appsinformers "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	appslisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

const (
	SuccessSynced         = "Synced"
	ErrResourceExists     = "ErrResourceExists"
	MessageResourceExists = "Resource %q already exists and is not managed by KubeApi"
	MessageResourceSynced = "KubeApi synced successfully"
)

type Controller struct {
	kubeclientset    kubernetes.Interface
	kubeapiclientset clientset.Interface

	deploymentLister appslisters.DeploymentLister
	deploymentSynced cache.InformerSynced
	kubeapiLister    listers.KubeApiLister
	kubeapiSynced    cache.InformerSynced

	workqueue workqueue.RateLimitingInterface
	recorder  record.EventRecorder
}

func NewController(
	kubeclientset kubernetes.Interface,
	kubeapiclientset clientset.Interface,
	deploymentInformer appsinformers.DeploymentInformer,
	kubeinformer informers.KubeApiInformer) *Controller {
	utilruntime.Must(samplescheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster()")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&corev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})
	controller := &Controller{
		kubeclientset:    kubeclientset,
		kubeapiclientset: kubeapiclientset,
		deploymentLister: deploymentLister,
		deploymentSynced: deploymentSynced,
		kubeapiLister:    kubeapiLister,
		kubeapiSynced:    kubeapiSynced,
		workqueue:        workqueue,
		recorder:         recorder,
	}
}
