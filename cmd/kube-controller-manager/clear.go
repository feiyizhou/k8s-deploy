package kube_controller_manager

import (
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/utils"
)

var ClearControllerManagerCmd = &cobra.Command{
	Use:   "kube-controller-manager",
	Short: "kube-controller-manager",
	Long:  "kube-controller-manager",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DoOrDieWithMsg(utils.ExecShell([]string{common.GetClearShellPath(
			common.KubeControllerManager)}, nil, nil, nil),
			"Failed to remove kube-controller-manager configuration")
	},
}
