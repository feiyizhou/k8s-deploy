package cmd

import (
	"github.com/spf13/cobra"
	"k8s-deploy/cmd/cni"
	"k8s-deploy/cmd/etcd"
	"k8s-deploy/cmd/flannel"
	kubeApiServer "k8s-deploy/cmd/kube-apiserver"
	kubeControllerManager "k8s-deploy/cmd/kube-controller-manager"
	kubeProxy "k8s-deploy/cmd/kube-proxy"
	kubeScheduler "k8s-deploy/cmd/kube-scheduler"
	"k8s-deploy/cmd/kubelet"
	rootServiceAccount "k8s-deploy/cmd/root_service_account"
)

func NewDeployRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "deploy",
		Long:  "deploy",
	}
	cmd.AddCommand(
		etcd.DeployEtcdCmd,
		cni.DeployCNICmd,
		flannel.DeployFlannelCmd,
		kubeApiServer.DeployApiServerCmd,
		kubeControllerManager.DeployControllerManagerCmd,
		kubeProxy.DeployProxyCmd,
		kubeScheduler.DeploySchedulerCmd,
		kubelet.DeployKubeletCmd,
		kubelet.JoinClusterCmd,
		rootServiceAccount.DeployRootSACmd,
	)
	return cmd
}
