package common

import "fmt"

const (
	ETCD = "etcd"

	Kubernetes            = "kubernetes"
	KubeApiServer         = "kube-apiserver"
	KubeControllerManager = "kube-controller-manager"
	KubeProxy             = "kube-proxy"
	KubeScheduler         = "kube-scheduler"
	Kubelet               = "kubelet"

	RootServiceAccount = "root-service-account"

	CNI     = "cni"
	Flannel = "flannel"
)

const (
	ShellHomePath = "/opt/k8s-deploy/shell"

	OsInitShell = "/os/os_init.sh"

	RemoveShell = "/remove/remove.sh"

	CertShellPath         = ShellHomePath + "/%s/cert.sh"
	DeployShellPath       = ShellHomePath + "/%s/deploy.sh"
	SingleDeployShellPath = ShellHomePath + "/%s/single_deploy.sh"
	ClearShellPath        = ShellHomePath + "/%s/clear.sh"

	YamlConfigHomePath = "/opt/k8s-deploy/yaml"
	YamlConfigType     = "yaml"
)

func GetOsInitShellPath() string {
	return fmt.Sprintf("%s%s", ShellHomePath, OsInitShell)
}

func GetCertShellPath(typ string) string {
	return fmt.Sprintf(CertShellPath, typ)
}

func GetDeployShellPath(typ string) string {
	return fmt.Sprintf(DeployShellPath, typ)
}

func GetSingleDeployShellPath(typ string) string {
	return fmt.Sprintf(SingleDeployShellPath, typ)
}

func GetClearShellPath(typ string) string {
	return fmt.Sprintf(ClearShellPath, typ)
}

func GetRemoveShellPath() string {
	return fmt.Sprintf("%s%s", ShellHomePath, RemoveShell)
}
