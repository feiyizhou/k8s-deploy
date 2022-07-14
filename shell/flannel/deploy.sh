#!/bin/bash

Network=$1
TYPE=$2
KUBE_APISERVER=$3
KUBE_ROOT_CA_FILE="/opt/kubernetes/ssl/ca.pem"
HOST_NAME=$4
INTERFACE_NAME=$5

mkdir /etc/kube-flannel/ -p

cat > /etc/kube-flannel/net-conf.json <<EOF
{
    "Network": "${Network}",
    "Backend": {
        "Type": "${TYPE}",
        "DirectRouting": true
    }
}
EOF

cat <<EOF | kubectl apply -f -
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: flannel
rules:
- apiGroups: ['extensions']
  resources: ['podsecuritypolicies']
  verbs: ['use']
  resourceNames: ['psp.flannel.unprivileged']
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
- apiGroups:
  - ""
  resources:
  - nodes
  verbs:
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - nodes/status
  verbs:
  - patch
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: flannel
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: flannel
subjects:
- kind: ServiceAccount
  name: flannel
  namespace: kube-system
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: flannel
  namespace: kube-system
EOF

mkdir /opt/flannel

cd /opt/flannel || exit

kubectl config set-cluster kubernetes \
--kubeconfig=flannel.conf \
--embed-certs \
--server="${KUBE_APISERVER}" \
--certificate-authority=${KUBE_ROOT_CA_FILE}

kubectl config set-credentials flannel \
--kubeconfig=flannel.conf \
--token=$(kubectl get sa -n kube-system flannel -o jsonpath={.secrets[0].name} | xargs kubectl get secret -n kube-system  -o jsonpath={.data.token} | base64 -d)

kubectl config set-context kubernetes \
--kubeconfig=flannel.conf \
--user=flannel \
--cluster=kubernetes

kubectl config use-context kubernetes \
--kubeconfig=flannel.conf

cat > /usr/lib/systemd/system/flanneld.service <<EOF
[Unit]
Description=Flanneld
After=network.target
After=network-online.target
Wants=network-online.target
After=etcd.service

[Service]
Type=notify
Environment=NODE_NAME=${HOST_NAME}
ExecStart=/usr/local/bin/flanneld \\
  --iface=${INTERFACE_NAME} \\
  --ip-masq \\
  --kube-subnet-mgr=true \\
  --kubeconfig-file=/opt/flannel/flannel.conf

Restart=always
RestartSec=5
StartLimitInterval=0

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl start flanneld
systemctl enable flanneld