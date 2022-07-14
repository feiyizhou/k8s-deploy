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
)

func NewClearCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "clear",
		Long:  "clear",
	}
	cmd.AddCommand(
		etcd.ClearEtcdCmd,
		cni.ClearCNICmd,
		flannel.ClearFlannelCmd,
		kubeApiServer.ClearApiServerCmd,
		kubeControllerManager.ClearControllerManagerCmd,
		kubeProxy.ClearProxyCmd,
		kubeScheduler.ClearSchedulerCmd,
		kubelet.ClearKubeletCmd,
	)
	return cmd
}
