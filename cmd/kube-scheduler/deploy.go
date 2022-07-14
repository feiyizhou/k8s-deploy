package kube_scheduler

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/config"
	"k8s-deploy/utils"
)

var DeploySchedulerCmd = &cobra.Command{
	Use:   "kube-scheduler",
	Short: "kube-scheduler",
	Long:  "kube-scheduler",
	Run: func(cmd *cobra.Command, args []string) {

		conf := config.GetKubernetesConf()

		certArgs := []string{common.GetCertShellPath(common.KubeScheduler)}
		utils.DoOrDieWithMsg(utils.ExecShell(certArgs, nil, nil, nil),
			"Failed to generate the cert of kube-scheduler")

		startArgs := []string{
			common.GetDeployShellPath(common.KubeScheduler),
			fmt.Sprintf("https://%s:6443", conf.KubeApiServer),
		}
		utils.DoOrDieWithMsg(utils.ExecShell(startArgs, nil, nil, nil),
			"Failed to start kube-scheduler service")
	},
}
