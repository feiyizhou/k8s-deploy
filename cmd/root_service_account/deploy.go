package root_service_account

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/config"
	"k8s-deploy/utils"
)

var DeployRootSACmd = &cobra.Command{
	Use:   "root-sa",
	Short: "root-sa",
	Long:  "root-sa",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.GetKubernetesConf()

		startArgs := []string{
			common.GetDeployShellPath(common.RootServiceAccount),
			fmt.Sprintf("https://%s:6443", conf.KubeApiServer),
		}
		utils.DoOrDieWithMsg(utils.ExecShell(startArgs, nil, nil, nil),
			"Failed to deploy root service account on kubernetes")
	},
}
