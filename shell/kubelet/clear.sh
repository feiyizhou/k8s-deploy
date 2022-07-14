#!/bin/bash

systemctl stop kubelet

rm -f /usr/lib/systemd/system/kubelet.service

rm -f /opt/kubernetes/cfg/kubelet*
rm -f /opt/kubernetes/ssl/kubelet*
rm -f ~/TLS/k8s/kubelet*

systemctl daemon-reload