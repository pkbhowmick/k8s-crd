package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"time"

	"k8s.io/client-go/kubernetes"

	"k8s.io/apimachinery/pkg/util/runtime"

	"github.com/pkbhowmick/k8s-crd/pkg/apis/stable.example.com/v1alpha1"

	kubeapiClientset "github.com/pkbhowmick/k8s-crd/pkg/client/clientset/versioned"

	"k8s.io/client-go/util/homedir"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/fields"

	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/apimachinery/pkg/util/wait"

	v1 "k8s.io/api/core/v1"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

// Controller demonstrates how to implement a controller with client-go
type Controller struct {
	// indexer is a thread safe store to store objects and their keys
	indexer cache.Indexer
	// queue represents a controller to maintain the events coming from kubernetes api server
	queue workqueue.RateLimitingInterface
	// informer pops from queue to deliver to custom controller
	informer cache.Controller

	kClient   kubernetes.Interface
	crdClient kubeapiClientset.Interface
}

// NewController creates a new controller.
func NewController(
	indexer cache.Indexer,
	queue workqueue.RateLimitingInterface,
	informer cache.Controller,
	kClient kubernetes.Interface,
	crdClient kubeapiClientset.Interface) *Controller {
	return &Controller{
		indexer:   indexer,
		queue:     queue,
		informer:  informer,
		kClient:   kClient,
		crdClient: crdClient,
	}
}

// processNextItem processes items from controller
func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	// actual logic will go in below function
	err := c.syncHandler(key.(string))
	c.handleErr(err, key)
	return true
}

// syncToStdout is a major function which contains actual business logic
func (c *Controller) syncHandler(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		klog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}
	if !exists {
		fmt.Printf("Kubeapi %s does not exist anymore\n", key)
	} else {
		deepCopyObj := obj.(*v1alpha1.KubeApi).DeepCopy()

		// Get deployment to check if it already exists
		depl, err := c.kClient.AppsV1().Deployments(v1.NamespaceDefault).Get(context.TODO(), deepCopyObj.Name, metav1.GetOptions{})

		if err == nil {
			// As err is nil, so deployment already exists, then update it with new replica count
			depl.Spec.Replicas = deepCopyObj.Spec.Replicas
			updatedDepl, err := c.kClient.AppsV1().Deployments(v1.NamespaceDefault).Update(context.TODO(), depl, metav1.UpdateOptions{})
			if err != nil {
				return err
			}
			fmt.Printf("Deployment %q created\n", updatedDepl.GetObjectMeta().GetName())
		} else {
			// As err is not nil, so deployment doesn't exist, then creating a new deployment

			deployment := CreateDeploymentObj(deepCopyObj)
			deployedObj, err := c.kClient.AppsV1().Deployments(v1.NamespaceDefault).Create(context.TODO(), deployment, metav1.CreateOptions{})
			if err != nil {
				return err
			}
			fmt.Printf("Deployment %q created\n", deployedObj.GetObjectMeta().GetName())
		}

		// Creating the service according to deployed object
		serviceObj := CreateServiceObj(deepCopyObj)
		svc, err := c.kClient.CoreV1().Services(v1.NamespaceDefault).Create(context.TODO(), serviceObj, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		fmt.Printf("Service %q created\n", svc.GetObjectMeta().GetName())
	}
	return nil
}

func CreateServiceObj(obj *v1alpha1.KubeApi) *v1.Service {
	serviceObj := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: obj.Name,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(obj, v1alpha1.SchemeGroupVersion.WithKind("KubeApi")),
			},
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceType(obj.Spec.ServiceType),
			Selector: map[string]string{
				"app": obj.Name,
			},
			Ports: []v1.ServicePort{
				{
					Protocol: v1.ProtocolTCP,
					Port:     obj.Spec.Container.ContainerPort,
				},
			},
		},
	}
	return serviceObj
}

func CreateDeploymentObj(obj *v1alpha1.KubeApi) *appsv1.Deployment {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: obj.Name,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(obj, v1alpha1.SchemeGroupVersion.WithKind("KubeApi")),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: obj.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": obj.Name,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": obj.Name,
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  obj.Name,
							Image: obj.Spec.Container.Image,
							Ports: []v1.ContainerPort{
								{
									ContainerPort: obj.Spec.Container.ContainerPort,
								},
							},
						},
					},
				},
			},
		},
	}
	return deployment
}

func (c *Controller) handleErr(err error, key interface{}) {
	if err == nil {
		c.queue.Forget(key)
		return
	}
	if c.queue.NumRequeues(key) < 2 {
		klog.Infof("Error syncing pod %v: %v", key, err)
		c.queue.AddRateLimited(key)
		return
	}
	c.queue.Forget(key)
	//runtime.HandleError(err)
	klog.Infof("Dropping pod %q out of the queue: %v", key, err)
}

func (c *Controller) Run(threadiness int, stopCh chan struct{}) {
	defer runtime.HandleCrash()
	defer c.queue.ShutDown()

	klog.Info("Starting KubeApi controller")

	go c.informer.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	klog.Info("Stopping Pod controller")
}

func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
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
		klog.Fatal(err)
	}

	clientset, err := kubeapiClientset.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}
	var kClientset kubernetes.Interface
	kClientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}
	kubeapiListWatcher := cache.NewListWatchFromClient(clientset.StableV1alpha1().RESTClient(), "kubeapis", v1.NamespaceDefault, fields.Everything())

	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	indexer, informer := cache.NewIndexerInformer(kubeapiListWatcher, &v1alpha1.KubeApi{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old interface{}, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(new)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})
	controller := NewController(indexer, queue, informer, kClientset, clientset)

	addErr := indexer.Add(&v1alpha1.KubeApi{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: v1.NamespaceDefault,
		},
	})

	if addErr != nil {
		panic(addErr)
	}

	stop := make(chan struct{})
	defer close(stop)
	go controller.Run(1, stop)

	select {}
}

func intPtr32(i int32) *int32 {
	return &i
}
