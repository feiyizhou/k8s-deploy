package flannel

import (
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/utils"
)

var ClearFlannelCmd = &cobra.Command{
	Use:   "flannel",
	Short: "flannel",
	Long:  "flannel",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DoOrDieWithMsg(utils.ExecShell([]string{
			common.GetClearShellPath(common.Flannel)}, nil, nil, nil),
			"Failed to config cni")
	},
}
