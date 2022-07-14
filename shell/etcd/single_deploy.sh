#!/bin/bash

ETCD_NODE_NAME=$1

cat > /opt/etcd/cfg/etcd.conf << EOF
#[Member]
ETCD_NAME="${ETCD_NODE_NAME}"
ETCD_DATA_DIR="/var/lib/etcd/default.etcd"
ETCD_LISTEN_PEER_URLS="https://0.0.0.0:2380"
ETCD_LISTEN_CLIENT_URLS="https://0.0.0.0:2379"

#[Clustering]
ETCD_INITIAL_ADVERTISE_PEER_URLS="https://0.0.0.0:2380,https://127.0.0.1:2380"
ETCD_ADVERTISE_CLIENT_URLS="https://0.0.0.0:2379,https://127.0.0.1:2379"
ETCD_ENABLE_V2="true"
EOF

cat > /usr/lib/systemd/system/etcd.service << EOF
[Unit]
Description=Etcd Server
After=network.target
After=network-online.target
Wants=network-online.target

[Service]
Type=notify
EnvironmentFile=/opt/etcd/cfg/etcd.conf
ExecStart=/opt/etcd/bin/etcd \
--cert-file=/opt/etcd/ssl/server.pem \
--key-file=/opt/etcd/ssl/server-key.pem \
--peer-cert-file=/opt/etcd/ssl/server.pem \
--peer-key-file=/opt/etcd/ssl/server-key.pem \
--trusted-ca-file=/opt/etcd/ssl/ca.pem \
--peer-trusted-ca-file=/opt/etcd/ssl/ca.pem \
--logger=zap
Restart=on-failure
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl start etcd
systemctl enable etcd