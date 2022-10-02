#!/usr/bin/env bash

docker build -t k8s-configmap-operator hack/kubebuilder-build/

docker run -it --rm -v ${PWD}/:/workdir -u ${UID}:${UID} --workdir /workdir --entrypoint bash k8s-configmap-operator