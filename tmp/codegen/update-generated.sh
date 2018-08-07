#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
deepcopy \
github.com/thatinfraguy/redeploy-operator/pkg/generated \
github.com/thatinfraguy/redeploy-operator/pkg/apis \
app:v1alpha1 \
--go-header-file "./tmp/codegen/boilerplate.go.txt"
