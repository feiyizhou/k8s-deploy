#!/bin/bash

systemctl stop kube-scheduler

rm -f /usr/lib/systemd/system/kube-scheduler.service

rm -f /opt/kubernetes/cfg/kube-scheduler*
rm -f /opt/kubernetes/ssl/kube-scheduler*
rm -f ~/TLS/k8s/kube-scheduler*

systemctl daemon-reload