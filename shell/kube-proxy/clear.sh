#!/bin/bash

systemctl stop kube-proxy

rm -f /usr/lib/systemd/system/kube-proxy.service

rm -f /opt/kubernetes/cfg/kube-proxy*
rm -f /opt/kubernetes/ssl/kube-proxy*
rm -f ~/TLS/k8s/kube-proxy*

systemctl daemon-reload