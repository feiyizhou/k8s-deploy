#!/bin/bash

systemctl stop flanneld

rm -rf /etc/kube-flannel

rm -rf /opt/flannel

rm -f /usr/lib/systemd/system/flanneld.service

rm -rf /run/flannel

systemctl daemon-reload