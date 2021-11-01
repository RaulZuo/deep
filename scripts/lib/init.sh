#!/usr/bin/env bash

set -o errexit
set +o nounset
set -o pipefail

# Unset CDPATH so that path interpolation can work correctly
unset CDPATH

# Default use go module
export GO111MODULE=on

# The root of the build/dist directory
DEEP_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"

source "${DEEP_ROOT}/scripts/lib/util.sh"
source "${DEEP_ROOT}/scripts/lib/logging.sh"
source "${DEEP_ROOT}/scripts/lib/color.sh"

deep::log::install_errexit

source "${DEEP_ROOT}/scripts/lib/version.sh"
source "${DEEP_ROOT}/scripts/lib/golang.sh"
