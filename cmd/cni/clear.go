package cni

import (
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/utils"
)

var ClearCNICmd = &cobra.Command{
	Use:   "cni",
	Short: "cni",
	Long:  "cni",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DoOrDieWithMsg(utils.ExecShell([]string{
			common.GetClearShellPath(common.CNI)}, nil, nil, nil),
			"Failed to config cni")
	},
}
