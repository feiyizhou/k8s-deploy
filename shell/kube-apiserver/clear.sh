#!/bin/bash

systemctl stop kube-apiserver

rm -f /usr/lib/systemd/system/kube-apiserver.service

rm -f /opt/kubernetes/cfg/kube-apiserver*
rm -f /opt/kubernetes/ssl/kube-apiserver*
rm -f ~/TLS/k8s/kube-apiserver*

systemctl daemon-reload