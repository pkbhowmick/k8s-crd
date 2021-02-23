#!/bin/bash

vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/pkbhowmick/k8s-crd/pkg/client \
github.com/pkbhowmick/k8s-crd/pkg/apis \
stable.example.com:v1alpha1 \
--go-header-file "./hack/boilerplate.go.txt"

