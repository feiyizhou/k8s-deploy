package cmd

import (
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/utils"
)

var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove",
	Long:  "remove",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DoOrDieWithMsg(utils.ExecShell([]string{common.GetRemoveShellPath()},
			nil, nil, nil), "")
	},
}
