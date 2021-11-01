#!/usr/bin/env bash

# 本脚本功能：根据 scripts/environment.sh 配置，生成 DEEP_PROCESS 组件 DEEP_PROCESS 组件 YAML 配置问卷。
# 示例：./genconfig.sh scripts/environment.sh configs/deep-apiserver.yaml

env_file="$1"
template_file="$2"

DEEP_PROCESS_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

source "${DEEP_PROCESS_ROOT}/scripts/lib/init.sh"

if [ $# -ne 2 ];then
    deep::log::error "Usage: genconfig.sh scripts/environment.sh configs/deep-apiserver.yaml"
    exit 1
fi

source "${env_file}"

declare -a envs

set +u
for env in $(sed -n 's/^[^#].*${\(.*\)}.*/\1/p' ${template_file})
do
    if [ -z "$(eval echo \$${env})" ];then
        iam::log::error "environment variable '${env}' not set"
        missing=true
    fi
done

if [ "${missing}" ];then
    iam::log::error 'You may run `source scripts/environment.sh` to set these environment'
    exit 1
fi

eval "cat << EOF
$(cat ${template_file})
EOF"
