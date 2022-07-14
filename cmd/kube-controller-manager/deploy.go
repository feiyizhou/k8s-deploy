package kube_controller_manager

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/config"
	"k8s-deploy/utils"
)

type kubeControllerManagerService struct{}

func newKubeControllerManagerService() *kubeControllerManagerService {
	return &kubeControllerManagerService{}
}

var DeployControllerManagerCmd = &cobra.Command{
	Use:   "kube-controller-manager",
	Short: "kube-controller-manager",
	Long:  "kube-controller-manager",
	Run: func(cmd *cobra.Command, args []string) {

		conf := config.GetKubernetesConf()

		utils.DoOrDieWithMsg(utils.ExecShell(
			[]string{common.GetCertShellPath(common.KubeControllerManager)},
			nil, nil, nil), "Generate kube-controller-manager cert failed")

		utils.DoOrDieWithMsg(newKubeControllerManagerService().startService(conf),
			"Filed to start kube-controller-manager")
	},
}

func (kcs *kubeControllerManagerService) startService(
	conf *config.KubernetesConf) error {
	startArgs := []string{
		common.GetDeployShellPath(common.KubeControllerManager),
		conf.ClusterCIDR,
		conf.ServiceClusterIPRange,
		fmt.Sprintf("https://%s:6443", conf.KubeApiServer),
	}
	return utils.ExecShell(startArgs, nil, nil, nil)
}
