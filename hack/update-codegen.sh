#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# corresponding to go mod init <module>
DOMAIN=example.org
MODULE=example.org/multi-clusters
# api package
APIS_PKG=api
# generated output package
OUTPUT_PKG=generated/app
# group-version such as foo:v1alpha1
GROUP_VERSION=app:v1

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..


# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
bash ./generate-groups.sh "client,lister,informer" \
    example.org/multi-clusters/core-api/app example.org/multi-clusters/api \
    app:v1 \
    --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
    --output-base ../

cd .. && mv core-api core-api-old && mv example.org/multi-clusters/core-api ./ && rm -rf example.org && rm -rf core-api-old


#  ${MODULE}/${OUTPUT_PKG} ${MODULE}/${APIS_PKG} \
#  ${GROUP_VERSION} \
#  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
#  --output-base "${SCRIPT_ROOT}"
