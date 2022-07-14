package kube_apiserver

import (
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/utils"
)

var ClearApiServerCmd = &cobra.Command{
	Use:   "kube-apiserver",
	Short: "kube-apiserver",
	Long:  "kube-apiserver",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DoOrDieWithMsg(utils.ExecShell(
			[]string{common.GetClearShellPath(common.KubeApiServer)},
			nil, nil, nil), "Filed to remove kube-apiserver configuration")
	},
}
