#!/usr/bin/bash

# DEEP_PROCESS 项目源码根目录
DEEP_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..

# 生成文件存放目录
LOCAL_OUTOUT_ROOT="${DEEP_ROOT}/${OUT_DIR:-_output}"

# 设置统一的密码，方便记忆
readonly PASSWORD=${PASSWORD:-'deep'}

# Linux系统 going 用户
readonly LINUX_USERNAME=${LINUX_USERNAME:-deep-process}
# Linux root & going 用户密码
readonly LINUX_PASSWORD=${LINUX_PASSWORD:-${PASSWORD}}

# 设置安装目录
readonly INSTALL_DIR=${INSTALL_DIR:-/tmp/installation}
mkdir -p ${INSTALL_DIR}
readonly ENV_FILE=${DEEP_ROOT}/scripts/install/environment.sh

# MariaDB 配置信息
readonly MARIADB_ADMIN_USERNAME=${MARIADB_ADMIN_USERNAME:-root} # MariaDB root 用户
readonly MARIADB_ADMIN_PASSWORD=${MARIADB_ADMIN_PASSWORD:-${PASSWORD}} # MariaDB root 用户密码
readonly MARIADB_HOST=${MARIADB_HOST:-127.0.0.1:3306} # MariaDB 主机地址
readonly MARIADB_DATABASE=${MARIADB_DATABASE:-deep} # MariaDB deep 应用使用的数据库名
readonly MARIADB_USERNAME=${MARIADB_USERNAME:-deep} # deep 数据库用户名
readonly MARIADB_PASSWORD=${MARIADB_PASSWORD:-${PASSWORD}} # deep 数据库密码

# deep 配置
readonly DEEP_DATA_DIR=${DEEP_DATA_DIR:-/data/deep} # deep 各组件数据目录
readonly DEEP_INSTALL_DIR=${DEEP_INSTALL_DIR:-/opt/deep} # deep 安装文件存放目录
readonly DEEP_CONFIG_DIR=${DEEP_CONFIG_DIR:-/etc/deep} # deep 配置文件存放目录
readonly DEEP_LOG_DIR=${DEEP_LOG_DIR:-/var/log/deep} # deep 日志文件存放目录

# deep-apiserver 配置
readonly DEEP_APISERVER_HOST=${DEEP_APISERVER_HOST:-127.0.0.1} # deep-apiserver 部署机器 IP 地址
readonly DEEP_APISERVER_GRPC_BIND_ADDRESS=${DEEP_APISERVER_GRPC_BIND_ADDRESS:-0.0.0.0}
readonly DEEP_APISERVER_GRPC_BIND_PORT=${DEEP_APISERVER_GRPC_BIND_PORT:-8081}
#readonly DEEP_APISERVER_INSECURE_BIND_ADDRESS=${DEEP_APISERVER_INSECURE_BIND_ADDRESS:-127.0.0.1}
#readonly DEEP_APISERVER_INSECURE_BIND_PORT=${DEEP_APISERVER_INSECURE_BIND_PORT:-8080}
#readonly DEEP_APISERVER_SECURE_BIND_ADDRESS=${DEEP_APISERVER_SECURE_BIND_ADDRESS:-0.0.0.0}
#readonly DEEP_APISERVER_SECURE_BIND_PORT=${DEEP_APISERVER_SECURE_BIND_PORT:-8443}
#readonly DEEP_APISERVER_SECURE_TLS_CERT_KEY_CERT_FILE=${DEEP_APISERVER_SECURE_TLS_CERT_KEY_CERT_FILE:-${DEEP_CONFIG_DIR}/cert/iam-apiserver.pem}
#readonly DEEP_APISERVER_SECURE_TLS_CERT_KEY_PRIVATE_KEY_FILE=${DEEP_APISERVER_SECURE_TLS_CERT_KEY_PRIVATE_KEY_FILE:-${DEEP_CONFIG_DIR}/cert/iam-apiserver-key.pem}
