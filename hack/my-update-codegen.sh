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
OUTPUT_DIR=generated
# group-version such as foo:v1alpha1
GROUP=example
VERSION=v1
GROUP_VERSION=example:v1
OUTPUT_PKG=${OUTPUT_DIR}/${GROUP}

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}
rm -rf "${APIS_PKG}/${GROUP}" && mkdir -p "${APIS_PKG}/${GROUP}/${VERSION}" && cp -r "${APIS_PKG}/${VERSION}/" "${APIS_PKG}/${GROUP}/${VERSION}"

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
bash hack/generate-groups.sh "client,lister,informer" \
  ${MODULE}/${OUTPUT_PKG} ${MODULE}/${APIS_PKG} \
  ${GROUP_VERSION} \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --output-base "${SCRIPT_ROOT}"

mv "${MODULE}/${OUTPUT_DIR}" "${SCRIPT_ROOT}"
rm -r "${DOMAIN}"