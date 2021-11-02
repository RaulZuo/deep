# Systemd 配置、安装和启动

- [Systemd 配置、安装和启动](#systemd-配置安装和启动)
    - [1. 前置操作（需要 root 权限）](#前置操作需要-root-权限)
    - [1. 创建 iam-apiserver systemd unit 模板文件](#创建-iam-apiserver-systemd-unit-模板文件)
    - [4. 复制 systemd unit 模板文件到 sysmted 配置目录(需要有root权限)](#复制-systemd-unit-模板文件到-sysmted-配置目录需要有root权限)
    - [5. 启动 systemd 服务](#启动-systemd-服务)

## 1. 前置操作（需要 root 权限）

1. 根据注释配置 `environment.sh`

2. 创建 data 目录

```
mkdir -p ${DEEP_DATA_DIR}/deep-apiserver
```

3. 创建 bin 目录，并将 `deep-apiserver` 可执行文件复制过去

```bash
source ./environment.sh
mkdir -p ${DEEP_INSTALL_DIR}/bin
cp deep-apiserver ${DEEP_INSTALL_DIR}/bin
```