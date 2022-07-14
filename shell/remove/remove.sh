#!/bin/bash

ifconfig cni0 down
ifconfig flannel.1 down

ip link del cni0
ip link del flannel.1


rm -rf ~/TLS/

rm -rf /opt/kubernetes/ssl
rm -rf /opt/kubernetes/cfg
rm -rf /opt/kubernetes/logs

rm -rf /opt/etcd/ssl
rm -rf /opt/etcd/cfg

rm -rf /var/lib/cni/cache
rm -rf /var/lib/cni/flannel
rm -rf /var/lib/cni/networks

rm -rf /root/.kube