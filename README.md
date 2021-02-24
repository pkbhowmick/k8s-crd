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
    - jsonPath: .spec.hostUrl
      name: HostUrl
      type: string
    name: v1alpha1
    schema:
      openAPIV3Schema:
        description: KubeApi defines an Rest Api Server
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
                description: Container spec of the Api server
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
              hostUrl:
                description: Host url of the ingress
                type: string
              replicas:
                default: 1
                description: Number of replicas of api server
                format: int32
                type: integer
              serviceType:
                default: ClusterIP
                description: Service type of the api server service
                type: string
              version:
                description: Version of kubeapi to be deployed
                type: string
            required:
            - container
            - hostUrl
            type: object
          status:
            properties:
              phase:
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []

```

### Sample custom resource manifest file 
```
apiVersion: stable.example.com/v1alpha1
kind: KubeApi
metadata:
  name: go-rest-api
spec:
  replicas: 2
  hostUrl: api.github.local
  serviceType: ClusterIP
  container:
    image: pkbhowmick/go-rest-api:2.0.1
    containerPort: 8080
```

