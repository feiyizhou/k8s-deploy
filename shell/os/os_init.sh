#!/bin/bash

HOSTNAME="$1"

systemctl stop firewalld
systemctl disable firewalld

setenforce 0

swapoff -a
sed -ri 's/.*swap.*/#&/' /etc/fstab

if [[ "${HOSTNAME}" != "" ]]; then
  hostnamectl set-hostname "${HOSTNAME}"
fi

cat >> /etc/sysctl.d/k8s.conf << EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF

sysctl --system

mkdir -p ~/TLS/{etcd,k8s}

mkdir -p /opt/kubernetes/{ssl,cfg,logs}

mkdir -p /opt/etcd/{cfg,ssl}