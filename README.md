# KubeApi
KubeApi is a Kubernetes Custom Resource which runs REST API server in Kubernetes. 



### Available commands
```shell
$ make generate         // to build the deepcopy funcs, clientset, informers, listers
$ make manifests        // to build the manifest file for kubernetes custom resource definition
```

### Kubernetes Custom Resource Definition Manifest file
Manifest file generated from above command.
```
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: kubeapis.stable.example.com
spec:
  group: stable.example.com
  names:
    kind: KubeApi
    listKind: KubeApiList
    plural: kubeapis
    shortNames:
    - kapi
    singular: kubeapi
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    - jsonPath: .status.deploymentName
      name: Deployment
      type: string
    - jsonPath: .status.serviceName
      name: Service
      type: string
    - jsonPath: .status.replicas
      name: Replicas
      type: integer
    - jsonPath: .status.phase
      name: Status
      type: string
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: KubeApiSpec Defines KubeApi Object spec
            properties:
              container:
                properties:
                  containerPort:
                    description: Container port of Api Server
                    format: int32
                    type: integer
                  image:
                    description: Container image of the Api Server
                    type: string
                required:
                - containerPort
                - image
                type: object
              deploymentName:
                type: string
              replicas:
                default: 1
                format: int32
                type: integer
              serviceName:
                type: string
              serviceType:
                default: ClusterIP
                enum:
                - ClusterIP
                - NodePort
                type: string
              version:
                type: string
            required:
            - container
            - serviceType
            type: object
          status:
            properties:
              phase:
                type: string
              replicas:
                description: Conditions         []metav1.Condition `json:"conditions"`
                format: int32
                type: integer
            required:
            - phase
            - replicas
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []

```
 
### Kubebuilder markers
- ```// +kubebuilder:object:root=true``` tells the object generator that this type represents a kind.
- ```// +kubebuilder:subresource:status``` tells the object generator that we want a status subresource. 
- ```// +kubebuilder:object:generator=true``` is a package level marker.
- ```// +groupName=stable.example.com``` is a package level marker that represent the API group.

### Controller
- Controllers are the core of kubernetes, and of any operator. 
- It's controller job to ensure that, for any given object, the actual state of the world matches the desired state in the object.
- Each controller focuses on one root Kind, but may interact with other Kinds.

### References:
- [Kubebuilder book](https://book.kubebuilder.io/quick-start.html)
- [Code generation for CustomResources](https://www.openshift.com/blog/kubernetes-deep-dive-code-generation-customresources)
- [Kubernetes sample controller](https://github.com/kubernetes/sample-controller)
