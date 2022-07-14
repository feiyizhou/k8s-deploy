# Deploy kubernetes cluster with shell

# 构建

```shell
git clone git@github.com:feiyizhou/k8s-deploy.git
cd k8s-deploy
go build -o sobeyadm
```

## 配置

### 1、ETCD配置

```shell
etcd:
  currentName: xxxx        # etcd-1    
  nodes:
    - name: xxxx           # etcd-1
      ip: xxx.xxx.xxx.xxx  # 172.16.200.235
```

**字段说明：**

- `curremtName`：当前`ETCD`节点的名称
- `node`：`ETCD`集群中每个节点的名称及对应的节点`IP`

### 2、Kubernetes配置

```shell
kubernetes:
  hostnameOverride:
    xxxx                  # master01
  advertiseAddress:
    xxx.xxx.xxx.xxx       # 172.16.200.235
  serverHosts:
    - xxx.xxx.xxx.1       # 172.16.200.1
    - xxx.xxx.xxx.xxx     # 172.16.200.235
  etcdHosts:
    - xxx.xxx.xxx.xxx     # 172.16.200.235
  serviceClusterIPRange:
    xxx.xxx.xxx.0/24      # 172.16.201.0/24
  clusterCIDR:
    xxx.xxx.0.0/16        # 172.16.0.0/16
  kubeApiServer:
    xxx.xxx.xxx.xxx       # 172.16.200.235
  clusterDNS:
    xxx.xxx.xxx.2         # 172.16.200.2
```

**字段说明：**

- `hostnameOverride`：`Kubernetes`集群中当前节点的名称
- `advertiseAddress`：`Kubernetes`集群中当前节点的`IP`
- `serverHosts`：`Kubernetes`集群中各个节点的`IP`，第一个元素是集群网段首`IP`
- `etcdHosts`：`ETCD`集群中各个节点的`IP`
- `serviceClusterIPRange`：`Kubernetes`集群中`Service`映射后的`IP`网段区间
- `clusterCIDR`：需要与网络插件中的网段配置一致
- `kubeApiServer`：集群中部署的第一个`master`的`IP`
- `clusterDNS`：集群`DNS`

### 3、Flannel配置

```shell
flannel:
  network:
    xxx.xxx.0.0/16       # 172.16.0.0/16
  type:
    vxlan                # 默认选择vxlan
  kubeApiServer:
    xxx.xxx.xxx.xxx      # 172.16.200.235
  hostNameOverride:
    xxxx                 # master01
  ifaceName:
    xxxx                 # enp1s0
```

**字段说明：**

- `network`：`Flannel`可进行分配的子网段，与`Kubernetes`配置中的`clusterCIDR`一致
- `type`：`Flannel`子网络的类型，默认填写`vxlan`
- `kubeApiServer`：集群中部署的第一个`master`的`IP`
- `hostnameOverride`：`Kubernetes`集群中当前节点的名称
- `ifaceName`：当前节点的物理网卡名称，可通过`ip a`查看

## 部署

### 1、系统

```shell
# 1、系统初始化
sobeyadm os-init
```

### 2、ETCD

```shell
# 1、部署ETCD
sobeyadm deploy etcd
```

### 3、Kubernetes

```shell
# 1、部署kube-apiserver
sobeyadm deploy kube-apiserver

# 2、部署kube-controller-manager
sobeyadm deploy kube-controller-manager

# 3、部署kube-scheduler
sobeyadm deploy kube-scheduler

# 4、创建root service account
sobeyadm deploy root-sa

# 5、部署kubelet并加入集群
sobeyadm deploy kubelet
sobeyadm deploy join-cluster

# 6、部署kube-proxy
sobeyadm deploy kube-proxy
```

### 4、Flannel

```shell
# 1、部署cni
sobeyadm deploy cni

# 2、部署flannel
sobeyadm deploy flannel
```

## 卸载

### 1、ETCD

```shell
# 1、卸载ETCD
sobeyadm clear etcd
```

### 2、Kubernetes

```shell
# 1、卸载kube-proxy
sobeyadm clear kube-proxy

# 2、卸载kubelet
sobeyadm clear kubelet

# 3、卸载kube-scheduler
sobeyadm clear kube-scheduler

# 4、卸载kube-controller-manager
sobeyadm clear kube-controller-manager

# 5、卸载kube-apiserver
sobeyadm clear kube-apiserver
```

### 3、配置

```shell
# 1、移除创建的公共配置文件夹及文件
sobeyadm remove
```

