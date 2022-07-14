package kube_proxy

import (
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/utils"
)

var ClearProxyCmd = &cobra.Command{
	Use:   "kube-proxy",
	Short: "kube-proxy",
	Long:  "kube-proxy",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DoOrDieWithMsg(utils.ExecShell([]string{common.GetClearShellPath(common.KubeProxy)},
			nil, nil, nil), "Failed to remove kube-proxy configuration")
	},
}
