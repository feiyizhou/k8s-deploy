#!/bin/bash

HOSTS=""
for  i in "$@"; do
    if [[ "${HOSTS}" == "" ]]; then
        HOSTS=\"$i\"
    else
        tmp=${HOSTS},\"$i\"
        HOSTS=$tmp
    fi
done

# 1、创建工作目录
cd ~/TLS/k8s || exit

# 2、创建根证书
cat > ca-config.json << EOF
{
  "signing": {
    "default": {
      "expiry": "87600h"
    },
    "profiles": {
      "kubernetes": {
         "expiry": "87600h",
         "usages": [
            "signing",
            "key encipherment",
            "server auth",
            "client auth"
        ]
      }
    }
  }
}
EOF

cat > ca-csr.json << EOF
{
    "CN": "kubernetes",
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "L": "Beijing",
            "ST": "Beijing",
            "O": "k8s",
            "OU": "System"
        }
    ]
}
EOF

# 3、生成证书：生成ca.pem和ca-key.pem文件
cfssl gencert -initca ca-csr.json | cfssljson -bare ca -

# 4、创建证书请求文件
cat > server-csr.json << EOF
{
    "CN": "kubernetes",
    "hosts": [
      "127.0.0.1",
      ${HOSTS},
      "kubernetes",
      "kubernetes.default",
      "kubernetes.default.svc",
      "kubernetes.default.svc.cluster",
      "kubernetes.default.svc.cluster.local"
    ],
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "L": "BeiJing",
            "ST": "BeiJing",
            "O": "k8s",
            "OU": "System"
        }
    ]
}
EOF

# 5、生成证书，生成server.pem和server-key.pem
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes server-csr.json | cfssljson -bare server

# 6、把刚才生成的证书拷贝到配置文件中的路径
cp ~/TLS/k8s/ca*pem ~/TLS/k8s/server*pem /opt/kubernetes/ssl/