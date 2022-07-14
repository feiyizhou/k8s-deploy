#!/bin/bash

systemctl stop kube-controller-manager

rm -f /usr/lib/systemd/system/kube-controller-manager.service

rm -f /opt/kubernetes/cfg/kube-controller-manager*
rm -f /opt/kubernetes/ssl/kube-controller-manager*
rm -f ~/TLS/k8s/kube-controller-manager*

systemctl daemon-reload