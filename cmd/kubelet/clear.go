package kubelet

import (
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/utils"
)

var ClearKubeletCmd = &cobra.Command{
	Use:   "kubelet",
	Short: "kubelet",
	Long:  "kubelet",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DoOrDieWithMsg(utils.ExecShell([]string{common.GetClearShellPath(common.Kubelet)},
			nil, nil, nil), "Failed to remove kube-proxy configuration")
	},
}
