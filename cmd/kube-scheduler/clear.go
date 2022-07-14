package kube_scheduler

import (
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/utils"
)

var ClearSchedulerCmd = &cobra.Command{
	Use:   "kube-scheduler",
	Short: "kube-scheduler",
	Long:  "kube-scheduler",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DoOrDieWithMsg(utils.ExecShell([]string{common.GetClearShellPath(common.KubeScheduler)},
			nil, nil, nil), "Failed to remove kube-proxy configuration")
	},
}
