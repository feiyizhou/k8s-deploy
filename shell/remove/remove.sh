#!/bin/bash

ifconfig cni0 down
ifconfig flannel.1 down

ip link del cni0
ip link del flannel.1

yum remove cfssl.aarch64 -y
yum remove cni-plugins.aarch64 -y
yum remove etcd.aarch64 -y
yum remove flanneld.aarch64 -y
yum remove kubernetes.aarch64 -y
yum remove sobeyadm.aarch64 -y


rm -rf ~/TLS/

rm -rf /opt/kubernetes/

rm -rf /opt/k8s-deploy/

rm -rf /opt/etcd/

rm -rf /var/lib/cni/

rm -rf /var/lib/kubelet/

rm -rf /root/.kube