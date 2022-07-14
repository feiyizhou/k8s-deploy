package etcd

import (
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/utils"
)

var ClearEtcdCmd = &cobra.Command{
	Use:   "etcd",
	Short: "etcd",
	Long:  "etcd",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DoOrDieWithMsg(utils.ExecShell(
			[]string{common.GetClearShellPath(common.ETCD)},
			nil, nil, nil), "Failed to remove etcd service configuration")
	},
}
