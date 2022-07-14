#!/bin/bash

systemctl stop etcd

rm -rf /var/lib/etcd

rm -f /usr/lib/systemd/system/etcd.service

rm -rf ~/TLS/etcd/*
rm -rf /opt/etcd/cfg/*
rm -rf /opt/etcd/ssl/*

systemctl daemon-reload