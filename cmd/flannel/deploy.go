package flannel

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/config"
	"k8s-deploy/utils"
)

var DeployFlannelCmd = &cobra.Command{
	Use:   "flannel",
	Short: "flannel",
	Long:  "flannel",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.GetFlannelConf()
		startArgs := []string{
			common.GetDeployShellPath(common.Flannel),
			conf.Network,
			conf.Type,
			fmt.Sprintf("https://%s:6443", conf.KubeApiServer),
			conf.HostnameOverride,
			conf.IfaceName,
		}
		utils.DoOrDieWithMsg(utils.ExecShell(startArgs, nil, nil, nil),
			"Failed to config cni")
	},
}
